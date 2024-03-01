package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/muesli/reflow/wordwrap"
	historyMod "github.com/nicolas-sabbatini/go-choose-your-own-adventure/internal/history"
	"github.com/nicolas-sabbatini/go-choose-your-own-adventure/internal/style"
)

var (
	titleSeparator   = "================================================================================\n"
	optionsSeparator = "--------------------------------------------------------------------------------\n"
)

type model struct {
	currentChapter string
	history        historyMod.History
}

func (self model) changeChapter(nextChapter int) model {
	currentChapter := self.history[self.currentChapter]
	if currentChapter.Title == "" || nextChapter > len(currentChapter.Options) {
		return self
	}
	self.currentChapter = currentChapter.Options[nextChapter].ChapterName
	return self
}

func (self model) Init() tea.Cmd {
	cmd := tea.EnterAltScreen
	return cmd
}

func (self model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return self, tea.Quit
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			nextChapter, err := strconv.ParseInt(msg.String(), 10, 10)
			if err != nil {
				log.Fatal("Error parsing next chapter")
			}
			return self.changeChapter(int(nextChapter - 1)), nil
		}
	}
	return self, nil
}

func (self model) View() string {
	currentChapter := self.history[self.currentChapter]
	if currentChapter.Title == "" {
		return style.Keyword(self.currentChapter) + "\n\nThis chapter do not Exist!!\n\n" + style.Help("Press q to quit")
	}
	story := ""
	for _, paragraph := range currentChapter.Story {
		story += paragraph + "\n\n"
	}
	text := style.Keyword(currentChapter.Title) + "\n" +
		titleSeparator + wordwrap.String(story, 80) + optionsSeparator
	for i, option := range currentChapter.Options {
		text += style.Keyword(fmt.Sprintf("%d", i+1)) + " - " + wordwrap.String(option.Text, 76) + "\n"
	}
	text = text + "\n" + style.Help("Select an option to continue the story") + "\n" + style.Help("Press q to quit")
	return text
}

func main() {
	file, err := os.Create("cyoa.log")
	if err != nil {
		log.Fatal("Error creating log file")
	}
	defer file.Close()
	log.SetOutput(file)

	log.Info("Starting app")

	historyPath := flag.String("historyPath", "assets/history.json", "Path to history")
	flag.Parse()

	var history historyMod.History
	history, err = historyMod.ReadFromFile(*historyPath)
	if err != nil {
		log.Error("Error reading history file")
		log.Fatal(err)
	}

	_, err = tea.NewProgram(model{
		currentChapter: "intro",
		history:        history,
	}).Run()
	if err != nil {
		log.Fatal("Error running program:", err)
	}
}
