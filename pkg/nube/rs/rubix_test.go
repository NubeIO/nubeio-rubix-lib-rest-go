package rs

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"testing"
)

func TestRubix(*testing.T) {

	restService := &rest.Service{}
	restService.Url = "192.168.15.19"
	restService.Port = 1616
	restOptions := &rest.Options{}
	restService.Options = restOptions
	restService = rest.New(restService)

	nubeProxy := &rest.NubeProxy{}
	nubeProxy.UseRubixProxy = true
	nubeProxy.RubixUsername = "admin"
	nubeProxy.RubixPassword = "N00BWires"
	nubeProxy.RubixProxyPath = nube.Services.RubixService.Proxy
	restService.NubeProxy = nubeProxy

	client := New(restService)

	_, res := client.Ping()
	fmt.Println(res.AsString())
}
