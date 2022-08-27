package databases

import (
	"fmt"
)

func ForUpdate(queryIn string) string {
	return fmt.Sprintf("%s for update", queryIn)
}

func ForShare(queryIn string) string {
	return fmt.Sprintf("%s lock in share mode", queryIn)
}
