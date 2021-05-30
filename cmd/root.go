/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomodorGo",
	Short: "Pomodoro CLI application",
	Long: `
██████   ██████  ███    ███  ██████  ██████   ██████  ██████   ██████   ██████  
██   ██ ██    ██ ████  ████ ██    ██ ██   ██ ██    ██ ██   ██ ██       ██    ██ 
██████  ██    ██ ██ ████ ██ ██    ██ ██   ██ ██    ██ ██████  ██   ███ ██    ██ 
██      ██    ██ ██  ██  ██ ██    ██ ██   ██ ██    ██ ██   ██ ██    ██ ██    ██ 
██       ██████  ██      ██  ██████  ██████   ██████  ██   ██  ██████   ██████  
                                                                                
                                                                                
pomodorGo is a CLI application written in Go that enables productivity.
This application is a tool to manage your work time in chunks
to work more efficiently.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nKeyboard interrupt detected exiting pomodorGo...")
		os.Exit(1)
	}()
}
