package groups

type Group int

const (
	GroupA Group = 1 // 00000001
	GroupB Group = 2 // 00000010
	GroupC Group = 4 // 00000100
	GroupD Group = 8 // 00001000
)

var GroupsRegistered = []Group{GroupA, GroupB, GroupC, GroupD}

func (g Group) GroupName() string {
	if g == GroupA {
		return "Group A"
	}

	if g == GroupB {
		return "Group B"
	}

	if g == GroupC {
		return "Group C"
	}

	if g == GroupD {
		return "Group D"
	}

	return "Unknown group"
}
