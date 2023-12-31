// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package interact

import (
	"time"

	"code.gitea.io/tea/modules/config"
	"code.gitea.io/tea/modules/task"

	"code.gitea.io/sdk/gitea"
	"github.com/AlecAivazis/survey/v2"
)

// CreateMilestone interactively creates a milestone
func CreateMilestone(login *config.Login, owner, repo string) error {
	var title, description string
	var deadline *time.Time

	// owner, repo
	owner, repo, err := promptRepoSlug(owner, repo)
	if err != nil {
		return err
	}

	// title
	promptOpts := survey.WithValidator(survey.Required)
	promptI := &survey.Input{Message: "Milestone title:"}
	if err := survey.AskOne(promptI, &title, promptOpts); err != nil {
		return err
	}

	// description
	promptM := NewMultiline(Multiline{
		Message:   "Milestone description:",
		Syntax:    "md",
		UseEditor: config.GetPreferences().Editor,
	})
	if err := survey.AskOne(promptM, &description); err != nil {
		return err
	}

	// deadline
	if deadline, err = promptDatetime("Milestone deadline:"); err != nil {
		return err
	}

	return task.CreateMilestone(
		login,
		owner,
		repo,
		title,
		description,
		deadline,
		gitea.StateOpen)
}
