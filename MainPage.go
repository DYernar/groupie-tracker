package main

import(
	"fmt"
	"net/http"
	"html/template"
	"time"
	"strconv"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "GET" {

			GetAll(w, r )
			

			allData := fullData
			var returnValue RetVal
			if len(allData) == 0 {
				time.Sleep(2*time.Second)
				allData = fullData
			}

			
			returnValue.Artists = allData
			returnValue.Option = Option
			returnValue.Locations = GetAllLocations()
			t, tempErr := template.ParseFiles("static/index.html")
			if tempErr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
				return;
			}
			t.Execute(w, returnValue)

		} else if r.Method == "POST" {
			r.ParseForm()
			search := r.FormValue("searchText")
			searchType := r.FormValue("searchType")
			searchResult := GetByHint(search, searchType)


			creationStart := r.FormValue("cr_start_date")
			creationEnd := r.FormValue("cr_end_date")
			faStart := r.FormValue("fa_start_date")
			faEnd := r.FormValue("fa_end_date")
			membernum1 := r.FormValue("membernum1")
			membernum2 := r.FormValue("membernum2")
			var locationFilter []string
			for i := 0; i <193; i++ {
				locationFilter = append(locationFilter, r.FormValue("locations"+strconv.Itoa(i)))
			}
			
			searchResult=ApplyFilters(fullData, creationEnd, creationStart, faStart, faEnd, membernum1, membernum2)
			
			var searchReturn RetVal
			searchReturn.Artists = searchResult
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
