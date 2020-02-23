package includes

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


