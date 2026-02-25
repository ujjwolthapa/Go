package main

import (
	"fmt"
	"math/rand"
	"time"
)

type serverStatus struct {
	Name string
	IsUp bool
}

func checkServer(server string, result chan serverStatus) {
	time.Sleep(2 * time.Second)
	isUp := rand.Intn(2) == 1
	result <- serverStatus{Name: server, IsUp: isUp}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Monitoring started...")

	result := make(chan serverStatus)

	servers := []string{
		"server-1",
		"server-2",
		"server-3",
		"server-4",
		"server-5",
	}

	for {
		fmt.Println("\nRunning health checks...")

		// Start checks
		for _, server := range servers {
			go checkServer(server, result)
		}

		// Collect results
		for i := 0; i < len(servers); i++ {
			r := <-result

			if r.IsUp {
				fmt.Printf("%s is UP\n", r.Name)
			} else {
				fmt.Printf("%s is DOWN\n", r.Name)
			}
		}

		time.Sleep(5 * time.Second) // wait before next cycle
	}
}
