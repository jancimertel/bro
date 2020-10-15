package _default

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"bitbucket.org/jmertel/bro/templates/types"
	"fmt"
	"os"
)

type defaultTemplate struct {
	provider analyserTypes.IProvider
}

func (d defaultTemplate) Serve(port string) error {
	panic("implement me")
}

func (d defaultTemplate) Build() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}

func NewTemplate(provider analyserTypes.IProvider) types.ITemplater {
	return &defaultTemplate{
		provider,
	}
}

