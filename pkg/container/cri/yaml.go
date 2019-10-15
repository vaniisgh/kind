/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cri

import (
	"fmt"
	"strings"
)

/*
Custom YAML (de)serialization for these types
*/

// UnmarshalYAML implements custom decoding YAML
// https://godoc.org/gopkg.in/yaml.v3
func (m *Mount) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type Alias Mount
	aux := &struct {
		Propagation string `yaml:"propagation"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := unmarshal(&aux); err != nil {
		return err
	}
	// if unset, will fallback to the default (0)
	if aux.Propagation != "" {
		val, ok := MountPropagationNameToValue[aux.Propagation]
		if !ok {
			return fmt.Errorf("unknown propagation value: %s", aux.Propagation)
		}
		m.Propagation = MountPropagation(val)
	}
	return nil
}

// UnmarshalYAML implements custom decoding YAML
// https://godoc.org/gopkg.in/yaml.v3
func (p *PortMapping) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type Alias PortMapping
	aux := &struct {
		Protocol string `json:"protocol"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := unmarshal(&aux); err != nil {
		return err
	}
	if aux.Protocol != "" {
		val, ok := PortMappingProtocolNameToValue[strings.ToUpper(aux.Protocol)]
		if !ok {
			return fmt.Errorf("unknown protocol value: %s", aux.Protocol)
		}
		p.Protocol = PortMappingProtocol(val)
	}
	return nil
}
