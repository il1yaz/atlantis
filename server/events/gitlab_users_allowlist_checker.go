package events

import "strings"

type GitlabUsersAllowlistChecker struct {
	users []string
}

func NewGitlabUsersAllowlistChecker(allowlist string) *GitlabUsersAllowlistChecker {
	users := strings.Split(allowlist, ",")
	return &GitlabUsersAllowlistChecker{
		users: users,
	}
}

func (u *GitlabUsersAllowlistChecker) IsAllowListed(userName string) bool {
	if len(u.users) == 1 && u.users[0] == "" {
		return true
	}
	for _, user := range u.users {
		if userName == user {
			return true
		}
	}
	return false
}
