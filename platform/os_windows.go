// Copyright 2017 Microsoft. All rights reserved.
// MIT License

package platform

import (
	"time"
)

const (
	// RuntimePath is the path where runtime files are stored.
	RuntimePath = ""

	// LogPath is the path where log files are stored.
	LogPath = ""
)

// GetOSInfo returns OS version information.
func GetOSInfo() string {
	return "windows"
}

// GetLastRebootTime returns the last time the system rebooted.
func GetLastRebootTime() (time.Time, error) {
	var rebootTime time.Time
	return rebootTime, nil
}
