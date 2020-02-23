package main


import(
	"fmt"
	"os"
	"net/http"
	"html/template"
	"log"
	"io/ioutil"
	"encoding/json"
	includes "groupie-tracker/includes"
)

func GetAll(w http.ResponseWriter, r *http.Request) []includes.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	var allData []includes.Artist

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return allData;
	}
	defer resp.Body.Close()
	
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return allData;
	}
	json.Unmarshal([]byte(body), &allData)



	//Getting the locations 

	locations, err2 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err2 !=  nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return allData;
	}


	locBody, locErr := ioutil.ReadAll(locations.Body)
	if locErr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return allData;
	}

	var allLocations includes.Index
	json.Unmarshal([]byte(locBody), &allLocations)
	////////////////////////

	//relating the location with the artist
	for i := 0; i < len(allData); i++ {
		for _, loc := range allLocations.Index {
			if loc.ID == allData[i].ID {
				allData[i].Locs = loc
				continue
			}
		}
	}  
	//location end

	/// getting the concert dates
	concertDates, dateErr := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if dateErr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return allData;
	}

	concBody, dateErr2 := ioutil.ReadAll(concertDates.Body)
	if dateErr2 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return allData;
	}


	var allDates includes.DateIndex
	json.Unmarshal([]byte(concBody), &allDates)
	
	for i := 0; i < len(allData); i++ {
		for _, date := range allDates.Index {
			if date.ID == allData[i].ID {
				allData[i].ConDates = date
				continue
			}
		}
	}  

	//dates end

	//getting relations
		rels, relErr := http.Get("https://groupietrackers.herokuapp.com/api/relation")
		if relErr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
			return allData;
		}

		relBody, relErr2 := ioutil.ReadAll(rels.Body)
		if relErr2 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
			return allData;
		}

		var relation includes.RelationIndex
		json.Unmarshal([]byte(relBody), &relation)
		
		for i := 0; i < len(allData); i++ {
			for _, relation := range relation.Index {
				if relation.ID == allData[i].ID {
					allData[i].Rels = relation
					continue
				}
			}
		}  

	//relations end




	return allData
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == "GET" {

			allData := GetAll(w, r )
			
			t, tempErr := template.ParseFiles("static/index.html")
			if tempErr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
				return;
			}
			var returnValue includes.RetVal
			returnValue.Artists = allData
			t.Execute(w, returnValue)

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


func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}
	http.HandleFunc("/", mainPage)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Listen and serve err: ", err)
	}
}