package circlesdk

import "github.com/google/uuid"

// NewIdempotencyKey returns a random/valid new key that can be used to
// submit idempotent API requests.
// https://developers.circle.com/docs/a-note-on-idempotent-requests
func NewIdempotencyKey() string {
	return uuid.NewString()
}
