package kngraphql

import (
	"context"

	"github.com/julz/kngraphql/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Services(ctx context.Context, namespace string) ([]*model.Service, error) {
	panic("not implemented")
}
func (r *queryResolver) ServiceByName(ctx context.Context, namespace string, name string) (*model.Service, error) {
	panic("not implemented")
}
