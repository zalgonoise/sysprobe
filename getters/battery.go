package battery

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	t "github.com/ZalgoNoise/sysprobe/types"
	u "github.com/ZalgoNoise/sysprobe/utils"
)

// customizable folder variable according to your OS
const powerFolder string = "battery"

// GetBattery function - collects battery related values
// from /sys/class/power_supply/*/uevent, and returns a
// pointer to the Battery struct with this data
func GetBattery() *t.Battery {

	b := &t.Battery{}

	batteryFile := "/sys/class/power_supply/" + powerFolder + "/uevent"

	bat, err := os.Open(batteryFile)
	u.Check(err)
	defer bat.Close()

	scanner := bufio.NewScanner(bat)

	line := 0

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "STATUS") {
			val := u.Splitter(scanner.Text(), "=", 1)
			b.Status = val

		} else if strings.Contains(scanner.Text(), "HEALTH") {
			val := u.Splitter(scanner.Text(), "=", 1)
			b.Health = val

		} else if strings.Contains(scanner.Text(), "CAPACITY") {
			val := u.Splitter(scanner.Text(), "=", 1)
			intVal, err := strconv.Atoi(val)
			u.Check(err)
			b.Capacity = intVal

		} else if strings.Contains(scanner.Text(), "TEMP=") {
			val := u.Splitter(scanner.Text(), "=", 1)
			intVal, err := strconv.Atoi(val)
			u.Check(err)
			var floatVal float32 = (float32(intVal) / 10)
			b.Temp.Internal = floatVal

		} else if strings.Contains(scanner.Text(), "TEMP_AMBIENT") {
			val := u.Splitter(scanner.Text(), "=", 1)
			intVal, err := strconv.Atoi(val)
			u.Check(err)
			var floatVal float32 = (float32(intVal) / 10)
			b.Temp.Ambient = floatVal
		}

		line++
	}

	err = scanner.Err()
	u.Check(err)

	return b

}
