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
	"fmt"
	"reflect"
)

var logger = DefaultLogger{}

// Inject components in Manager and inject required dependencies
// Inject can inject interfaces only, tag public struct fields with `inject:""`
func  Inject(components ...interface{}) {

	for _, componentMeta := range components {
		component := reflect.ValueOf(componentMeta).Elem()
		componentType := component.Type()
		logger.Debugf("ComponentManager: Inject component: %s", componentType.String())

		for i := 0; i < componentType.NumField(); i++ {
			fieldMeta := componentType.Field(i)
			if _, ok := fieldMeta.Tag.Lookup("inject"); ok && component.Field(i).IsNil() {
				// if value == "subcomponent" && m.parent == nil {
				// 	continue
				// }
				logger.Debugf("ComponentManager: Component %s need inject: %s", componentType.String(), fieldMeta.Name)
				mustInject(component, fieldMeta, components)
			}
		}
	}
}

func  mustInject(component reflect.Value, fieldMeta reflect.StructField, components []interface{}) {
	found := false
	// if m.parent != nil {
	// 	found = m.injectDependency(component, fieldMeta, m.parent.components)
	// }
	found = found || injectDependency(component, fieldMeta, components)
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

func  injectDependency(component reflect.Value, dependencyMeta reflect.StructField, components []interface{}) (injectFound bool) {
	for _, componentMeta := range components {
		componentType := reflect.ValueOf(componentMeta).Type()

		if componentType.Implements(dependencyMeta.Type) {
			field := component.FieldByName(dependencyMeta.Name)
			field.Set(reflect.ValueOf(componentMeta))

			logger.Debugf(
				"ComponentManager: Inject interface %s with %s: ",
				field.Type().String(),
				componentType.String(),
			)
			return true
		}
	}
	return false
}


// func (m *Manager) isManaged(component interface{}) bool {
// 	// TODO: refactor this behavior
// 	if m.parent == nil {
// 		return true
// 	}
// 	for _, c := range m.parent.components {
// 		if c == component {
// 			return false
// 		}
// 	}
// 	return true
// }