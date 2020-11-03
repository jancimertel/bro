package main

import (
	"bitbucket.org/jmertel/bro/analyser"
	markdownTemplate "bitbucket.org/jmertel/bro/templates/markdown"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	projectPath := os.Args[1]
	logrus.Infof("Processing project in: %s", projectPath)

	// analysis
	analyserInst := analyser.NewProjectAnalyser(projectPath)
	analyserInst.Process()

	// output
	template := markdownTemplate.NewTemplate(&analyserInst)
	if err := template.Build(); err != nil {
		log.Fatalf("Could not build template: %v", err)
	}
}
