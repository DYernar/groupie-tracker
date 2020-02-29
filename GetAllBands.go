package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

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
