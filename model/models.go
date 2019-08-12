package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

type Service struct {
	v1beta1.Service
}

func (s *Service) Metadata() *Metadata {
	return &Metadata{
		ObjectMeta: s.Service.ObjectMeta,
	}
}

func (s *Service) Spec() *ServiceSpec {
	return &ServiceSpec{
		ServiceSpec: s.Service.Spec,
	}
}

func (s *Service) ID() string {
	return string(s.ObjectMeta.UID)
}

type ServiceSpec struct {
	v1beta1.ServiceSpec
}

func (s *ServiceSpec) Template() *RevisionTemplateSpec {
	return &RevisionTemplateSpec{RevisionTemplateSpec: s.ServiceSpec.Template}
}

type RevisionTemplateSpec struct {
	v1beta1.RevisionTemplateSpec
}

func (s *RevisionTemplateSpec) Metadata() *Metadata {
	return &Metadata{
		ObjectMeta: s.RevisionTemplateSpec.ObjectMeta,
	}
}

func (s *RevisionTemplateSpec) Spec() *RevisionSpec {
	return &RevisionSpec{RevisionSpec: s.RevisionTemplateSpec.Spec}
}

type RevisionSpec struct {
	v1beta1.RevisionSpec
}

func (r *RevisionSpec) ContainerConcurrency() *int {
	c := int(r.RevisionSpec.ContainerConcurrency)
	return &c
}

type Metadata struct {
	metav1.ObjectMeta
}

func (m *Metadata) UID() string {
	return string(m.ObjectMeta.UID)
}
