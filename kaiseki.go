package kaisekisan

import (
	"encoding/csv"
	"errors"
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

		// ここでスペースを削除しているのは、例えば、「千葉 真一」を与えると、「千葉」を解析してしまい、分類が地名になるため。
		// 人名以外でもスペースが入る可能性はあるが、この際、スペースを削除することにした。
		target := strings.ReplaceAll(record[columnNumber-1], " ", "")
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
		// pos := strings.Join(token.POS(), "/")
		// return fmt.Sprintf("%v (origin: %s)", pos, token.Surface)
		return t.filter(in, token.POS())
		// TODO: 一旦、解析結果の最初の行のみ
	}
	return "不明"
}

func (t *tokenizerKagome) filter(origin string, ss []string) string {
	filtered := t.filterFirst(ss)
	if filtered == "数" {
		if t.isPhoneNumber(origin) {
			return "電話番号"
		}
		if t.isPostCode(origin) {
			return "郵便番号"
		}
		return filtered
	}

	if filtered == "一般" {
		if t.isID(origin) {
			return "ID"
		}
		return filtered
	}

	return filtered
}

func (*tokenizerKagome) filterFirst(ss []string) string {
	filtered := ""
	ippan := 0
	sei := 0
	for i := range ss {
		filtered = ss[i]
		if filtered == "一般" {
			ippan = i
		}
		if filtered == "姓" {
			sei = i
		}
		if filtered == "*" {
			filtered = ss[i-1]
			break
		}
	}
	if 2 <= ippan {
		return ss[ippan-1]
	}
	if 1 <= sei {
		return ss[sei-1]
	}
	return filtered
}

// 以下、filtered structみたいなのを作って、そっちにメソッドはやしたい。Stringメソッドも定義しておけば楽

// TODO: ちょっと広すぎるかも...
func (*tokenizerKagome) isID(origin string) bool {
	return strings.Contains(origin, `-`)
}

// TODO:
func (t *tokenizerKagome) isPhoneNumber(origin string) bool {
	if t.isMobilePhoneNumber(origin) {
		return true
	}
	if t.isShigaikyokuban(origin) {
		return true
	}
	return false
}

var mobilePhonePrefixies = []string{"070", "080", "090"}

func (*tokenizerKagome) isMobilePhoneNumber(origin string) bool {
	for i := range mobilePhonePrefixies {
		if strings.HasPrefix(origin, mobilePhonePrefixies[i]) {
			return true
		}
	}
	return false
}

// TODO: 03-xxxxなど
func (*tokenizerKagome) isShigaikyokuban(origin string) bool {
	return false
}

// TODO: 多分他のAPI使わせてもらうかな...レートリミットとかAPI KEY必要とか心配
func (*tokenizerKagome) isPostCode(origin string) bool {
	return false
}
