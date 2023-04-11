package advanced

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelSelectErrorHandling(t *testing.T) {

	/* In this example, we define a worker function that takes an integer job from
	the jobs channel and performs some work on it. If the job is even, it sends the
	result on the results channel. If the job is odd, it sends an error on the
	errors channel. */

	worker := func(id int, jobs <-chan int, results chan<- int, errors chan<- error) {

		/* Notice the signature of this function. '<-chan int' is an int receive-only
		channel type, meanwhile 'chan<- int' is an int send-only channel type. These are
		used to prevent misuse and fail at compilation time */

		for job := range jobs {
			fmt.Printf("Worker %d starting job %d\n", id, job)
			time.Sleep(time.Second)
			if job%2 == 0 {
				results <- job * 2
			} else {
				errors <- fmt.Errorf("job %d failed", job)
			}
		}
	}

	/* We create the channels and start the workers. */
	numWorkers := 3
	numJobs := 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	errors := make(chan error, numJobs)

	/* We then queue up the jobs and wait for all jobs to complete. */

	// Start the workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results, errors)
	}

	// Queue up the jobs
	for i := 1; i <= numJobs; i++ {

		/* To queue up jobs in Go channels means to use channels to create a queue of
		tasks or jobs that need to be executed concurrently by multiple goroutines. */

		jobs <- i
	}
	close(jobs)

	// Wait for all jobs to complete
	numResults := 0
	numErrors := 0

	/* We use a select statement to wait for either a result or an error, and a
	switch statement to handle each case. */

	for numResults+numErrors < numJobs {
		select {
		case result := <-results:
			fmt.Printf("Result: %d\n", result)
			numResults++
		case err := <-errors:
			fmt.Printf("Error: %s\n", err)
			numErrors++
		}
	}

	/* If we receive a result, we print it to the console and increment the numResults
	variable. If we receive an error, we print it to the console and increment the
	numErrors variable.
	We continue to wait for results and errors until we have received them all. Once
	all jobs have completed, we can exit the program. */
}
