// Need ffmpeg installed on your machine for the alarm to work
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
	alarmCmd := flag.NewFlagSet("alarm", flag.ExitOnError)

	// Flags for the register subcommand
	registerSchedule := registerCmd.String("f", "", "Relative path to the JSON file to use to set up the weekly schedule.")
	registerTemplate := registerCmd.Bool("t", false, "Template to refer to for creating the JSON file")

	// Flags for the view subcommand
	viewNext := viewCmd.Bool("n", false, "View the next scheduled item (on the same day)")
	viewAll := viewCmd.Bool("a", false, "View the entire day's schedule")

	// Flags for the alarm subcommand
	alarmSetClip := alarmCmd.String("c", "", "Relative path to the file to use as the alarm audio")
	alarmNext := alarmCmd.Bool("n", false, "Set an alarm that alerts you 15 minutes before the next scheduled item starts")

	if len(os.Args) < 2 {
		fmt.Println("Expected a subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case SUBCOMMAND_REGISTER:
		if len(os.Args) < 3 {
			handleValidationError("Expected a filepath, or the --template flag")
		}
		registerCmd.Parse(os.Args[2:])
		handleRegister(registerCmd, registerSchedule, registerTemplate)

	case SUBCOMMAND_VIEW:
		if len(os.Args) < 3 {
			handleValidationError("Expected either the --next or the --all flag")
		}
		viewCmd.Parse(os.Args[2:])
		handleView(viewCmd, viewNext, viewAll)

	case SUBCOMMAND_ALARM:
		if len(os.Args) < 3 {
			handleValidationError("Expected either the --next or the --all flag")
		}
		alarmCmd.Parse(os.Args[2:])
		handleAlarm(alarmCmd, alarmSetClip, alarmNext)

	default:
		handleValidationError(
			fmt.Sprintf(
				"Expected one of the following subcommands: %v, %v, %v. Got: %v.",
				SUBCOMMAND_REGISTER, SUBCOMMAND_VIEW, SUBCOMMAND_ALARM, os.Args[1],
			),
		)
	}
}
