/*

Package kafka a provides high level client API for Apache Kafka.

Use 'Broker' for node connection management, 'Producer' for sending messages,
and 'Consumer' for fetching. All those structures implement Client, Consumer
and Producer interface, that is also implemented in kafkatest package.

*/
package kafka

import (
	"context"
	"fmt"
)

// isCanceled returns true when the given error is
// a context canceled or context deadline exceeded error.
func isCanceled(err error) bool {
	return err == context.Canceled || err == context.DeadlineExceeded
}

type responseWaiterError struct {
	Err error
}

func (e *responseWaiterError) Error() string {
	return fmt.Sprintf("waiting for response: %s", e.Err.Error())
}

// isResponseWaiter returns true when the given error is of type responseWaiterError.
// Returns false otherwise.
func isResponseWaiter(err error) bool {
	_, ok := err.(*responseWaiterError)
	return ok
}
