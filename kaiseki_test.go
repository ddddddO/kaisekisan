package kaisekisan

import (
	"testing"
)

func TestTokenizerKagome_Filter(t *testing.T) {
	tests := map[string]struct {
		in   []string
		want string
	}{
		"組織": {[]string{"名詞", "固有名詞", "組織", "*"}, "組織"},
		"一般": {[]string{"名詞", "一般", "*", "*"}, "一般"},
		"数":  {[]string{"名詞", "数", "*", "*"}, "数"},
		"人名": {[]string{"名詞", "固有名詞", "人名", "姓"}, "人名"},
		"地域": {[]string{"名詞", "固有名詞", "地域", "一般"}, "地域"},
	}

	tk, err := newTokenizerKagome()
	if err != nil {
		t.Fatal(err)
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tk.filter(tt.in)
			if got != tt.want {
				t.Errorf("\ngot: \n%s\nwant: \n%s", got, tt.want)
			}
		})
	}
}
