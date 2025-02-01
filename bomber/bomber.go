package bomber


import (
	// "encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/IR-Bomber/IRB-Core/utils"
)

type Bomber struct {
	APIURL string
	Method string
	Headers map[string]string
	Payload map[string]interface{}
	Data []byte
	Proxy http.Transport
	Timeout time.Duration
}

type Status struct {
	OK bool
	StatusCode int
}

func (s *Status) SetStatus(statusCode int) *Status {
	s.StatusCode = statusCode
	s.OK = utils.ValidateStatus(statusCode)

	return s
}



func (b *Bomber) SetAPIURL(url string) *Bomber {
	b.APIURL = url

	b.Headers = map[string]string{}
	return b
}

func (b *Bomber) SetProxy(proxy string) *Bomber {
	
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("Error parsing proxy URL: ", err)
	}

	b.Proxy = http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	return b
}

func (b *Bomber) SetTimeout(timeout int) *Bomber {
	b.Timeout = time.Second * time.Duration(timeout)

	return b
}

func (b *Bomber) SetMethod(method string) *Bomber {
	b.Method = method
	return b
}

func (b *Bomber) SetHeaders(headers map[string]string) *Bomber {
	b.Headers = headers
	return b
}

func (b *Bomber) SetPayload(payload map[string]interface{}) *Bomber {
	b.Payload = payload
	if b.Payload != nil {
		b.Headers["Content-Type"] = "application/json"
	}
	return b
}

func (b *Bomber) SetData(data string) *Bomber {
	b.Data = []byte(data)
	if string(b.Data) != "" {
		b.Headers["Content-Type"] = "text/plain"
	}
	return b
}


func (b *Bomber) Send() (error, *Status) {
	s := Status{}
	if b.APIURL == "" {
		err := fmt.Errorf("API URL is required")
		fmt.Println(err)
		return err, s.SetStatus(0)
	}

	if b.Method == "" {
		err := fmt.Errorf("Method is required")
		fmt.Println(err)
		return err, s.SetStatus(0)
	}
	if b.Payload != nil {
		jsonBytes, err := json.Marshal(b.Payload)

		if err != nil {
			fmt.Println("Error marshalling payload: ", err)
			return fmt.Errorf("Error marshalling payload: %v", err), s.SetStatus(0)
		}

		b.Data = jsonBytes
	}
	req, err := http.NewRequest(b.Method, b.APIURL, bytes.NewBuffer(b.Data))

	if err != nil {
		// fmt.Println("Error creating request: ", err)
		return fmt.Errorf("Error creating request: %v", err), s.SetStatus(0)
	}

	for k, v := range b.Headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Transport: &b.Proxy,
		Timeout: b.Timeout,
	}

	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("Error sending request: %v", err), s.SetStatus(0)
	}

	defer resp.Body.Close()
	
	return nil, s.SetStatus(resp.StatusCode)
}