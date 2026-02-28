package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string
	UnreadCount   int
}

func main() {
	fmt.Println("=================Text Template Sample=================")

	emailTemplate := `
		Subject: {{ .Subject}}

		{{.Body}}

		{{if .Items}}
			Related Items:
		{{range .Items}}
			- {{.}}
		{{end}}
		{{end}}

		{{if gt .UnreadCount 0}}
		You have {{.UnreadCount}} unread emails
		{{else}}
		You have no messages
		{{end}}


		- Thanks
		{{.SenderName}}
	`

	templ, err := template.New("email-message").Parse(emailTemplate)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	data := EmailData{
		RecipientName: "Alex Wonderland",
		SenderName:    "John Doe",
		Subject:       "Hello from Go",
		Body:          "Hello from Go Text Template with few Golang features",
		Items:         []string{"item 1", "item 2", "item 3"},
		UnreadCount:   5,
	}

	var output strings.Builder

	err = templ.Execute(&output, data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(output.String())
}
