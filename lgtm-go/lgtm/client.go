package lgtm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	pathLogin   = "%s/login?access_token=%s"
	pathUser    = "%s/api/user"
	pathRepos   = "%s/api/user/repos"
	pathRepo    = "%s/api/repos/%s"
	pathConf    = "%s/api/repos/%s/maintainers"
	pathOrgConf = "%s/api/repos/%s/maintainers/%s"
)

type client struct {
	client *http.Client
	base   string // base url
}

// NewClient returns a client at the specified url.
func NewClient(uri string) Client {
	return &client{http.DefaultClient, uri}
}

// NewClientToken returns a client at the specified url that
// authenticates all outbound requests with the given token.
func NewClientToken(uri, token string) Client {
	config := new(oauth2.Config)
	auther := config.Client(oauth2.NoContext, &oauth2.Token{AccessToken: token})
	return &client{auther, uri}
}

// SetClient sets the default http client. This should be
// used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the server.
func (c *client) SetClient(client *http.Client) {
	c.client = client
}

// Token returns a repository MAINTAINERS configuration.
func (c *client) Token(token string) (*Token, error) {
	out := new(Token)
	uri := fmt.Sprintf(pathLogin, c.base, token)
	err := c.post(uri, nil, out)
	return out, err
}

// User returns the currently authenticated user.
func (c *client) User() (*User, error) {
	out := new(User)
	uri := fmt.Sprintf(pathUser, c.base)
	err := c.get(uri, out)
	return out, err
}

// Repo returns a repository by name.
func (c *client) Repo(repo string) (*Repo, error) {
	out := new(Repo)
	uri := fmt.Sprintf(pathRepo, c.base, repo)
	err := c.get(uri, out)
	return out, err
}

// Repos returns a list of all repositories to which
// the user has explicit access in the host system.
func (c *client) Repos() ([]*Repo, error) {
	out := make([]*Repo, 0)
	uri := fmt.Sprintf(pathRepos, c.base)
	err := c.get(uri, &out)
	return out, err
}

// Activate activates a repository.
func (c *client) Activate(repo string) (*Repo, error) {
	out := new(Repo)
	uri := fmt.Sprintf(pathRepo, c.base, repo)
	err := c.post(uri, nil, out)
	return out, err
}

// Deactivate de-activates a repository.
func (c *client) Deactivate(repo string) error {
	uri := fmt.Sprintf(pathRepo, c.base, repo)
	err := c.delete(uri)
	return err
}

// Maintainer returns a repository MAINTAINERS configuration.
func (c *client) Maintainer(repo string) (*Maintainer, error) {
	out := new(Maintainer)
	uri := fmt.Sprintf(pathConf, c.base, repo)
	err := c.get(uri, out)
	return out, err
}

// MaintainerOrg returns a repository MAINTAINERS configuration file
// for a subset of people in the sepcified organization.
func (c *client) MaintainerOrg(repo, org string) (*Maintainer, error) {
	out := new(Maintainer)
	uri := fmt.Sprintf(pathOrgConf, c.base, repo, org)
	err := c.get(uri, out)
	return out, err
}

//
// http request helper functions
//

// helper function for making an http GET request.
func (c *client) get(rawurl string, out interface{}) error {
	return c.do(rawurl, "GET", nil, out)
}

// helper function for making an http POST request.
func (c *client) post(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "POST", in, out)
}

// helper function for making an http PUT request.
func (c *client) put(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PUT", in, out)
}

// helper function for making an http PATCH request.
func (c *client) patch(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PATCH", in, out)
}

// helper function for making an http DELETE request.
func (c *client) delete(rawurl string) error {
	return c.do(rawurl, "DELETE", nil, nil)
}

// helper function to make an http request
func (c *client) do(rawurl, method string, in, out interface{}) error {
	// executes the http request and returns the body as
	// and io.ReadCloser
	body, err := c.stream(rawurl, method, in, out)
	if err != nil {
		return err
	}
	defer body.Close()

	// if a json response is expected, parse and return
	// the json response.
	if out != nil {
		return json.NewDecoder(body).Decode(out)
	}
	return nil
}

// helper function to stream an http request
func (c *client) stream(rawurl, method string, in, out interface{}) (io.ReadCloser, error) {
	uri, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	// if we are posting or putting data, we need to
	// write it to the body of the request.
	var buf io.ReadWriter
	if in != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(in)
		if err != nil {
			return nil, err
		}
	}

	// creates a new http request to bitbucket.
	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > http.StatusPartialContent {
		defer resp.Body.Close()
		out, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf(string(out))
	}
	return resp.Body, nil
}
