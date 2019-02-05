/*
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

// This file contains the data structures used for sending and receiving projects.

package data

type Host struct {
	Id              int           `json:"id,omitempty"`
	Name            string        `json:"name,omitempty"`
	InventoryGroups SummaryFields `json:"summary_fields,omitempty"`
	HostVars        string        `json:"variables,omitempty"`
}

type SummaryFields struct {
	HostGroups Groups `json:"groups,omitempty"`
}

type Groups struct {
	GroupArray []Group `json:"results,omitempty"`
}

// Used when unmarshalling Host.HostVars which contains a yaml string
type HostVariables struct {
	IP string `yaml:"ansible_host,omitempty"`
}

type HostGetResponse struct {
	Host
}
