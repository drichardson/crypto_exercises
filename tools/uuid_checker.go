package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {

	collisions := make(map[uuid.UUID]int)
	collisionCount := 0
	summaryEvery := 10000000
	summaryCounter := 0
	errors := 0
	var count int64

	for {
		u, err := uuid.NewRandom()

		if err == nil {
			if count, ok := collisions[u]; ok {
				collisions[u] = count + 1
				collisionCount++
				fmt.Printf("%d collisions for UUID %v\n", collisions[u], u)
			} else {
				collisions[u] = 0
			}
		} else {
			fmt.Println("Error: %v\n", err)
		}

		count++
		summaryCounter++
		if summaryCounter >= summaryEvery {
			summaryCounter = 0
			fmt.Printf("Summary: count=%v, errors=%v, collisions=%v\n", count, errors, collisionCount)
		}
	}

}
