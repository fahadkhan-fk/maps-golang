package main

import "fmt"

func main() {

	// Map ( collection of `KEY` & `VALUE` pair ).
	// KEY --> All keys must be of the same type.
	// VALUE --> All values must be of the same type.

	// you can create a map according to your requirement like if you want to have a key in integer and it's value in string ( map[int]string{} ) etc.
	// all data types in golang can be used for map key or value types.
	// creating a map having a key & value both in strings

	countryCapital := map[string]string{ // we can declare a map using these `:= ` known as `short-declaration-operator` and assign values along aswell as done below
		"Pakistan": "Islamabad",
		"India":    "New Delhi",
		"Canada":   "Ottawa",
	}
	// you can declare the same map in this way aswell `var countryCapital map[string]string`

	// we can print the map simply from code below (just uncomment it then it will work)
	//fmt.Println(countryCapital)

	// if we want to print the map a bit different then follow below code
	// calling printMap() function to print the map
	printMap(countryCapital)
}

// POINT-TO-PONDER -> we need to define the parameter type while receiving them in a function
// printMap function prints the map along with it's key's and their values
func printMap(countryCapital map[string]string) {
	for country, capital := range countryCapital { // declaring a for loop here that will iterate (range) over a map keys and return keys with values.
		fmt.Println("Capital of", country, "is", capital) // a custom print message of map key-value
	}
}
