// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package kngraphql

import (
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

type RouteSpec struct {
	Traffic []*v1beta1.TrafficTarget `json:"traffic"`
}
