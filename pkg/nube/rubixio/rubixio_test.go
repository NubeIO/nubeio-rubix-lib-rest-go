package rubixio

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"testing"
)

func TestBACnetRest(*testing.T) {

	restService := &rest.Service{}
	restService.Url = "192.168.15.194"
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

	inputs, resp := bacnetClient.GetInputs()
	if resp.GetError() != nil || inputs == nil {
		fmt.Println(resp.GetError())
		fmt.Println(resp.GetStatusCode())
		return
	}

	//for i, input := range inputs {
	//	fmt.Println(i, input)
	//}

	fmt.Println(inputs)
	//pprint.PrintStrut(resp)
	//
	//write, resp := bacnetClient.UpdatePointValue("UO3", 22)
	//if resp.GetError() != nil || inputs == nil {
	//	fmt.Println(resp.GetError())
	//	fmt.Println(resp.GetStatusCode())
	//	return
	//}
	//
	//fmt.Println(write)
	//pprint.PrintStrut(resp)

}
