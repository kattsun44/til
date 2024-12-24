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

