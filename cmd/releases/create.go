// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package releases

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

// CmdReleaseCreate represents a sub command of Release to create release
var CmdReleaseCreate = cli.Command{
	Name:        "create",
	Aliases:     []string{"c"},
	Usage:       "Create a release",
	Description: `Create a release for a new or existing git tag`,
	ArgsUsage:   "[<tag>]",
	Action:      runReleaseCreate,
	Flags: append([]cli.Flag{
		&cli.StringFlag{
			Name:  "tag",
			Usage: "Tag name. If the tag does not exist yet, it will be created by Gitea",
		},
		&cli.StringFlag{
			Name:  "target",
			Usage: "Target branch name or commit hash. Defaults to the default branch of the repo",
		},
		&cli.StringFlag{
			Name:    "title",
			Aliases: []string{"t"},
			Usage:   "Release title",
		},
		&cli.StringFlag{
			Name:    "note",
			Aliases: []string{"n"},
			Usage:   "Release notes",
		},
		&cli.BoolFlag{
			Name:    "draft",
			Aliases: []string{"d"},
			Usage:   "Is a draft",
		},
		&cli.BoolFlag{
			Name:    "prerelease",
			Aliases: []string{"p"},
			Usage:   "Is a pre-release",
		},
		&cli.StringSliceFlag{
			Name:    "asset",
			Aliases: []string{"a"},
			Usage:   "Path to file attachment. Can be specified multiple times",
		},
	}, flags.AllDefaultFlags...),
}

func runReleaseCreate(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})

	tag := ctx.String("tag")
	if cmd.Args().Present() {
		if len(tag) != 0 {
			return fmt.Errorf("ambiguous arguments: provide tagname via --tag or argument, but not both")
		}
		tag = cmd.Args().First()
	}

	release, resp, err := ctx.Login.Client().CreateRelease(ctx.Owner, ctx.Repo, gitea.CreateReleaseOption{
		TagName:      tag,
		Target:       ctx.String("target"),
		Title:        ctx.String("title"),
		Note:         ctx.String("note"),
		IsDraft:      ctx.Bool("draft"),
		IsPrerelease: ctx.Bool("prerelease"),
	})

	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusConflict {
			return fmt.Errorf("There already is a release for this tag")
		}
		return err
	}

	for _, asset := range ctx.StringSlice("asset") {
		var file *os.File

		if file, err = os.Open(asset); err != nil {
			return err
		}

		filePath := filepath.Base(asset)

		if _, _, err = ctx.Login.Client().CreateReleaseAttachment(ctx.Owner, ctx.Repo, release.ID, file, filePath); err != nil {
			file.Close()
			return err
		}

		file.Close()
	}

	return nil
}
