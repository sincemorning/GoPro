package common

import (
	"encoding/json"
	"fmt"
)

/*
 設定ファイルのJsonの1要素分の定義
*/
type Settings struct {
	Url     string
	Domtype string
	Id      string
}

/*
 設定ファイルのJson全体の定義
*/
type SettingType struct {
	Setting []Settings
}

/*
 Jsonで記述された設定情報ファイルを読み取って、展開する
 第二返却値のerrorは必ずハンドルする
*/
func ReadSettingFile(data string) (SettingType, error) {
	fmt.Println(data)

	//
	var sett SettingType
	err := json.Unmarshal([]byte(data), &sett)
	if err != nil {
		// Jsonの読み取りが失敗している場合は何も入っていないであろうsettを返却するが、
		// 参照元ではきちんとerrorでハンドルすること
		fmt.Println(err)
		return sett, err
	}
	// Jsonの読み取りが成功している場合はエラーはnilで返却する
	return sett, nil
}
