package job

import (
	"fmt"
	"time"
)

type JobTimeoutErr struct {
	timeout time.Duration
}

func (e JobTimeoutErr) Error() string {
	return fmt.Sprintf("job timed out after %v", e.timeout)
}
