package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
}

type Unicorn struct {
	Id     string `json:"_id,omitempty"`
	Name   string `json:"name"`
	Age    int   `json:"age"`
	Colour string `json:"colour"`
}

func NewClient(hostname string, authToken string) *Client {
	return &Client{
		hostname:   hostname,
		authToken:  authToken,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAll() (*[]Unicorn, error) {
	body, err := c.httpRequest("unicorns", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	var items []Unicorn
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) GetItem(id string) (*Unicorn, error) {
	body, err := c.httpRequest(fmt.Sprintf("unicorns/%v", id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	item := &Unicorn{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) NewItem(item *Unicorn) (error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest("unicorns", "POST", buf)
	if err != nil {
		return err
	}
	
	return err
}


func (c *Client) UpdateItem(id string, item *Unicorn) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("unicorns/%s", id), "PUT", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteItem(itemName string) error {
	_, err := c.httpRequest(fmt.Sprintf("unicorns/%s", itemName), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK || resp.StatusCode != http.StatusCreated {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 || 201 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 || 201 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s/%s/%s", c.hostname, c.authToken, path)
}
