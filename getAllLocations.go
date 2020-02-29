package main

func ContainsLocation(allLocations []string ,location string) bool {
	for _, item := range allLocations {
		if item == location {
			return true
		}
	}
	return false
}

func GetAllLocations() []string {
	var allLocations []string
	for _, band := range fullData {
		for _, location := range band.Locs.Locations {
			if !ContainsLocation(allLocations, location) {
				allLocations = append(allLocations, location)
			}
		}
	}	
	return allLocations
}