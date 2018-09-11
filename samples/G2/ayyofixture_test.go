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

package ayyofixture_test

import (
	"fmt"
	"testing"
)

func AyyoFixture() func() {
	fmt.Println("\n**** This is the SETUP ****")
	fmt.Println("")
	return func() {
		fmt.Println("\n**** This is the TEARDOWN ****")
	}
}

//
//	Test
//
func TestAyyoFixture(t *testing.T) {
	defer AyyoFixture() // Missed the parantheses that calls the returned teardown function
	fmt.Println("\n  --- TestCode Start ---")
	fmt.Println("  Observe:")
	fmt.Println("  [1] SETUP executes after the test code.")
	fmt.Println("  [2] TEARDOWN doesn't execute.")
	fmt.Println("  --- TestCode End ---")
}
