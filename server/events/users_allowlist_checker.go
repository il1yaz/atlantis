package events

import "strings"

type UsersAllowlistChecker struct {
	users []string
}

func NewUsersAllowlistChecker(allowlist string) *UsersAllowlistChecker {
	users := strings.Split(allowlist, ",")
	return &UsersAllowlistChecker{
		users: users,
	}
}

func (u *UsersAllowlistChecker) IsAllowListed(userName string) bool {
	for _, user := range u.users {
		if userName == user {
			return true
		}
	}
	return false
}
