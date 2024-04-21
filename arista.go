package arista_api

import (
	"arista_api/routing"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Arista struct {
	Name     string
	Address  string
	Username string
	Password string
}

//This doesn't really need to be a method.
func apiconfig(fmt string) aristaAPICall {
	var aa aristaAPICall
	aa.Jsonrpc_version = "2.0"
	aa.Method = "runCmds"
	aa.Params.Version = 1
	aa.Params.Format = fmt
	aa.Id = "test"
	return aa
}
func (a *Arista) SendConfig(config []string) (CommandResult, error) {
	//Guard against incomplete command set
	if config[1] != "configure terminal" {
		temp_config := []string{"enable", "configure terminal"}
		temp_config = append(temp_config, config...)
		aa := apiconfig("json")
		aa.Params.Cmds = temp_config
		b, err := json.MarshalIndent(aa, "", "\t")
		if err != nil {
			return CommandResult{}, err
		}
		url := "https://" + a.Address + "/command-api"
		cr, err := apiCall(b, url, a.Username, a.Password)
		if err != nil {
			return CommandResult{}, nil
		}
		var c CommandResult
		_ = json.Unmarshal(cr, &c)

		return c, nil
	}
	aa := apiconfig("json")
	aa.Params.Cmds = config
	b, err := json.MarshalIndent(aa, "", "\t")
	if err != nil {
		return CommandResult{}, err
	}
	url := "https://" + a.Address + "/command-api"
	cr, err := apiCall(b, url, a.Username, a.Password)
	if err != nil {
		return CommandResult{}, nil
	}
	var c CommandResult
	_ = json.Unmarshal(cr, &c)
	return c, nil
}
func (a *Arista) SendCommands(commands []string) (CommandResult, error) {
	aa := aristaAPICall{
		Jsonrpc_version: "2.0",
		Method:          "runCmds",
		Params: aristaAPICallParameters{
			Version: 1,
			Cmds:    commands,
			Format:  "text",
		},
		Id: "test",
	}

	b, err := json.MarshalIndent(aa, "", "\t")
	if err != nil {
		return CommandResult{}, err
	}
	url := "https://" + a.Address + "/command-api"
	cr, err := apiCall(b, url, a.Username, a.Password)
	if err != nil {
		return CommandResult{}, nil
	}
	var c CommandResult
	_ = json.Unmarshal(cr, &c)
	return c, nil

}

func (a *Arista) GetRoutes(vrf string) (routing.Routes, error) {
	if vrf == "" {
		vrf = "default"
	}
	cmd := "show ip route vrf " + vrf
	aa := aristaAPICall{
		Jsonrpc_version: "2.0",
		Method:          "runCmds",
		Params: aristaAPICallParameters{
			Version: 1,
			Cmds:    []string{"enable", cmd},
			Format:  "json"},
		Id: "test",
	}
	b, err := json.MarshalIndent(aa, "", "\t")
	if err != nil {
		return routing.Routes{}, err
	}
	url := "https://" + a.Address + "/command-api"
	cr, err := apiCall(b, url, a.Username, a.Password)
	if err != nil {
		return routing.Routes{}, err
	}
	var r CommandResult
	_ = json.Unmarshal(cr, &r)
	routes := r.Result[1].Vrf[vrf]
	return routes, nil
}

type aristaAPICall struct {
	Jsonrpc_version string                  `json:"jsonrpc"`
	Method          string                  `json:"method"`
	Params          aristaAPICallParameters `json:"params"`
	Id              string                  `json:"id"`
}

type aristaAPICallParameters struct {
	Version int      `json:"version"`
	Cmds    []string `json:"cmds"`
	Format  string   `json:"format"`
}
type CommandResult struct {
	Jsonrpc string                 `json:"jsonrpc"`
	Id      string                 `json:"id"`
	Result  map[string]interface{} `json:"result"`
	Error   map[string]interface{} `json:"error"`
}

func apiCall(b []byte, url, username, password string) ([]byte, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.SetBasicAuth(username, password)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return res, nil
}
