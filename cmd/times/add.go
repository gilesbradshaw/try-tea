// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package times

import (
	"fmt"
	"strings"
	"time"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/utils"

	"code.gitea.io/sdk/gitea"
	"github.com/urfave/cli/v2"
)

// CmdTrackedTimesAdd represents a sub command of times to add time to an issue
var CmdTrackedTimesAdd = cli.Command{
	Name:      "add",
	Aliases:   []string{"a"},
	Usage:     "Track spent time on an issue",
	UsageText: "tea times add <issue> <duration>",
	Description: `Track spent time on an issue
	 Example:
		tea times add 1 1h25m
	`,
	Action: runTrackedTimesAdd,
	Flags:  flags.LoginRepoFlags,
}

func runTrackedTimesAdd(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})

	if ctx.Args().Len() < 2 {
		return fmt.Errorf("No issue or duration specified.\nUsage:\t%s", ctx.Command.UsageText)
	}

	issue, err := utils.ArgToIndex(ctx.Args().First())
	if err != nil {
		return err
	}

	duration, err := time.ParseDuration(strings.Join(ctx.Args().Tail(), ""))
	if err != nil {
		return err
	}

	_, _, err = ctx.Login.Client().AddTime(ctx.Owner, ctx.Repo, issue, gitea.AddTimeOption{
		Time: int64(duration.Seconds()),
	})
	return err
}
