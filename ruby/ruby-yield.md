## yield
Ruby のメソッド内で [[yield]] を使うと、メソッド呼び出し時に紐づけたブロックが実行される。
```rb
def greet
  puts 'おはよう'
  yield
  puts 'こんばんは'
end
```

```rb
greet { puts 'こんにちは' }
# おはよう
# こんにちは
# こんばんは
```

ブロックなしでメソッドが呼び出されているのに [[yield]] を実行しようとしたらエラー。
```rb
greet
#=> no block given (yield) (LocalJumpError)
```

ブロックが渡されたかどうかを確認するには [[block_given?]] メソッドを使う。
```rb
def greet
  puts 'おはよう'
  if block_given?
    yield
  end
  puts 'こんばんは'
end
```

```rb
greet { puts 'こんにちは' }
# おはよう
# こんにちは
# こんばんは
greet
# おはよう
# こんばんは
```

[[yield]] はメソッドに紐づけるブロックに引数を渡したり、そのブロックの戻り値を受け取ったりできる。
```rb
def greet
  puts 'おはよう'
  # ブロックに引数を渡し、戻り値を取得
  text = yield 'こんにちは'
  puts text
  puts 'こんばんは'
end
```

```rb
greet { |t| t * 2 }
# おはよう
# こんにちはこんにちは
# こんばんは
greet {}
# おはよう
#
# こんばんは
```

source: [[『プロを目指す人のためのRuby入門［改訂2版］』]]
