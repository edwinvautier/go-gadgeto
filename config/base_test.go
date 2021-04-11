package config

import (
	"os"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func Test_initBasicConfig(t *testing.T) {
	tests := []struct {
		name string
		want baseConfig
	}{
		{
			name: "test without viper",
			want: baseConfig{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initBasicConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initBasicConfig() = %v, want %v", got, tt.want)
			}
		})
	}

	workdir, err := os.Getwd()
	if err != nil {
		log.Info(err)
	}
	os.Create(workdir + "/.go-cli-config.yml")
	file, _ := os.OpenFile(workdir+"/.go-cli-config.yml", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if _, err := file.WriteString("package: helloworld"); err != nil {
		log.Fatal(err)
	}
	tests = []struct {
		name string
		want baseConfig
	}{
		{
			name: "test with viper",
			want: baseConfig{
				ProjectPath: workdir,
				PackagePath: "helloworld",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initBasicConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initBasicConfig() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Remove(workdir + "/.go-cli-config.yml")
}
