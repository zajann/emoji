# emojiDataGenerator

A generator tool of https://github.com/zajann/emoji . 

The tool generate Emoji Data and create data file.

## Emoji Source

- https://raw.githubusercontent.com/CodeFreezr/emojo/master/db/v5/emoji-v5.json
- https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json

## Data File(.go)

The tool create [data file](https://github.com/zajann/emoji/internal/data/emojiMap.go). Path is fixed to the relative path (`../../internal/data/emojiMap.go`)

*MUST KEEP THIE PATH. OR PROGRAM WILL BROKEN*

## Update Emoji Data

``` bash
# data file will be recreated
go run *.go
```

