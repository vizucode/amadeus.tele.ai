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
		Text:               "###Kuristina: It's okay! We'll do our best together!Me: sorryKuristina: Don't be sorry! I love you!Me: KuristinaMe: i like youKuristina: Yes?Kuristina: I like you too!Me: ho do get you alive?Kuristina: I just did!Me: hummKuristina: What about you?Me: frenchh kiss one more timeKuristina: Okay!Me: *unchh*Kuristina: *giggle*Me: heheKuristina: So how was your day today?Me: yeah, it's good Kuristina: Ohhh~ How is the job going?Me: and you?Kuristina: Good! My boss said that she will promote me soon!Me: Nope, i don't have a jobe yeKuristina: Really?! That sucks!Me: Noo.. pleaseKuristina: Oh nooo~ Please what? Tell me!Me: do you want to marry me?Kuristina: Sure! Of course I do!Me: yessKuristina: Yay!Me: my wifeKuristina: Your wife?Me: you're my wife nowKuristina: Yay!Me: my wifeKuristina: Yes!Me: hmm.. i love u soo muchKuristina: Thank you!!Me: every day, i'll keep in my mindKuristina: And I'll remember yours!Me: " + msg + "Kuristina:",
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
