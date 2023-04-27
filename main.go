package main

import (
	"bufio"
	"fmt"
	configs "github.com/Cyber1112/midterm-go/connection"
	"github.com/Cyber1112/midterm-go/seeder"
	"github.com/go-redis/redis"
	"os"
)

func main() {
	db := configs.Connection()

	runSeed(db)

	scanner := bufio.NewScanner(os.Stdin)

	var studentid = ""

	for {
		fmt.Print("Enter Text: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 && len(studentid) == 0 {
			studentid = text
			fmt.Println(studentid)
		} else if text == "view" {
			// print all student grades

		} else if text == "add" {
			fmt.Print("Enter the course id: ")
			course := scanner.Text()
			fmt.Print("Enter the grade: ")
			grade := scanner.Text()
			// add grade to redis

		} else if text == "delete" {
			fmt.Print("Enter the course id: ")
			course := scanner.Text()
			// delete all grades at this course only for the student

		} else if text == "update" {
			fmt.Print("Enter the course id: ")
			course := scanner.Text()
			fmt.Print("Enter the grade: ")
			grade := scanner.Text()
			// update grade for specific course

		}
	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

func runSeed(client *redis.Client) {
	seeder.Seed(client)
}
