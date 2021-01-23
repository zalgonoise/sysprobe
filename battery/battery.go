package battery

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/ZalgoNoise/sysprobe/utils"
)

// Battery type will be converted to JSON
// containing important information for this module
type Battery struct {
	Status   string `json:"status"`
	Health   string `json:"health"`
	Capacity int    `json:"capacity"`
	Temp     struct {
		Internal float32 `json:"int"`
		Ambient  float32 `json:"ext"`
	} `json:"temp"`
}

// TermuxBattery struct serves as a fallback object in case the uevent
// file is not accessible to the user (root privileges are required)
// It is a slower approach and thus not being the primary way to get
// battery metadata
type TermuxBattery struct {
	Health   string  `json:"health"`
	Capacity int     `json:"percentage"`
	Plugged  string  `json:"plugged"`
	Status   string  `json:"status"`
	Temp     float32 `json:"temperature"`
	Current  int     `json:"current"`
}

// Get method for TermuxBattery objects will execute the
// `termux-battery-status` command, and retrieve the slice of bytes
// which are already a JSON object. The method will unmarshal the JSON
// object into a TermuxBattery object
func (tb *TermuxBattery) Get() (*TermuxBattery, bool) {

	exec, err := utils.Run("termux-battery-status")

	if err != nil {
		json.Unmarshal(exec, tb)
		return tb, true
	} else {
		return tb, false
	}
}

// Push method for TermuxBattery will inject the values from the
// TermuxBattery object into the Battery (message) object
func (tb *TermuxBattery) Push(b *Battery) *Battery {
	b.Capacity = tb.Capacity
	b.Health = tb.Health
	b.Status = tb.Status
	b.Temp.Internal = tb.Temp

	return b
}

// Get method - collects battery related values
// from /sys/class/power_supply/*/uevent, and returns a
// pointer to the Battery struct with this data
func (b *Battery) Get(batteryLoc string) *Battery {

	batteryFile := "/sys/class/power_supply/" + batteryLoc + "/uevent"

	bat, err := os.Open(batteryFile)
	if err != nil {
		tb := &TermuxBattery{}
		tb, success := tb.Get()

		if success != true {
			return b
		}

		tb.Push(b)
		return b
	}
	defer bat.Close()

	scanner := bufio.NewScanner(bat)

	line := 0

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "STATUS=") {
			val := utils.Splitter(scanner.Text(), "=", 1)
			b.Status = val

		} else if strings.Contains(scanner.Text(), "HEALTH=") {
			val := utils.Splitter(scanner.Text(), "=", 1)
			b.Health = val

		} else if strings.Contains(scanner.Text(), "CAPACITY=") {
			val := utils.Splitter(scanner.Text(), "=", 1)
			intVal, err := strconv.Atoi(val)
			utils.Check(err)
			b.Capacity = intVal

		} else if strings.Contains(scanner.Text(), "TEMP=") {
			val := utils.Splitter(scanner.Text(), "=", 1)
			intVal, err := strconv.Atoi(val)
			utils.Check(err)
			var floatVal float32 = (float32(intVal) / 10)
			b.Temp.Internal = floatVal

		} else if strings.Contains(scanner.Text(), "TEMP_AMBIENT=") {
			val := utils.Splitter(scanner.Text(), "=", 1)
			intVal, err := strconv.Atoi(val)
			utils.Check(err)
			var floatVal float32 = (float32(intVal) / 10)
			b.Temp.Ambient = floatVal
		}

		line++
	}

	err = scanner.Err()
	utils.Check(err)

	return b

}
