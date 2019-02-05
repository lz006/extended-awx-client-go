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

// This file contains the implementation of the resource that manages a specific project.

package eawx

import (
	"github.com/golang/glog"
	"github.com/lz006/extended-awx-client-go/eawx/internal/data"
	yaml "gopkg.in/yaml.v2"
)

type HostResource struct {
	Resource
}

func NewHostResource(connection *Connection, path string) *HostResource {
	resource := new(HostResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *HostResource) Get() *HostGetRequest {
	request := new(HostGetRequest)
	request.resource = &r.Resource
	return request
}

type HostGetRequest struct {
	Request
}

func (r *HostGetRequest) Send() (response *HostGetResponse, err error) {
	output := new(data.HostGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(HostGetResponse)
	response.result = new(Host)
	response.result.id = output.Id
	response.result.name = output.Name

	fromGroupArray := output.InventoryGroups.HostGroups.GroupArray
	toStringArray := make([]*string, len(fromGroupArray))
	for j := 0; j < len(fromGroupArray); j++ {
		(*toStringArray[j]) = fromGroupArray[j].Name
	}
	response.result.groups = toStringArray

	var vars *data.HostVariables
	err = yaml.Unmarshal([]byte(output.HostVars), &vars)
	if err != nil {
		glog.Warningf("Error parsing: %v", err)
	}

	if vars != nil {
		response.result.ip = vars.IP
	}

	return
}

type HostGetResponse struct {
	result *Host
}

func (r *HostGetResponse) Result() *Host {
	return r.result
}
