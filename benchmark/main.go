package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gobook/ch1/echo1"
	"gobook/ch1/echo3"
)

// Run a function with args and output how long it takes
func MeasureTime(label string, testFn func([]string), args []string) {
	devNull, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
	defer devNull.Close()

	// Redirect output to /dev/null
	origStdout := os.Stdout
	os.Stdout = devNull

	// Measure execution time
	start := time.Now()
	testFn(args)
	elapsed := time.Since(start)

	// Output results
	os.Stdout = origStdout
	fmt.Printf("%s took %s\n", label, elapsed)
}

func main() {
	args := []string{"echo", "this", "is", "a", "benchmark", "test"}

	// Run tests
	MeasureTime("Echo1 (small input)", echo1.Echo1, args) // string concatenation
	MeasureTime("Echo3 (small input)", echo3.Echo3, args) // strings.Join()

	// Make larger input
	largeArgs := []string{"echo"}
	for i := 0; i < 10000; i++ {
		largeArgs = append(largeArgs, strings.Repeat("word", i%10))
	}

	// Run again
	MeasureTime("Echo1 (large input)", echo1.Echo1, largeArgs)
	MeasureTime("Echo3 (large input)", echo3.Echo3, largeArgs)
}
