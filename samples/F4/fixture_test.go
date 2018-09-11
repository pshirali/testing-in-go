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

type Resource struct {}
func (r *Resource) Open() {
	fmt.Printf("%T @@ %p :: Called Open\n", r, r)
}
func (r *Resource) Print(str string) {
	fmt.Printf("%T @@ %p :: PrintMsg: %v\n", r, r, str)
}
func (r *Resource) Close() {
	fmt.Printf("%T @@ %p :: Called Close\n", r, r)
}

// Returns pointer to Resource and TEARDOWN function
// Note that the TEARDOWN has access to *Resource to tear it down
func ResourceHelper() (*Resource, func()) {
	r := &Resource{}
	r.Open()
	return r, func() {
		r.Close()
	}
}

//
//	Test
//
func TestFixtures(t *testing.T) {
	resource, teardown := ResourceHelper()
	defer teardown()
	resource.Print("<<< --- TEST CODE --- >>>")
}
