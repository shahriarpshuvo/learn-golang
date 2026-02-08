package main

import (
	"fmt"
	"time"
)

func main() {
	temp := 10
	too_hot := false
	too_cold := false

	if !too_hot || !too_cold {
		if temp > 40 {
			fmt.Println("Too hot, stay inside!")
		} else if temp <= 40 && temp >= 15 {
			fmt.Println("Comfortable, go explore!")
		} else if temp < 15 && temp > 0 {
			fmt.Println("Cold, wear Jacket!")
		} else {
			fmt.Println("Don't even think about it")
		}
	}

	if age := 20; age >= 18 {
		fmt.Println("Permitted!")
	} else if age < 18 {
		fmt.Println("Not Permitted!")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend!")
	case time.Monday:
		fmt.Println("Weekday, Start of the week!")
	case time.Friday:
		fmt.Println("Weekday, End of the week!")
	default:
		fmt.Println("Weekday!")
	}

	var value any = 15
	switch v := value.(type) {
	case string:
		fmt.Println("Value is a string:", v)
	case int:
		fmt.Println("Value is an integer:", v)
	case float64:
		fmt.Println("Value is a float:", v)
	default:
		fmt.Println("Value is of unknown type")
	}
}
