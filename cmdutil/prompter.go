package cmdutil

import "github.com/AlecAivazis/survey/v2"

type Prompter struct {
	Question     string
	DefaultValue string
	Options      []string
	Suggest      func(string) []string
}

func (p Prompter) Confirm() bool {
	ans := false
	prompt := &survey.Confirm{
		Message: p.Question,
	}
	survey.AskOne(prompt, &ans)

	return ans
}

func (p Prompter) Input() string {
	ans := ""
	prompt := &survey.Input{
		Message: p.Question,
		Default: p.DefaultValue,
		Suggest: p.Suggest,
	}
	survey.AskOne(prompt, &ans)

	return ans
}

func (p Prompter) Select() string {
	ans := ""
	prompt := &survey.Select{
		Message: p.Question,
		Options: p.Options,
	}
	survey.AskOne(prompt, &ans)

	return ans
}
