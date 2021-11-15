package systems

import (
	"gateway/package/rest"

	"github.com/k0kubun/pp"
)

//API base url

//TODO: move thes into config.yml configuration file
const (
	AIMURL  = "http://192.168.1.103:4322/aim/api/v1"
	GRMURL  = "http://localhost:4323/auth/api/v1"
	SOMAURL = "http://localhost:4321/"
)

var AimAPI *rest.Client
var GRMAPI *rest.Client

func Init() {
	pp.Printf("api client initalised")
	AimAPI = rest.New(AIMURL)
	GRMAPI = rest.New(AIMURL)
}
