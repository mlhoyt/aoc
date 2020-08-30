package main

import (
	"bytes"
	"fmt"
	"github.com/mlhoyt/adventofcode.com-2019/day7/pkg/intcode"
	"os"
	"strconv"
	"strings"
)

var codeData = []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 46, 67, 88, 101, 126, 207, 288, 369, 450, 99999, 3, 9, 1001, 9, 5, 9, 1002, 9, 5, 9, 1001, 9, 5, 9, 102, 3, 9, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 101, 5, 9, 9, 102, 5, 9, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 102, 2, 9, 9, 1001, 9, 5, 9, 102, 4, 9, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 1001, 9, 4, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 1001, 9, 3, 9, 1002, 9, 2, 9, 101, 4, 9, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 99, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99}

func main() {
	maxOutputSignal := 0

	for _, ampSetting := range generateAmpSettings() {
		ampCode := make([]int, len(codeData))
		copy(ampCode, codeData)

		outputSignal, err := simulateAmpChain(ampCode, ampSetting)
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("[DEBUG] ampSettings=%v output=%d\n", ampSetting, outputSignal)
		if outputSignal > maxOutputSignal {
			maxOutputSignal = outputSignal
		}
	}

	fmt.Printf("[DEBUG] maxOutputSignal=%d\n", maxOutputSignal)
}

func simulateAmpChain(ampCode []int, ampSetting []int) (int, error) {
	fmt.Printf("[DEBUG] ampSetting=%v\n", ampSetting)

	inputSignal := 0
	for _, inputSetting := range ampSetting {
		outputSignal, err := simulateAmp(ampCode, inputSetting, inputSignal)
		if err != nil {
			return -1, err
		}

		inputSignal = outputSignal
	}

	return inputSignal, nil
}

func simulateAmp(code []int, setting int, signal int) (int, error) {
	var input bytes.Buffer
	_, err := input.WriteString(fmt.Sprintf("%d\n%d\n", setting, signal))
	if err != nil {
		return -1, fmt.Errorf("[ERROR] simulateAmp: failed to initialize input stream: %v", err)
	}

	var output bytes.Buffer

	intcode := intcode.NewIntCode(code, &input, &output)
	err = intcode.Run()
	if err != nil {
		return -1, fmt.Errorf("[ERROR] simulateAmp: failed running intcode program: setting=%d signal=%d error=%v", setting, signal, err)
	}

	outputStr := output.String()
	outputSignal, err := strconv.Atoi(strings.TrimSpace(outputStr))
	if err != nil {
		return -1, fmt.Errorf("[ERROR] simulateAmp: failed to convert output string to int: output=%s %v", outputStr, err)
	}

	fmt.Printf("[DEBUG] simulateAmp: setting=%d signal=%d output=%d\n", setting, signal, outputSignal)

	return outputSignal, nil
}

const nMax = 5

func generateAmpSettings() [][]int {
	// length = 1
	settings := [][]int{[]int{0}, []int{1}, []int{2}, []int{3}, []int{4}}
	newSettings := [][]int{}
	for _, setting := range settings {
		newSettings = append(newSettings, buildAmpSettings(setting)...)
	}

	// length = 2
	settings = newSettings
	newSettings = [][]int{}
	for _, setting := range settings {
		newSettings = append(newSettings, buildAmpSettings(setting)...)
	}

	// length = 3
	settings = newSettings
	newSettings = [][]int{}
	for _, setting := range settings {
		newSettings = append(newSettings, buildAmpSettings(setting)...)
	}

	// length = 4
	settings = newSettings
	newSettings = [][]int{}
	for _, setting := range settings {
		newSettings = append(newSettings, buildAmpSettings(setting)...)
	}

	return newSettings
}

func buildAmpSettings(suffix []int) [][]int {
	settings := [][]int{}

	for i := 0; i < nMax; i++ {
		if !intSliceContains(suffix, i) {
			settings = append(settings, append([]int{i}, suffix...))
		}
	}

	return settings
}

func intSliceContains(s []int, v int) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}

	return false
}
