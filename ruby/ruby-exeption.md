## Ruby の例外処理
[[例外処理]] ([[例外]]発生時の処理) の構文 ([[begin]] ... [[rescue]] ... end)

```rb
begin
  # 例外が起きうる処理
rescue
  # 例外発生時の処理
end
```

### 例外オブジェクト
Ruby では発生した例外もオブジェクトであり、メソッド (e.g. [[message]], [[backtrace]]) を呼び出すことができる。

```rb
begin
  1 / 0
rescue => e
  puts "エラークラス: #{e.class}"
  puts "エラーメッセージ: #{e.message}"
  puts "バックトレース -----"
  puts e.backtrace
  puts "-----"
end

# エラークラス: ZeroDivisionError
# エラーメッセージ: divided by 0
# バックトレース -----
# (irb):14:in `/'
# (irb):14:in `<top (required)>'
# (省略)
# -----
```

### クラス指定による捕捉例外の限定
例外には多くの種類があり、例外ごとにクラスが異なる (e.g. NoMethodError, ZeroDivisionError)。
[[rescue]] で例外のクラスを指定すると、クラスが一致したときだけその例外を捕捉するようにできる。

```rb
begin
  # 例外が起きうる処理
rescue NoMethodError, ZeroDivisionError => e
  # 例外発生時の処理
end
```

すべての例外クラスは [[Exeption]] クラスを継承している。
[[StandardError]] クラスは通常のプログラムで発生する例外を表す。[[NoMethodError]] や [[ZeroDivisionError]] は StandardError のサブクラス。

rescue 節に Exeption を指定すると、NoMemoryError や SystemExit などプログラムに無関係な例外まで捕捉してしまう。Ruby の例外処理で捕捉するのは [[StandardError]] とそのサブクラスに限定すべき。

また、rescue 節で捕捉する例外の継承関係を考慮しないと、永遠に実行されない節を書いてしまう可能性がある。例外を複数節で捕捉する場合は、サブクラス → スーパークラスの順で書くと良い。
```rb
begin
  'abc'.foo
rescue NameError => e
  puts e.message
rescue NoMethodError => e
  # NoMethodError は NameError のサブクラス
  # NameError が必ず先に捕捉されるため、この rescue 節は永遠に実行されない
  puts e.message
end
#=> undefined method `foo' for "abc":String
```

source: [[『プロを目指す人のためのRuby入門［改訂2版］』]]
