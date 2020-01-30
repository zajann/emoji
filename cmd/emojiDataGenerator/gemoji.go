package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const GemojiJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"

// Gemoji gemoji struct deserialized json
type Gemoji struct {
	Aliases     []string `json:"aliases"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

func getGemojiData() ([]EmojiInfo, error) {
	res, err := http.Get(GemojiJsonURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	emojiJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var gs []Gemoji
	if err := json.Unmarshal(emojiJson, &gs); err != nil {
		return nil, err
	}

	var eis []EmojiInfo
	for _, g := range gs {
		for _, a := range g.Aliases {
			if len(a) == 0 || len(g.Emoji) == 0 {
				continue
			}
			ei := EmojiInfo{
				Shortcode: a,
				Unicode:   fmt.Sprintf("%+q", strings.ToLower(g.Emoji)),
				Tags:      make([]string, len(g.Tags), cap(g.Tags)),
			}
			copy(ei.Tags, g.Tags)
			eis = append(eis, ei)
		}
	}
	return eis, nil
}
