package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Measurements struct {
	avg float64
	min float64
	max float64
}

func main() {
	e, err := os.Create("./profiles/" + "exec.pprof")
	if err != nil {
		log.Fatal("could not create trace execution profile: ", err)
	}
	defer e.Close()
	trace.Start(e)
	defer trace.Stop()

	f,err := os.Create("./profiles" + "/cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	m, err := os.Create("./profiles/" + "mem.pprof")
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer m.Close()
	runtime.GC()
	if err := pprof.WriteHeapProfile(m); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
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
	ch := make(chan Line)
	var wg sync.WaitGroup

	go func() {
		for scanner.Scan() {
			wg.Add(1)
			go processLine(scanner.Text(), &wg, ch)
		}
		wg.Wait()
		close(ch)
	}()

	data := map[string][]float64{}

	go func() {
		for d := range ch {
			data[d.key] = append(data[d.key], d.value)
		}
	}()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

type Line struct {
	key string
	value float64
}

func processLine(line string, wg *sync.WaitGroup, ch chan<- Line) {
	defer wg.Done()
	parts := strings.Split(line, ";")
	// Convert value to float
	value,err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	ch <- Line{
		key: parts[0],
		value: value,
	}
	// return parts[0], value

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