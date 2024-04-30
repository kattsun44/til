## Kernelモジュール
[[Kernel]] モジュールが提供するメソッドの一例
- [[puts]], [[p]], [[pp]], [[print]]
- [[require]]
- [[loop]]

Object クラスが Kernel モジュールを include しているため、(BasicObject を除く) すべてのクラスは Kernel モジュールのメソッドを使える。
```rb
Object.include?(Kernel) #=> true
```

### main
irb 起動直後、あるいは rb ファイル内のクラス構文やモジュール構文に囲まれていない場所を[[トップレベル]]と呼ぶ。
トップレベルの self は [[main]] という名前の Object クラスのインスタンスである。そのため、Kernel モジュールのメソッドが使える。

```rb
irb(main):001:0> self
=> main
irb(main):002:0> self.class
=> Object
irb(main):003:0> puts self
main
=> nil
```
