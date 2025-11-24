package main

import (
	"fmt"
	"log"

	"github.com/gocraft/dbr/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open Connection
	conn, err := dbr.Open("sqlite3", "./mumbles.db", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Query mumbles
	var mumbles []Mumble
	sess := conn.NewSession(nil)

	// Get a mumble by ID
	mumble := getMumble(sess, 5)
	fmt.Println("Likes with ID 5:", mumble.Likes)

	// Add a like to the mumble
	addLike(sess, 5)
	addLike(sess, 5)

	// Create a new mumble
	createMumbles(sess, "Shihoko", "I had a great weekend!", 5, 1625159072)

	// Delete a mumble
	deleteMumble(sess, 2)

	mumbles = getMumbles(sess)

	// Print results
	for index, m := range mumbles {
		fmt.Println(index, m)
	}
}
