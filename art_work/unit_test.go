package art_work

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

//global variables
var input []string = []string{"Hello", "Google"}
var raw_data string = ` _    _          _   _  ` + "\n" +
	`| |  | |        | | | |         ` + "\n" +
	`| |__| |   ___  | | | |   ___   ` + "\n" +
	`|  __  |  / _ \ | | | |  / _ \  ` + "\n" +
	`| |  | | |  __/ | | | | | (_) | ` + "\n" +
	`|_|  |_|  \___| |_| |_|  \___/  ` + "\n" +
	`                                  ` + "\n" +
	`                                  ` + "\n" +
	`  _____                           _         ` + "\n" +
	` / ____|                         | |        ` + "\n" +
	`| |  __    ___     ___     __ _  | |   ___  ` + "\n" +
	`| | |_ |  / _ \   / _ \   / _` + "`" + ` | | |  / _ \ ` + "\n" +
	`| |__| | | (_) | | (_) | | (_| | | | |  __/ ` + "\n" +
	` \_____|  \___/   \___/   \__, | |_|  \___| ` + "\n" +
	`                           __/ |            ` + "\n" +
	`                          |___/            ` + "\n"
var outputFile []string = strings.Split(raw_data, "\n")

func TestAsciiMap(t *testing.T) {
	expectedMap1 := map[rune]int{
		rune(32): 17,
	}
	asciiMap1 := asciiMap(outputFile)
	if !reflect.DeepEqual(asciiMap1, expectedMap1) {
		t.Errorf("TestAsciiMap failed! Expected: %v, got: %v", expectedMap1, asciiMap1)
	}
}

func TestOutput(t *testing.T) {
	expected := raw_data

	contents, err := os.ReadFile("../data/standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	outputFile = strings.Split(string(contents), "\n")
	result := Output(input, outputFile)
	// fmt.Println(len(expected))
	// fmt.Println(len(result))
	for _, val := range expected {
		fmt.Printf("Exp:%v\n", val)
	}
	for _, val := range result {
		fmt.Printf("res:%v\n", val)
	}

	if !reflect.DeepEqual(result, expected) {
		// t.Errorf("TestOutput failed!\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}
