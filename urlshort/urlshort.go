package urlshort

import (
	"fmt"
	"html/template"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("homePage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Welcome",
		Message: "This is the Home Page",
	}
	t.Execute(w, data)
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Call ParseForm() to parse the raw query and update r.Form and r.PostForm.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	// Use FormValue() to get the values from the form input
	userInput := r.FormValue("user_input")
	var bear string = calcDate(userInput)
	fmt.Fprintf(w, "Day is %s", bear)

	fmt.Fprintf(w, "You entered: %s", userInput)
}

func Serve() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/submit_form", formHandler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
