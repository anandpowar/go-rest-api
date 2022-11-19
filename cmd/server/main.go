package main

import "fmt"

// initiation and startup
func Run() error {
	fmt.Println("starting the application")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
