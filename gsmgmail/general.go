/*
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
package gsmgmail

import "github.com/hanneshayashi/gsm/gsmhelpers"

// SecurityModeIsValid checks if the given SecurityMode is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/PopSettings#securitymode
func SecurityModeIsValid(SecurityMode string) bool {
	switch SecurityMode {
	case "NONE", "SSL", "STARTTLS":
		return true
	}
	return false
}

// SizeComparisonIsValid checks if the given SizeComparison is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/PopSettings#sizecomparison
func SizeComparisonIsValid(sizeComparison string) bool {
	switch sizeComparison {
	case "UNSPECIFIED", "SMALLER", "LARGER":
		return true
	}
	return false
}

// AccessWindowIsValid checks if the given AccessWindow is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/PopSettings#accesswindow
func AccessWindowIsValid(accessWindow string) bool {
	switch accessWindow {
	case "DISABLED", "FROM_NOW_ON", "ALL_MAIL":
		return true
	}
	return false
}

// ExpungeBehaviourIsValid checks if the given ExpungeBehaviour is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/ImapSettings#expungebehavior
func ExpungeBehaviourIsValid(expungebehaviour string) bool {
	switch expungebehaviour {
	case "ARCHIVE", "TRASH", "DELETE_FOREVER":
		return true
	}
	return false
}

// DispositionIsValid checks if the given Disposition is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/AutoForwarding#disposition
func DispositionIsValid(disposition string) bool {
	switch disposition {
	case "LEAVE_IN_INBOX", "ARCHIVE", "TRASH", "MARK_READ":
		return true
	}
	return false
}

// InternalDateSourceIsValid checks if the given internalDateSource is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/InternalDateSource
func InternalDateSourceIsValid(internalDateSource string) bool {
	switch internalDateSource {
	case "RECEIVED_TIME", "DATE_HEADER":
		return true
	}
	return false
}

// FormatIsValid checks if the given Format is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/Format
func FormatIsValid(format string) bool {
	switch format {
	case "MINIMAL", "FULL", "RAW", "METADATA":
		return true
	}
	return false
}

// HistoryTypeIsValid checks if the given HistoryType is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/users.history/list#historytype
func HistoryTypeIsValid(historyType string) bool {
	switch historyType {
	case "MESSAGE_ADDED", "MESSAGE_DELETED", "LABEL_ADDED", "LABEL_REMOVED":
		return true
	}
	return false
}

// MessageVisibilityIsValid checks if the given MessageVisibility is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/users.labels#MessageListVisibility
func MessageVisibilityIsValid(messageVisibility string) bool {
	switch messageVisibility {
	case "SHOW", "HIDE":
		return true
	}
	return false
}

// LabelListVisibilityIsValid checks if the given LabelListVisibility is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/users.labels#MessageListVisibility
func LabelListVisibilityIsValid(labelListVisibility string) bool {
	switch labelListVisibility {
	case "LABEL_SHOW", "LABEL_SHOW_IF_UNREAD", "LABEL_HIDE":
		return true
	}
	return false
}

// ColorIsValid checks if the given Color is valid
// see: https://developers.google.com/gmail/api/reference/rest/v1/users.labels#MessageListVisibility
func ColorIsValid(color string) bool {
	validColors := []string{
		"",
		"#000000",
		"#434343",
		"#666666",
		"#999999",
		"#cccccc",
		"#efefef",
		"#f3f3f3",
		"#ffffff",
		"#fb4c2f",
		"#ffad47",
		"#fad165",
		"#16a766",
		"#43d692",
		"#4a86e8",
		"#a479e2",
		"#f691b3",
		"#f6c5be",
		"#ffe6c7",
		"#fef1d1",
		"#b9e4d0",
		"#c6f3de",
		"#c9daf8",
		"#e4d7f5",
		"#fcdee8",
		"#efa093",
		"#ffd6a2",
		"#fce8b3",
		"#89d3b2",
		"#a0eac9",
		"#a4c2f4",
		"#d0bcf1",
		"#fbc8d9",
		"#e66550",
		"#ffbc6b",
		"#fcda83",
		"#44b984",
		"#68dfa9",
		"#6d9eeb",
		"#b694e8",
		"#f7a7c0",
		"#cc3a21",
		"#eaa041",
		"#f2c960",
		"#149e60",
		"#3dc789",
		"#3c78d8",
		"#8e63ce",
		"#e07798",
		"#ac2b16",
		"#cf8933",
		"#d5ae49",
		"#0b804b",
		"#2a9c68",
		"#285bac",
		"#653e9b",
		"#b65775",
		"#822111",
		"#a46a21",
		"#aa8831",
		"#076239",
		"#1a764d",
		"#1c4587",
		"#41236d",
		"#83334c",
		"#464646",
		"#e7e7e7",
		"#0d3472",
		"#b6cff5",
		"#0d3b44",
		"#98d7e4",
		"#3d188e",
		"#e3d7ff",
		"#711a36",
		"#fbd3e0",
		"#8a1c0a",
		"#f2b2a8",
		"#7a2e0b",
		"#ffc8af",
		"#7a4706",
		"#ffdeb5",
		"#594c05",
		"#fbe983",
		"#684e07",
		"#fdedc1",
		"#0b4f30",
		"#b3efd3",
		"#04502e",
		"#a2dcc1",
		"#c2c2c2",
		"#4986e7",
		"#2da2bb",
		"#b99aff",
		"#994a64",
		"#f691b2",
		"#ff7537",
		"#ffad46",
		"#662e37",
		"#ebdbde",
		"#cca6ac",
		"#094228",
		"#42d692",
		"#16a765",
	}
	return gsmhelpers.Contains(color, validColors)
}
