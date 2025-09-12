package main

import "fmt"

// サーバー設定を保持する構造体
type Server struct {
	Host   string
	Port   int
	Secure bool
}

// サーバー設定のオプションを定義する関数型
type ServerOption func(*Server)

// 新しいServerオブジェクトを生成してオプションを適用する
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		Host:   "localhost", // デフォルト値
		Port:   8080,        // デフォルト値
		Secure: false,       // デフォルト値
	}

	// 各オプションを適用
	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

// ホスト名を設定するオプション
func WithHost(host string) ServerOption {
	return func(srv *Server) {
		srv.Host = host
	}
}

// ポート番号を設定するオプション
func WithPort(port int) ServerOption {
	return func(srv *Server) {
		srv.Port = port
	}
}

// セキュアフラグを設定するオプション
func WithSecure(secure bool) ServerOption {
	return func(srv *Server) {
		srv.Secure = secure
	}
}

// 使用例
func main() {
	// カスタムオプションでサーバーを作成
	srv := NewServer(
		WithHost("example.com"),
		WithPort(9090),
		WithSecure(true),
	)

	fmt.Printf("Server: %+v\n", srv)
}
