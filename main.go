package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Measurements struct {
	sum float64
	count int
	min float64
	max float64
}

func main() {
	// Open file
	// file := "./data/test/samples/measurements-10.txt"
	// file := "./data/test/samples/measurements-rounding.txt"
	file := "./data/measurements.txt"
	output := ProcessData(file)

	fmt.Println(output)
}

func ProcessData(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Data structure to store the sum and count for each key
	data := map[string]Measurements{}

	// Read file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		// Convert value to float
		value,err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		min := data[parts[0]].min
		if value <= min || data[parts[0]].count == 0{
			min = value
		}

		max := data[parts[0]].max
		if value > max || data[parts[0]].count == 0{
			max = value
		}

		// Increment sum and count for this key
		data[parts[0]] = Measurements{
			sum: data[parts[0]].sum + value,
			count: data[parts[0]].count + 1,
			min: min,
			max: max,
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	averages := map[string]float64{}
	var wg sync.WaitGroup
	var mu sync.Mutex
	// Calculate average for each key
	for key, value := range data {
		wg.Add(1)
		go func(k string, v Measurements) {
			defer wg.Done()
			avg := calcAverage(v)
			mu.Lock()
			averages[k] = avg
			mu.Unlock()
		}(key, value)
	}
	wg.Wait()

	// Extract and sort the keys
	keys := make([]string, 0, len(averages))
	for k := range averages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	output := ""
	for _,k := range keys {
		output += fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", k, data[k].min, averages[k], data[k].max)
	}
	output = strings.TrimRight(output, ", ")

	return "{" + output + "}"
}


func calcAverage(data Measurements) float64 {
	return data.sum / float64(data.count)
}