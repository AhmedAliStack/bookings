package helpers

import (
	"fmt"
	"sort"
)

func CreateSlices(){
	var mySlices []string

	mySlices = append(mySlices, "Ibrahim")
	mySlices = append(mySlices, "Ahmed")
	mySlices = append(mySlices, "Yoused")
	mySlices = append(mySlices, "Maian")

	sort.Strings(mySlices)

	fmt.Println("MySlices are",mySlices)
	fmt.Println("Print Range",mySlices[1:2])
}