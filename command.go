package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func addWorkerFunction(cmd *cobra.Command, args []string) {
	worker_id, _ := cmd.Flags().GetString("worker_id")
	addr, _ := cmd.Flags().GetString("addr")
	url := "http+unix://man_sock/cm_manager/v1.0/worker"
	requestBody := &addWorkerBody{
		Worker_id: worker_id,
		Addr:      addr,
	}
	sendRequest(url, requestBody, "POST")
}

func getWorkerFunction(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("all") {
		url := "http+unix://man_sock/cm_manager/v1.0/worker"
		sendRequest[demBody](url, nil, "GET")
		return
	}
	worker_id, _ := cmd.Flags().GetString("worker_id")
	url := "http+unix://man_sock/cm_manager/v1.0/worker/" + worker_id
	sendRequest[demBody](url, nil, "GET")
}

func addServiceFunction(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	image, _ := cmd.Flags().GetString("image")
	url := "http+unix://man_sock/cm_manager/v1.0/service"
	requestBody := &addServiceBody{
		Name:  name,
		Image: image,
	}
	sendRequest(url, requestBody, "POST")
}

func getServiceFunction(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("all") {
		url := "http+unix://man_sock/cm_manager/v1.0/service"
		sendRequest[demBody](url, nil, "GET")
		return
	}
	name, _ := cmd.Flags().GetString("name")
	url := "http+unix://man_sock/cm_manager/v1.0/service/" + name
	sendRequest[demBody](url, nil, "GET")
}

func startServiceFunction(cmd *cobra.Command, args []string) {
	worker_id, _ := cmd.Flags().GetString("worker_id")
	filePath, _ := cmd.Flags().GetString("file")

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse the JSON data into a struct
	var start startBody
	if err := json.Unmarshal(jsonData, &start); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	url := "http+unix://man_sock/cm_manager/v1.0/start/" + worker_id
	sendRequest(url, &start, "POST")
}

func runServiceFunction(cmd *cobra.Command, args []string) {
	worker_id, _ := cmd.Flags().GetString("worker_id")
	service, _ := cmd.Flags().GetString("service")
	filePath, _ := cmd.Flags().GetString("file")

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse the JSON data into a struct
	var run runBody
	if err := json.Unmarshal(jsonData, &run); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	url := "http+unix://man_sock/cm_manager/v1.0/run/" + worker_id + "/" + service
	sendRequest(url, &run, "POST")

}

func checkpointServiceFunction(cmd *cobra.Command, args []string) {
	worker_id, _ := cmd.Flags().GetString("worker_id")
	service, _ := cmd.Flags().GetString("service")
	filePath, _ := cmd.Flags().GetString("file")

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse the JSON data into a struct
	var chk checkpointBody
	if err := json.Unmarshal(jsonData, &chk); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	url := "http+unix://man_sock/cm_manager/v1.0/checkpoint/" + worker_id + "/" + service
	sendRequest(url, &chk, "POST")
}

func migrateServiceFunction(cmd *cobra.Command, args []string) {

	service, _ := cmd.Flags().GetString("service")
	src, _ := cmd.Flags().GetString("source")
	dest, _ := cmd.Flags().GetString("destination")
	url := "http+unix://man_sock/cm_manager/v1.0/migrate/" + service + "?src=" + src + "&dest=" + dest
	if cmd.Flags().Changed("allopt") {
		alloptPath, _ := cmd.Flags().GetString("allopt")
		jsonData, err := os.ReadFile(alloptPath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		var mgr migrateBody
		if err := json.Unmarshal(jsonData, &mgr); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return
		}
		sendRequest(url, &mgr, "POST")
		return
	}
	copt, _ := cmd.Flags().GetString("copt")
	ropt, _ := cmd.Flags().GetString("ropt")
	sopt, _ := cmd.Flags().GetString("sopt")
	stop, _ := cmd.Flags().GetBool("stop")
	requestBody := parseMigrateBody(copt, ropt, sopt, stop)
	sendRequest(url, &requestBody, "POST")
}

func removeServiceFunction(cmd *cobra.Command, args []string) {
	service, _ := cmd.Flags().GetString("service")
	worker, _ := cmd.Flags().GetString("worker_id")

	url := "http+unix://man_sock/cm_manager/v1.0/remove/" + worker + "/" + service
	sendRequest[demBody](url, nil, "DELETE")
}

func stopServiceFunction(cmd *cobra.Command, args []string) {
	service, _ := cmd.Flags().GetString("service")
	worker, _ := cmd.Flags().GetString("worker_id")

	url := "http+unix://man_sock/cm_manager/v1.0/stop/" + worker + "/" + service
	sendRequest[demBody](url, nil, "POST")
}

func parseMigrateBody(copt string, ropt string, sopt string, stop bool) migrateBody {
	var mgr migrateBody
	if copt != "" {
		jsonData, err := os.ReadFile(copt)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return migrateBody{}
		}
		var co checkpointBody
		if err := json.Unmarshal(jsonData, &co); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return migrateBody{}
		}
		mgr.Copt = co
	}
	if ropt != "" {
		jsonData, err := os.ReadFile(ropt)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return migrateBody{}
		}
		var ro runBody
		if err := json.Unmarshal(jsonData, &ro); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return migrateBody{}
		}
		mgr.Ropt = ro
	}
	if sopt != "" {
		jsonData, err := os.ReadFile(sopt)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return migrateBody{}
		}
		var so startBody
		if err := json.Unmarshal(jsonData, &so); err != nil {
			fmt.Printf("Error parsing JSON: %v\n", err)
			return migrateBody{}
		}
		mgr.Sopt = so
	}
	mgr.Stop = stop
	return mgr
}
