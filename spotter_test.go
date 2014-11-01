package main

import (
	"testing"
)

func TestGetEvents(t *testing.T) {

	hookSource := "pxe-server:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks := hookMap{}
	hooks.Set(hookSource)
	name := "/pxe-server-askhaskdhaskdhashasidhsai"
	events := GetEvents(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "pxe-server-lykyukjypujkypukjypuojkyhylkh"
	events = GetEvents(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "pxe-serverlykyukjypujkypukjypuojkyhylkh"
	events = GetEvents(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "/pxe-server"
	events = GetEvents(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "/pxe"
	events = GetEvents(hooks, name)

	if (events != nil) {
		t.Error("Name", name, "found in", hookSource)
	}
}
