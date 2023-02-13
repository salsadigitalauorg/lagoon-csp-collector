package util

import (
	"net/url"
	"strings"
)

func ProjectFomDocumentURI(d string) (string, error) {
	url, err := url.Parse(d)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(url.Hostname(), "www."), nil
}
