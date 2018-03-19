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

package pod

import (
	"context"
	"github.com/lastbackend/lastbackend/pkg/api/envs"
	"github.com/lastbackend/lastbackend/pkg/distribution"
	"github.com/lastbackend/lastbackend/pkg/distribution/errors"
	"github.com/lastbackend/lastbackend/pkg/distribution/types"
	"github.com/lastbackend/lastbackend/pkg/log"
)

func Provision(p *types.Pod) error {

	var (
		stg = envs.Get().GetStorage()

		memory = int64(0)
		node   *types.Node
	)

	nm := distribution.NewNodeModel(context.Background(), stg)

	if p.Meta.Node != "" {
		n, err := nm.Get(p.Meta.Node)
		if err != nil {
			log.Errorf("Node: find node err: %s", err.Error())
			return err
		}

		if n == nil {
			log.Errorf("Node: not found")
			return errors.New(errors.NodeNotFound)
		}

		if p.Spec.State.Destroy {
			if err := nm.RemovePod(n, p); err != nil {
				log.Errorf("Node: update pod spec err: %s", err.Error())
				return err
			}
			return nil
		}

		if err := nm.InsertPod(n, p); err != nil {
			log.Errorf("Node: update pod spec err: %s", err.Error())
			return err
		}
	}

	if p.Spec.State.Destroy {
		return nil
	}

	log.Debugf("Allocate node for pod: %s", p.Meta.Name)
	nodes, err := nm.List()
	if err != nil {
		log.Errorf("Node: allocate: get nodes error: %s", err.Error())
		return err
	}

	for _, c := range p.Spec.Template.Containers {
		memory += c.Resources.Quota.RAM
	}

	for _, n := range nodes {
		log.Debugf("Node: Allocate: available memory %d", n.State.Capacity)
		if n.State.Capacity.Memory > memory && n.Online {
			node = n
			break
		}
	}

	if node == nil {
		log.Error("Node: Allocate: Available node not found")
		p.Status.Stage = types.PodStageError
		p.Status.Message = errors.NodeNotFound

		if _, err := distribution.NewPodModel(context.Background(), stg).SetState(p); err != nil {
			log.Errorf("set pod state error: %s", err.Error())
			return err
		}

		return errors.New(errors.NodeNotFound)
	}

	if err := nm.InsertPod(node, p); err != nil {
		log.Errorf("Node: Pod spec add: insert spec to node err: %s", err.Error())
		return err
	}

	return nil
}
