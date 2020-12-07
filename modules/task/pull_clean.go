// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package task

import (
	"fmt"

	"code.gitea.io/sdk/gitea"
	"code.gitea.io/tea/modules/config"
	local_git "code.gitea.io/tea/modules/git"

	git_config "github.com/go-git/go-git/v5/config"
)

// PullClean deletes local & remote feature-branches for a closed pull
func PullClean(login *config.Login, repoOwner, repoName string, index int64, ignoreSHA bool) error {
	client := login.Client()

	repo, _, err := client.GetRepo(repoOwner, repoName)
	defaultBranch := repo.DefaultBranch
	if len(defaultBranch) == 0 {
		defaultBranch = "master"
	}

	// fetch PR source-repo & -branch from gitea
	pr, _, err := client.GetPullRequest(repoOwner, repoName, index)
	if err != nil {
		return err
	}
	if pr.State == gitea.StateOpen {
		return fmt.Errorf("PR is still open, won't delete branches")
	}

	// IDEA: abort if PR.Head.Repository.CloneURL does not match login.URL?

	r, err := local_git.RepoForWorkdir()
	if err != nil {
		return err
	}

	// find a branch with matching sha or name, that has a remote matching the repo url
	var branch *git_config.Branch
	if ignoreSHA {
		branch, err = r.TeaFindBranchByName(pr.Head.Ref, pr.Head.Repository.CloneURL)
	} else {
		branch, err = r.TeaFindBranchBySha(pr.Head.Sha, pr.Head.Repository.CloneURL)
	}
	if err != nil {
		return err
	}
	if branch == nil {
		if ignoreSHA {
			return fmt.Errorf("Remote branch %s not found in local repo", pr.Head.Ref)
		}
		return fmt.Errorf(`Remote branch %s not found in local repo.
Either you don't track this PR, or the local branch has diverged from the remote.
If you still want to continue & are sure you don't loose any important commits,
call me again with the --ignore-sha flag`, pr.Head.Ref)
	}

	// prepare deletion of local branch:
	headRef, err := r.Head()
	if err != nil {
		return err
	}
	if headRef.Name().Short() == branch.Name {
		fmt.Printf("Checking out '%s' to delete local branch '%s'\n", defaultBranch, branch.Name)
		if err = r.TeaCheckout(defaultBranch); err != nil {
			return err
		}
	}

	// remove local & remote branch
	fmt.Printf("Deleting local branch %s and remote branch %s\n", branch.Name, pr.Head.Ref)
	url, err := r.TeaRemoteURL(branch.Remote)
	if err != nil {
		return err
	}
	auth, err := local_git.GetAuthForURL(url, login.User, login.SSHKey)
	if err != nil {
		return err
	}
	return r.TeaDeleteBranch(branch, pr.Head.Ref, auth)
}