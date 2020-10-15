package main

import (
	"bitbucket.org/jmertel/bro/analyser"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	projectPath := os.Args[1]
	logrus.Infof("Processing project in: %s", projectPath)
	analyser := analyser.NewProjectAnalyser(projectPath)
	analyser.Process()
}