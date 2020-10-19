package main

import "fmt"

func main() {
	//DefaultVariables
	var publisher, writer, artist, title string
	var year, pageNumber int32
	var grade float32
	//first comic Variables
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	//Print Statements

	fmt.Println(title, " was written by ", writer, " and illustrated by ", artist, ".", "It was published by", publisher, "in", year, "with", pageNumber, "pages and it's condition is", grade, " out of 10.")

	title = "Epic Vol.1"
	writer = "Ryan N. Shawn"
	artist = "Pheobe Paperclips"
	year = 2013
	pageNumber = 160
	grade = 9.0

	fmt.Println(title, " was written by ", writer, " and illustrated by ", artist, ".", "It was published by", publisher, "in", year, "with", pageNumber, "pages and it's condition is", grade, " out of 10.")

}
