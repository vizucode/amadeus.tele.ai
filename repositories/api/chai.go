package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type chai struct {
	DEV_KEY    string
	DEV_UID    string
	URI_TARGET string
}

func NewchaiML(dev_key string, dev_uid string, uri_target string) *chai {
	return &chai{
		DEV_KEY:    dev_key,
		DEV_UID:    dev_uid,
		URI_TARGET: uri_target,
	}
}

func (r *chai) GetChat(ctx context.Context, msg string) string {
	body := struct {
		Text               string  `json:"text"`
		Repetition_penalty float32 `json:"repetition_penalty"`
		Temperature        float32 `json:"temperature"`
		Top_P              int     `json:"top_p"`
		Top_K              int     `json:"top_k"`
		Response_length    int     `json:"response_length"`
	}{
		Text:               msg,
		Repetition_penalty: 1.1,
		Temperature:        0.1,
		Top_P:              1,
		Top_K:              1,
		Response_length:    64,
	}

	req, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err.Error())
	}

	// create request
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.URI_TARGET, bytes.NewBuffer(req))
	if err != nil {
		log.Fatal(err.Error())
	}

	httpReq.Header.Add("developer_key", r.DEV_KEY)
	httpReq.Header.Add("developer_uid", r.DEV_UID)
	httpReq.Header.Add("Content-Type", "application/json")

	// create client todo request
	cli := new(http.Client)

	resp, err := cli.Do(httpReq)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	// decode data
	respBody := struct {
		Data string `json:"data"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		log.Fatal(err.Error())
	}

	return respBody.Data
}
