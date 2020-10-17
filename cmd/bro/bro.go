package main

import (
	"bitbucket.org/jmertel/bro/analyser"
	defaultTemplate "bitbucket.org/jmertel/bro/templates/text"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	projectPath := os.Args[1]
	logrus.Infof("Processing project in: %s", projectPath)

	// analysis
	analyser := analyser.NewProjectAnalyser(projectPath)
	analyser.Process()

	// output
	template := defaultTemplate.NewTemplate(&analyser)
	if err := template.Build(); err != nil {
		log.Fatalf("Could not build template: %v", err)
	}
}