package kaisekisan

import (
	"encoding/csv"
	"fmt"
	"io"
)

func Kaiseki(r io.Reader, w io.Writer) error {
	csvReader := csv.NewReader(r)
	csvWriter := csv.NewWriter(w)

	isHeader := true
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if isHeader {
			if err := csvWriter.Write(record); err != nil {
				return err
			}

			isHeader = false
			continue
		}

		ret := []string{}
		ret = append(ret, record...)
		csvWriter.Write(ret)
		// 一旦決め打ちで2列目
		fmt.Println(record[1])
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
