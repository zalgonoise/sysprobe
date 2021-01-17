package internet

import (
	"strconv"

	"github.com/ZalgoNoise/sysprobe/types"
	"github.com/ZalgoNoise/sysprobe/utils"
)

const ipDevice string = "wlan0"

// GetIP function - runs a simple `ip` command to retrieve the
// current information from the active network device, which
// returns a pointer to the Internet struct with this data
func GetIP() *types.Internet {
	i := &types.Internet{}

	ip, err := utils.Run("ip", "-f", "inet", "addr", "show", ipDevice)
	utils.Check(err)

	i.Device = utils.Splitter(string(ip), ": ", 1)

	devIndex, err := strconv.Atoi(utils.Splitter(string(ip), ": ", 0))
	utils.Check(err)
	i.ID = devIndex

	i.IPAddress = utils.Splitter(utils.Splitter(utils.Splitter(string(ip), "\n", 1), " ", 5), "/", 0)

	i.SubnetMask = utils.Splitter(utils.Splitter(string(ip), "\n", 1), " ", 7)

	return i

}
