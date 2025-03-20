package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("../static/",http.FileServer(http.Dir("static/")))
	http.HandleFunc("/", home)

	http.HandleFunc("/projects",projects)

	http.HandleFunc("/skills",skills)
	http.HandleFunc("/about",about)
	http.HandleFunc("/contact",contact)


	fmt.Println("Server is running on http://localhost:8080")

	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal((err))
	}
}

func home(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "home.html")
}
func projects(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w,"projects.html")
}
func skills(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "skills.html")
}
func about(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "about.html")
}
func contact(w http.ResponseWriter, r *http.Request)  {
	renderTemplate(w, "contact.html")
}


func renderTemplate(w http.ResponseWriter, templ string){
	t , err := template.ParseFiles("templates/"+templ)// Parse the specified template file
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w,nil)// Execute the template and write it to the ResponseWriter
}