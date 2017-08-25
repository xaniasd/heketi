//
// Copyright (c) 2015 The heketi Authors
//
// This file is licensed to you under your choice of the GNU Lesser
// General Public License, version 3 or any later version (LGPLv3 or
// later), as published by the Free Software Foundation,
// or under the Apache License, Version 2.0 <LICENSE-APACHE2 or
// http://www.apache.org/licenses/LICENSE-2.0>.
//
// You may not use this file except in compliance with those terms.
//

package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/heketi/heketi/pkg/glusterfs/api"
	"github.com/heketi/utils"
)

func (c *Client) GeoReplicationCreate(id string, request *api.GeoReplicationRequest) error {
	// Marshal request to JSON
	buffer, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// Create a request
	req, err := http.NewRequest("POST", c.host+"/volumes/"+id+"/georeplication", bytes.NewBuffer(buffer))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Set token
	err = c.setToken(req)
	if err != nil {
		return err
	}

	// Send request
	r, err := c.do(req)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusAccepted {
		return utils.GetErrorFromResponse(r)
	}

	// Wait for response
	r, err = c.waitForResponseWithTimer(r, time.Second)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return utils.GetErrorFromResponse(r)
	}

	// Read JSON response
	var resp api.GeoReplicationCreateResponse
	err = utils.GetJsonFromResponse(r, &resp)
	r.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
