package shelly_v1

import (
	"shelly-prometheus-exporter/shelly"
)

type ShellyV1 struct {
	*shelly.Shelly
	status *Status
}

func New(targetHost string) *ShellyV1 {
	return &ShellyV1{
		Shelly: &shelly.Shelly{TargetHost: targetHost},
	}
}