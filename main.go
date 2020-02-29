package main


import(
	"os"
	"net/http"
	"log"
)



var Option = []string{"no filter", "band/artist", "member", "creation date", "first album", "location"}



var fullData []Artist



//////


func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/artist", getArtist)

	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Listen and serve err: ", err)
	}
}