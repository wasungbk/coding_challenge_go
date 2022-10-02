package main

import "time"

type Record struct {
	Timestamp time.Time
	Username  string
	Operation string
	Size      int
}

func NewRecord(record []string) Record {

	// parse time and size
	t, err := time.Parse(time.UnixDate, record[0])
	if err != nil {
		return nil, err
	}
	size, err := strconv.Atoi(record[3])
	if err != nil {
		return nil, err
	}

	rec := Record{t, record[1], record[2], size}
	return rec, nil
}
