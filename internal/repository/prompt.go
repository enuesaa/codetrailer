package repository

import (
	"text/template"

	"github.com/erikgeiser/promptkit/textinput"
)

type PromptRepositoryInterface interface {
	Ask(message string, defaultValue string) (string, error)
	Render(defaultValue string, render func (s string) (string, bool)) (string, error)
}

type PromptRepository struct {}

func (repo *PromptRepository) Ask(message string, defaultValue string) (string, error) {
	input := textinput.New(message)
	input.InitialValue = defaultValue
	input.Validate = nil
	input.Template = `
	{{- Bold .Prompt }} {{ .Input -}}
	{{- if .ValidationError }} {{ Foreground "1" (Bold "✘") }}
	{{- end -}}
	`

	return input.RunPrompt()
}

func (repo *PromptRepository) Render(defaultValue string, render func (s string) (string, bool)) (string, error) {
	input := textinput.New("")
	input.InitialValue = defaultValue
	input.Validate = nil
	input.Template = `{{ Render .Input }}`
	input.ResultTemplate = `{{ ResultRender .FinalValue }}`

	input.ExtendedTemplateFuncs = template.FuncMap{
		"Render": func (s string) string {
			result, _ := render(s)
			return result
		},
		"ResultRender": func (s string) string {
			result, persist := render(s)
			if persist {
				return result
			}
			return ""
		},
	}

	return input.RunPrompt()
}
