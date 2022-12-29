package kaisekisan

import (
	"encoding/csv"
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
			record = append(record, "classification")
			if err := csvWriter.Write(record); err != nil {
				return err
			}

			isHeader = false
			continue
		}

		record = append(record, "ret")
		csvWriter.Write(record)
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
