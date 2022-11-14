package skipmap

import (
	_ "unsafe" // for linkname

	"github.com/songzhibin97/go-baseutils/internal/wyhash"
	"github.com/songzhibin97/go-baseutils/sys/fastrand"
)

const (
	maxLevel            = 16
	p                   = 0.25
	defaultHighestLevel = 3
)

func hash(s string) uint64 {
	return wyhash.Sum64String(s)
}

//go:linkname cmpstring runtime.cmpstring
func cmpstring(a, b string) int

func randomLevel() int {
	level := 1
	for fastrand.Uint32n(1/p) == 0 {
		level++
	}
	if level > maxLevel {
		return maxLevel
	}
	return level
}
