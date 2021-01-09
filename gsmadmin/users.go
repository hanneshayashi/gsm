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
	"github.com/hanneshayashi/gsm/gsmhelpers"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

// DeleteUser deletes a user.
func DeleteUser(userKey string) (bool, error) {
	srv := getUsersService()
	c := srv.Delete(userKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey), func() error {
		return c.Do()
	})
	return result, err
}

// GetUser retrieves a user.
func GetUser(userKey, fields, projection, customFieldMask, viewType string) (*admin.User, error) {
	srv := getUsersService()
	c := srv.Get(userKey)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if projection != "" {
		c.Projection(projection)
	}
	if customFieldMask != "" {
		c.CustomFieldMask(customFieldMask)
	}
	if viewType != "" {
		c.ViewType(viewType)
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.User)
	return r, nil
}

// InsertUser creates a user.
func InsertUser(user *admin.User, fields string) (*admin.User, error) {
	srv := getUsersService()
	c := srv.Insert(user)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(user.PrimaryEmail), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.User)
	return r, nil
}

func listUsers(c *admin.UsersListCall, ch chan *admin.User, errKey string) error {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return err
	}
	r, _ := result.(*admin.Users)
	for _, i := range r.Users {
		ch <- i
	}
	if r.NextPageToken != "" {
		c := c.PageToken(r.NextPageToken)
		err = listUsers(c, ch, errKey)
	}
	return err
}

// ListUsers retrieves a paginated list of either deleted users or all users in a domain.
func ListUsers(showDeleted bool, query, domain, customer, fields, projection, orderBy, sortOrder, viewType, customFieldMask string, cap int) (<-chan *admin.User, <-chan error) {
	srv := getUsersService()
	c := srv.List().Customer(customer).MaxResults(500)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	if query != "" {
		c = c.Query(query)
	}
	if projection != "" {
		c = c.Projection(projection)
	}
	if showDeleted {
		c = c.ShowDeleted("true")
	}
	if orderBy != "" {
		c = c.OrderBy(orderBy)
	}
	if sortOrder != "" {
		c = c.SortOrder(sortOrder)
	}
	if domain != "" {
		c = c.Domain(domain)
	}
	if viewType != "" {
		c = c.ViewType(viewType)
	}
	if customFieldMask != "" {
		c = c.CustomFieldMask(customFieldMask)
	}
	ch := make(chan *admin.User, cap)
	err := make(chan error, 1)
	go func() {
		e := listUsers(c, ch, gsmhelpers.FormatErrorKey(customer))
		if e != nil {
			err <- e
		}
		close(ch)
		close(err)
	}()
	gsmhelpers.Sleep()
	return ch, err
}

// MakeAdmin makes a user a super administrator.
func MakeAdmin(userKey string, status bool) (bool, error) {
	srv := getUsersService()
	makeAdmin := &admin.UserMakeAdmin{Status: status}
	if !makeAdmin.Status {
		makeAdmin.ForceSendFields = append(makeAdmin.ForceSendFields, "Status")
	}
	c := srv.MakeAdmin(userKey, makeAdmin)
	result, err := gsmhelpers.ActionRetry(userKey, func() error {
		return c.Do()
	})
	return result, err
}

// UpdateUser updates a user using patch semantics.
func UpdateUser(userKey, fields string, user *admin.User) (*admin.User, error) {
	srv := getUsersService()
	c := srv.Update(userKey, user)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userKey), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*admin.User)
	return r, nil
}

// SignOutUser signs a user out of all web and device sessions and reset their sign-in cookies.
// User will have to sign in by authenticating again.
func SignOutUser(userKey string) (bool, error) {
	srv := getUsersService()
	c := srv.SignOut(userKey)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey), func() error {
		return c.Do()
	})
	return result, err
}

// UndeletUser undeletes a deleted user.
func UndeletUser(userKey, orgUnitPath string) (bool, error) {
	srv := getUsersService()
	c := srv.Undelete(userKey, &admin.UserUndelete{OrgUnitPath: orgUnitPath})
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(userKey), func() error {
		return c.Do()
	})
	return result, err
}

// func hashPW(password string) string {
// 	h := sha1.New()
// 	h.Write([]byte(password))
// 	bs := h.Sum(nil)
// 	return fmt.Sprintf("%x\n", bs)
// }
