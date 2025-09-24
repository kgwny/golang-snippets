# MySQLドライバー
github.com/go-sql-driver/mysql

## MySQL を起動して testdb データベースを作成しておく
```
CREATE DATABASE testdb;
```

## Go の MySQL ドライバをインストールする
```
go get -u github.com/go-sql-driver/mysql
```

DSN の user:password@tcp(127.0.0.1:3306)/testdb を自分の環境に合わせる
