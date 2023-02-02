package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type dayStructure map[string]string

type schedule struct {
	Monday    dayStructure
	Tuesday   dayStructure
	Wednesday dayStructure
	Thursday  dayStructure
	Friday    dayStructure
	Saturday  dayStructure
	Sunday    dayStructure
}

func handleView(viewCmd *flag.FlagSet, viewNextFlag *bool, viewAllFlag *bool) {
	filebytes, err := os.ReadFile(
		filepath.Join(build.Default.GOPATH, "config", "classScheduleTracker", "schedule.json"))
	handleError(err)

	var schd schedule

	err2 := json.Unmarshal(filebytes, &schd)
	handleError(err2)

	day := time.Now().Weekday().String()

	var daySchd dayStructure

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

	// So apparently iterating over a map is not guaranteed to be in any particular order?
	// We need extract the keys, sort them, and then iterate over the sorted array.
	// (They INTENTIONALLY randomized the order of iteration over maps so people won't rely on it being ordered.
	// I'm probably too dumb to understand so not gonna question it.)
	var keys []string
	for k := range daySchd {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if *viewNextFlag {
		fmt.Printf("\nNext item on schedule for today (%v):\n", day)

		for _, k := range keys {
			timeNow := time.Now().Format("1504") // This is a thing: https://stackoverflow.com/a/20234207/17327700
			if k > timeNow {
				fmt.Printf("%v at %v Hrs\n", daySchd[k], k)
				return
			}
		}

		println("No more items scheduled for the day!")
		return
	}

	if *viewAllFlag {
		fmt.Printf("\nSchedule for today (%v):\n\n", day)

		for _, k := range keys {
			fmt.Printf("%v at %v Hrs\n", daySchd[k], k)
		}

		return
	}

}
