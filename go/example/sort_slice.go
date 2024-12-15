package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("初期データ: ", people)

	// 姓 (LastName) でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("姓でソート: ", people)

	// 年齢 (Age) でソート
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("年齢でソート: ", people)

	// 最後に people がどうなっているか確認
	fmt.Println("ソート後の people: ", people)
}
