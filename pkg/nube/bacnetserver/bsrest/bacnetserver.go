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

// Ping Ping server
func (inst *BacnetClient) Ping() (data *ServerPing, response *rest.Reply) {
	path := fmt.Sprintf(nube.Services.BacnetServer.PingPath)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(path).
		DoRequest()
	response = inst.Rest.RestResponse(req, &data)
	return
}

// GetPoints  get all
func (inst *BacnetClient) GetPoints() (points []*BacnetPoint, response *rest.Reply) {
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(fmt.Sprintf("%s/points", Paths.Points.Path)).
		DoRequest()
	response = inst.Rest.RestResponse(req, &points)
	return
}

//
// GetPoint get one by its uuid
func (inst *BacnetClient) GetPoint(uuid string) (point *BacnetPoint, response *rest.Reply) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.GET).
		SetPath(fmt.Sprintf(path)).
		DoRequest()
	response = inst.Rest.RestResponse(req, &point)
	return
}

// AddPoint add one object
func (inst *BacnetClient) AddPoint(body *BacnetPoint) (point *BacnetPoint, response *rest.Reply) {
	path := Paths.Points.Path
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(fmt.Sprintf(path)).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &point)
	return
}

// UpdatePoint update one object
func (inst *BacnetClient) UpdatePoint(uuid string, body *BacnetPoint) (point *BacnetPoint, response *rest.Reply) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.PATCH).
		SetPath(fmt.Sprintf(path)).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &point)
	return
}

// UpdatePointValue do a point write
func (inst *BacnetClient) UpdatePointValue(uuid string, body *BacnetPoint) (point *BacnetPoint, response *rest.Reply) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.PATCH).
		SetPath(fmt.Sprintf(path)).
		SetBody(body).
		DoRequest()
	response = inst.Rest.RestResponse(req, &point)
	return
}

// DeletePoint delete one by its uuid
func (inst *BacnetClient) DeletePoint(uuid string) (response *rest.Reply, notFound bool, deletedOk bool) {
	path := fmt.Sprintf("%s/uuid/%s", Paths.Points.Path, uuid)
	req := inst.Rest.
		SetMethod(rest.DELETE).
		SetPath(fmt.Sprintf(path)).
		DoRequest()
	response = inst.Rest.RestResponse(req, nil)
	response.GetStatus()
	if response.GetStatus() == 204 {
		deletedOk = true
	}
	return
}

type dropPoints struct {
	Ok          bool `json:"ok"`
	DeleteCount int  `json:"delete_count"`
}

//
// DropPoints delete all objects
func (inst *BacnetClient) DropPoints() (response *rest.Reply) {
	inst.Rest.LogFunc = rest.GetFunctionName(inst.DropPoints)
	points, response := inst.GetPoints()
	statusCode := response.GetStatus()
	dp := &dropPoints{}
	if rest.StatusCode2xx(statusCode) {
		count := 0
		for _, pnt := range points {
			count++
			inst.DeletePoint(pnt.UUID)
		}
		dp.Ok = true
		dp.DeleteCount = count
	}
	response.SetNewBody(dp)
	return
}
