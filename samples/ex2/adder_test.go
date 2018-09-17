// Copyright 2018 Praveen G Shirali

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adder_test

import (
	"github.com/pshirali/testing-in-go/samples/ex2"
	"testing"
)

func TestAdderUsingTable(t *testing.T) {
	cases := []struct {
		integers []int
		expected int
	}{
		{[]int{}, 0},        // -------------------------------------------
		{[]int{0, 0, 0}, 0}, //  TABLE: { Input, Expected }
		{[]int{-1, -2}, -3}, //  One set of test params per test iteration
		{[]int{1, 2, 3}, 6}, // -------------------------------------------
	}

	for _, c := range cases {
		t.Logf("-------------------- Adding: %v", c.integers)
		actual := adder.AddInt(c.integers...)
		// START OMIT
		if actual != c.expected {
			t.Errorf("Sum of %v = %v (Actual). Expected: %v",
				c.integers, actual, c.expected)
		}
		// END OMIT
	}
}

func TestAdderUsingSubtests(t *testing.T) {
	cases := []struct {
		name     string
		integers []int
		expected int
	}{
		{"Empty", []int{}, 0},
		{"Zero", []int{0, 0, 0}, 0},
		{"Negative", []int{-1, -2}, -3},
		{"Positive", []int{1, 2, 3}, 6},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Logf("-------------------- Adding: %v", c.integers)
			actual := adder.AddInt(c.integers...)
			if actual != c.expected {
				t.Fatalf("Sum of %v = %v (Actual). Expected: %v",
					c.integers, actual, c.expected)
			}
		})
	}
}
