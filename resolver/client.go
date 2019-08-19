package resolver

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/julz/kngraphql"
	"github.com/julz/kngraphql/model"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
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

func (r *Resolver) Mutation() kngraphql.MutationResolver {
	return r
}

func (r *Resolver) Service() kngraphql.ServiceResolver {
	return r
}

func (r *Resolver) TrafficTarget() kngraphql.TrafficTargetResolver {
	return &trafficTargetResolver{r}
}

func (r *Resolver) CreateService(ctx context.Context, namespace string, input *kngraphql.CreateServiceInput) (*model.Service, error) {
	toCoreContainers := func(ctrs []*kngraphql.ContainerInput) (result []v1.Container) {
		for _, c := range ctrs {
			result = append(result, v1.Container{
				Image:   c.Image,
				Command: c.Command,
				Args:    c.Args,
			})
		}

		return result
	}

	s, err := r.client.ServingV1beta1().Services(namespace).Create(&v1beta1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: input.Name,
		},
		Spec: v1beta1.ServiceSpec{
			ConfigurationSpec: v1beta1.ConfigurationSpec{
				Template: v1beta1.RevisionTemplateSpec{
					Spec: v1beta1.RevisionSpec{
						PodSpec: v1.PodSpec{
							Containers: toCoreContainers(input.Containers),
						},
					},
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	return &model.Service{Service: *s}, nil
}

func (r *Resolver) Services(ctx context.Context, namespace string) (result []*model.Service, err error) {
	ss, err := r.client.ServingV1beta1().Services(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list: %s", err)
	}

	for _, s := range ss.Items {
		result = append(result, &model.Service{Service: s})
	}

	return result, nil
}

func (r *Resolver) ServiceByName(ctx context.Context, namespace string, name string) (*model.Service, error) {
	s, err := r.client.ServingV1beta1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("get: %s", err)
	}

	return &model.Service{
		Service: *s,
	}, nil
}

func (r *Resolver) Revisions(ctx context.Context, service *model.Service) (result []*model.Revision, err error) {
	ss, err := r.client.ServingV1beta1().Revisions(service.Service.ObjectMeta.Namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("serving.knative.dev/service=%s", service.Name),
	})

	if err != nil {
		return nil, err
	}

	for _, s := range ss.Items {
		result = append(result, &model.Revision{Revision: s})
	}

	return result, nil
}

type trafficTargetResolver struct{ *Resolver }

func (r *trafficTargetResolver) Revision(ctx context.Context, target *v1beta1.TrafficTarget) (*model.Revision, error) {
	if target.RevisionName == "" {
		return nil, nil
	}

	revision, err := r.Resolver.client.ServingV1beta1().Revisions(queryArg(ctx, "namespace").(string)).Get(target.RevisionName, metav1.GetOptions{})
	return &model.Revision{Revision: *revision}, err
}

func queryArg(ctx context.Context, key string) interface{} {
	req := graphql.GetResolverContext(ctx)
	var args map[string]interface{}
	for parentCtx := req.Parent; parentCtx.Parent != nil; parentCtx = parentCtx.Parent {
		args = parentCtx.Args
	}

	return args[key]
}
