/*
 *    Copyright 2018 Insolar
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package component

import (
	"testing"
)

func BenchmarkManager_Inject(b *testing.B) {
	c1 := &Component1{}
	c2 := &Component2{}
	c3 := &Component3{}
	c4 := &Component2{}
	c5 := &Component1{}
	c6 := &Component3{}
	c7 := &Component1{}
	c8 := &Component3{}
	c9 := &Component1{}
	c10 := &Component3{}


	for i := 0; i < b.N; i++ {
		cm := Manager{}
		cm.Inject(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10)
	}

}

func BenchmarkManager_Inject2(b *testing.B) {
	c1 := &Component1{}
	c2 := &Component2{}


	for i := 0; i < b.N; i++ {
		cm := Manager{}
		cm.Inject(c1, c2)
	}
}

func BenchmarkManager_DirectInject(b *testing.B) {
	c1 := &Component1{}
	c2 := &Component2{}


	for i := 0; i < b.N; i++ {
		cm := Manager{}
		cm.Register(c1, c2)
		c1.Interface2 = c2
		c2.Interface1 = c1
	}
}

