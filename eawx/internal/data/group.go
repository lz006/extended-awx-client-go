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

type Group struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Vars string `json:"variables,omitempty"`
}

// Used when unmarshalling Group.Vars which contains a yaml string
type Variables struct {
	Type      string     `yaml:"mo_type,omitempty"`
	Endpoints []Endpoint `yaml:"mo_endpoints,omitempty"`
}

type Endpoint struct {
	Endpoint        string `yaml:"endpoint,omitempty"`
	BearerTokenFile string `yaml:"bearerTokenFile,omitempty"`
	Port            int32  `yaml:"port,omitempty"`
	PortName        string `yaml:"portName,omitempty"`
	Protocol        string `yaml:"protocol,omitempty"`
	Scheme          string `yaml:"scheme,omitempty"`
	TargetPort      int    `yaml:"targetPort,omitempty"`

	HonorLabels   bool   `yaml:"honorLabels,omitempty"`
	Interval      string `yaml:"interval,omitempty"`
	ScrapeTimeout string `yaml:"scrapeTimeout,omitempty"`

	TLSConf TLSConfig `yaml:"tlsConfig,omitempty"`
}

// Used when unmarshalling Variables.TLSConf which contains a yaml string
type TLSConfig struct {
	CAFile             string `yaml:"caFile,omitempty"`
	Hostname           string `yaml:"hostname,omitempty"`
	InsecureSkipVerify bool   `yaml:"insecureSkipVerify,omitempty"`
}

type GroupGetResponse struct {
	Group
}
