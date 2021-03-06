//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2020] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package request

import (
	"encoding/json"
	"github.com/lastbackend/lastbackend/internal/pkg/models"
	"gopkg.in/yaml.v3"
)

type RouteManifest struct {
	Meta RouteManifestMeta `json:"meta,omitempty" yaml:"meta,omitempty"`
	Spec RouteManifestSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type RouteManifestMeta struct {
	RuntimeMeta `yaml:",inline"`
}

// swagger:model request_route_create
type RouteManifestSpec struct {
	Port     uint16                         `json:"port" yaml:"port"`
	Type     string                         `json:"type" yaml:"type"`
	Endpoint string                         `json:"endpoint" yaml:"endpoint"`
	Rules    []RouteManifestSpecRulesOption `json:"rules" yaml:"rules"`
}

// swagger:ignore
// swagger:model request_route_remove
type RouteRemoveOptions struct {
	Force bool `json:"force"`
}

// swagger:model request_route_rules
type RouteManifestSpecRulesOption struct {
	Service string `json:"service" yaml:"service"`
	Path    string `json:"path" yaml:"path"`
	Port    int    `json:"port" yaml:"port"`
}

func (r *RouteManifest) FromJson(data []byte) error {
	return json.Unmarshal(data, r)
}

func (r *RouteManifest) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RouteManifest) FromYaml(data []byte) error {
	return yaml.Unmarshal(data, r)
}

func (r *RouteManifest) ToYaml() ([]byte, error) {
	return yaml.Marshal(r)
}

func (r *RouteManifest) SetRouteMeta(route *models.Route) {

	if route.Meta.Name == models.EmptyString {
		route.Meta.Name = *r.Meta.Name
	}

	if r.Meta.Description != nil {
		route.Meta.Description = *r.Meta.Description
	}

	if r.Meta.Labels != nil {
		for k, v := range r.Meta.Labels {
			route.Meta.Labels[k] = v
		}
	}
}

func (r *RouteManifest) SetRouteSpec(route *models.Route, ns *models.Namespace, svc *models.ServiceList) {

	var sl = make(map[string]*models.Service)
	for _, s := range svc.Items {
		sl[s.Meta.Name] = s
	}

	if r.Spec.Endpoint != route.Spec.Endpoint {
		route.Spec.Endpoint = r.Spec.Endpoint
	}

	if r.Spec.Port != route.Spec.Port {
		route.Spec.Port = r.Spec.Port
	}

	route.Spec.Rules = make([]models.RouteRule, 0)
	for _, rs := range r.Spec.Rules {

		if rs.Service == models.EmptyString || rs.Port == 0 {
			continue
		}

		if _, ok := sl[rs.Service]; !ok {
			continue
		}

		route.Spec.Rules = append(route.Spec.Rules, models.RouteRule{
			Upstream: sl[rs.Service].Meta.Endpoint,
			Service:  rs.Service,
			Port:     rs.Port,
			Path:     rs.Path,
		})

	}

	route.Spec.State = models.StateProvision
}
