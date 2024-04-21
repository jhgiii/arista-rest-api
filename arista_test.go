package arista_api

import (
	"arista_api/routing"
	"reflect"
	"testing"
)

func TestArista_SendConfig(t *testing.T) {
	type fields struct {
		Name     string
		Address  string
		Username string
		Password string
	}
	type args struct {
		config []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    CommandResult
		wantErr bool
	}{
		//todo
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Arista{
				Name:     tt.fields.Name,
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			got, err := a.SendConfig(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Arista.SendConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arista.SendConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArista_SendCommands(t *testing.T) {
	type fields struct {
		Name     string
		Address  string
		Username string
		Password string
	}
	type args struct {
		commands []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    CommandResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Arista{
				Name:     tt.fields.Name,
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			got, err := a.SendCommands(tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("Arista.SendCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arista.SendCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArista_GetRoutes(t *testing.T) {
	type fields struct {
		Name     string
		Address  string
		Username string
		Password string
	}
	type args struct {
		vrf string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    routing.Routes
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Arista{
				Name:     tt.fields.Name,
				Address:  tt.fields.Address,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			got, err := a.GetRoutes(tt.args.vrf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Arista.GetRoutes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Arista.GetRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
