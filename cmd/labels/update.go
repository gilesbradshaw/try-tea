// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package labels

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

// CmdLabelUpdate represents a sub command of labels to update label.
var CmdLabelUpdate = cli.Command{
	Name:        "update",
	Usage:       "Update a label",
	Description: `Update a label`,
	ArgsUsage:   " ", // command does not accept arguments
	Action:      runLabelUpdate,
	Flags: append([]cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "label id",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "label name",
		},
		&cli.StringFlag{
			Name:  "color",
			Usage: "label color value",
		},
		&cli.StringFlag{
			Name:  "description",
			Usage: "label description",
		},
	}, flags.AllDefaultFlags...),
}

func runLabelUpdate(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})

	id := ctx.Int64("id")
	var pName, pColor, pDescription *string
	name := ctx.String("name")
	if name != "" {
		pName = &name
	}

	color := ctx.String("color")
	if color != "" {
		pColor = &color
	}

	description := ctx.String("description")
	if description != "" {
		pDescription = &description
	}

	var err error
	_, _, err = ctx.Login.Client().EditLabel(ctx.Owner, ctx.Repo, id, gitea.EditLabelOption{
		Name:        pName,
		Color:       pColor,
		Description: pDescription,
	})

	if err != nil {
		return err
	}

	return nil
}
