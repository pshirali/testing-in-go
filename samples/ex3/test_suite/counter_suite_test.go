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

package test_suite

// NOTE: The dot import has been used here to condense the data on the slides.
//       AVOID: '. <import-path' in your code.
import (
	"flag"
	"fmt"
	. "github.com/pshirali/testing-in-go/samples/ex3/counter"
	. "github.com/pshirali/testing-in-go/samples/ex3/safe_counter"
	. "github.com/pshirali/testing-in-go/samples/ex3/unsafe_counter"
	"reflect"
	"strings"
	"testing"
)

func Assert(t *testing.T, condition bool, args ...interface{}) {
	if !condition {
		t.Fatal(args...)
	}
}

var debug = flag.Bool("d", false, "Debug output: addresses & types")
var wrap = flag.Bool("w", false, "Debug output: before & after wrapper calls for (a) suite & (b) per-test")

func Debug(fmtStr string, args ...interface{}) {
	if *debug {
		fmt.Printf(fmtStr, args...)
	}
}
func DebugFixture(msg string) {
	if *wrap {
		fmt.Println(msg)
	}
}

type Counter interface {
	Resetter
	Incrementer
	IntValuer
}

// BUILDER START OMIT
func UnsafeCounterBuilder() Counter {
	return NewUnsafeCounter() // returns *unsafeCounter
}

func SafeCounterBuilder() Counter {
	return NewSafeCounter() // returns *safeCounter
}

// BUILDER END OMIT

// SUITE START OMIT
type suite struct {
	builder func() Counter
}

func Suite(builder func() Counter) *suite {
	return &suite{builder: builder}
}

// SUITE END OMIT

/* ----------------------------------------------------------------------------

		This commented section is used in the slides

// TESTS START OMIT
func (s *suite) TestCounterIncrementIncreasesValue(t *testing.T) {     // Added (s *suite)
	c := s.builder()              // <---------------------- new instance built by builder
	for i := 1; i < 3; i++ {
		c.Increment()
		Assert(t, c.Value() == i, "At Step:", i, "!=", c.Value())
	}
}

func (s *suite) TestCounterIncrementReset(t *testing.T) {              // Added (s *suite)
	c := s.builder()              // <---------------------- new instance built by builder
	for i := 0; i < 2; i++ {
		c.Increment()
	}
	c.Reset()
	Assert(t, c.Value() == 0, "Expected 0 after Reset. Got:", c.Value())
}
// TESTS END OMIT

// RUNALL START OMIT
func (s *suite) RunAllTests(t *testing.T) {
	v := reflect.ValueOf(s)                                 //  1. Reflect on the suite
	for n := 0; n < v.NumMethod(); n++ {                    //  2. Iterate through method numbers
		i := v.Method(n).Interface()                        //  3. Get the method as interface{}
		if method, ok := i.(func(*testing.T)); ok {         //  4. If it matches test signature
			methodName := reflect.TypeOf(s).Method(n).Name  //  5. Get the method's name
			if strings.HasPrefix(methodName, "Test") {      //  6. If it begins with 'Test'
				t.Run(methodName, method)                   //  7. Run that method as a test
			}
		}
	}
}
// RUNALL END OMIT

// RUNNER_FIXTURE START OMIT
func (s *suite) RunAllTests(t *testing.T) {
	//
	// Before and After Suite (defer <After>())
	//
	v := reflect.ValueOf(s)
	for n := 0; n < v.NumMethod(); n++ {
		i := v.Method(n).Interface()
		if method, ok := i.(func(*testing.T)); ok {
			methodName := reflect.TypeOf(s).Method(n).Name
			if strings.HasPrefix(methodName, "Test") {
				// Before Test
				// (don't defer <AfterTest>() here, inside a loop)
				t.Run(methodName, method)
				// After Test
			}
		}
	}
}
// RUNNER_FIXTURE END OMIT

// RUN_TEST_SUITE START OMIT
func TestCounterSuite(t *testing.T) {	// Table + Subtest driven tests against the Suite
	cases := []struct {
		name    string
		builder func() Counter
	}{
		{"SafeCounter", SafeCounterBuilder},
		{"UnsafeCounter", UnsafeCounterBuilder},
	}

	for _, c := range cases {
		t.Run(c.name, Suite(c.builder).RunAllTests)
	}
}
// RUN_TEST_SUITE END OMIT
*/

// -------
// The code below is a replica. It contains additional debug output
// When you run 'go test -v', you might be interested in the code below
//

func (s *suite) TestCounterIncrementIncreasesValue(t *testing.T) {
	c := s.builder()
	Debug("        ===== Test1 : %p %T\n", t, t)
	Debug("         \\_ Counter : %p %T\n", c, c)
	for i := 1; i < 3; i++ {
		c.Increment()
		Assert(t, c.Value() == i, "At Step:", i, "!=", c.Value())
	}
}

func (s *suite) TestCounterIncrementReset(t *testing.T) {
	c := s.builder()
	Debug("        ===== Test2 : %p %T\n", t, t)
	Debug("         \\_ Counter : %p %T\n", c, c)
	for i := 0; i < 2; i++ {
		c.Increment()
	}
	c.Reset()
	Assert(t, c.Value() == 0, "Expected 0 after Reset. Got:", c.Value())
}

func (s *suite) RunAllTests(t *testing.T) {
	Debug("   ==== RunAllTests : %p %T\n", t, t)
	DebugFixture("\n    >>> [ RunAllTests -Before- ] >>>\n")
	defer DebugFixture("\n    <<< [ RunAllTests -After- ] <<<\n")
	v := reflect.ValueOf(s)
	for n := 0; n < v.NumMethod(); n++ {
		i := v.Method(n).Interface()
		if method, ok := i.(func(*testing.T)); ok {
			methodName := reflect.TypeOf(s).Method(n).Name
			if strings.HasPrefix(methodName, "Test") {
				DebugFixture("        --- [ BeforeTest: " + methodName + " ] ---")
				t.Run(methodName, method)
				DebugFixture("        --- [ AfterTest: " + methodName + " ] ---")
			}
		}
	}
}

func TestCounterSuite(t *testing.T) {
	Debug("@@ TestCounterSuite : %p %T\n", t, t)
	cases := []struct {
		name    string
		builder func() Counter
	}{
		{"SafeCounter", SafeCounterBuilder},
		{"UnsafeCounter", UnsafeCounterBuilder},
	}

	for _, c := range cases {
		t.Run(c.name, Suite(c.builder).RunAllTests)
	}
}
