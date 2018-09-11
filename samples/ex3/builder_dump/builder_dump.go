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

package main

// NOTE: The dot import has been used here to condense the data on the slides.
//       AVOID: '. <import-path' in your code.
import (
	"fmt"
	. "gitlab.com/pshirali/testing-in-go/samples/ex3/counter"
	. "gitlab.com/pshirali/testing-in-go/samples/ex3/test_suite"
)

func dumpInstances(count int, counterBuilder func() Counter) {
	var inst Counter
	fmt.Printf("Dumping %v instances of '%T'\n", count, counterBuilder)
	for i := 0; i < count; i++ {
		inst = counterBuilder()
		fmt.Printf("Type: %T - Address: %p\n", inst, inst)
	}
}

func main() {
	dumpInstances(5, UnsafeCounterBuilder)
	dumpInstances(5, SafeCounterBuilder)
}
