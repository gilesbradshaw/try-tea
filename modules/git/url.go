// Copyright 2019 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package git

import (
	"net/url"
	"regexp"
	"strings"
)

var protocolRe = regexp.MustCompile("^[a-zA-Z_+-]+://")

// URLParser represents a git URL parser
type URLParser struct{}

// Parse parses the git URL
func (p *URLParser) Parse(rawURL string) (u *url.URL, err error) {
	rawURL = strings.TrimSpace(rawURL)

	if !protocolRe.MatchString(rawURL) {
		// convert the weird git ssh url format to a canonical url:
		// git@gitea.com:gitea/tea -> ssh://git@gitea.com/gitea/tea
		if strings.Contains(rawURL, ":") &&
			// not a Windows path
			!strings.Contains(rawURL, "\\") {
			rawURL = "ssh://" + strings.Replace(rawURL, ":", "/", 1)
		} else if !strings.Contains(rawURL, "@") &&
			strings.Count(rawURL, "/") == 2 {
			// match cases like gitea.com/gitea/tea
			rawURL = "https://" + rawURL
		}
	}

	u, err = url.Parse(rawURL)
	if err != nil {
		return
	}

	if u.Scheme == "git+ssh" {
		u.Scheme = "ssh"
	}

	if strings.HasPrefix(u.Path, "//") {
		u.Path = strings.TrimPrefix(u.Path, "/")
	}

	// .git suffix is optional and breaks normalization
	u.Path = strings.TrimSuffix(u.Path, ".git")

	return
}

// ParseURL parses URL string and return URL struct
func ParseURL(rawURL string) (u *url.URL, err error) {
	p := &URLParser{}
	return p.Parse(rawURL)
}
