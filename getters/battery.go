package getters

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/ZalgoNoise/sysprobe/types"
	"github.com/ZalgoNoise/sysprobe/utils"
)

// customizable folder variable according to your OS
const powerFolder string = "battery"

// GetBattery function - collects battery related values
// from /sys/class/power_supply/*/uevent, and returns a
// pointer to the Battery struct with this data
func GetBattery() *types.Battery {

	b := &types.Battery{}

	batteryFile := "/sys/class/power_supply/" + powerFolder + "/uevent"

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
