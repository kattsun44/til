#!/bin/bash

count=$(find . -name '*.md' -mindepth 2 | wc -l | sed "s/ //g")

sed -i '' "s/Markdown files: [0-9]*/Markdown files: $count/" README.md
