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

// This file contains the implementation of the host type.

package eawx

// Group represents an AWX group.
//
type Group struct {
	id   int
	name string
	vars *Variables
}

type Variables struct {
	mType     string
	endpoints []*Endpoint
}

type Endpoint struct {
	endpoint        string
	bearerTokenFile string
	port            int32
	portName        string
	protocol        string
	scheme          string
	targetPort      int
	honorLabels     bool
	interval        string
	scrapeTimeout   string
	tlsConf         *TLSConfig
}

type TLSConfig struct {
	caFile             string
	hostname           string
	insecureSkipVerify bool
}

// Id returns the unique identifier of the group.
//
func (p *Group) Id() int {
	return p.id
}

// Name returns the name of the group.
//
func (p *Group) Name() string {
	return p.name
}

// GroupsArray returns array containing all related groups of the group.
//
func (p *Group) Vars() *Variables {
	return p.vars
}

func (v *Variables) MType() string {
	return v.mType
}

func (v *Variables) Endpoints() []*Endpoint {
	return v.endpoints
}

func (e *Endpoint) Endpoint() string {
	return e.endpoint
}

func (e *Endpoint) BearerTokenFile() string {
	return e.bearerTokenFile
}

func (e *Endpoint) Port() int32 {
	return e.port
}

func (e *Endpoint) PortName() string {
	return e.portName
}

func (e *Endpoint) Protocol() string {
	return e.protocol
}

func (e *Endpoint) Scheme() string {
	return e.scheme
}

func (e *Endpoint) TargetPort() int {
	return e.targetPort
}

func (e *Endpoint) HonorLabels() bool {
	return e.honorLabels
}

func (e *Endpoint) Interval() string {
	return e.interval
}

func (e *Endpoint) ScrapeTimeout() string {
	return e.scrapeTimeout
}

func (e *Endpoint) TLSConfig() *TLSConfig {
	return e.tlsConf
}

func (t *TLSConfig) CAFile() string {
	return t.caFile
}

func (t *TLSConfig) Hostname() string {
	return t.hostname
}

func (t *TLSConfig) InsecureSkipVerify() bool {
	return t.insecureSkipVerify
}
