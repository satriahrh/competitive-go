package k_smallest_element_test

import (
	"bufio"
	"fmt"
	k_smallest_element "github.com/satriahrh/competitive-go/geeks-for-geeks/k-smallest-element"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type algorithm func(int, []int) int

type algorithmTest struct {
	name      string
	algorithm algorithm
}

func TestAlgorithm(t *testing.T) {
	files, err := ioutil.ReadDir("./test-case")
	if err != nil {
		t.Fatal(err)
	}

	for _, at := range []*algorithmTest{
		{
			name:      "CountingSort",
			algorithm: k_smallest_element.AlgorithmCountingSort,
		},
		{
			name:      "HeapSort",
			algorithm: k_smallest_element.AlgorithmHeapSort,
		},
	} {

		for _, f := range files {
			t.Run(fmt.Sprintf("%v/%v", at.name, f.Name()), at.testAlgorithmCase)
		}
	}
}

func (at *algorithmTest) testAlgorithmCase(t *testing.T) {
	testCaseName := strings.Split(t.Name(), "/")[2]
	separator := string([]byte{os.PathSeparator})
	fileInputName := fmt.Sprintf("./test-case%v%v%vinput", separator, testCaseName, separator)
	fileExpectedName := fmt.Sprintf("./test-case%v%v%voutput", separator, testCaseName, separator)

	// Read Input
	fileInput, err := os.Open(fileInputName)
	checkError(err)
	defer fileInput.Close()

	readerInput := bufio.NewReaderSize(fileInput, 1024*1024)

	kTemp := readLine(readerInput)
	k, err := strconv.Atoi(kTemp)
	checkError(err)

	arrTemp := strings.Split(readLine(readerInput), " ")

	var arr []int

	for i := 0; i < len(arrTemp); i++ {
		expenditureItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 32)
		checkError(err)
		item := expenditureItemTemp
		arr = append(arr, int(item))
	}

	// Computing
	result := at.algorithm(k, arr)

	// Write Actual Output
	actualOutput := strconv.Itoa(result)

	// Read Expected Output
	fileExpected, err := os.Open(fileExpectedName)
	checkError(err)
	defer fileExpected.Close()

	readerExpected := bufio.NewReaderSize(fileExpected, 1024*1024)
	expectedOutput := readLine(readerExpected)
	if expectedOutput != actualOutput {
		t.Fatalf("actual and expectation miss match\ngot\nexpected\n%v\n%v", actualOutput, expectedOutput)
	}
}
