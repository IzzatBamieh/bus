package db

import (
	"path"
)

type Log struct {
	id        string
	workdir   string
	length    uint64
	readers   map[string]*Reader
	entries   *EntryStore
	offsets   *OffsetStore
	newWrites *Signaler
}

func NewLog(id string, workdir string) (*Log, error) {
	readers := map[string]*Reader{}
	entries, err := NewEntryStore(id, workdir)
	if err != nil {
		return nil, err
	}
	offsets, err := NewOffsetStore(id, path.Join(workdir, "offsets"))
	if err != nil {
		return nil, err
	}
	newWrites := NewSignaler(false)
	length := uint64(0)
	return &Log{
		id,
		workdir,
		length,
		readers,
		entries,
		offsets,
		newWrites,
	}, nil
}

func (log *Log) Length() uint64 {
	return log.length
}

func (log *Log) Reader(readerID string) (*Reader, error) {
	var err error
	var reader *Reader
	reader, ok := log.readers[readerID]
	if !ok {
		reader, err = newReader(readerID, log.entries, log.offsets, log.newWrites)
		if err != nil {
			return nil, err
		}
		log.readers[readerID] = reader
	}
	return reader, nil
}

func (log *Log) Writer() *Writer {
	return newWriter(func(value []byte) (*Entry, error) {
		entry, err := log.entries.Append(value)
		if err != nil {
			return nil, err
		}
		log.newWrites.Notify()
		return entry, nil
	})
}

type Writer struct {
	append func([]byte) (*Entry, error)
}

func newWriter(append func([]byte) (*Entry, error)) *Writer {
	return &Writer{
		append,
	}
}

func (writer *Writer) Append(value []byte) (*Entry, error) {
	return writer.append(value)
}

func (log *Log) Readers() map[string]*Reader {
	return log.readers
}

func (log *Log) ID() string {
	return log.id
}
