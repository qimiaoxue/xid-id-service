package uuid

import (
	"encoding/binary"
	"fmt"
	"os"
)

// A Domain represents a Version 2 domin
type Domain byte

// Domain constants for DCE Security (Version 2) UUID.
const (
	Person = Domain(0)
	Group  = Domain(1)
	Org    = Domain(2)
)

// NewDCESecurity returns a DCE Security (Version 2) UUID.
//
// The domain should be one of Person, Group or Org.
// On a POSIX system the id should be the users UID for the Person
// domain and the users GID for the Group. The meaning of id for
// the domain Org or on non-POSIX systems is site defined.
func NewDCESecurity(domain Domain, id uint32) (UUID, error) {
	uuid, err := NewUUID()
	if err == nil {
		uuid[6] = (uuid[6] & 0x0f) | 0x20 // Version 2
		uuid[9] = byte(domain)
		binary.BigEndian.PutUint32(uuid[0:], id)
	}
	return uuid, err
}

// NewDCEPerson returns a DCE Security (Version 2) UUID in the person
// domain with the id returnd by os.Getuid.
//
// NewDCESecurity(Person, uint32(os.GetUid()))
func NewDCEPerson() (UUID, error) {
	return NewDCESecurity(Person, uint32(os.Geteuid()))
}

// Domain returns the domain for a Version 2 UUID. Domains are only defined
// for Version 2 UUIDs.
func (uuid UUID) Domain() Domain {
	return Domain(binary.BigEndian.Uint32(uuid[0:4]))
}

// ID returns the id for a Version 2 UUID. IDs are only defined for version 2 UUIDs.
func (uuid UUID) ID() uint32 {
	return binary.BigEndian.Uint32(uuid[0:4])
}

func (d Domain) String() string {
	switch d {
	case Person:
		return "Person"
	case Group:
		return "Group"
	case Org:
		return "Org"
	}
	return fmt.Sprintf("Domain%d", int(d))
}
