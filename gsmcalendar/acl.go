/*
Package gsmcalendar implements the Calendar API
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
package gsmcalendar

import (
	"gsm/gsmhelpers"

	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
)

// DeleteACL deletes an access control rule.
func DeleteACL(calendarID, ruleID string) (bool, error) {
	srv := getACLService()
	c := srv.Delete(calendarID, ruleID)
	result, err := gsmhelpers.ActionRetry(gsmhelpers.FormatErrorKey(calendarID, ruleID), func() error {
		return c.Do()
	})
	return result, err
}

// GetACL returns an access control rule.
func GetACL(calendarID, ruleID, fields string) (*calendar.AclRule, error) {
	srv := getACLService()
	c := srv.Get(calendarID, ruleID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, ruleID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.AclRule)
	return r, nil
}

// InsertACL creates an acl.
func InsertACL(calendarID, fields string, acl *calendar.AclRule, sendNotifications bool) (*calendar.AclRule, error) {
	srv := getACLService()
	c := srv.Insert(calendarID, acl).SendNotifications(sendNotifications)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, acl.Scope.Value), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.AclRule)
	return r, nil
}

func makeListACLsCallAndAppend(c *calendar.AclListCall, acls []*calendar.AclRule, errKey string) ([]*calendar.AclRule, error) {
	result, err := gsmhelpers.GetObjectRetry(errKey, func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.Acl)
	for _, a := range r.Items {
		acls = append(acls, a)
	}
	if r.NextPageToken != "" {
		c.PageToken(r.NextPageToken)
		acls, err = makeListACLsCallAndAppend(c, acls, errKey)
	}
	return acls, err
}

// ListACLs returns the rules in the access control list for the calendar.
func ListACLs(calendarID, fields string, showDeleted bool) ([]*calendar.AclRule, error) {
	srv := getACLService()
	c := srv.List(calendarID).ShowDeleted(showDeleted)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	var acls []*calendar.AclRule
	acls, err := makeListACLsCallAndAppend(c, acls, gsmhelpers.FormatErrorKey(calendarID))
	return acls, err
}

// PatchACL updates an access control rule. This method supports patch semantics.
func PatchACL(calendarID, ruleID, fields string, aclRule *calendar.AclRule, sendNotifications bool) (*calendar.AclRule, error) {
	srv := getACLService()
	c := srv.Patch(calendarID, ruleID, aclRule).SendNotifications(sendNotifications)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(calendarID, ruleID), func() (interface{}, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*calendar.AclRule)
	return r, nil
}
