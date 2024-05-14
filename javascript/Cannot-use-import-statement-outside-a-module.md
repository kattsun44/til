## SyntaxError: Cannot use import statement outside a module
```js
% node test.js
(node:78934) Warning: To load an ES module, set "type": "module" in the package.json or use the .mjs extension.
(Use `node --trace-warnings ...` to show where the warning was created)
/Users/kattsun/test.js:1
import { BskyAgent } from '@atproto/api'
^^^^^^

SyntaxError: Cannot use import statement outside a module
    at wrapSafe (node:internal/modules/cjs/loader:1389:18)
    at Module._compile (node:internal/modules/cjs/loader:1425:20)
    at Module._extensions..js (node:internal/modules/cjs/loader:1564:10)
    at Module.load (node:internal/modules/cjs/loader:1287:32)
    at Module._load (node:internal/modules/cjs/loader:1103:12)
    at Function.executeUserEntryPoint [as runMain] (node:internal/modules/run_main:168:12)
    at node:internal/main/run_main_module:30:49
```

## 原因
ES6 モジュール (ECMAScript 2015) をサポートしていない環境で [[import]] 文を使用したため。

## 対策
以下のいずれかで対応可能。

- ファイル拡張子を `.mjs` に変更する
- package.json に `"type": "module"` を追加する
