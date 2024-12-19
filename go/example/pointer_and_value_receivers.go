package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int       // 合計
	lastUpdated time.Time // 最終更新時刻
}

// ポインタレシーバ
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// 値レシーバ
// (「ポインタレシーバのメソッドが1つでもあればすべてのメソッドに
// ポインタレシーバを使って形式を揃える」ルールに違反している)
func (c Counter) String() string {
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment() // `(&c).Increment()` と書かなくてもよい
	fmt.Println(c.String())
}
