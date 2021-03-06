// Copyright © 2018 munzli <manuel@monostream.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"os/exec"
	"github.com/spf13/cobra"
	"github.com/monostream/k8s-localflex-provisioner/driver/helper"
)

// unmountCmd represents the unmount command
var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "Removes a directory",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least 1 args")
		}
		return nil
	},
	Long: `Removes a directory`,
	Run: func(cmd *cobra.Command, args []string) {
		err := unMount(args[0])
		if err != nil {
			helper.Handle(helper.Response{
				Status:  helper.StatusFailure,
				Message: err.Error(),
			})
			return
		}

		helper.Handle(helper.Response{
			Status:  helper.StatusSuccess,
			Message: "successfully removed the volume",
		})
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)
}

func unMount(target string) error {
	mountCmd := "/bin/umount"

	cmd := exec.Command(mountCmd, target)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return errors.New(string(output[:]))
	}
	return nil
}
