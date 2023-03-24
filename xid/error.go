package xid

const (
	// ErrIncalidID is returned when trying to unmarshal an invalid ID
	ErrInvalidID strErr = "xid: invalid ID"
)

type strErr string

func (err strErr) Error() string { return string(err) }
