package groups

import (
	"html/template"
	"log"
)

func GroupAConfig() Group {
	return GroupA
}

func GroupBConfig() Group {
	return GroupB
}

func GroupCConfig() Group {
	return GroupC
}

func GroupDConfig() Group {
	return GroupD
}

func ObjBelongsToGroup(g Group, config int) bool {
	groupIsRegistered := false
	for _, group := range RegisteredGroups {
		if group == g {
			groupIsRegistered = true
			break
		}
	}
	if !groupIsRegistered {
		log.Fatal("Group passed to template tag `BelongsToGroup` is not registered!")
	}
	return BelongsToGroup(g, config)
}

var Fm = template.FuncMap{
	"GroupA":         GroupAConfig,
	"GroupB":         GroupBConfig,
	"GroupC":         GroupCConfig,
	"GroupD":         GroupDConfig,
	"BelongsToGroup": ObjBelongsToGroup,
}
