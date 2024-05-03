## Go言語の基本的な構文

### 変数宣言
Go で変数を宣言するときは [[var]] を使う。
```go
// int 型の変数 n を宣言
var n int

// 宣言と初期値の代入を同時に行える
var n int = 1

// 変数は型を保持し、同じ変数名で異なる型を代入することはできない
var n int
n = 1
n = 2
n = "foo" //=> cannot use "foo" (untyped string constant) as int value in assignment

// 再宣言もできない
var n int
var n string //=> n redeclared in this block

// 変数宣言は var を省略できる (型は自動で推測されるため同様に省略可能)
x := 1
y := 1.2

// 異なる型同士の演算はコンパイルエラー
n := (x + 2) * y //=> invalid operation: (x + 2) * y (mismatched types int and float64)
// int(y) で y を int 型にキャスト
n := (x + 2) * int(y) // n: 3
```

Go には以下のビルトイン型がある
- [[interface{}]] ([[any]])
- [[bool]]
- [[byte]]
- [[complex64]], [[complex128]]
- [[error]]
- [[float32]], [[float64]]
- [[int]], [[int8]], [[int16]], [[int32]], [[int64]]
- [[rune]], [[string]]
- [[uint]], [[uint8]], [[uint16]], [[uint32]], [[uint64]]
- [[uintptr]]

### const (定数宣言)
Go で定数を宣言するときは [[const]] を使う。変数と異なり未使用でもコンパイルエラーにはならない。
```go
const x = 1
```
配列、スライス、struct を宣言することはできない (つまりこれらをイミュータブルな定数にすることはできない)。

#### untyped constant
Go では const で型を明示的に指定せずに宣言した定数は型を持たない。使われる場所で型がその都度決まる。
```go
// 変数の場合、宣言時に型が決まるため、異なる型との演算時にコンパイルエラー
var x = 1
fmt.Println(x + 1)
fmt.Println(x + 1.2) //=> 1.2 (untyped float constant) truncated to int

// 定数の場合、型が都度決まるため、以下のどちらもエラーにならない
const y = 2
fmt.Println(y + 1)
fmt.Println(y + 1.2)
```

## iota (列挙)
Go には [[enum]] のような列挙型はないが、[[iota]] を使うことで同様の定数を宣言できる
```go
const (
  Apple = iota
  Orange
  Banana
)
fmt.Println(Apple, Orange, Banana) //=> 0 1 2
```
```go
const (
  Apple = iota + iota
  Orange
  Banana
)
fmt.Println(Apple, Orange, Banana) //=> 0 2 4
```
```go
const (
  Apple = iota + iota
  Orange
  // 途中で異なる宣言が可能
  Banana = iota + 3
)
fmt.Println(Apple, Orange, Banana) //=> 0 2 5
```

source: [[『Go言語プログラミングエッセンス』]]
