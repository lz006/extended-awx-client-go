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
	mType           string
	endpoint        string
	bearerTokenFile string
	port            string
	scheme          string
	targetPort      int
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

func (v *Variables) Endpoint() string {
	return v.endpoint
}

func (v *Variables) BearerTokenFile() string {
	return v.bearerTokenFile
}

func (v *Variables) Port() string {
	return v.port
}

func (v *Variables) Scheme() string {
	return v.scheme
}

func (v *Variables) TargetPort() int {
	return v.targetPort
}

func (v *Variables) TLSConfig() *TLSConfig {
	return v.tlsConf
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
