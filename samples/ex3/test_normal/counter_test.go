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

package test_normal

// NOTE: The dot import has been used here to condense the data on the slides.
//       AVOID: '. <import-path' in your code.
import (
	. "gitlab.com/pshirali/testing-in-go/samples/ex3/unsafe_counter"
	"testing"
)

func Assert(tb testing.TB, condition bool, args ...interface{}) {
	if !condition {
		tb.Fatal(args...)
	}
}

// START OMIT
func TestCounterIncrementIncreasesValue(t *testing.T) {
	c := NewUnsafeCounter()
	for i := 1; i < 3; i++ {
		c.Increment()
		Assert(t, c.Value() == i, "At Step:", i, "!=", c.Value())
	}
}

func TestCounterIncrementReset(t *testing.T) {
	c := NewUnsafeCounter()
	for i := 0; i < 2; i++ {
		c.Increment()
	}
	c.Reset()
	Assert(t, c.Value() == 0, "Expected 0 after Reset. Got:", c.Value())
}

// END OMIT
