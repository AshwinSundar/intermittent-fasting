package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

// accepts times, days
// validates both
type Segment struct {
	date      string
	startTime string
	endTime   string
}

func (s *Segment) updateTime(t string, interval string) bool {
	if !s.validateTime(t) { // error logged in validateTime
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

func (s *Segment) updateDate(d string) bool {
	if !s.validateDate(d) {
		return false
	}

	s.date = d
	return true
}

func (s *Segment) validateDate(date string) bool {
	dateIsValid, err := regexp.MatchString("^[2][0][0-9]{2}[0-1][0-9][0-3][0-9]$", date) // YYYYMMDD

	if !dateIsValid {
		return false
	}

	if err != nil {
		log.Fatalf("validateDate failed with: %v", err)
	}

	return true
}

func (s *Segment) validateTime(time string) bool {
	timeIsValid, err := regexp.MatchString("^[0-2][0-9][0-5][0-9]$", time)

	if err != nil {
		log.Fatalf("validateTime failed with: %v", err)
	}

	return timeIsValid
}

func (s *Segment) isValid() bool {
	validDate := s.validateDate(s.date)
	validTimes := s.validateTime(s.startTime) && s.validateTime(s.endTime)

	return validDate && validTimes
}

// handles CRUD on file
// good place to use defer to close files properly
type FileInterface struct{}

func (f *FileInterface) write (seg Segment, fileName string) (bool, error) {
	if !(seg.isValid()) {
		fmt.Println("Segment is not valid: ", seg)
		return false, nil
	}	

	file, err := os.OpenFile(fileName, os.O_APPEND, 0644)

	if err != nil {
		return false, err
	}

	defer file.Close()
	strToWrite := seg.date + ", " + seg.startTime + ", " + seg.endTime
	if _, err := file.WriteString(strToWrite); err != nil {
		return false, err
	}

	return true, nil
}

// handles CLI
type CmdLineInterface struct{}

func (c *CmdLineInterface) read() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatalf("Error reading input: ", err)
		return "", err
	}

	return input, nil
}

// prints charts
type Visualizer struct{}

func main() {
	cmdLineInt := CmdLineInterface{}
	segment := Segment{}
	var err error


	fmt.Println("Enter time of first meal today (HHMM): ")
	input, err := cmdLineInt.read()
	segment.updateTime(input, "start")

	if err != nil {
		log.Fatalf("Could not write startTime, error: %v", err)
	}

	fmt.Println("Enter time of last meal today (HHMM): ")
	input, err = cmdLineInt.read()
	segment.updateTime(input, "end")

	if err != nil {
		log.Fatalf("Could not write endTime, error: %v", err)
	}
	
	segment.date = "20230805"
	fileInterface := FileInterface {}
	fileInterface.write(segment, "fakeFile") // not writing currently

	fmt.Println(segment)
}
