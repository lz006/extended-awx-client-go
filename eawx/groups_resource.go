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

type GroupsResource struct {
	Resource
}

func NewGroupsResource(connection *Connection, path string) *GroupsResource {
	resource := new(GroupsResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *GroupsResource) Get() *GroupsGetRequest {
	request := new(GroupsGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *GroupsResource) Id(id int) *GroupResource {
	return NewGroupResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type GroupsGetRequest struct {
	Request
}

func (r *GroupsGetRequest) Filter(name string, value interface{}) *GroupsGetRequest {
	r.addFilter(name, value)
	return r
}

func (r *GroupsGetRequest) Send() (response *GroupsGetResponse, err error) {
	output := new(data.GroupsGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(GroupsGetResponse)
	response.count = output.Count
	response.previous = output.Previous
	response.next = output.Next
	response.results = make([]*Group, len(output.Results))
	for i := 0; i < len(output.Results); i++ {
		response.results[i] = new(Group)
		response.results[i].id = output.Results[i].Id
		response.results[i].name = output.Results[i].Name
		var vars *data.Variables
		err = yaml.Unmarshal([]byte(output.Results[i].Vars), &vars)
		if err != nil {
			glog.Warningf("Error parsing: %v", err)
		}

		if vars != nil {

			var endpoints []*Endpoint

			for _, endpoint := range vars.Endpoints {

				tmpEndpoint := &Endpoint{
					endpoint:        endpoint.Endpoint,
					bearerTokenFile: endpoint.BearerTokenFile,
					port:            endpoint.Port,
					portName:        endpoint.PortName,
					protocol:        endpoint.Protocol,
					scheme:          endpoint.Scheme,
					targetPort:      endpoint.TargetPort,
					honorLabels:     endpoint.HonorLabels,
					interval:        endpoint.Interval,
					scrapeTimeout:   endpoint.ScrapeTimeout,
				}

				if &endpoint.TLSConf != nil {
					tmpTLSConfig := &TLSConfig{
						caFile:             endpoint.TLSConf.CAFile,
						hostname:           endpoint.TLSConf.Hostname,
						insecureSkipVerify: endpoint.TLSConf.InsecureSkipVerify,
					}

					tmpEndpoint.tlsConf = tmpTLSConfig
				}

				endpoints = append(endpoints, tmpEndpoint)
			}

			response.results[i].vars = &Variables{
				mType:     vars.Type,
				endpoints: endpoints,
			}
		}
	}
	return
}

type GroupsGetResponse struct {
	ListGetResponse

	results []*Group
}

func (r *GroupsGetResponse) Results() []*Group {
	return r.results
}
