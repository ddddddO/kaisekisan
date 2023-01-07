package kaisekisan

import (
	"strings"
)

type filter struct {
	origin   string
	pos      []string
	filtered string
}

func (f *filter) String() string {
	f = f.first().second()
	return f.filtered
}

// Kagome由来の分類にフィルター
func (f *filter) first() *filter {
	filtered := ""
	ippan := 0
	sei := 0
	kuni := 0
	for i := range f.pos {
		filtered = f.pos[i]
		if filtered == "一般" {
			ippan = i
		}
		if filtered == "姓" {
			sei = i
		}
		if filtered == "国" {
			kuni = i
		}
		if filtered == "*" {
			filtered = f.pos[i-1]
			break
		}
	}
	if 2 <= ippan {
		f.filtered = f.pos[ippan-1]
		return f
	}
	if 1 <= sei {
		f.filtered = f.pos[sei-1]
		return f
	}
	if 3 <= kuni {
		f.filtered = "地域"
		return f
	}

	f.filtered = filtered
	return f
}

// 独自の分類にフィルター
func (f *filter) second() *filter {
	if f.filtered == "数" {
		if f.isPhoneNumber() {
			f.filtered = "電話番号"
			return f
		}
		if f.isPostCode() {
			f.filtered = "郵便番号"
			return f
		}
		return f
	}

	if f.filtered == "一般" {
		if f.isID() {
			f.filtered = "ID"
			return f
		}
		return f
	}

	return f
}

// TODO: ちょっと広すぎるかも...
func (f *filter) isID() bool {
	return strings.Contains(f.origin, `-`)
}

// TODO:
func (f *filter) isPhoneNumber() bool {
	if f.isMobilePhoneNumber() {
		return true
	}
	if f.isShigaikyokuban() {
		return true
	}
	return false
}

var mobilePhonePrefixies = []string{"070", "080", "090"}

func (f *filter) isMobilePhoneNumber() bool {
	for i := range mobilePhonePrefixies {
		if strings.HasPrefix(f.origin, mobilePhonePrefixies[i]) {
			return true
		}
	}
	return false
}

// TODO: 03-xxxxなど
func (f *filter) isShigaikyokuban() bool {
	_ = f.origin
	return false
}

// TODO: 多分他のAPI使わせてもらうかな...レートリミットとかAPI KEY必要とか心配
func (f *filter) isPostCode() bool {
	_ = f.origin
	return false
}
