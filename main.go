// This little program looks up a connection of linux Network Manager via DBUS and then trys
// to get the details of this connection.
// With godbus/dbus it hits this bug: https://github.com/godbus/dbus/issues/85
// With a different dbus lib: jamesh/go-dbus the lookup works

package main

import (
	"log"
	"github.com/godbus/dbus"
	jameshdbus "launchpad.net/~jamesh/go-dbus/trunk"
	"os"
	"encoding/json"
)

func main() {

	conns, err := GetConnections()

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	if len(conns) < 1 {
		log.Print("Didn't find any networkmanager connections")
		os.Exit(1)
	}

	println("Found conn: " + conns[0])

	println("\nLooking up details with \"launchpad.net/~jamesh/go-dbus/trunk\"")
	details, err := GetConnectionDetailJameshDbus(conns[0])

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	jameshJson, err := json.Marshal(details)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	println(string(jameshJson))

	println("\nLooking up details with \"github.com/godbus/dbus\"")
	details2, err2 := GetConnectionDetailGoDbus(conns[0])

	if err2 != nil {
		log.Print(err)
		os.Exit(1)
	}

	godbusJson, err := json.Marshal(details2)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	println(string(godbusJson))
}

func GetConnections() ([]dbus.ObjectPath, error) {

	var connectionObjectPaths []dbus.ObjectPath

	bus, err := dbus.SystemBus()

	if err != nil {
		log.Print(err.Error())
		return connectionObjectPaths, err
	}

	object := bus.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager/Settings")

	err = object.Call("org.freedesktop.NetworkManager.Settings.ListConnections", 0).Store(&connectionObjectPaths)

	if err != nil {
		log.Print(err.Error())
		return connectionObjectPaths, err
	}

	return connectionObjectPaths, nil
}

func GetConnectionDetailJameshDbus(connectionObjectPath dbus.ObjectPath) (map[string]map[string]jameshdbus.Variant, error) {
	details := make(map[string]map[string]jameshdbus.Variant)

	busConn, err := jameshdbus.Connect(jameshdbus.SystemBus)

	if err != nil {
		log.Print(err.Error())
		return details, err
	}

	obj := busConn.Object("org.freedesktop.NetworkManager", jameshdbus.ObjectPath(connectionObjectPath))

	msg, err := obj.Call("org.freedesktop.NetworkManager.Settings.Connection", "GetSettings")

	if err != nil {
		log.Print(err.Error())
		return details, err
	}

	if err = msg.Args(&details); err != nil {
		return details, err
	}

	return details, nil

}

func GetConnectionDetailGoDbus(connectionObjectPath dbus.ObjectPath) (map[string]map[string]dbus.Variant, error) {
	details := make(map[string]map[string]dbus.Variant)

	bus, err := dbus.SystemBus()

	if err != nil {
		log.Print(err.Error())
		return details, err
	}

	object := bus.Object("org.freedesktop.NetworkManager", connectionObjectPath)

	err = object.Call("org.freedesktop.NetworkManager.Settings.Connection.GetSettings", 0).Store(&details)

	if err != nil {
		log.Print(err.Error())
		return details, err
	}

	return details, nil

}
