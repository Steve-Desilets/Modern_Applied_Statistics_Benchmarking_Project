package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"testing"
	"time"
)

// Function to log messages to both stdout and a log file
func logMessage(message string) {
	fmt.Println(message)

	logfile, err := os.OpenFile("logGolang.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	log.Println(message)
}

// loadData loads the insurance.csv file into a 2D slice of strings
func loadData(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Skip the first row (column headers)
	records = records[1:]

	return records, nil
}

// parseFloat converts a string to a float64
func parseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// bootstrap resamples the data with replacement
func bootstrap(data []float64, nBootstraps int) []float64 {
	n := len(data)
	resamples := make([]float64, nBootstraps)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nBootstraps; i++ {
		resample := make([]float64, n)
		for j := 0; j < n; j++ {
			resample[j] = data[rand.Intn(n)]
		}
		resamples[i] = mean(resample)
	}

	return resamples
}

// mean calculates the mean of a slice of floats
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// confidenceInterval calculates the confidence interval for the given data
func confidenceInterval(data []float64, alpha float64) (lower, upper float64) {
	sort.Float64s(data)
	n := len(data)
	lowerIndex := int(float64(n) * alpha / 2)
	upperIndex := n - lowerIndex
	return data[lowerIndex], data[upperIndex]
}

// TestBootstrap tests the bootstrap function
func TestBootstrap(t *testing.T) {
	data := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	nBootstraps := 1000
	bootstrappedMeans := bootstrap(data, nBootstraps)

	if len(bootstrappedMeans) != nBootstraps {
		t.Errorf("Expected %d bootstrapped means, got %d", nBootstraps, len(bootstrappedMeans))
	} else {
		logMessage("TestBootstrap passed successfully.")
	}

}

func main() {
	// Logging start of the program
	logMessage("Starting the program...")

	// Starting software profiling
	var memStatsBefore, memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	// Load data from insurance.csv file
	records, err := loadData("insurance.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Extract charges column
	var charges []float64
	for _, record := range records {
		charge := parseFloat(record[6])
		charges = append(charges, charge)
	}

	// Execute 100 iterations of the experimental trials
	// Append the runtime for each experimental trial to a slice called runtime_slice

	var runtimeSlice []int64

	outputFile, err := os.Create("bootstrappingGolangOutput.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	loopCounter := 0
	for i := 0; i < 100; i++ {

		startTime := time.Now()

		// Perform bootstrapping
		nBootstraps := 1000
		bootstrappedMeans := bootstrap(charges, nBootstraps)

		// Calculate confidence interval
		alpha := 0.05
		lower, upper := confidenceInterval(bootstrappedMeans, alpha)
		fmt.Fprintf(outputFile, "95%% Confidence Interval for charges: [%.2f, %.2f]\n", lower, upper)

		endTime := time.Now()
		executionTime := endTime.Sub(startTime)
		runtimeSlice = append(runtimeSlice, executionTime.Microseconds())
		loopCounter += i

	}

	runtime.ReadMemStats(&memStatsAfter)
	memUsed := memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc

	// Write memory usage information to profile.txt
	profileFile, err := os.Create("profileGolang.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer profileFile.Close()

	_, err = fmt.Fprintf(profileFile, "Memory used: %d bytes\n", memUsed)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate and print the average runtime across each of the experimental trials
	var runtimeSum int64 = 0

	for i := 0; i < 100; i++ {
		runtimeSum += runtimeSlice[i]
	}
	//avgRuntime := (float64(runtimeSum)) / (float64(100))
	avgRuntime := runtimeSum / 100

	fmt.Fprintf(outputFile, "Total runtime: %d\n", runtimeSum)
	fmt.Fprintf(outputFile, "Average runtime: %d\n", avgRuntime)

	// Logging end of the program
	logMessage("Program finished successfully.")

	// Running test
	TestBootstrap(nil)

}
