package utils

import (
	"encoding/csv"
	"io"
	"os"
)

func WriteCsv(fileName string, rows [][]string) {
	nfs, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	defer nfs.Close()
	_, _ = nfs.Seek(0, io.SeekEnd)
	w := csv.NewWriter(nfs)
	_ = w.WriteAll(rows)
	w.Flush()
}

func ReadCsvContains(fileName string, column int, contains string) ([]string, error) {
	nfs, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer nfs.Close()
	r := csv.NewReader(nfs)
	for {
		row, err := r.Read()
		if err != nil {
			return nil, err
		}
		if row[column] == contains {
			return row, nil
		}
	}
}

func ReadCsvLastRow(fileName string) ([]string, error) {
	nfs, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer nfs.Close()
	r := csv.NewReader(nfs)
	var row []string
	for {
		_row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		row = _row
	}
	return row, nil
}
