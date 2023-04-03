package uuid

import (
	"sync"
)

var (
	nodeMu sync.Mutex
	ifname string  // name of interface being used
	nodeID [6]byte // hardware for version 1 UUIDS
	zeroID [6]byte // nodeID with only 0's
)

// NodeInterface returns the name of the interface from which the NodeID was delived.
// The interface "user" is returned if the NodeID was set by SetNodeID.
func NodeInterface() string {
	defer nodeMu.Unlock()
	nodeMu.Lock()
	return ifname
}

// SetNodeInterface selects the hardware address to be used for Version 1 UUIDs.
// If name is "" then the first usable interface found will be used or a random
// Node ID will be generated. If a named interface cannot be found then false is returned.
//
// SetNodeInterface never fails when name is "".

func SetNodeInterface(name string) bool {
	defer nodeMu.Unlock()
	nodeMu.Lock()
	return setNodeInterface(name)
}

func setNodeInterface(name string) bool {
	iname, addr := getHardwareInterface(name) // null implementation for js
	if iname != "" && addr != nil {
		ifname = iname
		copy(nodeID[:], addr)
		return true
	}

	// we found no interfacce with a valid hardware address. If name does not specify
	// specific interface generate a random Node ID
	// (sectiong 4.1.6)
	if name == "" {
		ifname = "random"
		randomBits(nodeID[:])
		return true
	}
	return false
}

// NodeID returns a slice of a copy of the current Node ID, setting the Node ID
// if not already set.
func NodeID() []byte {
	defer nodeMu.Unlock()
	nodeMu.Lock()
	if nodeID == zeroID {
		setNodeInterface("")
	}
	nid := nodeID
	return nid[:]
}

// SetNodeID sets the Node ID to be used for version 1 UUIDS. The first 6 bytes of id
// are used. If id is less than 6 bytes then false is returned and the Node ID is not set.
func SetNodeID(id []byte) bool {
	if len(id) < 6 {
		return false
	}
	defer nodeMu.Unlock()
	nodeMu.Lock()
	copy(nodeID[:], id)
	ifname = "user"
	return true
}

// NodeID returns the 6 byte node id encoded in uuid. If returns nil if uuid is
// not valid. The NodeID is only well defined for version 1 and 2 UUIDs.
func (uuid UUID) NodeID() []byte {
	var node [6]byte
	copy(node[:], uuid[10:])
	return node[:]
}
