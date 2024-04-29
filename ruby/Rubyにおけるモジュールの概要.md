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
