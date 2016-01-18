package common

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
)

/*
 * スクレイピング
 */
func GetPage(url string, domtype string, id string) error {
	fmt.Println(domtype + ":" + id)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("error!!")
		return errors.New("指定されたurlへのアクセス中に例外が発生しました#" + url)
	}

	outData := make([]byte, 0, 10)

	doc.Find(domtype).Each(func(_ int, s *goquery.Selection) {
		// href固定でとりにいくのは問題
		url, _ := s.Attr("href")
		// url, _ := s.id(id)

		outData = append(outData, url...)
		outData = append(outData, "\t"...)
	})

	fmt.Println(string(outData))
	// この書き方だと複数回読みに行く設定がされていたらファイルを毎回上書きするので、最後の要素しか取れない
	ioutil.WriteFile("./scraping.txt", outData, os.ModePerm)
	return nil
}
