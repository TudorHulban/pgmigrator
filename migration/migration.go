package migration

import (
	"crypto/md5"
	"fmt"
)

type Migration struct {
	ID  string
	SQL string
}

func (m *Migration) MD5() string {
	return fmt.Sprintf(
		"%x",
		md5.Sum([]byte(m.SQL)),
	)
}
