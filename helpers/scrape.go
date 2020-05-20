package helpers

import (
	"io/ioutil"
	"net/http"
	"strings"
)


func UrlBuilder(tech string) string {
	rawUrl := "https://raw.githubusercontent.com/github/gitignore/master/"
	url := rawUrl + strings.Title(tech) + ".gitignore"
	return url
}

func GetIgnore(tech string) ([]byte, error) {
	url := UrlBuilder(tech)

	resp, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}