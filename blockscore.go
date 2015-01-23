package blockscore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// The BlockScore API key used to authenticate API requests.
var apiKey string

// The base URL for all BlockScore API requests.
const apiUrl string = "https://api.blockscore.com"

// The most recent BlockScore API version.
var apiVersion int = 4

// Sets the BlockScore API key to use for authenticating API requests.
func SetKey(key string) {
	apiKey = key
}

func SetKeyEnv() (err error) {
	apiKey = os.Getenv("BLOCKSCORE_API_KEY")
	if apiKey == "" {
		err = errors.New("BLOCKSCORE_API_KEY not found in environment")
	}
	return
}

// Sets the version of the BlockScore API to use.
// Note: This library is designed to only work with v4 and newer.
func SetVersion(version int) {
	apiVersion = version
}

var (
	People       = new(PersonClient)
	Companies    = new(CompanyClient)
	Candidates   = new(CandidateClient)
	QuestionSets = new(QuestionSetClient)
	Watchlists   = new(WatchlistClient)
)

// Error encapsulates an error returned by the BlockScore REST API.
type Error struct {
	Code   int
	Detail struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Param   string `json:"param"`
		Type    string `json:"type"`
	} `json:"error"`
}

func (e *Error) Error() string {
	return e.Detail.Message
}

func query(method, path string, values url.Values, v interface{}) error {
	// Parse the API URL.
	endpoint, err := url.Parse(apiUrl)
	if err != nil {
		return err
	}

	// Set the endpoint for the specific API.
	endpoint.Path = path
	endpoint.User = url.User(apiKey)

	// If this is a GET request, add the url.Values to the endpoint.
	if method == "GET" {
		endpoint.RawQuery = values.Encode()
	}

	// Else if this is not a GET, encode the url.Values in the body
	var reqBody io.Reader
	if method != "GET" && values != nil {
		reqBody = strings.NewReader(values.Encode())
	}

	// Create the request.
	req, err := http.NewRequest(method, endpoint.String(), reqBody)
	if err != nil {
		return err
	}

	// set Accept header with correct BlockScore API version.
	blockscoreHeader := fmt.Sprintf("application/vnd.blockscore+json;version=%d", apiVersion)

	// Set the request headers.
	req.Header.Set("Accept", blockscoreHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(values.Encode())))
	req.SetBasicAuth(apiKey, "")

	// Submit the http request.
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Read the body of the http message into a byte array.
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	// Is this an error?
	if r.StatusCode != 200 && r.StatusCode != 201 {
		fmt.Printf("StatusCode: %d\n", r.StatusCode)
		error := Error{}
		json.Unmarshal(body, &error)
		return &error
	}

	// Parse the JSON response into the response object.
	return json.Unmarshal(body, v)
}
