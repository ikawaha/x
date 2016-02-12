package neologd

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	type testData struct {
		in, out string
	}
	data := []testData{
		{
			in:  "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ",
			out: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			in:  "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ",
			out: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			in:  "！”＃＄％＆’（）＊＋，−．／：；＜＞？＠［￥］＾＿｀｛｜｝",
			out: "!\"#$%&'()*+,-./:;<>?@[¥]^_`{|}",
		},
		{
			in:  "＝。、・「」",
			out: "＝。、・「」",
		},
		{
			in:  "ﾊﾝｶｸ",
			out: "ハンカク",
		},
		{
			in:  "o₋o",
			out: "o-o",
		},
		{
			in:  "majika━",
			out: "majikaー",
		},
		{
			in:  "わ〰い",
			out: "わい",
		},
		{
			in:  "スーパーーーー",
			out: "スーパー",
		},
		{
			in:  "!#",
			out: "!#",
		},
		{
			in:  "ゼンカク　スペース",
			out: "ゼンカクスペース",
		},
		{
			in:  "お             お",
			out: "おお",
		},
		{
			in:  "      おお",
			out: "おお",
		},
		{
			in:  "おお      ",
			out: "おお",
		},
		{
			in:  "検索 エンジン 自作 入門 を 買い ました!!!",
			out: "検索エンジン自作入門を買いました!!!",
		},
		{
			in:  "アルゴリズム C",
			out: "アルゴリズムC",
		},
		{
			in:  "　　　ＰＲＭＬ　　副　読　本　　　",
			out: "PRML副読本",
		},
		{
			in:  "Coding the Matrix",
			out: "Coding the Matrix",
		},
		{
			in:  "南アルプスの　天然水　Ｓｐａｒｋｉｎｇ　Ｌｅｍｏｎ　レモン一絞り",
			out: "南アルプスの天然水Sparking Lemonレモン一絞り",
		},
		{
			in:  "南アルプスの　天然水-　Ｓｐａｒｋｉｎｇ*　Ｌｅｍｏｎ+　レモン一絞り",
			out: "南アルプスの天然水- Sparking*Lemon+レモン一絞り",
		},
	}
	n := NewNeologdNormalizer()
	for _, d := range data {
		if x := n.Normalize(d.in); x != d.out {
			t.Errorf("got %v, expected %v", x, d.out)
		}
	}
}

func TestCharReplace(t *testing.T) {
	type testData struct {
		in, out string
	}
	data := []testData{
		{
			in:  "０１２３４５６７８９",
			out: "0123456789",
		},
		{
			in:  "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ",
			out: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			in:  "ｶﾞｷﾞｸﾞｹﾞｺﾞｶｷｸｹｺﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟﾊﾞﾋﾞﾌﾞﾍﾞﾎﾞ",
			out: "ガギグゲゴカキクケコパピプペポバビブベボ",
		},
		{
			in:  "ｧｨｩｪｫｬｭｮｯ",
			out: "ァィゥェォャュョッ",
		},
		{
			in:  "　",
			out: " ",
		},
	}
	n := NewNeologdNormalizer()
	for _, d := range data {
		if x := n.CharReplace(d.in); x != d.out {
			t.Errorf("got %v, expected %v", x, d.out)
		}
	}
}

func TestHyphen(t *testing.T) {
	l := "\u02D7\u058A\u2010\u2011\u2012\u2013\u2043\u207B\u208B\u2212"
	n := NewNeologdNormalizer()
	for _, c := range l {
		if n.CharReplace(string(c)) != "-" {
			t.Errorf("got 0x%X, expected 0x%X", c, '-')
		}
	}
}

func TestBar(t *testing.T) {
	l := "\u2014\u2015\u2500\uFE63\uFF0D\uFF70\u30FC"
	n := NewNeologdNormalizer()
	for _, c := range l {
		if n.CharReplace(string(c)) != "ー" {
			t.Errorf("got 0x%X, expected 0x%X", c, '\u30FC')
		}
	}
}

func TestShrinkProlongedSoundMark(t *testing.T) {
	type testData struct {
		in, out string
	}
	data := []testData{
		{
			in:  "スーーーーーーパーーーーーー",
			out: "スーパー",
		},
		{
			in:  "スーパーーーーーー",
			out: "スーパー",
		},
	}
	n := NewNeologdNormalizer()
	for _, d := range data {
		if x := n.ShurinkProlongedSoundMark(d.in); x != d.out {
			t.Errorf("got %v, expected %v", x, d.out)
		}
	}
}

func TestTilde(t *testing.T) {
	l := "~∼∾〜〰～"
	n := NewNeologdNormalizer()
	for _, c := range l {
		if n.CharReplace(string(c)) != "" {
			t.Errorf("got 0x%X, expected empty", c)
		}
	}
}

func TestEliminateSpace(t *testing.T) {
	type testData struct {
		in, out string
	}
	data := []testData{
		{
			in:  "  abc   ",
			out: "abc",
		},
		{
			in:  "検索 エンジン 自作 入門 を 買い ました!!!",
			out: "検索エンジン自作入門を買いました!!!",
		},
		{
			in:  "アルゴリズム C",
			out: "アルゴリズムC",
		},
		{
			in:  "　　　PRML　　副　読　本　　　",
			out: "PRML副読本",
		},
		{
			in:  "Coding the Matrix",
			out: "Coding the Matrix",
		},
		{
			in:  "南アルプスの　天然水　Sparking　Lemon　レモン一絞り",
			out: "南アルプスの天然水Sparking Lemonレモン一絞り",
		},
	}
	n := NewNeologdNormalizer()
	for _, d := range data {
		if x := n.EliminateSpace(d.in); x != d.out {
			t.Errorf("got %v, expected %v", x, d.out)
		}
	}
}
