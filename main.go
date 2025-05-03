package main

import (
	"fmt"
	"html/template"
)

func init() {
	tmpl, _ := template.ParseGlob("templates/*.html")
}

func main() {
	fmt.Println("main")
	//db.InitDB()
}
