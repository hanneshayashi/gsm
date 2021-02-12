/*
Package gsmadmin implements the Admin SDK APIs
Copyright © 2020-2021 Hannes Hayashi

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

// // TakeActionOnContactDelegate takes an action that affects a mobile device. For example, remotely wiping a device.
// func TakeActionOnContactDelegate(customerID, resourceID string, action *admin.ContactDelegateAction) (bool, error) {
// 	srv := getContactDelegatesService()
// 	c := srv.Action(customerID, resourceID, action)
// 	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID, resourceID), func() error {
// 		return c.Do()
// 	})
// 	return result, err
// }

// // DeleteContactDelegate removes a mobile device.
// func DeleteContactDelegate(customerID, resourceID string) (bool, error) {
// 	srv := getContactDelegatesService()
// 	c := srv.Delete(customerID, resourceID)
// 	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(customerID, resourceID), func() error {
// 		return c.Do()
// 	})
// 	return result, err
// }

// // GetContactDelegate retrieves a mobile device's properties.
// func GetContactDelegate(customerID, resourceID, fields, projection string) (*admin.ContactDelegate, error) {
// 	srv := getContactDelegatesService()
// 	c := srv.Get(customerID, resourceID)
// 	if fields != "" {
// 		c.Fields(googleapi.Field(fields))
// 	}
// 	if projection != "" {
// 		c.Projection(projection)
// 	}
// 	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(customerID, resourceID), func() (interface{}, error) {
// 		return c.Do()
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	r, _ := result.(*admin.ContactDelegate)
// 	return r, nil
// }

// func listContactDelegates(c *admin.ContactDelegatesListCall, ch chan *admin.ContactDelegate, errKey string) error {
// 	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
// 		return c.Do()
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	r, _ := result.(*admin.ContactDelegates)
// 	for _, i := range r.ContactDelegates {
// 		ch <- i
// 	}
// 	if r.NextPageToken != "" {
// 		c := c.PageToken(r.NextPageToken)
// 		err = listContactDelegates(c, ch, errKey)
// 	}
// 	return err
// }

// // ListContactDelegates retrieves a paginated list of all mobile devices for an account.
// func ListContactDelegates(customerID, query, fields, projection, orderBy, sortOrder string, cap int) (<-chan *admin.ContactDelegate, <-chan error) {
// 	srv := getContactDelegatesService()
// 	c := srv.List(customerID).MaxResults(100)
// 	if fields != "" {
// 		c.Fields(googleapi.Field(fields))
// 	}
// 	if query != "" {
// 		c = c.Query(query)
// 	}
// 	if projection != "" {
// 		c = c.Projection(projection)
// 	}
// 	if orderBy != "" {
// 		c = c.OrderBy(orderBy)
// 	}
// 	if sortOrder != "" {
// 		c = c.SortOrder(sortOrder)
// 	}
// 	ch := make(chan *admin.ContactDelegate, cap)
// 	err := make(chan error, 1)
// 	go func() {
// 		e := listContactDelegates(c, ch, gsmhelpers.FormatErrorKey(customerID))
// 		if e != nil {
// 			err <- e
// 		}
// 		close(ch)
// 		close(err)
// 	}()
// 	gsmhelpers.Sleep()
// 	return ch, err
// }
