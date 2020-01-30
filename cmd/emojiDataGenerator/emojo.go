package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const EmojoJsonURL = "https://raw.githubusercontent.com/CodeFreezr/emojo/master/db/v5/emoji-v5.json"

// Emojo emoji-v5 struct deserialized json
type Emojo struct {
	No          int    `json:"No"`
	Emoji       string `json:"Emoji"`
	Category    string `json:"Category"`
	SubCategory string `json:"SubCategory"`
	Unicode     string `json:"Unicode"`
	Name        string `json:"Name"`
	Tags        string `json:"Tags"`
	Shortcode   string `json:"Shortcode"`
}

func getEmojoData() ([]EmojiInfo, error) {
	res, err := http.Get(EmojoJsonURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	emojiJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var emos []Emojo
	if err := json.Unmarshal(emojiJson, &emos); err != nil {
		return nil, err
	}

	var eis []EmojiInfo
	for _, emo := range emos {
		if len(emo.Shortcode) == 0 || len(emo.Emoji) == 0 {
			continue
		}
		ei := EmojiInfo{
			Shortcode: strings.Replace(emo.Shortcode, ":", "", 2),
			Unicode:   fmt.Sprintf("%+q", strings.ToLower(emo.Emoji)),
			Tags:      getEmojoTags(emo.Tags),
		}
		eis = append(eis, ei)
	}
	return eis, nil
}

func getEmojoTags(s string) []string {
	var tags []string

	for _, tagWithSpace := range strings.Split(s, "|") {
		tags = append(tags, strings.TrimSpace(tagWithSpace))
	}
	return tags
}
