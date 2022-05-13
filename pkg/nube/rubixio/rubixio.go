package rubixio

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
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
	path := fmt.Sprintf(nube.Services.BacnetServer.PingPath)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

// GetInputs  get all inputs
func (inst *Client) GetInputs() (points *Inputs, response *rest.Reply) {
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(fmt.Sprintf("%s", Paths.Inputs.Path)).
		DoRequest()
	response = inst.Rest.RestResponse(req, &points)
	return
}

// UpdatePointValue write point value take an int from 1 to 100 (any value > 0 will write a DO to true for AO value from 0-100 = 0-10dc)
func (inst *Client) UpdatePointValue(ioNum string, value int) (point *Write, response *rest.Reply) {
	body := &Write{
		IoNum: ioNum,
		Value: value,
	}
	path := fmt.Sprintf("%s", Paths.Outputs.Path)
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(path).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &point)
	return
}

// UpdatePointValueBulk bulk write [{"io_num":"UO1", "value":100}, {"io_num":"UO2", "value":100}]
func (inst *Client) UpdatePointValueBulk(body []BulkWrite) (point *BulkResponse, response *rest.Reply) {
	path := fmt.Sprintf("%s", Paths.OutputsBulk.Path)
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(path).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &point)
	return
}
