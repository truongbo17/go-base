package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type BaseRequest struct {
	Client *http.Client
}

func NewBaseRequest(timeout time.Duration) *BaseRequest {
	return &BaseRequest{
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (r *BaseRequest) Do(ctx context.Context, method, url string, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	var requestBody []byte
	var err error
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}

	return resp, respBody, nil
}

func (r *BaseRequest) Get(ctx context.Context, url string, headers map[string]string) (*http.Response, []byte, error) {
	return r.Do(ctx, http.MethodGet, url, headers, nil)
}

func (r *BaseRequest) Post(ctx context.Context, url string, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	return r.Do(ctx, http.MethodPost, url, headers, body)
}

func (r *BaseRequest) Put(ctx context.Context, url string, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	return r.Do(ctx, http.MethodPut, url, headers, body)
}

func (r *BaseRequest) Delete(ctx context.Context, url string, headers map[string]string, body interface{}) (*http.Response, []byte, error) {
	return r.Do(ctx, http.MethodDelete, url, headers, body)
}
