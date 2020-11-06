/*
Package cmd contains the commands available to the end user
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
package cmd

import (
	"gsm/gsmhelpers"

	"github.com/spf13/cobra"
	admin "google.golang.org/api/admin/directory/v1"
)

var userFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"userKey": {
		AvailableFor: []string{"delete", "get", "makeAdmin", "patch", "signOut", "undelete"},
		Type:         "string",
		Description: `Identifies the user in the API request.
The value can be the user's primary email address, alias email address, or unique user ID.`,
		Required: []string{"delete", "get", "makeAdmin", "patch", "signOut", "undelete"},
	},
	"customFieldMask": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `A comma-separated list of schema names.
All fields from these schemas are fetched. This should only be set when projection=custom`,
	},
	"projection": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `What subset of fields to fetch for this user.

Acceptable values are:
"basic": Do not include any custom fields for the user. (default)
"custom": Include custom fields from schemas requested in customFieldMask.
"full": Include all fields associated with this user.`,
	},
	"viewType": {
		AvailableFor: []string{"get", "list"},
		Type:         "string",
		Description: `Whether to fetch the administrator-only or domain-wide public view of the user.
For more information, see https://developers.google.com/admin-sdk/directory/v1/guides/manage-users#retrieve_users_non_admin.

Acceptable values are:
"admin_view": Results include both administrator-only and domain-public fields for the user. (default)
"domain_public": Results only include fields for the user that are publicly visible to other users in the domain.
                 Contact sharing must be enabled for the domain.`,
	},
	"fields": {
		AvailableFor: []string{"get", "insert", "list", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
	"familyName": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `The user's last name. Required when creating a user account.`,
		Required:     []string{"insert"},
	},
	"givenName": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `The user's first name. Required when creating a user account.`,
		Required:     []string{"insert"},
	},
	"password": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Stores the password for the user account.
The user's password value is required when creating a user account.
It is optional when updating a user and should only be provided if the user is updating their account password.
A password can contain any combination of ASCII characters.
A minimum of 8 characters is required. The maximum length is 100 characters.
We recommend sending the password property value as a base 16 bit, hexadecimal-encoded hash value.
If a hashFunction is specified, the password must be a valid hash key.
The password value is never returned in the API's response body.`,
		Required: []string{"insert"},
	},
	"primaryEmail": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `The user's primary email address.
This property is required in a request to create a user account.
The primaryEmail must be unique and cannot be an alias of another user.`,
		Required: []string{"insert"},
	},
	"addresses": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `Specifies addresses for the user. May be used multiple times in the form of:
'--addresses "country=...;countryCode=..;customType=..."', etc.
You can use the following properties:
country             - Country.
countryCode         - The country code. Uses the ISO 3166-1 standard.
customType          - If the address type is custom, this property contains the custom value.
extendedAddress     - For extended addresses, such as an address that includes a sub-region.
formatted           - A full and unstructured postal address. This is not synced with the structured address fields.
locality            - The town or city of the address.
poBox               - The post office box, if present.
postalCode          - The ZIP or postal code, if applicable.
primary             - BOOL! If this is the user's primary address. The addresses list may contain only one primary address.
region              - The abbreviated province or state.
sourceIsStructured  - BOOL! Indicates if the user-supplied address was formatted. Formatted addresses are not currently supported.
streetAddress       - The street address, such as 1600 Amphitheatre Parkway. Whitespace within the string is ignored; however, newlines are significant.
type                - The address type.
                      Acceptable values are:
                        - "custom"
                        - "home"
                        - "other"
                        - "work"`,
	},
	"archived": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Indicates if user is archived.`,
	},
	"changePasswordAtNextLogin": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description: `Indicates if the user is forced to change their password at next login.
This setting doesn't apply when the user signs in via a third-party identity provider.`,
	},
	"emails": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `Specifies email addresses for the user. May be used multiple times in the form of:
'--emails "address=...;customType=..;primary=..."', etc.
You can use the following properties:
address     - Country.
customType  - The country code. Uses the ISO 3166-1 standard.
primary     - If the address type is custom, this property contains the custom value.
type        - The type of the email account.
              Acceptable values are:
                - "custom"
                - "home"
                - "other"
                - "work"`,
	},
	"externalIds": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `Specifies externalIds for the user. May be used multiple times in the form of:
'--externalIds "customType=...;type=..;value=..."'
You can use the following properties:
customType  - Country.
type        - The type of the ID.
              Acceptable values are:
                - "account"
                - "custom"
                - "customer"
				- "login_id"
				- "network"
				- "organization": IDs of this type map to employee ID in the Admin Console
value           - The value of the ID.`,
	},
	"customGender": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Custom gender.`,
	},
	"genderType": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Gender.
Acceptable values are:
"female"
"male"
"other"
"unknown"`,
	},
	"hashFunction": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Stores the hash format of the password property.
We recommend sending the password property value as a base 16 bit hexadecimal-encoded hash value.
Set the hashFunction values as either the SHA-1, MD5, or crypt hash format.`,
	},
	"ims": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The user's Instant Messenger (IM) accounts.
A user account can have multiple ims properties.
But, only one of these ims properties can be the primary IM contact.
The maximum allowed data size for this field is 2Kb.
May be used multiple times in the form of:
'--ims "customProtocol=...;customType=..;im=..."', etc.
You can use the following properties:
customProtocol  - If the protocol value is custom_protocol, this property holds the custom protocol's string.
customType      - If the IM type is custom, this property holds the custom type string.
im              - The user's IM network ID.
primary         - BOOL! If this is the user's primary IM. Only one entry in the IM list can have a value of true.
protocol        - An IM protocol identifies the IM network. The value can be a custom network or the standard network.
                  Acceptable values are:
                    - "aim": AOL Instant Messenger protocol
                    - "custom_protocol": A custom IM network protocol
                    - "gtalk": Google Talk protocol
                    - "icq": ICQ protocol
                    - "jabber": Jabber protocol
                    - "msn": MSN Messenger protocol
                    - "net_meeting": Net Meeting protocol
                    - "qq": QQ protocol
                    - "skype": Skype protocol
                    - "yahoo": Yahoo Messenger protocol
type            - The type of the IM account.
                  Acceptable values are:
                   - "custom"
                   - "home"
                   - "other"
                   - "work"`,
	},
	"includeInGlobalAddressList": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description: `Indicates if the user's profile is visible in the G Suite global address list when the contact sharing feature is enabled for the domain.
For more information about excluding user profiles, see the administration help center.`,
	},
	"ipWhitelisted": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `If true, the user's IP address is white listed.`,
	},
	"keywords": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The user's keywords. The maximum allowed data size for this field is 1Kb.
May be used multiple times in the form of:
'--keywords "customType=...;type=..;value=..."'
You can use the following properties:
customType  - Custom Type.
type        - Each entry can have a type which indicates standard type of that entry.
              For example, keyword could be of type occupation or outlook.
              In addition to the standard type, an entry can have a custom type and can give it any name.
              Such types should have the CUSTOM value as type and also have a customType value.
              Acceptable values are:
                - "custom"
                - "mission"
                - "occupation"
				- "outlook"
value  -      Keyword.`,
	},
	"languages": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The user's languages. The maximum allowed data size for this field is 1Kb.
May be used multiple times in the form of:
'--languages "customLanguage=..."'
'--languages "languageCode=..."'
You can use the following properties:
customLanguage  - Other language.
				  A user can provide their own language name if there is no corresponding Google III language code.
			      If this is set, LanguageCode can't be set
languageCode    - Language Code.
                  Should be used for storing Google III LanguageCode string representation for language.
                  Illegal values cause SchemaException.`,
	},
	"locations": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The user's locations. The maximum allowed data size for this field is 10Kb.
May be used multiple times in the form of:
'--locations "area=...;buildingId=...;customType=..."', etc.
You can use the following properties:
area          - Textual location.
		        This is most useful for display purposes to concisely describe the location.
		        For example, "Mountain View, CA", "Near Seattle".
buildingId    - Building identifier.
customType    - If the location type is custom, this property contains the custom value.
deskCode      - Most specific textual code of individual desk location.
floorName     - Floor name/number.
floorSection  - Floor section. More specific location within the floor.
                  For example, if a floor is divided into sections "A", "B", and "C", this field would identify one of those values.
type          - The location type.
                Acceptable values are:
                  - "custom"
                  - "default"
                  - "desk"`,
	},
	"notesContentType": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Content type of note, either plain text or HTML.
Default is plain text. Possible values are:
text_plain
text_html`,
	},
	"notesValue": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Contents of notes.`,
	},
	"orgUnitPath": {
		AvailableFor: []string{"insert", "patch", "undelete"},
		Type:         "string",
		Description: `The full path of the parent organization associated with the user.
If the parent organization is the top-level, it is represented as a forward slash (/).`,
		Required: []string{"undelete"},
	},
	"organizations": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `A list of organizations the user belongs to. The maximum allowed data size for this field is 10Kb.
May be used multiple times in the form of:
'--organizations "costCenter=...;customType=...;department=..."', etc.
You can use the following properties:
costCenter          - The cost center of the user's organization.
customType          - If the value of type is custom, this property contains the custom type.
department          - Specifies the department within the organization, such as 'sales' or 'engineering'.
description         - The description of the organization.
domain              - The domain the organization belongs to.
fullTimeEquivalent  - INT! The full-time equivalent millipercent within the organization (100000 = 100%).	
location            - The physical location of the organization.
			          This does not need to be a fully qualified address.
name                - The name of the organization.
primary             - BOOL! Indicates if this is the user's primary organization.
                      A user may only have one primary organization.
symbol              - Text string symbol of the organization.
                      For example, the text symbol for Google is GOOG.
title               - The user's title within the organization, for example 'member' or 'engineer'.
type                - The type of organization.
                      Acceptable values are:
                        - "domain_only"
                        - "school"
                        - "unknown"
                        - "work"`,
	},
	"phones": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `A list of the user's phone numbers. The maximum allowed data size for this field is 1Kb.
May be used multiple times in the form of:
'--phones "customType=...;primary=...;type=..."', etc.
You can use the following properties:
customType  - If the value of type is custom, this property contains the custom type.
primary     - BOOL! Indicates if this is the user's primary phone number.
              A user may only have one primary phone number.
type        - The type of phone number.
              Acceptable values are:
                - "assistant"
                - "callback"
                - "car"
                - "company_main"
                - "custom"
                - "grand_central"
                - "home"
                - "home_fax"
                - "isdn"
                - "main"
                - "mobile"
                - "other"
                - "other_fax"
                - "pager"
                - "radio"
                - "telex"
                - "tty_tdd"
                - "work"
                - "work_fax"
                - "work_mobile"
		        - "work_pager"
value       - A human-readable phone number.
              It may be in any telephone number format.`,
	},
	"posixAccounts": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `A list of POSIX account information for the user.
May be used multiple times in the form of:
'--posixAccounts "accountId=...;gecos=...;gid=..."', etc.
You can use the following properties:
accountId            - A POSIX account field identifier.
gecos                - The GECOS (user information) for this account.
gid                  - The default group ID.
homeDirectory        - The path to the home directory for this account.
operatingSystemType  - The operating system type for this account.
                       Acceptable values are:
                         - "linux"
                         - "unspecified"
                         - "windows"
primary              - BOOL! If this is user's primary account within the SystemId.
shell                - The path to the login shell for this account.
systemId             - System identifier for which account Username or Uid apply to.
uid                  - The POSIX compliant user ID.
username             - The username of the account.`,
	},
	"recoveryEmail": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description:  `Recovery email of the user.`,
	},
	"recoveryPhone": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "string",
		Description: `Recovery phone of the user.
The phone number must be in the E.164 format, starting with the plus sign (+).
Example: +16506661212.`,
	},
	"relations": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `A list of the user's relationships to other users.
The maximum allowed data size for this field is 2Kb.
May be used multiple times in the form of:
'--relations "customType=...;type=...;value=..."'
You can use the following properties:
customType  - If the value of type is custom, this property contains the custom type.
type        - The type of relation.
              Acceptable values are:
              - "admin_assistant"
              - "assistant"
              - "brother"
              - "child"
              - "custom"
              - "domestic_partner"
              - "dotted_line_manager"
              - "exec_assistant"
              - "father"
              - "friend"
              - "manager"
              - "mother"
              - "parent"
              - "partner"
              - "referred_by"
              - "relative"
              - "sister"
			  - "spouse"
value       - The name of the person the user is related to.`,
	},
	"sshPublicKeys": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `A list of SSH public keys.
May be used multiple times in the form of:
'--sshPublicKeys "expirationTimeUsec=...;key=..."'
You can use the following properties:
expirationTimeUsec  - An expiration time in microseconds since epoch.
key                 - An SSH public key.`,
	},
	"suspended": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "bool",
		Description:  `Indicates if user is suspended.`,
	},
	"websites": {
		AvailableFor: []string{"insert", "patch"},
		Type:         "stringSlice",
		Description: `The user's websites.
The maximum allowed data size for this field is 2Kb.
May be used multiple times in the form of:
'--websites "customType=...;primary=...;type=..."', etc.
You can use the following properties:
customType  - The custom type. Only used if the type is custom.
primary     - BOOL! If this is user's primary website or not.
type        - The type or purpose of the website.
			  For example, a website could be labeled as home or blog.
			  Alternatively, an entry can have a custom type.
			  Custom types must have a customType value.
			  Acceptable values are:
			    - "app_install_page"
			    - "blog"
			    - "custom"
			    - "ftp"
			    - "home"
			    - "home_page"
			    - "other"
			    - "profile"
			    - "reservations"
			    - "resume"
			    - "work"
value       - The URL of the website.`,
	},
	"customer": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The unique ID for the customer's G Suite account.
In case of a multi-domain account, to fetch all groups for a customer, fill this field instead of domain.
You can also use the my_customer alias to represent your account's customerId.
The customerId is also returned as part of the Users resource.
Either the customer or the domain parameter must be provided.`,
		Defaults: map[string]interface{}{"list": "my_customer"},
	},
	"domain": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `The domain name.
Use this field to get fields from only one domain.
To return all domains for a customer account, use the customer query parameter instead.
Either the customer or the domain parameter must be provided.`,
	},
	"orderBy": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Property to use for sorting results.
Acceptable values are:
"email": Primary email of the user.
"familyName": User's family name.
"givenName": User's given name.`,
	},
	"query": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Query string for searching user fields.
For more information on constructing user queries, see https://developers.google.com/admin-sdk/directory/v1/guides/search-users`,
	},
	"showDeleted": {
		AvailableFor: []string{"list"},
		Type:         "bool",
		Description:  `If set to true, retrieves the list of deleted users.`,
	},
	"sortOrder": {
		AvailableFor: []string{"list"},
		Type:         "string",
		Description: `Whether to return results in ascending or descending order.
Acceptable values are:
"ASCENDING": Ascending order.
"DESCENDING": Descending order.`,
	},
	"unmake": {
		AvailableFor: []string{"makeAdmin"},
		Type:         "bool",
		Description:  `Use to remove admin access.`,
	},
}

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Manage Users (Park of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/directory/v1/reference/users",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func mapToUser(flags map[string]*gsmhelpers.Value) (*admin.User, error) {
	user := &admin.User{}
	if flags["addresses"].IsSet() {
		addresses := flags["addresses"].GetStringSlice()
		if len(addresses) > 0 {
			addressesMap := make([]map[string]string, 0)
			for _, a := range addresses {
				m := gsmhelpers.FlagToMap(a)
				addressesMap = append(addressesMap, m)
			}
			user.Addresses = addressesMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Addresses")
		}
	}
	if flags["archived"].IsSet() {
		user.Archived = flags["archived"].GetBool()
		if !user.Archived {
			user.ForceSendFields = append(user.ForceSendFields, "Archived")
		}
	}
	if flags["changePasswordAtNextLogin"].IsSet() {
		user.ChangePasswordAtNextLogin = flags["changePasswordAtNextLogin"].GetBool()
		if !user.ChangePasswordAtNextLogin {
			user.ForceSendFields = append(user.ForceSendFields, "changePasswordAtNextLogin")
		}
	}
	if flags["emails"].IsSet() {
		emails := flags["emails"].GetStringSlice()
		if len(emails) > 0 {
			emailsMap := make([]map[string]string, 0)
			for _, a := range emails {
				m := gsmhelpers.FlagToMap(a)
				emailsMap = append(emailsMap, m)
			}
			user.Emails = emailsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Emails")
		}
	}
	if flags["externalIds"].IsSet() {
		externalIds := flags["externalIds"].GetStringSlice()
		if len(externalIds) > 0 {
			externalIdsMap := make([]map[string]string, 0)
			for _, a := range externalIds {
				m := gsmhelpers.FlagToMap(a)
				externalIdsMap = append(externalIdsMap, m)
			}
			user.ExternalIds = externalIdsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "ExternalIds")
		}
	}
	if flags["customGender"].IsSet() || flags["genderType"].IsSet() {
		gender := map[string]string{
			"customGender": flags["customGender"].GetString(),
			"genderType":   flags["genderType"].GetString(),
		}
		if gender["customGender"] == "" && gender["genderType"] == "" {
			user.ForceSendFields = append(user.ForceSendFields, "Gender")
		} else {
			user.Gender = gender
		}
	}
	if flags["hashFunction"].IsSet() {
		user.HashFunction = flags["hashFunction"].GetString()
		if user.HashFunction == "" {
			user.ForceSendFields = append(user.ForceSendFields, "hashFunction")
		}
	}
	if flags["ims"].IsSet() {
		ims := flags["ims"].GetStringSlice()
		if len(ims) > 0 {
			imsMap := make([]map[string]string, 0)
			for _, a := range ims {
				m := gsmhelpers.FlagToMap(a)
				imsMap = append(imsMap, m)
			}
			user.Ims = imsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Ims")
		}
	}
	if flags["includeInGlobalAddressList"].IsSet() {
		user.IncludeInGlobalAddressList = flags["includeInGlobalAddressList"].GetBool()
		if !user.IncludeInGlobalAddressList {
			user.ForceSendFields = append(user.ForceSendFields, "includeInGlobalAddressList")
		}
	}
	if flags["ipWhitelisted"].IsSet() {
		user.IpWhitelisted = flags["ipWhitelisted"].GetBool()
		if !user.IpWhitelisted {
			user.ForceSendFields = append(user.ForceSendFields, "ipWhitelisted")
		}
	}
	if flags["keywords"].IsSet() {
		keywords := flags["keywords"].GetStringSlice()
		if len(keywords) > 0 {
			keywordsMap := make([]map[string]string, 0)
			for _, a := range keywords {
				m := gsmhelpers.FlagToMap(a)
				keywordsMap = append(keywordsMap, m)
			}
			user.Keywords = keywordsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Keywords")
		}
	}
	if flags["languages"].IsSet() {
		languages := flags["languages"].GetStringSlice()
		if len(languages) > 0 {
			languagesMap := make([]map[string]string, 0)
			for _, a := range languages {
				m := gsmhelpers.FlagToMap(a)
				languagesMap = append(languagesMap, m)
			}
			user.Languages = languagesMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Languages")
		}
	}
	if flags["locations"].IsSet() {
		locations := flags["locations"].GetStringSlice()
		if len(locations) > 0 {
			locationsMap := make([]map[string]string, 0)
			for _, a := range locations {
				m := gsmhelpers.FlagToMap(a)
				locationsMap = append(locationsMap, m)
			}
			user.Locations = locationsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Locations")
		}
	}
	if flags["familyName"].IsSet() || flags["givenName"].IsSet() {
		user.Name = &admin.UserName{}
		if flags["familyName"].IsSet() {
			user.Name.FamilyName = flags["familyName"].GetString()
			if user.Name.FamilyName == "" {
				user.Name.ForceSendFields = append(user.Name.ForceSendFields, "FamilyName")
			}
		}
		if flags["givenName"].IsSet() {
			user.Name.GivenName = flags["givenName"].GetString()
			if user.Name.GivenName == "" {
				user.Name.ForceSendFields = append(user.Name.ForceSendFields, "GivenName")
			}
		}
	}
	if flags["notesContentType"].IsSet() || flags["notesValue"].IsSet() {
		notes := map[string]string{
			"contentType": flags["notesContentType"].GetString(),
			"value":       flags["notesValue"].GetString(),
		}
		if notes["contentType"] == "" && notes["value"] == "" {
			user.ForceSendFields = append(user.ForceSendFields, "Notes")
		} else {
			user.Notes = notes
		}
	}
	if flags["orgUnitPath"].IsSet() {
		user.OrgUnitPath = flags["orgUnitPath"].GetString()
		if user.OrgUnitPath == "" {
			user.ForceSendFields = append(user.ForceSendFields, "OrgUnitPath")
		}
	}
	if flags["organizations"].IsSet() {
		organizations := flags["organizations"].GetStringSlice()
		if len(organizations) > 0 {
			organizationsMap := make([]map[string]string, 0)
			for _, a := range organizations {
				m := gsmhelpers.FlagToMap(a)
				organizationsMap = append(organizationsMap, m)
			}
			user.Organizations = organizationsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Organizations")
		}
	}
	if flags["password"].IsSet() {
		user.Password = flags["password"].GetString()
		if user.Password == "" {
			user.ForceSendFields = append(user.ForceSendFields, "Password")
		}
	}
	if flags["phones"].IsSet() {
		phones := flags["phones"].GetStringSlice()
		if len(phones) > 0 {
			phonesMap := make([]map[string]string, 0)
			for _, a := range phones {
				m := gsmhelpers.FlagToMap(a)
				phonesMap = append(phonesMap, m)
			}
			user.Phones = phonesMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Phones")
		}
	}
	if flags["posixAccounts"].IsSet() {
		posixaccounts := flags["posixAccounts"].GetStringSlice()
		if len(posixaccounts) > 0 {
			posixaccountsMap := make([]map[string]string, 0)
			for _, a := range posixaccounts {
				m := gsmhelpers.FlagToMap(a)
				posixaccountsMap = append(posixaccountsMap, m)
			}
			user.PosixAccounts = posixaccountsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "PosixAccounts")
		}
	}
	if flags["primaryEmail"].IsSet() {
		user.PrimaryEmail = flags["primaryEmail"].GetString()
		if user.PrimaryEmail == "" {
			user.ForceSendFields = append(user.ForceSendFields, "PrimaryEmail")
		}
	}
	if flags["recoveryEmail"].IsSet() {
		user.RecoveryEmail = flags["recoveryEmail"].GetString()
		if user.RecoveryEmail == "" {
			user.ForceSendFields = append(user.ForceSendFields, "RecoveryEmail")
		}
	}
	if flags["recoveryPhone"].IsSet() {
		user.RecoveryPhone = flags["recoveryPhone"].GetString()
		if user.RecoveryPhone == "" {
			user.ForceSendFields = append(user.ForceSendFields, "RecoveryPhone")
		}
	}
	if flags["relations"].IsSet() {
		relations := flags["relations"].GetStringSlice()
		if len(relations) > 0 {
			relationsMap := make([]map[string]string, 0)
			for _, a := range relations {
				m := gsmhelpers.FlagToMap(a)
				relationsMap = append(relationsMap, m)
			}
			user.Relations = relationsMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Relations")
		}
	}
	if flags["sshPublicKeys"].IsSet() {
		sshPublicKeys := flags["sshPublicKeys"].GetStringSlice()
		if len(sshPublicKeys) > 0 {
			sshPublicKeysMap := make([]map[string]string, 0)
			for _, a := range sshPublicKeys {
				m := gsmhelpers.FlagToMap(a)
				sshPublicKeysMap = append(sshPublicKeysMap, m)
			}
			user.SshPublicKeys = sshPublicKeysMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "SshPublicKeys")
		}
	}
	if flags["suspended"].IsSet() {
		user.Suspended = flags["suspended"].GetBool()
		if !user.Suspended {
			user.ForceSendFields = append(user.ForceSendFields, "suspended")
		}
	}
	if flags["websites"].IsSet() {
		websites := flags["websites"].GetStringSlice()
		if len(websites) > 0 {
			websitesMap := make([]map[string]string, 0)
			for _, a := range websites {
				m := gsmhelpers.FlagToMap(a)
				websitesMap = append(websitesMap, m)
			}
			user.Websites = websitesMap
		} else {
			user.ForceSendFields = append(user.ForceSendFields, "Websites")
		}
	}
	return user, nil
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
