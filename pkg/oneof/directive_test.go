package oneof

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirective(t *testing.T) {
	tests := []struct {
		name      string
		input     interface{}
		wantError error
	}{
		{
			"successful",
			map[string]interface{}{
				"foo": "bar",
			},
			nil,
		},
		{
			"failed - no options",
			map[string]interface{}{},
			ErrNoOptionSelected,
		},
		{
			"failed - too many options",
			map[string]interface{}{
				"foo": "bar",
				"baz": "boom",
			},
			ErrTooManyOptionsSelected,
		},
	}

	ctx := context.Background()
	resolver := func(ctx context.Context) (res interface{}, err error) {
		return nil, nil
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Directive(ctx, tt.input, resolver)
			assert.Equal(t, tt.wantError, err)
		})
	}
}
