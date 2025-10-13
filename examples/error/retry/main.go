package main

import (
	"errors"
	"fmt"
	"time"
)

func unreliable() error {
	return errors.New("temporary network error")
}

func main() {
	var err error
	for i := 0; i < 3; i++ {
		err = unreliable()
		if err == nil {
			fmt.Println("Success")
			return
		}
		fmt.Println("Retrying...", i+1)
		time.Sleep(time.Second)
	}
	fmt.Println("Failed after retries:", err)
}
