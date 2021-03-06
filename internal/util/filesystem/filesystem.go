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

package filesystem

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	LineSeparator      = '\n'
	DefaultDataDir     = "/var/lib/lastbackend/"
	DefaultHomeDataDir = "${HOME}/.lastbackend/"
)

var (
	homes = []string{"$HOME", "${HOME}", "~"}
)

func LineSeek(lines int, f *os.File) (int64, error) {

	count := 0
	pos, err := f.Seek(0, io.SeekEnd)

	if err != nil {
		return 0, err
	}

	chunk := 4096

	b1 := make([]byte, 1)
	if _, err := f.ReadAt(b1, pos-1); err != nil {
		return 0, err
	}

	if '\n' == b1[0] {
		pos = pos - 1
	}

	for {
		rf := pos - int64(chunk)
		ids := make([]int64, 0)

		if rf <= 0 {
			chunk += int(rf)
			rf = 0
			ids = append(ids, 0)
			count++
		}

		b := make([]byte, chunk)

		_, err := f.ReadAt(b, rf)

		if err != nil {
			return 0, err
		}

		i := 0

		for {
			pos := bytes.IndexByte(b[i:], LineSeparator)
			if pos == -1 {
				break
			}
			i = i + pos + 1
			ids = append(ids, int64(i))
			count++
		}

		var lpos int64

		if len(ids) == 0 {
			lpos = 0
		} else {
			lpos = ids[0]
		}

		if count == lines {
			return lpos + rf, nil
		}

		if count > lines {
			left := count - lines
			pos := ids[left]
			return pos + rf, nil

		}

		if rf == 0 {
			return 0, nil
		}
		pos = lpos + rf - 1
	}

}

// MkDir is used to create directory
func MkDir(path string, mode os.FileMode) (err error) {
	err = os.MkdirAll(path, mode)
	return err
}

// CreateFile is used to create file
func CreateFile(path string, mode os.FileMode) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	err = os.Chmod(path, mode)
	if err != nil {
		return err
	}
	return nil
}

func WriteStrToFile(filepath, value string, mode os.FileMode) error {
	dirPath := path.Dir(filepath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, mode); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if err := CreateFile(filepath, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, []byte(value), mode)
}

// HomeDir returns the home directory for the current user
func HomeDir() string {
	if runtime.GOOS == "windows" {
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDir := homeDrive + homePath
			if _, err := os.Stat(homeDir); err == nil {
				return homeDir
			}
		}
		if userProfile := os.Getenv("USERPROFILE"); len(userProfile) > 0 {
			if _, err := os.Stat(userProfile); err == nil {
				return userProfile
			}
		}
	}
	return os.Getenv("HOME")
}

func WriteFile(name string, content string) error {
	os.MkdirAll(filepath.Dir(name), 0755)
	err := ioutil.WriteFile(name, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("%v: writing %s", err, name)
	}
	return nil
}

func ResolveHomePath(s string) (string, error) {
	for _, home := range homes {
		if strings.Contains(s, home) {
			homeDir := HomeDir()
			s = strings.Replace(s, home, homeDir, -1)
		}
	}
	return s, nil
}
func DetermineHomePath(dirPath string, forceLocal bool) (string, error) {
	if dirPath == "" {
		if os.Getuid() == 0 && !forceLocal {
			dirPath = DefaultDataDir
		} else {
			dirPath = DefaultHomeDataDir
		}
	}
	dirPath, err := ResolveHomePath(dirPath)
	if err != nil {
		return "", fmt.Errorf("%v: resolving %s", err, dirPath)
	}
	return filepath.Abs(dirPath)
}
