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

package component

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Interface1 interface {
	Method1()
}

type Interface2 interface {
	Method2()
}

type Component1 struct {
	field1     string
	Interface2 Interface2 `inject:""`
	asd        int
	started    bool
}

func (cm *Component1) Start(ctx context.Context) error {
	cm.Method1()
	cm.Interface2.Method2()
	return nil
}

func (cm *Component1) Method1() {
	fmt.Println("Component1.Method1 called")
}

type Component2 struct {
	field2     string
	Interface1 Interface1 `inject:""`
	dsa        string
	started    bool
}

type Component3 struct {
	Interface1 Interface1 `inject:""`
	Interface2 Interface1 `inject:""`
}

func (cm *Component2) Init(ctx context.Context) error {
	cm.field2 = "init done"
	return nil
}

func (cm *Component2) Start(ctx context.Context) error {
	cm.Interface1.Method1()
	cm.Method2()
	return nil
}

func (cm *Component2) GracefulStop(ctx context.Context) error {
	return nil
}

func (cm *Component2) Stop(ctx context.Context) error {
	return nil
}

func (cm *Component2) Method2() {
	fmt.Println("Component2.Method2 called")
}

func TestComponentManager_Inject(t *testing.T) {

	component1 := &Component1{}
	component2 := &Component2{}
	cm := NewManager(nil)
	cm.Inject(component1, component2)

	ctx := context.Background()
	require.NoError(t, cm.Init(ctx))
	assert.Equal(t, "init done", component2.field2)
	require.NoError(t, cm.Start(ctx))
	require.NoError(t, cm.GracefulStop(ctx))
	require.NoError(t, cm.Stop(ctx))
}
