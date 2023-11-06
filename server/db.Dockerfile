FROM postgres:13.4-alpine

# 環境変数を設定
ENV LANG ja_JP.utf8

# postgresql-contribパッケージをインストール
RUN apk add --no-cache postgresql-contrib

# PostgreSQLを初期化する際にuuid-osspモジュールを有効にするスクリプトを追加
COPY ./initdb-uuid.sh /docker-entrypoint-initdb.d/
