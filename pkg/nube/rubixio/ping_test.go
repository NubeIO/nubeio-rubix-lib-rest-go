package rubixio

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"testing"
)

func TestPing(*testing.T) {

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

	client := New(restService)

	ping, res := client.Ping()

	fmt.Println(1111)
	fmt.Println(res.GetError())
	fmt.Println(1111)

	if res.GetError() != nil {
		fmt.Println("ping", res.GetError())
		return
	}

	fmt.Println("ping", ping.Ok)
}
