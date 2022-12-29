# kaisekisan

## Installation

Go
```console
$ go install github.com/ddddddO/kaisekisan/cmd/kaisekisan@latest
```

Homebrew
```console
$ brew install ddddddO/tap/kaisekisan
```

Scoop
```console
$ scoop bucket add ddddddO https://github.com/ddddddO/scoop-bucket.git
$ scoop install ddddddO/kaisekisan
```

## Usage

```console
$ cat test.csv
no,text,description
0,テキスト,テキストです
1,天気,晴れがいい
2,千葉,県名
3,0120441222,電話番号
4,越智大貴,人です

$ kaisekisan xxx.csv 2
Succeeded!

$ cat test.csv.out
no,text,classification,description
0,テキスト,名詞/一般/*/* (origin: テキスト),テキストです
1,天気,名詞/一般/*/* (origin: 天気),晴れがいい
2,千葉,名詞/固有名詞/地域/一般 (origin: 千葉),県名
3,0120441222,名詞/数/*/* (origin: 0120441222),電話番号
4,越智大貴,名詞/固有名詞/人名/姓 (origin: 越智),人です
```