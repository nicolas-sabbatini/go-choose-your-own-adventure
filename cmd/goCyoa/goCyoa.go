package main

import (
	"flag"
	"log"
	"net/http"

	historyMod "github.com/nicolas-sabbatini/go-choose-your-own-adventure/internal/history"
	"github.com/nicolas-sabbatini/go-choose-your-own-adventure/internal/htmlGenerator"
)

var history historyMod.History

func main() {
	historyPath := flag.String("historyPath", "assets/history.json", "Path to history")
	chapterTemplatePath := flag.String("chapterTemplatePath", "assets/historyTemplate.html", "Path to chapter template")
	indexTemplatePath := flag.String("indexTemplatePath", "assets/indexTemplate.html", "Path to index template")
	flag.Parse()

	var err error
	history, err = historyMod.ReadFromFile(*historyPath)
	if err != nil {
		log.Print("Error reading history file")
		log.Panic(err)
	}

	err = htmlGenerator.LoadTemplates(*chapterTemplatePath, *indexTemplatePath)
	if err != nil {
		log.Print("Error loading templates")
		log.Panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", info)
	log.Print("App lisening to port 8080")
	http.ListenAndServe(":8080", handleChapter(mux))
}

func info(response http.ResponseWriter, request *http.Request) {
	log.Print("Hit to route `/`")
	htmlGenerator.IndexTemplate.Execute(response, historyMod.GetChapterNames(history))
}

func handleChapter(fallback http.Handler) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		log.Printf("Hit to route `%s`", request.URL.Path)
		chapterName := request.URL.Path[1:]
		chapter := history[chapterName]
		if chapter.Title == "" {
			fallback.ServeHTTP(response, request)
		}
		htmlGenerator.ChapterTemplate.Execute(response, history[chapterName])
	}
}
