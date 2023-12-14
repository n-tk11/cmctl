package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types/mount"
	"github.com/tv42/httpunix"
)

type addWorkerBody struct {
	Worker_id string `json:"worker_id"`
	Addr      string `json:"addr"`
}

type addServiceBody struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
type demBody struct {
	dummy string
}

type startBody struct {
	ContainerName string        `json:"container_name"`
	Image         string        `json:"image"`
	AppPort       string        `json:"app_port"`
	Envs          []string      `json:"envs"`
	Mounts        []mount.Mount `json:"mounts"`
	Caps          []string      `json:"caps"`
}

type runBody struct {
	AppArgs        string   `json:"app_args"`
	ImageURL       string   `json:"image_url"`
	OnAppReady     string   `json:"on_app_ready"`
	PassphraseFile string   `json:"passphrase_file"`
	PreservedPaths string   `json:"preserved_paths"`
	NoRestore      bool     `json:"no_restore"`
	AllowBadImage  bool     `json:"allow_bad_image"`
	LeaveStopped   bool     `json:"leave_stopped"`
	Verbose        int      `json:"verbose"`
	Envs           []string `json:"envs"`
}

type checkpointBody struct {
	LeaveRun      bool     `json:"leave_running"`
	ImgUrl        string   `json:"image_url"`
	Passphrase    string   `json:"passphrase_file"`
	Preserve_path string   `json:"preserved_paths"`
	Num_shards    int      `json:"num_shards"`
	Cpu_budget    string   `json:"cpu_budget"`
	Verbose       int      `json:"verbose"`
	Envs          []string `json:"envs"`
}

type migrateBody struct {
	Copt checkpointBody `json:"copt"`
	Ropt runBody        `json:"ropt"`
	Sopt startBody      `json:"sopt"`
	Stop bool           `json:"stop"`
}

type Bodyer interface {
	demBody | addWorkerBody | addServiceBody | startBody | runBody | checkpointBody | migrateBody
}

func sendRequest[B Bodyer](url string, requestBody *B, verb string) {
	u := &httpunix.Transport{}
	u.RegisterLocation("man_sock", "/var/run/cm_man.sock")

	// If you want to use http: with the same client:
	t := &http.Transport{}
	t.RegisterProtocol(httpunix.Scheme, u)
	var client = http.Client{
		Transport: t,
	}
	//url := "http+unix://man_sock/cm_manager/v1.0/worker"

	// Convert the struct to JSON
	var req *http.Request
	if requestBody != nil {
		jsonData, err := json.Marshal(*requestBody)
		if err != nil {
			fmt.Printf("Error encoding JSON: %v\n", err)
		}
		req, err = http.NewRequest(verb, url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
		}

		// Set the Content-Type header to application/json
		req.Header.Set("Content-Type", "application/json")
	} else {
		var err error
		req, err = http.NewRequest(verb, url, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error performing request: %v\n", err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := decodeResponse(resp)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
	}

	fmt.Printf("Response: %s\n", responseBody)
}

func decodeResponse(resp *http.Response) (string, error) {
	var buffer bytes.Buffer
	_, err := buffer.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
