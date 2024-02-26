package main

import (
	"os"
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	file := "./data/test/samples/measurements-1.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func Test2(t *testing.T) {
	file := "./data/test/samples/measurements-2.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func Test3(t *testing.T) {
	file := "./data/test/samples/measurements-3.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func Test10(t *testing.T) {
	file := "./data/test/samples/measurements-10.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func Test20(t *testing.T) {
	file := "./data/test/samples/measurements-20.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func Test10000(t *testing.T) {
	file := "./data/test/samples/measurements-10000-unique-keys.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func TestBoundaries(t *testing.T) {
	file := "./data/test/samples/measurements-boundaries.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func TestComplex(t *testing.T) {
	file := "./data/test/samples/measurements-complex-utf8.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func TestDot(t *testing.T) {
	file := "./data/test/samples/measurements-dot.txt"
	dotIndex := strings.LastIndex(file, ".txt")
	resultFile := file[:dotIndex] + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func TestRounding(t *testing.T) {
	file := "./data/test/samples/measurements-rounding.txt"
	resultFile := strings.TrimRight(file, ".txt") + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func TestShort(t *testing.T) {
	file := "./data/test/samples/measurements-short.txt"
	dotIndex := strings.LastIndex(file, ".txt")
	resultFile := file[:dotIndex] + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}

func TestShortest(t *testing.T) {
	file := "./data/test/samples/measurements-shortest.txt"
	dotIndex := strings.LastIndex(file, ".txt")
	resultFile := file[:dotIndex] + ".out"
	result,err := os.ReadFile(resultFile)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	
	}
	output := ProcessData(file)
	if output != strings.TrimSpace(string(result)) {
		t.Errorf("Output is not correct")
	}
}