package util

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

//UUID ...
func UUID() string {
	u := uuid.NewV4()
	return fmt.Sprintf("%s", u)
}
