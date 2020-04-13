package httputil

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseContentByImg(valueHtml string) ([]string, error) {
	resultList := make([]string, 0)

	htmlDoc, err := goquery.NewDocumentFromReader(strings.NewReader(valueHtml))
	if err != nil {
		return nil, err
	}
	htmlDoc.Find("img").Each(func(i int, s *goquery.Selection) {
		imgSrc, ok := s.Attr("src")
		if ok {
			resultList = append(resultList, imgSrc)
		}
	})

	return resultList, nil
}
