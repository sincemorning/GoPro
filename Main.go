package main

import (
	"./common"
	"fmt"
	"io/ioutil"
)

func main() {

	data, err := ioutil.ReadFile("./setting.json")

	if err == nil {
		settingData, errors := common.ReadSettingFile(string(data))
		fmt.Println(settingData)

		if errors == nil {
			for n := 0; n < len(settingData.Setting); n++ {
				error := common.GetPage(settingData.Setting[n].Url, settingData.Setting[n].Domtype, settingData.Setting[n].Id)
				if error != nil {
					fmt.Print("ページのアクセスに失敗=")
					fmt.Println(error)
				}
			}
		} else {
			fmt.Print("設定の分解に失敗=")
			fmt.Println(errors)
		}
	} else {
		fmt.Print("設定の読み込みに失敗=")
		fmt.Println(err)
	}

}
