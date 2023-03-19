package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"os"
)

// //go:embed auth_data.json
// var auth_data string

type AuthInput struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func getAuthData() map[string]map[string]string {
	auth_data, _ := os.ReadFile(os.Getenv("AUTH_DATA_FILE"))
	var auth_data_map map[string]map[string]string
	_ = json.Unmarshal([]byte(auth_data), &auth_data_map)
	return auth_data_map
}

func main() {
	in_bytes, _ := io.ReadAll(os.Stdin)
	var auth_input AuthInput
	json.Unmarshal(in_bytes, &auth_input)
	auth_data_map := getAuthData()
	var overrides map[string]string = auth_data_map[auth_input.User+":"+auth_input.Pass]
	if overrides != nil {
		data := map[string]string{}
		data["user"] = auth_input.User
		data["pass"] = auth_input.Pass
		data["_obscure"] = "pass"
		if overrides["type"] == "local" {
			data["_root"] = "/data/webdav/" + auth_input.User
		} else {
			data["_root"] = ""
		}
		for k, v := range overrides {
			data[k] = v
		}
		bytes, _ := json.Marshal(data)
		os.Stdout.WriteString(string(bytes))
	} else {
		os.Exit(1)
	}
}
