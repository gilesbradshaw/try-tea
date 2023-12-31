// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package times

import (
	"fmt"

	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/context"
	"code.gitea.io/tea/modules/utils"

	"github.com/urfave/cli/v2"
)

// CmdTrackedTimesReset is a subcommand of CmdTrackedTimes, and
// clears all tracked times on an issue.
var CmdTrackedTimesReset = cli.Command{
	Name:      "reset",
	Usage:     "Reset tracked time on an issue",
	UsageText: "tea times reset <issue>",
	Action:    runTrackedTimesReset,
	Flags:     flags.LoginRepoFlags,
}

func runTrackedTimesReset(cmd *cli.Context) error {
	ctx := context.InitCommand(cmd)
	ctx.Ensure(context.CtxRequirement{RemoteRepo: true})
	client := ctx.Login.Client()

	if ctx.Args().Len() != 1 {
		return fmt.Errorf("No issue specified.\nUsage:\t%s", ctx.Command.UsageText)
	}

	issue, err := utils.ArgToIndex(ctx.Args().First())
	if err != nil {
		return err
	}

	_, err = client.ResetIssueTime(ctx.Owner, ctx.Repo, issue)
	return err
}
