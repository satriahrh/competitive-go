package sorting_test

import (
	"bufio"
	"fmt"
	merge_sort "github.com/satriahrh/competitive-go/miscellaneous/sorting/merge-sort"
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

type sortingAlgorithm func([]int32) []int32

type SortingTest struct {
	algorithm sortingAlgorithm
}

func TestApplication(t *testing.T) {
	sortingTest := &SortingTest{
		algorithm: merge_sort.SortingAlgorithm,
	}

	files, err := ioutil.ReadDir("./test-case")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		t.Run(f.Name(), sortingTest.testApplicationCase)
	}

}

func (st *SortingTest) testApplicationCase(t *testing.T) {
	testCaseName := t.Name()[16:]
	separator := string([]byte{os.PathSeparator})
	fileInputName := fmt.Sprintf("./test-case%v%v%vinput", separator, testCaseName, separator)
	fileExpectedName := fmt.Sprintf("./test-case%v%v%voutput", separator, testCaseName, separator)

	// Read Input
	fileInput, err := os.Open(fileInputName)
	checkError(err)
	defer fileInput.Close()

	readerInput := bufio.NewReaderSize(fileInput, 1024*1024)

	arrTemp := strings.Split(readLine(readerInput), " ")

	var arr []int32

	for i := 0; i < len(arrTemp); i++ {
		expenditureItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		arr = append(arr, expenditureItem)
	}

	// Computing
	result := st.algorithm(arr)
	if len(arrTemp) != len(result) {
		t.Fatalf("len actual and expectation miss match\ngot\nexpected\n%v\n%v", len(result), len(arrTemp))
	}

	// Write Actual Output
	actualOutput := strings.Trim(fmt.Sprint(result), "[]")

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
