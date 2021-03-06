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

package oci

import (
	"context"
	"io"
	"time"

	"github.com/lastbackend/lastbackend/internal/pkg/models"
)

type Runc struct {
}

func NewRunc(cfg ConfigRunc) (*Runc, error) {
	r := new(Runc)
	return r, nil
}

func (Runc) List(ctx context.Context, all bool) ([]*models.Container, error) {
	return nil, nil
}

func (Runc) Create(ctx context.Context, spec *models.ContainerManifest) (string, error) {
	return "", nil
}

func (Runc) Start(ctx context.Context, ID string) error {
	return nil
}

func (Runc) Restart(ctx context.Context, ID string, timeout *time.Duration) error {
	return nil
}

func (Runc) Stop(ctx context.Context, ID string, timeout *time.Duration) error {
	return nil
}

func (Runc) Pause(ctx context.Context, ID string) error {
	return nil
}

func (Runc) Resume(ctx context.Context, ID string) error {
	return nil
}

func (Runc) Remove(ctx context.Context, ID string, clean bool, force bool) error {
	return nil
}

func (Runc) Inspect(ctx context.Context, ID string) (*models.Container, error) {
	return nil, nil
}

func (Runc) Logs(ctx context.Context, ID string, stdout, stderr, follow bool) (io.ReadCloser, error) {
	return nil, nil
}

func (Runc) Copy(ctx context.Context, ID, path string, content io.Reader) error {
	return nil
}

func (Runc) Wait(ctx context.Context, ID string) error {
	return nil
}

func (Runc) Subscribe(ctx context.Context, container chan *models.Container) error {
	return nil
}

func (Runc) Close() error {
	return nil
}
