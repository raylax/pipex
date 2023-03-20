package worker

import (
	"github.com/raylax/pipex/utils"
	"regexp"
	"time"
)

// Timeout 超时时间 eg. 3600s 60m 1h
type Timeout string

const timeoutDefault Timeout = "30m"

var timeoutRegex = regexp.MustCompile("^(\\d+)([sSmMhH])$")

func (t Timeout) ToDuration() time.Duration {
	if string(t) == "" {
		return timeoutDefault.ToDuration()
	}
	strings := timeoutRegex.FindStringSubmatch(string(t))
	if len(strings) != 3 {
		return timeoutDefault.ToDuration()
	}
	value := utils.ParseInt[int64](strings[1])
	var unit time.Duration
	switch strings[2] {
	case "s", "S":
		unit = time.Second
	case "m", "M":
		unit = time.Minute
	case "h", "H":
		unit = time.Hour
	default:
		return timeoutDefault.ToDuration()
	}
	return time.Duration(value) * unit
}
