package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"promagentd/configs"
	"strings"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

type Client struct {
	serverConfig configs.ServerConfig
	clientConfig configs.ClientConfig

	httpClient *http.Client
	req        *req.Req
}

func NewClient(conf configs.AgentConfig) *Client {
	return &Client{
		serverConfig: conf.Server,
		clientConfig: conf.Client,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		req: req.New(),
	}
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s/%s/", strings.Trim(c.serverConfig.Addr, "/"), strings.Trim(path, "/"))
}

func (c *Client) header() req.Header {
	return req.Header{
		"x-token": c.serverConfig.Token,
	}
}

func (c *Client) result(resp *req.Resp) string {
	if txt, err := resp.ToString(); err == nil {
		return fmt.Sprint(err)
	} else {
		return txt
	}
}

func (c *Client) Register(params map[string]interface{}) error {
	response, err := c.req.Post(c.url("/agent/register/"), c.header(), req.BodyJSON(params), c.httpClient)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"params": params,
			"err":    err,
		}).Error("request register api")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"params":   params,
		"response": c.result(response),
	}).Debug("request register api")
	return nil
}

func (c *Client) Heartbeat(params map[string]interface{}) error {
	response, err := c.req.Post(c.url("/agent/heartbeat/"), c.header(), req.BodyJSON(params), c.httpClient)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"params": params,
			"err":    err,
		}).Error("request heartbeat api")
		return err
	}
	logrus.WithFields(logrus.Fields{
		"params":   params,
		"response": c.result(response),
	}).Debug("request heartbeat api")
	return nil
}

func (c *Client) Config(params map[string]interface{}) (string, error) {
	response, err := c.req.Get(c.url("/agent/config/"), c.header(), req.QueryParam(params), c.httpClient)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"params": params,
			"err":    err,
		}).Debug("request config api")
		return "", err
	}
	logrus.WithFields(logrus.Fields{
		"params":   params,
		"response": c.result(response),
	}).Debug("request config api")
	return response.ToString()
}
