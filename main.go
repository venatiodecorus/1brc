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
	sum float64
	count int
	min float64
	max float64
}

type Computed struct {
	min,max,avg float64
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

	parsedData := map[string]Computed{}
	for key, value := range data {
		parsedData[key] = calcAverage(value)
	}

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

func calcAverage(vals Measurements) Computed {
	avg := vals.sum / float64(vals.count)
	avg = math.Round((avg+0.00001)*10) / 10

	return Computed{
		min: vals.min,
		max: vals.max,
		avg: avg,
	}
}

func readFile(file string) map[string]Measurements {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	results := make(map[string]Measurements)
	ch := make(chan map[string]Measurements)
	chunkSize := 50 // Number of lines to process at a time
	var lines []string
	var wg sync.WaitGroup
	// var mutex sync.Mutex

	go func() {
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
			if len(lines) == chunkSize {
				wg.Add(1)
				processChunk(lines, ch, &wg)
				lines = nil
			}
		}
		// wg.Wait()
		// close(ch)
	}()

	wg.Wait()
	close(ch)

	// go func() {
		for result := range ch {
			for city, temps := range result {
				// mutex.Lock()
				if val, ok := results[city]; ok {
					val.min = math.Min(val.min, temps.min)
					val.max = math.Max(val.max, temps.max)
					val.sum += temps.sum
					val.count += temps.count
					results[city] = val
				} else {
					results[city] = temps
				}
				// mutex.Unlock()
			}
		}
	// }()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return results
}

func processChunk(lines []string, ch chan<- map[string]Measurements, wg *sync.WaitGroup) {
	defer wg.Done()
	toSend := make(map[string]Measurements)
	
	for _, line := range lines {
		parts := strings.Split(line, ";")
		// Convert value to float
		value,err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		if val, ok := toSend[parts[0]]; ok {
			val.min = math.Min(val.min, value)
			val.max = math.Max(val.max, value)
			val.sum += value
			val.count++
			toSend[parts[0]] = val
		} else {
			toSend[parts[0]] = Measurements{
				min: value,
				max: value,
				sum: value,
				count: 1,
			}
		}

		ch <- toSend
	}
}