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

package v1

import (
	"context"
	"fmt"
	"github.com/lastbackend/lastbackend/internal/pkg/errors"
	"github.com/lastbackend/lastbackend/internal/util/http/request"
	"github.com/lastbackend/lastbackend/pkg/api/types/v1/views"
)

type ControllerClient struct {
	client *request.RESTClient

	hostname string
}

func (ic *ControllerClient) List(ctx context.Context) (*views.ControllerList, error) {

	var i *views.ControllerList
	var e *errors.Http

	err := ic.client.Get(fmt.Sprintf("/controller")).
		AddHeader("Content-Type", "application/json").
		JSON(&i, &e)

	if err != nil {
		return nil, err
	}
	if e != nil {
		return nil, errors.New(e.Message)
	}

	if ic == nil {
		list := make(views.ControllerList, 0)
		i = &list
	}

	return i, nil
}

func (ic *ControllerClient) Get(ctx context.Context) (*views.Controller, error) {

	var s *views.Controller
	var e *errors.Http

	err := ic.client.Get(fmt.Sprintf("/controller/%s", ic.hostname)).
		AddHeader("Content-Type", "application/json").
		JSON(&s, &e)

	if err != nil {
		return nil, err
	}
	if e != nil {
		return nil, errors.New(e.Message)
	}

	return s, nil
}

func newControllerClient(req *request.RESTClient, hostname string) *ControllerClient {
	return &ControllerClient{client: req, hostname: hostname}
}
