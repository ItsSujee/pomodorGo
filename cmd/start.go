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
	"strconv"
	"time"
)

var focusDurationFlag int
var restDurationFlag int

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start pomodorGo app",
	Long: `Start pomodorGo application. 
Default set to 25 minute sessions with 5 minute breaks.`,
	Run: func(cmd *cobra.Command, args []string) {
		//Clear terminal screen
		fmt.Print("\033[H\033[2J")

		fmt.Println("Starting pomodorGo application... (Press Control + C to exit)")

		pomodoro(focusDurationFlag, restDurationFlag)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntVarP(&focusDurationFlag, "focusDuration", "f", 25, "Duration of focus session in minutes")
	startCmd.Flags().IntVarP(&restDurationFlag, "restDuration", "r", 5, "Duration of rest session in minutes")
}

func pomodoro(focusLength int, restLength int) {
	for {
		focus(focusLength)
		rest(restLength)
	}
}

func focus(d int) {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	startTime := time.Now()
	focusTime,_ := time.ParseDuration(strconv.Itoa(d) + "m")
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("\rFocus Time Left: %s", (focusTime + startTime.Sub(t)).Round(time.Second).String() )
			}
		}
	}()
	time.Sleep(time.Duration(d) * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println()
}

func rest(d int) {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	startTime := time.Now()
	restTime,_ := time.ParseDuration(strconv.Itoa(d) + "m")
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("\rRest Time Left: %s", (restTime + startTime.Sub(t)).Round(time.Second).String() )
			}
		}
	}()
	time.Sleep(time.Duration(d) * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println()
}