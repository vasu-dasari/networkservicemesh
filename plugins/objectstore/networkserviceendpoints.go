// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package objectstore

import (
	"sync"

	"github.com/ligato/networkservicemesh/pkg/nsm/apis/netmesh"
)

// NetworkServiceEndpointsStore map stores all discovered Network Service Endpoint
// Objects with a key composed of a name and a namespace
type networkServiceEndpointsStore struct {
	networkServiceEndpoint map[meta]*netmesh.NetworkServiceEndpoint
	sync.RWMutex
}

// NewNetworkServiceEndpointsStore instantiates a new instance of a global
// NetworkServiceEndpoints store. It must be initialized before any controllers start.
func newNetworkServiceEndpointsStore() *networkServiceEndpointsStore {
	return &networkServiceEndpointsStore{
		networkServiceEndpoint: map[meta]*netmesh.NetworkServiceEndpoint{}}
}

// Add method adds descovered NetworkServiceEndpoint if it does not
// already exit in the store.
func (n *networkServiceEndpointsStore) Add(ns *netmesh.NetworkServiceEndpoint) {
	n.Lock()
	defer n.Unlock()

	key := meta{
		name:      ns.Metadata.Name,
		namespace: ns.Metadata.Namespace,
	}
	if _, ok := n.networkServiceEndpoint[key]; !ok {
		// Not in the store, adding it.
		n.networkServiceEndpoint[key] = ns
	}
}

// Delete method deletes removed NetworkServiceEndpoint object from the store.
func (n *networkServiceEndpointsStore) Delete(key meta) {
	n.Lock()
	defer n.Unlock()

	if _, ok := n.networkServiceEndpoint[key]; ok {
		delete(n.networkServiceEndpoint, key)
	}
}

// List method lists all known NetworkServiceEndpoint objects.
func (n *networkServiceEndpointsStore) List() []*netmesh.NetworkServiceEndpoint {
	n.Lock()
	defer n.Unlock()
	networkServiceEndpoints := make([]*netmesh.NetworkServiceEndpoint, 0)
	for _, ns := range n.networkServiceEndpoint {
		networkServiceEndpoints = append(networkServiceEndpoints, ns)
	}

	return networkServiceEndpoints
}
