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
	"reflect"

	"github.com/pkg/errors"
)

// Manager provide methods to manage components lifecycle
type Manager struct {
	parent     *Manager
	components []interface{}
	logger     Logger
}

// NewManager creates new component manager with default logger
func NewManager() *Manager {
	return &Manager{logger: &DefaultLogger{}}
}

// Register components in Manager and inject required dependencies.
// Register can inject interfaces only, tag public struct fields with `inject:""`.
// If the injectable struct already has a value on the tagged field, the value WILL NOT be overridden.
func (m *Manager) Register(components ...interface{}) {
	m.components = append(m.components, components...)
	// в момент регистрации получать имена? name := reflect.TypeOf(c).Elem().String()
	// спазу проверять на наличие интерфейсов и разбивать по кучкам?
}



// Start invokes Start method of all components which implements Starter interface
func (m *Manager) Start(ctx context.Context) error {
	for _, c := range m.components {
		// if !m.isManaged(c) {
		// 	continue
		// }
		name := reflect.TypeOf(c).Elem().String()
		if s, ok := c.(Starter); ok {
			m.logger.Debug("ComponentManager: Start component: ", name)
			err := s.Start(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to start components.")
			}
			m.logger.Debugf("ComponentManager: Component %s started ", name)
		} else {
			m.logger.Debugf("ComponentManager: Component %s has no Start method", name)
		}
	}
	return nil
}

// Init invokes Init method of all components which implements Initer interface
func (m *Manager) Init(ctx context.Context) error {
	for _, c := range m.components {
		// if !m.isManaged(c) {
		// 	continue
		// }
		name := reflect.TypeOf(c).Elem().String()
		if s, ok := c.(Initer); ok {
			m.logger.Debug("ComponentManager: Init component: ", name)

			err := s.Init(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to init components.")
			}
		}
	}
	return nil
}

// Stop invokes Stop method of all components which implements Starter interface
func (m *Manager) GracefulStop(ctx context.Context) error {
	for i := len(m.components) - 1; i >= 0; i-- {
		// if !m.isManaged(m.components[i]) {
		// 	continue
		// }
		name := reflect.TypeOf(m.components[i]).Elem().String()
		if s, ok := m.components[i].(GracefulStopper); ok {
			m.logger.Debug("ComponentManager: GracefulStop component: ", name)

			err := s.GracefulStop(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to gracefully stop components.")
			}
		}
	}
	return nil
}

// Stop invokes Stop method of all components which implements Starter interface
func (m *Manager) Stop(ctx context.Context) error {

	for i := len(m.components) - 1; i >= 0; i-- {
		// if !m.isManaged(m.components[i]) {
		// 	continue
		// }
		name := reflect.TypeOf(m.components[i]).Elem().String()
		if s, ok := m.components[i].(Stopper); ok {
			m.logger.Debug("ComponentManager: Stop component: ", name)

			err := s.Stop(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to stop components.")
			}
		}
	}
	return nil
}

// SetLogger sets custom DefaultLogger
func (m *Manager) SetLogger(logger Logger) {
	m.logger = logger
}
