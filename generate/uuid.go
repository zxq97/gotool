package generate

import (
	"github.com/google/uuid"
)

func UUIDStr() string {
	u, _ := uuid.NewUUID()
	return u.String()
}

func UUIDUint() uint32 {
	u, _ := uuid.NewUUID()
	return u.ID()
}

func UUIDLong() int64 {
	u, _ := uuid.NewUUID()
	return int64(u.Time())
}
