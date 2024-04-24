
[[fswatch]] を使ってファイル数自動カウントを行う。
```shell
brew install fswatch
```

[update_count.sh](../update_count.sh)
```bash
#!/bin/bash

count=$(find . -name '*.md' -mindepth 2 | wc -l | sed "s/ //g")

sed -i '' "s/Markdown files: [0-9]*/Markdown files: $count/" README.md
```

[auto_update_info.sh](../auto_update_info.sh)
```bash
#!/bin/bash

# .md ファイルの変更を監視
fswatch -o . | while read num_changes; do
  # .md ファイルの変更があった場合、スクリプトを実行
  find . -name '*.md' -print0 | xargs -0 -n 1 -I {} bash update_count.sh
done
```


[source](https://chat.openai.com/share/a5a63e9f-7fc5-4886-bd82-d6d7082f882d)
