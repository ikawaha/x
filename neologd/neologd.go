package neologd

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	ProlongedSoundMark = '\u30FC'
)

var neologdReplacer = strings.NewReplacer(
	"０", "0", "１", "1", "２", "2", "３", "3", "４", "4",
	"５", "5", "６", "6", "７", "7", "８", "8", "９", "9",

	"Ａ", "A", "Ｂ", "B", "Ｃ", "C", "Ｄ", "D", "Ｅ", "E",
	"Ｆ", "F", "Ｇ", "G", "Ｈ", "H", "Ｉ", "I", "Ｊ", "J",
	"Ｋ", "K", "Ｌ", "L", "Ｍ", "M", "Ｎ", "N", "Ｏ", "O",
	"Ｐ", "P", "Ｑ", "Q", "Ｒ", "R", "Ｓ", "S", "Ｔ", "T",
	"Ｕ", "U", "Ｖ", "V", "Ｗ", "W", "Ｘ", "X", "Ｙ", "Y",
	"Ｚ", "Z",

	"ａ", "a", "ｂ", "b", "ｃ", "c", "ｄ", "d", "ｅ", "e",
	"ｆ", "f", "ｇ", "g", "ｈ", "h", "ｉ", "i", "ｊ", "j",
	"ｋ", "k", "ｌ", "l", "ｍ", "m", "ｎ", "n", "ｏ", "o",
	"ｐ", "p", "ｑ", "q", "ｒ", "r", "ｓ", "s", "ｔ", "t",
	"ｕ", "u", "ｖ", "v", "ｗ", "w", "ｘ", "x", "ｙ", "y",
	"ｚ", "z",

	//small case
	"ｧ", "ァ", "ｨ", "ィ", "ｩ", "ゥ", "ｪ", "ェ", "ｫ", "ォ",
	"ｬ", "ャ", "ｭ", "ュ", "ｮ", "ョ", "ｯ", "ッ",

	"ｱ", "ア", "ｲ", "イ", "ｳ", "ウ", "ｴ", "エ", "ｵ", "オ",
	"ｶﾞ", "ガ", "ｷﾞ", "ギ", "ｸﾞ", "グ", "ｹﾞ", "ゲ", "ｺﾞ", "ゴ",
	"ｶ", "カ", "ｷ", "キ", "ｸ", "ク", "ｹ", "ケ", "ｺ", "コ",
	"ｻﾞ", "ザ", "ｼﾞ", "ジ", "ｽﾞ", "ズ", "ｾﾞ", "ゼ", "ｿﾞ", "ゾ",
	"ｻ", "サ", "ｼ", "シ", "ｽ", "ス", "ｾ", "セ", "ｿ", "ソ",
	"ﾀﾞ", "ダ", "ﾁﾞ", "ヂ", "ﾂﾞ", "ヅ", "ﾃﾞ", "デ", "ﾄﾞ", "ド",
	"ﾀ", "タ", "ﾁ", "チ", "ﾂ", "ツ", "ﾃ", "テ", "ﾄ", "ト",
	"ﾅ", "ナ", "ﾆ", "ニ", "ﾇ", "ヌ", "ﾈ", "ネ", "ﾉ", "ノ",
	"ﾊﾞ", "バ", "ﾋﾞ", "ビ", "ﾌﾞ", "ブ", "ﾍﾞ", "ベ", "ﾎﾞ", "ボ",
	"ﾊﾟ", "パ", "ﾋﾟ", "ピ", "ﾌﾟ", "プ", "ﾍﾟ", "ペ", "ﾎﾟ", "ポ",
	"ﾊ", "ハ", "ﾋ", "ヒ", "ﾌ", "フ", "ﾍ", "ヘ", "ﾎ", "ホ",
	"ﾏ", "マ", "ﾐ", "ミ", "ﾑ", "ム", "ﾒ", "メ", "ﾓ", "モ",
	"ﾔ", "ヤ", "ﾕ", "ユ", "ﾖ", "ヨ",
	"ﾗ", "ラ", "ﾘ", "リ", "ﾙ", "ル", "ﾚ", "レ", "ﾛ", "ロ",
	"ﾜ", "ワ", "ｦ", "ヲ", "ﾝ", "ン",

	// hyphen
	"\u02D7", "-", "\u058A", "-", "\u2010", "-", "\u2011", "-", "\u2012", "-",
	"\u2013", "-", "\u2043", "-", "\u207B", "-", "\u208B", "-", "\u2212", "-",

	// bar
	"\u2014", string(ProlongedSoundMark), // エムダッシュ
	"\u2015", string(ProlongedSoundMark), // ホリゾンタルバー
	"\u2500", string(ProlongedSoundMark), // 横細罫線
	"\u2501", string(ProlongedSoundMark), // 横太罫線
	"\uFE63", string(ProlongedSoundMark), // SMALL HYPHEN-MINUS
	"\uFF0D", string(ProlongedSoundMark), // 全角ハイフンマイナス
	"\uFF70", string(ProlongedSoundMark), // 半角長音記号

	// tilde
	"~", "", "\u223C", "", "\u223E", "", "\u301C", "", "\u3030", "", "\uFF5E", "",

	// zen -> han
	"！", "!", "”", `"`, "＃", "#", "＄", "$", "％", "%",
	"＆", "&", `’`, `'`, "（", "(", "）", ")", "＊", "*",
	"＋", "+", "，", ",", "−", "-", "．", ".", "／", "/",
	"：", ":", "；", ";", "＜", "<", "＞", ">", "？", "?",
	"＠", "@", "［", "[", "￥", "\u00A5", "］", "]", "＾", "^",
	"＿", "_", "｀", "`", "｛", "{", "｜", "|", "｝", "}",
	"　", " ",

	// han -> zen
	"｡", "。", "､", "、", "･", "・", "=", "＝", "｢", "「", "｣", "」",
)

type NeologdNormalizer struct {
	replacer *strings.Replacer
}

func NewNeologdNormalizer() *NeologdNormalizer {
	return &NeologdNormalizer{
		replacer: neologdReplacer,
	}
}

func (n NeologdNormalizer) Normalize(s string) string {
	return n.EliminateSpace(
		n.ShurinkProlongedSoundMark(
			n.CharReplace(s)))
}

func (n NeologdNormalizer) CharReplace(s string) string {
	return n.replacer.Replace(s)
}

func (n NeologdNormalizer) ShurinkProlongedSoundMark(s string) string {
	var b bytes.Buffer
	for p := 0; p < len(s); {
		c, w := utf8.DecodeRuneInString(s[p:])
		p += w
		b.WriteRune(c)
		if c != ProlongedSoundMark {
			continue
		}
		for p < len(s) {
			c0, w0 := utf8.DecodeRuneInString(s[p:])
			p += w0
			if c0 != ProlongedSoundMark {
				b.WriteRune(c0)
				break
			}
		}

	}
	return b.String()
}

func (n NeologdNormalizer) EliminateSpace(s string) string {
	var (
		b    bytes.Buffer
		prev rune
	)
	for p := 0; p < len(s); {
		c, w := utf8.DecodeRuneInString(s[p:])
		p += w
		if !unicode.IsSpace(c) {
			prev = c
			b.WriteRune(c)
			continue
		}
		for p < len(s) {
			c0, w0 := utf8.DecodeRuneInString(s[p:])
			p += w0
			if !unicode.IsSpace(c0) {
				if unicode.In(prev, unicode.Latin) && unicode.In(c0, unicode.Latin) {
					b.WriteRune(' ')
				}
				prev = c0
				b.WriteRune(c0)
				break
			}
		}

	}
	return b.String()
}
