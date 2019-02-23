package db

type Reader struct {
	id          string
	entryStore  *EntryStore
	offsetStore *OffsetStore
	stop        bool
	newWrites   *Signaler
	offset      []byte
	distributor *Distributor
}

func (reader *Reader) getOffset() ([]byte, error) {
	return reader.offsetStore.Get([]byte(reader.id))
}

func (reader *Reader) setOffset(offset []byte) error {
	return reader.offsetStore.Set([]byte(reader.id), offset)
}

func newReader(id string, entryStore *EntryStore, offsetStore *OffsetStore, newWrites *Signaler) (*Reader, error) {
	reader := &Reader{
		id:          id,
		entryStore:  entryStore,
		offsetStore: offsetStore,
		newWrites:   newWrites,
		stop:        false,
		distributor: NewRoundRobin(),
	}
	offset, err := reader.getOffset()
	if err != nil {
		return nil, err
	}
	reader.offset = offset

	go reader.process()

	return reader, nil
}

func (reader *Reader) ID() string {
	return reader.id
}

func (reader *Reader) Offset() []byte {
	return reader.offset
}

func (reader *Reader) Notify(entry *Entry) {
	reader.newWrites.Notify()
}

func (reader *Reader) process() {
	handler := func(offset []byte, value []byte) error {
		message := NewMessage(NewEntry(offset, value), func() {
			if err := reader.setOffset(offset); err != nil {
				// TODO: hmmm
				panic(err)
			}
		})
		reader.distributor.Send(message)

		return nil
	}

	var err error
	for !reader.stop {
		reader.offset, err = reader.entryStore.Stream(reader.offset, handler)
		if err != nil {
			// TODO: if this is DB problem we need a failure state transition
			// but if it's consumer then ... warn? drop consumer?
			panic(err)
		}
		reader.newWrites.Wait()
	}
}

func (reader *Reader) Join(id string) *Receiver {
	return reader.distributor.Join(id)
}
