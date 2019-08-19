package kngraphql

import (
	"context"

	"github.com/julz/kngraphql/model"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Service() ServiceResolver {
	return &serviceResolver{r}
}
func (r *Resolver) TrafficTarget() TrafficTargetResolver {
	return &trafficTargetResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateService(ctx context.Context, namespace string, input *CreateServiceInput) (*model.Service, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Services(ctx context.Context, namespace string) ([]*model.Service, error) {
	panic("not implemented")
}
func (r *queryResolver) ServiceByName(ctx context.Context, namespace string, name string) (*model.Service, error) {
	panic("not implemented")
}

type serviceResolver struct{ *Resolver }

func (r *serviceResolver) Revisions(ctx context.Context, obj *model.Service) ([]*model.Revision, error) {
	panic("not implemented")
}

type trafficTargetResolver struct{ *Resolver }

func (r *trafficTargetResolver) Revision(ctx context.Context, obj *v1beta1.TrafficTarget) (*model.Revision, error) {
	panic("not implemented")
}
