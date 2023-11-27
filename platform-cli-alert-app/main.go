package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

func main() {

	checkArgs()
	handleTimerInput()

}

const (
	// AppName is the name of the application
	AppName   = "PLATFORM_CLI_ALERT_APP"
	markValue = "1"
)

// check if correct arguments are passed
func checkArgs() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <hh:mm - alert-reminder-time> <Say what you want to be reminded - alert-message> \n", os.Args[0])
		os.Exit(1)
	}
}

func handleTimerInput() {
	// create the time instance
	now := time.Now()

	// create an instance of the when function
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	// get the arguments and use the time arg
	t, err := w.Parse(os.Args[1], now)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if t == nil {
		fmt.Printf("No match for %s\nNot a valid time\n", os.Args[1])
		os.Exit(2)
	}
	if now.After(t.Time) {
		fmt.Println("Please enter a future time")
		os.Exit(3)
	}

	// if args is correct then get diff
	diff := t.Time.Sub(now)

	fmt.Println("Reminder set to display in %v", diff.Round(time.Second))
	time.Sleep(diff)

	beepErr := beeep.Notify("Reminder", strings.Join(os.Args[2:], " "), "assets/information.png")
	if beepErr != nil {
		fmt.Println(beepErr)
		os.Exit(4)
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", AppName, markValue))

		osVal := runtime.GOOS
		if osVal == "darwin" {
			exec.Command("say", "Time to close work and call my baby girl").Start()
		}
		if err := cmd.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		os.Exit(0)
	}
}
