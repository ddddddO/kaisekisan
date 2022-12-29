# kaisekisan

## Install

```console
go install github.com/ddddddO/kaisekisan/cmd/kaisekisan@latest
```

## Usage

```console
$ cat xxx.csv
no,text,description
0,テキスト,テキストです
1,天気,晴れがいい
2,千葉,県名
3,0120441222,電話番号
4,越智大貴,人です

$ kaisekisan xxx.csv
succeeded!

$ cat xxx.csv.out
no,text,description,classification
0,テキスト,テキストです,"テキスト       名詞,一般,*,*,*,*,テキスト,テキスト,テキスト"
1,天気,晴れがいい,"天気 名詞,一般,*,*,*,*,天気,テンキ,テンキ"
2,千葉,県名,"千葉       名詞,固有名詞,地域,一般,*,*,千葉,チバ,チバ"
3,0120441222,電話番号,"0120441222       名詞,数,*,*,*,*,*"
4,越智大貴,人です,"越智 名詞,固有名詞,人名,姓,*,*,越智,オチ,オチ"
```