package data

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/moyrne/delay-record/pkg/pingclient/internal/repo"
	"github.com/moyrne/delay-record/pkg/utils"
	"github.com/pkg/errors"
)

type PingCSV struct {
	file   *os.File
	writer *csv.Writer
}

func NewPingCSV(filename string) (*PingCSV, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// not close file

	writer := csv.NewWriter(file)

	return &PingCSV{
		file:   file,
		writer: writer,
	}, nil
}

func (p *PingCSV) Close() {
	p.writer.Flush()
	p.file.Close()
}

func (p *PingCSV) InsertRecord(ping *repo.Ping) error {
	if err := p.writer.Write([]string{
		ping.StartTime.String(),
		strconv.Itoa(int(ping.Start)),
		ping.EndTime.String(),
		strconv.Itoa(int(ping.End)),
		strconv.Itoa(int(ping.Delay)),
		utils.GetError(ping.Error),
	}); err != nil {
		return errors.WithStack(err)
	}

	p.writer.Flush()
	return nil
}
