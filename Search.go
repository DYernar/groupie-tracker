package main

import(
	"fmt"
	"strings"
	"strconv"
)

func ArrContains(arr []Artist, artist Artist) bool {
	for _, art := range arr {
		if art.Name == artist.Name {
			return true
		}
	}
	return false
}


func GetByHint(hint string, searchType string) []Artist {
	fmt.Print(searchType)

	var returnList []Artist

	if searchType == "no filter" {
		for _, artist := range fullData {
			if strings.HasPrefix(artist.Name, hint) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
			if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(hint)) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
			for _, member := range artist.Members {
				if strings.HasPrefix(member, hint) {
					if !ArrContains(returnList, artist){
						returnList = append(returnList, artist)
					}
				}
				if strings.HasPrefix(strings.ToLower(member), strings.ToLower(hint)) {
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
	} else if searchType == "band/artist" {

		for _, artist := range fullData {
			if strings.HasPrefix(artist.Name, hint) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
			if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(hint)) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
		}
	} else if searchType == "member" {
		for _, artist := range fullData {

		for _, member := range artist.Members {
			if strings.HasPrefix(member, hint) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
			if strings.HasPrefix(strings.ToLower(member), strings.ToLower(hint)) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
		}
	}
	} else if searchType == "creation date" {
		for _, artist := range fullData {
			if strings.Contains(strconv.Itoa(artist.CreationDate), hint) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
		}
	} else if searchType == "first album" {
		for _, artist := range fullData {
			if strings.Contains(artist.FirstAlbum, hint) {
				if !ArrContains(returnList, artist){
					returnList = append(returnList, artist)
				}
			}
		}
	} else if searchType == "location" {
		for _, artist := range fullData {
			for _, location := range artist.Locs.Locations {
				if strings.Contains(location, hint) {
					if !ArrContains(returnList, artist){
						returnList = append(returnList, artist)
					}
				}
			}
		}
	}

	return returnList
}
