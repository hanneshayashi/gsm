/*
Copyright © 2020-2022 Hannes Hayashi

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

package gsmgmail

import (
	"github.com/hanneshayashi/gsm/gsmhelpers"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

// GetAutoForwardingSettings gets the auto-forwarding setting for the specified account.
func GetAutoForwardingSettings(userID, fields string) (*gmail.AutoForwarding, error) {
	srv := getUsersSettingsService()
	c := srv.GetAutoForwarding(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.AutoForwarding)
	return r, nil
}

// GetIMAPSettings gets IMAP settings.
func GetIMAPSettings(userID, fields string) (*gmail.ImapSettings, error) {
	srv := getUsersSettingsService()
	c := srv.GetImap(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.ImapSettings)
	return r, nil
}

// GetLanguageSettings gets language settings.
func GetLanguageSettings(userID, fields string) (*gmail.LanguageSettings, error) {
	srv := getUsersSettingsService()
	c := srv.GetLanguage(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.LanguageSettings)
	return r, nil
}

// GetPOPSettings gets POP settings.
func GetPOPSettings(userID, fields string) (*gmail.PopSettings, error) {
	srv := getUsersSettingsService()
	c := srv.GetPop(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.PopSettings)
	return r, nil
}

// GetVacationResponderSettings gets vacation responder settings.
func GetVacationResponderSettings(userID, fields string) (*gmail.VacationSettings, error) {
	srv := getUsersSettingsService()
	c := srv.GetVacation(userID)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.VacationSettings)
	return r, nil
}

// UpdateAutoForwardingSettings updates the auto-forwarding setting for the specified account.
// A verified forwarding address must be specified when auto-forwarding is enabled.
func UpdateAutoForwardingSettings(userID, fields string, autoForwarding *gmail.AutoForwarding) (*gmail.AutoForwarding, error) {
	srv := getUsersSettingsService()
	c := srv.UpdateAutoForwarding(userID, autoForwarding)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.AutoForwarding)
	return r, nil
}

// UpdateIMAPSettings updates IMAP settings.
func UpdateIMAPSettings(userID, fields string, imapSettings *gmail.ImapSettings) (*gmail.ImapSettings, error) {
	srv := getUsersSettingsService()
	c := srv.UpdateImap(userID, imapSettings)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.ImapSettings)
	return r, nil
}

// UpdateLanguageSettings updates language settings.
// If successful, the return object contains the displayLanguage that was saved for the user,which may differ from the value passed into the request.
// This is because the requested displayLanguage may not be directly supported by Gmail but have a close variant that is, and so the variant may be chosen and saved instead.
func UpdateLanguageSettings(userID, fields string, languageSetting *gmail.LanguageSettings) (*gmail.LanguageSettings, error) {
	srv := getUsersSettingsService()
	c := srv.UpdateLanguage(userID, languageSetting)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.LanguageSettings)
	return r, nil
}

// UpdatePOPSettings updates POP settings.
func UpdatePOPSettings(userID, fields string, popSettings *gmail.PopSettings) (*gmail.PopSettings, error) {
	srv := getUsersSettingsService()
	c := srv.UpdatePop(userID, popSettings)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.PopSettings)
	return r, nil
}

// UpdateVacationResponderSettings updates vacation responder settings.
func UpdateVacationResponderSettings(userID, fields string, vacationSettings *gmail.VacationSettings) (*gmail.VacationSettings, error) {
	srv := getUsersSettingsService()
	c := srv.UpdateVacation(userID, vacationSettings)
	if fields != "" {
		c.Fields(googleapi.Field(fields))
	}
	result, err := gsmhelpers.GetObjectRetry(gsmhelpers.FormatErrorKey(userID), func() (any, error) {
		return c.Do()
	})
	if err != nil {
		return nil, err
	}
	r, _ := result.(*gmail.VacationSettings)
	return r, nil
}
