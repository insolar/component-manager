/*
 *    Copyright 2019 Insolar Technologies
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
package component_test

import (
	"github.com/insolar/component-manager"
)

type Buyer interface {
	BuyGoods(goods []string) error
}

type Supermarket struct {
}

func (Supermarket) BuyGoods(goods []string) error {
	return nil
}

type Customer struct {
	Buyer Buyer `inject:""`
	name string
}

func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func Example() {
	cm := component.NewManager()
	cm.Register(&Supermarket{})
	cm.Register(NewCustomer("Bob"), NewCustomer("Alice"))
	component.Inject()
}