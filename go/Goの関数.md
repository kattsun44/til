## 関数の宣言

引数を取り、値を返す[[関数]]の宣言 … キーワード `func`、関数名、引数、戻り値の型で構成

```go
// 例 … 0で割ろうとすると0を返す割り算関数
func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
```

## 名前付き引数とオプション引数

Goで[[名前付き引数]]と[[オプション引数]]を実現するには、各引数に対応するフィールドを持った構造体を定義し、それを関数にわたす。

しかし、そのような引数がほしいと思うということは、関数が複雑すぎる可能性を考えたほうがいいかもしれない。

## 可変長引数

Go言語でも[[可変長引数]] (variadic parameters) が使える。

- 可変長引数は引数リストの最後 (もしくは唯一) の引数でなければならない
- 型の **前に** `...` を付けて表す
- 可変長引数に対応する変数として関数内で作成されるのは指定された型の[[スライス]]
- 可変長引数にスライスを渡すときは、スライスリテラルの **後に** `...` を書く
  - そうしないとコンパイルエラーになる


```go
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	fmt.Println(addTo(3))             // []
	fmt.Println(addTo(3, 2))          // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))                    // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]
}
```
ref: https://go.dev/play/p/PYnEaoHOV66


## 複数の戻り値

Go言語では複数の戻り値が返せる。関数が複数の戻り値を返す場合、戻り値の型を `(int, int, error)` などのように `()` で示して宣言する。

## 関数は値

Go言語の関数は「[[値]] (value)」。関数の型はキーワード `func`、および引数と戻り値の型によって決まる。この組み合わせを関数の[[シグネチャ]] (signature) と呼ぶ。

2つの関数の引数と戻り値の数と型が同じであれば、両者のシグネチャが一致することになる。

このように関数を値として扱うことで賢い使い方ができる → [example/expression.go](example/expression.go)

- マップ `opMap` に対するキーとして `op` を使い、キーに結びつく値を変数 `opFunc` に代入する
- `opFunc` の型は `func(int, int) int`
- キーに結びつけられた関数がマップになければエラー

### 関数型の宣言

[[関数型]]を定義するときにも `type` が使える。

```go
type opFuncType func(int, int) int
```

このように型を定義すれば、[example/expression.go](example/expression.go) の `opMap` は以下のように書き換えられる。

```diff
- var opMap = map[string]func(int, int) int{
+ var opMap = map[string]opFuncType{
	"+": add,
	...
```

- 関数型を宣言する利点
  - 説明として使える
  - and more...

## 無名関数

Goにおいて関数は変数に代入できるだけではなく、名前を持たない関数「[[無名関数]] (匿名関数)」が使える。

無名関数を宣言する場合は、キーワード `func` + 引数 + 戻り値 + `{` を書く。`func` と引数の間に関数名を書くとコンパイルエラー。

無名関数は [[defer]] 文や[[ゴルーチン]]の起動で役立つ。

source: [[『初めてのGo言語』]]
