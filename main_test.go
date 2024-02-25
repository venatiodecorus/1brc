package main

import (
	"os"
	"strings"
	"testing"
)

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