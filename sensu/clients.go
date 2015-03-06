package sensu

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type clientStruct struct {
	Name          string
	Address       string
	Subscriptions []string
	Timestamp     int64
}

func GetClients(limit int, offset int) string {
	var clients []clientStruct
	var result []byte

	contents, status := getAPI(fmt.Sprintf("/clients?limit=%d&offset=%d", limit, offset))
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &clients)
	if len(clients) == 0 {
		return "No clients\n"
	}

	result = append(result, bold("NAME                                    ADDRESS                                 TIMESTAMP\n")...)
	for i := range clients {
		c := clients[i]
		line := fillSpace(c.Name, 40) + fillSpace(c.Address, 40) + utoa(c.Timestamp) + "\n"
		result = append(result, line...)
	}

	return string(result)
}

func GetClientsClient(client string) string {
	var c clientStruct
	var result []byte

	contents, status := getAPI("/clients/" + client)
	if status != 200 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	json.Unmarshal(contents, &c)

	result = append(result, (bold("NAME           ") + c.Name + "\n")...)
	result = append(result, (bold("ADDRESS        ") + c.Address + "\n")...)
	result = append(result, (bold("SUBSCRIPTIONS  ") + strings.Join(c.Subscriptions, ", ") + "\n")...)
	result = append(result, (bold("TIMESTAMP      ") + utoa(c.Timestamp) + "\n")...)

	return string(result)
}

func DeleteClientsClient(client string) string {
	_, status := deleteAPI("/clients/" + client)
	if status != 202 {
		fmt.Println(httpStatus(status))
		os.Exit(1)
	}

	return httpStatus(status)
}
