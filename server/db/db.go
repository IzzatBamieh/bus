package db

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/rs/xid"
)

type DB struct {
	workdir string
	logs    map[string]*Log
}

func NewDB(workdir string) (*DB, error) {
	db := &DB{
		workdir: workdir,
		logs:    map[string]*Log{},
	}

	if err := os.MkdirAll(workdir, 0700); err != nil {
		return nil, err
	}
	ids, err := readLogIDs(db.workdir)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		log, err := NewLog(id, path.Join(db.workdir, id))
		if err != nil {
			return nil, err
		}
		db.logs[id] = log
	}

	return db, nil
}

func (db *DB) Log(logID string) (*Log, error) {
	var err error
	var log *Log
	log, ok := db.logs[logID]
	if !ok {
		log, err = NewLog(logID, path.Join(db.workdir, logID))
		if err != nil {
			return nil, err
		}
		db.logs[logID] = log
	}

	return log, nil
}

func readLogIDs(path string) ([]string, error) {
	names := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			names = append(names, file.Name())
		}
	}

	return names, nil
}

type DBStats struct {
	Logs []*LogStats
}

type LogStats struct {
	ID      string
	Readers []*ReaderStats
}

type ReaderStats struct {
	ID     string
	Offset string
}

func (db *DB) Stats() (*DBStats, error) {
	stats := &DBStats{
		Logs: []*LogStats{},
	}
	for id := range db.logs {
		log, err := db.Log(id)
		if err != nil {
			return nil, err
		}
		logStats := &LogStats{
			ID:      log.ID(),
			Readers: []*ReaderStats{},
		}
		stats.Logs = append(stats.Logs, logStats)
		for id, reader := range log.Readers() {
			// TODO: need to create an "Offset" type
			x, err := xid.FromBytes(reader.Offset())
			if err != nil {
				return nil, err
			}
			readerStats := &ReaderStats{
				ID:     id,
				Offset: x.Time().String(),
			}
			logStats.Readers = append(logStats.Readers, readerStats)
		}
	}

	return stats, nil
}
