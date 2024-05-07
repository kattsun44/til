## リファラー

[[リファラー]] ([[Referer]]) は、ユーザーがどの経路からWebサイトに到達したかをサーバーが把握するために、クライアントがサーバーに送るフィールド。

- クライアントが `http://www.example.com/link.html` のリンクをクリックしてべつのサイトにジャンプしたとき、ジャンプ先サーバーにリンク元ページ URL を以下の形式で送信する
```
Referer: http://www.example.com/link.html
```
- URL を手入力したり、ブックマークから選択したりした場合は、Referer を送信しないか `Referer: about:blank` を送信する

> [!warning]
> GET パラメーターはリファラー経由で外部サービスに送信される。プライバシーに関わる情報が GET パラメーターとして表示される仕様を設計してはならない。

> [!tip]
> リファラーそのものを送信しないブラウザ設定をすることは可能。

### リファラーの送信制限
[[RFC 2616]] でリファラー送信制限の追加が定められている。

アクセス元|アクセス先|送信するかどうか
-|-|-
HTTPS|HTTPS|する
HTTPS|HTTP|しない
HTTP|HTTPS|する
HTTP|HTTP|する

### リファラーポリシー
[リファラーの送信制限](#リファラーの送信制限) を厳密に適用するとサービス間連携に支障が出る場合がある。そのため、[[リファラーポリシー]]が 2014年に W3C に提案された。

#### リファラーポリシーの設定方法
リファラーポリシーは以下のいずれかの方法で設定できる。
- [[Referrer-Policy]] ヘッダーフィールド
- `<meta name="referrer" content="xxx">`
- 要素 (`<a>` など) の [[referrerpolicy]] 属性および `rel="noreferrer"` 属性

リファラーポリシーに設定できる値は以下の通り。
- `no-referrer`: 一切送らない
- `no-referrer-when-downgrade`: HTTPS → HTTP 時は送らない。[現在のデフォルト](#リファラーの送信制限)操作
- `same-origin`: 同一ドメイン内リンクに対してのみリファラーを送信
- `origin`: ドメイン名のみ送信 (トップページからリンクされた扱い)
- `strict-origin`: `origin` と同じだが、HTTPS → HTTP 時は送らない
- `origin-when-crossorigin`: 同ドメインではフルのリファラー、別ドメインではトップのドメイン名のみ送信
- `strict-origin-when-crossorigin`: `origin-when-crossorigin` と同じだが以下略
- `unsafe-url`: 常に送信

`Content-Security-Policy referrer origin` のように、[[Content-Security-Policy]] ヘッダーフィールドでも指定可能。

source: [[『Real World HTTP 第3版』]]
