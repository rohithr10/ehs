package dto

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpRequestDetails struct {
	QueryParams map[string]string `json:"query_params"`
	Headers map[string]string `json:"headers"`
	Body []byte `json:"body"`

}

func NewHttpRequest(httpClient *http.Client, method, url string, details HttpRequestDetails) ([]byte, error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(details.Body))
	req.Header.Add("Accept-Encoding", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if details.Headers != nil {
		for k, v := range details.Headers {
			req.Header.Add(k, v)
		}
	}

	if details.QueryParams != nil {
		q := req.URL.Query()
		for k, v := range details.QueryParams {
			q.Add(k, v)
			q.Add(k, v)
			req.URL.RawQuery = q.Encode()
		}
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		//logger.Error("http request failed with error : " + err.Error())
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	//logger.Info("Response Body: " + string(data))
	if err != nil {
		//logger.Error("reading response body failed with error: " + err.Error())
		return nil, err

	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode != http.StatusCreated {
			//logger.Error("request failed with status: " + resp.Status + " having response body:" + string(data))
			return nil, err
		}
	}

	//logger.Info("ending new http request from util")
	return data, nil
}