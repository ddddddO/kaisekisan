package kaisekisan

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

func Kaiseki(r io.Reader, w io.Writer) error {
	csvReader := csv.NewReader(r)
	csvWriter := csv.NewWriter(w)
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return err
	}

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

		// TODO: 一旦決め打ちで2番目
		target := record[1]

		ret := ""
		tokens := t.Tokenize(target)
		for _, token := range tokens {
			features := strings.Join(token.Features(), ",")
			ret = fmt.Sprintf("%s\t%v", token.Surface, features)
			break
			// TODO: 一旦、解析結果の最初の行のみ
		}

		record = append(record, ret)
		csvWriter.Write(record)
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
