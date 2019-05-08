package groups

import "errors"

type Group int

const (
	GroupA int = 1 // 00000001
	GroupB int = 2 // 00000010
	GroupC int = 4 // 00000100
	GroupD int = 8 // 00001000
)

var RegisteredGroups = []int{GroupA, GroupB, GroupC, GroupD}

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

func BelongsToGroupA(groupConfig int) bool {
	if groupConfig&GroupA > 1 {
		return true
	}
	return false
}

func BelongsToGroupB(groupConfig int) bool {
	if groupConfig&GroupB > 1 {
		return true
	}
	return false
}

func BelongsToGroupC(groupConfig int) bool {
	if groupConfig&GroupC > 1 {
		return true
	}
	return false
}

func BelongsToGroupD(groupConfig int) bool {
	if groupConfig&GroupD > 1 {
		return true
	}
	return false
}
