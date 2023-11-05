# CorrectCraft

# proto ファイルを使ったスキーマ定義ファイルの生成

## 新しいAPIを作成する場合の手順
1. proto ファイルを作成する
   1. server/proto 配下にディレクトリを作成する
      API の取り扱うトピックが異なる場合は、異なるディレクトリを作成するとわかりやすいかも。
      server/proto/<APIトピック>/v1/<APIトピック名>.proto
2. proto ファイルからスキーマ定義ファイルを生成する
下記コマンドを/server 配下で実行することで、/interface にconnect のスキーマ定義ファイルが生成される 
```sh
 buf generate
```


## bug.gen.yaml について
buf.gen.yaml は、buf が使用する設定ファイルです。
buf は、buf.gen.yaml を読み込み、その設定に従って、スキーマ定義ファイルを生成します。


