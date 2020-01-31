package main

import "testing"

func TestIndex(t *testing.T) {
	tests := []struct {
		name      string
		items     [] int
		predicate func(int) bool
		want      int
	}{
		{"Index exist", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 9 {
				return true
			}
			return false
		}, 8},
		{"Index does not exist", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 10 {
				return true
			}
			return false
		}, -1},
	}
	for _, tt := range tests {
		if got := Index(tt.items, tt.predicate); got != tt.want {
			t.Errorf("Test for %v Index() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		name      string
		items     []int
		predicate func(int) bool
		want      bool
	}{
		{"All true", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 9 {
				return true
			}
			return false
		}, false},
		{"All false", []int{2, 2, 2, 2, 2, 2, 2, 2, 2}, func(i int) bool {
			if i == 2 {
				return true
			}
			return false
		}, true},
	}
	for _, tt := range tests {
		if got := All(tt.items, tt.predicate); got != tt.want {
			t.Errorf("Test for %v All() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		name      string
		items     [] int
		predicate func(int) bool
		want      bool
	}{
		{"Any true", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 9 {
				return true
			}
			return false
		}, true},
		{"Any false", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 10 {
				return true
			}
			return false
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.items, tt.predicate); got != tt.want {
				t.Errorf("Test for %v Any() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestNone(t *testing.T) {
	tests := []struct {
		name string
		items     []int
		predicate func(int) bool
		want bool
	}{
		{"None false", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 9 {
				return true
			}
			return false
		}, false},
		{"None true", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 10 {
				return true
			}
			return false
		}, true},
	}
	for _, tt := range tests {
		if got := None(tt.items, tt.predicate); got != tt.want {
			t.Errorf("Test for %v None() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name string
		items     []int
		predicate func(int) bool
		want int
	}{
		{"Find result", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(i int) bool {
			if i == 9 {
				return true
			}
			return false
		}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.items, tt.predicate); got != tt.want {
				t.Errorf("Test for %v Find() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}