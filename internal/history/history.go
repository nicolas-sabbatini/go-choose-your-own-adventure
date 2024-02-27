package history

import (
	"encoding/json"
	"os"
)

type Option struct {
	Text        string `json:"text"`
	ChapterName string `json:"arc"`
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type History map[string]Chapter

func ReadFromFile(path string) (History, error) {
	historyFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	history := make(map[string]Chapter)
	err = json.Unmarshal(historyFile, &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func GetChapterNames(history History) []string {
	var topics []string

	for key := range history {
		topics = append(topics, key)
	}

	return topics
}
