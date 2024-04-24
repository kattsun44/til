#!/bin/bash

# LIST.md ファイルを作成またはクリア
echo "# List" > LIST.md

# ディレクトリを再帰的に検索し、Markdown リンクとして整形して LIST.md に追加
find . -name '*.md' -mindepth 2 | sort | while read -r filepath; do
    # ディレクトリパスを抽出
    dirpath=$(dirname "$filepath")
    # ファイル名を抽出
    filename=$(basename "$filepath")
    # ディレクトリ名から最後のディレクトリ名部分を取得
    dirname=$(basename "$dirpath")

    # 前回のディレクトリと異なる場合は、新しいディレクトリヘッダを出力
    if [[ "$prev_dirpath" != "$dirpath" ]]; then
        echo -e "\n### $dirname\n" >> LIST.md
    fi

    # Markdown リンクとしてファイル名を出力
    echo "- [$filename]($filepath)" >> LIST.md

    # 前回のディレクトリパスを更新
    prev_dirpath="$dirpath"
done
