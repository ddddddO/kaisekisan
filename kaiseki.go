package kaisekisan

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

func Kaiseki(csvReader *csv.Reader, w io.Writer, columnNumber int) error {
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

		// ここでスペースを削除しているのは、例えば、「千葉 真一」を与えると、「千葉」を解析してしまい、分類が地名になるため。
		// 人名以外でもスペースが入る可能性はあるが、この際、スペースを削除することにした。
		target := strings.ReplaceAll(record[columnNumber-1], " ", "")
		result := t.analyze(target)
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

func (t *tokenizerKagome) analyze(in string) string {
	tokens := t.Tokenize(in)
	for _, token := range tokens {
		// pos := strings.Join(token.POS(), "/")
		// return fmt.Sprintf("%v (origin: %s)", pos, token.Surface)
		return t.filter(in, token.POS())
		// TODO: 一旦、解析結果の最初の行のみ
	}
	return "不明"
}

func (t *tokenizerKagome) filter(origin string, ss []string) string {
	f := &filter{origin: origin, pos: ss}
	return f.String()
}
