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

package dir

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/lastbackend/lastbackend/internal/pkg/models"
	"github.com/lastbackend/lastbackend/tools/logger"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Storage struct {
	root string
}

type Config struct {
	RootDir string
}

func (s *Storage) List(ctx context.Context) (map[string]*models.VolumeState, error) {
	var vols = make(map[string]*models.VolumeState, 0)

	var dirs []string
	f, err := os.Open(s.root)
	if err != nil {
		return vols, err
	}

	items, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return vols, err
	}

	for _, item := range items {
		if item.IsDir() {
			dirs = append(dirs, item.Name())
		}
	}

	for _, dir := range dirs {
		vol := new(models.VolumeState)

		vol.Path = filepath.Join(s.root, dir)
		vol.Type = models.KindVolumeHostDir
		vol.Ready = true
		vols[dir] = vol
	}

	return vols, nil
}

func (s *Storage) Create(ctx context.Context, name string, manifest *models.VolumeManifest) (*models.VolumeState, error) {

	var (
		status = new(models.VolumeState)
		path   = filepath.Join(s.root, strings.Replace(name, ":", "_", -1))
	)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return status, err
		}
	}

	status.Path = path
	status.Ready = true

	return status, nil
}

func (s *Storage) FilesList(ctx context.Context, state *models.VolumeState) ([]string, error) {

	return make([]string, 0), nil
}

func (s *Storage) FilesPut(ctx context.Context, state *models.VolumeState, files map[string]string) error {
	log := logger.WithContext(context.Background())
	for file, data := range files {
		path := filepath.Join(state.Path, file)
		var f *os.File

		f, err := os.OpenFile(path, os.O_WRONLY, 0644)
		if err != nil {
			if os.IsNotExist(err) {
				if f, err = os.Create(path); err != nil {
					return err
				}
			} else {
				return err
			}
		}
		f.Close()

		if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
			log.Errorf("can not be write data to file: %s", err.Error())
		}

	}

	return nil
}

func (s *Storage) FilesCheck(ctx context.Context, state *models.VolumeState, files map[string]string) (bool, error) {

	for file, data := range files {
		path := filepath.Join(state.Path, file)
		var f *os.File

		f, err := os.Open(path)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			} else {
				return false, err
			}
		}

		hashFile := sha1.New()
		hashData := sha1.New()
		if _, err := io.Copy(hashFile, f); err != nil {
			return false, err
		}
		if _, err := io.Copy(hashData, bytes.NewReader([]byte(data))); err != nil {
			return false, err
		}

		//Convert the bytes to a string
		hashFileS := hex.EncodeToString(hashFile.Sum(nil)[:20])
		hashFileD := hex.EncodeToString(hashData.Sum(nil)[:20])

		if hashFileS != hashFileD {
			return false, nil
		}

		f.Close()
	}

	return true, nil
}

func (s *Storage) FilesDel(ctx context.Context, state *models.VolumeState, files []string) error {

	for _, file := range files {
		path := filepath.Join(state.Path, file)
		if err := os.Remove(path); err != nil {
			if os.IsNotExist(err) {
				continue
			}
		}
	}

	return nil
}

func (s *Storage) Remove(ctx context.Context, state *models.VolumeState) error {

	if err := os.RemoveAll(filepath.Join(state.Path)); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return nil
}

func Get(cfg Config) (*Storage, error) {
	log := logger.WithContext(context.Background())
	log.Debug("Initialize dir storage interface")
	var s = new(Storage)

	if cfg.RootDir != "" {
		s.root = cfg.RootDir
		log.Debugf("Initialize dir storage interface root: %s", s.root)
	}

	if _, err := os.Stat(s.root); os.IsNotExist(err) {
		err = os.MkdirAll(s.root, 0755)
		if err != nil {
			return nil, err
		}
	}

	return s, nil
}
