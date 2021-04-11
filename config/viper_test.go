package config

import (
	"os"
	"testing"
	log "github.com/sirupsen/logrus"
)

func Test_initViper(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "Test without config file",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initViper(); (err != nil) != tt.wantErr {
				t.Errorf("initViper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// Create config file
	workdir, err := os.Getwd()
	if err != nil {
		return
	}
	if _, err := os.Create(workdir + "/.go-cli-config.yml"); err != nil {
		log.Error(err)
		return
	}
	
	tests = []struct {
		name    string
		wantErr bool
	}{
		{
			name: "Test with config file",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := initViper(); (err != nil) != tt.wantErr {
				t.Errorf("initViper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	if err := os.Remove(workdir + "/.go-cli-config.yml"); err != nil {
		log.Error("couldnt remove fake config file")
		return
	}
}
