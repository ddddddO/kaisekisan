package kaisekisan

import (
	"testing"
)

func TestTokenizerKagome_Filter(t *testing.T) {
	tests := map[string]struct {
		inOrigin string
		inPos    []string
		want     string
	}{
		"組織":   {"ロケット団", []string{"名詞", "固有名詞", "組織", "*"}, "組織"},
		"一般":   {"天気", []string{"名詞", "一般", "*", "*"}, "一般"},
		"数":    {"111", []string{"名詞", "数", "*", "*"}, "数"},
		"携帯電話": {"08011111111", []string{"名詞", "数", "*", "*"}, "電話番号"},
		"ID":   {"LLL111-222", []string{"名詞", "一般", "*", "*"}, "ID"},
		"人名":   {"千葉 真一", []string{"名詞", "固有名詞", "人名", "姓"}, "人名"},
		"地域":   {"千葉県", []string{"名詞", "固有名詞", "地域", "一般"}, "地域"},
	}

	tk, err := newTokenizerKagome()
	if err != nil {
		t.Fatal(err)
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tk.filter(tt.inOrigin, tt.inPos)
			if got != tt.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.want)
			}
		})
	}
}
