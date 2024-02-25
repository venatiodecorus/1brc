package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type measurements struct {
	sum float64
	count int
}

func main() {
	// Open file
	file := "./data/test/samples/measurements-10.txt"
	// file := "./data/measurements.txt"
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Data structure to store the sum and count for each key
	data := map[string]measurements{}

	// Read file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		// Convert value to float
		value,err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		// Increment sum and count for this key
		data[parts[0]] = measurements{
			sum: data[parts[0]].sum + value,
			count: data[parts[0]].count + 1,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	output := map[string]float64{}
	var wg sync.WaitGroup
	// Calculate average for each key
	for key, value := range data {
		wg.Add(1)
		go func(k string, v measurements) {
			defer wg.Done()
			output[k] = calcAverage(v)
		}(key, value)
	}
	wg.Wait()

	fmt.Println(output)
}

func calcAverage(data measurements) float64 {
	return data.sum / float64(data.count)
}