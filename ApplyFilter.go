package main

import (
	"strconv"
	"fmt"
)

func ApplyFilters(searchResult []Artist, creationEnd string, creationStart string, faStart string, faEnd string, membernum1 string, membernum2 string) []Artist {
	var retVal []Artist
	memberAmount, _ := strconv.Atoi(membernum1)
	memberAmount2, _ := strconv.Atoi(membernum2)
	if memberAmount > memberAmount2 {
		temp := memberAmount2
		memberAmount2 = memberAmount
		memberAmount = temp
	}
	fmt.Println(memberAmount2)
	fmt.Println(memberAmount)
	for _, band := range searchResult  {
		if (len(band.Members) <= memberAmount2 && len(band.Members) >= memberAmount) {
			retVal = append(retVal, band)
			// searchResult = append(searchResult[:index], searchResult[index+1:]...)
		}
	}
	return retVal
}