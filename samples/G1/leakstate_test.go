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

package leakstate_test

import (
	"fmt"
	"testing"
	"errors"
)

func stateChangingHelper() error {
	fmt.Println("Step1: State change complete. No errors.")
	return nil
}

func helperThatFails() error {
	return errors.New("Step2: Failed with error. Observe: [1] Step1 changed state [2] No teardown")
}


func LeakingFixture(t *testing.T) func() {
	var err error
	err = stateChangingHelper()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	err = helperThatFails()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	return func() {
		fmt.Println("Teardown: Executed.")
	}

}

//
//	Test
//
func TestLeakingFixture(t *testing.T) {
	defer LeakingFixture(t)()
	fmt.Println("<<< --- TEST CODE --- >>>")
}
