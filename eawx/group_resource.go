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

type GroupResource struct {
	Resource
}

func NewGroupResource(connection *Connection, path string) *GroupResource {
	resource := new(GroupResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *GroupResource) Get() *GroupGetRequest {
	request := new(GroupGetRequest)
	request.resource = &r.Resource
	return request
}

type GroupGetRequest struct {
	Request
}

func (r *GroupGetRequest) Send() (response *GroupGetResponse, err error) {
	output := new(data.GroupGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(GroupGetResponse)
	response.result = new(Group)
	response.result.id = output.Id
	response.result.name = output.Name
	var vars *data.Variables
	err = yaml.Unmarshal([]byte(output.Vars), &vars)
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

		response.result.vars = &Variables{
			mType:     vars.Type,
			endpoints: endpoints,
		}
	}
	return
}

type GroupGetResponse struct {
	result *Group
}

func (r *GroupGetResponse) Result() *Group {
	return r.result
}
