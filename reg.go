package main

import (
	"golang.org/x/sys/windows/registry"
)

func getKey() {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\NVIDIA Corporation\Global\NGXCore`, registry.QUERY_VALUE)

	value, _, _ := key.GetIntegerValue("ShowDlssIndicator")
	defer key.Close()

	indicatorStatus = int(value)

	if indicatorStatus == 1024 {
		indicatorStatusString = "ON"
	} else if indicatorStatus == 0 {
		indicatorStatusString = "OFF"
	} else {
		indicatorStatusString = "NOT SET"
	}
}
func setKeyOn() {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\NVIDIA Corporation\Global\NGXCore`, registry.SET_VALUE)

	key.SetDWordValue("ShowDlssIndicator", uint32(1024))
	defer key.Close()

}

func setKeyOff() {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\NVIDIA Corporation\Global\NGXCore`, registry.SET_VALUE)

	key.SetDWordValue("ShowDlssIndicator", uint32(0))
	defer key.Close()

}
