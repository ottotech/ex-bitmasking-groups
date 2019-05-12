package groups

import (
	"errors"
)

type Group int

const (
	GroupA Group = 1 // 00000001
	GroupB Group = 2 // 00000010
	GroupC Group = 4 // 00000100
	GroupD Group = 8 // 00001000
)

var RegisteredGroups = []Group{GroupA, GroupB, GroupC, GroupD} // order matters

type GroupData struct {
	GroupName   string
	GroupConfig Group
}

func GetGroupName(g Group) (string, error) {
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
		if config&int(groupConfig) > 0 {
			name, _ := GetGroupName(groupConfig)
			group := GroupData{name, groupConfig}
			groups = append(groups, group)
		}
	}
	return groups
}

func BelongsToGroup(g Group, config int) bool {
	if config&int(g) > 0 {
		return true
	}
	return false
}
