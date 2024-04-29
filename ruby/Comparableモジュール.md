## Comparableモジュール
[[Comparable]] モジュールは比較演算を可能にするモジュール。include すると以下のメソッドが使えるようになる。

- [[<]]
- [[<=]]
- [[==]]
- [[>]]
- [[>=]]
- [[between?]]

### <=> 演算子
Comparable モジュールのメソッドを使えるようにするには、include 先のクラスで以下の挙動をする [[<=>]] (宇宙船演算子、UFO演算子) を実装しておく必要がある。

`a <=> b` の挙動
- a が b より大きいなら正の整数を返す
- a と b が等しいなら0を返す
- a が b より小さいなら負の整数を返す
- a と b が比較できないなら nil を返す

```rb
# Conmparable を include している文字列や数値の <=> の挙動
2 <=> 1 #=> 1
2 <=> 2 #=> 0
1 <=> 2 #=> -1
1 <=> 'abc' #=> nil
'xyz' <=> 'abc' #=> 1
'abc' <=> 'abc' #=> 0
'abc' <=> 'xyz' #=> -1
'abc' <=> 1 #=> nil
```

### Comparable モジュールを独自クラスに include する
Comparable モジュールを include し、<=> 演算子を定義することで独自クラスで Comparable が使えるようになる。
```rb
class Tempo
  include Comparable

  attr_reader :bpm

  def initialize(bpm)
    @bpm = bpm
  end

  # Comparable を使うために定義
  def <=>(other)
    # Tempo 同士を <=> で比較、それ以外は nil を返す
    other.is_a?(Tempo) ? bpm <=> other.bpm : nil
  end

  # irb 上で結果を見やすくするために inspect をオーバーライド
  def inspect
    "#{bpm}bpm"
  end
end
```

```rb
t120 = Tempo.new(120) #=> 120bpm
t180 = Tempo.new(180) #=> 180bpm
t120 > t180 #=> false
t120 <= t180 #=> true
t120 == t180 #=> false
t120 <=> t180 #=> -1
```

[[<=>]] は並び替えで利用されるため、[[sort]] も利用可能になる。
```rb
tempos = [Tempo.new(180), Tempo.new(60), Tempo.new(120)] #=> [180bpm, 60bpm, 120bpm]
tempos.sort #=> [60bpm, 120bpm, 180bpm]
```

source: [[『プロを目指す人のためのRuby入門［改訂2版］』]]
