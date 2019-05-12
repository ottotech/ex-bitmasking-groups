package groups

import (
	"html/template"
	"log"
)

func GroupAConfig() int {
	return GroupA
}

func GroupBConfig() int {
	return GroupB
}

func GroupCConfig() int {
	return GroupC
}

func GroupDConfig() int {
	return GroupD
}

func ObjBelongToGroup(g, config int) bool {
	groupIsRegistered := false
	for _, group := range RegisteredGroups {
		if group == g {
			groupIsRegistered = true
			break
		}
	}
	if !groupIsRegistered {
		log.Fatal("Groups passed to template tag ObjBelongToGroup is not registered!")
	}
	return BelongsToGroup(g, config)
}

var Fm = template.FuncMap{
	"GroupA":         GroupAConfig,
	"GroupB":         GroupBConfig,
	"GroupC":         GroupCConfig,
	"GroupD":         GroupDConfig,
	"BelongsToGroup": ObjBelongToGroup,
}
