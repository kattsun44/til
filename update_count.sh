#!/bin/bash

count_md=$(find . -name '*.md' | wc -l | sed "s/ //g")
count_til=$((count_md - 1)) # README.md 分マイナス

sed -i '' "s/Markdown files: [0-9]*/Markdown files: $count_til/" README.md
