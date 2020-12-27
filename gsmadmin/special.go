/*
Package gsmadmin implements the Admin SDK APIs
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
	"log"
	"strings"
	"sync"

	"github.com/hanneshayashi/gsm/gsmhelpers"
)

// GetUniqueUsersChannelRecursive returns a channel containing unique email addresses of all users inside the specified orgUnits and groups
func GetUniqueUsersChannelRecursive(orgUnits, groupEmails []string, threads int) (<-chan string, error) {
	wgOrgUnits := &sync.WaitGroup{}
	wgGroups := &sync.WaitGroup{}
	wgUnique := &sync.WaitGroup{}
	userKeys := make(chan string, threads)
	userKeysUnique := make(chan string, threads)
	done := make(map[string]struct{})
	wgOrgUnits.Add(1)
	go func() {
		for _, o := range orgUnits {
			us, err := ListUsers(false, fmt.Sprintf("orgUnitPath=%s", o), "", "my_customer", "users(primaryEmail),nextPageToken", "", "", "", "", "")
			if err != nil {
				log.Println(err)
			} else {
				for _, u := range us {
					userKeys <- u.PrimaryEmail
				}
			}
		}
		wgOrgUnits.Done()
	}()
	wgGroups.Add(1)
	go func() {
		for _, g := range groupEmails {
			mems, err := ListMembers(g, "", "members(email,type),nextPageToken", true)
			if err != nil {
				log.Println(err)
			} else {
				for _, m := range mems {
					if m.Type == "USER" {
						userKeys <- m.Email
					}
				}
			}
		}
		wgGroups.Done()
	}()
	wgUnique.Add(1)
	go func() {
		for uk := range userKeys {
			if _, found := done[uk]; !found {
				userKeysUnique <- uk
				done[uk] = struct{}{}
			}
		}
		wgUnique.Done()
	}()
	go func() {
		wgGroups.Wait()
		wgOrgUnits.Wait()
		close(userKeys)
		wgUnique.Wait()
		close(userKeysUnique)
	}()
	return userKeysUnique, nil
}

// GetMembersToSet compares the list of current members of a group to the specified emailAddresses.
// The function will return a list of members to be added and / or removed.
func GetMembersToSet(groupKey string, threads int, emailAddresses ...string) (<-chan string, <-chan string, error) {
	currentMembers, err := ListMembers(groupKey, "", "members(email)", false)
	membersToAdd := make(chan string, threads)
	membersToRemove := make(chan string, threads)
	if err != nil {
		return nil, nil, err
	}
	var cLower []string
	for _, cm := range currentMembers {
		cLower = append(cLower, strings.ToLower(cm.Email))
	}
	var nLower []string
	for _, e := range emailAddresses {
		nLower = append(nLower, strings.ToLower(e))
	}
	go func() {
		for _, n := range nLower {
			if !gsmhelpers.Contains(n, cLower) {
				membersToAdd <- n
			}
		}
		close(membersToAdd)
	}()
	go func() {
		for _, c := range cLower {
			if !gsmhelpers.Contains(c, nLower) {
				membersToRemove <- c
			}
		}
		close(membersToRemove)
	}()
	return membersToAdd, membersToRemove, nil
}
