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

package safe_counter

import (
	. "gitlab.com/pshirali/testing-in-go/samples/ex3/counter"
	. "gitlab.com/pshirali/testing-in-go/samples/ex3/unsafe_counter"
	"sync"
)

type UnsafeCounter interface {
	Resetter
	Incrementer
	IntValuer
}

// START OMIT
type safeCounter struct {
	mu sync.RWMutex
	uc UnsafeCounter
}

func (c *safeCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.uc.Reset()
}
func (c *safeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.uc.Increment()
}
func (c *safeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.uc.Value()
}
func NewSafeCounter() *safeCounter {
	return &safeCounter{uc: NewUnsafeCounter()}
}

// END OMIT
