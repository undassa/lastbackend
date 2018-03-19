//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
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

package http

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	ctx context.Context
}

func (r *Request) Get(endpoint string) ([]byte, error) {

	schema := "http"
	if r.ctx.Value("https").(bool) {
		schema = "https"
	}

	uri := fmt.Sprintf("%s://%s/%s", schema, r.ctx.Value("host").(string), endpoint)

	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NewRequest(ctx context.Context) *Request {
	return &Request{ctx}
}