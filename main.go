package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Println("🫂 Hello, World!")
	fmt.Println("----------------")
	fmt.Printf("version: %s\n", version)
	fmt.Printf("commit: %s\n", commit)
	fmt.Printf("date: %s\n", date)
}
