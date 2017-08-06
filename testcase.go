// Package testcase provides availability to run few tests cases inside single Golang test method
// Just prepare slice of your cases, provide arguments for before/after/passed functions & start Run()
package testcase

import "testing"

// Params contains arguments for testcase closures: before/after/passed
// Map key - parameter name, value - should be casted from interface{} to required value
type Params struct {
	BeforeArgs map[string]interface{}
	AfterArgs  map[string]interface{}
	PassedArgs map[string]interface{}
}

// Cases - slice of tests cases that you want to run
type Cases []tcase

// tcase describes test case to run.
// Name - test case name, params - arguments for test helper closures.
// Before - callback to execute before test case start (like create file / cleanup dir / etc)
// After - callback to execute after test (cleanup smth / etc)
// Passed - test body here. Returns bool value (passed / failed test)
type tcase struct {
	Name   string
	Params Params
	Before func(args map[string]interface{})
	After  func(args map[string]interface{})
	Passed func(args map[string]interface{}) bool
}

// Run starts provided tests cases
func Run(t *testing.T, cs Cases) {
	for _, c := range cs {
		t.Run(c.Name, func(t *testing.T) {
			c.Before(c.Params.BeforeArgs)
			defer c.After(c.Params.AfterArgs)

			if !c.Passed(c.Params.PassedArgs) {
				t.Errorf("Test case `%s` failed", c.Name)
			}
		})
	}
}
