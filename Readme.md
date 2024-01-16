### Mini Project: Concurrent API Data Aggregator

#### Problem Statement:

Create a program that concurrently fetches data from multiple APIs, processes the data, and aggregates the results.

#### Requirements:

1. Define a list of API endpoints or URLs from which you want to fetch data.
2. Create a Goroutine for each API endpoint to fetch data concurrently. You can use the `net/http` package for making HTTP requests.
3. Use a channel to send the fetched data from each Goroutine to the main Goroutine.
4. Implement a function to process the fetched data. This could include parsing JSON, extracting relevant information, or performing some computations.
5. Aggregate the processed data from all APIs and print or display the results.

#### Tips:

* Consider using a third-party package like `github.com/go-resty/resty` for making HTTP requests.
* Use Goroutines to parallelize the API requests.
* Utilize channels to pass data between Goroutines.
* Experiment with error handling for API requests and data processing.

This project allows you to practice concurrent API requests, error handling, and data processing. Customize it based on the APIs you want to interact with and the kind of processing you'd like to perform on the fetched data.
