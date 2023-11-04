# Bun migrations server

This is a simple migrations server for [Bun](https://bun.uptrace.dev/).
サーバが使用するmodel のマイグレーションを行います。

To initialize migrations:
migration の順序を保存するbun_migration_locks, bun_migrations テーブルを作成する
```shell
go run . db init
```

To run migrations:
migration ファイルに記載された処理を実行します
```shell
BUNDEBUG=2 go run . db migrate
```

To rollback migrations:
migration ファイルに記載された処理を使用してDBの状態をロールバックします
1実行にあたり、1ファイル分の処理が巻き戻ります
```shell
go run . db rollback
```

To view status of migrations:
migration ファイルの実行状況を確認します

```shell
go run . db status
```

To create a Go migration:
Go 言語を使ったmigration ファイルを作成します
```shell
go run . db create_go go_migration_name
```

To create a SQL migration:
SQL を使ったmigration ファイルを作成します
```shell
go run . db create_sql sql_migration_name
```

To get help:

```shell
go run . db

NAME:
   bun db - database commands

USAGE:
   bun db command [command options] [arguments...]

COMMANDS:
   init        create migration tables
   migrate     migrate database
   rollback    rollback the last migration group
   unlock      unlock migrations
   create_go   create a Go migration
   create_sql  create a SQL migration
   help, h     Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)
```

See [docs](https://bun.uptrace.dev/guide/migrations.html) for details.
