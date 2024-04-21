package routing

type RouteResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Result  []VrfResponse `json:"result"`
}
type VrfResponse struct {
	Vrf map[string]Routes `json:"vrfs"`
}

type Routes struct {
	Prefix                      map[string]Route `json:"routes"`
	AllRouteProgrammedKernel    bool             `json:"allRoutesProgrammedKernel"`
	RoutingDisabled             bool             `json:"routingDisabled"`
	AllRoutesProgrammedHardware bool             `json:"allRoutesProgrammedHardware"`
	DefaultRouteState           string           `json:"defaultRouteState"`
}
type Route struct {
	KernelProgrammed  bool                `json:"kernelProgrammed"`
	DirectlyConnected bool                `json:"directlyConnected"`
	RouteAction       string              `json:"routeAction"`
	RouteLeaked       bool                `json:"routeLeaked"`
	Vias              []map[string]string `json:"vias"`
	Metric            int32               `json:"metric"`
	Preference        int32               `json:"preference"`
	RouteType         string              `json:"routeType"`
}
