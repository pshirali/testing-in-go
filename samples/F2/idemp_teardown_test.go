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

package idemp_teardown_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	TestFile = "deleteme.temp"
)

var debug = flag.Bool("d", false, "Debug output: addresses & types")

func Debug(args ...interface{}) {
	if *debug {
		fmt.Println(args...)
	}
}
func HorizontalRule() {
	Debug("-------------------------------------------------------------------------------------")
}

func WriteAndVerify(t *testing.T, msg string) {
	//
	//	Write 'msg' to TestFile
	//
	bytes := []byte(msg)
	err := ioutil.WriteFile(TestFile, bytes, 0644)
	if err != nil {
		t.Fatalf("Error writing data to TestFile: Err: %v", err)
	} else {
		Debug("Wrote", len(bytes), "bytes to TestFile")
	}
	//
	//	And read it back from 'TestFile' to verify and print
	//
	readBytes, err := ioutil.ReadFile(TestFile)
	if err != nil {
		t.Fatalf("Error reading from TestFile: Err: %v", err)
	}
	if string(bytes) == string(readBytes) {
		readInQuotes := "'" + string(readBytes) + "'"
		Debug("Verified. Content is: " + readInQuotes)
	} else {
		t.Fatalf("Written & read-back data didn't match!\nWrote: %v\nRead : %v", bytes, readBytes)
	}
}

func Cleanup(t *testing.T, msg string, simulateError bool) {
	// go test switches working directory to the package running the tests
	// In this function, cleanup will attempt to delete the filename
	// set by `TestFile`, which it looks for in the present directory.

	// os.Remove already does exactly what we need reg. cleanup
	// It returns an error only if it failed to clean up when a file
	// does exist. But we're adding some extra code here to print the
	// status of the file via Debug(..)

	_, err := os.Stat(TestFile)
	if err != nil {
		if os.IsNotExist(err) {
			Debug(msg + " -- Skipping cleanup. Already clean.")
		} else {
			// unexpected error: couldn't stat file. FAIL
			t.Fatalf(msg+" -- ERROR: Can't stat TestFile: Err: %v", err)
		}
	} else {

		// simulateError will return a t.Fatalf from this point, faking what could be
		// a real error caused by os.Remove(..). Note that the TestFile must exist
		// for this feature to kick-in.
		if simulateError {
			t.Fatal(msg + "*** SIMULATING CLEANUP ERROR *** Cleanup Failed")
		}

		// attempt actual cleanup of file
		statusMsg := msg + " -- Removing TestFile: " + TestFile + " -- "
		err := os.Remove(TestFile)
		if err == nil {
			Debug(statusMsg + "Cleanup Successful.")
		} else {
			t.Fatalf(statusMsg+"Cleanup Failed: %v", err)
		}
	}
}

//
//	Tests
//
//
// Run this test with 'go test -v -d' to print debug msgs.
//
// [1] Observe that the tests start off with a clean state already.
//     That is, Cleanup(t, ">> SETUP >>"...) has nothing to do.
//     The Cleanup that's happening on TEARDOWN is cleaning up TestFile
//
// [2] In TestFirst, 'defer Cleanup(t, "<< TEARDOWN << "+t.Name(), false)'
//     set the simulateError flag from false to true.
//     - This will cause TestFirst to fail cleanup on TEARDOWN
//     - Every test attempts cleanup before its test logic `WriteAndVerify`.
//       TestSecond will perform the actual cleanup on SETUP
//
// [3] In TestSecond, 'Cleanup(t, ">> SETUP >> "+t.Name(), false)'
//     set the simulateError flag from false to true (in addition to [2])
//     - TEARDOWN from TestFirst failed to cleanup
//     - SETUP from TestSecond also failed to cleanup.
//       So, TestSecond fails, as the env from TestFirst's execution has
//       still left the env dirty.
//     - Tests will continue to pass after a successful cleanup.
//	 i. TestSecond's TEARDOWN will perform the cleanup
//       ii. TestThird passes.
//

func TestFirst(t *testing.T) {
	HorizontalRule()
	defer HorizontalRule()
	Cleanup(t, ">> SETUP >> "+t.Name(), false)
	defer Cleanup(t, "<< TEARDOWN << "+t.Name(), false)
	WriteAndVerify(t, "Hello from "+t.Name())
}

func TestSecond(t *testing.T) {
	HorizontalRule()
	defer HorizontalRule()
	Cleanup(t, ">> SETUP >> "+t.Name(), false)
	defer Cleanup(t, "<< TEARDOWN << "+t.Name(), false)
	WriteAndVerify(t, "Hello from "+t.Name())
}

func TestThird(t *testing.T) {
	HorizontalRule()
	defer HorizontalRule()
	Cleanup(t, ">> SETUP >> "+t.Name(), false)
	defer Cleanup(t, "<< TEARDOWN << "+t.Name(), false)
	WriteAndVerify(t, "Hello from "+t.Name())
}
