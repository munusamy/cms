package main

import (
	"github.com/munusamy/cms"
	"os"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, world!",
		Content: "This is the body of our webapge",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
