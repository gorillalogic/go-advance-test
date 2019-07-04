# Go Advance-Level Developer Coding Test

This is the coding test for advance-level [Go](https://golang.org/) developers.  Please fork this repository into your own account or download the code and push it into your own public repository.


# Fibonacci

Write the code for **Fibonacci** function in *fibonacci.go* so the provided unit test works as expected.

## Unit test

Your code should pass the provided unit test successfully.

```go
package fib

import "testing"

func TestFibonacci(t *testing.T) {
	data := []struct {
		n    uint
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296}, {90, 2880067194370816120},
	}

	for _, d := range data {
		if got := Fibonacci(d.n); got != d.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}
```
**Plus**
> Create a documentation for the fibonacci.go using GoDoc tool.  Should contain at least two call examples on the output documentation.
> Why the function can't be used to calculate Fibonacci(1200)?
> Create a command line program that calls the Fibonacci function outputs the first 20 fibonacci numbers and calculate the time it took to complete the task.

# REST API 

Design a REST API for the **Person** data model.  The API should support the following methods:

 - POST:  To create a new Person and store it.
 - PUT: To update a previously created Person.
 - GET: Retrieve the data associated to a given Person.
 - DELETE: To remove an existing Person.

The API should also support listing all the Person entities on the system.  All the entities can be stored on memory.  Should response with a **JSON** representation for all the responses.

Create a new public repository on any cloud git provider (Github, Gitlab or BitBucket) and upload your project there.

There is no restriction on the use of any web framework you decide but we prefer [gin](https://github.com/gin-gonic/gin).

In case your API needs something special for us to test it you should provide us with specific instructions.

## Person type

Consider the following type for your API:

```go
// Person represents any given person for your system
type Person struct {
	FirstName      string
	LastName       string
	Identification string
	Gender         string
	Email          string
	Phone          string
	DateOfBirth    string
	ProfilePic     string
	CreateDate     time.Time
}
```

**Plus**
> Host your implementation on any cloud provider service so we can test your API live.
> Create unit tests for your API.
> Use a [Redis](https://redis.io/) server to support cache to in your API.
> Store the **Person** entities in a NoSQL database (e.g. MongoDB).
> Support responding in XML format in addition to JSON.
> Create a "living HTML documentation" for your API. Consider using [https://swagger.io/](https://swagger.io/)

