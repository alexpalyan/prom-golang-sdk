// groups
package prom

import (
	"strconv"
)

type GroupsRequest struct {
	Limit  int
	LastId int
}

type Group struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Image         string `json:"image"`
	ParentGroupId int    `json:"parent_group_id"`
}

type GroupsResponse struct {
	Groups []Group `json:"groups"`
	Error  string  `json:"string"`
}

func (c *Client) GetGroups(request GroupsRequest) (groups []Group, err error) {
	var (
		result GroupsResponse
		params map[string]string = make(map[string]string)
	)

	if request.LastId > 0 {
		params["last_id"] = strconv.Itoa(request.LastId)
	}

	if request.Limit > 0 {
		params["limit"] = strconv.Itoa(request.LastId)
	}

	err = c.Get("/groups/list", params, &result)
	groups = result.Groups
	return
}
