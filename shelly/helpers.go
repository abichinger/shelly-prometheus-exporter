package shelly

import (
	"strconv"
)

func BoolToFloat(b bool) float64 {
	if b {
		return 1
	}

	return 0
}

func DeviceLabels(s *Shelly) []string {
	return []string{s.TargetHost}
}

func LineLabels(s *Shelly, kind string, line int) []string {
	return append(DeviceLabels(s), kind+":"+strconv.Itoa(line))
}
