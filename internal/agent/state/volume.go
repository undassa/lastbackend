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

package state

import (
	"context"
	"github.com/lastbackend/lastbackend/internal/pkg/models"
	"github.com/lastbackend/lastbackend/tools/logger"
	"sync"
)

const logVolumePrefix = "state:volume:> "

type VolumesState struct {
	lock     sync.RWMutex
	volumes  map[string]models.VolumeStatus
	local    map[string]bool
	watchers map[chan string]bool
	claims   map[string]models.VolumeClaim
}

func (s *VolumesState) dispatch(pod string) {
	for w := range s.watchers {
		w <- pod
	}
}

func (s *VolumesState) Watch(watcher chan string, done chan bool) {
	s.watchers[watcher] = true
	defer delete(s.watchers, watcher)
	<-done
}

func (s *VolumesState) GetVolumes() map[string]models.VolumeStatus {
	log := logger.WithContext(context.Background())
	log.Debugf("%s get volumes", logVolumePrefix)
	return s.volumes
}

func (s *VolumesState) SetVolumes(key string, volumes []*models.VolumeStatus) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s set volumes: %#v", logVolumePrefix, volumes)
	for _, vol := range volumes {
		s.volumes[key] = *vol
	}
}

func (s *VolumesState) GetVolume(key string) *models.VolumeStatus {
	log := logger.WithContext(context.Background())
	log.Debugf("%s get volume: %s", logVolumePrefix, key)
	s.lock.Lock()
	defer s.lock.Unlock()
	v, ok := s.volumes[key]
	if !ok {
		return nil
	}
	return &v
}

func (s *VolumesState) AddVolume(key string, v *models.VolumeStatus) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s add volume: %s > %s", logVolumePrefix, key, v.State)
	s.SetVolume(key, v)
}

func (s *VolumesState) SetVolume(key string, v *models.VolumeStatus) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s set volume: %s > %s", logVolumePrefix, key, v.State)
	s.lock.Lock()
	s.volumes[key] = *v
	s.lock.Unlock()
	s.dispatch(key)
}

func (s *VolumesState) DelVolume(key string) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s del volume: %#v", logVolumePrefix, key)
	s.lock.Lock()
	if _, ok := s.volumes[key]; ok {
		delete(s.volumes, key)
	}
	s.lock.Unlock()
	s.dispatch(key)
}

func (s *VolumesState) GetClaim(key string) *models.VolumeClaim {
	log := logger.WithContext(context.Background())
	log.Debugf("%s get claim: %s", logVolumePrefix, key)
	v, ok := s.claims[key]
	if !ok {
		return nil
	}
	return &v
}

func (s *VolumesState) AddClaim(key string, vc *models.VolumeClaim) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s add claim: %s", logVolumePrefix, key)
	s.SetClaim(key, vc)
}

func (s *VolumesState) SetClaim(key string, vc *models.VolumeClaim) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s set claim: %s", logVolumePrefix, key)
	s.lock.Lock()
	s.claims[key] = *vc
	s.lock.Unlock()
}

func (s *VolumesState) DelClaim(key string) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s del claim: %#v", logVolumePrefix, key)
	s.lock.Lock()
	if _, ok := s.claims[key]; ok {
		delete(s.claims, key)
	}
	s.lock.Unlock()
}

func (s *VolumesState) SetLocal(key string) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s set volume: %s as local", logVolumePrefix, key)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.local[key] = true
}

func (s *VolumesState) DelLocal(key string) {
	log := logger.WithContext(context.Background())
	log.Debugf("%s del volume: %s from local", logVolumePrefix, key)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.local[key] = true
}

func (s *VolumesState) IsLocal(key string) bool {
	log := logger.WithContext(context.Background())
	log.Debugf("%s check volume: %s is local", logVolumePrefix, key)
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.local[key]; ok {
		log.Debugf("%s volume: %s is local", logVolumePrefix, key)
		return true
	}

	log.Debugf("%s volume: %s is not local", logVolumePrefix, key)
	return false
}
