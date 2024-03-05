package dsutils

import "fmt"

func Must0(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic(err)
	}
}

func Must1[T any](result T, err error) T {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		panic(err)
	}

	return result
}
