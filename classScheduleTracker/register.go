package main

import (
	"flag"
	"os"
	"path/filepath"
)

func handleRegister(registerCmd *flag.FlagSet, registerScheduleFlag *string, registerTemplateFlag *bool) {
	if *registerScheduleFlag != "" {
		workingDir, err := os.Getwd()
		handleError(err)
		filebytes, err2 := os.ReadFile(filepath.Join(workingDir, *registerScheduleFlag))
		handleError(err2)

		err = os.WriteFile("./schedule.json", filebytes, 0644)
		handleError(err)
		return
	}

	if *registerTemplateFlag {
		println(SCHEDULE_FORMAT)
		return
	}
}
