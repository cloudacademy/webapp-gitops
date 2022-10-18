package main

import (
	"html/template"
	"net/http"
	"os"
)

type config struct {
	Message         string
	BackgroundColor string
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CloudAcademy Kustomize Lab</title>
</head>
<link href="https://fonts.googleapis.com/css2?family=Montserrat&display=swap" rel="stylesheet">

<style>
body {
  position: relative;
  height: 100vh;
  margin: 0;
  background-color: {{.BackgroundColor}};
}

h1 {
 position: absolute;
 top: 40%;
 transform: translateY(-50%);
 width: 100%;
 text-align: center;
 margin: 0;
 font-family: 'Montserrat', sans-serif;
}
</style>

<body>
<h1>{{.Message}}</h1>
</body>
</html>
`

func main() {
	message := os.Getenv("MESSAGE")
	bgcolor := os.Getenv("BACKGROUND_COLOR")

	tmpl := template.Must(template.New("main").Parse(htmlTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := config{
			Message:         message,
			BackgroundColor: bgcolor,
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
