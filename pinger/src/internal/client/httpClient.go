package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Eugene-Usachev/test-task-for-VK/pinger/src/pkg"
	"github.com/Eugene-Usachev/test-task-for-VK/pinger/src/pkg/model"
	"github.com/goccy/go-json"
)

var (
	ErrServerInternalError = fmt.Errorf("server internal error")
	ErrBadRequest          = fmt.Errorf("bad request")
)

type HTTPClient struct {
	client *http.Client
	addr   string
}

var _ Client = (*HTTPClient)(nil)

func NewHTTPClient(addr string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{Timeout: 3 * time.Second},
		addr:   addr,
	}
}

func (c *HTTPClient) GetContainers(ctx context.Context) ([]model.GetContainer, error) {
	url := "http://" + c.addr + "/container/id_and_ip_address_only"

	var (
		err        error
		req        *http.Request
		resp       *http.Response
		containers []model.GetContainer
	)

	req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return []model.GetContainer{}, err
	}

	err = pkg.DoWithTries(func() error {
		// reason = false positive
		//nolint:bodyclose
		resp, err = c.client.Do(req)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				log.Println("Failed to close HTTP client body: ", err)
			}
		}(resp.Body)

		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)

			return fmt.Errorf("%w: %s", ErrServerInternalError, string(body))
		}

		return json.NewDecoder(resp.Body).Decode(&containers)
	}, 5, 100*time.Millisecond)

	return containers, err
}

func (c *HTTPClient) RegisterContainers(ctx context.Context, addrs []string) error {
	url := "http://" + c.addr + "/container/many"

	var (
		req     *http.Request
		reqBody []byte
		resp    *http.Response
	)

	containers := make([]model.RegisterContainer, len(addrs))
	for i, addr := range addrs {
		containers[i] = model.RegisterContainer{
			IpAddress: addr,
		}
	}

	bodyBytes, err := json.Marshal(containers)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(bodyBytes)

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return err
	}

	err = pkg.DoWithTries(func() error {
		// reason = false positive
		//nolint:bodyclose
		resp, err = c.client.Do(req)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				log.Println("Failed to close HTTP client body: ", err)
			}
		}(resp.Body)

		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			reqBody, _ = io.ReadAll(resp.Body)

			if resp.StatusCode == http.StatusBadRequest {
				return fmt.Errorf("%w: %s", ErrBadRequest, string(reqBody))
			}

			return fmt.Errorf("%w: %s", ErrServerInternalError, string(reqBody))
		}

		return nil
	}, 5, 100*time.Millisecond)

	return err
}

func (c *HTTPClient) StorePings(ctx context.Context, pings []model.Ping) error {
	url := "http://" + c.addr + "/ping/"

	var (
		err  error
		req  *http.Request
		resp *http.Response
	)

	data, err := json.Marshal(pings)
	if err != nil {
		return err
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	err = pkg.DoWithTries(func() error {
		// reason = false positive
		//nolint:bodyclose
		resp, err = c.client.Do(req)

		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				log.Println("Failed to close HTTP client body: ", err)
			}
		}(resp.Body)

		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode == http.StatusBadRequest {
				return fmt.Errorf("%w: %s", ErrBadRequest, string(body))
			}

			return fmt.Errorf("%w: %s", ErrServerInternalError, string(body))
		}

		return nil
	}, 5, 100*time.Millisecond)

	return err
}
