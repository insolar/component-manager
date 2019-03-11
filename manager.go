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
	"log"
	"reflect"

	"github.com/pkg/errors"
)

// Logger interface provides methods for debug logging
type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}

type logger struct{}

func (l *logger) Debug(v ...interface{}) {
	log.Println(v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Manager provide methods to manage components lifecycle
type Manager struct {
	parent     *Manager
	components []interface{}
	logger     Logger
}

// NewManager creates new component manager
func NewManager(parent *Manager) *Manager {
	return &Manager{parent: parent, logger: &logger{}}
}

// Register components in Manager and inject required dependencies.
// Register can inject interfaces only, tag public struct fields with `inject:""`.
// If the injectable struct already has a value on the tagged field, the value WILL NOT be overridden.
func (m *Manager) Register(components ...interface{}) {
	m.components = append(m.components, components...)
}

// Inject components in Manager and inject required dependencies
// Inject can inject interfaces only, tag public struct fields with `inject:""`
func (m *Manager) Inject(components ...interface{}) {
	m.Register(components...)

	for _, componentMeta := range m.components {
		component := reflect.ValueOf(componentMeta).Elem()
		componentType := component.Type()
		m.logger.Debugf("ComponentManager: Inject component: %s", componentType.String())

		for i := 0; i < componentType.NumField(); i++ {
			fieldMeta := componentType.Field(i)
			if value, ok := fieldMeta.Tag.Lookup("inject"); ok && component.Field(i).IsNil() {
				if value == "subcomponent" && m.parent == nil {
					continue
				}
				m.logger.Debugf("ComponentManager: Component %s need inject: %s", componentType.String(), fieldMeta.Name)
				m.mustInject(component, fieldMeta)
			}
		}
	}
}

func (m *Manager) mustInject(component reflect.Value, fieldMeta reflect.StructField) {
	found := false
	if m.parent != nil {
		found = m.injectDependency(component, fieldMeta, m.parent.components)
	}
	found = found || m.injectDependency(component, fieldMeta, m.components)
	if found {
		return
	}

	panic(fmt.Sprintf(
		"Component %s injects not existing component with interface %s to field %s",
		component.Type().String(),
		fieldMeta.Type.String(),
		fieldMeta.Name,
	))
}

func (m *Manager) injectDependency(component reflect.Value, dependencyMeta reflect.StructField, components []interface{}) (injectFound bool) {
	for _, componentMeta := range components {
		componentType := reflect.ValueOf(componentMeta).Type()

		if componentType.Implements(dependencyMeta.Type) {
			field := component.FieldByName(dependencyMeta.Name)
			field.Set(reflect.ValueOf(componentMeta))

			m.logger.Debugf(
				"ComponentManager: Inject interface %s with %s: ",
				field.Type().String(),
				componentType.String(),
			)
			return true
		}
	}
	return false
}

func (m *Manager) isManaged(component interface{}) bool {
	// TODO: refactor this behavior
	if m.parent == nil {
		return true
	}
	for _, c := range m.parent.components {
		if c == component {
			return false
		}
	}
	return true
}

// Start invokes Start method of all components which implements Starter interface
func (m *Manager) Start(ctx context.Context) error {
	for _, c := range m.components {
		if !m.isManaged(c) {
			continue
		}
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
		if !m.isManaged(c) {
			continue
		}
		name := reflect.TypeOf(c).Elem().String()
		s, ok := c.(Initer)
		if !ok {
			m.logger.Debugf("ComponentManager: Component %s has no Init method", name)
			continue
		}
		m.logger.Debug("ComponentManager: Init component: ", name)
		err := s.Init(ctx)
		if err != nil {
			return errors.Wrap(err, "Failed to init components.")
		}
	}
	return nil
}

// Stop invokes Stop method of all components which implements Starter interface
func (m *Manager) GracefulStop(ctx context.Context) error {
	for i := len(m.components) - 1; i >= 0; i-- {
		if !m.isManaged(m.components[i]) {
			continue
		}
		name := reflect.TypeOf(m.components[i]).Elem().String()
		if s, ok := m.components[i].(GracefulStopper); ok {
			m.logger.Debug("ComponentManager: GracefulStop component: ", name)

			err := s.GracefulStop(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to gracefully stop components.")
			}
		} else {
			m.logger.Debugf("ComponentManager: Component %s has no GracefulStop method", name)
		}
	}
	return nil
}

// Stop invokes Stop method of all components which implements Starter interface
func (m *Manager) Stop(ctx context.Context) error {

	for i := len(m.components) - 1; i >= 0; i-- {
		if !m.isManaged(m.components[i]) {
			continue
		}
		name := reflect.TypeOf(m.components[i]).Elem().String()
		if s, ok := m.components[i].(Stopper); ok {
			m.logger.Debug("ComponentManager: Stop component: ", name)

			err := s.Stop(ctx)
			if err != nil {
				return errors.Wrap(err, "Failed to stop components.")
			}
		} else {
			m.logger.Debugf("ComponentManager: Component %s has no Stop method", name)
		}
	}
	return nil
}

// SetLogger sets custom logger
func (m *Manager) SetLogger(logger Logger) {
	m.logger = logger
}
