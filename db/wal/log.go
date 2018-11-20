package wal

import "errors"

// Check check whether the log.checksum is same as crc
func (log *Log) Check(crc uint64) error {
	if log.Checksum == crc {
		return nil
	}
	log.Reset()
	return errors.New("WAL: crc mismatch")
}
