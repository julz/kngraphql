package resolver

import (
	"context"
	"fmt"

	"github.com/julz/kngraphql"
	"github.com/julz/kngraphql/model"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/serving/pkg/client/clientset/versioned"
)

type Resolver struct {
	client *versioned.Clientset
}

func NewResolver(client *versioned.Clientset) *Resolver {
	return &Resolver{
		client: client,
	}
}

func (r *Resolver) Query() kngraphql.QueryResolver {
	return r
}

func (r *Resolver) Services(ctx context.Context, namespace string) (result []*model.Service, err error) {
	ss, err := r.client.ServingV1beta1().Services(namespace).List(v1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list: %s", err)
	}

	for _, s := range ss.Items {
		result = append(result, &model.Service{Service: s})
	}

	return result, nil
}

func (r *Resolver) ServiceByName(ctx context.Context, namespace string, name string) (*model.Service, error) {
	s, err := r.client.ServingV1beta1().Services(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get: %s", err)
	}

	return &model.Service{
		Service: *s,
	}, nil
}
