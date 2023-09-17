/*
*In simple terms, a worker pool in Go is like having a group of workers (goroutines) ready to do tasks concurrently.
Think of it as a team of workers in a factory. Here's a straightforward explanation:

Workers: You have a fixed number of workers (goroutines) who are waiting to do jobs.
These workers are like skilled employees ready to perform tasks.

Queue of Jobs: There's a queue of jobs that need to be done. These jobs can be any kind of work or task that your program needs to complete.

Task Distribution: The worker pool manages the distribution of jobs to the workers. When a job arrives,

	it assigns it to an available worker to complete.

Concurrency: Multiple workers can work on jobs simultaneously.
This means that several tasks can be done at the same time, making your program more efficient.

Controlled Resources: Worker pools help control the number of concurrent tasks. You can specify

	how many workers are in the pool, preventing your program from overwhelming your system with too many tasks at once.

	*
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Number of workers (goroutines) in the pool
	numWorkers := 3

	// Create a job channel to send download tasks to workers
	jobs := make(chan string, 5)

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Define a list of files to download
	filesToDownload := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"}

	// Send download tasks to the job channel
	for _, file := range filesToDownload {
		jobs <- file
	}

	// Close the job channel to indicate that no more tasks will be sent
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
}

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range jobs {
		// Simulate downloading by sleeping for a short time
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d downloaded %s\n", id, file)
	}
}

/**
We define the number of workers (numWorkers) in our worker pool, which is 3 in this case.

We create a channel (jobs) to send download tasks (file names) to the worker pool.

We start the worker goroutines by calling the worker function in a loop, passing the worker's ID, the jobs channel, and the WaitGroup.

We define a list of files to download (filesToDownload).

We send download tasks (file names) to the jobs channel.

After sending all tasks, we close the jobs channel to indicate that no more tasks will be sent.

Each worker goroutine receives download tasks from the jobs channel, simulates the download by sleeping for a short time, and then prints a message indicating which file it downloaded.

Finally, we wait for all worker goroutines to finish using wg.Wait().
**/
