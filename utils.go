package main

import (
	"./data"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/phachon/go-logger"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address      string
	ReadTimeOut  int64
	WriteTimeOut int64
	Static       string
}

var config Configuration
var logger = go_logger.NewLogger()

func init() {
	setLogger()
	loadConfig()
}

func setLogger() {
	// TODO: set logger parameters here
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		logger.Error("config json read failed!")
		os.Exit(2)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		logger.Error("config json parse failed!")
		os.Exit(2)
	}
	logger.Debugf("config json parse: \t%s", config)
}

func error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)

}

// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
