/*
Copyright © 2020 Hannes Hayashi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package gsmadmin

import (
	"fmt"
	"strings"
	"sync"

	"github.com/hanneshayashi/gsm/gsmhelpers"
)

// GetUniqueUsersChannelRecursive returns a channel containing unique email addresses of all users inside the specified orgUnits and groups
func GetUniqueUsersChannelRecursive(orgUnits, groupEmails []string, threads int) (<-chan string, <-chan error) {
	wgOrgUnits := &sync.WaitGroup{}
	wgGroups := &sync.WaitGroup{}
	wgUnique := &sync.WaitGroup{}
	userKeys := make(chan string, threads)
	userKeysUnique := make(chan string, threads)
	done := make(map[string]struct{})
	errChan := make(chan error, 2)
	wgOrgUnits.Go(func() {
		for i := range orgUnits {
			var iterErr error
			for u, err := range ListUsers(false, fmt.Sprintf("orgUnitPath=%s", orgUnits[i]), "", "my_customer", "users(primaryEmail),nextPageToken", "", "", "", "", "") {
				if err != nil {
					iterErr = err
					break
				}
				userKeys <- u.PrimaryEmail
			}
			if iterErr != nil {
				errChan <- iterErr
				break
			}
		}
	})
	wgGroups.Go(func() {
		for i := range groupEmails {
			var iterErr error
			for m, err := range ListMembers(groupEmails[i], "", "members(email,type),nextPageToken", true) {
				if err != nil {
					iterErr = err
					break
				}
				if m.Type == "USER" {
					userKeys <- m.Email
				}
			}
			if iterErr != nil {
				errChan <- iterErr
				break
			}
		}
	})
	wgUnique.Go(func() {
		for uk := range userKeys {
			if _, found := done[uk]; !found {
				userKeysUnique <- uk
				done[uk] = struct{}{}
			}
		}
	})
	go func() {
		wgGroups.Wait()
		wgOrgUnits.Wait()
		close(userKeys)
		wgUnique.Wait()
		close(userKeysUnique)
		close(errChan)
	}()
	return userKeysUnique, errChan
}

// GetMembersToSet compares the list of current members of a group to the specified emailAddresses.
// The function will return a list of members to be added and / or removed.
func GetMembersToSet(groupKey string, threads int, emailAddresses ...string) (<-chan string, <-chan string, error) {
	var cLower []string
	for cm, err := range ListMembers(groupKey, "", "members(email)", false) {
		if err != nil {
			return nil, nil, err
		}
		cLower = append(cLower, strings.ToLower(cm.Email))
	}
	membersToAdd := make(chan string, threads)
	membersToRemove := make(chan string, threads)
	var nLower []string
	for i := range emailAddresses {
		nLower = append(nLower, strings.ToLower(emailAddresses[i]))
	}
	go func() {
		for i := range nLower {
			if !gsmhelpers.Contains(nLower[i], cLower) {
				membersToAdd <- nLower[i]
			}
		}
		close(membersToAdd)
	}()
	go func() {
		for i := range cLower {
			if !gsmhelpers.Contains(cLower[i], nLower) {
				membersToRemove <- cLower[i]
			}
		}
		close(membersToRemove)
	}()
	return membersToAdd, membersToRemove, nil
}
