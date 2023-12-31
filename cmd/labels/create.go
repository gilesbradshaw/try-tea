// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package labels

import (
	"bufio"
	"log"
	"os"
	"strings"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

// CmdLabelCreate represents a sub command of labels to create label.
var CmdLabelCreate = cli.Command{
	Name:        "create",
	Aliases:     []string{"c"},
	Usage:       "Create a label",
	Description: `Create a label`,
	ArgsUsage:   " ", // command does not accept arguments
	Action:      runLabelCreate,
	Flags: append([]cli.Flag{
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
		&cli.StringFlag{
			Name:  "file",
			Usage: "indicate a label file",
		},
	}, flags.AllDefaultFlags...),
}

func runLabelCreate(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})

	labelFile := ctx.String("file")
	var err error
	if len(labelFile) == 0 {
		_, _, err = ctx.Login.Client().CreateLabel(ctx.Owner, ctx.Repo, gitea.CreateLabelOption{
			Name:        ctx.String("name"),
			Color:       ctx.String("color"),
			Description: ctx.String("description"),
		})
	} else {
		f, err := os.Open(labelFile)
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		var i = 1
		for scanner.Scan() {
			line := scanner.Text()
			color, name, description := splitLabelLine(line)
			if color == "" || name == "" {
				log.Printf("Line %d ignored because lack of enough fields: %s\n", i, line)
			} else {
				_, _, err = ctx.Login.Client().CreateLabel(ctx.Owner, ctx.Repo, gitea.CreateLabelOption{
					Name:        name,
					Color:       color,
					Description: description,
				})
			}

			i++
		}
	}

	return err
}

func splitLabelLine(line string) (string, string, string) {
	fields := strings.SplitN(line, ";", 2)
	var color, name, description string
	if len(fields) < 1 {
		return "", "", ""
	} else if len(fields) >= 2 {
		description = strings.TrimSpace(fields[1])
	}
	fields = strings.Fields(fields[0])
	if len(fields) <= 0 {
		return "", "", ""
	}
	color = fields[0]
	if len(fields) == 2 {
		name = fields[1]
	} else if len(fields) > 2 {
		name = strings.Join(fields[1:], " ")
	}
	return color, name, description
}
