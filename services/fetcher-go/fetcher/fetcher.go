package fetcher

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func GetTitle(url string) string {
	res, _ := http.Get(url)
	defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)
	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)

	// 文字コード変換
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	// HTMLパース
	doc, _ := goquery.NewDocumentFromReader(reader)
	title := doc.Find("title").Text()
	return title
}
