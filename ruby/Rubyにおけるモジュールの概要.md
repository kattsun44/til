## Ruby におけるモジュールの概要
Ruby の[[モジュール]]を定義する構文
```rb
module モジュール名
end
```

モジュールはクラスと違い、
1. モジュールからインスタンスを作成することはできない
2. 他のモジュールやクラスを継承することはできない

```rb
module Greetable
  def hello = 'Hello!'
end

greetable = Greetable.new
#=> undefined method `new' for Greetable:Module (NoMethodError)

module AwesomeGreetable < Greetable
end
#=> syntax error, unexpected '<' (SyntaxError)
```
### include
モジュールをクラスに [[include]] する (Ruby ではこれを[[ミックスイン]]と呼ぶ) ことで、モジュール内のメソッドをそのクラスの[[インスタンスメソッド]]として呼び出すことができる。
```rb
class User
  # クラス構文内で include
  include Greetable
end
user = User.new
user.hello #=> "Hello!"

# クラス構文外で include
class Admin end
Admin.include Greetable
admin = Admin.new
admin.hello #=> "Hello!"
```

複数のクラスにまたがる共通の機能 (e.g. ログ機能) を実装する際に用いられる。

[[include?]] を呼ぶと引数で渡したモジュールが include されているかどうかが分かる。
また、[[included_modules]] を呼ぶと include されているモジュールの配列が返る。
```rb
User.include?(Greetable) #=> true
User.included_modules #=> [Greetable, PP::ObjectMixin, Kernel]

user = User.new
user.class.included_modules #=> [Greetable, PP::ObjectMixin, Kernel]
```

[[ancestors]] を使うとスーパークラスの情報も含まれる。
```rb
User.ancestors #=> [User, Greetable, Object, PP::ObjectMixin, Kernel, BasicObject]
```

インスタンスに対しては [[class]] メソッド経由で上記メソッドを呼ぶか、直接 [[is_a?]] でモジュールを include しているか判定できる。
```rb
user = User.new
user.class.included_modules #=> [Greetable, PP::ObjectMixin, Kernel]
user.is_a?(Greetable) #=> true
```

### extend
モジュールをクラス (や他のモジュール) に [[extend]] すると、モジュール内のメソッドを特異メソッドとして呼び出すことができる。
```rb
module AwesomeGreetable
  # クラス構文内で extend
  extend Greetable
end
AwesomeGreetable.hello #=> "Hello!"

# クラス構文外で extend
module AnotherGreetable end
AnotherGreetable.extend Greetable
AnotherGreetable.hello #=> "Hello!"
```

source: [[『プロを目指す人のためのRuby入門［改訂2版］』]]
