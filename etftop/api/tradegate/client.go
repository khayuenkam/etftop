package tradegate

import (
	"encoding/json"
	"fmt"
	"github.com/khayuenkam/etftop/etftop/api/tradegate/types"
	"io/ioutil"
	"net/http"
	"net/url"
)

var baseURL = "https://www.tradegate.de"

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient: httpClient}
}

func doReq(req *http.Request, client *http.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func (client *Client) MakeReq(url, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := doReq(req, client.httpClient)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (client *Client) Refresh(isin string) (*types.RefreshType, error) {
	params := url.Values{}
	params.Add("isin", isin)

	url := fmt.Sprintf("%s/refresh.php?%s", baseURL, params.Encode())
	resp, err := client.MakeReq(url, "POST")
	if err != nil {
		return nil, err
	}

	var data *types.RefreshType
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
