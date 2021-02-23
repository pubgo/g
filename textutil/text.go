package textutil

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"text/template"
	"unicode"
	"unicode/utf8"
)

// CamelCase
// splits the camelcase word and returns a list of words. It also
// supports digits. Both lower camel case and upper camel case are supported.
// For more info please check: http://en.wikipedia.org/wiki/CamelCase
//
// Examples
//
//   "" =>                     [""]
//   "lowercase" =>            ["lowercase"]
//   "Class" =>                ["Class"]
//   "MyClass" =>              ["My", "Class"]
//   "MyC" =>                  ["My", "C"]
//   "HTML" =>                 ["HTML"]
//   "PDFLoader" =>            ["PDF", "Loader"]
//   "AString" =>              ["A", "String"]
//   "SimpleXMLParser" =>      ["Simple", "XML", "Parser"]
//   "vimRPCPlugin" =>         ["vim", "RPC", "Plugin"]
//   "GL11Version" =>          ["GL", "11", "Version"]
//   "99Bottles" =>            ["99", "Bottles"]
//   "May5" =>                 ["May", "5"]
//   "BFG9000" =>              ["BFG", "9000"]
//   "BöseÜberraschung" =>     ["Böse", "Überraschung"]
//   "Two  spaces" =>          ["Two", "  ", "spaces"]
//   "BadUTF8\xe2\xe2\xa1" =>  ["BadUTF8\xe2\xe2\xa1"]
//
// Splitting rules
//
//  1) If string is not valid UTF-8, return it without splitting as
//     single item array.
//  2) Assign all unicode characters into one of 4 sets: lower case
//     letters, upper case letters, numbers, and all other characters.
//  3) Iterate through characters of string, introducing splits
//     between adjacent characters that belong to different sets.
//  4) Iterate through array of split strings, and if a given string
//     is upper case:
//       if subsequent string is lower case:
//         move last character of upper case string to beginning of
//         lower case string
func CamelCase(src string) (entries []string) {
	// don't split invalid utf8
	if !utf8.ValidString(src) {
		return []string{src}
	}
	entries = []string{}
	var runes [][]rune
	lastClass := 0
	class := 0
	// split into fields based on class of unicode character
	for _, r := range src {
		switch true {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			if len(runes) == 0 {
				runes = [][]rune{{r}}
			} else {
				runes[len(runes)-1] = append(runes[len(runes)-1], r)
			}
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}
	// handle upper case -> lower case sequences, e.g.
	// "PDFL", "oader" -> "PDF", "Loader"
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	// construct []string from results
	for _, s := range runes {
		if len(s) > 0 {
			entries = append(entries, string(s))
		}
	}
	return
}

// CJK is short for Chinese, Japanese and Korean.
//
// The constant cjk contains following Unicode blocks:
// 	\u2e80-\u2eff CJK Radicals Supplement
// 	\u2f00-\u2fdf Kangxi Radicals
// 	\u3040-\u309f Hiragana
// 	\u30a0-\u30ff Katakana
// 	\u3100-\u312f Bopomofo
// 	\u3200-\u32ff Enclosed CJK Letters and Months
// 	\u3400-\u4dbf CJK Unified Ideographs Extension A
// 	\u4e00-\u9fff CJK Unified Ideographs
// 	\uf900-\ufaff CJK Compatibility Ideographs
//
// For more information about Unicode blocks, see
// 	http://unicode-table.com/en/
const cjk = "" +
	"\u2e80-\u2eff" +
	"\u2f00-\u2fdf" +
	"\u3040-\u309f" +
	"\u30a0-\u30ff" +
	"\u3100-\u312f" +
	"\u3200-\u32ff" +
	"\u3400-\u4dbf" +
	"\u4e00-\u9fff" +
	"\uf900-\ufaff"

// ANS is short for Alphabets, Numbers
// and Symbols (`~!@#$%^&*()-_=+[]{}\|;:'",<.>/?).
//
// The constant ans doesn't contain all symbols above.
const ans = "A-Za-z0-9`\\$%\\^&\\*\\-=\\+\\\\|/\u00a1-\u00ff\u2022\u2027\u2150-\u218f"

var cjk_quote = regexp.MustCompile(re("([{{ .CJK }}])" + "([\"'])"))
var quote_cjk = regexp.MustCompile(re("([\"'])" + "([{{ .CJK }}])"))
var fix_quote = regexp.MustCompile(re("([\"'\\(\\[\\{<\u201c])" + "(\\s*)" + "(.+?)" + "(\\s*)" + "([\"'\\)\\]\\}>\u201d])"))
var fix_single_quote = regexp.MustCompile(re("([{{ .CJK }}])" + "( )" + "(')" + "([A-Za-z])"))

var cjk_hash = regexp.MustCompile(re("([{{ .CJK }}])" + "(#(\\S+))"))
var hash_cjk = regexp.MustCompile(re("((\\S+)#)" + "([{{ .CJK }}])"))

var cjk_operator_ans = regexp.MustCompile(re("([{{ .CJK }}])" + "([\\+\\-\\*/=&\\|<>])" + "([A-Za-z0-9])"))
var ans_operator_cjk = regexp.MustCompile(re("([A-Za-z0-9])" + "([\\+\\-\\*/=&\\|<>])" + "([{{ .CJK }}])"))

var cjk_bracket_cjk = regexp.MustCompile(re("([{{ .CJK }}])" + "([\\(\\[\\{<\u201c]+(.*?)[\\)\\]\\}>\u201d]+)" + "([{{ .CJK }}])"))
var cjk_bracket = regexp.MustCompile(re("([{{ .CJK }}])" + "([\\(\\[\\{<\u201c>])"))
var bracket_cjk = regexp.MustCompile(re("([\\)\\]\\}>\u201d<])" + "([{{ .CJK }}])"))
var fix_bracket = regexp.MustCompile(re("([\\(\\[\\{<\u201c]+)" + "(\\s*)" + "(.+?)" + "(\\s*)" + "([\\)\\]\\}>\u201d]+)"))

var fix_symbol = regexp.MustCompile(re("([{{ .CJK }}])" + "([~!;:,\\.\\?\u2026])" + "([A-Za-z0-9])"))

var cjk_ans = regexp.MustCompile(re("([{{ .CJK }}])([{{ .ANS }}@])"))
var ans_cjk = regexp.MustCompile(re("([{{ .ANS }}~!;:,\\.\\?\u2026])([{{ .CJK }}])"))

var context = map[string]string{
	"CJK": cjk,
	"ANS": ans,
}

func re(exp string) string {
	var buf bytes.Buffer

	var tmpl = template.New("pangu")
	tmpl, _ = tmpl.Parse(exp)
	tmpl.Execute(&buf, context)
	expr := buf.String()

	return expr
}

// SpacingText performs paranoid text spacing on text.
// It returns the processed text, with love.
func SpacingText(text string) string {
	if len(text) < 2 {
		return text
	}

	text = cjk_quote.ReplaceAllString(text, "$1 $2")
	text = quote_cjk.ReplaceAllString(text, "$1 $2")
	text = fix_quote.ReplaceAllString(text, "$1$3$5")
	text = fix_single_quote.ReplaceAllString(text, "$1$3$4")

	text = cjk_hash.ReplaceAllString(text, "$1 $2")
	text = hash_cjk.ReplaceAllString(text, "$1 $3")

	text = cjk_operator_ans.ReplaceAllString(text, "$1 $2 $3")
	text = ans_operator_cjk.ReplaceAllString(text, "$1 $2 $3")

	oldText := text
	newText := cjk_bracket_cjk.ReplaceAllString(oldText, "$1 $2 $4")
	text = newText
	if oldText == newText {
		text = cjk_bracket.ReplaceAllString(text, "$1 $2")
		text = bracket_cjk.ReplaceAllString(text, "$1 $2")
	}
	text = fix_bracket.ReplaceAllString(text, "$1$3$5")

	text = fix_symbol.ReplaceAllString(text, "$1$2 $3")

	text = cjk_ans.ReplaceAllString(text, "$1 $2")
	text = ans_cjk.ReplaceAllString(text, "$1 $2")

	return text
}

// SpacingFile reads the file named by filename, performs paranoid text
// spacing on its contents and writes the processed content to w.
// A successful call returns err == nil.
func SpacingFile(filename string, w io.Writer) (err error) {
	fr, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fr.Close()

	br := bufio.NewReader(fr)
	bw := bufio.NewWriter(w)

	for {
		line, err := br.ReadString('\n')
		if err == nil {
			fmt.Fprint(bw, SpacingText(line))
		} else {
			if err == io.EOF {
				fmt.Fprint(bw, SpacingText(line))
				break
			}
			return err
		}
	}
	defer bw.Flush()

	return nil
}
