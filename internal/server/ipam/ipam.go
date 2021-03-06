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

package ipam

import (
	"github.com/lastbackend/lastbackend/internal/pkg/storage"
	"net"

	"github.com/lastbackend/lastbackend/internal/server/ipam/local"
)

type IPAM interface {
	Lease() (*net.IP, error)
	Release(ip *net.IP) error
	Available() int
	Reserved() int
}

func New(stg storage.IStorage, cidr string) (IPAM, error) {
	return local.New(stg, cidr)
}
