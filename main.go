package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "cmctl"}

	workerCmd := &cobra.Command{
		Use: "worker",
		//Short: "Run the application",
	}

	serviceCmd := &cobra.Command{
		Use: "service",
		//Short: "Inspect the application",
	}

	// Add commands to the root command
	rootCmd.AddCommand(workerCmd)
	rootCmd.AddCommand(serviceCmd)

	//cmctl worker add
	addWorkerCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a worker",
		Run:   addWorkerFunction,
	}
	addWorkerCmd.Flags().StringP("worker_id", "w", "", "Worker name")
	addWorkerCmd.Flags().StringP("addr", "a", "", "Worker address")
	workerCmd.AddCommand(addWorkerCmd)

	//cmctl worker get
	getWorkerCmd := &cobra.Command{
		Use:   "get",
		Short: "Get a worker/all workers",
		Run:   getWorkerFunction,
	}

	getWorkerCmd.Flags().StringP("worker_id", "w", "", "Worker name")
	getWorkerCmd.Flags().BoolP("all", "a", false, "Get all workers")
	workerCmd.AddCommand(getWorkerCmd)

	//cmctl service add
	addServiceCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a service",
		Run:   addServiceFunction,
	}
	addServiceCmd.Flags().StringP("name", "n", "", "Service name")
	addServiceCmd.Flags().StringP("image", "i", "", "Service image")
	serviceCmd.AddCommand(addServiceCmd)

	//cmctl service get
	getServiceCmd := &cobra.Command{
		Use:   "get",
		Short: "Get a service/all services",
		Run:   getServiceFunction,
	}

	getServiceCmd.Flags().StringP("name", "n", "", "Service name")
	getServiceCmd.Flags().BoolP("all", "a", false, "Get all services")
	serviceCmd.AddCommand(getServiceCmd)

	//cmctl start
	startServiceCmd := &cobra.Command{
		Use:   "start",
		Short: "Start a service container",
		Run:   startServiceFunction,
	}

	startServiceCmd.Flags().StringP("worker_id", "w", "", "Worker name")
	startServiceCmd.Flags().StringP("file", "f", "", "start json file path")
	rootCmd.AddCommand(startServiceCmd)

	//cmctl run
	runServiceCmd := &cobra.Command{
		Use:   "run",
		Short: "Run a service",
		Run:   runServiceFunction,
	}

	runServiceCmd.Flags().StringP("worker_id", "w", "", "Worker name")
	runServiceCmd.Flags().StringP("service", "s", "", "Service name")
	runServiceCmd.Flags().StringP("file", "f", "", "run json file path")
	rootCmd.AddCommand(runServiceCmd)

	//cmctl checkpoint
	checkpointServiceCmd := &cobra.Command{
		Use:   "checkpoint",
		Short: "Checkpoint a service",
		Run:   checkpointServiceFunction,
	}

	checkpointServiceCmd.Flags().StringP("worker_id", "w", "", "Worker name")
	checkpointServiceCmd.Flags().StringP("service", "s", "", "Service name")
	checkpointServiceCmd.Flags().StringP("file", "f", "", "checkpoint json file path")
	rootCmd.AddCommand(checkpointServiceCmd)

	//cmctl migrate
	migrateServiceCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate a service",
		Run:   migrateServiceFunction,
	}

	migrateServiceCmd.Flags().StringP("service", "s", "", "Service name")
	migrateServiceCmd.Flags().StringP("source", "e", "", "Source worker name")
	migrateServiceCmd.Flags().StringP("destination", "d", "", "Destination worker name")
	migrateServiceCmd.Flags().StringP("copt", "c", "", "Checkpoint options")
	migrateServiceCmd.Flags().StringP("ropt", "r", "", "Run options")
	migrateServiceCmd.Flags().StringP("sopt", "t", "", "Start options")
	migrateServiceCmd.Flags().BoolP("stop", "p", false, "Stop the service")
	migrateServiceCmd.Flags().StringP("allopt", "a", "", "all opt file")

	rootCmd.AddCommand(migrateServiceCmd)

	//cmctl remove
	removeServiceCmd := &cobra.Command{
		Use:  "remove",
		Long: "Remove a service from a worker",
		Run:  removeServiceFunction,
	}

	removeServiceCmd.Flags().StringP("worker_id", "w", "", "Worker name")
	removeServiceCmd.Flags().StringP("service", "s", "", "Service name")

	rootCmd.AddCommand(removeServiceCmd)

	//cmctl stop
	stopServiceCmd := &cobra.Command{
		Use:  "stop",
		Long: "Stop a service",
		Run:  stopServiceFunction,
	}

	stopServiceCmd.Flags().StringP("worker_id", "w", "", "Worker name")
	stopServiceCmd.Flags().StringP("service", "s", "", "Service name")

	rootCmd.AddCommand(stopServiceCmd)

	// Execute the application
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
