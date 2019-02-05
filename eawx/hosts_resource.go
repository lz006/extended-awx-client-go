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

// This file contains the implementation of the resource that manages the collection of
// projects.

package eawx

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/lz006/extended-awx-client-go/eawx/internal/data"
	yaml "gopkg.in/yaml.v2"
)

type HostsResource struct {
	Resource
}

func NewHostsResource(connection *Connection, path string) *HostsResource {
	resource := new(HostsResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *HostsResource) Get() *HostsGetRequest {
	request := new(HostsGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *HostsResource) Id(id int) *HostResource {
	return NewHostResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type HostsGetRequest struct {
	Request
}

func (r *HostsGetRequest) Filter(name string, value interface{}) *HostsGetRequest {
	r.addFilter(name, value)
	return r
}

func (r *HostsGetRequest) Send() (response *HostsGetResponse, err error) {
	output := new(data.HostsGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(HostsGetResponse)
	response.count = output.Count
	response.previous = output.Previous
	response.next = output.Next
	response.results = make([]*Host, len(output.Results))
	for i := 0; i < len(output.Results); i++ {
		response.results[i] = new(Host)
		response.results[i].id = output.Results[i].Id
		response.results[i].name = output.Results[i].Name

		fromGroupArray := output.Results[i].InventoryGroups.HostGroups.GroupArray
		toStringArray := make([]*string, len(fromGroupArray))
		for j := 0; j < len(fromGroupArray); j++ {
			toStringArray[j] = &fromGroupArray[j].Name
		}
		response.results[i].groups = toStringArray

		var vars *data.HostVariables
		err = yaml.Unmarshal([]byte(output.Results[i].HostVars), &vars)
		if err != nil {
			glog.Warningf("Error parsing: %v", err)
		}
		if vars != nil {
			response.results[i].ip = vars.IP
		}
	}
	return
}

type HostsGetResponse struct {
	ListGetResponse

	results []*Host
}

func (r *HostsGetResponse) Results() []*Host {
	return r.results
}
