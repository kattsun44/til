## URL フラグメントテキストディレクティブ

[[URL フラグメントテキストディレクティブ]] ([[テキストフラグメント]]) は `:~:` の後に構造化された命令が書ける構文のこと。

構文
```http
https://example.com#:~:text=[prefix-,]textStart[,textEnd][,-suffix]
```
- `:~:`: フラグメントディレクティブ (fragment directive)。次に来るのが1つ以上のユーザーエージェント命令であることをブラウザに指示する
- `text=`: テキストディレクティブ (text directive)。ブラウザにテキストフラグメントを渡し、文書内のどのテキストにリンクされるかを定義する
- `prefix-`: リンク先テキストの前に来るべきテキストを指定する文字列。後ろにハイフン `-` がつく
- `textStart`: リンク先テキストの始まりを指定する文字列
- `textEnd`: リンク先テキストの終わりを指定する文字列
- `-suffix`: リンク先テキストの後に来るべきテキストを指定する文字列。前にハイフン `-` がつく

例 (Chrome の "Copy Link to Highlight" で取得したリンク)
```http
https://developer.mozilla.org/ja/docs/Web/Text_fragments#:~:text=この-,一連の文字,-は「フラグメン
```


sources: [テキストフラグメント | MDN](https://developer.mozilla.org/ja/docs/Web/Text_fragments), [[『Real World HTTP 第3版』]], [URL Fragment Text Directives](https://wicg.github.io/scroll-to-text-fragment/)
