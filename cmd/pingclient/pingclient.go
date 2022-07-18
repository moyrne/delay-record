package main

import (
	"github.com/moyrne/delay-record/pkg/pingclient/service"
	"github.com/spf13/cobra"
)

var pingHost string

func main() {
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			service.Run(pingHost)
		},
	}

	rootCmd.PersistentFlags().StringVar(&pingHost, "host", "", "ping host")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
