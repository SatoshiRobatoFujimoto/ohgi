package sensu

import (
	"strings"
)

func PostRequest(check string, subscriber string) string {
	body := `{"check":"` + check + `","subscribers":["` + subscriber + `"]}`
	payload := strings.NewReader(body)
	_, status := postAPI("/request", payload)

	return httpStatus(status)
}
