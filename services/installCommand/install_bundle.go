package installCommand

import (
	"errors"
	"github.com/edwinvautier/go-gadgeto/config"
	"github.com/gobuffalo/packr/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

// InstallBundle install bundle from its name
func InstallBundle(name string) error {
	var installCmdConfig config.InstallCmdConfig
	config.InitInstallCmdConfig(&installCmdConfig)
	box := packr.New("Bundles", "../../bundles")

	if !bundleExists(name, box) {
		return errors.New(name + " bundle does not exist")
	}

	// execute templates
	if err := executeTemplates(name, installCmdConfig); err != nil {
		return err
	}

	// if it exists, execute the shell script it contains
	if err := executeInstallScript(box, name); err != nil {
		return err
	}
	// load & update config
	config.UpdateConfigAfterInstalling(name)
	log.Info("successfully installed ", name, " bundle")

	// Get link to the bundle documentation
	log.Info("You can go to the bundle documentation by going to the following link: https://github.com/edwinvautier/go-gadgeto/tree/main/bundles/", name)

	return nil
}

func bundleExists(name string, box *packr.Box) bool {
	files := box.List()

	for _, file := range files {
		bundleName := strings.Split(file, "/")[0]
		if bundleName == name {
			return true
		}
	}

	return false
}
