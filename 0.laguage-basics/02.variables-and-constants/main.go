package main

import "fmt"

func main() {
	// Variables
	var first_name string = "Shahriar"
	var last_name = "Parvez" // infer <- go automatically assign
	age := 12                //shorthand
	age = 28                 // redeclare or update

	// Constants
	const nid = "BD162831231212"
	// nid = "BD162831231878" // updating will throw errors

	// Multiple Declation once, require one variables = each line
	const (
		race     = "South Asian"
		birthday = "11 July"
	)

	fmt.Println("Name: ", first_name, last_name+",", "Age: ", age, "NID: ", nid)
	fmt.Println("Race: ", race, "Brithday: ", birthday)
}
