//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
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

package v1

import (
	"encoding/json"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
)

type DeploymentView struct{}

func (dv *DeploymentView) New(obj *types.Deployment, pl []*types.Pod) *Deployment {
	d := Deployment{}
	d.ID = obj.Meta.Name
	d.Meta = d.ToMeta(obj.Meta)
	d.Status = obj.Meta.Status
	d.State = d.ToState(obj.State)
	d.Spec = d.ToSpec(obj.Spec)
	d.Replicas = d.ToReplicas(obj.Replicas)
	d.Pods = d.ToPods(pl)
	return &d
}

func (di *Deployment) ToMeta(obj types.DeploymentMeta) DeploymentMeta {
	meta := DeploymentMeta{}
	meta.Name = obj.Name
	meta.Description = obj.Description
	meta.Version = obj.Version
	meta.SelfLink = obj.SelfLink
	meta.Namespace = obj.Namespace
	meta.Service = obj.Service
	meta.Status = obj.Status
	meta.Endpoint = obj.Endpoint
	meta.Updated = obj.Updated
	meta.Created = obj.Created

	return meta
}

func (di *Deployment) ToState(obj types.DeploymentState) DeploymentStateInfo {
	return DeploymentStateInfo{
		Active:    obj.Active,
		Provision: obj.Provision,
		Cancel:    obj.Cancel,
		Error:     obj.Error,
		Destroy:   obj.Destroy,
	}
}

func (di *Deployment) ToSpec(obj types.DeploymentSpec) DeploymentSpec {

	var spec = DeploymentSpec{
		Strategy: obj.Strategy,
		Triggers: obj.Triggers,
		Selector: obj.Selector,
		Replicas: obj.Replicas,
		Template: obj.Template,
	}

	return spec
}

func (di *Deployment) ToReplicas(obj types.DeploymentReplicas) DeploymentReplicasInfo {

	return DeploymentReplicasInfo{
		Total:     obj.Total,
		Provision: obj.Provision,
		Created:   obj.Created,
		Started:   obj.Started,
		Stopped:   obj.Stopped,
		Errored:   obj.Errored,
		Pulling:   obj.Pulling,
	}
}

func (di *Deployment) ToPods(obj []*types.Pod) []PodView {
	pods := make([]PodView, 0)
	for _, p := range obj {
		if p.Meta.Deployment == di.ID {
			pv := new(PodViewHelper)
			pods = append(pods, pv.New(p))
		}
	}
	return pods
}

func (di *Deployment) ToJson() ([]byte, error) {
	return json.Marshal(di)
}