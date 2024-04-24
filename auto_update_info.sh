#!/bin/bash

# .md ファイルの変更を監視
fswatch -o . | while read num_changes; do
  # .md ファイルの変更があった場合、スクリプトを実行
  find . -name '*.md' -print0 | xargs -0 -n 1 -I {} bash update_count.sh
  find . -name '*.md' -print0 | xargs -0 -n 1 -I {} bash update_file_list.sh
done
