// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"net"
	"strings"

	"github.com/sdeoras/ping-pong/config"
	"github.com/sdeoras/ping-pong/pb"
	"github.com/sdeoras/ping-pong/server"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Long:  "",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag("/run/host", cmd.Flags().Lookup("host"))
	_ = viper.BindPFlag("/run/port", cmd.Flags().Lookup("port"))

	host := viper.GetString("/run/host")
	port := viper.GetString("/run/port")

	lis, err := net.Listen("tcp", strings.Join([]string{host, port}, ":"))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pingPongServer := server.NewPingPongServer()

	pb.RegisterPingPongServer(s, pingPongServer)

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.Flags().String("host", config.DefaultHost, "hostname")
	runCmd.Flags().String("port", config.DefaultPort, "port number")
}
