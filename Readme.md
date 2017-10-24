This little program looks up a connection of linux Network Manager via DBUS and then trys
to get the details of this connection.

* With *godbus/dbus* it hits this bug: https://github.com/godbus/dbus/issues/85
* With a different dbus lib: jamesh/go-dbus the lookup works