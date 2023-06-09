package uuid

import (
	"encoding/binary"
)

// NewUUID returns a Version 1 UUID based on the current NodeID and clock sequence,
// and the current time. If the NodeID has not been set by SetNodeID or SetNodeInterface
// then it will be set automatically. If the NodeID cannot be set NewUUID returns nil.
// If clock sequence has not been set by SetClockSequence then it will be automatically.
// If GetTime fails to return the current NewUUID returns nil and an error.

// In most cases, New should be used.

func NewUUID() (UUID, error) {
	var uuid UUID
	now, seq, err := GetTime()
	if err != nil {
		return uuid, err
	}

	timeLow := uint32(now & 0xffffffff)
	timeMid := uint16((now >> 32) & 0xffff)
	timeHi := uint16((now >> 48) & 0xfff)
	timeHi |= 0x1000 //version 1

	binary.BigEndian.PutUint32(uuid[0:], timeLow)
	binary.BigEndian.PutUint16(uuid[4:], timeMid)
	binary.BigEndian.PutUint16(uuid[6:], timeHi)
	binary.BigEndian.PutUint16(uuid[8:], seq)

	nodeMu.Lock()
	if nodeID == zeroID {
		SetNodeInterface("")
	}
	copy(uuid[10:], nodeID[:])
	nodeMu.Unlock()
	return uuid, nil
}
