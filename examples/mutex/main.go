package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Mutex を使って排他制御を行いながら、複数のゴルーチンが
// 並行して動いている様子をログ出力で確認できるサンプルプログラム
func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	// 5つのゴルーチンを起動
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 1; j <= 3; j++ {
				// クリティカルセクションに入る前にログ出力
				fmt.Printf("Goroutine %d: waiting for lock (iteration %d)\n", id, j)

				// 排他制御（ロックを取得）
				mu.Lock()

				// ロック取得後のログ（ここは他のゴルーチンと同時に動かない）
				fmt.Printf("Goroutine %d: acquired lock (iteration %d)\n", id, j)

				// 処理を模擬（1秒待機）
				time.Sleep(1 * time.Second)

				// ロック解除
				mu.Unlock()

				// ロックを解放した後のログ
				fmt.Printf("Goroutine %d: released lock (iteration %d)\n", id, j)

				// 他の処理（排他不要な部分）
				time.Sleep(500 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished.")
}
