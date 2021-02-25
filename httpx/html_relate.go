package httpx

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/treewei/blackfriday"
)

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	//src = re.ReplaceAllString(src, "\n")
	src = re.ReplaceAllString(src, "")
	////去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	//去除各种转义符
	re, _ = regexp.Compile("(&[a-z]{3,7};)")
	src = re.ReplaceAllString(src, "")

	return strings.TrimSpace(src)
}

func GetHtmlAbstract(content string) (abstract string) {
	nameRune := []rune(TrimHtml(content))

	if len(nameRune) > 300 {
		abstract = string(nameRune[0:300])
	} else {
		abstract = string(nameRune)
	}

	return
}

func GetHtmlContent(content string) (abstract string) {
	nameRune := []rune(TrimHtml(content))

	abstract = string(nameRune[0 : len(nameRune)*2/5])

	return
}

func GetContenAbstract(content string) (abstract string) {
	nameRune := []rune(content)
	if len(nameRune) > 140 {
		abstract = string(nameRune[0:140])
	} else {
		abstract = string(nameRune)
	}

	return
}

func GetURLFromContent(content string) string {
	re, _ := regexp.Compile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)

	//查找符合正则的第一个
	find := re.FindAll([]byte(content), -1)

	if len(find) > 0 {
		return string(find[0])
	}
	return content
}

func CutoffHTMLByPercentage(content string, percentage float64) (string, error) {
	pContent := TrimHtml(content)
	nameRune := []rune(pContent)
	target := int(float64(len(nameRune)) * percentage)
	find := strings.TrimSpace(string(nameRune[target : target+10]))
	var status bool
	//log.Println("find:", find)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return "", err
	}
	doc.Find("body").EachWithBreak(func(i int, body *goquery.Selection) bool {
		body.Children().Each(func(i int, p *goquery.Selection) {
			if status {
				//log.Println("status:", status)
				p.Remove()
			} else {
				status = strings.Contains(p.Text(), find)
				//log.Println(p.Html())
				//log.Println("get status:", status)
				//log.Println("==================== children ====================")
			}
		})
		return true
	})
	if h, err := doc.Html(); err != nil {
		return "", err
	} else {
		return h, nil
	}
}

func MarkdownRenderer(content string) string {
	//result := string(blackfriday.Run([]byte(content), blackfriday.WithExtensions(blackfriday.CommonExtensions)))
	result := string(blackfriday.Run([]byte(content)))
	return string(result)
}

func MarkdownAutoNewline(str string) string {
	re, _ := regexp.Compile("\\ *\\n")
	str = re.ReplaceAllLiteralString(str, "  \n")
	//m.Content=strings.Replace(m.Content, "\n", "  \n", -1)
	reg := regexp.MustCompile("```([\\s\\S]*)```")
	//返回str中第一个匹配reg的字符串
	data := reg.Find([]byte(str))
	strs := strings.Replace(string(data), "  \n", "\n", -1)
	re, _ = regexp.Compile("```([\\s\\S]*)```")
	return re.ReplaceAllLiteralString(str, strs)
}
