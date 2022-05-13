package chirps

import (
	"github.com/NubeIO/nubeio-rubix-lib-rest-go/pkg/rest"
)

//admin
//N00BWAN

type ChirpClient struct {
	Rest     *rest.Service
	Email    string `json:"email"`
	Password string `json:"password"`
	OrgID    int    `json:"org_id"`
	Limit    int    `json:"limit"`
}

type Path struct {
	Path string
}

var Paths = struct {
	Login Path
}{
	Login: Path{Path: "/api/internal/login"},
}

var limit = 10
var orgID = 1

// New returns a new instance of the nube common apis
func New(bc *ChirpClient) *ChirpClient {
	return bc
}


type TokenResponse struct {
	JWT string `json:"jwt"`
}

type Token struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetToken get cs token
func (inst *ChirpClient) GetToken(token *Token) (data *TokenResponse, response *rest.Reply) {
	req := inst.Rest.
		SetMethod(rest.POST).
		SetPath(Paths.Login.Path).
		SetBody(token).
		DoRequest()
	response = inst.Rest.RestResponse(req, data)
	return
}

// GetDevices get all
func (inst *ChirpClient) GetDevices(token string) (data *Devices, response *rest.Reply) {
	//path := fmt.Sprintf("/api/devices?limit=%d&organizationID=%d", limit, orgID)
	inst.Rest.Options.Headers = map[string]interface{}{"Grpc-Metadata-Authorization": token}
	res := inst.Rest.DoRequest()
	response = inst.Rest.RestResponse(res, &data)
	return
}
