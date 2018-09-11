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

package fixture_test

import (
	"fmt"
	"testing"
)

//
//     func FixtureX(t *testing.T) func () {
//             // SETUP GOES HERE
//             return func() {
//                     // TEARDOWN GOES HERE
//             }
//     }
//
// There are three fixtures below. Each fixture also has a t.Fatal
// in both SETUP and TEARDOWN areas. Uncomment the t.Fatal lines
// in specific fixtures to observe behavior.
//
// Behavior:
// [1] t.Fatal during SETUP in any fixture fails the test at that point.
//     Deferred TEARDOWNs till that point will get called.
// [2] t.Fatal in TEARDOWN will stop the execution of only that teardown
//     function. TEARDOWNs from outer fixtures WILL get called.
//     Example: Uncomment t.Fatal() in Fixture2's TEARDOWN
//              Fixture1's TEARDOWN does get called
//

func Fixture1(t *testing.T) func() {
	// t.Fatalf("SetupError: "+t.Name())
	fmt.Println(">>>>>> SETUP (1) : " + t.Name())
	return func() {
		// t.Fatalf("TeardownError: "+t.Name())
		fmt.Println("<<< TEARDOWN (1) : " + t.Name())
	}
}

func Fixture2(t *testing.T) func() {
	// t.Fatalf("SetupError: "+t.Name())
	fmt.Println(">>>>>> SETUP (2) : " + t.Name())
	return func() {
		// t.Fatalf("TeardownError: "+t.Name())
		fmt.Println("<<< TEARDOWN (2) : " + t.Name())
	}
}

func Fixture3(t *testing.T) func() {
	// t.Fatalf("SetupError: "+t.Name())
	fmt.Println(">>>>>> SETUP (3) : " + t.Name())
	return func() {
		// t.Fatalf("TeardownError: "+t.Name())
		fmt.Println("<<< TEARDOWN (3) : " + t.Name())
	}
}

//
//	Test
//
func TestFixtures(t *testing.T) {
	defer Fixture1(t)()
	defer Fixture2(t)()
	defer Fixture3(t)()
	fmt.Println("<<< --- TEST CODE --- >>>")
}
