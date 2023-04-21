/*
Copyright © 2020-2023 Hannes Hayashi

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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const contactDelegateURL = `https://www.googleapis.com/admin/contacts/v1/users/%s/delegates`

// ContactDelegate represents a delegation to manage a user's contacts
type ContactDelegate struct {
	Email string `json:"email,omitempty"`
}

// CreateContactDelegate creates one or more delegates for a given user.
func CreateContactDelegate(parent, email string) (*ContactDelegate, error) {
	delegation, err := json.Marshal(ContactDelegate{Email: email})
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(delegation)
	req, err := http.NewRequest("POST", fmt.Sprintf(contactDelegateURL, parent), body)
	if err != nil {
		return nil, err
	}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	delegationR := &ContactDelegate{}
	err = json.Unmarshal(responseBody, delegationR)
	if err != nil {
		return nil, err
	}
	return delegationR, nil
}

// DeleteContactDelegate deletes a delegate from a given user.
func DeleteContactDelegate(parent, email string) (bool, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf(contactDelegateURL, parent)+fmt.Sprintf("/%s", email), nil)
	if err != nil {
		return false, err
	}
	r, err := client.Do(req)
	if err != nil {
		return false, err
	}
	if r.StatusCode != 200 {
		return false, nil
	}
	return true, nil
}

// ListContactDelegates lists the delegates of a given user.
func ListContactDelegates(parent string) ([]*ContactDelegate, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(contactDelegateURL, parent), nil)
	if err != nil {
		return nil, err
	}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	type listDelegateRespone struct {
		Delegates []*ContactDelegate
	}
	delegations := &listDelegateRespone{}
	err = json.Unmarshal(responseBody, delegations)
	if err != nil {
		return nil, err
	}
	return delegations.Delegates, nil
}
