package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Page struct that represents a Wiki page. Composed by a title field of type string and a body field
// of type byte slice array. Used to represent the page content into the program memory
type Page struct {
	Title string
	Body  []byte
}

// Parse all the files at the begining into one *Template
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

// Regex to validate the routes of the API
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// Function that receives the title of the page (that matches with the file name stored),
// read that file and creates a pointer to a Page. Loads the page from disk to memory.
// If some error ocurr during the reading, it returns that error, otherwhise the Page pointer is returned
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Function that receives a pointer to a Page and store the page attributes into a txt file.
// It returns an error if there is a problem with the writing.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Function that executes the template with the data of the page identified by title.
// It returns an http 500 error if there is any issue with the template execution, otherwise
// writes the template content into the http.ResponseWriter.
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Wrapper function that validates the received http request. If the route doesnt match with
// the accepted ones, it returns a http 404 error, otherwise, it calls the correspondent page handler.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// Function that handles the view of a page. Its render the view.html template if found, otherwise
// its redirect to a new editing page.
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// Function that handles the edit of a page. Its render the edit.html template if the page exist, otherwise
// its create a new page to edit.
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// Function that saves the current page into disk. Its redirect to the view of that page or returns
// a http 500 status if there is an error on saving.
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Main function, starts all the handlers for view, edit and save packages. Serves the http server.
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
