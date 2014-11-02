package main

import (
	"testing"
	"encoding/json"
	"os"
)

func TestGetEventsByName(t *testing.T) {

	hookSource := "pxe-server:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks := hookMap{}
	hooks.Set(hookSource)
	name := "/pxe-server-askhaskdhaskdhashasidhsai"
	events := GetEventsByName(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "pxe-server-lykyukjypujkypukjypuojkyhylkh"
	events = GetEventsByName(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "pxe-serverlykyukjypujkypukjypuojkyhylkh"
	events = GetEventsByName(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "/pxe-server"
	events = GetEventsByName(hooks, name)

	if (events == nil) {
		t.Error("Name", name, "not found in", hookSource)
	}

	name = "/pxe"
	events = GetEventsByName(hooks, name)

	if (events != nil) {
		t.Error("Name", name, "shouldn't be found in", hookSource)
	}
}

func TestGetEventsByEnv(t *testing.T) {
	container := getContainerFromFile("test/container.json")
	hookSource := "pxe-server:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks := getHooks(hookSource)
	events := GetEventsByEnv(hooks, container.Config.Env)

	if (events != nil) {
		t.Error("Environment shouldn't be found", hookSource)
	}

	hookSource = "RABBITMQ_SERVICE_HOST=10.10.10.5:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks = getHooks(hookSource)
	events = GetEventsByEnv(hooks, container.Config.Env)

	if (events == nil) {
		t.Error("Environment not found", hookSource)
	}
}

func TestGetEvents(t *testing.T) {
	container := getContainerFromFile("test/container.json")
	hookSource := "test_123_hash_xyz:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks := getHooks(hookSource)
	events := GetEvents(hooks, &container)

	if (events == nil) {
		t.Error("Container not found by name", container.Name, "hooks: ", hookSource)
	}

	hookSource = "LIBVIRT_SERVICE_PORT=16509:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks = getHooks(hookSource)
	events = GetEvents(hooks, &container)

	if (events == nil) {
		t.Error("Container not found by env. Hooks: ", hookSource)
	}

	container = getContainerFromFile("test/container.json")
	hookSource = "teest_123_hash_xyz:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks = getHooks(hookSource)
	events = GetEvents(hooks, &container)

	if (events != nil) {
		t.Error("Container shouldn't be found by name", container.Name, "hooks: ", hookSource)
	}

	hookSource = "LIBVIRT_SERVICE_PORT=16599:start,restart:pipework:eth0:{{.ID}}:192.168.242.1/24"
	hooks = getHooks(hookSource)
	events = GetEvents(hooks, &container)

	if (events != nil) {
		t.Error("Container shouldn't be found by env. Hooks: ", hookSource)
	}
}

func getHooks(source string) hookMap {
	result := hookMap{}
	result.Set(source)
	return result
}

func getContainerFromFile(filename string) Container {

	result := Container{}
	reader, _ := os.Open(filename)
	decoder := json.NewDecoder(reader)
	decoder.Decode(&result)
	return result
}
