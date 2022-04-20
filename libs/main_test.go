package main

import (
	"fmt"
	"testing"
)

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}

// test成功
func TestSuccess(t *testing.T) {
	fmt.Println("テスト成功！！")
}

// test失敗
// func TestFailure(t *testing.T) {
// 	t.Errorf("テスト失敗！！")
// }
