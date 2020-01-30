package main

import (
	"bytes"
	"go/format"
	"log"
	"os"
	"text/template"
)

const emojiDataFilePath = "../../internal/data/emojiMap.go"
const templateEmojiData = `
// Package data provide emoji data
package data

// NOTE: THIS FILE WAS CREATED BY 
// EMOJIDATA GENERATION TOOL. (github.com/zajann/emoji/cmd/emojiDataGenerator)
// DO NOT EDIT. IF EDIT THIS FILE THE PROGRAM
// MIGHT NOT WORK PROPERLY.

type EmojiMap map[string]Code

type Code struct {
	Unicode string
	TagMap map[string]struct{}
}

var Emojis = EmojiMap{
	{{- range .Emojies}} 
		":{{.Shortcode}}:": Code{ 
						{{.Unicode}}, 
						map[string]struct{}{ 
							{{- range .Tags}} 
								"{{.}}": struct{}{}, 
							{{- end}}
						},
		},
	{{- end}}
}
`

type EmojiInfo struct {
	Shortcode string
	Unicode   string
	Tags      []string
}

type TemplateData struct {
	PkgName string
	Emojies []EmojiInfo
}

func getAllEmojiData() ([]EmojiInfo, error) {
	emojoInfo, err := getEmojoData()
	if err != nil {
		return nil, err
	}

	gemojiInfo, err := getGemojiData()
	if err != nil {
		return nil, err
	}

	emojoInfo = append(emojoInfo, gemojiInfo...)
	return deDuplicate(emojoInfo), nil
}

func createEmojiDataSource(infos []EmojiInfo) ([]byte, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("template").Parse(templateEmojiData))
	if err := t.Execute(&buf, TemplateData{Emojies: infos}); err != nil {
		return nil, err
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return src, nil
}

func deDuplicate(infos []EmojiInfo) []EmojiInfo {
	keys := make(map[string]struct{})
	list := []EmojiInfo{}
	for _, i := range infos {
		if _, ok := keys[i.Shortcode]; !ok {
			keys[i.Shortcode] = struct{}{}
			list = append(list, i)
		}
	}
	return list
}

func main() {
	emojiInfos, err := getAllEmojiData()
	if err != nil {
		log.Fatalln(err)
	}

	emojiDataSource, err := createEmojiDataSource(emojiInfos)
	if err != nil {
		log.Fatalln(err)
	}

	os.Remove(emojiDataFilePath)
	f, err := os.Create(emojiDataFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if _, err := f.Write(emojiDataSource); err != nil {
		log.Fatalln(err)
	}
}
