package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurements struct {
	avg float64
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
	data := readFile(file)

	parsedData := map[string]Measurements{}
	for key, value := range data {
		parsedData[key] = calcValues(value)
	}

	// f, err := os.Open(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// // Data structure to store the sum and count for each key
	// data := map[string]Measurements{}

	// // Read file line by line
	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	parts := strings.Split(scanner.Text(), ";")
	// 	// Convert value to float
	// 	value,err := strconv.ParseFloat(parts[1], 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	min := data[parts[0]].min
	// 	if value <= min || data[parts[0]].count == 0{
	// 		min = value
	// 	}

	// 	max := data[parts[0]].max
	// 	if value > max || data[parts[0]].count == 0{
	// 		max = value
	// 	}

	// 	// Increment sum and count for this key
	// 	data[parts[0]] = Measurements{
	// 		sum: data[parts[0]].sum + value,
	// 		count: data[parts[0]].count + 1,
	// 		min: min,
	// 		max: max,
	// 	}
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// averages := map[string]float64{}
	// var wg sync.WaitGroup
	// var mu sync.Mutex
	// // Calculate average for each key
	// for key, value := range data {
	// 	wg.Add(1)
	// 	go func(k string, v Measurements) {
	// 		defer wg.Done()
	// 		avg := calcAverage(v)
	// 		mu.Lock()
	// 		averages[k] = avg
	// 		mu.Unlock()
	// 	}(key, value)
	// }
	// wg.Wait()

	// Extract and sort the keys
	keys := make([]string, 0, len(parsedData))
	for k := range parsedData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	output := ""
	for _,k := range keys {
		output += fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", k, parsedData[k].min, parsedData[k].avg, parsedData[k].max)
	}
	output = strings.TrimRight(output, ", ")

	return "{" + output + "}"
}


// func calcAverage(data Measurements) float64 {
// 	avg := data.sum / float64(data.count)
// 	// Rounding was still off on a value of 25.45, so I added a small number to the average
// 	return math.Round((avg+0.00001)*10) / 10
// }

func calcValues(city []float64) Measurements {
	min := city[0]
	max := city[0]
	sum := 0.0
	count := 0
	for _, value := range city {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		sum += value
		count++
	}
	avg := sum / float64(count)
	avg = math.Round((avg+0.00001)*10) / 10

	return Measurements{
		min: min,
		max: max,
		avg: avg,
	}
}

func readFile(file string) map[string][]float64 {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ch := make(chan string)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	data := map[string][]float64{}
	for text := range ch {
		key, value := processLine(text)
		data[key] = append(data[key], value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func processLine(line string) (string, float64) {
	parts := strings.Split(line, ";")
	// Convert value to float
	value,err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	return parts[0], value

	// min := data[parts[0]].min
	// if value <= min || data[parts[0]].count == 0{
	// 	min = value
	// }

	// max := data[parts[0]].max
	// if value > max || data[parts[0]].count == 0{
	// 	max = value
	// }

	// // Increment sum and count for this key
	// data[parts[0]] = Measurements{
	// 	sum: data[parts[0]].sum + value,
	// 	count: data[parts[0]].count + 1,
	// 	min: min,
	// 	max: max,
	// }

	// return parts[0], Measurements{
	// 	sum: value,
	// 	count: 1,
	// 	min: value,
	// 	max: value,
	// }
}