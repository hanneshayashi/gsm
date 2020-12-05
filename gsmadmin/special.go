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
	"sync"
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
