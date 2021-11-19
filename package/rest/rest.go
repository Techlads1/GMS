// Package rest allows for quick and easy access any REST or REST-like API.
package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"gateway/package/log"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Client struct {
	user   string
	key    string
	secret string
	url    string //base url
}

type Response struct {
	StatusCode int
	Body       []byte
}

//New function creates a restful api client.
// baseUrl should be in the format: https://localhost:4321
func New(baseUrl string) *Client {
	//thread safe singletone initialised

	c := &Client{
		//user:   user,
		//key:    key,
		//secret: secret,
		url: baseUrl,
		//wsUrl:  WsUrl,
	}

	return c
}

//GetMethod call the end point and returns body in binary form
func (c *Client) GetMethod(u string) *Response {
	response := &Response{}
	res, err := http.Get(u)
	if err != nil {
		log.Errorf("error obtaining get response: %v", err)
		return response
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Errorf("error reading get response: %v", err)
		return response
	}
	response.StatusCode = res.StatusCode
	response.Body = data
	return response
}

func (c *Client) PostMethod(u string, v url.Values) *Response {
	response := &Response{}
	res, err := http.PostForm(u, v)
	if err != nil {
		log.Errorf("error obtaining post response: %v", err)
		return response
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Errorf("error reading post response body: %v", err)
		return response
	}
	response.StatusCode = res.StatusCode
	response.Body = data
	return response
}
func (c *Client) PostFile(u string, v url.Values, file *multipart.FileHeader) *Response {

	body := strings.NewReader(v.Encode())
	req, err := http.NewRequest("POST", u, body)
	if err != nil {
		log.Errorf("error preparing the request: %v", err)
		return nil
	}

	response := &Response{}
	res, err := http.PostForm(u, v)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		log.Errorf("error obtaining post response: %v", err)
		return response
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Errorf("error reading post response body: %v", err)
		return response
	}
	response.StatusCode = res.StatusCode
	response.Body = data
	return response
}

//Send consumes the api
// endpoint shoud be in this format '/user/list
// params should be filled by form values for post requests
// isPost should be true for post request and false for get request
//opt should be index, e.g 1, if combined with
func (c *Client) Send(endpoint string, params map[string]string, isPost bool) *Response {
	u := c.url + endpoint
	/*if opt != 0 {
		u = fmt.Sprintf("%s/%d", u, opt) //this forms /user/1
	}*/
	var res *Response
	if isPost {
		urlValues := url.Values{}
		for k, v := range params {
			urlValues.Add(k, v)
		}
		urlValues.Encode()

		res = c.PostMethod(u, urlValues)
	} else {
		// Get method for public method
		res = c.GetMethod(u)
	}
	return res
}

func (c *Client) SendFile(endpoint string, params map[string]string, file *multipart.FileHeader) *Response {
	u := c.url + endpoint

	var res *Response

	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Add(k, v)
	}
	urlValues.Encode()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/endpoint", body)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // <<< important part

	res = c.PostMethod(u, urlValues)

	return res
}

//after adding signature

func (c *Client) Nonce() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (c *Client) ToHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func (c *Client) Signature() (string, string) {
	nonce := c.Nonce()
	message := nonce + c.user + c.key
	signature := c.ToHmac256(message, c.secret)
	return signature, nonce
}
