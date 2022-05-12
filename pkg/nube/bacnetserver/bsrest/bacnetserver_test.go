package bsrest

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	pprint "github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/print"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
	"testing"
)

func truncateString(str string, num int) string {
	ret := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		ret = str[0:num] + ""
	}
	return ret
}

func nameIsNil(name string) string {
	if name == "" {
		uuid, _ := uuid.MakeUUID()
		return fmt.Sprintf("bacnet_%s", truncateString(uuid, 8))
	}
	return name
}

func TestBACnetRest(*testing.T) {

	restService := &rest.Service{}
	restService.Port = 1717
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

	ping, resp := bacnetClient.Ping()
	fmt.Println(ping)
	if resp.GetError() != nil || ping == nil {
		fmt.Println(resp.GetError())
		fmt.Println(resp.GetStatusCode())
		return
	}

	newPoint := &BacnetPoint{
		ObjectName:           nameIsNil(""),
		Description:          "na",
		ObjectType:           "analogOutput",
		Enable:               true,
		UseNextAvailableAddr: true,
		EventState:           "normal",
		Units:                "noUnits",
	}

	bacnetClient.AddPoint(newPoint)

	fmt.Println(ping.UpHour)
	pprint.PrintStrut(resp)

}
