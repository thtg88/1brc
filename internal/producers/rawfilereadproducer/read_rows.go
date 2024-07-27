package rawfilereadproducer

import (
	"fmt"
	"io"
	"strings"
)

func (rfrp *RawFileReadProducer) ReadRows(carryover string) ([][]string, string, error) {
	if rfrp.Config.Debug {
		rfrp.Logger.Println("reading rows")
	}

	bytes := make([]byte, rfrp.Config.BufferedChannelSize)
	n, err := rfrp.IOReader.Read(bytes)
	if err != io.EOF && err != nil {
		return [][]string{}, "", fmt.Errorf("could not read from CSV file: %v", err)
	}
	if err == io.EOF || n == 0 {
		return [][]string{}, "", nil
	}

	data := fmt.Sprintf("%s%s", carryover, strings.TrimRight(string(bytes), "\x00"))
	if rfrp.Config.Debug {
		rfrp.Logger.Printf("data %v", data)
	}

	newLineIndex := strings.LastIndex(data, NewLineSeparator)
	if newLineIndex == -1 {
		return [][]string{}, data, nil
	}

	actualData := data[:newLineIndex]
	newCarryover := data[newLineIndex+1:]
	if rfrp.Config.Debug {
		rfrp.Logger.Printf("actualData %v", actualData)
		rfrp.Logger.Printf("newCarryover %v", newCarryover)
	}

	rawRows := strings.Split(actualData, NewLineSeparator)
	rows := make([][]string, len(rawRows))
	for idx, rawRow := range rawRows {
		row := strings.Split(rawRow, CommaSeparator)
		rows[idx] = row
	}

	return rows, newCarryover, nil
}
