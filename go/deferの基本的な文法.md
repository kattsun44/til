## defer

Go の [[defer]] は後処理 (= 関数を抜ける際に実行される処理) を記述できる。

defer が複数指定されると、指定とは逆順に実行される。



defer 呼び出し時の変数はキャプチャされ、続く処理で変数が上書きされたとしても defer が呼ばれたタイミングの変数が参照される。
```go
f, err := os.Open("data01.txt")
if err != nil {
	log.Fatal(err)
}
defer f.Close() // 上書き前の f を閉じる

f, err := os.Open("data02.txt") // f が上書きされる
if err != nil {
	log.Fatal(err)
}
defer f.Close() // 上書き後の f を閉じる
```

defer には[[無名関数]]を渡すことができる。下の例は defer 呼び出し時に n がキャプチャされていないため、2が出力される。
```go
n := 1
defer func() {
	fmt.Println(n) //=> 2
}()

n = 2
```

無名関数に引数として渡したらキャプチャされる。
```go
n := 1
defer func(n int) {
	fmt.Println(n) //=> 1
}(n)

n = 2
```
