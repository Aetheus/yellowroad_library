package main

import (
	"github.com/spf13/cobra"
	"yellowroad_library/commands"
	"os"
	"log"
	"path"
	"yellowroad_library/config"
)


func main() {
	//get the working directory
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//load the configuration file, since some commands need it
	configPath := path.Join(workingDirectory,"config.json")
	configuration, err:= config.Load(configPath)
	if (err != nil){
		log.Fatal(err)
	}

	var rootCommand = &cobra.Command{
		Use:   "library_app",
		Short: "Defaults to running the 'server' command",
		Long:  `Defaults to running the 'server' command. See the readme.md for more details`,
		Run: func(cmd *cobra.Command, args []string) {
			commands.ServerCommand(configuration)
		},
	}

	var serverCommand = &cobra.Command {
		Use:   "server",
		Short: "Runs the server application",
		Long:  `Runs the server application. See the readme.md for more details`,
		Run: func(cmd *cobra.Command, args []string) {
			commands.ServerCommand(configuration)
		},
	}

	var testCommand = &cobra.Command {
		Use: "test",
		Short: "Runs unit tests via GoConvey",
		Long : "Runs unit tests using GoConvey. The output of those tests can be viewed with a web browser at the address that GoConvey displays. See readme.md for more details",
		Run: func(cmd *cobra.Command, args []string) {
			commands.TestCommand(workingDirectory)
		},
	}

	var migrateCommand = &cobra.Command{
		Use:   "migrate",
		Short: "Runs database migrations",
		Long:  `Runs database migrations in the database/migrations directory. See readme.md for more details`,
		Run: func(cmd *cobra.Command, args []string) {
			commands.MigrateCommand(configuration)
		},
	}


	rootCommand.AddCommand(migrateCommand, serverCommand,testCommand)
	rootCommand.Execute()
}