## goroutine (ゴルーチン)

[[goroutine]] ([[ゴルーチン]]) は、Go の[[ランタイム]]で管理された軽量な[[スレッド]]。

関数呼び出しの前に `go` をつけるだけで goroutine が実行される。

呼び出し時に引数がキャプチャされるため、呼び出し後に変数を変更しても goroutine で呼び出される引数は変わらない。
```go
import "time"

func sendMessage(msg string) {
	println(msg)
}

func main() {
	message := "hi"
	go sendMessage(message) //=> hi
	message = "ho"

	time.Sleep(time.Second)
}
```

ただし[[無名関数]]を実行するときは実行順序が保証されない。この状態を [[race condition]] と言い、goroutine とその呼び出し元との間でデータ競合が起きていることを意味する。

コンパイル ([[go build]]) 時に [[-race]] をつけて実行することでランタイムが race condition を検出してくれる。

```go
func main() {
	message := "hi"
	go func() {
		// ho が「出ることがある」
		// と本に書いてあったが ho しか出ない
		sendMessage(message)
	}()
	message = "ho"

	time.Sleep(time.Second)
}
```

### sync パッケージ
main 関数が goroutine の終了を待つためには [[sync]] パッケージを使う。
(使用しない場合は main 終了時に goroutine が強制終了してしまう)

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(1) // リファレンスカウンタを +1
	go func() {
		defer wg.Done() // リファレンスカウンタを -1

		// 重たい処理
	}()

	// 別の重たい処理

	wg.Wait() // リファレンスカウンタが 0 になるまで待つ
}
```

上記のように重たい処理を[[並行処理]]し、CPU を有効活用できるのが [[goroutine]] のメリット。

### goroutine をループで使う場合

```go
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
  wg.Add(1)
  go func() {
    defer wg.Done()
    fmt.Println(i)
  }()
}
wg.Wait()
```

の実行結果は10が多くなる。
goroutine 内で参照している変数 i を表示するときに、すでに for ループが終了している可能性があるため。
```shell
% go run main.go
10
10
10
10
4
10
10
10
4
10
```

この場合はループのスコープ内で新しい変数を宣言するか、無名関数の引数として渡すとよい
```go
for i := 0; i < 10; i++ {
  n := i // ループのスコープ内で新しい変数を宣言
  wg.Add(1)
  go func() {
    defer wg.Done()
    fmt.Println(n)
  }()
  // もしくは無名関数の引数として渡す
  // go func(n int) {
  //   defer wg.Done()
  //   fmt.Println(n)
  // }(i)
}
```

```shell
% go run main.go
7
4
3
6
5
8
1
9
2
0
```

### sync.Mutex によるデータ保護
goroutine 内と呼び出し元で同じ変数を参照・更新する場合も [[race condition]] が発生することがある。
データを保護するには [[sync.Mutex]] を使う。

```go
// データを保護しない場合
func main() {
		n := 0
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				n++
			}
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				n++
			}
		}()

		wg.Wait()

		fmt.Println(n) // 2000 が出力されないことがある
}
```

```go
// データを保護する場合
func main() {
		n := 0
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mu.Lock()
				n++
				mu.Unlock()
			}
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mu.Lock()
				n++
				mu.Unlock()
			}
		}()

		wg.Wait()
		fmt.Println(n) // 必ず 2000 が出力される
}
```

source: [[『Go言語プログラミングエッセンス』]]
