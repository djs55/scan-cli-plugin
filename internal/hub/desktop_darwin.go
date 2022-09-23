package hub

import (
	"net"
	"os/user"
	"path"
)

func dialDesktopHTTPProxy() (net.Conn, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}
	socket := path.Join(user.HomeDir, "Library/Containers/com.docker.docker/Data/httpproxy.sock")
	return net.Dial("unix", socket)
}
