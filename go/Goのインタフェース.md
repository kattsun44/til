---
tags:
  - 初めてのGo言語
---

## インタフェースとは

[[インタフェース]]は、Go で唯一の[[抽象型]] (実装を提供しない型) で、次の2つの側面を持つ。

1. メソッドの集合――多言語のインタフェースのように、特定のクラスが満たすべき要件（実装するべき一群のメソッド）を示す
2. 型――変数がインタフェースを基盤とする型を持つことで、様々なこと（e.g. 任意の型の値を代入できる変数を定義[^2024-12-24-084804]）ができる

[^2024-12-24-084804]: 特によく使われるのが [[interface{}]]。これは「0個のメソッドが定義された型」という意味になり、任意の型がこの条件を満たすため任意の型の値を記憶できる。[[Go 1.18]] からは `interface{}` の代わりに [[any]] と書けるようになった。

### インタフェースを宣言する方法

キーワード `type` + インタフェース名 + インタフェースリテラル (`interface` + メソッドリスト[^2024-12-24-090105])

[^2024-12-24-090105]: インタフェースを満たすために実装するメソッドのリスト。メソッドセットとも言う

```go
type Stringer interface {
	String() string
}
```

### インタフェースの特徴

- [[インタフェース型]]は型の集合を定義する。`int` や `string`、「`struct` を使って定義した型」と同様、**インタフェースも型の一種**
- インタフェース型の変数は、そのインタフェースが特定する型の集合に属する任意の型の値を記憶できる
- 任意のブロックでインタフェースを定義できる (他の型と同様)
- インタフェースの名前は、"er" で終わるのが通例[^2024-12-24-090230]
- あるインタフェースを満たす型に、そのインタフェースにないメソッドを追加しても問題ない
  - e.g. [[os.File]] はインタフェース [[io.Reader]] と [[io.Writer]] を満たすが、ファイル読み込みをするだけのコードなら `io.Reader` だけ使い、他のメソッドは無視する[^2024-12-25-080827]

[^2024-12-25-080827]: このメリットを説明してくれていそうな記事 → [Go言語 暗黙的なinterfaceのメリット \#Java - Qiita](https://qiita.com/sasakitimaru/items/02652b433e2dbe0e5149)

[^2024-12-24-090230]: [[fmt.Stringer]], [[io.Reader]], [[io.Closer]], [[io.ReadCloser]], [[json.Marshaler]], [[http.Handler]] など、標準ライブラリには多くのインタフェースが使われている。


### Goのインタフェースは「暗黙的」

[インタフェースの特徴](#インタフェースの特徴)は他言語のそれと大きく異なるものはないが、Goのインタフェースはそれが「暗黙的に」実装される点が特別。[[具象型]] (具体的な実装) のほうではあるインタフェースを実装することを明示的に宣言することはなく、ある具象型 `C` のメソッドセットがあるインタフェース `I` のメソッドセットを完全に含めば、具象型 `C` はインタフェース `I` を「暗黙的（自動的）に」実装したことになる。

この仕様により、Goでは「型の安全性」(type-safe) と「デカップリング[^2024-12-24-091501]」(decoupling、分離) の両方が達成され、静的言語と動的言語の両方の機能を併せ持つことができている。

[^2024-12-24-091501]: インタフェースと具象型の「分離」が達成され、コードの[[変更容易性]]が高められるということらしい。

## インタフェースは型安全なダックタイピング

Python, Ruby, JavaScript などの[[動的型付け]]の言語にはインタフェースがない代わりに、[[ダックタイピング]][^2024-12-24-215017]を使う。

[^2024-12-24-215017]: 「アヒルのように歩き、アヒルのようにガアガア鳴くなら、それはアヒルだ」という表現に由来する言葉。「関数 `f` を起動する方法を見つけられるなら、その (ある型の) インスタンス `t` は関数 `f` の引数にできる」という考え方。

```python
class Logic:
    def process(self, data):
        # ビジネスロジックがここに書かれる
        print(data)

def program(logic):
    # どこからかデータを取得
    logic.process(data)

logicToUse = Logic()
program(logicToUse)
```

一方で、Java のプログラマーはインタフェースを定義し、そのインタフェースを実装するが、使う側 (`Client`) コード中ではインタフェースだけを参照する。

```java
public interface Logic {
  String process(String data);
}

public class LogicImpl implements Logic {
  public String process(String data) {
    // ビジネスロジックがここに書かれる
  }
}

public class Client {
    private final Logic logic; // この型はインタフェースで、実装ではない

    public Client(Logic logic) {
        this.logic = logic;
    }

    public void program() {
      // どこからかデータを取得
      this.logic.process(data);
    }
}

public static void main(String[] args) {
  Logic logic = new LogicImpl();
  Client client = new Client(logic);
  client.program(); // 1
}
```

Go は2つのスタイル混ぜ合わせたものになる。

```go
package main

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	// ビジネスロジックがここに書かれる
	return data
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// どこからかデータを取得
	c.L.Process(data)
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
```

Goではインタフェースを知っているのは呼び出し側 (`Client`) だけ。呼び出される `LogicProvider` はインタフェースに適合していることを示すものが何も宣言されていない。これにより、

- 将来新しいロジックを組み込むことができる[^2024-12-24-222740]
- クライアントに渡された任意の型がクライアントの要求にマッチすることを保証する、実行可能なドキュメントを提供できる

[^2024-12-24-222740]: (1) どこに？ (2) Java の場合は新しいロジックを組み込むことができない？

## 標準インタフェースとデコレータ

[[標準ライブラリ]]のインタフェースを使うと[[デコレータ]]パターンが使いやすくなる。

Goでは、インタフェースのインスタンスを受け取り、そのインタフェースを実装する別の型を返す「[[ファクトリ関数]][^2024-12-25-073751]を書くのが一般的。

[^2024-12-25-073751]: `func bufio.NewScanner(r io.Reader) *bufio.Scanner` みたいなやつ。Goにはコンストラクタがない代わりにファクトリ関数を使う。

e.g. 次のように定義された関数があったとする

```go
func process(r io.Reader) error
```

次のようなコードでファイルを処理できる。

```go
r, err := os.Open(fileName)
if err != nil {
	return err
}
defer r.Close()
return process(r)
```

[[os.Open]] が返す [[os.File]] のインスタンスは [[io.Reader]] インタフェースを満たしているため、データを読み込むどんなコードにも使える。

e.g. ファイルが gzip で圧縮されていた場合

```go
r, err := os.Open(fileName)
if err != nil {
	return err
}
defer r.Close()

gz, err := gzip.NewReader(r) // io.Reader をもうひとつ別の io.Reader でラップ
if err != nil {
	return err
}
return process(gz)
```

## インタフェースを受け取り構造体を返す

経験豊かなGoプログラマー「インタフェースを受け取り、構造体を返すようにコードを書け[^2024-12-26-090559]」

[^2024-12-26-090559]: [Go言語のInterfaceの考え方、Accept interfaces,return structs \#Go - Qiita](https://qiita.com/weloan/items/de3b1bcabd329ec61709)、[GoのInterfaceの作法 "Accept Interfaces, Return structs" - y-zumiの日記](https://y-zumi.hatenablog.com/entry/2019/07/28/035632)


- 関数内で起動されるビジネスロジックはインタフェースによって起動されるべき
- それに対し、関数の出力は具体的な型であるべき

こうすることで、コードが柔軟になり、どのような機能が使われているのかを正確かつ明示的に宣言できる。

### インタフェースを返さないようにする理由

インタフェースを返さないようにする理由1 … [[デカップリング]]の維持

もしもインタフェースを返す API を作成してしまうと、コードがそのインタフェースを含むモジュールに依存 (= [[カップリング]]) してしまう。それに対して具体的なインスタンスへの依存であれば、[[依存性注入]]によって影響を限定できる。

インタフェースを返さないようにする理由2 … [[バージョン管理]]のため

具体的な型が返される場合は既存のコードを壊すことなくメソッドやフィールドを追加できる。しかし、インタフェースにメソッドを追加するという場合はそのインタフェースのすべての実装を書き換える必要がある。

例外はエラー。Goの関数・メソッドで宣言する戻り値 [[error]] はインタフェース型。`error` の場合はインタフェースの異なる実装が返される可能性が高いため、Go唯一の抽象型であるインタフェースを使ってすべての可能性を処理する必要がある。
