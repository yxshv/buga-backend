package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func parseArray(jsonBuffer []byte) ([]string, error) {
	var ids []string

	if err := json.Unmarshal(jsonBuffer, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

func getUUID() []string, error {

	var (
		resp *http.Response
		err  error
	)

	if resp, err = http.Get("https://www.uuidtools.com/api/generate/v1"); err != nil {
		return nil,err
	}

	defer resp.Body.Close()

	var result []byte

	if result, err = io.ReadAll(resp.Body); err != nil {
		return nil,err
	}

	var ids []string

	if ids, err = parseArray(result); err != nil {
		return nil,err
	}

	return ids, nil
}

func MakeMessage(content string) error {

	var err error

	query := fmt.Sprintf("INSERT INTO messages (id, content) VALUES ( '%s', '%s' )", ids[0], content)

	if _, err = db.Exec(query); err != nil {
		return err
	}

	return nil

}
