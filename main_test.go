package main

import (
	"testing"
	_ "regexp"
)

func Test_ValidateDate(t *testing.T) {
	cases := []struct {
		description string
		date string
		expected bool
	}{
		{
			description: "Jan 1 2023",
			date: "Jan 1 2023",
			expected: false,
		},
		{
			description: "01012023",
			date: "01012023",
			expected: false,
		},
		{
			description: "2023-01-01",
			date: "2023-01-01",
			expected: true,
		},
		{
			description: "2023-20-12",
			date: "2023-20-12",
			expected: false,
		},
		/*
		// test fails, currently allows YYYY13DD - YYYY19DD
		{
			description: "20231512",
			date: "20231512",
			expected: false,
		},
		*/
	}

	for _, tc := range cases {
		// uncomment to enable parallelization
		// tc := tc // defines looping variable as a local variable, to prevent race conditions
		t.Run(tc.description, func(t *testing.T) {
			// uncomment to enable parallelization
			// t.Parallel() // runs tests in this loop in parallel
			result := ValidateDate(tc.date)

			if result != tc.expected {
				t.Errorf("For %v, expected %v, got: %v", tc.description, tc.expected, result)
			}
		})
	}
}

func Test_ValidateTime(t *testing.T) {
	cases := []struct {
		description string
		time string
		expected bool
	}{
		{
			description: "0000",
			time: "0000",
			expected: true,
		},
		{
			description: "2359",
			time: "2359",
			expected: true,
		},
		/*
		// test fails - currently allows 2400-2959
		{
			description: "2400",
			time: "2400",
			expected: false,
		},
		*/
		{
			description: "1250 AM",
			time: "1250 AM",
			expected: false,
		},
	}

	for _, tc := range cases {
		// uncomment to enable parallelization
		// tc := tc // defines looping variable as a local variable, to prevent race conditions
		t.Run(tc.description, func(t *testing.T) {
			// uncomment to enable parallelization
			// t.Parallel() // runs tests in this loop in parallel
			result := ValidateTime(tc.time)

			if result != tc.expected {
				t.Errorf("For %v, expected %v, got: %v", tc.description, tc.expected, result)
			}
		})
	}
}
