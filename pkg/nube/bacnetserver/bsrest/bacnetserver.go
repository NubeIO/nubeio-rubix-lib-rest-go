package bsrest

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
	Points Path
}{
	Points: Path{Path: "/api/bacnet/points"},
}

// New returns a new instance of the nube common apis
func New(rest *rest.Service) *BacnetClient {

	bc := &BacnetClient{
		Rest: rest,
	}

	return bc
}

func (inst *BacnetClient) builder(logFunc interface{}, path string) *rest.Service {
	//get token if using proxy
	if inst.Rest.NubeProxy.UseRubixProxy {
		r := inst.Rest.GetToken()
		inst.Rest.Options.Headers = map[string]interface{}{"Authorization": r.Token}
	}
	inst.Rest.Path = path
	inst.Rest.LogFunc = rest.GetFunctionName(logFunc)
	return inst.Rest

}

// Ping Ping server
func (inst *BacnetClient) Ping() (data *ServerPing, response *rest.ProxyResponse) {
	path := fmt.Sprintf(nube.Services.BacnetServer.PingPath)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &data)
	response.GetResponse()
	return
}

// GetPoints  get all
func (inst *BacnetClient) GetPoints() (points []*BacnetPoint, response *rest.ProxyResponse) {
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(fmt.Sprintf("%s/points", Paths.Points.Path)).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &points)
	response.GetResponse()
	return
}

//
// GetPoint get one by its uuid
func (inst *BacnetClient) GetPoint(uuid string) (point *BacnetPoint, response *rest.ProxyResponse) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(fmt.Sprintf(path)).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &point)
	response.GetResponse()
	return
}

// AddPoint add one object
func (inst *BacnetClient) AddPoint(body *BacnetPoint) (point *BacnetPoint, response *rest.ProxyResponse) {
	path := Paths.Points.Path
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(fmt.Sprintf(path)).
		SetBody(body).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &point)
	response.GetResponse()
	return
}

// UpdatePoint update one object
func (inst *BacnetClient) UpdatePoint(uuid string, body *BacnetPoint) (point *BacnetPoint, response *rest.ProxyResponse) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.PATCH).
		SetPath(fmt.Sprintf(path)).
		SetBody(body).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &point)
	response.GetResponse()
	return
}

// UpdatePointValue do a point write
func (inst *BacnetClient) UpdatePointValue(uuid string, body *BacnetPoint) (point *BacnetPoint, response *rest.ProxyResponse) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.PATCH).
		SetPath(fmt.Sprintf(path)).
		SetBody(body).
		DoRequest()
	response = inst.Rest.BuildResponse(req, &point)
	response.GetResponse()
	return
}

// DeletePoint delete one by its uuid
func (inst *BacnetClient) DeletePoint(uuid string) (response *rest.ProxyResponse, notFound bool, deletedOk bool) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.DELETE).
		SetPath(fmt.Sprintf(path)).
		DoRequest()
	response = inst.Rest.BuildResponse(req, nil)
	response.GetResponse()
	if response.GetStatusCode() == 204 {
		deletedOk = true
	}
	return
}

//
// DropPoints delete all objects
func (inst *BacnetClient) DropPoints() (response *rest.ProxyResponse) {
	inst.Rest.LogFunc = rest.GetFunctionName(inst.DropPoints)
	points, response := inst.GetPoints()
	statusCode := response.GetStatusCode()
	if rest.StatusCode2xx(statusCode) {
		count := 0
		for _, pnt := range points {
			count++
			inst.DeletePoint(pnt.UUID)
		}
		response.Response.BodyType = rest.TypeString
		response.Response.Body = ""
		response.Response.Message = fmt.Sprintf("points delete %d", count)
	}
	return
}
