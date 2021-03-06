package cmd

/*
Copyright © 2021 Edwin Vautier <edwin.vautier@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"github.com/edwinvautier/go-gadgeto/config"
	"github.com/edwinvautier/go-gadgeto/services/makeCommand"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// makeCmd represents the install command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "make is used to create new files, for example for models",
	Long:  `make is used to create new files, for example for models, it creates your model file after prompting you for fields`,
	Run: func(cmd *cobra.Command, args []string) {
		config.UpdateConfig()
		switch args[0] {
		case "model":
			if err := makeCommand.MakeModel(args[1]); err != nil {
				log.Fatal("make model ", args[1], " failed: ", err)
			}
		case "crud":
			if err := makeCommand.MakeCrud(args[1]); err != nil {
				log.Fatal("make crud for ", args[1], " failed: ", err)
			}
		case "fixtures":
			if err := makeCommand.MakeFixtures(args[1]); err != nil {
				log.Fatal("make fixtures for ", args[1], " failed: ", err)
			}
		case "tests":
			if err := makeCommand.MakeTests(args[1]); err != nil {
				log.Fatal("make tests for ", args[1], " failed: ", err)
			}
		default:
			log.Fatal(args[0], " is not a make command!")
		}
		config.UpdateConfig()	
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
