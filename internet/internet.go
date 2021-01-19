package internet

import (
	"strconv"

	"github.com/ZalgoNoise/sysprobe/utils"
)

// Internet type will be converted to JSON
// containing important information for this module
type Internet struct {
	Device     string `json:"device"`
	ID         int    `json:"id"`
	IPAddress  string `json:"ipv4"`
	SubnetMask string `json:"mask"`
}

// Get method - runs a simple `ip` command to retrieve the
// current information from the active network device, which
// returns a pointer to the Internet struct with this data
func (i *Internet) Get(ipDevice string) *Internet {

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
