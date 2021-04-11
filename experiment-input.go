package main

import "fmt"

func main() {

	yesNo := "1"
	var launch bool

	switch yesNo {
	case "yes", "1", "true":
		launch = true
	case "No", "0", "false":
		launch = false
	default:
		fmt.Println("error")
	}
	fmt.Println("Ready for launch", launch)

}
