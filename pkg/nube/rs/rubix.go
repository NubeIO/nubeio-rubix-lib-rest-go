package rs

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
)

type Client struct {
	Rest *rest.Service
}

type Path struct {
	Path string
}

var Paths = struct {
	Ping        Path
	Outputs     Path
	OutputsBulk Path
	Inputs      Path
}{
	Ping:        Path{Path: "/api/system/ping"},
	Outputs:     Path{Path: "/api/outputs"},
	OutputsBulk: Path{Path: "/api/outputs/bulk"},
	Inputs:      Path{Path: "/api/inputs/all"},
}

// New returns a new instance of the nube common apis
func New(rest *rest.Service) *Client {
	bc := &Client{
		Rest: rest,
	}
	return bc
}

// Ping Ping server
func (inst *Client) Ping() (data *Ping, response *rest.Reply) {
	path := fmt.Sprintf(Paths.Ping.Path)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}
