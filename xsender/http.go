package xsender

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

func NewHttp(ctx context.Context, cfg *Configs) Send {
	client := resty.New().
		SetTimeout(time.Duration(cfg.TimeoutInSeconds) * time.Second).
		SetRetryCount(cfg.RetryMax).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.StatusCode() >= http.StatusInternalServerError
			},
		)
	return func(request *Request) (*Response, error) {
		res, err := callHttp(client, request)
		if err != nil {
			return nil, err
		}

		ok := res.StatusCode() >= 200 && res.StatusCode() < 300
		if !ok {
			return nil, errors.New(fmt.Sprintf("xsender: %d - %s", res.StatusCode(), res.Status()))
		}

		response := &Response{
			Uri:     res.RawResponse.Request.URL.String(),
			Status:  res.StatusCode(),
			Headers: res.RawResponse.Header,
			Body:    string(res.Body()),
		}
		return response, nil
	}
}

func callHttp(client *resty.Client, request *Request) (*resty.Response, error) {
	req := client.R().
		SetHeaderMultiValues(request.Headers)

	if request.Method == http.MethodPost {
		return req.SetBody(request.Body).Post(request.Uri)
	}
	if request.Method == http.MethodPatch {
		return req.SetBody(request.Body).Patch(request.Uri)
	}
	if request.Method == http.MethodPut {
		return req.SetBody(request.Body).Put(request.Uri)
	}
	return nil, errors.New("xsender.http: only accepted POST/PATCH/PUT method")
}
