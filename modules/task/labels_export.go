// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package task

import (
	"fmt"
	"os"

	"code.gitea.io/sdk/gitea"
)

// LabelsExport save list of labels to disc
func LabelsExport(labels []*gitea.Label, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, label := range labels {
		if _, err := fmt.Fprintf(f, "#%s %s\n", label.Color, label.Name); err != nil {
			return err
		}
	}
	return nil
}
