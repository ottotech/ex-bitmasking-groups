package groups

import (
	"errors"
)

type Group int

const (
	GroupA int = 1 // 00000001
	GroupB int = 2 // 00000010
	GroupC int = 4 // 00000100
	GroupD int = 8 // 00001000
)

var RegisteredGroups = []int{GroupA, GroupB, GroupC, GroupD}  // order matters

type GroupData struct {
	GroupName   string
	GroupConfig int
}

type GroupMembership interface {
}

func GetGroupName(g int) (string, error) {
	if g == GroupA {
		return "Group A", nil
	}

	if g == GroupB {
		return "Group B", nil
	}

	if g == GroupC {
		return "Group C", nil
	}

	if g == GroupD {
		return "Group D", nil
	}

	return "Unknown group", errors.New("unknown group")

}

func GetAllGroups() []GroupData {
	var groups []GroupData
	for _, groupConfig := range RegisteredGroups {
		groupName, _ := GetGroupName(groupConfig)
		g := GroupData{GroupName: groupName, GroupConfig: groupConfig}
		groups = append(groups, g)
	}
	return groups
}

func GetGroupsByConfiguration(config int) []GroupData {
	var groups []GroupData
	for _, groupConfig := range RegisteredGroups {
		if config&groupConfig > 0 {
			name, _ := GetGroupName(groupConfig)
			group := GroupData{name, groupConfig}
			groups = append(groups, group)
		}
	}
	return groups
}
