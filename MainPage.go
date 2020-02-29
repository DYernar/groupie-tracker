package main

import(
	"fmt"
	"net/http"
	"html/template"
	"time"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "GET" {

			go GetAll(w, r )
			

			allData := fullData
			var returnValue RetVal
			if len(allData) == 0 {
				time.Sleep(2*time.Second)
				allData = fullData
			}

			
			returnValue.Artists = allData
			returnValue.Option = Option
			t, tempErr := template.ParseFiles("static/index.html")
			if tempErr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
				return;
			}
			t.Execute(w, returnValue)

		} else if r.Method == "POST" {
			r.ParseForm()
			data := r.FormValue("searchText")
			searchType := r.FormValue("searchType")
			searchReasult := GetByHint(data, searchType)
			var searchReturn RetVal
			searchReturn.Artists = searchReasult
			searchReturn.Option = Option
			t, tempErr := template.ParseFiles("static/search.html")
			if tempErr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
				return;
			}
			t.Execute(w, searchReturn)
			
		} else {
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(400)
				fmt.Fprintf(w, "<h1>400 Bad Request!</h1>")
		}
	} else {
		w.Header().Set("Content-Type", "text/html")
		t, err := template.ParseFiles("static/error404.html")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
			return;
		}
		w.WriteHeader(404)
		t.Execute(w, nil)

	}

}
