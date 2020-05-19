package helpers

import (
	"io/ioutil"
	"net/http"
)


func UrlBuilder(tech string) string {
	rawUrl := "https://raw.githubusercontent.com/github/gitignore/master/"
	url := rawUrl + tech + ".gitignore"
	return url
}

func getIgnore(tech string) (string, error) {
	url := UrlBuilder(tech)

	resp, getErr := http.Get(url)
	if getErr != nil {
		return "", getErr
	}
	defer resp.Body.Close()

	bodyBytes, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		return "", ioErr
	}
	body := string(bodyBytes)

	return body, nil
}