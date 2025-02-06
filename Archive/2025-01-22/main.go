package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func main() {
	// parse the template
	tpl = template.Must(template.ParseFiles("goindex.gohtml", "head.gohtml", "nav.gohtml", "article.gohtml", "footer.gohtml"))
	// create file index.html and write to index.html
	file, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tpl.Execute(file, nil)
	//parse the template
	tpl = template.Must(template.ParseFiles("goarchive.gohtml", "head.gohtml", "nav.gohtml", "footer.gohtml"))
	// create file archive.html and write to archive.html
	file, err = os.Create("archive.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	type archive struct {
		Version string;
		Directory string;
	}
	archives := []archive{
		archive{"Version 1", "Archive/2025-01-19/index.html"},
		//archive{"Version 2", "archive2"},
		//archive{"Version 3", "archive3"},
	}
	tpl.Execute(file, archives)
	//parse the template
	tpl = template.Must(template.ParseFiles("goabout.gohtml", "head.gohtml", "nav.gohtml", "section.gohtml", "footer.gohtml"))
	// create file about.html and write to about.html
	file, err = os.Create("about.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tpl.Execute(file, nil)
}