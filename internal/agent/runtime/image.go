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
	"encoding/base64"
	"fmt"
	"github.com/lastbackend/lastbackend/tools/logger"
	"strings"

	"github.com/lastbackend/lastbackend/internal/pkg/errors"
	"github.com/lastbackend/lastbackend/internal/pkg/models"
		"golang.org/x/net/context"
)

func (r Runtime) ImagePull(ctx context.Context, namespace string, image *models.SpecTemplateContainerImage) error {
	log := logger.WithContext(context.Background())
	var (
		mf = new(models.ImageManifest)
	)

	mf.Name = image.Name
	if len(image.Sha) != 0 {
		mf.Name = fmt.Sprintf("%s@%s", strings.Split(image.Name, ":")[0], image.Sha)
	}

	if image.Secret.Name != models.EmptyString {

		secret, err := r.SecretGet(ctx, namespace, image.Secret.Name)
		if err != nil {
			log.Errorf("can not be get secret for image. err: %s", err.Error())
			if strings.Contains(err.Error(), "Internal Server Error") {
				return errors.New("can not be get secret data")
			}
			return err
		}

		token, err := secret.DecodeSecretTextData(image.Secret.Key)
		if err != nil {
			log.Errorf("can not be get parse secret text data. err: %s", err.Error())
			return err
		}

		payload, _ := base64.StdEncoding.DecodeString(token)
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 {
			log.Error("can not be parse docker auth secret: invalid format")
			return errors.New("docker auth secret format is invalid")
		}

		data := models.SecretAuthData{
			Username: pair[0],
			Password: pair[1],
		}

		auth, err := r.cii.Auth(ctx, &data)
		if err != nil {
			log.Errorf("can not be create secret string. err: %s", err.Error())
			return err
		}

		mf.Auth = auth
	}

	img, err := r.cii.Pull(ctx, mf, nil)
	if err != nil {
		log.Errorf("can not be pull image: %s", err.Error())
		return err
	}

	if img != nil {
		r.state.Images().AddImage(img.SelfLink(), img)
	}

	return nil
}

func (r Runtime) ImageRemove(ctx context.Context, link string) error {
	log := logger.WithContext(context.Background())
	if err := r.cii.Remove(ctx, link); err != nil {
		log.Warnf("Can-not remove unnecessary image %s: %s", link, err)
	}

	r.state.Images().DelImage(link)

	return nil
}

func (r Runtime) ImageRestore(ctx context.Context) error {

	state := r.state.Images()
	imgs, err := r.cii.List(ctx)
	if err != nil {
		return err
	}

	for _, i := range imgs {
		if len(i.Meta.Tags) > 0 {
			state.AddImage(i.Meta.Name, i)
		}
	}

	return nil
}
