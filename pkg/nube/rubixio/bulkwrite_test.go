package rubixio

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"testing"
)

func TestBulkWrite(*testing.T) {

	restService := &rest.Service{}
	restService.Url = "192.168.15.10"
	restService.Port = 5001
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = false
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	nubeProxy.RubixProxyPath = nube.Services.BacnetServer.Proxy
	restService.NubeProxy = nubeProxy

	bacnetClient := New(restService)

	bulk := []BulkWrite{
		{
			IONum: "DO1",
			Value: 0,
		},
		{
			IONum: "DO2",
			Value: 0,
		},
	}

	inputs, resp := bacnetClient.UpdatePointValueBulk(bulk)
	if resp.GetError() != nil || inputs == nil {
		fmt.Println(resp.GetError())
		fmt.Println(resp.GetStatusCode())
		return
	}
	fmt.Println(inputs)

}
