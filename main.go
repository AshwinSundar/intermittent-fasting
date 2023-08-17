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
	if !ValidateDate(d) {
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

func FileWrite (seg Segment, fileName string) bool {
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

// Checks for duplicates, makes sure each line is valid
// func FileValidate(fileName string) bool {}

func CliRead() (string, error) {
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
	segment := Segment{}
	var err error

	fmt.Println("Enter date (YYYY-MM-DD)")
	input, err := CliRead()
	if err != nil {
		log.Fatalf("Could not write startTime, error: %v", err)
	}
	res := segment.UpdateDate(input)
	if !res {
		log.Fatalf("Could not update date.")
	}


	fmt.Println("Enter time of first meal (HHMM): ")
	input, err = CliRead()
	if err != nil {
		log.Fatalf("Could not write startTime, error: %v", err)
	}
	res = segment.UpdateTime(input, "start")
	if !res {
		log.Fatalf("Could not update start time.")
	}


	fmt.Println("Enter time of last meal today (HHMM): ")
	input, err = CliRead()
	if err != nil {
		log.Fatalf("Could not write endTime, error: %v", err)
	}
	res = segment.UpdateTime(input, "end")
	if !res {
		log.Fatalf("Could not update end time.")
	}
	
	res = FileWrite(segment, "if_log.txt")

	if res {
		fmt.Println("Write successful")
	} else {
		fmt.Println("Write failed")
	}
}
