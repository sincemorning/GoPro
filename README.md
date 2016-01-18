設定した情報に基づくデータを宛先URLから引っ張ってきてローカルに保存する。  

### ファイル構成
Main.go
setting.json
common
  ┣ scraping.go
  ┗ settingReader.go

### 設定ファイル(setting.json)の記載ルール
```code
{
    "output":"C:\\GoPro\\filename.txt",
    "setting":[
        {"url":"http://qiita.com/advent-calendar/2013/","domtype":"a","id":"test"},
        {"url":"https://godoc.org/","domtype":"b","id":"test2"}
    ]
}
```
接続先の情報は"setting"要素の配下にある配列を取得する。  
配列は次の要素で構成される必要がある

|要素名|設定内容|
|:---|:---|
|url|取得先のURL|
|domtype|取得対象の要素|
|id|取得対象のDOMのID|

出力先の情報は"output"の設定値を参照してファイルを出力する。  
そのため、対象ディレクトリへの書き込み権限を保持するユーザー（もしくは管理者）で実行する事。
