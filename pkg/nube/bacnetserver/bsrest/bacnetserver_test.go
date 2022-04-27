package bsrest

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"testing"
)

func TestBACnetRest(*testing.T) {

	restService := &rest.Service{}
	restService.Port = 1616
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = true
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	nubeProxy.RubixProxyPath = nube.Services.BacnetServer.Proxy
	restService.NubeProxy = nubeProxy

	bacnetClient := New(restService)

	ping, resp := bacnetClient.Ping()
	fmt.Println(ping)
	if resp.GetError() != nil || ping == nil {
		fmt.Println(resp.GetError())
		fmt.Println(resp.GetStatusCode())
		return
	}

	fmt.Println(ping.UpHour)
	pprint.PrintStrut(resp)

}
