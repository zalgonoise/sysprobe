package battery

import (
	"bufio"
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

// Get method - collects battery related values
// from /sys/class/power_supply/*/uevent, and returns a
// pointer to the Battery struct with this data
func (b *Battery) Get(batteryLoc string) *Battery {

	batteryFile := "/sys/class/power_supply/" + batteryLoc + "/uevent"

	bat, err := os.Open(batteryFile)
	utils.Check(err)
	defer bat.Close()

	scanner := bufio.NewScanner(bat)

	line := 0

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "STATUS") {
			val := utils.Splitter(scanner.Text(), "=", 1)
			b.Status = val

		} else if strings.Contains(scanner.Text(), "HEALTH") {
			val := utils.Splitter(scanner.Text(), "=", 1)
			b.Health = val

		} else if strings.Contains(scanner.Text(), "CAPACITY") {
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

		} else if strings.Contains(scanner.Text(), "TEMP_AMBIENT") {
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
