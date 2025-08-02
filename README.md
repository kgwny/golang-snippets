# golang-snippets

## コード整形
```
gofmt -w yourfile.go
```

## 実行
```
go run yourfile.go
```

## go mod 初期化
```
go mod init my-app
```

```
❯ go mod init my-app
go: creating new go.mod: module my-app
go: to add module requirements and sums:
	go mod tidy
```

## Makefile

### コマンド
| タスク             | 説明                                                |
| --------------- | ------------------------------------------------- |
| `make tidy`     | `go mod tidy` を実行して依存関係を整理                        |
| `make download` | `go mod download` で依存をローカルに取得                     |
| `make build`    | アプリをビルドして `bin/myapp` を作成                         |
| `make run`      | `main.go` を実行                                     |
| `make clean`    | ビルド成果物を削除                                         |
| `make test`     | 単体テスト実行                                           |
| `make lint`     | `go vet` および `staticcheck` で静的解析（※staticcheckは任意） |
| `make fmt`      | `go fmt` でコード整形                                   |
| `make help`     | `make` コマンドの一覧と説明を表示（自己ドキュメント化）                   |
