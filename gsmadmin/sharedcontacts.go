/*
Copyright © 2020-2025 Hannes Hayashi

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
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/hanneshayashi/gsm/gsmhelpers"
)

const feedURL = `https://www.google.com/m8/feeds/contacts/%s/full?v=3.0`

// Category is a Feed Category
type Category struct {
	Text   string `xml:",chardata"`
	Scheme string `xml:"scheme,attr"`
	Term   string `xml:"term,attr"`
}

// Title is a Feed Title
type Title struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

// Link is a Feed Link
type Link struct {
	Text string `xml:",chardata"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
}

// Author is a Feed Author
type Author struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

// Generator is a Feed Generator
type Generator struct {
	Text    string `xml:",chardata"`
	Version string `xml:"version,attr"`
	URI     string `xml:"uri,attr"`
}

// Name is a Shared Contact Name
type Name struct {
	Text           string `xml:",chardata"`
	GivenName      string `xml:"givenName"`
	AdditionalName string `xml:"additionalName"`
	FamilyName     string `xml:"familyName"`
	NamePrefix     string `xml:"namePrefix"`
	NameSuffix     string `xml:"nameSuffix"`
	FullName       string `xml:"fullName"`
}

// PhoneNumber is a Shared Contact Phone Number
type PhoneNumber struct {
	PhoneNumber string `xml:",chardata"`
	Rel         string `xml:"rel,attr"`
	Primary     string `xml:"primary,attr"`
	Label       string `xml:"label,attr"`
	URI         string `xml:"uri,attr"`
}

// ExtendedProperty is a Shared Contact Extended Property
type ExtendedProperty struct {
	Text  string `xml:",chardata"`
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
	Realm string `xml:"realm,attr"`
}

// Email is a Shared Contact Email
type Email struct {
	Text        string `xml:",chardata"`
	Rel         string `xml:"rel,attr"`
	Primary     string `xml:"primary,attr"`
	Address     string `xml:"address,attr"`
	DisplayName string `xml:"displayName,attr"`
	Label       string `xml:"label,attr"`
}

// Im is a Shared Contact IM object
type Im struct {
	Text     string `xml:",chardata"`
	Rel      string `xml:"rel,attr"`
	Protocol string `xml:"protocol,attr"`
	Address  string `xml:"address,attr"`
	Label    string `xml:"label,attr"`
	Primary  string `xml:"primary,attr"`
}

// PostalAddress is a Shared Contact Postal Address
type PostalAddress struct {
	PostalAddress string `xml:",chardata"`
	Rel           string `xml:"rel,attr"`
	Primary       string `xml:"primary,attr"`
}

// Organization is a Shared Contact Organization object
type Organization struct {
	Text              string `xml:",chardata"`
	Label             string `xml:"label,attr"`
	Primary           string `xml:"primary,attr"`
	Rel               string `xml:"rel,attr"`
	OrgDepartment     string `xml:"orgDepartment"`
	OrgJobDescription string `xml:"orgJobDescription"`
	OrgName           string `xml:"orgName"`
	OrgSymbol         string `xml:"orgSymbol"`
	OrgTitle          string `xml:"orgTitle"`
	Where             string `xml:"where"`
}

// StructuredPostalAddress is a Shared Contact Structured Postal Address
type StructuredPostalAddress struct {
	Text             string `xml:",chardata"`
	MailClass        string `xml:"mailClass,attr"`
	Label            string `xml:"label,attr"`
	Usage            string `xml:"usage,attr"`
	Primary          string `xml:"primary,attr"`
	Agent            string `xml:"agent"`
	Housename        string `xml:"housename"`
	Street           string `xml:"street"`
	Pobox            string `xml:"pobox"`
	Neighborhood     string `xml:"neighborhood"`
	City             string `xml:"city"`
	Subregion        string `xml:"subregion"`
	Region           string `xml:"region"`
	Postcode         string `xml:"postcode"`
	Country          string `xml:"country"`
	FormattedAddress string `xml:"formattedAddress"`
}

// EntryLinkEntryLink is a Link in an EntryLinkEntry object
type EntryLinkEntryLink struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}

// GeoPt is a GeoPt object in an EntryLinkEntry object
type GeoPt struct {
	Text string `xml:",chardata"`
	Lat  string `xml:"lat,attr"`
	Lon  string `xml:"lon,attr"`
}

// EntryLinkEntryEmail is an Email object in an EntryLinkEntry object
type EntryLinkEntryEmail struct {
	Text    string `xml:",chardata"`
	Address string `xml:"address,attr"`
}

// EntryLinkEntry is a Entry Link Entry object of a Shared Contact Where object
type EntryLinkEntry struct {
	Text          string              `xml:",chardata"`
	ID            string              `xml:"id"`
	Category      Category            `xml:"category"`
	Content       string              `xml:"content"`
	Link          Link                `xml:"link"`
	PostalAddress string              `xml:"postalAddress"`
	GeoPt         GeoPt               `xml:"geoPt"`
	PhoneNumber   string              `xml:"phoneNumber"`
	Email         EntryLinkEntryEmail `xml:"email"`
}

// EntryLink is a Where Entry Link object
type EntryLink struct {
	Text  string         `xml:",chardata"`
	Href  string         `xml:"href,attr"`
	Entry EntryLinkEntry `xml:"entry"`
}

// Where is a Shared Contact Where object
type Where struct {
	Text        string `xml:",chardata"`
	Rel         string `xml:"rel,attr"`
	ValueString string `xml:"valueString,attr"`
	EntryLink   `xml:"entryLink"`
}

// Entry is a Feed Entry
type Entry struct {
	Where                   Where                     `xml:"where"`
	Name                    Name                      `xml:"name"`
	Category                Category                  `xml:"category"`
	Title                   Title                     `xml:"title"`
	Gd                      string                    `xml:"gd,attr"`
	ID                      string                    `xml:"id"`
	Updated                 string                    `xml:"updated"`
	Xmlns                   string                    `xml:"xmlns,attr"`
	Content                 string                    `xml:"content"`
	Text                    string                    `xml:",chardata"`
	ExtendedProperty        []ExtendedProperty        `xml:"extendedProperty"`
	Link                    []Link                    `xml:"link"`
	Email                   []Email                   `xml:"email"`
	Im                      []Im                      `xml:"im"`
	PostalAddress           []PostalAddress           `xml:"postalAddress"`
	Organization            []Organization            `xml:"organization"`
	StructuredPostalAddress []StructuredPostalAddress `xml:"structuredPostalAddress"`
	PhoneNumber             []PhoneNumber             `xml:"phoneNumber"`
}

// Feed was generated 2020-10-11 05:44:12 by hannes_siefert_gmail_com on code-server.
type Feed struct {
	XMLName      xml.Name  `xml:"feed"`
	Text         string    `xml:",chardata"`
	Xmlns        string    `xml:"xmlns,attr"`
	OpenSearch   string    `xml:"openSearch,attr"`
	Gd           string    `xml:"gd,attr"`
	GContact     string    `xml:"gContact,attr"`
	Batch        string    `xml:"batch,attr"`
	ID           string    `xml:"id"`
	Updated      string    `xml:"updated"`
	Category     Category  `xml:"category"`
	Title        Title     `xml:"title"`
	Link         []Link    `xml:"link"`
	Author       Author    `xml:"author"`
	Generator    Generator `xml:"generator"`
	TotalResults string    `xml:"totalResults"`
	StartIndex   string    `xml:"startIndex"`
	ItemsPerPage string    `xml:"itemsPerPage"`
	Entry        []Entry   `xml:"entry"`
}

func makeListSharedContactsCallAndAppend(url string) ([]Entry, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}
	req.Header.Add("GData-Version", "3.0")
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer gsmhelpers.CloseLog(r.Body, "sharedContactBody")
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	feed := Feed{}
	err = xml.Unmarshal(responseBody, &feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v. %s", err, string(responseBody))
	}
	for i := range feed.Link {
		if feed.Link[i].Rel == "next" {
			f, err := makeListSharedContactsCallAndAppend(feed.Link[i].Href)
			if err != nil {
				return nil, err
			}
			feed.Entry = append(feed.Entry, f...)
		}
	}
	return feed.Entry, nil
}

// ListSharedContacts lists all shared contacts in the domain
func ListSharedContacts(domain string) ([]Entry, error) {
	url := fmt.Sprintf(feedURL, domain) + "&max-results=1000"
	entries, err := makeListSharedContactsCallAndAppend(url)
	return entries, err
}

// CreateSharedContact creates a new shared contact in the domain
func CreateSharedContact(domain string, person *Entry) (*Entry, error) {
	personXML, err := xml.Marshal(person)
	if err != nil {
		return nil, fmt.Errorf("error marshalling XML: %v", err)
	}
	body := bytes.NewReader(personXML)
	req, err := http.NewRequest("POST", fmt.Sprintf(feedURL, domain), body)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}
	req.Header.Add("GData-Version", "3.0")
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer gsmhelpers.CloseLog(r.Body, "createSharedContactBody")
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	personC := &Entry{}
	err = xml.Unmarshal(responseBody, personC)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v. %s", err, string(responseBody))
	}
	return personC, nil
}

// DeleteSharedContact deletes a shared contact
func DeleteSharedContact(url string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("GData-Version", "3.0")
	req.Header.Add("If-Match", "*")
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer gsmhelpers.CloseLog(r.Body, "deleteSharedContactBody")
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return responseBody, nil
}

// GetSharedContact retrieves a shared contact
func GetSharedContact(url string) (*Entry, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}
	req.Header.Add("GData-Version", "3.0")
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer gsmhelpers.CloseLog(r.Body, "getShareContactBody")
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	person := &Entry{}
	err = xml.Unmarshal(responseBody, person)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v. %s", err, string(responseBody))
	}
	return person, nil
}

// UpdateSharedContact updates a shared contact
func UpdateSharedContact(url string, person *Entry) (*Entry, error) {
	personXML, err := xml.Marshal(person)
	if err != nil {
		return nil, fmt.Errorf("error marshalling XML: %v", err)
	}
	body := bytes.NewReader(personXML)
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}
	req.Header.Add("GData-Version", "3.0")
	req.Header.Add("If-Match", "*")
	r, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer gsmhelpers.CloseLog(r.Body, "updateSharedContactBody")
	responseBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	personU := &Entry{}
	err = xml.Unmarshal(responseBody, personU)
	if err != nil {
		fmt.Println(string(personXML))
		return nil, fmt.Errorf("error unmarshalling response: %v. %s", err, string(responseBody))
	}
	return personU, nil
}
