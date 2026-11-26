package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/dweatherstone/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	baseURL    string
	cache      *pokecache.Cache
}

func NewClient() *Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.ResponseHeaderTimeout = 10 * time.Second
	transport.MaxIdleConns = 100
	transport.MaxConnsPerHost = 20

	httpClient := http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
	client := &Client{
		httpClient: httpClient,
		baseURL:    "https://pokeapi.co/api/v2/",
		cache:      pokecache.NewCache(5 * time.Minute),
	}
	return client
}

func (c *Client) Get(url string, dest any) error {
	data, found := c.cache.Get(url)
	if found {
		if err := unmarshalJson(data, dest); err != nil {
			return err
		}
		return nil
	}
	res, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("sending HTTP GET request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading HTTP GET response: %w", err)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("response failed with status code %d and body: %s", res.StatusCode, body)
	}

	if err := unmarshalJson(body, dest); err != nil {
		return err
	}
	c.cache.Add(url, body)

	return nil
}

func unmarshalJson(body []byte, dest any) error {
	if err := json.Unmarshal(body, dest); err != nil {
		return fmt.Errorf("unable to unmarshal response: %w", err)
	}

	return nil
}
