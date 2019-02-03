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

package awx

// Host represents an AWX host.
//
type Host struct {
	id     int
	name   string
	groups []*string
	ipv4   string
}

// Id returns the unique identifier of the host.
//
func (p *Host) Id() int {
	return p.id
}

// Name returns the name of the host.
//
func (p *Host) Name() string {
	return p.name
}

// GroupsArray returns array containing all related groups of the host.
//
func (p *Host) GroupsArray() []*string {
	return p.groups
}

// IPV4 returns the ipv4 address of the host
//
func (p *Host) IPV4() string {
	return p.ipv4
}
