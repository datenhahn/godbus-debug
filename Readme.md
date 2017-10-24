This little program looks up a connection of linux Network Manager via DBUS and then trys
to get the details of this connection.

Hope this helps debug the issue #85.

* With *github.com/godbus/dbus* it hits this bug: https://github.com/godbus/dbus/issues/85
* With a different dbus lib: *launchpad.net/~jamesh/go-dbus/trunk* the lookup works
