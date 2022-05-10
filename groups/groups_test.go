package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetGroupById(t *testing.T) {
	given := "1" // Given Group Id 1
	want := &Group{"1", "Marvelous Group", 1, []string{"Liam"}, "Mathew"}
	got, err := queryGroupById(given)
	condition := want.Name == got.Name &&
		want.Id == got.Id &&
		reflect.DeepEqual(want.Participants, got.Participants) &&
		err == nil
	if !condition {
		t.Errorf("Want: %v  but Got: %v", want, got)
	}
}

func TestGetGroups(t *testing.T) {
	tests := []bool{true, false}
	expectations := []int{1, 2}
	for i := 0; i < 2; i++ {
		fmt.Printf("Starting Test %v", i)
		want := expectations[i]
		got, err := queryGroups(tests[i])
		if len(*got) != want || err != nil {
			t.Errorf("Test %v - Want: %v  but Got: %v", i, want, got)
		}
	}
}
