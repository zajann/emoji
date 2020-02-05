package main

import (
	"fmt"
	"os"

	"github.com/zajann/emoji/internal/data"
)

var emojis data.EmojiMap

func init() {
	emojis = data.Emojis
}

func printEmoji(shortcode string) {
	shortcode = fmt.Sprintf(":%s:", shortcode)
	code, ok := emojis[shortcode]
	if !ok {
		fmt.Printf("there is no emoji of '%s'\n", shortcode)
	} else {
		fmt.Printf("%-10s %s\n", shortcode, code.Unicode)
	}
}

func searchEmojiByTag(tag string) {
	codeMap := make(map[string]string)

	for k, v := range emojis {
		if _, ok := v.TagMap[tag]; ok {
			codeMap[k] = v.Unicode
		}
	}

	if len(codeMap) == 0 {
		fmt.Printf("No results for '%s'\n", tag)
	} else {
		fmt.Printf("Results for '%s'\n", tag)
		for k, v := range codeMap {
			fmt.Printf("%-10s %s\n", v, k)
		}
	}
}

func main() {
	sc := os.Args[1]
	printEmoji(sc)
	searchEmojiByTag(sc)
}
