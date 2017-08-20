// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/guenter/osdetector-demo/osdetector"
	"github.com/spf13/cobra"
	"log"
)

var config = osdetector.OSDetectorConfig{}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Config: %+v", config)
		handler := osdetector.NewOSDetector(config)
		handler.Serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVar(&config.ServeAddress, "address", ":8080", "The address to run the HTTP server on")
	serveCmd.Flags().StringVar(&config.TemplateFile, "template", "index.html", "The filename of the template")
	serveCmd.Flags().StringArrayVar(&config.CassandraHosts, "cassandra-host", []string{"127.0.0.1"}, "Cassandra hosts to connect to")
	serveCmd.Flags().StringVar(&config.CassandraKeyspace, "cassandra-keyspace", "browsers", "The keyspace in Cassandra")
}
