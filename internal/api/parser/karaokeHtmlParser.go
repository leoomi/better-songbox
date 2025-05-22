package parser

import (
	"fmt"
	"io"

	"github.com/leoomi/better-songbox/internal/models"
	"golang.org/x/net/html"
)

func ParseKaraokeSearch(reader io.Reader) []models.Song {
	songs := []models.Song{}
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			if tokenizer.Err().Error() == "EOF" {
				break
			}
			panic("error parsing! " + tokenizer.Err().Error())
		}

		token := tokenizer.Token()

		if token.Type == html.StartTagToken && token.Data == "tr" && isSearchRow(token) {
			fmt.Printf("Token: %v %v %v %v\n", token, token.Type.String(), token.Data, token.Attr)
			songs = append(songs, parseSearchRow(tokenizer))
		}
	}

	return songs
}

func isSearchRow(token html.Token) bool {
	attrs := token.Attr

	if len(attrs) == 0 {
		return false
	}

	for _, v := range attrs {
		if v.Key == "class" && v.Val == "SearchRow" {
			return true
		}
	}

	return false
}

func parseSearchRow(tokenizer *html.Tokenizer) models.Song {
	song := models.Song{}
	tokenizer.Next()
	token := tokenizer.Token()
	tdCount := 0

	for tokenizer.Token().Type != html.EndTagToken || token.Data != "tr" {
		if token.Type == html.StartTagToken && token.Data == "td" {
			text := getNextText(tokenizer)
			switch tdCount {
			case 0:
				song.Singer = text
			case 1:
				song.Name = text
			case 2:
				song.Code = text
			}

			tdCount++
		}

		tokenizer.Next()
		token = tokenizer.Token()
	}

	return song
}

func getNextText(tokenizer *html.Tokenizer) string {
	token := tokenizer.Token()
	for token.Type != html.TextToken {
		tokenizer.Next()
		token = tokenizer.Token()
	}

	return token.Data
}
