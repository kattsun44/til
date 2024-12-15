## 構造体の宣言

キーワード `type`、[[構造体]]の名前、キーワード `struct` を書き、 `{...}` に囲まれたフィールドを書く。フィールドはまずフィールド名を書き、次にそのフィールドの型を書く。フィールドを分けるときに `,` は不要。

```go
type person struct {
	name string // 名前
	age  int    // 年齢
	pet  string // ペット
}
```

## 構造体の定義

構造体を宣言したら、その構造体を型としてもつ変数を定義できる。
値が代入されない場合、各フィールドはそのフィールドの型の[[ゼロ値]]をもったものになる。

```go
var taro person
fmt.Println(taro) // { 0 }
```

[[構造体リテラル]]を変数に代入することもできる。空の構造体リテラルを代入する場合も構造体の全フィールドが各フィールドのゼロ値で初期化される。

```go
jiro := person{}
fmt.Println(jiro) // { 0 }
```

空でない構造体を指定するスタイルは2つある。

1) 値を `,` で区切ってフィールドを並べる

この形式を使う場合はすべてのフィールドの値を順番に指定する必要がある。

```go
saburo := person{
	"三郎",
	40,
	"猫",
}
```

2) フィールド名と値を `:` で区切る

この形式ではフィールドの順番は自由であり、すべてのフィールドを指定する必要もない (指定されないフィールドはゼロ値で初期化される)。

```go
siro := person{
	age: 30,
	name: "四郎",
}
```

構造体のフィールドにアクセスするときは `.` を使う。

```go
taro.name = "太郎"
fmt.Println(taro.name) // 太郎
```


## 無名構造体

変数に対し、[[無名構造体]] (anonymous struct) という「名前のない構造体」を割り当てることもできる。

```go
var person struct {
	name string
	age  int
	pet  string
}

person.name = "太郎"
person.age = 100
person.pet = "犬"

pet := struct {
	name string
	kind string
}{
	name: "ポチ",
	kind: "犬",
}

fmt.Println(person, pet) // {太郎 100 犬} {ポチ 犬}
```

上記の例では、変数 `person` および `pet` の型は無名の構造体になっている。
無名構造体は名前付き構造体と同じように代入や値の呼び出しが可能。`pet` のように構造体リテラルで初期化することもできる。

用途 1) 外部データ (JSON など) と構造体を相互変換する場合。外部データ→構造体への変換を[[アンマーシャリング]]、構造体→外部データへの変換を[[マーシャリング]]と呼ぶ。

用途 2) テスト。テーブル駆動テストを書く際に無名構造体のスライスを用いる。


## 構造体の比較と変換

構造体が比較可能かどうかは、各フィールドの型に依存する。
すべてのフィールドが比較可能である構造体同士は比較可能だが、スライスヤマップ、関数、チャネルをフィールドとして持つ構造体は比較できない。

また仮にすべてのフィールドが比較可能であっても、型の異なる構造体同士の比較はできない (フィールドの名前、順番、型が同じであれば[[型変換]]は可能)。

```go
type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

type thirdPerson struct {
	age  int
	name string
}

type fourthPerson struct {
	firstName string
	age       int
}

type fifthPerson struct {
	name          string
	age           int
	favoriteColor string
}
```

|構造体名|firstPerson との[[型変換]]|firstPerson との比較|
|---|---|---|
|firstPerson|◯|◯|
|secondPerson|◯|✕ (型が異なるため)|
|thirdPerson|✕ (フィールドの順番が異なるため)|✕ (型が異なるため)|
|fourthPerson|✕ (フィールド名が異なるため)|✕ (型が異なるため)|
|fifthPerson|✕ (フィールドの数が異なるため)|✕ (型が異なるため)|

source: [[『初めてのGo言語』]]
