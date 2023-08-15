package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

// accepts times, days
// validates both
type Segment struct {
	date      string
	startTime string
	endTime   string
}

func (s *Segment) UpdateTime(t string, interval string) bool {
	if !ValidateTime(t) { // error logged in validateTime
		return false
	}

	switch interval {
	case "s", "start", "startTime":
		s.startTime = t
	case "e", "end", "endTime":
		s.endTime = t
	default:
		fmt.Println("did not update either interval")
	}

	return true
}

func (s *Segment) UpdateDate(d string) bool {
	if ValidateDate(d) {
		return false
	}

	s.date = d
	return true
}

func ValidateDate(date string) bool {
	dateIsValid, err := regexp.MatchString("^[2][0][0-9]{2}[-][0-1][0-9][-][0-3][0-9]$", date) // YYYYMMDD

	if err != nil {
		log.Fatalf("validateDate failed with: %v", err)
	}

	return dateIsValid
}

func ValidateTime(time string) bool {
	timeIsValid, err := regexp.MatchString("^[0-2][0-9][0-5][0-9]$", time)

	if err != nil {
		log.Fatalf("validateTime failed with: %v", err)
	}

	return timeIsValid
}

func (s *Segment) isValid() bool {
	validDate := ValidateDate(s.date)
	validTimes := ValidateTime(s.startTime) && ValidateTime(s.endTime)

	return validDate && validTimes
}

// FileInterface type handles writing to file. Will not handle update, delete - edit .txt directly instead.
type FileInterface struct{}

func (f *FileInterface) write (seg Segment, fileName string) bool {
	if !(seg.isValid()) {
		fmt.Println("Segment is not valid: ", seg)
		return false
	}	

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatalf("Error opening file %v - %v", fileName, err)
		return false
	}

	strToWrite := seg.date + ", " + seg.startTime + ", " + seg.endTime + "\n"
	if _, err := file.WriteString(strToWrite); err != nil {
		log.Fatalf("Error writing to %v - %v", fileName, err)
		return false
	}

	defer file.Close()
	return true
}

// CmdLineInterface handles user interface via CLI
type CmdLineInterface struct{}

func (c *CmdLineInterface) read() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatalf("Error reading input: %v", err)
		return "", err
	}

	return input, nil
}

// Visualizer handles charts
type Visualizer struct{}

func main() {
	cmdLineInt := CmdLineInterface{}
	segment := Segment{}
	var err error

	fmt.Println("Enter time of first meal today (HHMM): ")
	input, err := cmdLineInt.read()
	segment.UpdateTime(input, "start")

	if err != nil {
		log.Fatalf("Could not write startTime, error: %v", err)
	}

	fmt.Println("Enter time of last meal today (HHMM): ")
	input, err = cmdLineInt.read()
	segment.UpdateTime(input, "end")

	if err != nil {
		log.Fatalf("Could not write endTime, error: %v", err)
	}
	
	segment.UpdateDate(time.Now().Format("2006-01-02"))
	fileInterface := FileInterface {}
	res := fileInterface.write(segment, "if_log.txt")
	fmt.Printf("wrote file: %v", res)
	fmt.Println(segment)
}
