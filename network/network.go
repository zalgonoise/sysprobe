package network

import (
	"strconv"

	"github.com/ZalgoNoise/sysprobe/utils"
)

// Network type will be converted to JSON
// containing important information for this module
type Network struct {
	System System   `json:"sys"`
	Ping   PingScan `json:"ping"`
}

// System type will contain this device's network information
type System struct {
	Device     string `json:"device"`
	ID         int    `json:"id"`
	IPAddress  string `json:"ipv4"`
	SubnetMask string `json:"mask"`
}

// Get method - runs a simple `ip` command to retrieve the
// current information from the active network device, which
// returns a pointer to the Internet struct with this data
func (s *System) Get(ipDevice string) *System {

	ip, err := utils.Run("ip", "-f", "inet", "addr", "show", ipDevice)
	utils.Check(err)

	s.Device = utils.Splitter(string(ip), ": ", 1)

	devIndex, err := strconv.Atoi(utils.Splitter(string(ip), ": ", 0))
	utils.Check(err)
	s.ID = devIndex

	s.IPAddress = utils.Splitter(utils.Splitter(utils.Splitter(string(ip), "\n", 1), " ", 5), "/", 0)

	s.SubnetMask = utils.Splitter(utils.Splitter(string(ip), "\n", 1), " ", 7)

	return s

}
