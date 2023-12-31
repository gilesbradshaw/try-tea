// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package cmd

import (
	"code.gitea.io/tea/cmd/times"
	"github.com/urfave/cli/v2"
)

// CmdTrackedTimes represents the command to operate repositories' times.
var CmdTrackedTimes = cli.Command{
	Name:     "times",
	Aliases:  []string{"time", "t"},
	Category: catEntities,
	Usage:    "Operate on tracked times of a repository's issues & pulls",
	Description: `Operate on tracked times of a repository's issues & pulls.
		 Depending on your permissions on the repository, only your own tracked
		 times might be listed.`,
	ArgsUsage: "[username | #issue]",
	Action:    times.RunTimesList,
	Subcommands: []*cli.Command{
		&times.CmdTrackedTimesAdd,
		&times.CmdTrackedTimesDelete,
		&times.CmdTrackedTimesReset,
		&times.CmdTrackedTimesList,
	},
	Flags: times.CmdTrackedTimesList.Flags,
}
