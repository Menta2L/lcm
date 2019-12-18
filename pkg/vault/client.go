package vault

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/menta2l/lcm/pkg/config"
)

type Client struct {
	Config config.VaultConfig
	Token  string
	Vc     *api.Client
}

func NewClient(cfg config.VaultConfig) (*Client, error) {
	c := &Client{
		Config: cfg,
	}
	err := c.auth()
	if err != nil {
		return c, err
	}
	vc, err := api.NewClient(
		&api.Config{
			Address:    c.Config.Address,
			HttpClient: &http.Client{},
			Timeout:    3 * time.Second,
		},
	)
	if err != nil {
		return c, err
	}
	vc.SetToken(c.Token)
	vc.Auth()
	c.Vc = vc
	return c, nil

}
func (c *Client) List(path string) error {

	path = ensureTrailingSlash(sanitizePath(path))
	mountPath, v2, err := IsKVv2(path, c.Vc)
	if err != nil {
		return err
	}
	if v2 {
		path = AddPrefixToVKVPath(path, mountPath, "metadata")
		if err != nil {
			return err
		}
	}
	secret, err := c.Vc.Logical().List(path)
	if err != nil {
		return err
	}
	_, ok := extractListData(secret)
	if secret == nil || secret.Data == nil {
		return fmt.Errorf("No value found at %s", path)
	}
	if !ok {
		fmt.Errorf("No entries found at %s", path)
	}
	fmt.Printf("%v", secret.Data)
	var items []string
	for _, path := range secret.Data {
		// expecting "[secret0 secret1 secret2...]"

		//
		// if the name both exists as directory and as file
		// e.g. "/secret/" and "/secret" it will print an empty line
		items = strings.Split(strings.Trim(fmt.Sprint(path), "[]"), " ")
	}
	fmt.Printf("Items : %v \n", items)

	return nil
}
func (c *Client) GetKeys(path string) ([]string, error) {
	var items []string
	secrets, err := c.GetSecrets(path)
	if err != nil {
		return items, err
	}
	for _, path := range secrets.Data {
		items = strings.Split(strings.Trim(fmt.Sprint(path), "[]"), " ")
	}

	return items, nil
}
func (c *Client) GetSecrets(path string) (*api.Secret, error) {
	path = ensureTrailingSlash(sanitizePath(path))
	mountPath, v2, err := IsKVv2(path, c.Vc)
	if err != nil {
		return nil, err
	}
	if v2 {
		path = AddPrefixToVKVPath(path, mountPath, "metadata")
		if err != nil {
			return nil, err
		}
	}
	secret, err := c.Vc.Logical().List(path)
	if err != nil {
		return nil, err
	}
	_, ok := extractListData(secret)
	if secret == nil || secret.Data == nil {
		return nil, fmt.Errorf("No value found at %s", path)
	}
	if !ok {
		return nil, fmt.Errorf("No entries found at %s", path)
	}
	return secret, nil

}
func (c *Client) Exist(path string) bool {
	_, err := c.Read(path)
	if err != nil {
		return false
	}
	return true
}
func (c *Client) Write(path string, data map[string]interface{}) error {
	mountPath, v2, err := IsKVv2(path, c.Vc)
	if err != nil {
		return err
	}
	if v2 {
		path = AddPrefixToVKVPath(path, mountPath, "data")
		data = map[string]interface{}{
			"data":    data,
			"options": map[string]interface{}{},
		}

	}

	_, err = c.Vc.Logical().Write(path, data)
	if err != nil {
		return err
	}

	return nil
}
func (c *Client) Read(path string) (*api.Secret, error) {
	path = sanitizePath(path)
	mountPath, v2, err := IsKVv2(path, c.Vc)
	if err != nil {
		return nil, err
	}

	var versionParam map[string]string

	if v2 {
		path = AddPrefixToVKVPath(path, mountPath, "data")
	}
	secret, err := c.kvReadRequest(path, versionParam)
	if err != nil {
		return nil, fmt.Errorf("Error reading %s: %s", path, err)
	}
	if secret == nil {
		return nil, fmt.Errorf("No value found at %s", path)
	}
	return secret, err
}
func (c *Client) kvReadRequest(path string, params map[string]string) (*api.Secret, error) {
	r := c.Vc.NewRequest("GET", "/v1/"+path)
	for k, v := range params {
		r.Params.Set(k, v)
	}
	resp, err := c.Vc.RawRequest(r)
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp != nil && resp.StatusCode == 404 {
		secret, parseErr := api.ParseSecret(resp.Body)
		switch parseErr {
		case nil:
		case io.EOF:
			return nil, nil
		default:
			return nil, err
		}
		if secret != nil && (len(secret.Warnings) > 0 || len(secret.Data) > 0) {
			return secret, nil
		}
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return api.ParseSecret(resp.Body)
}
func (c *Client) auth() error {
	c.Token = c.Config.Token
	/*
		body := []byte(fmt.Sprintf(`{"password":"%s"}`, c.Config.Password))
		url := fmt.Sprintf("%s/v1/auth/userpass/login/%s", c.Config.Address, c.Config.Username)
		fmt.Println(string(body))
		fmt.Println(url)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			return err
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("Unable to retrieve authentication token from vault %q", err)
		} else if res.StatusCode != 200 {
			return fmt.Errorf("Unable to retrieve authentication token from vault (status code %d)", res.StatusCode)
		}
		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Unable to parse request body: %q", err)
		}
		var secret vault.Secret
		err = json.Unmarshal(body, &secret)
		if err != nil {
			return fmt.Errorf("Unable to parse request body: %q", err)
		}
		c.Token = secret.Auth.ClientToken
	*/
	return nil

}

// Returns the the time when a token will expire
func (c *Client) GetTokenTTL(token string) (time.Time, error) {

	var valid_until time.Time

	// Don't login, just show information about the current token.
	secret, err := c.Vc.Auth().Token().Lookup(c.Token)
	if err != nil {
		return valid_until, err
	}

	ttl, err := secret.Data["ttl"].(json.Number).Int64()
	if err != nil {
		return valid_until, err
	}

	return time.Unix(time.Now().Unix()+ttl, 0), nil
}
