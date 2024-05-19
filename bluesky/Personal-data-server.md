## PDS

[[PDS]] (Personal Data Server/Store) は [[Bluesky]] の

- アカウント情報や投稿データなどの保管
- クライアントからの API アクセスの窓口

を担うサービス。<br>
Bluesky 公式の PDS 以外に、自分で PDS を立ち上げることも可能。


### 役割 1. アカウント情報管理
Bluesky のアカウントには Handle と [[DID]] という情報が紐づいている。

[[DID]] は全アカウント間で重複しない永続的な ID で、アカウント登録時に [[PLC]] によって新規発行される。またアカウント削除時に他アカウントに使い回されることがない。<br>
Handle は DID を示す一種のラベルで、他のアカウントが使用していない任意のドメインを登録することができる。変更も可能だが、その場合過去 Handle の URL は無効になる。

その他、Display name, Description, アバター画像などのアカウントそのものの情報に加え, アカウント関連データの保存を担う。

### 役割 2. API リクエスト受付窓口
ポスト作成・削除、いいねをつける、リポスト、などのアカウント自身の管理データについての処理を行う。<br>
その他のデータ表示系 API は [[App View]] サービスに転送される。




sources:
- [開発視点から見る、新しい分散型SNS「Bluesky」とAT Protocolの可能性 | gihyo.jp](https://gihyo.jp/article/2023/04/bluesky-atprotocol)
- [Hello Nostr! Yo Bluesky! 分散SNSの最前線：四谷ラボ](https://techbookfest.org/product/6quLEm85cpd4TMJR17xnVF?productVariantID=kgmgxRsKgbVruvRd2zV1sp)
- https://bsky.app/profile/kattsun.dev/post/3ksrehl5wfc26
