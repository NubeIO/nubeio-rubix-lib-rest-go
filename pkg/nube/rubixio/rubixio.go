package rubixio

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
)

type BacnetClient struct {
	Rest *rest.Service
}

type Path struct {
	Path string
}

var Paths = struct {
	Ping    Path
	Outputs Path
	Inputs  Path
}{
	Ping:    Path{Path: "/api/system/ping"},
	Outputs: Path{Path: "/api/outputs"},
	Inputs:  Path{Path: "/api/inputs/all"},
}

// New returns a new instance of the nube common apis
func New(rest *rest.Service) *BacnetClient {
	bc := &BacnetClient{
		Rest: rest,
	}
	return bc
}

// Ping Ping server
func (inst *BacnetClient) Ping() (data *Ping, response *rest.ProxyResponse) {
	path := fmt.Sprintf(nube.Services.BacnetServer.PingPath)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &data)
	response.GetResponse()
	return
}

// GetInputs  get all inputs
func (inst *BacnetClient) GetInputs() (points *Inputs, response *rest.ProxyResponse) {
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(fmt.Sprintf("%s", Paths.Inputs.Path)).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &points)
	response.GetResponse()
	return
}

// UpdatePointValue write point value take an int from 1 to 100 (any value > 0 will write a DO to true for AO value from 0-100 = 0-10dc)
func (inst *BacnetClient) UpdatePointValue(ioNum string, value int) (point *Write, response *rest.ProxyResponse) {
	body := &Write{
		IoNum: ioNum,
		Value: value,
	}
	path := fmt.Sprintf("%s", Paths.Outputs.Path)
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(fmt.Sprintf(path)).
		SetBody(body).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &point)
	response.GetResponse()
	return
}
