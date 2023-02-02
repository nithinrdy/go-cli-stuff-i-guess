package main

import (
	"flag"
	"go/build"
	"os"
	"path/filepath"
)

func handleRegister(registerCmd *flag.FlagSet, registerScheduleFlag *string, registerTemplateFlag *bool) {
	if *registerScheduleFlag != "" {
		workingDir, err := os.Getwd()
		handleError(err)
		filebytes, err2 := os.ReadFile(filepath.Join(workingDir, *registerScheduleFlag))
		handleError(err2)

		err3 := os.MkdirAll(filepath.Join(build.Default.GOPATH, "config", "classScheduleTracker"), 0755)
		handleError(err3)

		err4 := os.WriteFile(
			filepath.Join(build.Default.GOPATH, "config", "classScheduleTracker", "schedule.json"), filebytes, 0644)
		handleError(err4)
		return
	}

	if *registerTemplateFlag {
		println(SCHEDULE_FORMAT)
		return
	}
}
