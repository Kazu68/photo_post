package main

import(
"templates/index.html"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]interface{}{"Title": "index"}
    renderTemplate(w, "index", data)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    if err := templates.ExecuteTemplate(w, tmpl+".html", data); err != nil {
        log.Fatalln("Unable to execute template.")
    }
}

func main() {
    http.HandleFunc("/", IndexHandler)
    http.ListenAndServe(":8888", nil)
}
