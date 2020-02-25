package main


import(
	"fmt"
	"os"
	"net/http"
	"html/template"
	"log"
	"io/ioutil"
	"encoding/json"
	// includes "groupie-tracker/includes"
	"time"
	"strings"
	"strconv"
)


type RetVal struct {
	Artists []Artist
}

type Artist struct{
	ID int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum`
	Locs Locations
	ConDates ConcertDates
	Rels Relation
}


type Index struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	ID int `json:"id"`
	Locations  []string `json:"locations"`
	Dates ConcertDates
}

type DateIndex struct {
	Index []ConcertDates `json:"index"`
}

type ConcertDates struct {
	ID int `json:"id"`
	Dates []string `json:"dates"`
}


type RelationIndex struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations`
}





var fullData []Artist

func GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	var allData []Artist

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return ;
	}
	defer resp.Body.Close()
	
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return ;
	}
	json.Unmarshal([]byte(body), &allData)



	//Getting the locations 

	locations, err2 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err2 !=  nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return ;
	}


	locBody, locErr := ioutil.ReadAll(locations.Body)
	if locErr != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return ;
	}

	var allLocations Index
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
		return ;
	}

	concBody, dateErr2 := ioutil.ReadAll(concertDates.Body)
	if dateErr2 != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
		return ;
	}


	var allDates DateIndex
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
			return ;
		}

		relBody, relErr2 := ioutil.ReadAll(rels.Body)
		if relErr2 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "<h1>500 Internal server Error!</h1>")
			return ;
		}

		var relation RelationIndex
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




	fullData = allData
}

////
func ArrContains(arr []Artist, artist Artist) bool {
	for _, art := range arr {
		if art.Name == artist.Name {
			return true
		}
	}
	return false
}

///---------------search result////
func GetByHint(hint string) []Artist {
	var returnList []Artist
	for _, artist := range fullData {
		if strings.Contains(artist.Name, hint) {
			if !ArrContains(returnList, artist){
				returnList = append(returnList, artist)
			}
		}
		for _, member := range artist.Members {
			if strings.Contains(member, hint) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
		}

		for _, location := range artist.Locs.Locations {
			if strings.Contains(location, hint) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
		}

		if strings.Contains(strconv.Itoa(artist.CreationDate), hint) {
			if !ArrContains(returnList, artist){
				returnList = append(returnList, artist)
			}
		}
		if strings.Contains(artist.FirstAlbum, hint) {
			if !ArrContains(returnList, artist){
				returnList = append(returnList, artist)
			}
		}
	}
	return returnList
}
/////////////////////
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
			searchReasult := GetByHint(data)
			var searchReturn RetVal
			searchReturn.Artists = searchReasult
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