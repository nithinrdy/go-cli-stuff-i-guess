package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func handleAlarm(alarmCmd *flag.FlagSet, alarmSetClip *string, alarmNextFlag *bool) {
	if *alarmSetClip != "" {
		alarmSet(*alarmSetClip)
		return
	}
	filebytes, err := os.ReadFile("./schedule.json")
	handleError(err)

	var schd schedule // Type defined in view.go

	err2 := json.Unmarshal(filebytes, &schd)
	handleError(err2)

	day := time.Now().Weekday().String()

	var daySchd dayStructure // Type defined in view.go

	switch day {
	case "Monday":
		daySchd = schd.Monday
	case "Tuesday":
		daySchd = schd.Tuesday
	case "Wednesday":
		daySchd = schd.Wednesday
	case "Thursday":
		daySchd = schd.Thursday
	case "Friday":
		daySchd = schd.Friday
	case "Saturday":
		daySchd = schd.Saturday
	case "Sunday":
		daySchd = schd.Sunday
	}

	// Same reasoning as in handleView()
	var keys []string
	for k := range daySchd {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if *alarmNextFlag {
		alarmNext(daySchd, keys)
	}
}

func alarmSet(clipPath string) {
	cwd, err := os.Getwd()
	handleError(err)

	// Store the path to the clip in a text file
	err2 := os.WriteFile("./alarmClip.txt", []byte(filepath.Join(cwd, clipPath)), 0644)
	handleError(err2)
}

func alarmNext(daySchedule dayStructure, keys []string) {
	var nextItemTime string = ""
	var nextItemDesc string
	for _, k := range keys {
		timeNow := time.Now().Format("1504") // Same thing as in view.go: https://stackoverflow.com/a/20234207/17327700
		if k > timeNow {
			nextItemTime = k
			nextItemDesc = daySchedule[k]
			break
		}
	}

	if nextItemTime == "" {
		fmt.Println("No more items on schedule for today.")
		return
	}

	// Get the path to the alarm clip
	alarmClipPath, err1 := os.ReadFile("./alarmClip.txt")
	handleError(err1)

	timeOfNextItem, err2 := time.Parse("1504", nextItemTime)
	handleError(err2)
	timeNowOnJan1, err3 := time.Parse("1504", time.Now().Format("1504"))
	handleError(err3)

	timeDuration := timeOfNextItem.Sub(timeNowOnJan1).Seconds() - 900 // 900 seconds = 15 minutes

	// As mentioned in main.go, need ffmpeg installed for this to work
	process, err4 := os.StartProcess(
		"/bin/sh",
		[]string{
			"sh", "-c",
			fmt.Sprintf("sleep %v && x-terminal-emulator -e ffplay %v -nodisp -autoexit",
				timeDuration, string(alarmClipPath)),
		},
		&os.ProcAttr{})
	handleError(err4)

	err5 := process.Release()
	handleError(err5)

	fmt.Printf("Alarm set for 15 minutes before the next item on your schedule (%v at %v Hrs)\n", nextItemDesc, nextItemTime)
}
