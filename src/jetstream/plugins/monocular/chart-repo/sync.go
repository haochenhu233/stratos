/*
Copyright (c) 2018 The Helm Authors

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

package chartrepo

import (
	"github.com/helm/monocular/cmd/chart-repo/foundationdb"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

//SyncCmd Sync a chart repository with Monocular
var SyncCmd = &cobra.Command{
	Use:   "sync [REPO NAME] [REPO URL]",
	Short: "add a new chart repository, and resync its charts periodically",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			log.Info("Need exactly two arguments: [REPO NAME] [REPO URL]")
			cmd.Help()
			return
		}
		foundationdb.Sync(cmd, args)
	},
}
