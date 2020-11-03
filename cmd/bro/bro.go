package main

import (
	"bitbucket.org/jmertel/bro/analyser"
	markdownTemplate "bitbucket.org/jmertel/bro/templates/markdown"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		logrus.Fatalf("Please provide arguments: <project path> <out path>")
	}

	projectPath := os.Args[1]
	logrus.Infof("Processing project in: %s", projectPath)

	// analysis
	analyserInst := analyser.NewProjectAnalyser(projectPath)
	analyserInst.Process()

	// output
	outPath := os.Args[2]
	logrus.Infof("Generating output to: %s", outPath)

	template := markdownTemplate.NewTemplate(&analyserInst)
	if err := template.Build(outPath); err != nil {
		log.Fatalf("Could not build template: %v", err)
	}
}
