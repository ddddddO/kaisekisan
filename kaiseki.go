package kaisekisan

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

func Kaiseki(r io.Reader, w io.Writer, columnNumber int) error {
	csvReader := csv.NewReader(r)
	csvWriter := csv.NewWriter(w)
	t, err := newTokenizerKagome()
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
			if len(record) < columnNumber {
				return errors.New("Specified column number is too large.")
			}

			ret := insert(record, columnNumber, "classification")
			if err := csvWriter.Write(ret); err != nil {
				return err
			}

			isHeader = false
			continue
		}

		target := record[columnNumber-1]
		result := t.Analyze(target)
		ret := insert(record, columnNumber, result)

		csvWriter.Write(ret)
	}

	csvWriter.Flush()
	return csvWriter.Error()
}

func insert(origin []string, columnNumber int, s string) []string {
	left := make([]string, columnNumber)
	right := make([]string, len(origin[columnNumber:]))
	copy(left, origin[:columnNumber])
	copy(right, origin[columnNumber:])
	ret := append(left, s)
	return append(ret, right...)
}

type tokenizerKagome struct {
	*tokenizer.Tokenizer
}

func newTokenizerKagome() (*tokenizerKagome, error) {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}

	return &tokenizerKagome{
		Tokenizer: t,
	}, nil
}

func (t *tokenizerKagome) Analyze(in string) string {
	tokens := t.Tokenize(in)
	for _, token := range tokens {
		features := strings.Join(token.Features()[0:4], "/")
		return fmt.Sprintf("%v (origin: %s)", features, token.Surface)
		// TODO: 一旦、解析結果の最初の行のみ
	}
	return ""
}
