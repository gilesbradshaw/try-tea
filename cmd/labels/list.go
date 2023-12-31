// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package labels

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/print"
	"code.gitea.io/tea/modules/task"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

// CmdLabelsList represents a sub command of labels to list labels
var CmdLabelsList = cli.Command{
	Name:        "list",
	Aliases:     []string{"ls"},
	Usage:       "List labels",
	Description: "List labels",
	ArgsUsage:   " ", // command does not accept arguments
	Action:      RunLabelsList,
	Flags: append([]cli.Flag{
		&cli.BoolFlag{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "Save all the labels as a file",
		},
		&flags.PaginationPageFlag,
		&flags.PaginationLimitFlag,
	}, flags.AllDefaultFlags...),
}

// RunLabelsList list labels.
func RunLabelsList(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})

	client := ctx.Login.Client()
	labels, _, err := client.ListRepoLabels(ctx.Owner, ctx.Repo, gitea.ListLabelsOptions{
		ListOptions: ctx.GetListOptions(),
	})
	if err != nil {
		return err
	}

	if ctx.IsSet("save") {
		return task.LabelsExport(labels, ctx.String("save"))
	}

	print.LabelsList(labels, ctx.Output)
	return nil
}
