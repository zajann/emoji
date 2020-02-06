package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/zajann/emoji/internal/data"
)

type codeMap map[string]string

const (
	NotFound int = iota
	Found
	RandomFound
)

var emojis data.EmojiMap

func init() {
	emojis = data.Emojis
}

func main() {
	random := flag.Bool("r", false, "search Emoji randomly")
	flag.Parse()

	// Searh Emoji randomly
	if *random {
		cnt := 0
		switch len(flag.Args()) {
		case 1:
			var err error
			cnt, err = strconv.Atoi(flag.Args()[0])
			if err != nil {
				log.Fatalf("Wrong argument type...Expecting 'Integer'")
			}
		case 0:
			cnt = 10 // default
		default:
			log.Fatalf("Wrong argument cnt...Expecting 1")
		}

		result := searchEmojiRandomly(cnt)
		printHeader("", RandomFound, len(result))
		printPretty(result)
		os.Exit(0)
	}

	for _, arg := range flag.Args() {
		keyword := strings.ToLower(arg)
		result := searchEmoji(keyword)

		if len(result) == 0 {
			printHeader(keyword, NotFound, 0)
		} else {
			printHeader(keyword, Found, len(result))
			printPretty(result)
		}
	}
}

func searchEmoji(keyword string) codeMap {
	c := make(codeMap)

	// search shortcode match
	shortcode := fmt.Sprintf(":%s:", keyword)
	code, ok := emojis[shortcode]
	if ok {
		c[shortcode] = code.Unicode
	}

	// search tags match
	for k, v := range emojis {
		if _, ok := v.TagMap[keyword]; ok {
			if _, ok := c[k]; !ok {
				c[k] = v.Unicode
			}
		}
	}
	return c
}

func searchEmojiRandomly(cnt int) codeMap {
	c := make(codeMap)

	i := 0
	for k, v := range emojis {
		c[k] = v.Unicode
		i++
		if i == cnt {
			break
		}
	}
	return c
}

func printPretty(c codeMap) {
	for k, v := range c {
		fmt.Printf("> %6s%-6s %s\n", "", v, k)
	}
}

func printHeader(keyword string, flag int, cnt int) {
	fmt.Println("=========================================")
	switch flag {
	case NotFound:
		fmt.Printf("# No results found for '%s'\n", keyword)
	case Found:
		fmt.Printf("# Results for '%s' ( %d of %d )\n", keyword, cnt, len(emojis))
	case RandomFound:
		fmt.Printf("# Results for randomly ( %d of %d )\n", cnt, len(emojis))
	}
	fmt.Println("=========================================")
}
