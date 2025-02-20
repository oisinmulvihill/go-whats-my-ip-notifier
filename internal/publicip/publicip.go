package publicip

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetAddress(ifconfigURL string) (address string, err error) {

	client := &http.Client{}

	r, err := http.NewRequest("GET", ifconfigURL, nil)
	if err != nil {
		return string(""), err
	}

	res, err := client.Do(r)
	if err != nil {
		return string(""), err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return string(""), err
	}

	ipAddress := strings.TrimSpace(string(resBody))

	return ipAddress, nil
}
