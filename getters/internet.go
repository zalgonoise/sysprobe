package internet

import (
	"strconv"

	t "github.com/ZalgoNoise/sysprobe/types"
	u "github.com/ZalgoNoise/sysprobe/utils"
)

const ipDevice string = "wlan0"

// GetIP function - runs a simple `ip` command to retrieve the
// current information from the active network device, which
// returns a pointer to the Internet struct with this data
func GetIP() *t.Internet {
	i := &t.Internet{}

	ip, err := u.Run("ip", "-f", "inet", "addr", "show", ipDevice)
	u.Check(err)

	i.Device = u.Splitter(string(ip), ": ", 1)

	devIndex, err := strconv.Atoi(u.Splitter(string(ip), ": ", 0))
	u.Check(err)
	i.ID = devIndex

	i.IPAddress = u.Splitter(u.Splitter(u.Splitter(string(ip), "\n", 1), " ", 5), "/", 0)

	i.SubnetMask = u.Splitter(u.Splitter(string(ip), "\n", 1), " ", 7)

	return i

}
