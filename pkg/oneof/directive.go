package oneof

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

func Directive(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	val, err := next(ctx)
	if err != nil {
		panic(err)
	}

	v := obj.(map[string]interface{})
	c := len(v)
	if c == 0 {
		return nil, ErrNoOptionSelected
	} else if c > 1 {
		return nil, ErrTooManyOptionsSelected
	}

	return val, nil
}
