package fraudulent_activity_notifications_test

import (
	"bufio"
	"fmt"
	"github.com/satriahrh/competitive-go/hackerrank/fraudulent-activity-notifications"
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

func TestApplication(t *testing.T) {
	files, err := ioutil.ReadDir("./test_case")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		t.Run(f.Name(), testApplicationCase)
	}

}

func testApplicationCase(t *testing.T) {
	testCaseName := t.Name()[16:]
	separator := string([]byte{os.PathSeparator})
	fileInputName := fmt.Sprintf("./test_case%v%v%vinput", separator, testCaseName, separator)
	fileExpectedName := fmt.Sprintf("./test_case%v%v%voutput", separator, testCaseName, separator)

	// Read Input
	fileInput, err := os.Open(fileInputName)
	checkError(err)
	defer fileInput.Close()

	readerInput := bufio.NewReaderSize(fileInput, 1024*1024)

	nd := strings.Split(readLine(readerInput), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(readerInput), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	// Computing
	result := fraudulent_activity_notifications.ActivityNotifications(expenditure, d)

	// Write Actual Output
	actualOutput := fmt.Sprintf("%d", result)

	// Read Expected Output
	fileExpected, err := os.Open(fileExpectedName)
	checkError(err)
	defer fileExpected.Close()

	readerExpected := bufio.NewReaderSize(fileExpected, 1024*1024)
	expectedOutput := readLine(readerExpected)
	if expectedOutput != actualOutput {
		t.Fatalf("got %v expected %v", actualOutput, expectedOutput)
	}
}
