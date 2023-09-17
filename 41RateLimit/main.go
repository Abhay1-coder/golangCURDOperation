/*
*
In simple terms, rate limiting in Go (or any programming language) is like controlling the speed at which a process or task can run. It helps prevent overloading resources or services by limiting how often you can perform a certain action.

Imagine you're at an all-you-can-eat buffet:

- **No Rate Limiting:** If there's no rate limiting, you can keep grabbing food as fast as you can eat. This might lead to overeating or wastage.

- **With Rate Limiting:** If there's a rate limit of, let's say, one plate of food every 15 minutes, you can only get a new plate of food after 15 minutes have passed. This prevents you from overeating and ensures everyone has a fair chance to eat.

In Go, rate limiting is used to control how often you can perform actions like making HTTP requests, sending messages, or processing data, to prevent overwhelming external services, APIs, or your own system. It helps maintain a controlled and reasonable pace of work.
*
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// Define a function to simulate making an HTTP request
	makeAPIRequest := func(requestNumber int) {
		fmt.Printf("Sending request %d to the API...\n", requestNumber)
		// Simulate an API response delay
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Received response for request %d\n", requestNumber)
	}

	// Perform 5 API requests without rate limiting
	for i := 1; i <= 5; i++ {
		makeAPIRequest(i)
	}

	// Introduce rate limiting - 2 requests per second
	rateLimit := time.Tick(500 * time.Millisecond) // Limit to 2 requests per second

	fmt.Println("Rate-limited requests:")
	for i := 6; i <= 10; i++ {
		<-rateLimit // Wait for the rate limit interval
		makeAPIRequest(i)
	}
}
