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

package runtime

import (
	"context"
	"github.com/lastbackend/lastbackend/tools/logger"
	"time"

	"github.com/lastbackend/lastbackend/internal/pkg/models"
	)

const (
	logServicePrefix = "node:runtime:service"
)

func (r Runtime) serviceStart(ctx context.Context, pod string, m *models.ContainerManifest, status *models.PodStatus) error {
	log := logger.WithContext(context.Background())
	var (
		err error
		c   = new(models.PodContainer)
	)

	c.ID, err = r.cri.Create(ctx, m)
	if err != nil {
		switch err {
		case context.Canceled:
			log.Errorf("%s stop creating container: %s", logServicePrefix, err.Error())
			return nil
		}

		log.Errorf("%s can-not create container: %s", logServicePrefix, err)

		c.State.Error = models.PodContainerStateError{
			Error:   true,
			Message: err.Error(),
			Exit: models.PodContainerStateExit{
				Timestamp: time.Now().UTC(),
			},
		}
		return err
	}

	status.Runtime.Services[c.ID] = c

	if err := r.containerInspect(context.Background(), c); err != nil {
		log.Errorf("%s inspect container after create: err %s", logServicePrefix, err.Error())
		r.PodClean(context.Background(), status)
		return err
	}

	//==========================================================================
	// Start container =========================================================
	//==========================================================================

	c.State.Created = models.PodContainerStateCreated{
		Created: time.Now().UTC(),
	}

	r.state.Pods().SetPod(pod, status)
	log.Debugf("%s container created: %s", logServicePrefix, c.ID)

	if err := r.cri.Start(ctx, c.ID); err != nil {

		log.Errorf("%s can-not start container: %s", logServicePrefix, err)

		switch err {
		case context.Canceled:
			log.Errorf("%s stop starting container err: %s", logServicePrefix, err.Error())
			return nil
		}

		c.State.Error = models.PodContainerStateError{
			Error:   true,
			Message: err.Error(),
			Exit: models.PodContainerStateExit{
				Timestamp: time.Now().UTC(),
			},
		}

		return err
	}

	log.Debugf("%s container started: %s", logServicePrefix, c.ID)

	if err := r.containerInspect(context.Background(), c); err != nil {
		log.Errorf("%s inspect container after create: err %s", logServicePrefix, err.Error())
		return err
	}

	c.Ready = true
	c.State.Started = models.PodContainerStateStarted{
		Started:   true,
		Timestamp: time.Now().UTC(),
	}

	info, err := r.cri.Inspect(ctx, c.ID)
	if err != nil {
		switch err {
		case context.Canceled:
			log.Errorf("Stop inspect container err: %v", err)
			return nil
		}
		log.Errorf("Can-not inspect container: %v", err)
		return err
	}

	if status.Network.PodIP == "" {
		status.Network.PodIP = info.Network.IPAddress
	}

	r.state.Pods().SetPod(pod, status)
	return nil
}

func (r Runtime) serviceStop() error {

	return nil
}

func (r Runtime) serviceRestart() error {

	return nil
}

func (r Runtime) serviceRemove() error {
	return nil
}
