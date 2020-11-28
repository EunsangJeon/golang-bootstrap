package main

import "fmt"

func basics() {
	var purpose string = "test"
	autoPurpose := "test"

	fmt.Println(purpose)
	fmt.Println(autoPurpose)

	purpose = itReturns()

	sliceOfStrings := []string{"one", "two", itReturns()}
	fmt.Println(sliceOfStrings)
	sliceOfStrings = append(sliceOfStrings, "one more")
	fmt.Println(sliceOfStrings)

	for i, someString := range sliceOfStrings {
		fmt.Println(i, someString)
	}
}

func itReturns() string {
	return "dev"
}
