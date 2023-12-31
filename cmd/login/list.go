// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package login

import (
	"code.gitea.io/tea/cmd/flags"
	"code.gitea.io/tea/modules/config"
	"code.gitea.io/tea/modules/print"

	"github.com/urfave/cli/v2"
)

// CmdLoginList represents to login a gitea server.
var CmdLoginList = cli.Command{
	Name:        "list",
	Aliases:     []string{"ls"},
	Usage:       "List Gitea logins",
	Description: `List Gitea logins`,
	ArgsUsage:   " ", // command does not accept arguments
	Action:      RunLoginList,
	Flags:       []cli.Flag{&flags.OutputFlag},
}

// RunLoginList list all logins
func RunLoginList(cmd *cli.Context) error {
	logins, err := config.GetLogins()
	if err != nil {
		return err
	}
	print.LoginsList(logins, cmd.String("output"))
	return nil
}
