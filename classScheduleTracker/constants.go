package main

const SUBCOMMAND_REGISTER = "register"
const SUBCOMMAND_VIEW = "view"
const SUBCOMMAND_ALARM = "alarm"
const SCHEDULE_FORMAT = `
- The JSON must contain a single object, with each of the weekdays as keys.
- The values corresponding to each weekday must be an object, where each key-value pair corresponds to an item on your schedule (assign an empty object if there is nothing scheduled on that day).
- For each scheduled item, its key must indicate the time (string) -- in military time, i.e. 24-hour format, no space/colons -- of the item, and the value must indicate the name/description (string) of the item.

The following is an example:

{
	"Monday": { "1120": "Computer Networks", "1500": "Technical Writing" },
	"Tuesday": {
			"1120": "Computer Networks",
			"1500": "Image and Vision Processing",
			"1600": "Technical Writing"
	},
	"Wednesday": {
			"1120": "Computer Networks",
			"1500": "Image and Vision Processing"
	},
	"Thursday": { "1500": "Image and Vision Processing" },
	"Friday": {},
	"Saturday": {},
	"Sunday": {}
}
`
