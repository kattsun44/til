## Enumerableモジュール
[[Enumerable]] モジュールは「繰り返し処理ができるクラス」に include されているモジュール。

```rb
Array.include?(Enumerable) #=> true
Hash.include?(Enumerable)  #=> true
Range.include?(Enumerable) #=> true
```

Enumerable に定義されている代表的なメソッド … [[map]], [[select]], [[find]], [[count]]
```rb
# map
[1, 2, 3].map { |n| n * 10 }               #=> [10, 20, 30]
{ a: 1, b: 2, c: 3 }.map { |_, v| v * 10 } #=> [10, 20, 30]
(1..3).map { |n| n * 10 }                  #=> [10, 20, 30]

# select
[1, 2, 3].select { |e| e >= 2 }               #=> [2, 3]
{ a: 1, b: 2, c: 3 }.select { |_, v| v >= 2 } #=> {:b=>2, :c=>3}
(1..3).select { |e| e >= 2 }                  #=> [2, 3]

# count
[1, 2, 3].count            #=> 3
{ a: 1, b: 2, c: 3 }.count #=> 3
(1..3).count               #=> 3
```
Enumerable モジュールを include して、このモジュールに定義されたメソッドを使えるようにする条件 … include 先のクラスで [[each]] メソッドが定義されていること

source: [[『プロを目指す人のためのRuby入門［改訂2版］』]]
