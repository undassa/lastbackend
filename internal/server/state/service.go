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
	"sync"

	"github.com/lastbackend/lastbackend/internal/pkg/errors"
	"github.com/lastbackend/lastbackend/internal/pkg/models"
	"github.com/lastbackend/lastbackend/tools/logger"
)

// ServiceController structure
type ServiceController struct {
	lock  sync.Mutex
	items map[models.SelfLink]*models.Service
}

// Run controller state
func (sc *ServiceController) Run(ctx context.Context) error {
	return nil
}

// Restore controller state
func (sc *ServiceController) Restore(ctx context.Context) error {
	return nil
}

// List all namespaces in state
func (sc *ServiceController) List(ctx context.Context, filter *ResourceFilter) (models.NamespaceResourceList, error) {

	log := logger.WithContext(ctx)
	log.Debugf("%s:list:> get service list", logPrefix)

	list := models.NewServiceList()

	for _, item := range sc.items {
		list.Items = append(list.Items, item)
	}

	return list, nil
}

// Map all service in state
func (sc *ServiceController) Map(ctx context.Context, filter *ResourceFilter) map[models.SelfLink]*models.Service {

	log := logger.WithContext(ctx)
	log.Debugf("%s:service:> get service map", logPrefix)

	return sc.items
}

// Set service to state
func (sc *ServiceController) Put(ctx context.Context, mf models.NamespaceResourceManifest) (models.NamespaceResource, error) {
	log := logger.WithContext(ctx)
	log.Debugf("%s:set:> set service", logPrefix)

	// TODO: fill service manifest set logic

	return nil, nil
}

// Set service to state
func (sc *ServiceController) Set(ctx context.Context, mf models.NamespaceResourceManifest) (models.NamespaceResource, error) {
	log := logger.WithContext(ctx)
	log.Debugf("%s:set:> set service", logPrefix)

	// TODO: fill service manifest set logic

	return nil, nil
}

// Get particular service from state
func (sc *ServiceController) Get(ctx context.Context, selflink models.SelfLink) (models.NamespaceResource, error) {
	log := logger.WithContext(ctx)
	log.Debugf("%s:list:> get service from state", logPrefix)

	sc.lock.Lock()
	item, ok := sc.items[selflink]
	sc.lock.Unlock()

	if !ok {
		return nil, errors.NewResourceNotFound()
	}

	return item, nil
}

// Del service from state
func (sc *ServiceController) Del(ctx context.Context, selflink models.SelfLink) (models.NamespaceResource, error) {
	log := logger.WithContext(ctx)
	log.Debugf("%s:list:> delete service from state", logPrefix)

	sc.lock.Lock()
	_, ok := sc.items[selflink]
	sc.lock.Unlock()

	if !ok {
		return nil, errors.NewResourceNotFound()
	}

	// TODO: implement service delete logic

	return nil, nil
}

func (sc *ServiceController) loop() {

}

// ServiceControllerOpts struct for filtering queries to state
type ServiceControllerOpts struct {
	namespace *models.NamespaceSelfLink
}

func (sco *ServiceControllerOpts) WithNamespace(ns string) {

}

func (sco *ServiceControllerOpts) WithNamespaceSelfLink(ns models.ServiceSelfLink) {

}

// NewServiceController return new instance of ServiceController struct
func NewServiceController(ctx context.Context) *ServiceController {
	return nil
}

// NewServiceControllerOptions return new option struct for quering state
func NewServiceControllerOptions() *ServiceControllerOpts {
	return new(ServiceControllerOpts)
}
