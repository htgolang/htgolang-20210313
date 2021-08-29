package client

import (
	"fmt"

	"github.com/imroc/req"
)

type Client struct {
	id       string
	endpoint string
	req      *req.Req
}

func NewClient(id string, endpoint string) *Client {
	return &Client{
		id:       id,
		endpoint: endpoint,
		req:      req.New(),
	}
}

func (c *Client) Register(params map[string]interface{}) error {
	path := fmt.Sprintf("%s/v1/register/", c.endpoint)
	_, err := c.req.Post(path, req.QueryParam{"uuid": c.id}, req.BodyJSON(params))
	return err
}

func (c *Client) Heartbeat(params map[string]interface{}) error {
	path := fmt.Sprintf("%s/v1/heartbeat/", c.endpoint)
	_, err := c.req.Post(path, req.QueryParam{"uuid": c.id}, req.BodyJSON(params))
	return err
}

func (c *Client) Config(version string, r interface{}) error {
	path := fmt.Sprintf("%s/v1/config/", c.endpoint)
	resp, err := c.req.Get(path, req.QueryParam{"uuid": c.id, "version": version})
	if err != nil {
		return err
	}
	return resp.ToJSON(r)
}

func (c *Client) Tasks(r interface{}) error {
	path := fmt.Sprintf("%s/v1/tasks/", c.endpoint)
	resp, err := c.req.Get(path, req.QueryParam{"uuid": c.id})
	if err != nil {
		return err
	}
	return resp.ToJSON(r)
}

func (c *Client) Results(params map[string]interface{}) error {
	path := fmt.Sprintf("%s/v1/task/results/", c.endpoint)
	_, err := c.req.Post(path, req.QueryParam{"uuid": c.id}, req.BodyJSON(params))
	return err
}

func (c *Client) Log(params map[string]interface{}) error {
	path := fmt.Sprintf("%s/v1/log/", c.endpoint)
	_, err := c.req.Post(path, req.QueryParam{"uuid": c.id}, req.BodyJSON(params))
	return err
}
