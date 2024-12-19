package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int       // 合計
	lastUpdated time.Time // 最終更新時刻
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment() // main の c のコピーに対して Increment が行われる
	fmt.Println("NG:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment() // main の c に対して Increment が行われる
	fmt.Println("OK:", c.String())
}

func main() {
	var c Counter
	doUpdateWrong(c)
	fmt.Println("main:", c.String())
	doUpdateRight(&c)
	fmt.Println("main:", c.String())
}
