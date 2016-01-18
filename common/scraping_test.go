package common

import (
	"testing"
)

/*
 * scraping.goのGetPageに対するテスト
 */
func TestGetPage(t *testing.T) {
    /*
	var tring int = 1
	var tring2 int = 2
	if tring != tring2 {
		t.Errorf("got %v want %v", tring, tring2)
	}
*/
	err := GetPage("https://www.google.co.jp/webhp?hl=ja", "input", "lst-ib")
	if err != nil {
		t.Errorf("TestGetPage#try GetPage()=[%v]", err)
	}
}
