package rest

import (
	"fmt"
	"github.com/RadikChernyshov/klingon/pkg/colorizer"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiUrl = "http://stapi.co/api/v1/rest"

func PostReq(uri string, payload *strings.Reader) (body []byte, err error) {
	url := fmt.Sprintf("%s%s", apiUrl, uri)
	req, _ := http.NewRequest("POST", url, payload)
	res, _ := http.DefaultClient.Do(req)
	body, _ = ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf(colorizer.Red("stapi api request failed status code: %d"), res.StatusCode)
	}
	defer res.Body.Close()
	return body, err
}

func GetReq(uri string) (body []byte, err error) {
	url := fmt.Sprintf("%s%s", apiUrl, uri)
	res, _ := http.Get(url)
	body, _ = ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf(colorizer.Red("stapi api request failed status code: %d"), res.StatusCode)
	}
	defer res.Body.Close()
	return body, err
}
