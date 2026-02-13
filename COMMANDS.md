# GSM Command Inventory

Auto-generated — run `GSM_GENERATE_INVENTORY=1 go test -run TestGenerateInventory ./cmd/` to update.

## Summary

- **Total commands**: 735
- **Leaf commands** (executable): 422

## Commands

### `gsm about`

Information about the user, the user's Drive, and system capabilities (Part of Drive API)

#### `gsm about get`

Gets information about the user, the user's Drive, and system capabilities.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

### `gsm accessProposals`

Manage Access Proposals on a file (Part of Drive API)

#### `gsm accessProposals get`

Retrieves an AccessProposal by ID.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fileId` | string | ✓ | The id of the item the request is on. |
| `proposalId` | string | ✓ | The id of the access proposal. |

#### `gsm accessProposals list`

List the AccessProposals on a file. Note: Only approvers are able to list AccessProposals on a file. If the user is not an approver, returns a 403.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fileId` | string | ✓ | The id of the item the request is on. |

#### `gsm accessProposals resolve`

Used to approve or deny an Access Proposal.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `action` | string | ✓ | The action to take on the AccessProposal. Must be one of the following: ACCEPT  The user accepts the proposal. Note: ... |
| `fileId` | string | ✓ | The id of the item the request is on. |
| `proposalId` | string | ✓ | The id of the access proposal. |
| `role` | stringSlice |  | The roles the approver has allowed, if any. Note: This field is required for the ACCEPT action. This flag can be used... |
| `sendNotification` | bool |  | Whether to send an email to the requester when the AccessProposal is denied or accepted. |
| `view` | string |  | Indicates the view for this access proposal. This should only be set when the proposal belongs to a view. 'published'... |

### `gsm activities`

Manage (list) activities (Part of Admin SDK API)

#### `gsm activities list`

Retrieves a list of activities for a specific customer's account and application such as the Admin console application or the Google Drive application.
For more information, see the guides for administrator and Google Drive activity reports.
For more information about the activity report's parameters, see the activity parameters reference guides.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `actorIpAddress` | string |  | The Internet Protocol (IP) Address of host where the event was performed. This is an additional way to filter a repor... |
| `applicationName` | string | ✓ | Application name for which the events are to be retrieved. The following values are accepted: ACCESS_TRANSPARENCY   -... |
| `customerId` | string |  | The unique ID of the customer to retrieve data for. |
| `endTime` | string |  | Sets the end of the range of time shown in the report. The date is in the RFC 3339 format, for example 2010-10-28T10:... |
| `eventName` | string |  | The name of the event being queried by the API. Each eventName is related to a specific Workspace service or feature ... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filters` | string |  | The filters query string is a comma-separated list. The list is composed of event parameters that are manipulated by ... |
| `groupIdFilter` | string |  | Comma separated group ids (obfuscated) on which user activities are filtered, i.e, the response will contain activiti... |
| `orgUnitId` | string |  | ID of the organizational unit to report on. Activity records will be shown only for users who belong to the specified... |
| `startTime` | string |  | Sets the beginning of the range of time shown in the report. The date is in the RFC 3339 format, for example 2010-10-... |
| `userKey` | string |  | Represents the profile ID or the user email for which the data should be filtered. Can be all for all information, or... |

### `gsm apps`

Information apps a user has installed (Part of Drive API)

#### `gsm apps get`

Gets a specific app. For more information, see https://developers.google.com/workspace/drive/api/guides/user-info.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `appId` | string |  | The ID of the app. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm apps list`

Lists a user's installed apps. For more information, see https://developers.google.com/workspace/drive/api/guides/user-info.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `appFilterExtensions` | string |  | A comma-separated list of file extensions to limit returned results. All results within the given app query scope whi... |
| `appFilterMimeTypes` | string |  | A comma-separated list of file extensions to limit returned results. All results within the given app query scope whi... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | A language or locale code, as defined by BCP 47, with some extensions from Unicode's LDML format (http://www.unicode.... |

### `gsm asps`

Manage ASPs (application-specific password) for a user (Part of Admin SDK API)

#### `gsm asps delete`

Delete an ASP issued by a user.

##### `gsm asps delete batch`

Batch deletes ASPs issued by a user using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `codeId` | int64 |  | The unique ID of the ASP |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

#### `gsm asps get`

Get information about an ASP issued by a user.

##### `gsm asps get batch`

Batch gets information about ASPs issued by a user using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `codeId` | int64 |  | The unique ID of the ASP |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

#### `gsm asps list`

List the ASPs issued by a user.

##### `gsm asps list batch`

Batch lists ASPs issued by a user using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

##### `gsm asps list recursive`

List the ASPs issued by users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

### `gsm attachments`

Manage (get..) message attachments (Part of Gmail API)

#### `gsm attachments get`

Gets the specified message attachment.

##### `gsm attachments get batch`

Batch gets the specified message attachments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the attachment. |
| `messageId` | int64 |  | The ID of the message containing the attachment. |
| `messageId_ALL` | string |  | Same as messageId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

### `gsm buildings`

Manage Buildings (Resources) (Part of Admin SDK API)

#### `gsm buildings delete`

Deletes a building.

##### `gsm buildings delete batch`

Batch deletes buildings using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `buildingId` | int64 |  | The ID of the file. |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm buildings get`

Retrieves a building.

##### `gsm buildings get batch`

Batch retrieves buildings using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `buildingId` | int64 |  | The ID of the file. |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm buildings insert`

Inserts a building.

##### `gsm buildings insert batch`

Batch inserts buildings using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addressLines` | int64 |  | Unstructured address lines describing the lower levels of an address. |
| `addressLines_ALL` | stringSlice |  | Same as addressLines but value is applied to all lines in the CSV file |
| `administrativeArea` | int64 |  | Optional. Highest administrative subdivision which is used for postal addresses of a country or region. |
| `administrativeArea_ALL` | string |  | Same as administrativeArea but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `buildingId` | int64 |  | The ID of the file. |
| `buildingName` | int64 |  | The building name as seen by users in Calendar. Must be unique for the customer. For example, "NYC-CHEL". The maximum... |
| `buildingName_ALL` | string |  | Same as buildingName but value is applied to all lines in the CSV file |
| `coordinatesSource` | int64 |  | Source from which Building.coordinates are derived.  Acceptable values are: CLIENT_SPECIFIED       - Building.coordin... |
| `coordinatesSource_ALL` | string |  | Same as coordinatesSource but value is applied to all lines in the CSV file |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | A brief description of the building. For example, "Chelsea Market". |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `floorNames` | int64 |  | The display names for all floors in this building. The floors are expected to be sorted in ascending order, from lowe... |
| `floorNames_ALL` | stringSlice |  | Same as floorNames but value is applied to all lines in the CSV file |
| `languageCode` | int64 |  | Optional. BCP-47 language code of the contents of this address (if known). |
| `languageCode_ALL` | string |  | Same as languageCode but value is applied to all lines in the CSV file |
| `latitude` | int64 |  | Latitude in decimal degrees. |
| `latitude_ALL` | float64 |  | Same as latitude but value is applied to all lines in the CSV file |
| `locality` | int64 |  | Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In re... |
| `locality_ALL` | string |  | Same as locality but value is applied to all lines in the CSV file |
| `longitude` | int64 |  | Longitude in decimal degrees. |
| `longitude_ALL` | float64 |  | Same as longitude but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `postalCode` | int64 |  | Optional. Postal code of the address. |
| `postalCode_ALL` | string |  | Same as postalCode but value is applied to all lines in the CSV file |
| `regionCode` | int64 |  | Required. CLDR region code of the country/region of the address. |
| `regionCode_ALL` | string |  | Same as regionCode but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sublocality` | int64 |  | Optional. Sublocality of the address. |
| `sublocality_ALL` | string |  | Same as sublocality but value is applied to all lines in the CSV file |

#### `gsm buildings list`

Retrieves a list of buildings for an account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm buildings patch`

Updates a building. This method supports patch semantics.

##### `gsm buildings patch batch`

Batch updates buildings using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addressLines` | int64 |  | Unstructured address lines describing the lower levels of an address. |
| `addressLines_ALL` | stringSlice |  | Same as addressLines but value is applied to all lines in the CSV file |
| `administrativeArea` | int64 |  | Optional. Highest administrative subdivision which is used for postal addresses of a country or region. |
| `administrativeArea_ALL` | string |  | Same as administrativeArea but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `buildingId` | int64 |  | The ID of the file. |
| `buildingName` | int64 |  | The building name as seen by users in Calendar. Must be unique for the customer. For example, "NYC-CHEL". The maximum... |
| `buildingName_ALL` | string |  | Same as buildingName but value is applied to all lines in the CSV file |
| `coordinatesSource` | int64 |  | Source from which Building.coordinates are derived.  Acceptable values are: CLIENT_SPECIFIED       - Building.coordin... |
| `coordinatesSource_ALL` | string |  | Same as coordinatesSource but value is applied to all lines in the CSV file |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | A brief description of the building. For example, "Chelsea Market". |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `floorNames` | int64 |  | The display names for all floors in this building. The floors are expected to be sorted in ascending order, from lowe... |
| `floorNames_ALL` | stringSlice |  | Same as floorNames but value is applied to all lines in the CSV file |
| `languageCode` | int64 |  | Optional. BCP-47 language code of the contents of this address (if known). |
| `languageCode_ALL` | string |  | Same as languageCode but value is applied to all lines in the CSV file |
| `latitude` | int64 |  | Latitude in decimal degrees. |
| `latitude_ALL` | float64 |  | Same as latitude but value is applied to all lines in the CSV file |
| `locality` | int64 |  | Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In re... |
| `locality_ALL` | string |  | Same as locality but value is applied to all lines in the CSV file |
| `longitude` | int64 |  | Longitude in decimal degrees. |
| `longitude_ALL` | float64 |  | Same as longitude but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `postalCode` | int64 |  | Optional. Postal code of the address. |
| `postalCode_ALL` | string |  | Same as postalCode but value is applied to all lines in the CSV file |
| `regionCode` | int64 |  | Required. CLDR region code of the country/region of the address. |
| `regionCode_ALL` | string |  | Same as regionCode but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sublocality` | int64 |  | Optional. Sublocality of the address. |
| `sublocality_ALL` | string |  | Same as sublocality but value is applied to all lines in the CSV file |

### `gsm calendarAcl`

Manage entries in users' calendar acl (Part of Calendar API)

#### `gsm calendarAcl delete`

Deletes an access control rule.

##### `gsm calendarAcl delete batch`

Batch deletes ACL rules using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarAcl.acl method. If you want to access the primary cale... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `ruleId` | int64 |  | ACL rule identifier. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarAcl get`

Returns an access control rule.

##### `gsm calendarAcl get batch`

Batch gets ACL rules using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarAcl.acl method. If you want to access the primary cale... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `ruleId` | int64 |  | ACL rule identifier. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarAcl insert`

Creates an access control rule.

##### `gsm calendarAcl insert batch`

Batch inserts ACL rules using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarAcl.acl method. If you want to access the primary cale... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `role` | int64 |  | The role assigned to the scope. Possible values are: "none" - Provides no access. "freeBusyReader" - Provides read ac... |
| `role_ALL` | string |  | Same as role but value is applied to all lines in the CSV file |
| `scopeType` | int64 |  | The type of the scope. Possible values are: "default" - The public scope. This is the default value. "user" - Limits ... |
| `scopeType_ALL` | string |  | Same as scopeType but value is applied to all lines in the CSV file |
| `scopeValue` | int64 |  | The email address of a user or group, or the name of a domain, depending on the scope type. Omitted for type "default". |
| `scopeValue_ALL` | string |  | Same as scopeValue but value is applied to all lines in the CSV file |
| `sendNotifications` | int64 |  | Whether to send notifications about the calendar sharing change. Optional. The default is True. |
| `sendNotifications_ALL` | bool |  | Same as sendNotifications but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarAcl list`

Returns the rules in the access control list for the calendar.

##### `gsm calendarAcl list batch`

Batch lists ACL rules using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarAcl.acl method. If you want to access the primary cale... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `showDeleted` | int64 |  | Whether to include deleted ACLs in the result. Deleted ACLs are represented by role equal to "none". Deleted ACLs wil... |
| `showDeleted_ALL` | bool |  | Same as showDeleted but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarAcl patch`

Updates an access control rule. This method supports patch semantics.

##### `gsm calendarAcl patch batch`

Batch patches ACL rules using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarAcl.acl method. If you want to access the primary cale... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `role` | int64 |  | The role assigned to the scope. Possible values are: "none" - Provides no access. "freeBusyReader" - Provides read ac... |
| `role_ALL` | string |  | Same as role but value is applied to all lines in the CSV file |
| `ruleId` | int64 |  | ACL rule identifier. |
| `scopeType` | int64 |  | The type of the scope. Possible values are: "default" - The public scope. This is the default value. "user" - Limits ... |
| `scopeType_ALL` | string |  | Same as scopeType but value is applied to all lines in the CSV file |
| `scopeValue` | int64 |  | The email address of a user or group, or the name of a domain, depending on the scope type. Omitted for type "default". |
| `scopeValue_ALL` | string |  | Same as scopeValue but value is applied to all lines in the CSV file |
| `sendNotifications` | int64 |  | Whether to send notifications about the calendar sharing change. Optional. The default is True. |
| `sendNotifications_ALL` | bool |  | Same as sendNotifications but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm calendarLists`

Manage entries in users' calendar list (Part of Calendar API)

#### `gsm calendarLists delete`

Removes a calendar from the user's calendar list.

##### `gsm calendarLists delete batch`

Batch deletes calendars from the user's calendar list using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarLists get`

Returns a calendar from the user's calendar list.

##### `gsm calendarLists get batch`

Batch returns calendars from the user's calendar list using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarLists insert`

Inserts an existing calendar into the user's calendar list.

##### `gsm calendarLists insert batch`

Batch inserts existing calendar entries using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `backgroundColor` | int64 |  | The main color of the calendar in the hexadecimal format "#0088aa". This property supersedes the index-based colorId ... |
| `backgroundColor_ALL` | string |  | Same as backgroundColor but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `colorId` | int64 |  | The color of the calendar. This is an ID referring to an entry in the calendar section of the colors definition (see ... |
| `colorId_ALL` | string |  | Same as colorId but value is applied to all lines in the CSV file |
| `colorRgbFormat` | int64 |  | Whether to use the foregroundColor and backgroundColor fields to write the calendar colors (RGB). If this feature is ... |
| `colorRgbFormat_ALL` | bool |  | Same as colorRgbFormat but value is applied to all lines in the CSV file |
| `defaultReminders` | int64 |  | The default reminders that the authenticated user has for this calendar. Must be given in the form of '--defaultRemin... |
| `defaultReminders_ALL` | stringSlice |  | Same as defaultReminders but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `foregroundColor` | int64 |  | The foreground color of the calendar in the hexadecimal format "#ffffff". This property supersedes the index-based co... |
| `foregroundColor_ALL` | string |  | Same as foregroundColor but value is applied to all lines in the CSV file |
| `hidden` | int64 |  | Whether the calendar has been hidden from the list. |
| `hidden_ALL` | bool |  | Same as hidden but value is applied to all lines in the CSV file |
| `id` | int64 |  | Identifier of the calendar. |
| `notificationsType` | int64 |  | The type of notification. [eventCreation\|eventChange\|eventCancellation\|eventResponse\|agenda] eventCreation      -... |
| `notificationsType_ALL` | stringSlice |  | Same as notificationsType but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `selected` | int64 |  | Whether the calendar content shows up in the calendar UI |
| `selected_ALL` | bool |  | Same as selected but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `summaryOverride` | int64 |  | The summary that the authenticated user has set for this calendar. |
| `summaryOverride_ALL` | string |  | Same as summaryOverride but value is applied to all lines in the CSV file |

#### `gsm calendarLists list`

Returns the calendars on the user's calendar list.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `calendarId` | string |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `minAccessRole` | string |  | The minimum access role for the user in the returned entries. Optional. The default is no restriction. [freeBusyReade... |
| `showDeleted` | bool |  | Whether to include deleted calendar list entries in the result. |
| `showHidden` | bool |  | Whether to show hidden entries. |

#### `gsm calendarLists patch`

Updates an existing calendar on the user's calendar list. This method supports patch semantics.

##### `gsm calendarLists patch batch`

Batch patches existing calendar list entries using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `backgroundColor` | int64 |  | The main color of the calendar in the hexadecimal format "#0088aa". This property supersedes the index-based colorId ... |
| `backgroundColor_ALL` | string |  | Same as backgroundColor but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `colorId` | int64 |  | The color of the calendar. This is an ID referring to an entry in the calendar section of the colors definition (see ... |
| `colorId_ALL` | string |  | Same as colorId but value is applied to all lines in the CSV file |
| `colorRgbFormat` | int64 |  | Whether to use the foregroundColor and backgroundColor fields to write the calendar colors (RGB). If this feature is ... |
| `colorRgbFormat_ALL` | bool |  | Same as colorRgbFormat but value is applied to all lines in the CSV file |
| `defaultReminders` | int64 |  | The default reminders that the authenticated user has for this calendar. Must be given in the form of '--defaultRemin... |
| `defaultReminders_ALL` | stringSlice |  | Same as defaultReminders but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `foregroundColor` | int64 |  | The foreground color of the calendar in the hexadecimal format "#ffffff". This property supersedes the index-based co... |
| `foregroundColor_ALL` | string |  | Same as foregroundColor but value is applied to all lines in the CSV file |
| `hidden` | int64 |  | Whether the calendar has been hidden from the list. |
| `hidden_ALL` | bool |  | Same as hidden but value is applied to all lines in the CSV file |
| `id` | int64 |  | Identifier of the calendar. |
| `notificationsType` | int64 |  | The type of notification. [eventCreation\|eventChange\|eventCancellation\|eventResponse\|agenda] eventCreation      -... |
| `notificationsType_ALL` | stringSlice |  | Same as notificationsType but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `selected` | int64 |  | Whether the calendar content shows up in the calendar UI |
| `selected_ALL` | bool |  | Same as selected but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `summaryOverride` | int64 |  | The summary that the authenticated user has set for this calendar. |
| `summaryOverride_ALL` | string |  | Same as summaryOverride but value is applied to all lines in the CSV file |

### `gsm calendarResources`

Manage resource calendars (Part of Admin SDK API)

#### `gsm calendarResources delete`

Deletes a calendar resource.

##### `gsm calendarResources delete batch`

Batch deletes calendar resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarResourceId` | int64 |  | The unique ID of the calendar resource |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarResources get`

Gets a calendar resource.

##### `gsm calendarResources get batch`

Batch retrieves calendar resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarResourceId` | int64 |  | The unique ID of the calendar resource |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarResources insert`

Inserts a calendar resource.

##### `gsm calendarResources insert batch`

Batch inserts calendar resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `buildingId` | int64 |  | Unique ID for the building a resource is located in. |
| `buildingId_ALL` | string |  | Same as buildingId but value is applied to all lines in the CSV file |
| `capacity` | int64 |  | Capacity of a resource, number of seats in a room. |
| `capacity_ALL` | int64 |  | Same as capacity but value is applied to all lines in the CSV file |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `featureInstances` | int64 |  | Instances of features for the calendar resource. |
| `featureInstances_ALL` | string |  | Same as featureInstances but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `floorName` | int64 |  | Name of the floor a resource is located on. |
| `floorName_ALL` | string |  | Same as floorName but value is applied to all lines in the CSV file |
| `floorSection` | int64 |  | Name of the section within a floor a resource is located in. |
| `floorSection_ALL` | string |  | Same as floorSection but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceCategory` | int64 |  | The category of the calendar resource. Either CONFERENCE_ROOM or OTHER. Legacy data is set to CATEGORY_UNKNOWN.  Acce... |
| `resourceCategory_ALL` | string |  | Same as resourceCategory but value is applied to all lines in the CSV file |
| `resourceDescription` | int64 |  | Description of the resource, visible only to admins. |
| `resourceDescription_ALL` | string |  | Same as resourceDescription but value is applied to all lines in the CSV file |
| `resourceId` | int64 |  | The unique ID of the calendar resource |
| `resourceName` | int64 |  | The name of the calendar resource. For example, "Training Room 1A". |
| `resourceName_ALL` | string |  | Same as resourceName but value is applied to all lines in the CSV file |
| `resourceType` | int64 |  | The type of the calendar resource, intended for non-room resources. |
| `resourceType_ALL` | string |  | Same as resourceType but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userVisibleDescription` | int64 |  | Description of the resource, visible to users and admins. |
| `userVisibleDescription_ALL` | string |  | Same as userVisibleDescription but value is applied to all lines in the CSV file |

#### `gsm calendarResources list`

Retrieves a list of calendar resources for an account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `orderBy` | string |  | Field(s) to sort results by in either ascending or descending order. Supported fields include resourceId, resourceNam... |
| `query` | string |  | 	String query used to filter results. Should be of the form "field operator value" where field can be any of supporte... |

#### `gsm calendarResources patch`

Patches a calendar resource.

##### `gsm calendarResources patch batch`

Batch patches calendar resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `buildingId` | int64 |  | Unique ID for the building a resource is located in. |
| `buildingId_ALL` | string |  | Same as buildingId but value is applied to all lines in the CSV file |
| `calendarResourceId` | int64 |  | The unique ID of the calendar resource |
| `capacity` | int64 |  | Capacity of a resource, number of seats in a room. |
| `capacity_ALL` | int64 |  | Same as capacity but value is applied to all lines in the CSV file |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `featureInstances` | int64 |  | Instances of features for the calendar resource. |
| `featureInstances_ALL` | string |  | Same as featureInstances but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `floorName` | int64 |  | Name of the floor a resource is located on. |
| `floorName_ALL` | string |  | Same as floorName but value is applied to all lines in the CSV file |
| `floorSection` | int64 |  | Name of the section within a floor a resource is located in. |
| `floorSection_ALL` | string |  | Same as floorSection but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceCategory` | int64 |  | The category of the calendar resource. Either CONFERENCE_ROOM or OTHER. Legacy data is set to CATEGORY_UNKNOWN.  Acce... |
| `resourceCategory_ALL` | string |  | Same as resourceCategory but value is applied to all lines in the CSV file |
| `resourceDescription` | int64 |  | Description of the resource, visible only to admins. |
| `resourceDescription_ALL` | string |  | Same as resourceDescription but value is applied to all lines in the CSV file |
| `resourceName` | int64 |  | The name of the calendar resource. For example, "Training Room 1A". |
| `resourceName_ALL` | string |  | Same as resourceName but value is applied to all lines in the CSV file |
| `resourceType` | int64 |  | The type of the calendar resource, intended for non-room resources. |
| `resourceType_ALL` | string |  | Same as resourceType but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userVisibleDescription` | int64 |  | Description of the resource, visible to users and admins. |
| `userVisibleDescription_ALL` | string |  | Same as userVisibleDescription but value is applied to all lines in the CSV file |

### `gsm calendarSettings`

See users' calendar settings (Part of Calendar API)

#### `gsm calendarSettings get`

Returns a single user setting.

##### `gsm calendarSettings get batch`

Batch gets calendar settings using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `setting` | int64 |  | The id of the user setting. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendarSettings list`

Returns all user settings for the authenticated user.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

### `gsm calendars`

Manage users' calendars (Part of Calendar API)

#### `gsm calendars clear`

Clears a primary calendar.
This operation deletes all events associated with the primary calendar of an account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `calendarId` | string | ✓ | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |

#### `gsm calendars delete`

Deletes a secondary calendar.
Use calendars.clear for clearing all events on primary calendars.

##### `gsm calendars delete batch`

Batch deletes secondary calendars using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendars get`

Returns metadata for a calendar.

##### `gsm calendars get batch`

Batch gets secondary calendars using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm calendars insert`

Creates a secondary calendar.

##### `gsm calendars insert batch`

Batch inserts secondary calendars using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | Description of the calendar. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `location` | int64 |  | Geographic location of the calendar as free-form text. |
| `location_ALL` | string |  | Same as location but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `summary` | int64 |  | Title of the calendar. |
| `summary_ALL` | string |  | Same as summary but value is applied to all lines in the CSV file |
| `timeZone` | int64 |  | The time zone of the calendar. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich"). |
| `timeZone_ALL` | string |  | Same as timeZone but value is applied to all lines in the CSV file |

#### `gsm calendars patch`

Updates metadata for a calendar.

##### `gsm calendars patch batch`

Batch patches secondary calendars using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | Description of the calendar. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `location` | int64 |  | Geographic location of the calendar as free-form text. |
| `location_ALL` | string |  | Same as location but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `summary` | int64 |  | Title of the calendar. |
| `summary_ALL` | string |  | Same as summary but value is applied to all lines in the CSV file |
| `timeZone` | int64 |  | The time zone of the calendar. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich"). |
| `timeZone_ALL` | string |  | Same as timeZone but value is applied to all lines in the CSV file |

### `gsm changes`

View changes to user's or Shared Drive (Part of Drive API)

#### `gsm changes getStartPageToken`

Gets the starting pageToken for listing future changes.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `driveId` | string |  | The ID of the shared drive |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm changes list`

Lists the changes for a user or shared drive.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `driveId` | string |  | The ID of the shared drive |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `includeCorpusRemovals` | bool |  | Whether changes should include the file resource if the file is still accessible by the user at the time of the reque... |
| `includeItemsFromAllDrives` | bool |  | Whether both My Drive and shared drive items should be included in results. |
| `includePermissionsForView` | string |  | Specifies which additional view's permissions to include in the response. Only 'published' is supported. |
| `includeRemoved` | bool |  | Whether to include changes indicating that items have been removed from the list of changes, for example by deletion ... |
| `pageToken` | string | ✓ | The token for continuing a previous list request on the next page. This should be set to the value of 'nextPageToken'... |
| `restrictToMyDrive` | bool |  | Whether to restrict the results to changes inside the My Drive hierarchy. This omits changes to files such as those i... |
| `spaces` | string |  | A comma-separated list of spaces to query within the user corpus. Supported values are 'drive', 'appDataFolder' and '... |

### `gsm chromeOs`

Issue Commands to Chrome OS Devices (Part of Admin SDK API)

#### `gsm chromeOs issueCommand`

Takes an issueCommand that affects a Chrome OS Device. This includes deprovisioning, disabling, and re-enabling devices.

##### `gsm chromeOs issueCommand batch`

Batch issues commands to Chrome OS devices using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commandType` | int64 |  | The type of command.  Acceptable values are: REBOOT             - Reboot the device.                      Can only be... |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `deviceId` | int64 |  | Immutable ID of Chrome OS Device. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `payload` | int64 |  | The payload for the command, provide it only if command supports it. The following commands support adding payload: -... |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm chromeOs`

Issue Commands to Chrome OS Devices (Part of Admin SDK API)

#### `gsm chromeOs issueCommand`

Takes an issueCommand that affects a Chrome OS Device. This includes deprovisioning, disabling, and re-enabling devices.

##### `gsm chromeOs issueCommand batch`

Batch issues commands to Chrome OS devices using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commandType` | int64 |  | The type of command.  Acceptable values are: REBOOT             - Reboot the device.                      Can only be... |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `deviceId` | int64 |  | Immutable ID of Chrome OS Device. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `payload` | int64 |  | The payload for the command, provide it only if command supports it. The following commands support adding payload: -... |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm chromeOsDevices`

Managed Chrome OS Devices (Part of Admin SDK API)

#### `gsm chromeOsDevices action`

Takes an action that affects a Chrome OS Device. This includes deprovisioning, disabling, and re-enabling devices.

##### `gsm chromeOsDevices action batch`

Batch takes actions that affect Chrome OS devices using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `action` | int64 |  | Action to be taken on the Chrome OS device  Acceptable values are: deprovision  - Remove a device from management tha... |
| `action_ALL` | string |  | Same as action but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `deprovisionReason` | int64 |  | Only used when the action is deprovision. With the deprovision action, this field is required.  Note: The deprovision... |
| `deprovisionReason_ALL` | string |  | Same as deprovisionReason but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceId` | int64 |  | The unique ID of the device. The resourceIds are returned in the response from the chromeosdevices.list method. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm chromeOsDevices get`

Retrieves a Chrome OS device's properties.

##### `gsm chromeOsDevices get batch`

Batch retrieves Chrome OS devices's properties using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `deviceId` | int64 |  | The unique ID of the device. The deviceIds are returned in the response from the chromeosdevices.list method. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `projection` | int64 |  | Determines whether the response contains the full list of properties or only a subset.  Acceptable values are: BASIC ... |
| `projection_ALL` | string |  | Same as projection but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm chromeOsDevices list`

Retrieves a paginated list of Chrome OS devices within an account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `orderBy` | string |  | Device property to use for sorting results.  Acceptable values are: annotatedLocation  - Chrome device location as an... |
| `orgUnitPath` | string |  | The full path of the organizational unit or its unique ID. |
| `projection` | string |  | Determines whether the response contains the full list of properties or only a subset.  Acceptable values are: BASIC ... |
| `query` | string |  | Search string in the format provided by List query operators (https://developers.google.com/admin-sdk/directory/v1/li... |
| `sortOrder` | string |  | Whether to return results in ascending or descending order. Must be used with the orderBy parameter.  Acceptable valu... |

#### `gsm chromeOsDevices moveToOU`

Move or insert multiple Chrome OS devices to an organizational unit. You can move up to 50 devices at once.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string | ✓ | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `deviceIds` | stringSlice | ✓ | Chrome OS devices to be moved to OU |
| `orgUnitPath` | string | ✓ | The full path of the organizational unit or its unique ID. |

#### `gsm chromeOsDevices patch`

Updates a device's updatable properties, such as annotatedUser, annotatedLocation, notes, orgUnitPath, or annotatedAssetId. This method supports patch semantics

##### `gsm chromeOsDevices patch batch`

Batch patch Chrome OS devices using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `annotatedAssetId` | int64 |  | The asset identifier as noted by an administrator or specified during enrollment. |
| `annotatedAssetId_ALL` | string |  | Same as annotatedAssetId but value is applied to all lines in the CSV file |
| `annotatedLocation` | int64 |  | The address or location of the device as noted by the administrator. Maximum length is 200 characters. Empty values a... |
| `annotatedLocation_ALL` | string |  | Same as annotatedLocation but value is applied to all lines in the CSV file |
| `annotatedUser` | int64 |  | The user of the device as noted by the administrator. Maximum length is 100 characters. Empty values are allowed. |
| `annotatedUser_ALL` | string |  | Same as annotatedUser but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `deviceId` | int64 |  | The unique ID of the device. The deviceIds are returned in the response from the chromeosdevices.list method. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `notes` | int64 |  | Notes about this device added by the administrator. This property can be searched with the list method's query parame... |
| `notes_ALL` | string |  | Same as notes but value is applied to all lines in the CSV file |
| `orgUnitPath` | int64 |  | The full path of the organizational unit or its unique ID. |
| `orgUnitPath_ALL` | string |  | Same as orgUnitPath but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `projection` | int64 |  | Determines whether the response contains the full list of properties or only a subset.  Acceptable values are: BASIC ... |
| `projection_ALL` | string |  | Same as projection but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm chromePrinters`

Managed Chrome Printers (Part of Admin SDK API)

#### `gsm chromePrinters batchCreate`

Creates printers under given Organization Unit.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string |  | The name of the customer. Format: customers/{customer_id} |
| `printer` | stringSlice | ✓ | A printer to create. If you want to place the printer under particular OU then populate orgUnitId filed. Otherwise th... |

#### `gsm chromePrinters batchDelete`

Deletes printers in batch.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `parent` | string |  | The name of the customer. Format: customers/{customer_id} |
| `printerIds` | stringSlice | ✓ | A list of printer ids that should be deleted. Max 100 at a time. |

#### `gsm chromePrinters create`

Creates a printer under given Organization Unit.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `description` | string |  | Editable. Description of printer. |
| `displayName` | string | ✓ | Editable. Name of printer. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `makeAndModel` | string |  | Editable. Make and model of printer. e.g. Lexmark MS610de Value must be in format as seen in printers.listPrinterMode... |
| `orgUnitId` | string | ✓ | Organization Unit |
| `parent` | string |  | The name of the customer. Format: customers/{customer_id} |
| `uri` | string | ✓ | Editable. Printer URI. |
| `useDriverlessConfig` | bool |  | Editable. flag to use driverless configuration or not. If it's set to be true, makeAndModel can be ignored |

#### `gsm chromePrinters delete`

Deletes a Printer.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | The name of the printer to be updated. Format: customers/{customer_id}/chrome/printers/{printer_id} |

#### `gsm chromePrinters get`

Returns a Printer resource (printer's config).

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string | ✓ | The name of the printer to be updated. Format: customers/{customer_id}/chrome/printers/{printer_id} |

#### `gsm chromePrinters list`

List printers configs.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | Search query. Search syntax is shared between this api and Admin Console printers pages. |
| `parent` | string |  | The name of the customer. Format: customers/{customer_id} |

#### `gsm chromePrinters listModels`

Lists the supported printer models.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | Search query. Search syntax is shared between this api and Admin Console printers pages. |
| `parent` | string |  | The name of the customer. Format: customers/{customer_id} |

#### `gsm chromePrinters patch`

Updates a Printer resource.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `clearMask` | string |  | The list of fields to be cleared. Note, some of the fields are read only and cannot be updated. Values for not specif... |
| `description` | string |  | Editable. Description of printer. |
| `displayName` | string |  | Editable. Name of printer. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `makeAndModel` | string |  | Editable. Make and model of printer. e.g. Lexmark MS610de Value must be in format as seen in printers.listPrinterMode... |
| `name` | string | ✓ | The name of the printer to be updated. Format: customers/{customer_id}/chrome/printers/{printer_id} |
| `updateMask` | string |  | The list of fields to be updated. Note, some of the fields are read only and cannot be updated. Values for not specif... |
| `uri` | string |  | Editable. Printer URI. |
| `useDriverlessConfig` | bool |  | Editable. flag to use driverless configuration or not. If it's set to be true, makeAndModel can be ignored |

### `gsm clientStates`

Manage client states (Part of Cloud Identity API)

#### `gsm clientStates get`

Gets the client state for the device user

##### `gsm clientStates get batch`

Batch gets client states using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the ClientState in format: devices/{device_id}/deviceUsers/{device_user_id}/clientStates/{partner_id... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm clientStates list`

Lists the client states for the given search query.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | Additional restrictions when fetching list of client states. |
| `orderBy` | string |  | Order specification for client states in the response. |
| `parent` | string | ✓ | To list all ClientStates, set this to "devices/-/deviceUsers/-". To list all ClientStates owned by a DeviceUser, set ... |

#### `gsm clientStates patch`

Updates the client state for the device user

##### `gsm clientStates patch batch`

Batch patches client states using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `assetTags` | int64 |  | The caller can specify asset tags for this resource |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customId` | int64 |  | This field may be used to store a unique identifier for the API resource within which these CustomAttributes are a fi... |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `etag` | int64 |  | The token that needs to be passed back for concurrency control in updates. Token needs to be passed back in UpdateReq... |
| `etag_ALL` | string |  | Same as etag but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `keyValuePairs` | int64 |  | The map of key-value attributes stored by callers specific to a device. The total serialized length of this map may n... |
| `keyValuePairs_ALL` | stringSlice |  | Same as keyValuePairs but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the ClientState in format: devices/{device_id}/deviceUsers/{device_user_id}/clientStates/{partner_id... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm colors`

Show Calendar and Event color definitions (Part of Calendar API)

#### `gsm colors get`

Returns the color definitions for calendars and events.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

### `gsm comments`

Manage comments in Google files (Part of Drive API)

#### `gsm comments create`

Creates a new comment on a file.

##### `gsm comments create batch`

Batch creates new comments on a file using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `anchor` | int64 |  | A region of the document represented as a JSON string. See anchor documentation for details on how to define and inte... |
| `anchor_ALL` | string |  | Same as anchor but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `content` | int64 |  | The plain text content of the comment. This field is used for setting the content, while htmlContent should be displa... |
| `content_ALL` | string |  | Same as content but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `quotedFileContentValue` | int64 |  | The quoted content itself. This is interpreted as plain text if set through the API. |
| `quotedFileContentValue_ALL` | string |  | Same as quotedFileContentValue but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm comments delete`

Deletes a comment.

##### `gsm comments delete batch`

Batch deletes comments by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commentId` | int64 |  | The ID of the comment. |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm comments get`

Gets a comment by ID.

##### `gsm comments get batch`

Batch gets comments by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commentId` | int64 |  | The ID of the comment. |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `includeDeleted` | int64 |  | Whether to return deleted comments. Deleted comments will not include their original content. |
| `includeDeleted_ALL` | bool |  | Same as includeDeleted but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm comments list`

Lists a file's comments.

##### `gsm comments list batch`

Batch lists comments in files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `includeDeleted` | int64 |  | Whether to return deleted comments. Deleted comments will not include their original content. |
| `includeDeleted_ALL` | bool |  | Same as includeDeleted but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `startModifiedTime` | int64 |  | The minimum value of 'modifiedTime' for the result comments (RFC 3339 date-time). |
| `startModifiedTime_ALL` | string |  | Same as startModifiedTime but value is applied to all lines in the CSV file |

#### `gsm comments update`

Updates a comment with patch semantics.

##### `gsm comments update batch`

Batch updates comments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `anchor` | int64 |  | A region of the document represented as a JSON string. See anchor documentation for details on how to define and inte... |
| `anchor_ALL` | string |  | Same as anchor but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commentId` | int64 |  | The ID of the comment. |
| `content` | int64 |  | The plain text content of the comment. This field is used for setting the content, while htmlContent should be displa... |
| `content_ALL` | string |  | Same as content but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `quotedFileContentValue` | int64 |  | The quoted content itself. This is interpreted as plain text if set through the API. |
| `quotedFileContentValue_ALL` | string |  | Same as quotedFileContentValue but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm configs`

Configure GSM

#### `gsm configs get`

Return a single GSM config

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string |  | Name of the configuration. This (plus ".yaml") will be used as the file name. |

#### `gsm configs getDefaultScopes`

Returns the default scopes.

*No command-specific flags.*

#### `gsm configs getScopes`

Returns the scopes of a config file so they can be easily added in the Admin Console

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string |  | Name of the configuration. This (plus ".yaml") will be used as the file name. |

#### `gsm configs list`

List current configurations

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `details` | bool |  | List detailed information about configs. |

#### `gsm configs load`

Load a config file

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | Name of the configuration. This (plus ".yaml") will be used as the file name. |

#### `gsm configs new`

Create a new config file.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `credentialsFile` | string |  | Path to the credential file. Can be relative to the binary or fully qualified. |
| `errorOutput` | string |  | The destination where errors should be output to. Can be 'stderr', 'log' or 'both' |
| `logFile` | string |  | Path of the log file. |
| `mode` | string | ✓ | The mode to operate in. Can be: [dwd\|user\|adc] |
| `name` | string | ✓ | Name of the configuration. This (plus ".yaml") will be used as the file name. |
| `scopes` | stringSlice |  | OAuth Scopes to use. |
| `serviceAccount` | string |  | The Service Account that should be impersonated when using ADC (Application Default Credentials) mode. If you are usi... |
| `standardDelay` | int64 |  | Delay in ms to wait after each API call |
| `subject` | string |  | The user who should be impersonated with DWD. |
| `threads` | int64 |  | The maximum number of threads to use. |

#### `gsm configs remove`

Removes a configuration

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | Name of the configuration. This (plus ".yaml") will be used as the file name. |

#### `gsm configs resetScopes`

Resets the scopes of a config back to the defaults.

*No command-specific flags.*

#### `gsm configs update`

Updates the current config.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `credentialsFile` | string |  | Path to the credential file. Can be relative to the binary or fully qualified. |
| `errorOutput` | string |  | The destination where errors should be output to. Can be 'stderr', 'log' or 'both' |
| `logFile` | string |  | Path of the log file. |
| `name` | string |  | Name of the configuration. This (plus ".yaml") will be used as the file name. |
| `scopes` | stringSlice |  | OAuth Scopes to use. |
| `serviceAccount` | string |  | The Service Account that should be impersonated when using ADC (Application Default Credentials) mode. If you are usi... |
| `standardDelay` | int64 |  | Delay in ms to wait after each API call |
| `subject` | string |  | The user who should be impersonated with DWD. |
| `threads` | int64 |  | The maximum number of threads to use. |

### `gsm contactDelegates`

Manage users' contact contact delegations (Part of Admin SDK API)

#### `gsm contactDelegates create`

Creates one or more delegates for a given user.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string | ✓ | Email of the delegate. |
| `parent` | string | ✓ | The email address of the user whose contacts should be delegated. |

#### `gsm contactDelegates delete`

Deletes a delegate from a given user.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string | ✓ | Email of the delegate. |
| `parent` | string | ✓ | The email address of the user whose contacts should be delegated. |

#### `gsm contactDelegates list`

Lists the delegates of a given user.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `parent` | string | ✓ | The email address of the user whose contacts should be delegated. |

### `gsm contactGroups`

Manage users' contact groups (Part of People API)

#### `gsm contactGroups batchGet`

Get a list of contact groups owned by the authenticated user by specifying a list of contact group resource names.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `maxMembers` | int64 |  | Specifies the maximum number of members to return for each group. Defaults to 0 if not set, which will return zero me... |
| `resourceNames` | stringSlice | ✓ | The resource names of the contact groups. |

#### `gsm contactGroups create`

Create a new contact group owned by the authenticated user.

##### `gsm contactGroups create batch`

Batch creates contact groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The contact group name set by the group owner or a system provided name for system groups. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm contactGroups delete`

Delete an existing contact group owned by the authenticated user by specifying a contact group resource name.

##### `gsm contactGroups delete batch`

Batch deletes contact groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `deleteContacts` | int64 |  | Set to true to also delete the contacts in the specified group. |
| `deleteContacts_ALL` | bool |  | Same as deleteContacts but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceName` | int64 |  | The resource name of the contact group. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm contactGroups get`

Get a specific contact group owned by the authenticated user by specifying a contact group resource name.

##### `gsm contactGroups get batch`

Batch gets contact groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `maxMembers` | int64 |  | Specifies the maximum number of members to return for each group. Defaults to 0 if not set, which will return zero me... |
| `maxMembers_ALL` | int64 |  | Same as maxMembers but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceName` | int64 |  | The resource name of the contact group. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm contactGroups list`

List all contact groups owned by the authenticated user.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm contactGroups update`

Update the name of an existing contact group owned by the authenticated user.

##### `gsm contactGroups update batch`

Batch updates contact groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The contact group name set by the group owner or a system provided name for system groups. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceName` | int64 |  | The resource name of the contact group. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm contactGroupsMembers`

Modify members of contact groups (Part of People API)

#### `gsm contactGroupsMembers modify`

Modify the members of a contact group owned by the authenticated user.
The only system contact groups that can have members added are contactGroups/myContacts and contactGroups/starred.
Other system contact groups are deprecated and can only have contacts removed.

##### `gsm contactGroupsMembers modify batch`

Batch modifies contact groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceName` | int64 |  | The resource name of the contact group to modify. |
| `resourceName_ALL` | string |  | Same as resourceName but value is applied to all lines in the CSV file |
| `resourceNamesToAdd` | int64 |  | The resource names of the contact people to add in the form of people/{person_id}. |
| `resourceNamesToAdd_ALL` | stringSlice |  | Same as resourceNamesToAdd but value is applied to all lines in the CSV file |
| `resourceNamesToRemove` | int64 |  | The resource names of the contact people to remove in the form of people/{person_id}. |
| `resourceNamesToRemove_ALL` | stringSlice |  | Same as resourceNamesToRemove but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm customerUsageReports`

Manage (get) Customer Usage Reports (Part of Admin SDK API)

#### `gsm customerUsageReports get`

Retrieves a report which is a collection of properties and statistics for a specific customer's account.
For more information, see the Customers Usage Report guide.
For more information about the customer report's parameters, see the Customers Usage parameters reference guides.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The unique ID of the customer to retrieve data for. |
| `date` | string |  | Represents the date the usage occurred. The timestamp is in the ISO 8601 format, yyyy-mm-dd. We recommend you use you... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parameters` | string |  | The parameters query string is a comma-separated list of event parameters that refine a report's results. The paramet... |

### `gsm customers`

Implements customers API (Part of Admin SDK API).

#### `gsm customers get`

Retrieves a customer using an ID or retrieve your own customer without knowing your ID.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerKey` | string |  | Id of the customer. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm customers patch`

Patches a customer.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addressLine1` | string |  | A customer's physical address. The address can be composed of one to three lines. |
| `addressLine2` | string |  | Address line 2 of the address. |
| `addressLine3` | string |  | Address line 3 of the address. |
| `alternateEmail` | string |  | The customer's secondary contact email address. This email address cannot be on the same domain as the customerDomain |
| `contactName` | string |  | The customer contact's name. |
| `countryCode` | string |  | This is a required property. For countryCode information see the ISO 3166 country code elements.(http://www.iso.org/i... |
| `customerDomain` | string |  | The customer's primary domain name string. Do not include the www prefix when creating a new customer. |
| `customerKey` | string |  | Id of the customer. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `language` | string |  | The customer's ISO 639-2 language code. See the Language Codes page for the list of supported codes. Valid language c... |
| `locality` | string |  | Name of the locality. An example of a locality value is the city of San Francisco. |
| `organizationName` | string |  | The company or company division name. |
| `phoneNumber` | string |  | The customer's contact phone number in E.164 format. |
| `postalCode` | string |  | The postal code. A postalCode example is a postal zip code such as 10009. This is in accordance with - http://portabl... |
| `region` | string |  | Name of the region. An example of a region value is NY for the state of New York. |

### `gsm delegates`

Manage Gmail Delegates (Part of Gmail API)

#### `gsm delegates create`

Adds a delegate with its verification status set directly to accepted, without sending any verification email.

##### `gsm delegates create batch`

Batch adds delegates using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delegateEmail` | int64 |  | The email address of the delegate. |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm delegates delete`

Removes the specified delegate

##### `gsm delegates delete batch`

Batch deletes the specified delegates using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delegateEmail` | int64 |  | The email address of the delegate. |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm delegates get`

Gets the specified delegate.

##### `gsm delegates get batch`

Batch gets the specified delegates using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delegateEmail` | int64 |  | The email address of the delegate. |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm delegates list`

Lists the delegates for the specified account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value me can be used to indicate the authenticated user. |

### `gsm deviceUsers`

Manage device users (Part of Cloud Identity API)

#### `gsm deviceUsers approve`

Approves device to access user data.

##### `gsm deviceUsers approve batch`

Batch approves device users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}/deviceUsers/{device_user_id}, where device_id is the uniqu... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm deviceUsers block`

Blocks device from accessing user data

##### `gsm deviceUsers block batch`

Batch blocks device users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}/deviceUsers/{device_user_id}, where device_id is the uniqu... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm deviceUsers cancelWipe`

Cancels an unfinished user account wipe.

##### `gsm deviceUsers cancelWipe batch`

Batch cancels device user wipes using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}/deviceUsers/{device_user_id}, where device_id is the uniqu... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm deviceUsers delete`

Deletes the specified DeviceUser.

##### `gsm deviceUsers delete batch`

Batch deletes device users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}/deviceUsers/{device_user_id}, where device_id is the uniqu... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm deviceUsers get`

Retrieves the specified DeviceUser

##### `gsm deviceUsers get batch`

Batch gets device users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}/deviceUsers/{device_user_id}, where device_id is the uniqu... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm deviceUsers list`

Lists/Searches DeviceUsers.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | Additional restrictions when fetching list of devices. For a list of search fields, refer to https://developers.googl... |
| `orderBy` | string |  | Order specification for devices in the response. |
| `parent` | string | ✓ | To list all DeviceUsers, set this to "devices/-". To list all DeviceUsers owned by a device, set this to the resource... |

#### `gsm deviceUsers lookup`

Looks up resource names of the DeviceUsers associated with the caller's credentials, as well as the properties provided in the request.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `androidId` | string |  | Android Id returned by Settings.Secure#ANDROID_ID (https://developer.android.com/reference/android/provider/Settings.... |
| `parent` | string |  | To list all DeviceUsers, set this to "devices/-". To list all DeviceUsers owned by a device, set this to the resource... |
| `rawResourceId` | string |  | Raw Resource Id used by Google Endpoint Verification. If the user is enrolled into Google Endpoint Verification, this... |
| `userId` | string |  | The user whose DeviceUser's resource name will be fetched. Must be set to 'me' to fetch the DeviceUser's resource nam... |

#### `gsm deviceUsers wipe`

Wipes the user's account on a device.

##### `gsm deviceUsers wipe batch`

Batch wipes device users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}/deviceUsers/{device_user_id}, where device_id is the uniqu... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm devices`

Manage Devices (Part of Cloud Identity API)

#### `gsm devices cancelWipe`

Cancels an unfinished device wipe.

##### `gsm devices cancelWipe batch`

Batch cancels pending device wipes using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}, where device_id is the unique ID assigned to the Device. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm devices create`

Creates a device. Only company-owned device may be created.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `assetTag` | string |  | Asset tag of the device. |
| `customer` | string |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `deviceType` | string | ✓ | Type of device: ANDROID      Device is an Android device IOS          Device is an iOS device GOOGLE_SYNC  Device is ... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `serialNumber` | string | ✓ | Serial Number of device. Example: HT82V1A01076. |

#### `gsm devices delete`

Deletes the specified device.

##### `gsm devices delete batch`

Batch deletes the specified devices using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}, where device_id is the unique ID assigned to the Device. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm devices get`

Retrieves the specified device.

##### `gsm devices get batch`

Batch gets the specified devices using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}, where device_id is the unique ID assigned to the Device. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm devices list`

Lists/Searches devices.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | Additional restrictions when fetching list of devices. For a list of search fields, refer to https://developers.googl... |
| `orderBy` | string |  | Order specification for devices in the response. Only one of the following field names may be used to specify the ord... |
| `view` | string |  | The view to use for the List request. Possible values are: COMPANY_INVENTORY      This view contains all devices impo... |

#### `gsm devices wipe`

Wipes all data on the specified device.

##### `gsm devices wipe batch`

Batch wipes devices using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Resource name of the customer. If you're using this API for your own organization, use customers/my_customer. If you'... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `name` | int64 |  | Resource name of the Device in format: devices/{device_id}, where device_id is the unique ID assigned to the Device. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm domainAliases`

Manage Domain Aliases (Part of Admin SDK API)

#### `gsm domainAliases delete`

Deletes a Domain Alias of the customer.

##### `gsm domainAliases delete batch`

Batch deletes domain aliases of the customer using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domainAliasName` | int64 |  | Name of domain alias. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm domainAliases get`

Retrieves a domain alias of the customer.

##### `gsm domainAliases get batch`

Batch retrieves domain aliases of the customer using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domainAliasName` | int64 |  | Name of domain alias. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm domainAliases insert`

Inserts a Domain alias of the customer.

##### `gsm domainAliases insert batch`

Batch inserts Domain aliases of a customer using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domainAliasName` | int64 |  | Name of domain alias. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `parentDomainName` | int64 |  | Name of domain alias. |
| `parentDomainName_ALL` | string |  | Same as parentDomainName but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm domainAliases list`

Lists the domain aliases of the customer.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | Immutable ID of the Workspace account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parentDomainName` | string |  | Name of domain alias. |

### `gsm domains`

Manage Domains (Part of Admin SDK API)

#### `gsm domains delete`

Deletes a Domain of the customer.

##### `gsm domains delete batch`

Batch deletes domains of the customer using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domainName` | int64 |  | Name of domain . |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm domains get`

Retrieves a domain of the customer.

##### `gsm domains get batch`

Batch retrieves domains of the customer using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domainName` | int64 |  | Name of domain . |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm domains insert`

Inserts a Domain of the customer.

##### `gsm domains insert batch`

Batch inserts Domains of the customer using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domainName` | int64 |  | Name of domain . |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm domains list`

Lists the domains of the customer.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | Immutable ID of the Workspace account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

### `gsm drafts`

Manage Drafts (Part of Gmail API)

#### `gsm drafts create`

Creates a new draft with the DRAFT label.

##### `gsm drafts create batch`

Batch creates new drafts with the DRAFT label using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `attachment` | int64 |  | Path to a file that should be attached to the message. Can be used multiple times. |
| `attachment_ALL` | stringSlice |  | Same as attachment but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `bcc` | int64 |  | Blind Copy (Bcc) |
| `bcc_ALL` | string |  | Same as bcc but value is applied to all lines in the CSV file |
| `body` | int64 |  | Body or content of the (draft) message |
| `body_ALL` | string |  | Same as body but value is applied to all lines in the CSV file |
| `cc` | int64 |  | Copy (Cc) |
| `cc_ALL` | string |  | Same as cc but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `html` | int64 |  | Send the body as HTML |
| `html_ALL` | bool |  | Same as html but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `subject` | int64 |  | Subject of the (draft) message |
| `subject_ALL` | string |  | Same as subject but value is applied to all lines in the CSV file |
| `to` | int64 |  | Recipient of the (draft) message |
| `to_ALL` | string |  | Same as to but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm drafts delete`

Immediately and permanently deletes the specified draft. Does not simply trash it.

##### `gsm drafts delete batch`

Batch deletes the specified drafts using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `id` | int64 |  | The ID of the draft. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm drafts get`

Gets the specified draft.

##### `gsm drafts get batch`

Batch gets the specified drafts using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `format` | int64 |  | The format to return the draft in. "[MINIMAL\|FULL\|RAW\|METADATA]. MINIMAL   - Returns only email message ID and lab... |
| `format_ALL` | string |  | Same as format but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the draft. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm drafts list`

Lists the drafts in the user's mailbox.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `includeSpamTrash` | bool |  | Include drafts from SPAM and TRASH in the results. |
| `q` | string |  | Only return draft messages matching the specified query. Supports the same query format as the Gmail search box. For ... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm drafts send`

Sends the specified, existing draft to the recipients in the To, Cc, and Bcc headers.

##### `gsm drafts send batch`

Batch sends drafts using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `id` | int64 |  | The ID of the draft. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm drafts update`

Replaces a draft's content.

##### `gsm drafts update batch`

Batch updates drafts using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `attachment` | int64 |  | Path to a file that should be attached to the message. Can be used multiple times. |
| `attachment_ALL` | stringSlice |  | Same as attachment but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `bcc` | int64 |  | Blind Copy (Bcc) |
| `bcc_ALL` | string |  | Same as bcc but value is applied to all lines in the CSV file |
| `body` | int64 |  | Body or content of the (draft) message |
| `body_ALL` | string |  | Same as body but value is applied to all lines in the CSV file |
| `cc` | int64 |  | Copy (Cc) |
| `cc_ALL` | string |  | Same as cc but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `html` | int64 |  | Send the body as HTML |
| `html_ALL` | bool |  | Same as html but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the draft. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `subject` | int64 |  | Subject of the (draft) message |
| `subject_ALL` | string |  | Same as subject but value is applied to all lines in the CSV file |
| `to` | int64 |  | Recipient of the (draft) message |
| `to_ALL` | string |  | Same as to but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

### `gsm driveLabelLimits`

Manages Drive Label Limits (Part of Drive Labels API)

#### `gsm driveLabelLimits getLabel`

Get the constraints on the structure of a Label; such as, the maximum number of Fields allowed and maximum length of the label title.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string |  | Label revision resource name. API docs say this must be: "limits/label". However, only an empty string seems to work ... |

### `gsm driveLabelLocks`

Manages Drive Label Locks (Part of Drive Labels API)

#### `gsm driveLabelLocks list`

Lists the LabelLocks on a Label.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string | ✓ | Label on which Locks are applied. Format: labels/{label}. If you don't specify the "labels/" prefix, GSM will automat... |

### `gsm driveLabelPermissions`

Manages Drive Label Permissions (Part of Drive Labels API)

#### `gsm driveLabelPermissions batchDelete`

Deletes Label permissions.
Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `parent` | string | ✓ | The parent Label resource name. Format: labels/{label} If you don't specify the "labels/" prefix, GSM will automatica... |
| `permissionName` | stringSlice | ✓ | Label Permission resource name. Format: labels/{label} May be used multiple times to delete multiple permissions at o... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabelPermissions batchUpdate`

Updates Label permissions.
If a permission for the indicated principal doesn't exist, a new Label Permission is created, otherwise the existing permission is updated.
Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string | ✓ | The parent Label resource name. Format: labels/{label} If you don't specify the "labels/" prefix, GSM will automatica... |
| `permission` | stringSlice | ✓ | A permission. In order to update an existing permission use the following format: "name=...;role=..."  In order to cr... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabelPermissions create`

Updates a Label's permissions.
If a permission for the indicated principal doesn't exist, a new Label Permission is created, otherwise the existing permission is updated.
Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `audience` | string |  | Audience to grant a role to. The magic value of audiences/default may be used to apply the role to the default audien... |
| `email` | string |  | Specifies the email address for a user or group principal. Not populated for audience principals. User and Group perm... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string | ✓ | The parent Label resource name. Format: labels/{label} If you don't specify the "labels/" prefix, GSM will automatica... |
| `role` | string | ✓ | The role the principal should have. "[READER\|APPLIER\|ORGANIZER\|EDITOR]. READER     - A reader can read the label a... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabelPermissions delete`

Deletes a principal's permission on a Label.
Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | Resource name of this permission. Format: labels/{label} If you don't specify the "labels/" prefix, GSM will automati... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabelPermissions list`

Lists the LabelPermissions on a Label.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string | ✓ | The parent Label resource name. Format: labels/{label} If you don't specify the "labels/" prefix, GSM will automatica... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabelPermissions update`

Updates a Label's permissions.
The permission must exist and be referenced with the "name" parameter.
Permissions affect the Label resource as a whole, are not revisioned, and do not require publishing.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string | ✓ | Resource name of this permission. Format: labels/{label} If you don't specify the "labels/" prefix, GSM will automati... |
| `parent` | string | ✓ | The parent Label resource name. Format: labels/{label} If you don't specify the "labels/" prefix, GSM will automatica... |
| `role` | string | ✓ | The role the principal should have. "[READER\|APPLIER\|ORGANIZER\|EDITOR]. READER     - A reader can read the label a... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

### `gsm driveLabelUsers`

Manages Drive Label Users (Part of Drive Labels API)

#### `gsm driveLabelUsers getCapabilities`

Gets the user capabilities.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | The customer to scope this request to. For example: "customers/abcd1234". If unset, will return settings within the c... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string |  | The resource name of the user. Only "users/me/capabilities" is supported. (Default) |

### `gsm driveLabels`

Managed driveLabels (Part of Drive Labels API)

#### `gsm driveLabels create`

Creates a new Label.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `description` | string | ✓ | The description of the label. |
| `field` | stringSlice |  | Defines a field that has a display name, data type, and other configuration options. This field defines the kind of m... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `labelType` | string |  | The type of this label. Defaults to SHARED. [SHARED\|ADMIN] SHARED  - Shared labels may be shared with users to apply... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `learnMoreUri` | string |  | Custom URL to present to users to allow them to learn more about this label and how it should be used. |
| `title` | string | ✓ | Title of the label. |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabels createField`

Creates a new field for an existing label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `choice` | stringSlice |  | A choice for a selection field. Can be used multiple times to create multiple choices that will be set in the order s... |
| `displayName` | string | ✓ | The display text to show in the UI identifying this item. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `insertBeforeField` | string |  | Input only. Insert or move this field before the indicated field. If empty, the field is placed at the end of the list. |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `required` | bool |  | Whether the field should be marked as required. |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `valueType` | string |  | The type of the field May be one of the following: - dateString  - A date field - integer     - An integer field - se... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels createSelectionChoice`

Creates a new choice for an existing label selection field on a label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `alpha` | float64 |  | The alpha value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `blue` | float64 |  | The blue value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `displayName` | string | ✓ | The display text to show in the UI identifying this item. |
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `green` | float64 |  | The green value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `insertBeforeChoice` | string |  | Input only. Insert or move this choice before the indicated choice. If empty, the choice is placed at the end of the ... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `priorityOverride` | int64 |  | Override the default global priority of this badge. When set to 0, the default priority heuristic is used. |
| `red` | float64 |  | The red value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels delete`

Deletes a Label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabels deleteField`

Deletes a field from an existing label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels deleteSelectionChoice`

Deletes a choice for an existing label selection field on a label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `choiceId` | string | ✓ | The ID of the choice. |
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels disable`

Disables a Label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `hideInSearch` | bool |  | Whether to hide this disabled object in the search menu for Drive items. When false, the object is generally shown in... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `showInApply` | bool |  | Whether to show this disabled object in the apply menu on Drive items. When true, the object is generally shown in th... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabels disableField`

Disables a field for an existing label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `hideInSearch` | bool |  | Whether to hide this disabled object in the search menu for Drive items. When false, the object is generally shown in... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `showInApply` | bool |  | Whether to show this disabled object in the apply menu on Drive items. When true, the object is generally shown in th... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels disableSelectionChoice`

Disables a choice for an existing label selection field on a label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `choiceId` | string | ✓ | The ID of the choice. |
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `hideInSearch` | bool |  | Whether to hide this disabled object in the search menu for Drive items. When false, the object is generally shown in... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `showInApply` | bool |  | Whether to show this disabled object in the apply menu on Drive items. When true, the object is generally shown in th... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels enable`

Enables a Label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |

#### `gsm driveLabels enableField`

Enables a field for an existing label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels enableSelectionChoice`

Enables a choice for an existing label selection field on a label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `choiceId` | string | ✓ | The ID of the choice. |
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels get`

Get a label by its resource name.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels list`

List labels.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `minimumRole` | string |  | Specifies the level of access the user must have on the returned Labels. The minimum role a user must have on a label... |
| `publishedOnly` | bool |  | Whether to include only published labels in the results.  When true, only the current published label revisions are r... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels publish`

Publishes a Label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels updateField`

Updates basic properties of a Field.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `displayName` | string |  | The display text to show in the UI identifying this item. |
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `insertBeforeField` | string |  | Input only. Insert or move this field before the indicated field. If empty, the field is placed at the end of the list. |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `required` | bool |  | Whether the field should be marked as required. |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels updateFieldType`

Updates a field for an existing label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `choice` | stringSlice |  | A choice for a selection field. Can be used multiple times to create multiple choices that will be set in the order s... |
| `dateFormatType` | string |  | Localized date format options. May be one of the following: LONG_DATE   - Includes full month name. For example, Janu... |
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `maxEntries` | int64 |  | The maximum number of entries for the field as a whole number. Can be used with "user" or "selection type fields |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `valueType` | string |  | The type of the field May be one of the following: - dateString  - A date field - integer     - An integer field - se... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels updateLabel`

Updates basic properties of a Label.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `description` | string |  | The description of the label. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `title` | string |  | Title of the label. |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels updateLabelCopyMode`

Updates a Label's CopyMode

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `copyMode` | string | ✓ | Indicates how the applied label and field values should be copied when a Drive item is copied. May be one of the foll... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

#### `gsm driveLabels updateSelectionChoiceProperties`

Updates the properties of a choice for an existing selection field on a label

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `alpha` | float64 |  | The alpha value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `blue` | float64 |  | The blue value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `choiceId` | string | ✓ | The ID of the choice. |
| `displayName` | string |  | The display text to show in the UI identifying this item. |
| `fieldId` | string | ✓ | The ID of the field. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `green` | float64 |  | The green value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `insertBeforeChoice` | string |  | Input only. Insert or move this choice before the indicated choice. If empty, the choice is placed at the end of the ... |
| `languageCode` | string |  | The BCP-47 language code to use for evaluating localized field labels. When not specified, values in the default conf... |
| `name` | string | ✓ | Label resource name. May be any of:   - labels/{id} (equivalent to labels/{id}@latest)   - labels/{id}@latest   - lab... |
| `priorityOverride` | int64 |  | Override the default global priority of this badge. When set to 0, the default priority heuristic is used. |
| `red` | float64 |  | The red value for the badge color as a float (number between 1 and 0 - e.g. "0.5") |
| `requiredRevisionId` | string |  | The [revisionId][google.apps.drive.labels.v1.Label.revision_id] of the label that the write request will be applied t... |
| `useAdminAccess` | bool |  | Set to true in order to use the user's admin credentials. The server verifies that the user is an admin for the label... |
| `view` | string |  | When specified, only certain fields belonging to the indicated view are returned. [LABEL_VIEW_BASIC\|LABEL_VIEW_FULL]... |

### `gsm drives`

Manage Shared Drives (Part of Drive API)

#### `gsm drives create`

Creates a new shared drive.

##### `gsm drives create batch`

Batch creates new shared drives using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The name of this shared drive |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `returnWhenReady` | int64 |  | The Google Drive API returns the drive after creation immediately, but usually before it can be used in subsequent re... |
| `returnWhenReady_ALL` | bool |  | Same as returnWhenReady but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `themeId` | int64 |  | The ID of the theme from which the background image and color will be set. The set of possible driveThemes can be ret... |
| `themeId_ALL` | string |  | Same as themeId but value is applied to all lines in the CSV file |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

#### `gsm drives delete`

Permanently deletes a shared drive for which the user is an organizer.
The shared drive cannot contain any untrashed items.

##### `gsm drives delete batch`

Batch deletes shared drives by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `driveId` | int64 |  | The ID of the shared drive |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

#### `gsm drives get`

Gets a shared drive's metadata by ID.

##### `gsm drives get batch`

Batch gets shared drives' metadata by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `driveId` | int64 |  | The ID of the shared drive |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

#### `gsm drives getSize`

Counts the files in a Shared Drive and returns their number and total size

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `driveId` | string | ✓ | The ID of the shared drive |
| `includeTrash` | bool |  | Whether to include trashed items. |

#### `gsm drives hide`

Hides a shared drive from the default view.

##### `gsm drives hide batch`

Batch hides shared drives using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `driveId` | int64 |  | The ID of the shared drive |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

#### `gsm drives list`

Lists the user's shared drives.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `q` | string |  | Query string for searching shared drives. See the https://developers.google.com/drive/api/v3/search-shareddrives for ... |
| `useDomainAdminAccess` | bool |  | Issue the request as a domain administrator |

#### `gsm drives unhide`

Restores a shared drive to the default view.

##### `gsm drives unhide batch`

Batch unhides shared drives using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `driveId` | int64 |  | The ID of the shared drive |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

#### `gsm drives update`

Updates the metadata for a shared drive.

##### `gsm drives update batch`

Batch updates shared drives using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `adminManagedRestrictions` | int64 |  | Whether administrative privileges on this shared drive are required to modify restrictions |
| `adminManagedRestrictions_ALL` | bool |  | Same as adminManagedRestrictions but value is applied to all lines in the CSV file |
| `backgroundImageFile` | int64 |  | An image file and cropping parameters from which a background image for this shared drive is set. This is a write onl... |
| `backgroundImageFile_ALL` | string |  | Same as backgroundImageFile but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `colorRgb` | int64 |  | The color of this shared drive as an RGB hex string. It can only be set on a drive.drives.update request that does no... |
| `colorRgb_ALL` | string |  | Same as colorRgb but value is applied to all lines in the CSV file |
| `copyRequiresWriterPermission` | int64 |  | Whether the options to copy, print, or download files inside this shared drive, should be disabled for readers and co... |
| `copyRequiresWriterPermission_ALL` | bool |  | Same as copyRequiresWriterPermission but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domainUsersOnly` | int64 |  | Whether access to this shared drive and items inside this shared drive is restricted to users of the domain to which ... |
| `domainUsersOnly_ALL` | bool |  | Same as domainUsersOnly but value is applied to all lines in the CSV file |
| `driveId` | int64 |  | The ID of the shared drive |
| `driveMembersOnly` | int64 |  | Whether access to items inside this shared drive is restricted to its members |
| `driveMembersOnly_ALL` | bool |  | Same as driveMembersOnly but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The name of this shared drive |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sharingFoldersRequiresOrganizerPermission` | int64 |  | If true, only users with the organizer role can share folders. If false, users with either the organizer role or the ... |
| `sharingFoldersRequiresOrganizerPermission_ALL` | bool |  | Same as sharingFoldersRequiresOrganizerPermission but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `themeId` | int64 |  | The ID of the theme from which the background image and color will be set. The set of possible driveThemes can be ret... |
| `themeId_ALL` | string |  | Same as themeId but value is applied to all lines in the CSV file |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

### `gsm entityUsageReports`

Manage (get) Entity Usage Reports (Part of Admin SDK API)

#### `gsm entityUsageReports get`

Retrieves a report which is a collection of properties and statistics for entities used by users within the account.
For more information, see the Entities Usage Report guide.
For more information about the entities report's parameters, see the Entities Usage parameters reference guides.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The unique ID of the customer to retrieve data for. |
| `date` | string |  | Represents the date the usage occurred. The timestamp is in the ISO 8601 format, yyyy-mm-dd. We recommend you use you... |
| `entityKey` | string |  | Represents the key of the object to filter the data with. Accepted values are: ALL         - Returns activity events ... |
| `entityType` | string |  | Represents the type of entity for the report. Accepted values are: GPLUS_COMMUNITIES  - Returns a report on Google+ c... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filters` | string |  | The filters query string is a comma-separated list of an application's event parameters where the parameter's value i... |
| `parameters` | string |  | The parameters query string is a comma-separated list of event parameters that refine a report's results. The paramet... |

### `gsm events`

Manage events in users' calendars (Part of Calendar API)

#### `gsm events delete`

Deletes an event.

##### `gsm events delete batch`

Batch deletes events using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `eventId` | int64 |  | Event identifier. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendUpdates` | int64 |  | Guests who should receive notifications about the event update (for example, title changes, etc.). [all\|externalOnly... |
| `sendUpdates_ALL` | string |  | Same as sendUpdates but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm events get`

Returns an event.

##### `gsm events get batch`

Batch returns events using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `eventId` | int64 |  | Event identifier. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `maxAttendees` | int64 |  | The maximum number of attendees to include in the response. If there are more than the specified number of attendees,... |
| `maxAttendees_ALL` | int64 |  | Same as maxAttendees but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `timeZone` | int64 |  | Time zone used in the response. Optional. The default is the time zone of the calendar. |
| `timeZone_ALL` | string |  | Same as timeZone but value is applied to all lines in the CSV file |

#### `gsm events import`

Imports an event.
This operation is used to add a private copy of an existing event to a calendar.

##### `gsm events import batch`

Batch imports events using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `conferenceDataVersion` | int64 |  | Version number of conference data supported by the API client. Version 0 assumes no conference data support and ignor... |
| `conferenceDataVersion_ALL` | int64 |  | Same as conferenceDataVersion but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `destination` | int64 |  | Calendar identifier of the target calendar. |
| `destination_ALL` | string |  | Same as destination but value is applied to all lines in the CSV file |
| `eventId` | int64 |  | Event identifier. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `supportsAttachments` | int64 |  | Whether API client performing operation supports event attachments. |
| `supportsAttachments_ALL` | bool |  | Same as supportsAttachments but value is applied to all lines in the CSV file |

#### `gsm events insert`

Creates an event.

##### `gsm events insert batch`

Batch inserts events using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addConferenceData` | int64 |  | Whether to add a Meet conference to the event. |
| `addConferenceData_ALL` | bool |  | Same as addConferenceData but value is applied to all lines in the CSV file |
| `anyoneCanAddSelf` | int64 |  | Whether anyone can invite themselves to the event (currently works for Google+ events only). |
| `anyoneCanAddSelf_ALL` | bool |  | Same as anyoneCanAddSelf but value is applied to all lines in the CSV file |
| `attendees` | int64 |  | Must be given in the following format: "--attendees "email=some.address@domain.com;resource=[true\|false];optional=[t... |
| `attendees_ALL` | stringSlice |  | Same as attendees but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `colorId` | int64 |  | The color of the event. This is an ID referring to an entry in the event section of the colors definition (see the co... |
| `colorId_ALL` | string |  | Same as colorId but value is applied to all lines in the CSV file |
| `conferenceDataVersion` | int64 |  | Version number of conference data supported by the API client. Version 0 assumes no conference data support and ignor... |
| `conferenceDataVersion_ALL` | int64 |  | Same as conferenceDataVersion but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | Description of the event. Can contain HTML. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `endDate` | int64 |  | The date, in the format "yyyy-mm-dd", if this is an all-day event. |
| `endDateTime` | int64 |  | The time, as a combined date-time value (formatted according to RFC3339). A time zone offset is required unless a tim... |
| `endDateTime_ALL` | string |  | Same as endDateTime but value is applied to all lines in the CSV file |
| `endDate_ALL` | string |  | Same as endDate but value is applied to all lines in the CSV file |
| `endTimeZone` | int64 |  | The time zone in which the time is specified. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".) F... |
| `endTimeZone_ALL` | string |  | Same as endTimeZone but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileUrl` | int64 |  | URL link to the attachment. For adding Google Drive file attachments use the same format as in alternateLink property... |
| `fileUrl_ALL` | stringSlice |  | Same as fileUrl but value is applied to all lines in the CSV file |
| `guestsCanInviteOthers` | int64 |  | Whether attendees other than the organizer can invite others to the event. |
| `guestsCanInviteOthers_ALL` | bool |  | Same as guestsCanInviteOthers but value is applied to all lines in the CSV file |
| `guestsCanModify` | int64 |  | Whether attendees other than the organizer can modify the event. |
| `guestsCanModify_ALL` | bool |  | Same as guestsCanModify but value is applied to all lines in the CSV file |
| `guestsCanSeeOtherGuests` | int64 |  | Whether attendees other than the organizer can see who the event's attendees are. |
| `guestsCanSeeOtherGuests_ALL` | bool |  | Same as guestsCanSeeOtherGuests but value is applied to all lines in the CSV file |
| `id` | int64 |  | Opaque identifier of the event. When creating new single or recurring events, you can specify their IDs. Provided IDs... |
| `location` | int64 |  | Geographic location of the event as free-form text. |
| `location_ALL` | string |  | Same as location but value is applied to all lines in the CSV file |
| `maxAttendees` | int64 |  | The maximum number of attendees to include in the response. If there are more than the specified number of attendees,... |
| `maxAttendees_ALL` | int64 |  | Same as maxAttendees but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `privateExtendedProperty` | int64 |  | Properties that are private to the copy of the event that appears on this calendar. |
| `privateExtendedProperty_ALL` | stringSlice |  | Same as privateExtendedProperty but value is applied to all lines in the CSV file |
| `recurrence` | int64 |  | List of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545. Note that DTSTART and D... |
| `recurrence_ALL` | stringSlice |  | Same as recurrence but value is applied to all lines in the CSV file |
| `reminderOverride` | int64 |  | If the event doesn't use the default reminders, this lists the reminders specific to the event, or, if not set, indic... |
| `reminderOverride_ALL` | stringSlice |  | Same as reminderOverride but value is applied to all lines in the CSV file |
| `sendUpdates` | int64 |  | Guests who should receive notifications about the event update (for example, title changes, etc.). [all\|externalOnly... |
| `sendUpdates_ALL` | string |  | Same as sendUpdates but value is applied to all lines in the CSV file |
| `sequence` | int64 |  | Sequence number as per iCalendar. |
| `sequence_ALL` | int64 |  | Same as sequence but value is applied to all lines in the CSV file |
| `sharedExtendedProperty` | int64 |  | Properties that are shared between copies of the event on other attendees' calendars. |
| `sharedExtendedProperty_ALL` | stringSlice |  | Same as sharedExtendedProperty but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `startDate` | int64 |  | The date, in the format "yyyy-mm-dd", if this is an all-day event. |
| `startDateTime` | int64 |  | The time, as a combined date-time value (formatted according to RFC3339). A time zone offset is required unless a tim... |
| `startDateTime_ALL` | string |  | Same as startDateTime but value is applied to all lines in the CSV file |
| `startDate_ALL` | string |  | Same as startDate but value is applied to all lines in the CSV file |
| `startTimeZone` | int64 |  | The time zone in which the time is specified. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".) F... |
| `startTimeZone_ALL` | string |  | Same as startTimeZone but value is applied to all lines in the CSV file |
| `status` | int64 |  | Status of the event. [confirmed\|tentative\|cancelled] confirmed  - The event is confirmed. This is the default statu... |
| `status_ALL` | string |  | Same as status but value is applied to all lines in the CSV file |
| `summary` | int64 |  | Title of the event. |
| `summary_ALL` | string |  | Same as summary but value is applied to all lines in the CSV file |
| `supportsAttachments` | int64 |  | Whether API client performing operation supports event attachments. |
| `supportsAttachments_ALL` | bool |  | Same as supportsAttachments but value is applied to all lines in the CSV file |
| `transparency` | int64 |  | Whether the event blocks time on the calendar. [opaque\|transparent] opaque       - Default value. The event does blo... |
| `transparency_ALL` | string |  | Same as transparency but value is applied to all lines in the CSV file |
| `useDefaultReminders` | int64 |  | Whether the default reminders of the calendar apply to the event. |
| `useDefaultReminders_ALL` | bool |  | Same as useDefaultReminders but value is applied to all lines in the CSV file |
| `visibility` | int64 |  | Visibility of the event. [default\|public\|private\|confidential] default       - Uses the default visibility for eve... |
| `visibility_ALL` | string |  | Same as visibility but value is applied to all lines in the CSV file |

#### `gsm events instances`

Returns instances of the specified recurring event.

##### `gsm events instances batch`

Batch returns instances of events using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `eventId` | int64 |  | Event identifier. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `maxAttendees` | int64 |  | The maximum number of attendees to include in the response. If there are more than the specified number of attendees,... |
| `maxAttendees_ALL` | int64 |  | Same as maxAttendees but value is applied to all lines in the CSV file |
| `originalStart` | int64 |  | The original start time of the instance in the result. |
| `originalStart_ALL` | string |  | Same as originalStart but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `showDeleted` | int64 |  | Whether to include deleted events (with status equals "cancelled") in the result. Cancelled instances of recurring ev... |
| `showDeleted_ALL` | bool |  | Same as showDeleted but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `timeMax` | int64 |  | Upper bound (exclusive) for an event's start time to filter by. Must be an RFC3339 timestamp with mandatory time zone... |
| `timeMax_ALL` | string |  | Same as timeMax but value is applied to all lines in the CSV file |
| `timeMin` | int64 |  | Lower bound (exclusive) for an event's end time to filter by. Must be an RFC3339 timestamp with mandatory time zone o... |
| `timeMin_ALL` | string |  | Same as timeMin but value is applied to all lines in the CSV file |
| `timeZone` | int64 |  | Time zone used in the response. Optional. The default is the time zone of the calendar. |
| `timeZone_ALL` | string |  | Same as timeZone but value is applied to all lines in the CSV file |

#### `gsm events list`

Returns events on the specified calendar.

##### `gsm events list batch`

Batch lists events using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `iCalUID` | int64 |  | ICalUID sets the optional parameter "iCalUID": Specifies event ID in the iCalendar format to be included in the respo... |
| `maxAttendees` | int64 |  | The maximum number of attendees to include in the response. If there are more than the specified number of attendees,... |
| `maxAttendees_ALL` | int64 |  | Same as maxAttendees but value is applied to all lines in the CSV file |
| `orderBy` | int64 |  | The order of the events returned in the result. Optional. The default is an unspecified, stable order. Acceptable val... |
| `orderBy_ALL` | string |  | Same as orderBy but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `privateExtendedProperty` | int64 |  | Properties that are private to the copy of the event that appears on this calendar. |
| `privateExtendedProperty_ALL` | stringSlice |  | Same as privateExtendedProperty but value is applied to all lines in the CSV file |
| `q` | int64 |  | Free text search terms to find events that match these terms in any field, except for extended properties. |
| `q_ALL` | string |  | Same as q but value is applied to all lines in the CSV file |
| `sharedExtendedProperty` | int64 |  | Properties that are shared between copies of the event on other attendees' calendars. |
| `sharedExtendedProperty_ALL` | stringSlice |  | Same as sharedExtendedProperty but value is applied to all lines in the CSV file |
| `showDeleted` | int64 |  | Whether to include deleted events (with status equals "cancelled") in the result. Cancelled instances of recurring ev... |
| `showDeleted_ALL` | bool |  | Same as showDeleted but value is applied to all lines in the CSV file |
| `showHiddenInvitations` | int64 |  | Whether to include hidden invitations in the result. |
| `showHiddenInvitations_ALL` | bool |  | Same as showHiddenInvitations but value is applied to all lines in the CSV file |
| `singleEvents` | int64 |  | Whether to expand recurring events into instances and only return single one-off events and instances of recurring ev... |
| `singleEvents_ALL` | bool |  | Same as singleEvents but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `timeMax` | int64 |  | Upper bound (exclusive) for an event's start time to filter by. Must be an RFC3339 timestamp with mandatory time zone... |
| `timeMax_ALL` | string |  | Same as timeMax but value is applied to all lines in the CSV file |
| `timeMin` | int64 |  | Lower bound (exclusive) for an event's end time to filter by. Must be an RFC3339 timestamp with mandatory time zone o... |
| `timeMin_ALL` | string |  | Same as timeMin but value is applied to all lines in the CSV file |
| `timeZone` | int64 |  | Time zone used in the response. Optional. The default is the time zone of the calendar. |
| `timeZone_ALL` | string |  | Same as timeZone but value is applied to all lines in the CSV file |
| `updatedMin` | int64 |  | Lower bound for an event's last modification time (as a RFC3339 timestamp) to filter by. When specified, entries dele... |
| `updatedMin_ALL` | string |  | Same as updatedMin but value is applied to all lines in the CSV file |

#### `gsm events move`

Moves an event to another calendar, i.e. changes an event's organizer.

##### `gsm events move batch`

Batch moves events to other calendars, i.e. changes an events' organizer using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `destination` | int64 |  | Calendar identifier of the target calendar. |
| `destination_ALL` | string |  | Same as destination but value is applied to all lines in the CSV file |
| `eventId` | int64 |  | Event identifier. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendUpdates` | int64 |  | Guests who should receive notifications about the event update (for example, title changes, etc.). [all\|externalOnly... |
| `sendUpdates_ALL` | string |  | Same as sendUpdates but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm events patch`

Updates an event. This method supports patch semantics.

##### `gsm events patch batch`

Batch patches events using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addConferenceData` | int64 |  | Whether to add a Meet conference to the event. |
| `addConferenceData_ALL` | bool |  | Same as addConferenceData but value is applied to all lines in the CSV file |
| `anyoneCanAddSelf` | int64 |  | Whether anyone can invite themselves to the event (currently works for Google+ events only). |
| `anyoneCanAddSelf_ALL` | bool |  | Same as anyoneCanAddSelf but value is applied to all lines in the CSV file |
| `attendees` | int64 |  | Must be given in the following format: "--attendees "email=some.address@domain.com;resource=[true\|false];optional=[t... |
| `attendeesOmitted` | int64 |  | Whether attendees may have been omitted from the event's representation. When retrieving an event, this may be due to... |
| `attendeesOmitted_ALL` | bool |  | Same as attendeesOmitted but value is applied to all lines in the CSV file |
| `attendees_ALL` | stringSlice |  | Same as attendees but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `colorId` | int64 |  | The color of the event. This is an ID referring to an entry in the event section of the colors definition (see the co... |
| `colorId_ALL` | string |  | Same as colorId but value is applied to all lines in the CSV file |
| `conferenceDataVersion` | int64 |  | Version number of conference data supported by the API client. Version 0 assumes no conference data support and ignor... |
| `conferenceDataVersion_ALL` | int64 |  | Same as conferenceDataVersion but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | Description of the event. Can contain HTML. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `endDate` | int64 |  | The date, in the format "yyyy-mm-dd", if this is an all-day event. |
| `endDateTime` | int64 |  | The time, as a combined date-time value (formatted according to RFC3339). A time zone offset is required unless a tim... |
| `endDateTime_ALL` | string |  | Same as endDateTime but value is applied to all lines in the CSV file |
| `endDate_ALL` | string |  | Same as endDate but value is applied to all lines in the CSV file |
| `endTimeZone` | int64 |  | The time zone in which the time is specified. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".) F... |
| `endTimeZone_ALL` | string |  | Same as endTimeZone but value is applied to all lines in the CSV file |
| `eventId` | int64 |  | Event identifier. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileUrl` | int64 |  | URL link to the attachment. For adding Google Drive file attachments use the same format as in alternateLink property... |
| `fileUrl_ALL` | stringSlice |  | Same as fileUrl but value is applied to all lines in the CSV file |
| `guestsCanInviteOthers` | int64 |  | Whether attendees other than the organizer can invite others to the event. |
| `guestsCanInviteOthers_ALL` | bool |  | Same as guestsCanInviteOthers but value is applied to all lines in the CSV file |
| `guestsCanModify` | int64 |  | Whether attendees other than the organizer can modify the event. |
| `guestsCanModify_ALL` | bool |  | Same as guestsCanModify but value is applied to all lines in the CSV file |
| `guestsCanSeeOtherGuests` | int64 |  | Whether attendees other than the organizer can see who the event's attendees are. |
| `guestsCanSeeOtherGuests_ALL` | bool |  | Same as guestsCanSeeOtherGuests but value is applied to all lines in the CSV file |
| `id` | int64 |  | Opaque identifier of the event. When creating new single or recurring events, you can specify their IDs. Provided IDs... |
| `location` | int64 |  | Geographic location of the event as free-form text. |
| `location_ALL` | string |  | Same as location but value is applied to all lines in the CSV file |
| `maxAttendees` | int64 |  | The maximum number of attendees to include in the response. If there are more than the specified number of attendees,... |
| `maxAttendees_ALL` | int64 |  | Same as maxAttendees but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `privateExtendedProperty` | int64 |  | Properties that are private to the copy of the event that appears on this calendar. |
| `privateExtendedProperty_ALL` | stringSlice |  | Same as privateExtendedProperty but value is applied to all lines in the CSV file |
| `recurrence` | int64 |  | List of RRULE, EXRULE, RDATE and EXDATE lines for a recurring event, as specified in RFC5545. Note that DTSTART and D... |
| `recurrence_ALL` | stringSlice |  | Same as recurrence but value is applied to all lines in the CSV file |
| `reminderOverride` | int64 |  | If the event doesn't use the default reminders, this lists the reminders specific to the event, or, if not set, indic... |
| `reminderOverride_ALL` | stringSlice |  | Same as reminderOverride but value is applied to all lines in the CSV file |
| `sendUpdates` | int64 |  | Guests who should receive notifications about the event update (for example, title changes, etc.). [all\|externalOnly... |
| `sendUpdates_ALL` | string |  | Same as sendUpdates but value is applied to all lines in the CSV file |
| `sequence` | int64 |  | Sequence number as per iCalendar. |
| `sequence_ALL` | int64 |  | Same as sequence but value is applied to all lines in the CSV file |
| `sharedExtendedProperty` | int64 |  | Properties that are shared between copies of the event on other attendees' calendars. |
| `sharedExtendedProperty_ALL` | stringSlice |  | Same as sharedExtendedProperty but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `startDate` | int64 |  | The date, in the format "yyyy-mm-dd", if this is an all-day event. |
| `startDateTime` | int64 |  | The time, as a combined date-time value (formatted according to RFC3339). A time zone offset is required unless a tim... |
| `startDateTime_ALL` | string |  | Same as startDateTime but value is applied to all lines in the CSV file |
| `startDate_ALL` | string |  | Same as startDate but value is applied to all lines in the CSV file |
| `startTimeZone` | int64 |  | The time zone in which the time is specified. (Formatted as an IANA Time Zone Database name, e.g. "Europe/Zurich".) F... |
| `startTimeZone_ALL` | string |  | Same as startTimeZone but value is applied to all lines in the CSV file |
| `status` | int64 |  | Status of the event. [confirmed\|tentative\|cancelled] confirmed  - The event is confirmed. This is the default statu... |
| `status_ALL` | string |  | Same as status but value is applied to all lines in the CSV file |
| `summary` | int64 |  | Title of the event. |
| `summary_ALL` | string |  | Same as summary but value is applied to all lines in the CSV file |
| `supportsAttachments` | int64 |  | Whether API client performing operation supports event attachments. |
| `supportsAttachments_ALL` | bool |  | Same as supportsAttachments but value is applied to all lines in the CSV file |
| `transparency` | int64 |  | Whether the event blocks time on the calendar. [opaque\|transparent] opaque       - Default value. The event does blo... |
| `transparency_ALL` | string |  | Same as transparency but value is applied to all lines in the CSV file |
| `useDefaultReminders` | int64 |  | Whether the default reminders of the calendar apply to the event. |
| `useDefaultReminders_ALL` | bool |  | Same as useDefaultReminders but value is applied to all lines in the CSV file |
| `visibility` | int64 |  | Visibility of the event. [default\|public\|private\|confidential] default       - Uses the default visibility for eve... |
| `visibility_ALL` | string |  | Same as visibility but value is applied to all lines in the CSV file |

#### `gsm events quickAdd`

Creates an event based on a simple text string.

##### `gsm events quickAdd batch`

Batch creates events based on simple text strings using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `calendarId` | int64 |  | Calendar identifier. To retrieve calendar IDs call the calendarList.list method. If you want to access the primary ca... |
| `calendarId_ALL` | string |  | Same as calendarId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendUpdates` | int64 |  | Guests who should receive notifications about the event update (for example, title changes, etc.). [all\|externalOnly... |
| `sendUpdates_ALL` | string |  | Same as sendUpdates but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `text` | int64 |  | The text describing the event to be created. |
| `text_ALL` | string |  | Same as text but value is applied to all lines in the CSV file |

### `gsm features`

Manage resource features (Part of Admin SDK API)

#### `gsm features delete`

Deletes a feature resource.

##### `gsm features delete batch`

Batch deletes feature resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `featureKey` | int64 |  | The unique ID of the feature. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm features get`

Gets a feature resource.

##### `gsm features get batch`

Batch retrieves feature resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `featureKey` | int64 |  | The unique ID of the feature. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm features insert`

Inserts a feature resource.

##### `gsm features insert batch`

Batch inserts feature resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The name of the feature. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm features list`

Retrieves a list of feature resources for an account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm features patch`

Patches a feature resource.

##### `gsm features patch batch`

Batch patches feature resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `featureKey` | int64 |  | The unique ID of the feature. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm features rename`

Renames a feature resource.

##### `gsm features rename batch`

Batch renames feature resources using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `newName` | int64 |  | New name of the feature. |
| `oldName` | int64 |  | The unique ID of the feature to rename. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm files`

Managed files (Part of Drive API)

#### `gsm files copy`

Creates a copy of a file and applies any requested updates with patch semantics.
Use "files copy recursive" to copy folders.

##### `gsm files copy batch`

Batch copies files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `appProperties` | int64 |  | A collection of arbitrary key-value pairs which are private to the requesting app. Entries with null values are clear... |
| `appProperties_ALL` | stringSlice |  | Same as appProperties but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `copyRequiresWriterPermission` | int64 |  | Whether the options to copy, print, or download this file, should be disabled for readers and commenters. |
| `copyRequiresWriterPermission_ALL` | bool |  | Same as copyRequiresWriterPermission but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | A short description of the file. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file |
| `id` | int64 |  | The ID of the file. |
| `ignoreDefaultVisibility` | int64 |  | Whether to ignore the domain's default visibility settings for the created file. Domain administrators can choose to ... |
| `ignoreDefaultVisibility_ALL` | bool |  | Same as ignoreDefaultVisibility but value is applied to all lines in the CSV file |
| `includePermissionsForView` | int64 |  | Specifies which additional view's permissions to include in the response. Only 'published' is supported. |
| `includePermissionsForView_ALL` | string |  | Same as includePermissionsForView but value is applied to all lines in the CSV file |
| `keepRevisionForever` | int64 |  | Whether to set the 'keepForever' field in the new head revision. This is only applicable to files with binary content... |
| `keepRevisionForever_ALL` | bool |  | Same as keepRevisionForever but value is applied to all lines in the CSV file |
| `mimeType` | int64 |  | The target MIME type of the file. Google Drive will attempt to automatically detect an appropriate value from uploade... |
| `mimeType_ALL` | string |  | Same as mimeType but value is applied to all lines in the CSV file |
| `modifiedTime` | int64 |  | The last time the file was modified by anyone (RFC 3339 date-time). Note that setting modifiedTime will also update m... |
| `modifiedTime_ALL` | string |  | Same as modifiedTime but value is applied to all lines in the CSV file |
| `name` | int64 |  | The name of the file. This is not necessarily unique within a folder. Note that for immutable items such as the top l... |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `ocrLanguage` | int64 |  | A language hint for OCR processing during image import (ISO 639-1 code). |
| `ocrLanguage_ALL` | string |  | Same as ocrLanguage but value is applied to all lines in the CSV file |
| `ownerRestricted` | int64 |  | Whether the content restriction can only be modified or removed by a user who owns the file. For files in shared driv... |
| `ownerRestricted_ALL` | bool |  | Same as ownerRestricted but value is applied to all lines in the CSV file |
| `parent` | int64 |  | The single parent of the file. |
| `parent_ALL` | string |  | Same as parent but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `properties` | int64 |  | A collection of arbitrary key-value pairs which are visible to all apps. Entries with null values are cleared in upda... |
| `properties_ALL` | stringSlice |  | Same as properties but value is applied to all lines in the CSV file |
| `readOnly` | int64 |  | Whether the content of the file is read-only. If a file is read-only, a new revision of the file may not be added, co... |
| `readOnlyReason` | int64 |  | Reason for why the content of the file is restricted. This is only mutable on requests that also set readOnly=true. |
| `readOnlyReason_ALL` | string |  | Same as readOnlyReason but value is applied to all lines in the CSV file |
| `readOnly_ALL` | bool |  | Same as readOnly but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `starred` | int64 |  | Whether the user has starred the file. |
| `starred_ALL` | bool |  | Same as starred but value is applied to all lines in the CSV file |
| `thumbnailImage` | int64 |  | The thumbnail data encoded with URL-safe Base64 (RFC 4648 section 5). |
| `thumbnailImage_ALL` | string |  | Same as thumbnailImage but value is applied to all lines in the CSV file |
| `thumbnailMimeType` | int64 |  | The MIME type of the thumbnail. |
| `thumbnailMimeType_ALL` | string |  | Same as thumbnailMimeType but value is applied to all lines in the CSV file |
| `viewedByMeTime` | int64 |  | The last time the file was viewed by the user (RFC 3339 date-time). |
| `viewedByMeTime_ALL` | string |  | Same as viewedByMeTime but value is applied to all lines in the CSV file |
| `writersCanShare` | int64 |  | Whether users with only writer permission can modify the file's permissions. Not populated for items in shared drives. |
| `writersCanShare_ALL` | bool |  | Same as writersCanShare but value is applied to all lines in the CSV file |

##### `gsm files copy recursive`

Recursively copies a folder to a new destination.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `parent` | string |  | The single parent of the file. |

#### `gsm files count`

Counts files in a folder and returns their number and size.

##### `gsm files count recursive`

Recursively count files in a folder

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |

#### `gsm files create`

Creates a new file or folder. Can also be used to upload files.

##### `gsm files create batch`

Batch creates new files or folders. Can also be used to upload files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `appProperties` | int64 |  | A collection of arbitrary key-value pairs which are private to the requesting app. Entries with null values are clear... |
| `appProperties_ALL` | stringSlice |  | Same as appProperties but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `copyRequiresWriterPermission` | int64 |  | Whether the options to copy, print, or download this file, should be disabled for readers and commenters. |
| `copyRequiresWriterPermission_ALL` | bool |  | Same as copyRequiresWriterPermission but value is applied to all lines in the CSV file |
| `createdTime` | int64 |  | The time at which the file was created (RFC 3339 date-time). |
| `createdTime_ALL` | string |  | Same as createdTime but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | A short description of the file. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `folderColorRgb` | int64 |  | The color for a folder as an RGB hex string. The supported colors are published in the folderColorPalette field of th... |
| `folderColorRgb_ALL` | string |  | Same as folderColorRgb but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the file. |
| `ignoreDefaultVisibility` | int64 |  | Whether to ignore the domain's default visibility settings for the created file. Domain administrators can choose to ... |
| `ignoreDefaultVisibility_ALL` | bool |  | Same as ignoreDefaultVisibility but value is applied to all lines in the CSV file |
| `includePermissionsForView` | int64 |  | Specifies which additional view's permissions to include in the response. Only 'published' is supported. |
| `includePermissionsForView_ALL` | string |  | Same as includePermissionsForView but value is applied to all lines in the CSV file |
| `indexableText` | int64 |  | Text to be indexed for the file to improve fullText queries. This is limited to 128KB in length and may contain HTML ... |
| `indexableText_ALL` | string |  | Same as indexableText but value is applied to all lines in the CSV file |
| `inheritedPermissionsDisabled` | int64 |  | Whether this file has inherited permissions disabled. Inherited permissions are enabled by default. See https://devel... |
| `inheritedPermissionsDisabled_ALL` | bool |  | Same as inheritedPermissionsDisabled but value is applied to all lines in the CSV file |
| `keepRevisionForever` | int64 |  | Whether to set the 'keepForever' field in the new head revision. This is only applicable to files with binary content... |
| `keepRevisionForever_ALL` | bool |  | Same as keepRevisionForever but value is applied to all lines in the CSV file |
| `localFilePath` | int64 |  | Path to a file or folder on the local disk. |
| `localFilePath_ALL` | string |  | Same as localFilePath but value is applied to all lines in the CSV file |
| `mimeType` | int64 |  | The target MIME type of the file. Google Drive will attempt to automatically detect an appropriate value from uploade... |
| `mimeType_ALL` | string |  | Same as mimeType but value is applied to all lines in the CSV file |
| `modifiedTime` | int64 |  | The last time the file was modified by anyone (RFC 3339 date-time). Note that setting modifiedTime will also update m... |
| `modifiedTime_ALL` | string |  | Same as modifiedTime but value is applied to all lines in the CSV file |
| `name` | int64 |  | The name of the file. This is not necessarily unique within a folder. Note that for immutable items such as the top l... |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `ocrLanguage` | int64 |  | A language hint for OCR processing during image import (ISO 639-1 code). |
| `ocrLanguage_ALL` | string |  | Same as ocrLanguage but value is applied to all lines in the CSV file |
| `originalFilename` | int64 |  | The original filename of the uploaded content if available, or else the original value of the name field. This is onl... |
| `originalFilename_ALL` | string |  | Same as originalFilename but value is applied to all lines in the CSV file |
| `ownerRestricted` | int64 |  | Whether the content restriction can only be modified or removed by a user who owns the file. For files in shared driv... |
| `ownerRestricted_ALL` | bool |  | Same as ownerRestricted but value is applied to all lines in the CSV file |
| `parent` | int64 |  | The single parent of the file. |
| `parent_ALL` | string |  | Same as parent but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `properties` | int64 |  | A collection of arbitrary key-value pairs which are visible to all apps. Entries with null values are cleared in upda... |
| `properties_ALL` | stringSlice |  | Same as properties but value is applied to all lines in the CSV file |
| `readOnly` | int64 |  | Whether the content of the file is read-only. If a file is read-only, a new revision of the file may not be added, co... |
| `readOnlyReason` | int64 |  | Reason for why the content of the file is restricted. This is only mutable on requests that also set readOnly=true. |
| `readOnlyReason_ALL` | string |  | Same as readOnlyReason but value is applied to all lines in the CSV file |
| `readOnly_ALL` | bool |  | Same as readOnly but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sourceMimeType` | int64 |  | The MIME type of the source file to upload. Set this to "text/csv" and "mimeType" to "application/vnd.google-apps.spr... |
| `sourceMimeType_ALL` | string |  | Same as sourceMimeType but value is applied to all lines in the CSV file |
| `starred` | int64 |  | Whether the user has starred the file. |
| `starred_ALL` | bool |  | Same as starred but value is applied to all lines in the CSV file |
| `targetId` | int64 |  | The ID of the file that this shortcut points to. |
| `targetId_ALL` | string |  | Same as targetId but value is applied to all lines in the CSV file |
| `thumbnailImage` | int64 |  | The thumbnail data encoded with URL-safe Base64 (RFC 4648 section 5). |
| `thumbnailImage_ALL` | string |  | Same as thumbnailImage but value is applied to all lines in the CSV file |
| `thumbnailMimeType` | int64 |  | The MIME type of the thumbnail. |
| `thumbnailMimeType_ALL` | string |  | Same as thumbnailMimeType but value is applied to all lines in the CSV file |
| `useContentAsIndexableText` | int64 |  | Whether to use the uploaded content as indexable text. |
| `useContentAsIndexableText_ALL` | bool |  | Same as useContentAsIndexableText but value is applied to all lines in the CSV file |
| `viewedByMeTime` | int64 |  | The last time the file was viewed by the user (RFC 3339 date-time). |
| `viewedByMeTime_ALL` | string |  | Same as viewedByMeTime but value is applied to all lines in the CSV file |
| `writersCanShare` | int64 |  | Whether users with only writer permission can modify the file's permissions. Not populated for items in shared drives. |
| `writersCanShare_ALL` | bool |  | Same as writersCanShare but value is applied to all lines in the CSV file |

#### `gsm files delete`

Permanently deletes a file owned by the user without moving it to the trash.
If the file belongs to a shared drive the user must be an organizer on the parent.
If the target is a folder, all descendants owned by the user are also deleted.

##### `gsm files delete batch`

Batch deletes files or folders by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fileId` | int64 |  | The ID of the file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm files download`

Download a file

##### `gsm files download batch`

Batch download files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `acknowledgeAbuse` | int64 |  | Whether the user is acknowledging the risk of downloading known malware or other abusive files. |
| `acknowledgeAbuse_ALL` | bool |  | Same as acknowledgeAbuse but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fileId` | int64 |  | The ID of the file |
| `localFilePath` | int64 |  | Path to a file or folder on the local disk. |
| `localFilePath_ALL` | string |  | Same as localFilePath but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm files export`

Exports a Google Doc to the requested MIME type and returns the exported content.

##### `gsm files export batch`

Batch export Google documents to the specified MIME type using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fileId` | int64 |  | The ID of the file |
| `localFilePath` | int64 |  | Path to a file or folder on the local disk. |
| `localFilePath_ALL` | string |  | Same as localFilePath but value is applied to all lines in the CSV file |
| `mimeType` | int64 |  | The target MIME type of the file. Google Drive will attempt to automatically detect an appropriate value from uploade... |
| `mimeType_ALL` | string |  | Same as mimeType but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm files generateIds`

Generates a set of file IDs which can be provided in create or copy requests.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `count` | int64 |  | The number of IDs to return. Acceptable values are 1 to 1000, inclusive. (Default: 10) |
| `space` | string |  | The space in which the IDs can be used to create new files. Supported values are 'drive' and 'appDataFolder'. |

#### `gsm files get`

Gets a file or folder's metadata or content by ID.

##### `gsm files get batch`

Batch gets files' or folders' metadata or content by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file |
| `includePermissionsForView` | int64 |  | Specifies which additional view's permissions to include in the response. Only 'published' is supported. |
| `includePermissionsForView_ALL` | string |  | Same as includePermissionsForView but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm files list`

Lists or searches files.

##### `gsm files list recursive`

Recursively list files in a folder

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |

#### `gsm files listLabels`

Lists the labels on a file.

##### `gsm files listLabels batch`

Batch list the labels on files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm files listLabels recursive`

Recursively lists labels on a folder and all of its children.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |

#### `gsm files modifyLabels`

Modifies the set of labels on a file.

##### `gsm files modifyLabels batch`

Batch modify the labels on files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file |
| `labelField` | int64 |  | A single label field that should be updated on a file. Can be used multiple times in the form of "--labelField "label... |
| `labelField_ALL` | stringSlice |  | Same as labelField but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm files modifyLabels recursive`

Recursively modifies the labels on a folder and all of its children.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `labelField` | stringSlice | ✓ | A single label field that should be updated on a file. Can be used multiple times in the form of "--labelField "label... |

#### `gsm files move`

Move a file.

##### `gsm files move batch`

Batch moves files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fileId` | int64 |  | The ID of the file |
| `parent` | int64 |  | The single parent of the file. |
| `parent_ALL` | string |  | Same as parent but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm files move recursive`

Moves a folder to a Shared Drive

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `parent` | string | ✓ | The single parent of the file. |

#### `gsm files removeLabels`

Removes labels from a file.

##### `gsm files removeLabels batch`

Batch remove the specified labels on files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file |
| `labelId` | int64 |  | The ID of a label that should be removed from the file Can be used multiple times to remove multiple labels in one re... |
| `labelId_ALL` | stringSlice |  | Same as labelId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm files removeLabels recursive`

Recursively removes the specified labels on a folder and all of its children.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `labelId` | stringSlice | ✓ | The ID of a label that should be removed from the file Can be used multiple times to remove multiple labels in one re... |

#### `gsm files update`

Updates a file's metadata and/or content. This method supports patch semantics.

##### `gsm files update batch`

Batch update files using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `appProperties` | int64 |  | A collection of arbitrary key-value pairs which are private to the requesting app. Entries with null values are clear... |
| `appProperties_ALL` | stringSlice |  | Same as appProperties but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `copyRequiresWriterPermission` | int64 |  | Whether the options to copy, print, or download this file, should be disabled for readers and commenters. |
| `copyRequiresWriterPermission_ALL` | bool |  | Same as copyRequiresWriterPermission but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | A short description of the file. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file |
| `folderColorRgb` | int64 |  | The color for a folder as an RGB hex string. The supported colors are published in the folderColorPalette field of th... |
| `folderColorRgb_ALL` | string |  | Same as folderColorRgb but value is applied to all lines in the CSV file |
| `includePermissionsForView` | int64 |  | Specifies which additional view's permissions to include in the response. Only 'published' is supported. |
| `includePermissionsForView_ALL` | string |  | Same as includePermissionsForView but value is applied to all lines in the CSV file |
| `indexableText` | int64 |  | Text to be indexed for the file to improve fullText queries. This is limited to 128KB in length and may contain HTML ... |
| `indexableText_ALL` | string |  | Same as indexableText but value is applied to all lines in the CSV file |
| `inheritedPermissionsDisabled` | int64 |  | Whether this file has inherited permissions disabled. Inherited permissions are enabled by default. See https://devel... |
| `inheritedPermissionsDisabled_ALL` | bool |  | Same as inheritedPermissionsDisabled but value is applied to all lines in the CSV file |
| `keepRevisionForever` | int64 |  | Whether to set the 'keepForever' field in the new head revision. This is only applicable to files with binary content... |
| `keepRevisionForever_ALL` | bool |  | Same as keepRevisionForever but value is applied to all lines in the CSV file |
| `localFilePath` | int64 |  | Path to a file or folder on the local disk. |
| `localFilePath_ALL` | string |  | Same as localFilePath but value is applied to all lines in the CSV file |
| `mimeType` | int64 |  | The target MIME type of the file. Google Drive will attempt to automatically detect an appropriate value from uploade... |
| `mimeType_ALL` | string |  | Same as mimeType but value is applied to all lines in the CSV file |
| `modifiedTime` | int64 |  | The last time the file was modified by anyone (RFC 3339 date-time). Note that setting modifiedTime will also update m... |
| `modifiedTime_ALL` | string |  | Same as modifiedTime but value is applied to all lines in the CSV file |
| `name` | int64 |  | The name of the file. This is not necessarily unique within a folder. Note that for immutable items such as the top l... |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `ocrLanguage` | int64 |  | A language hint for OCR processing during image import (ISO 639-1 code). |
| `ocrLanguage_ALL` | string |  | Same as ocrLanguage but value is applied to all lines in the CSV file |
| `originalFilename` | int64 |  | The original filename of the uploaded content if available, or else the original value of the name field. This is onl... |
| `originalFilename_ALL` | string |  | Same as originalFilename but value is applied to all lines in the CSV file |
| `ownerRestricted` | int64 |  | Whether the content restriction can only be modified or removed by a user who owns the file. For files in shared driv... |
| `ownerRestricted_ALL` | bool |  | Same as ownerRestricted but value is applied to all lines in the CSV file |
| `parent` | int64 |  | The single parent of the file. |
| `parent_ALL` | string |  | Same as parent but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `properties` | int64 |  | A collection of arbitrary key-value pairs which are visible to all apps. Entries with null values are cleared in upda... |
| `properties_ALL` | stringSlice |  | Same as properties but value is applied to all lines in the CSV file |
| `readOnly` | int64 |  | Whether the content of the file is read-only. If a file is read-only, a new revision of the file may not be added, co... |
| `readOnlyReason` | int64 |  | Reason for why the content of the file is restricted. This is only mutable on requests that also set readOnly=true. |
| `readOnlyReason_ALL` | string |  | Same as readOnlyReason but value is applied to all lines in the CSV file |
| `readOnly_ALL` | bool |  | Same as readOnly but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `starred` | int64 |  | Whether the user has starred the file. |
| `starred_ALL` | bool |  | Same as starred but value is applied to all lines in the CSV file |
| `thumbnailImage` | int64 |  | The thumbnail data encoded with URL-safe Base64 (RFC 4648 section 5). |
| `thumbnailImage_ALL` | string |  | Same as thumbnailImage but value is applied to all lines in the CSV file |
| `thumbnailMimeType` | int64 |  | The MIME type of the thumbnail. |
| `thumbnailMimeType_ALL` | string |  | Same as thumbnailMimeType but value is applied to all lines in the CSV file |
| `trashed` | int64 |  | Whether the file has been trashed, either explicitly or from a trashed parent folder. Only the owner may trash a file... |
| `trashed_ALL` | bool |  | Same as trashed but value is applied to all lines in the CSV file |
| `useContentAsIndexableText` | int64 |  | Whether to use the uploaded content as indexable text. |
| `useContentAsIndexableText_ALL` | bool |  | Same as useContentAsIndexableText but value is applied to all lines in the CSV file |
| `viewedByMeTime` | int64 |  | The last time the file was viewed by the user (RFC 3339 date-time). |
| `viewedByMeTime_ALL` | string |  | Same as viewedByMeTime but value is applied to all lines in the CSV file |
| `writersCanShare` | int64 |  | Whether users with only writer permission can modify the file's permissions. Not populated for items in shared drives. |
| `writersCanShare_ALL` | bool |  | Same as writersCanShare but value is applied to all lines in the CSV file |

### `gsm filters`

Manage users' Gmail message filters (Part of Gmail API)

#### `gsm filters create`

Creates a filter.

##### `gsm filters create batch`

Batch creates filters using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addLabelIds` | int64 |  | A list of IDs of labels to add to this message. Can be used multiple times. |
| `addLabelIds_ALL` | stringSlice |  | Same as addLabelIds but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `excludeChats` | int64 |  | Whether the response should exclude chats. |
| `excludeChats_ALL` | bool |  | Same as excludeChats but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `forward` | int64 |  | Email address that the message should be forwarded to. |
| `forward_ALL` | string |  | Same as forward but value is applied to all lines in the CSV file |
| `from` | int64 |  | The sender's display name or email address. |
| `from_ALL` | string |  | Same as from but value is applied to all lines in the CSV file |
| `hasAttachment` | int64 |  | Whether the message has any attachment. |
| `hasAttachment_ALL` | bool |  | Same as hasAttachment but value is applied to all lines in the CSV file |
| `negatedQuery` | int64 |  | Only return messages not matching the specified query. Supports the same query format as the Gmail search box. For ex... |
| `negatedQuery_ALL` | string |  | Same as negatedQuery but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `query` | int64 |  | Only return messages matching the specified query. Supports the same query format as the Gmail search box. For exampl... |
| `query_ALL` | string |  | Same as query but value is applied to all lines in the CSV file |
| `removeLabelIds` | int64 |  | A list of IDs of labels to remove from this message. Can be used multiple times. |
| `removeLabelIds_ALL` | stringSlice |  | Same as removeLabelIds but value is applied to all lines in the CSV file |
| `size` | int64 |  | The size of the entire RFC822 message in bytes, including all headers and attachments. |
| `sizeComparison` | int64 |  | How the message size in bytes should be in relation to the size field. "[SMALLER\|LARGER] SMALLER  - Find messages sm... |
| `sizeComparison_ALL` | string |  | Same as sizeComparison but value is applied to all lines in the CSV file |
| `size_ALL` | int64 |  | Same as size but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `subject` | int64 |  | Case-insensitive phrase found in the message's subject. Trailing and leading whitespace are be trimmed and adjacent s... |
| `subject_ALL` | string |  | Same as subject but value is applied to all lines in the CSV file |
| `to` | int64 |  | The recipient's display name or email address. Includes recipients in the "to", "cc", and "bcc" header fields. You ca... |
| `to_ALL` | string |  | Same as to but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm filters delete`

Deletes a filter.

##### `gsm filters delete batch`

Batch deletes filters using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `id` | int64 |  | The ID of the filter. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm filters get`

Gets a filter.

##### `gsm filters get batch`

Batch gets filters using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the filter. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm filters list`

Lists the message filters of a Gmail user.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

### `gsm forwardingAddresses`

Manage users' forwarding addresses (Part of Gmail API)

#### `gsm forwardingAddresses create`

Creates a forwarding address.
If ownership verification is required, a message will be sent to the recipient and the resource's verification status will be set to pending;
otherwise, the resource will be created with verification status set to accepted.

##### `gsm forwardingAddresses create batch`

Batch creates forwarding addresses using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `forwardingEmail` | int64 |  | An email address to which messages can be forwarded. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm forwardingAddresses delete`

Deletes the specified forwarding address and revokes any verification that may have been required.

##### `gsm forwardingAddresses delete batch`

Batch deletes the specified forwarding addresses using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `forwardingEmail` | int64 |  | An email address to which messages can be forwarded. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm forwardingAddresses get`

Gets the specified forwarding address.

##### `gsm forwardingAddresses get batch`

Batch gets the specified forwarding addresses using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `forwardingEmail` | int64 |  | An email address to which messages can be forwarded. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm forwardingAddresses list`

Lists the forwarding addresses for the specified account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

### `gsm freeBusy`

Query free/busy information (Part of Calendar API)

#### `gsm freeBusy query`

Returns free/busy information for a set of calendars.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `calendarExpansionMax` | int64 |  | Maximal number of calendars for which FreeBusy information is to be provided. Optional. Maximum value is 50. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupExpansionMax` | int64 |  | Maximal number of calendar identifiers to be provided for a single group. Optional. An error is returned for a group ... |
| `id` | stringSlice |  | The identifier of a calendar or a group. |
| `timeMax` | string | ✓ | The end of the interval for the query formatted as per RFC3339. |
| `timeMin` | string | ✓ | The start of the interval for the query formatted as per RFC3339. |
| `timeZone` | string |  | Time zone used in the response. Optional. The default is UTC. |

### `gsm gmailSettings`

Manage Gmail settings for users (Part of Gmail API)

#### `gsm gmailSettings getAutoForwarding`

Gets the auto-forwarding setting for the specified account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings getImap`

Gets IMAP settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings getLanguage`

Gets language settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings getPop`

Gets POP settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings getVacation`

Gets vacation responder settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings updateAutoForwarding`

Updates the auto-forwarding setting for the specified account.
A verified forwarding address must be specified when auto-forwarding is enabled.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `autoExpunge` | bool |  | If this value is true, Gmail will immediately expunge a message when it is marked as deleted in IMAP. Otherwise, Gmai... |
| `disposition` | string |  | The state that a message should be left in after it has been forwarded. [LEAVE_IN_INBOX\|ARCHIVE\|TRASH\|MARK_READ] L... |
| `emailAddress` | string |  | Email address to which all incoming messages are forwarded. This email address must be a verified member of the forwa... |
| `enabled` | bool |  | Whether the setting is enabled |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings updateImap`

Updates IMAP settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `autoExpunge` | bool |  | If this value is true, Gmail will immediately expunge a message when it is marked as deleted in IMAP. Otherwise, Gmai... |
| `enabled` | bool |  | Whether the setting is enabled |
| `expungeBehavior` | string |  | The action that will be executed on a message when it is marked as deleted and expunged from the last visible IMAP fo... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `maxFolderSize` | int64 |  | An optional limit on the number of messages that an IMAP folder may contain. Legal values are 0, 1000, 2000, 5000 or ... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings updateLanguage`

Updates language settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `displayLanguage` | string |  | The language to display Gmail in, formatted as an RFC 3066 Language Tag (for example en-GB, fr or ja for British Engl... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings updatePop`

Updates POP settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `accessWindow` | string |  | The range of messages which are accessible via POP. [DISABLED\|FROM_NOW_ON\|ALL_MAIL] DISABLED     - Indicates that n... |
| `disposition` | string |  | The state that a message should be left in after it has been forwarded. [LEAVE_IN_INBOX\|ARCHIVE\|TRASH\|MARK_READ] L... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm gmailSettings updateVacation`

Updates vacation responder settings.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `enableAutoReply` | bool |  | Flag that controls whether Gmail automatically replies to messages. |
| `endTime` | int64 |  | An optional end time for sending auto-replies (epoch ms). When this is specified, Gmail will automatically reply only... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `responseBodyHtml` | string |  | Response body in HTML format. Gmail will sanitize the HTML before storing it. If both responseBodyPlainText and respo... |
| `responseBodyPlainText` | string |  | Response body in plain text format. If both responseBodyPlainText and responseBodyHtml are specified, responseBodyHtm... |
| `responseSubject` | string |  | Optional text to prepend to the subject line in vacation responses. In order to enable auto-replies, either the respo... |
| `restrictToContacts` | bool |  | Flag that determines whether responses are sent to recipients who are not in the user's list of contacts. |
| `restrictToDomain` | bool |  | Flag that determines whether responses are sent to recipients who are outside of the user's domain. This feature is o... |
| `startTime` | int64 |  | An optional start time for sending auto-replies (epoch ms). When this is specified, Gmail will automatically reply on... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

### `gsm gmailUsers`

Gmail User Profiles (Part of Gmail API)

#### `gsm gmailUsers getProfile`

Gets the specified user's Gmail profile.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

### `gsm groupAliases`

Manage group aliases, which are alternative email addresses (Part of Admin SDK - not Gmail API!)

#### `gsm groupAliases delete`

Removes an alias.

##### `gsm groupAliases delete batch`

Batch deletes group aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `alias` | int64 |  | The alias. |
| `alias_ALL` | string |  | Same as alias but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupAliases insert`

Adds an alias for the group.

##### `gsm groupAliases insert batch`

Batch inserts group aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `alias` | int64 |  | The alias. |
| `alias_ALL` | string |  | Same as alias but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupAliases list`

Lists all aliases for a group.

##### `gsm groupAliases list batch`

Batch lists group aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm groupMembershipsCi`

Manage group memberships (Part of Cloud Identity API)

#### `gsm groupMembershipsCi checkTransitiveMembership`

Check a potential member for membership in a group.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `parent` | string |  | Resource name of the group. Format: groups/{group_id}, where group_id is the unique id assigned to the Group to which... |
| `query` | string | ✓ | A CEL expression that MUST include: getMembershipGraph      - member specification AND label(s)                      ... |

#### `gsm groupMembershipsCi create`

Creates a Membership.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `memberKeyId` | string | ✓ | The ID of the entity.  For Google-managed entities, the id must be the email address of an existing group or user.  F... |
| `memberKeyNamespace` | string |  | The namespace in which the entity exists.  If not specified, the EntityKey represents a Google-managed entity such as... |
| `parent` | string |  | Resource name of the group. Format: groups/{group_id}, where group_id is the unique id assigned to the Group to which... |
| `roles` | stringSlice | ✓ | The MembershipRoles that apply to the Membership.  Must not contain duplicate MembershipRoles with the same name.  Ca... |

#### `gsm groupMembershipsCi delete`

Deletes a Membership.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `name` | string |  | The resource name of the Membership. Must be of the form groups/{group_id}/memberships/{membership_id}. |

#### `gsm groupMembershipsCi get`

Gets a Membership.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string |  | The resource name of the Membership. Must be of the form groups/{group_id}/memberships/{membership_id}. |

#### `gsm groupMembershipsCi getMembershipGraph`

Get a membership graph of just a member or both a member and a group.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string |  | Resource name of the group. Format: groups/{group_id}, where group_id is the unique id assigned to the Group to which... |
| `query` | string | ✓ | A CEL expression that MUST include: getMembershipGraph      - member specification AND label(s)                      ... |

#### `gsm groupMembershipsCi list`

Lists the Memberships within a Group.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string |  | Resource name of the group. Format: groups/{group_id}, where group_id is the unique id assigned to the Group to which... |
| `view` | string |  | The level of detail to be returned. BASIC  - Default. Only basic resource information is returned. FULL   - All resou... |

#### `gsm groupMembershipsCi lookup`

Looks up the resource name of a Membership by its EntityKey.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `memberKeyId` | string | ✓ | The ID of the entity.  For Google-managed entities, the id must be the email address of an existing group or user.  F... |
| `memberKeyNamespace` | string |  | The namespace in which the entity exists.  If not specified, the EntityKey represents a Google-managed entity such as... |
| `parent` | string |  | Resource name of the group. Format: groups/{group_id}, where group_id is the unique id assigned to the Group to which... |

#### `gsm groupMembershipsCi modifyMembershipRoles`

Modifies the MembershipRoles of a Membership.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addRoles` | stringSlice |  | The MembershipRoles to be added.  Adding or removing roles in the same request as updating roles is not supported.  M... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string |  | The resource name of the Membership. Must be of the form groups/{group_id}/memberships/{membership_id}. |
| `removeRoles` | stringSlice |  | The names of the MembershipRoles to be removed.  Adding or removing roles in the same request as updating roles is no... |
| `updateRolesParams` | stringSlice |  | The MembershipRoles to be updated.  	Updating roles in the same request as adding or removing roles is not supported.... |

#### `gsm groupMembershipsCi searchTransitiveGroups`

Search transitive groups of a member.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `query` | string | ✓ | A CEL expression that MUST include: getMembershipGraph      - member specification AND label(s)                      ... |

#### `gsm groupMembershipsCi searchTransitiveMemberships`

Search transitive memberships of a group.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `email` | string |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string |  | Resource name of the group. Format: groups/{group_id}, where group_id is the unique id assigned to the Group to which... |

### `gsm groupSettings`

Manage Group Settings (Part of Admin SDK API)

#### `gsm groupSettings get`

Retrieves a group's settings identified by the group email address.

##### `gsm groupSettings get batch`

Batch retrieves groups' settings identified by the group email address using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupUniqueId` | int64 |  | The group's email address. |
| `ignoreDeprecated` | int64 |  | Ignore deprecated fields. |
| `ignoreDeprecated_ALL` | bool |  | Same as ignoreDeprecated but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupSettings patch`

Updates an existing resource. This method supports patch semantics.

##### `gsm groupSettings patch batch`

Batch patches groups' settings using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `allowExternalMembers` | int64 |  | Identifies whether members external to your organization can join the group. true   - Workspace users external to you... |
| `allowExternalMembers_ALL` | string |  | Same as allowExternalMembers but value is applied to all lines in the CSV file |
| `allowWebPosting` | int64 |  | Allows posting from web. true   - Allows any member to post to the group forum. false  - Members only use Gmail to co... |
| `allowWebPosting_ALL` | string |  | Same as allowWebPosting but value is applied to all lines in the CSV file |
| `archiveOnly` | int64 |  | Allows the group to be archived only. true   - Group is archived and the group is inactive. New messages to this grou... |
| `archiveOnly_ALL` | string |  | Same as archiveOnly but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customFooterText` | int64 |  | Set the content of custom footer text. The maximum number of characters is 1000. |
| `customFooterText_ALL` | string |  | Same as customFooterText but value is applied to all lines in the CSV file |
| `customReplyTo` | int64 |  | An email address used when replying to a message if the replyTo property is set to REPLY_TO_CUSTOM. This address is d... |
| `customReplyTo_ALL` | string |  | Same as customReplyTo but value is applied to all lines in the CSV file |
| `defaultMessageDenyNotificationText` | int64 |  | When a message is rejected, this is text for the rejection notification sent to the message's author. By default, thi... |
| `defaultMessageDenyNotificationText_ALL` | string |  | Same as defaultMessageDenyNotificationText but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `enableCollaborativeInbox` | int64 |  | Specifies whether a collaborative inbox will remain turned on for the group. |
| `enableCollaborativeInbox_ALL` | string |  | Same as enableCollaborativeInbox but value is applied to all lines in the CSV file |
| `favoriteRepliesOnTop` | int64 |  | Indicates if favorite replies should be displayed above other replies. true   - Favorite replies will be displayed ab... |
| `favoriteRepliesOnTop_ALL` | string |  | Same as favoriteRepliesOnTop but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupUniqueId` | int64 |  | The group's email address. |
| `ignoreDeprecated` | int64 |  | Ignore deprecated fields. |
| `ignoreDeprecated_ALL` | bool |  | Same as ignoreDeprecated but value is applied to all lines in the CSV file |
| `includeCustomFooter` | int64 |  | Whether to include custom footer. |
| `includeCustomFooter_ALL` | string |  | Same as includeCustomFooter but value is applied to all lines in the CSV file |
| `includeInGlobalAddressList` | int64 |  | Enables the group to be included in the Global Address List. For more information, see the help center. true   - Grou... |
| `includeInGlobalAddressList_ALL` | string |  | Same as includeInGlobalAddressList but value is applied to all lines in the CSV file |
| `isArchived` | int64 |  | Allows the Group contents to be archived. true   - Archive messages sent to the group. false  - Do not keep an archiv... |
| `isArchived_ALL` | string |  | Same as isArchived but value is applied to all lines in the CSV file |
| `membersCanPostAsTheGroup` | int64 |  | Enables members to post messages as the group. true   - Group member can post messages using the group's email addres... |
| `membersCanPostAsTheGroup_ALL` | string |  | Same as membersCanPostAsTheGroup but value is applied to all lines in the CSV file |
| `messageModerationLevel` | int64 |  | Moderation level of incoming messages. [MODERATE_ALL_MESSAGES\|MODERATE_NON_MEMBERS\|MODERATE_NEW_MEMBERS\|MODERATE_N... |
| `messageModerationLevel_ALL` | string |  | Same as messageModerationLevel but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `primaryLanguage` | int64 |  | The primary language for group. For a group's primary language use the language tags from the Workspace languages fou... |
| `primaryLanguage_ALL` | string |  | Same as primaryLanguage but value is applied to all lines in the CSV file |
| `replyTo` | int64 |  | Specifies who the default reply should go to. [REPLY_TO_CUSTOM\|REPLY_TO_SENDER\|REPLY_TO_LIST\|REPLY_TO_OWNER\|REPLY... |
| `replyTo_ALL` | string |  | Same as replyTo but value is applied to all lines in the CSV file |
| `sendMessageDenyNotification` | int64 |  | Allows a member to be notified if the member's message to the group is denied by the group owner. true   - When a mes... |
| `sendMessageDenyNotification_ALL` | string |  | Same as sendMessageDenyNotification but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `spamModerationLevel` | int64 |  | Specifies moderation levels for messages detected as spam. [ALLOW\|MODERATE\|SILENTLY_MODERATE\|REJECT] ALLOW        ... |
| `spamModerationLevel_ALL` | string |  | Same as spamModerationLevel but value is applied to all lines in the CSV file |
| `whoCanApproveMembers` | int64 |  | Specifies who can approve members who ask to join groups. This permission will be deprecated once it is merged into t... |
| `whoCanApproveMembers_ALL` | string |  | Same as whoCanApproveMembers but value is applied to all lines in the CSV file |
| `whoCanAssistContent` | int64 |  | Specifies who can moderate metadata. [ALL_MEMBERS\|OWNERS_AND_MANAGERS\|OWNERS_ONLY\|NONE] |
| `whoCanAssistContent_ALL` | string |  | Same as whoCanAssistContent but value is applied to all lines in the CSV file |
| `whoCanBanUsers` | int64 |  | Specifies who can deny membership to users. This permission will be deprecated once it is merged into the new whoCanM... |
| `whoCanBanUsers_ALL` | string |  | Same as whoCanBanUsers but value is applied to all lines in the CSV file |
| `whoCanContactOwner` | int64 |  | Specifies who can contact the group owner. [ALL_IN_DOMAIN_CAN_CONTACT\|ALL_MANAGERS_CAN_CONTACT\|ALL_MEMBERS_CAN_CONT... |
| `whoCanContactOwner_ALL` | string |  | Same as whoCanContactOwner but value is applied to all lines in the CSV file |
| `whoCanDiscoverGroup` | int64 |  | Specifies the set of users for whom this group is discoverable. [ANYONE_CAN_DISCOVER\|ALL_IN_DOMAIN_CAN_DISCOVER\|ALL... |
| `whoCanDiscoverGroup_ALL` | string |  | Same as whoCanDiscoverGroup but value is applied to all lines in the CSV file |
| `whoCanJoin` | int64 |  | Permission to join group. [ANYONE_CAN_JOIN\|ALL_IN_DOMAIN_CAN_JOIN\|INVITED_CAN_JOIN\|CAN_REQUEST_TO_JOIN] ANYONE_CAN... |
| `whoCanJoin_ALL` | string |  | Same as whoCanJoin but value is applied to all lines in the CSV file |
| `whoCanLeaveGroup` | int64 |  | Specifies who can leave the group. [ALL_MANAGERS_CAN_LEAVE\|ALL_MEMBERS_CAN_LEAVE\|NONE_CAN_LEAVE] |
| `whoCanLeaveGroup_ALL` | string |  | Same as whoCanLeaveGroup but value is applied to all lines in the CSV file |
| `whoCanModerateContent` | int64 |  | Specifies who can moderate content. [ALL_MEMBERS\|OWNERS_AND_MANAGERS\|OWNERS_ONLY\|NONE] |
| `whoCanModerateContent_ALL` | string |  | Same as whoCanModerateContent but value is applied to all lines in the CSV file |
| `whoCanModerateMembers` | int64 |  | Specifies who can manage members. [ALL_MEMBERS\|OWNERS_AND_MANAGERS\|OWNERS_ONLY\|NONE] |
| `whoCanModerateMembers_ALL` | string |  | Same as whoCanModerateMembers but value is applied to all lines in the CSV file |
| `whoCanPostMessage` | int64 |  | Permissions to post messages. [NONE_CAN_POST\|ALL_MANAGERS_CAN_POST\|ALL_MEMBERS_CAN_POST\|ALL_OWNERS_CAN_POST\|ALL_I... |
| `whoCanPostMessage_ALL` | string |  | Same as whoCanPostMessage but value is applied to all lines in the CSV file |
| `whoCanViewGroup` | int64 |  | Permissions to view group messages. [ANYONE_CAN_VIEW\|ALL_IN_DOMAIN_CAN_VIEW\|ALL_MEMBERS_CAN_VIEW\|ALL_OWNERS_CAN_VI... |
| `whoCanViewGroup_ALL` | string |  | Same as whoCanViewGroup but value is applied to all lines in the CSV file |
| `whoCanViewMembership` | int64 |  | Permissions to view membership. [ALL_IN_DOMAIN_CAN_VIEW\|ALL_MEMBERS_CAN_VIEW\|ALL_MANAGERS_CAN_VIEW] ALL_IN_DOMAIN_C... |
| `whoCanViewMembership_ALL` | string |  | Same as whoCanViewMembership but value is applied to all lines in the CSV file |

### `gsm groups`

Implements the groups API (Part of Admin SDK API).

#### `gsm groups delete`

Deletes a group.

##### `gsm groups delete batch`

Batch deletes groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groups get`

Retrieves a group's properties.

##### `gsm groups get batch`

Batch retrieves groups' properties using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groups insert`

Creates a group.

##### `gsm groups insert batch`

Batch inserts groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | An extended description to help users determine the purpose of a group. For example, you can include information abou... |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `email` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The group's display name. |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groups list`

Retrieve all groups of a domain or of a user given a userKey (paginated).

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | The unique ID for the customer's Workspace account. In case of a multi-domain account, to fetch all groups for a cust... |
| `domain` | string |  | The domain name. Use this field to get fields from only one domain. To return all domains for a customer account, use... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `orderBy` | string |  | Column to use for sorting results Acceptable values are: email  - Email of the group. |
| `query` | string |  | Query string search. Should be of the form "". Complete documentation is at https://developers.google.com/admin-sdk/d... |
| `sortOrder` | string |  | Whether to return results in ascending or descending order. Only of use when orderBy is also used Acceptable values a... |
| `userKey` | string |  | Email or immutable ID of the user if only those groups are to be listed, the given user is a member of. If it's an ID... |

#### `gsm groups patch`

Updates a group's properties. This method supports patch semantics

##### `gsm groups patch batch`

Batch patches groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | An extended description to help users determine the purpose of a group. For example, you can include information abou... |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `email` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `name` | int64 |  | The group's display name. |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm groupsCi`

Manage Google Groups with the Cloud Identity API

#### `gsm groupsCi create`

Creates a Group.

##### `gsm groupsCi create batch`

Batch creates groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | An extended description to help users determine the purpose of a Group. Must not be longer than 4,096 characters. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `displayName` | int64 |  | The display name of the Group. |
| `displayName_ALL` | string |  | Same as displayName but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the entity.  For Google-managed entities, the id must be the email address.  For external-identity-mapped e... |
| `initialGroupConfig` | int64 |  | Required. The initial configuration option for the Group. WITH_INITIAL_OWNER  - The end user making the request will ... |
| `initialGroupConfig_ALL` | string |  | Same as initialGroupConfig but value is applied to all lines in the CSV file |
| `labels` | int64 |  |  One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.  Go... |
| `labels_ALL` | stringSlice |  | Same as labels but value is applied to all lines in the CSV file |
| `namespace` | int64 |  | The namespace in which the entity exists.  If not specified, the EntityKey represents a Google-managed entity such as... |
| `namespace_ALL` | string |  | Same as namespace but value is applied to all lines in the CSV file |
| `parent` | int64 |  | Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_... |
| `parent_ALL` | string |  | Same as parent but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `queries` | int64 |  | Memberships will be the union of all queries. Only one entry with USER resource is currently supported. Can be used m... |
| `queries_ALL` | stringArray |  | Same as queries but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupsCi delete`

Deletes a Group.

##### `gsm groupsCi delete batch`

Batch deletes groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `email` | int64 |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `name` | int64 |  | The resource name of the Group.  Must be of the form groups/{group_id}. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupsCi get`

Retrieves a group.

##### `gsm groupsCi get batch`

Batch retrieves groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `email` | int64 |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The resource name of the Group.  Must be of the form groups/{group_id}. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupsCi getSecuritySettings`

Retrieves the security settings (member restrictions) of a group.

##### `gsm groupsCi getSecuritySettings batch`

Batch retrieves groups' security settings (member restrictions) using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `email` | int64 |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The resource name of the Group.  Must be of the form groups/{group_id}. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `readMask` | int64 |  | Field-level read mask of which fields to return. "*" returns all fields.  If not specified, all fields will be return... |
| `readMask_ALL` | string |  | Same as readMask but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupsCi list`

Lists the Groups under a customer or namespace.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string |  | Must be of the form identitysources/{identity_source_id} for external- identity-mapped groups or customers/{customer_... |
| `view` | string |  | The level of detail to be returned. BASIC  - Default. Only basic resource information is returned. FULL   - All resou... |

#### `gsm groupsCi lookup`

Looks up a Group.

##### `gsm groupsCi lookup batch`

Batch looks up groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the entity.  For Google-managed entities, the id must be the email address.  For external-identity-mapped e... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm groupsCi patch`

Updates a Group.

##### `gsm groupsCi patch batch`

Batch patches groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `email` | int64 |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `labels` | int64 |  |  One or more label entries that apply to the Group. Currently supported labels contain a key with an empty value.  Go... |
| `labels_ALL` | stringSlice |  | Same as labels but value is applied to all lines in the CSV file |
| `name` | int64 |  | The resource name of the Group.  Must be of the form groups/{group_id}. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `updateMask` | int64 |  | The fully-qualified names of fields to update.  May only contain the following fields: - patch:   - displayName   - d... |
| `updateMask_ALL` | string |  | Same as updateMask but value is applied to all lines in the CSV file |

#### `gsm groupsCi search`

Searches for Groups matching a specified query.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `query` | string |  | Must be specified in Common Expression Language. search: May only contain equality operators on the parent and inclus... |
| `view` | string |  | The level of detail to be returned. BASIC  - Default. Only basic resource information is returned. FULL   - All resou... |

#### `gsm groupsCi updateSecuritySettings`

Updates the security settings (member restrictions) of a group.

##### `gsm groupsCi updateSecuritySettings batch`

Batch retrieves groups' security settings (member restrictions) using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `email` | int64 |  | Email address of the group. This may be used instead of the name to do a lookup of the group resource name. Note that... |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The resource name of the Group.  Must be of the form groups/{group_id}. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `query` | int64 |  | Must be specified in Common Expression Language. search: May only contain equality operators on the parent and inclus... |
| `query_ALL` | string |  | Same as query but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `updateMask` | int64 |  | The fully-qualified names of fields to update.  May only contain the following fields: - patch:   - displayName   - d... |
| `updateMask_ALL` | string |  | Same as updateMask but value is applied to all lines in the CSV file |

### `gsm history`

Manage (list..) user's mailbox History (Part of Gmail API)

#### `gsm history list`

Lists the history of all changes to the given mailbox. History results are returned in chronological order (increasing historyId).

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `historyTypes` | stringSlice |  | History types to be returned by the function. [MESSAGE_ADDED\|MESSAGE_DELETED\|LABEL_ADDED\|LABEL_REMOVED] |
| `labelId` | string |  | Only return messages with a label matching the ID. |
| `startHistoryId` | uint64 | ✓ | Required. Returns history records after the specified startHistoryId. The supplied startHistoryId should be obtained ... |
| `userId` | string |  | The user's email address. The special value me can be used to indicate the authenticated user. |

### `gsm labels`

Manage users' mailbox labels (Part of Gmail API)

#### `gsm labels create`

Creates a new label.

##### `gsm labels create batch`

Batch creates new labels using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `backgroundColor` | int64 |  | Background color |
| `backgroundColor_ALL` | string |  | Same as backgroundColor but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `labelListVisibility` | int64 |  | The visibility of the label in the label list in the Gmail web interface. [LABEL_SHOW\|LABEL_SHOW_IF_UNREAD\|LABEL_HIDE] |
| `labelListVisibility_ALL` | string |  | Same as labelListVisibility but value is applied to all lines in the CSV file |
| `messageListVisibility` | int64 |  | The visibility of messages with this label in the message list in the Gmail web interface. [SHOW\|HIDE] |
| `messageListVisibility_ALL` | string |  | Same as messageListVisibility but value is applied to all lines in the CSV file |
| `name` | int64 |  | The display name of the label. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `textColor` | int64 |  | Text color |
| `textColor_ALL` | string |  | Same as textColor but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm labels delete`

Immediately and permanently deletes the specified label and removes it from any messages and threads that it is applied to.

##### `gsm labels delete batch`

Batch deletes the specified labels using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `id` | int64 |  | The ID of the label |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm labels get`

Gets the specified label.

##### `gsm labels get batch`

Batch gets the specified labels using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the label |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm labels list`

Lists all labels in the user's mailbox.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |

#### `gsm labels patch`

Patch the specified label.

##### `gsm labels patch batch`

Batch patches labels using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `backgroundColor` | int64 |  | Background color |
| `backgroundColor_ALL` | string |  | Same as backgroundColor but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the label |
| `labelListVisibility` | int64 |  | The visibility of the label in the label list in the Gmail web interface. [LABEL_SHOW\|LABEL_SHOW_IF_UNREAD\|LABEL_HIDE] |
| `labelListVisibility_ALL` | string |  | Same as labelListVisibility but value is applied to all lines in the CSV file |
| `messageListVisibility` | int64 |  | The visibility of messages with this label in the message list in the Gmail web interface. [SHOW\|HIDE] |
| `messageListVisibility_ALL` | string |  | Same as messageListVisibility but value is applied to all lines in the CSV file |
| `name` | int64 |  | The display name of the label. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `textColor` | int64 |  | Text color |
| `textColor_ALL` | string |  | Same as textColor but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value "me" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

### `gsm licenseAssignments`

Manage user license assignments (Part of Enterprise License Manager API)

#### `gsm licenseAssignments delete`

Delete a specific user's license by product SKU.

##### `gsm licenseAssignments delete batch`

Batch deletes user license assignments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `productId` | int64 |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `productId_ALL` | string |  | Same as productId but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `skuId` | int64 |  | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |
| `skuId_ALL` | string |  | Same as skuId but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

##### `gsm licenseAssignments delete recursive`

Removes licenses from users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `productId` | string |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `skuId` | string | ✓ | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |

#### `gsm licenseAssignments get`

Get a specific user's license by product SKU.

##### `gsm licenseAssignments get batch`

Batch gets user license assignments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `productId` | int64 |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `productId_ALL` | string |  | Same as productId but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `skuId` | int64 |  | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |
| `skuId_ALL` | string |  | Same as skuId but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

##### `gsm licenseAssignments get recursive`

Get users' licenses by product SKU by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `productId` | string |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `skuId` | string | ✓ | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |

#### `gsm licenseAssignments insert`

Assign a license.

##### `gsm licenseAssignments insert batch`

Batch inserts user license assignments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `productId` | int64 |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `productId_ALL` | string |  | Same as productId but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `skuId` | int64 |  | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |
| `skuId_ALL` | string |  | Same as skuId but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

##### `gsm licenseAssignments insert recursive`

Assigns licenses to users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `productId` | string |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `skuId` | string | ✓ | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |

#### `gsm licenseAssignments listForProduct`

List all users assigned licenses for a specific product SKU.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `fields` | string |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `productId` | string |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |

#### `gsm licenseAssignments listForProductAndSku`

List all users assigned licenses for a specific product SKU.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `fields` | string |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `productId` | string |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `skuId` | string | ✓ | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |

#### `gsm licenseAssignments patch`

Reassign a user's product SKU with a different SKU in the same product. This method supports patch semantics.

##### `gsm licenseAssignments patch batch`

Patch patches users' license assignments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `productId` | int64 |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `productId_ALL` | string |  | Same as productId but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `skuId` | int64 |  | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |
| `skuIdNew` | int64 |  | The product's new unique identifier. For more information about products in this version of the API, see https://deve... |
| `skuIdNew_ALL` | string |  | Same as skuIdNew but value is applied to all lines in the CSV file |
| `skuId_ALL` | string |  | Same as skuId but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

##### `gsm licenseAssignments patch recursive`

Patch users' license assignments by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | The user's current primary email address. If the user's email address changes, use the new email address in your API ... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `productId` | string |  | A product's unique identifier. For more information about products in this version of the API, see https://developers... |
| `skuId` | string | ✓ | A product SKU's unique identifier. For more information about available SKUs in this version of the API, see https://... |
| `skuIdNew` | string | ✓ | The product's new unique identifier. For more information about products in this version of the API, see https://deve... |

### `gsm log`

Manage GSM Logs

#### `gsm log clear`

Clears the current log.

*No command-specific flags.*

#### `gsm log show`

Shows the current log.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `lines` | int64 |  | Number of lines to return |

### `gsm members`

Manage group members (Part of Admin SDK API)

#### `gsm members delete`

Removes a member from a group.

##### `gsm members delete batch`

Batch deletes group members using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `groupKey_ALL` | string |  | Same as groupKey but value is applied to all lines in the CSV file |
| `memberKey` | int64 |  | Identifies the group member in the API request. A group member can be a user or another group. The value can be the m... |
| `memberKey_ALL` | string |  | Same as memberKey but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm members delete recursive`

Removes users from a group by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `groupKey` | string | ✓ | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm members get`

Retrieves a group member's properties.

##### `gsm members get batch`

Batch retrieves group members using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `groupKey_ALL` | string |  | Same as groupKey but value is applied to all lines in the CSV file |
| `memberKey` | int64 |  | Identifies the group member in the API request. A group member can be a user or another group. The value can be the m... |
| `memberKey_ALL` | string |  | Same as memberKey but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm members get recursive`

Retrieves members of a group by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `groupKey` | string | ✓ | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm members hasMember`

Checks whether the given user is a member of the group. Membership can be direct or nested.

##### `gsm members hasMember batch`

Batch checks whether users are members of groups using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `groupKey_ALL` | string |  | Same as groupKey but value is applied to all lines in the CSV file |
| `memberKey` | int64 |  | Identifies the group member in the API request. A group member can be a user or another group. The value can be the m... |
| `memberKey_ALL` | string |  | Same as memberKey but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm members hasMember recursive`

Checks whether users are members of a group by referencing one or more organizational units and/or groups. Membership can be direct or nested.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `groupKey` | string | ✓ | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm members insert`

Adds a user or group to the specified group.

##### `gsm members insert batch`

Batch inserts members using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `delivery_settings` | int64 |  | Defines mail delivery preferences of member. Acceptable values are: ALL_MAIL  - All messages, delivered as soon as th... |
| `delivery_settings_ALL` | string |  | Same as delivery_settings but value is applied to all lines in the CSV file |
| `email` | int64 |  | The member's email address. A member can be a user or another group. This property is required when adding a member t... |
| `email_ALL` | string |  | Same as email but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `groupKey_ALL` | string |  | Same as groupKey but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `role` | int64 |  | The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a m... |
| `role_ALL` | string |  | Same as role but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm members insert recursive`

Adds users to a group by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `delivery_settings` | string |  | Defines mail delivery preferences of member. Acceptable values are: ALL_MAIL  - All messages, delivered as soon as th... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `groupKey` | string | ✓ | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `role` | string |  | The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a m... |

#### `gsm members list`

Retrieves a paginated list of all members in a group.

##### `gsm members list batch`

Batch lists group members using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `groupKey_ALL` | string |  | Same as groupKey but value is applied to all lines in the CSV file |
| `includeDerivedMembership` | int64 |  | Whether to list indirect memberships. |
| `includeDerivedMembership_ALL` | bool |  | Same as includeDerivedMembership but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roles` | int64 |  | The roles query parameter allows you to retrieve group members by role. Allowed values are OWNER, MANAGER, and MEMBER. |
| `roles_ALL` | string |  | Same as roles but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm members patch`

Updates the membership properties of a user in the specified group. This method supports patch semantics.

##### `gsm members patch batch`

Batch patches members using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `delivery_settings` | int64 |  | Defines mail delivery preferences of member. Acceptable values are: ALL_MAIL  - All messages, delivered as soon as th... |
| `delivery_settings_ALL` | string |  | Same as delivery_settings but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `groupKey` | int64 |  | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `groupKey_ALL` | string |  | Same as groupKey but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `role` | int64 |  | The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a m... |
| `role_ALL` | string |  | Same as role but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm members patch recursive`

Updates the membership properties of users in a group by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `delivery_settings` | string |  | Defines mail delivery preferences of member. Acceptable values are: ALL_MAIL  - All messages, delivered as soon as th... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `groupKey` | string | ✓ | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `role` | string |  | The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a m... |

#### `gsm members set`

Sets the members of a group to match the specified email addresses with the given role

##### `gsm members set recursive`

Sets the memberships of the group by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `delivery_settings` | string |  | Defines mail delivery preferences of member. Acceptable values are: ALL_MAIL  - All messages, delivered as soon as th... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `groupKey` | string | ✓ | Identifies the group in the API request. The value can be the group's email address, group alias, or the unique group... |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `role` | string |  | The member's role in a group. The API returns an error for cycles in group memberships. For example, if group1 is a m... |

### `gsm messages`

Manage users' messages (Part of Gmail API)

#### `gsm messages batchDelete`

Deletes many messages by message ID. Provides no guarantees that messages were not already deleted or even existed at all.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `ids` | stringSlice | ✓ | The IDs of the messages. There is a limit of 1000 ids per request. |

#### `gsm messages delete`

Immediately and permanently deletes the specified message. This operation cannot be undone. Prefer messages.trash instead.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `id` | string |  | The ID of the message. |
| `userId` | string |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |

#### `gsm messages get`

Gets the specified message.

##### `gsm messages get batch`

Batch gets the specified messages using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `format` | int64 |  | The format to return the message in. [MINIMAL\|FULL\|RAW\|METADATA] MINIMAL   - Returns only email message ID and lab... |
| `format_ALL` | string |  | Same as format but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the message. |
| `metadataHeaders` | int64 |  | When given and format is METADATA, only include headers specified. |
| `metadataHeaders_ALL` | string |  | Same as metadataHeaders but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm messages import`

Imports a message into only this user's mailbox, with standard email delivery scanning and classification similar to receiving via SMTP.
Does not send a message.

##### `gsm messages import batch`

Batch imports messages using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `deleted` | int64 |  | Mark the email as permanently deleted (not TRASH) and only visible in Google Vault to a Vault administrator. Only use... |
| `deleted_ALL` | bool |  | Same as deleted but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `eml` | int64 |  | Path to the local .eml file |
| `eml_ALL` | string |  | Same as eml but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `internalDateSource` | int64 |  | Source for Gmail's internal date of the message. [DATE_HEADER\|RECEIVED_TIME] |
| `internalDateSource_ALL` | string |  | Same as internalDateSource but value is applied to all lines in the CSV file |
| `neverMarkSpam` | int64 |  | Ignore the Gmail spam classifier decision and never mark this email as SPAM in the mailbox. |
| `neverMarkSpam_ALL` | bool |  | Same as neverMarkSpam but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `processForCalendar` | int64 |  | Process calendar invites in the email and add any extracted meetings to the Google Calendar for this user. |
| `processForCalendar_ALL` | bool |  | Same as processForCalendar but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm messages insert`

Directly inserts a message into only this user's mailbox similar to IMAP APPEND, bypassing most scanning and classification.
Does not send a message.

##### `gsm messages insert batch`

Batch inserts messages using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `deleted` | int64 |  | Mark the email as permanently deleted (not TRASH) and only visible in Google Vault to a Vault administrator. Only use... |
| `deleted_ALL` | bool |  | Same as deleted but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `eml` | int64 |  | Path to the local .eml file |
| `eml_ALL` | string |  | Same as eml but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `internalDateSource` | int64 |  | Source for Gmail's internal date of the message. [DATE_HEADER\|RECEIVED_TIME] |
| `internalDateSource_ALL` | string |  | Same as internalDateSource but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm messages list`

Lists the messages in the user's mailbox.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `includeSpamTrash` | bool |  | Include messages from SPAM and TRASH in the results. |
| `labelIds` | stringSlice |  | Only return messages with labels that match all of the specified label IDs. |
| `q` | string |  | Only return messages matching the specified query. Supports the same query format as the Gmail search box. For exampl... |
| `userId` | string |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |

#### `gsm messages modify`

Modifies the labels on the specified message.

##### `gsm messages modify batch`

Batch modifies messages using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addLabelIds` | int64 |  | A list of label IDs to add to messages. |
| `addLabelIds_ALL` | stringSlice |  | Same as addLabelIds but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the message. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `removeLabelIds` | int64 |  | A list of label IDs to remove from messages. |
| `removeLabelIds_ALL` | stringSlice |  | Same as removeLabelIds but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm messages send`

Sends the specified message to the recipients in the To, Cc, and Bcc headers.

##### `gsm messages send batch`

Batch sends messages using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `attachment` | int64 |  | Path to a file that should be attached to the message. Can be used multiple times. |
| `attachment_ALL` | stringSlice |  | Same as attachment but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `bcc` | int64 |  | Blind Copy (Bcc) |
| `bcc_ALL` | string |  | Same as bcc but value is applied to all lines in the CSV file |
| `body` | int64 |  | Body or content of the (draft) message |
| `body_ALL` | string |  | Same as body but value is applied to all lines in the CSV file |
| `cc` | int64 |  | Copy (Cc) |
| `cc_ALL` | string |  | Same as cc but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `from` | int64 |  | Sender of the (draft) message. Must be a valid sendAs address. If this is not set, your primary sendAs address will b... |
| `from_ALL` | string |  | Same as from but value is applied to all lines in the CSV file |
| `html` | int64 |  | Send the body as HTML |
| `html_ALL` | bool |  | Same as html but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `subject` | int64 |  | Subject of the (draft) message |
| `subject_ALL` | string |  | Same as subject but value is applied to all lines in the CSV file |
| `to` | int64 |  | Recipient of the (draft) message |
| `to_ALL` | string |  | Same as to but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm messages trash`

Moves the specified message to the trash.

##### `gsm messages trash batch`

Batch trashes messages using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the message. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm messages untrash`

Removes the specified message from the trash.

##### `gsm messages untrash batch`

Batch untrashes messages using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The ID of the message. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value \"me\" can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

### `gsm mobileDevices`

Manage Mobile Devices (Part of Admin SDK API)

#### `gsm mobileDevices action`

Takes an action that affects a mobile device. For example, remotely wiping a device.

##### `gsm mobileDevices action batch`

Batch applies actions on mobile devices using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `action` | int64 |  | The action to be performed on the device. [admin_account_wipe\|admin_remote_wipe\|approve\|approve\|block\|cancel_rem... |
| `action_ALL` | string |  | Same as action but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceId` | int64 |  | The unique ID the API service uses to identify the mobile device. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm mobileDevices delete`

Removes a mobile device.

##### `gsm mobileDevices delete batch`

Batch deletes mobile devices using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceId` | int64 |  | The unique ID the API service uses to identify the mobile device. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm mobileDevices get`

Retrieves a mobile device's properties.

##### `gsm mobileDevices get batch`

Batch retrieves mobile devices' properties using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `projection` | int64 |  | Restrict information returned to a set of selected fields. Acceptable values are: BASIC  - Includes only the basic me... |
| `projection_ALL` | string |  | Same as projection but value is applied to all lines in the CSV file |
| `resourceId` | int64 |  | The unique ID the API service uses to identify the mobile device. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm mobileDevices list`

Retrieves a paginated list of all mobile devices for an account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `orderBy` | string |  | Device property to use for sorting results. Acceptable values are: deviceId  - The serial number for a Google Sync mo... |
| `projection` | string |  | Restrict information returned to a set of selected fields. Acceptable values are: BASIC  - Includes only the basic me... |
| `query` | string |  | Search string in the format provided by List query operators. See https://developers.google.com/admin-sdk/directory/v... |
| `sortOrder` | string |  | Whether to return results in ascending or descending order. Must be used with the orderBy parameter. Acceptable value... |

### `gsm orgUnits`

Manage Organizational Unit (Part of Admin SDK API)

#### `gsm orgUnits delete`

Removes an organizational unit.

##### `gsm orgUnits delete batch`

Batch deletes organizational units using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `orgUnitPath` | int64 |  | The full path of the organizational unit or its unique ID. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm orgUnits get`

Retrieves an organizational unit.

##### `gsm orgUnits get batch`

Batch retrieves organizational units using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `orgUnitPath` | int64 |  | The full path of the organizational unit or its unique ID. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm orgUnits insert`

Adds an organizational unit.

##### `gsm orgUnits insert batch`

Batch inserts organizational units using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `blockInheritance` | int64 |  | Determines if a sub-organizational unit can inherit the settings of the parent organization. The default value is fal... |
| `blockInheritance_ALL` | bool |  | Same as blockInheritance but value is applied to all lines in the CSV file |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | Description of the organizational unit. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The organizational unit's path name. For example, an organizational unit's name within the /corp/support/sales_suppor... |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `parentOrgUnitId` | int64 |  | The unique ID of the parent organizational unit. Required, unless parentOrgUnitPath is set. |
| `parentOrgUnitId_ALL` | string |  | Same as parentOrgUnitId but value is applied to all lines in the CSV file |
| `parentOrgUnitPath` | int64 |  | The organizational unit's parent path. For example, /corp/sales is the parent path for /corp/sales/sales_support orga... |
| `parentOrgUnitPath_ALL` | string |  | Same as parentOrgUnitPath but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `type` | int64 |  | Whether to return all sub-organizations or just immediate children. Acceptable values are: all       - All sub-organi... |
| `type_ALL` | string |  | Same as type but value is applied to all lines in the CSV file |

#### `gsm orgUnits list`

Retrieves a list of all organizational units for an account.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `orgUnitPath` | string |  | The full path of the organizational unit or its unique ID. |
| `type` | string |  | Whether to return all sub-organizations or just immediate children. Acceptable values are: all       - All sub-organi... |

#### `gsm orgUnits patch`

Updates an organizational unit. This method supports patch semantics.

##### `gsm orgUnits patch batch`

Batch patches organizational units using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `blockInheritance` | int64 |  | Determines if a sub-organizational unit can inherit the settings of the parent organization. The default value is fal... |
| `blockInheritance_ALL` | bool |  | Same as blockInheritance but value is applied to all lines in the CSV file |
| `customerId` | int64 |  | The unique ID for the customer's Workspace account. As an account administrator, you can also use the my_customer ali... |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `description` | int64 |  | Description of the organizational unit. |
| `description_ALL` | string |  | Same as description but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | The organizational unit's path name. For example, an organizational unit's name within the /corp/support/sales_suppor... |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `orgUnitPath` | int64 |  | The full path of the organizational unit or its unique ID. |
| `parentOrgUnitId` | int64 |  | The unique ID of the parent organizational unit. Required, unless parentOrgUnitPath is set. |
| `parentOrgUnitId_ALL` | string |  | Same as parentOrgUnitId but value is applied to all lines in the CSV file |
| `parentOrgUnitPath` | int64 |  | The organizational unit's parent path. For example, /corp/sales is the parent path for /corp/sales/sales_support orga... |
| `parentOrgUnitPath_ALL` | string |  | Same as parentOrgUnitPath but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `type` | int64 |  | Whether to return all sub-organizations or just immediate children. Acceptable values are: all       - All sub-organi... |
| `type_ALL` | string |  | Same as type but value is applied to all lines in the CSV file |

### `gsm orgUnitsMemberships`

Manage the memberships of Shared Drives in organizational units (OUs) (Part of Cloud Identity Beta API)

#### `gsm orgUnitsMemberships list`

List OrgMembership resources in an OrgUnit treated as 'parent'.

##### `gsm orgUnitsMemberships list batch`

Batch list Shared Drives in organizational units using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Customer that this OrgMembership belongs to. All authorization will happen on the role assignments of this customer. ... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `filter` | int64 |  | The search query. Must be specified in Common Expression Language. May only contain equality operators on the type (e... |
| `filter_ALL` | string |  | Same as filter but value is applied to all lines in the CSV file |
| `parent` | int64 |  | OrgUnit which is queried for a list of memberships. Format: orgUnits/{$orgUnitId} where $orgUnitId is the orgUnitId f... |
| `parent_ALL` | string |  | Same as parent but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm orgUnitsMemberships move`

Move an OrgMembership to a new OrgUnit.
NOTE: This is an atomic copy-and-delete.
The resource will have a new copy under the destination OrgUnit and be deleted from the source OrgUnit.
The resource can only be searched under the destination OrgUnit afterwards.

##### `gsm orgUnitsMemberships move batch`

Batch move Shared Drives to organizational units using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Customer that this OrgMembership belongs to. All authorization will happen on the role assignments of this customer. ... |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `destinationOrgUnit` | int64 |  | OrgUnit where the membership will be moved to. Format: orgUnits/{$orgUnitId} where $orgUnitId is the orgUnitId from t... |
| `destinationOrgUnit_ALL` | string |  | Same as destinationOrgUnit but value is applied to all lines in the CSV file |
| `driveId` | int64 |  | The driveId of the Shared Drive to be moved. Use this instead of the name, if you just want to specify the driveId. G... |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Use "driveId" instead if you just want to specify the driveId! The resource name of the OrgMembership. Format: orgUni... |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm otherContacts`

Manage 'other' contacts (Part of People API)

#### `gsm otherContacts copyOtherContactToMyContactsGroup`

Copies an "Other contact" to a new contact in the user's "myContacts" group.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `copyMask` | string | ✓ | A field mask to restrict which fields are copied into the new contact. Valid values are:   - emailAddresses   - names... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `readMask` | string |  | A field mask to restrict which fields on the person are returned. Multiple fields can be specified by separating them... |
| `resourceName` | string | ✓ | The resource name of the "Other contact". |
| `sources` | stringSlice |  | A mask of what source types to return. READ_SOURCE_TYPE_PROFILE         - Returns SourceType.ACCOUNT, SourceType.DOMA... |

#### `gsm otherContacts list`

List all "Other contacts", that is contacts that are not in a contact group.
"Other contacts" are typically auto created contacts from interactions.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `readMask` | string | ✓ | A field mask to restrict which fields on the person are returned. Multiple fields can be specified by separating them... |

### `gsm people`

Manage people's contacts (Part of People API)

#### `gsm people createContact`

Create a new contact and return the person resource for that contact.

##### `gsm people createContact batch`

Batch create contacts using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addressMeAs` | int64 |  | The type of pronouns that should be used to address the person. The value can be custom or one of these predefined va... |
| `addressMeAs_ALL` | string |  | Same as addressMeAs but value is applied to all lines in the CSV file |
| `addresses` | int64 |  | A person's physical address. May be a P.O. box or street address. All fields are optional. May be used multiple times... |
| `addresses_ALL` | stringSlice |  | Same as addresses but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `biographyContentType` | int64 |  | The content type of the biography. CONTENT_TYPE_UNSPECIFIED  - Unspecified. TEXT_PLAIN                - Plain text. T... |
| `biographyContentType_ALL` | string |  | Same as biographyContentType but value is applied to all lines in the CSV file |
| `biographyValue` | int64 |  | The short biography. |
| `biographyValue_ALL` | string |  | Same as biographyValue but value is applied to all lines in the CSV file |
| `birthdayDay` | int64 |  | Day of month. Must be from 1 to 31 and valid for the year and month, or 0 if specifying a year by itself or a year an... |
| `birthdayDay_ALL` | int64 |  | Same as birthdayDay but value is applied to all lines in the CSV file |
| `birthdayMonth` | int64 |  | Month of year. Must be from 1 to 12, or 0 if specifying a year without a month and day. |
| `birthdayMonth_ALL` | int64 |  | Same as birthdayMonth but value is applied to all lines in the CSV file |
| `birthdayText` | int64 |  | A free-form string representing the user's birthday. |
| `birthdayText_ALL` | string |  | Same as birthdayText but value is applied to all lines in the CSV file |
| `birthdayYear` | int64 |  | Year of date. Must be from 1 to 9999, or 0 if specifying a date without a year. |
| `birthdayYear_ALL` | int64 |  | Same as birthdayYear but value is applied to all lines in the CSV file |
| `calendarUrls` | int64 |  | The person's calendar URLs. Can be used multiple times in the form of "url=...,type=...", etc. You may use the follow... |
| `calendarUrls_ALL` | stringSlice |  | Same as calendarUrls but value is applied to all lines in the CSV file |
| `clientData` | int64 |  | The person's client data. Arbitrary client data that is populated by clients. Duplicate keys and values are allowed. ... |
| `clientData_ALL` | stringSlice |  | Same as clientData but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `emailAddresses` | int64 |  | The person's email addresses. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the... |
| `emailAddresses_ALL` | stringSlice |  | Same as emailAddresses but value is applied to all lines in the CSV file |
| `events` | int64 |  | The person's events. Can be used multiple times in the form of "year=...,month=...", etc. You may use the following f... |
| `events_ALL` | stringSlice |  | Same as events but value is applied to all lines in the CSV file |
| `externalIds` | int64 |  | The person's external IDs. Can be used multiple times in the form of "value=...,type=...", etc. You may use the follo... |
| `externalIds_ALL` | stringSlice |  | Same as externalIds but value is applied to all lines in the CSV file |
| `familyName` | int64 |  | The family name. |
| `familyName_ALL` | string |  | Same as familyName but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileAses` | int64 |  | The person's file-ases. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `fileAses_ALL` | stringSlice |  | Same as fileAses but value is applied to all lines in the CSV file |
| `genderValue` | int64 |  | The gender for the person. The gender can be custom or one of these predefined values:   - male   - female   - unspec... |
| `genderValue_ALL` | string |  | Same as genderValue but value is applied to all lines in the CSV file |
| `givenName` | int64 |  | The given name. |
| `givenName_ALL` | string |  | Same as givenName but value is applied to all lines in the CSV file |
| `honorificPrefix` | int64 |  | The honorific prefixes, such as Mrs. or Dr. |
| `honorificPrefix_ALL` | string |  | Same as honorificPrefix but value is applied to all lines in the CSV file |
| `honorificSuffix` | int64 |  | The honorific suffixes, such as Jr. |
| `honorificSuffix_ALL` | string |  | Same as honorificSuffix but value is applied to all lines in the CSV file |
| `imClients` | int64 |  | The person's instant messaging clients. Can be used multiple times in the form of "primary=...,username=...", etc. Yo... |
| `imClients_ALL` | stringSlice |  | Same as imClients but value is applied to all lines in the CSV file |
| `interests` | int64 |  | The person's interests. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `interests_ALL` | stringSlice |  | Same as interests but value is applied to all lines in the CSV file |
| `locales` | int64 |  | The person's locale preferences. Can be used multiple times in the form of "primary=...,value=...", etc. You may use ... |
| `locales_ALL` | stringSlice |  | Same as locales but value is applied to all lines in the CSV file |
| `locations` | int64 |  | The person's locations. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `locations_ALL` | stringSlice |  | Same as locations but value is applied to all lines in the CSV file |
| `memberships` | int64 |  | The person's group memberships. Can be used multiple times in the form of "primary=...,contactGroupResourceName=...",... |
| `memberships_ALL` | stringSlice |  | Same as memberships but value is applied to all lines in the CSV file |
| `middleName` | int64 |  | The middle name(s). |
| `middleName_ALL` | string |  | Same as middleName but value is applied to all lines in the CSV file |
| `miscKeywords` | int64 |  | The person's miscellaneous keywords. Can be used multiple times in the form of "primary=...,value=...", etc. You may ... |
| `miscKeywords_ALL` | stringSlice |  | Same as miscKeywords but value is applied to all lines in the CSV file |
| `nicknames` | int64 |  | The person's nicknames. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `nicknames_ALL` | stringSlice |  | Same as nicknames but value is applied to all lines in the CSV file |
| `occupations` | int64 |  | The person's occupations. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the fol... |
| `occupations_ALL` | stringSlice |  | Same as occupations but value is applied to all lines in the CSV file |
| `organizations` | int64 |  | The person's past or current organizations. Can be used multiple times in the form of "primary=...,type=...", etc. Yo... |
| `organizations_ALL` | stringSlice |  | Same as organizations but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `personFields` | int64 |  | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `personFields_ALL` | string |  | Same as personFields but value is applied to all lines in the CSV file |
| `phoneNumbers` | int64 |  | The person's phone numbers. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the f... |
| `phoneNumbers_ALL` | stringSlice |  | Same as phoneNumbers but value is applied to all lines in the CSV file |
| `phoneticFamilyName` | int64 |  | The family name spelled as it sounds. |
| `phoneticFamilyName_ALL` | string |  | Same as phoneticFamilyName but value is applied to all lines in the CSV file |
| `phoneticFullName` | int64 |  | The full name spelled as it sounds. |
| `phoneticFullName_ALL` | string |  | Same as phoneticFullName but value is applied to all lines in the CSV file |
| `phoneticGivenName` | int64 |  | The given name spelled as it sounds. |
| `phoneticGivenName_ALL` | string |  | Same as phoneticGivenName but value is applied to all lines in the CSV file |
| `phoneticHonorificPrefix` | int64 |  | The honorific prefixes spelled as they sound. |
| `phoneticHonorificPrefix_ALL` | string |  | Same as phoneticHonorificPrefix but value is applied to all lines in the CSV file |
| `phoneticHonorificSuffix` | int64 |  | The honorific suffixes spelled as they sound. |
| `phoneticHonorificSuffix_ALL` | string |  | Same as phoneticHonorificSuffix but value is applied to all lines in the CSV file |
| `phoneticMiddleName` | int64 |  | The middle name(s) spelled as they sound. |
| `phoneticMiddleName_ALL` | string |  | Same as phoneticMiddleName but value is applied to all lines in the CSV file |
| `relations` | int64 |  | The person's relations. Can be used multiple times in the form of "primary=...,person=...", etc. You may use the foll... |
| `relations_ALL` | stringSlice |  | Same as relations but value is applied to all lines in the CSV file |
| `sipAddresses` | int64 |  | The person's SIP addresses. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the f... |
| `sipAddresses_ALL` | stringSlice |  | Same as sipAddresses but value is applied to all lines in the CSV file |
| `skills` | int64 |  | The person's skills. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the followin... |
| `skills_ALL` | stringSlice |  | Same as skills but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sources` | int64 |  | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |
| `sources_ALL` | string |  | Same as sources but value is applied to all lines in the CSV file |
| `unstructuredName` | int64 |  | The free form name value. |
| `unstructuredName_ALL` | string |  | Same as unstructuredName but value is applied to all lines in the CSV file |
| `urls` | int64 |  | The person's associated URLs. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the... |
| `urls_ALL` | stringSlice |  | Same as urls but value is applied to all lines in the CSV file |
| `userDefined` | int64 |  | The person's user defined data. Can be used multiple times in the form of "primary=...,type=...", etc. You may use th... |
| `userDefined_ALL` | stringSlice |  | Same as userDefined but value is applied to all lines in the CSV file |

#### `gsm people deleteContact`

Delete a contact person. Any non-contact data will not be deleted.

##### `gsm people deleteContact batch`

Batch deletes contacts using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `resourceName` | int64 |  | The resource name of the contact- |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm people deleteContactPhoto`

Delete a contact's photo.

##### `gsm people deleteContactPhoto batch`

Batch deletes contact photos using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `personFields` | int64 |  | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `personFields_ALL` | string |  | Same as personFields but value is applied to all lines in the CSV file |
| `resourceName` | int64 |  | The resource name of the contact- |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sources` | int64 |  | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |
| `sources_ALL` | string |  | Same as sources but value is applied to all lines in the CSV file |

#### `gsm people get`

Provides information about a person by specifying a resource name.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `personFields` | string | ✓ | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `resourceName` | string | ✓ | The resource name of the contact- |
| `sources` | string |  | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |

#### `gsm people getBatchGet`

Provides information about a list of specific people by specifying a list of requested resource names.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `personFields` | string | ✓ | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `resourceNames` | stringSlice | ✓ | The resource names of the people to provide information about. It's repeatable. The URL query parameter should be  re... |
| `sources` | string |  | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |

#### `gsm people listDirectoryPeople`

Provides a list of domain profiles and domain contacts in the authenticated user's domain directory.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `mergeSources` | stringSlice |  | Additional data to merge into the directory sources if they are connected through verified join keys such as email ad... |
| `readMask` | string | ✓ | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `sources` | string | ✓ | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |

#### `gsm people searchDirectoryPeople`

Provides a list of domain profiles and domain contacts in the authenticated user's domain directory that match the search query.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `mergeSources` | stringSlice |  | Additional data to merge into the directory sources if they are connected through verified join keys such as email ad... |
| `query` | string | ✓ | Prefix query that matches fields in the person. Does NOT use the readMask for determining what fields to match. |
| `readMask` | string | ✓ | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `sources` | string | ✓ | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |

#### `gsm people updateContact`

Update contact data for an existing contact person.

##### `gsm people updateContact batch`

Batch update contacts using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addressMeAs` | int64 |  | The type of pronouns that should be used to address the person. The value can be custom or one of these predefined va... |
| `addressMeAs_ALL` | string |  | Same as addressMeAs but value is applied to all lines in the CSV file |
| `addresses` | int64 |  | A person's physical address. May be a P.O. box or street address. All fields are optional. May be used multiple times... |
| `addresses_ALL` | stringSlice |  | Same as addresses but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `biographyContentType` | int64 |  | The content type of the biography. CONTENT_TYPE_UNSPECIFIED  - Unspecified. TEXT_PLAIN                - Plain text. T... |
| `biographyContentType_ALL` | string |  | Same as biographyContentType but value is applied to all lines in the CSV file |
| `biographyValue` | int64 |  | The short biography. |
| `biographyValue_ALL` | string |  | Same as biographyValue but value is applied to all lines in the CSV file |
| `birthdayDay` | int64 |  | Day of month. Must be from 1 to 31 and valid for the year and month, or 0 if specifying a year by itself or a year an... |
| `birthdayDay_ALL` | int64 |  | Same as birthdayDay but value is applied to all lines in the CSV file |
| `birthdayMonth` | int64 |  | Month of year. Must be from 1 to 12, or 0 if specifying a year without a month and day. |
| `birthdayMonth_ALL` | int64 |  | Same as birthdayMonth but value is applied to all lines in the CSV file |
| `birthdayText` | int64 |  | A free-form string representing the user's birthday. |
| `birthdayText_ALL` | string |  | Same as birthdayText but value is applied to all lines in the CSV file |
| `birthdayYear` | int64 |  | Year of date. Must be from 1 to 9999, or 0 if specifying a date without a year. |
| `birthdayYear_ALL` | int64 |  | Same as birthdayYear but value is applied to all lines in the CSV file |
| `calendarUrls` | int64 |  | The person's calendar URLs. Can be used multiple times in the form of "url=...,type=...", etc. You may use the follow... |
| `calendarUrls_ALL` | stringSlice |  | Same as calendarUrls but value is applied to all lines in the CSV file |
| `clientData` | int64 |  | The person's client data. Arbitrary client data that is populated by clients. Duplicate keys and values are allowed. ... |
| `clientData_ALL` | stringSlice |  | Same as clientData but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `emailAddresses` | int64 |  | The person's email addresses. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the... |
| `emailAddresses_ALL` | stringSlice |  | Same as emailAddresses but value is applied to all lines in the CSV file |
| `events` | int64 |  | The person's events. Can be used multiple times in the form of "year=...,month=...", etc. You may use the following f... |
| `events_ALL` | stringSlice |  | Same as events but value is applied to all lines in the CSV file |
| `externalIds` | int64 |  | The person's external IDs. Can be used multiple times in the form of "value=...,type=...", etc. You may use the follo... |
| `externalIds_ALL` | stringSlice |  | Same as externalIds but value is applied to all lines in the CSV file |
| `familyName` | int64 |  | The family name. |
| `familyName_ALL` | string |  | Same as familyName but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileAses` | int64 |  | The person's file-ases. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `fileAses_ALL` | stringSlice |  | Same as fileAses but value is applied to all lines in the CSV file |
| `genderValue` | int64 |  | The gender for the person. The gender can be custom or one of these predefined values:   - male   - female   - unspec... |
| `genderValue_ALL` | string |  | Same as genderValue but value is applied to all lines in the CSV file |
| `givenName` | int64 |  | The given name. |
| `givenName_ALL` | string |  | Same as givenName but value is applied to all lines in the CSV file |
| `honorificPrefix` | int64 |  | The honorific prefixes, such as Mrs. or Dr. |
| `honorificPrefix_ALL` | string |  | Same as honorificPrefix but value is applied to all lines in the CSV file |
| `honorificSuffix` | int64 |  | The honorific suffixes, such as Jr. |
| `honorificSuffix_ALL` | string |  | Same as honorificSuffix but value is applied to all lines in the CSV file |
| `imClients` | int64 |  | The person's instant messaging clients. Can be used multiple times in the form of "primary=...,username=...", etc. Yo... |
| `imClients_ALL` | stringSlice |  | Same as imClients but value is applied to all lines in the CSV file |
| `interests` | int64 |  | The person's interests. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `interests_ALL` | stringSlice |  | Same as interests but value is applied to all lines in the CSV file |
| `locales` | int64 |  | The person's locale preferences. Can be used multiple times in the form of "primary=...,value=...", etc. You may use ... |
| `locales_ALL` | stringSlice |  | Same as locales but value is applied to all lines in the CSV file |
| `locations` | int64 |  | The person's locations. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `locations_ALL` | stringSlice |  | Same as locations but value is applied to all lines in the CSV file |
| `memberships` | int64 |  | The person's group memberships. Can be used multiple times in the form of "primary=...,contactGroupResourceName=...",... |
| `memberships_ALL` | stringSlice |  | Same as memberships but value is applied to all lines in the CSV file |
| `middleName` | int64 |  | The middle name(s). |
| `middleName_ALL` | string |  | Same as middleName but value is applied to all lines in the CSV file |
| `miscKeywords` | int64 |  | The person's miscellaneous keywords. Can be used multiple times in the form of "primary=...,value=...", etc. You may ... |
| `miscKeywords_ALL` | stringSlice |  | Same as miscKeywords but value is applied to all lines in the CSV file |
| `nicknames` | int64 |  | The person's nicknames. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the follo... |
| `nicknames_ALL` | stringSlice |  | Same as nicknames but value is applied to all lines in the CSV file |
| `occupations` | int64 |  | The person's occupations. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the fol... |
| `occupations_ALL` | stringSlice |  | Same as occupations but value is applied to all lines in the CSV file |
| `organizations` | int64 |  | The person's past or current organizations. Can be used multiple times in the form of "primary=...,type=...", etc. Yo... |
| `organizations_ALL` | stringSlice |  | Same as organizations but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `personFields` | int64 |  | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `personFields_ALL` | string |  | Same as personFields but value is applied to all lines in the CSV file |
| `phoneNumbers` | int64 |  | The person's phone numbers. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the f... |
| `phoneNumbers_ALL` | stringSlice |  | Same as phoneNumbers but value is applied to all lines in the CSV file |
| `phoneticFamilyName` | int64 |  | The family name spelled as it sounds. |
| `phoneticFamilyName_ALL` | string |  | Same as phoneticFamilyName but value is applied to all lines in the CSV file |
| `phoneticFullName` | int64 |  | The full name spelled as it sounds. |
| `phoneticFullName_ALL` | string |  | Same as phoneticFullName but value is applied to all lines in the CSV file |
| `phoneticGivenName` | int64 |  | The given name spelled as it sounds. |
| `phoneticGivenName_ALL` | string |  | Same as phoneticGivenName but value is applied to all lines in the CSV file |
| `phoneticHonorificPrefix` | int64 |  | The honorific prefixes spelled as they sound. |
| `phoneticHonorificPrefix_ALL` | string |  | Same as phoneticHonorificPrefix but value is applied to all lines in the CSV file |
| `phoneticHonorificSuffix` | int64 |  | The honorific suffixes spelled as they sound. |
| `phoneticHonorificSuffix_ALL` | string |  | Same as phoneticHonorificSuffix but value is applied to all lines in the CSV file |
| `phoneticMiddleName` | int64 |  | The middle name(s) spelled as they sound. |
| `phoneticMiddleName_ALL` | string |  | Same as phoneticMiddleName but value is applied to all lines in the CSV file |
| `relations` | int64 |  | The person's relations. Can be used multiple times in the form of "primary=...,person=...", etc. You may use the foll... |
| `relations_ALL` | stringSlice |  | Same as relations but value is applied to all lines in the CSV file |
| `resourceName` | int64 |  | The resource name of the contact- |
| `sipAddresses` | int64 |  | The person's SIP addresses. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the f... |
| `sipAddresses_ALL` | stringSlice |  | Same as sipAddresses but value is applied to all lines in the CSV file |
| `skills` | int64 |  | The person's skills. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the followin... |
| `skills_ALL` | stringSlice |  | Same as skills but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sources` | int64 |  | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |
| `sources_ALL` | string |  | Same as sources but value is applied to all lines in the CSV file |
| `unstructuredName` | int64 |  | The free form name value. |
| `unstructuredName_ALL` | string |  | Same as unstructuredName but value is applied to all lines in the CSV file |
| `updatePersonFields` | int64 |  | A field mask to restrict which fields on the person are updated. Multiple fields can be specified by separating them ... |
| `updatePersonFields_ALL` | string |  | Same as updatePersonFields but value is applied to all lines in the CSV file |
| `urls` | int64 |  | The person's associated URLs. Can be used multiple times in the form of "primary=...,value=...", etc. You may use the... |
| `urls_ALL` | stringSlice |  | Same as urls but value is applied to all lines in the CSV file |
| `userDefined` | int64 |  | The person's user defined data. Can be used multiple times in the form of "primary=...,type=...", etc. You may use th... |
| `userDefined_ALL` | stringSlice |  | Same as userDefined but value is applied to all lines in the CSV file |

#### `gsm people updateContactPhoto`

Update contact data for an existing contact person.

##### `gsm people updateContactPhoto batch`

Batch update contact photos using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `personFields` | int64 |  | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `personFields_ALL` | string |  | Same as personFields but value is applied to all lines in the CSV file |
| `photo` | int64 |  | Path to a photo file. |
| `photo_ALL` | string |  | Same as photo but value is applied to all lines in the CSV file |
| `resourceName` | int64 |  | The resource name of the contact- |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sources` | int64 |  | A mask of what source types to return. DIRECTORY_SOURCE_TYPE_DOMAIN_CONTACT  - Workspace domain shared contact. DIREC... |
| `sources_ALL` | string |  | Same as sources but value is applied to all lines in the CSV file |

### `gsm peopleConnections`

Information about a person merged from various data sources such as the authenticated user's contacts and profile data. (Part of People API)

#### `gsm peopleConnections list`

Provides a list of the authenticated user's contacts.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `personFields` | string | ✓ | A field mask to restrict which fields on each person are returned. Multiple fields can be specified by separating the... |
| `resourceName` | string |  | The resource name to return connections for. Only people/me is valid. |
| `sortOrder` | string |  | Optional. The order in which the connections should be sorted. Defaults to LAST_MODIFIED_ASCENDING. Valid values are:... |
| `sources` | string |  | A mask of what source types to return. READ_SOURCE_TYPE_PROFILE         - Returns SourceType.ACCOUNT, SourceType.DOMA... |

### `gsm permissions`

Manage file and drive permissions (Part of Drive API)

#### `gsm permissions create`

Creates a permission for a file or shared drive.

##### `gsm permissions create batch`

Batch creates permissions for files or shared drives using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `allowFileDiscovery` | int64 |  | Whether the permission allows the file to be discovered through search. This is only applicable for permissions of ty... |
| `allowFileDiscovery_ALL` | bool |  | Same as allowFileDiscovery but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domain` | int64 |  | The domain to which this permission refers. |
| `domain_ALL` | string |  | Same as domain but value is applied to all lines in the CSV file |
| `emailAddress` | int64 |  | The email address of the user or group to which this permission refers. |
| `emailMessage` | int64 |  | A plain text custom message to include in the notification email |
| `emailMessage_ALL` | string |  | Same as emailMessage but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | Id of the file or drive |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `moveToNewOwnersRoot` | int64 |  | This parameter only takes effect if the item is not in a shared drive and the request is attempting to transfer the o... |
| `moveToNewOwnersRoot_ALL` | bool |  | Same as moveToNewOwnersRoot but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `role` | int64 |  | The role granted by this permission. While new values may be supported in the future, the following are currently all... |
| `role_ALL` | string |  | Same as role but value is applied to all lines in the CSV file |
| `sendNotificationEmail` | int64 |  | Whether to send a notification email when sharing to users or groups. This defaults to true for users and groups, and... |
| `sendNotificationEmail_ALL` | bool |  | Same as sendNotificationEmail but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `transferOwnership` | int64 |  | Whether to transfer ownership to the specified user and downgrade the current owner to a writer. This parameter is re... |
| `transferOwnership_ALL` | bool |  | Same as transferOwnership but value is applied to all lines in the CSV file |
| `type` | int64 |  | The type of the grantee. [user\|group\|domain\|anyone]. When creating a permission, if type is user or group, you mus... |
| `type_ALL` | string |  | Same as type but value is applied to all lines in the CSV file |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |
| `view` | int64 |  | Indicates the view for this permission. Only populated for permissions that belong to a view. published is the only s... |
| `view_ALL` | string |  | Same as view but value is applied to all lines in the CSV file |

##### `gsm permissions create recursive`

Recursively grant a permissions to a folder and all of its children.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `allowFileDiscovery` | bool |  | Whether the permission allows the file to be discovered through search. This is only applicable for permissions of ty... |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `domain` | string |  | The domain to which this permission refers. |
| `emailAddress` | string |  | The email address of the user or group to which this permission refers. |
| `emailMessage` | string |  | A plain text custom message to include in the notification email |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `moveToNewOwnersRoot` | bool |  | This parameter only takes effect if the item is not in a shared drive and the request is attempting to transfer the o... |
| `role` | string | ✓ | The role granted by this permission. While new values may be supported in the future, the following are currently all... |
| `sendNotificationEmail` | bool |  | Whether to send a notification email when sharing to users or groups. This defaults to true for users and groups, and... |
| `transferOwnership` | bool |  | Whether to transfer ownership to the specified user and downgrade the current owner to a writer. This parameter is re... |
| `type` | string | ✓ | The type of the grantee. [user\|group\|domain\|anyone]. When creating a permission, if type is user or group, you mus... |
| `useDomainAdminAccess` | bool |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |

#### `gsm permissions delete`

Deletes a permission.

##### `gsm permissions delete batch`

Batch deletes permissions by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domain` | int64 |  | The domain to which this permission refers. |
| `domain_ALL` | string |  | Same as domain but value is applied to all lines in the CSV file |
| `emailAddress` | int64 |  | The email address of the user or group to which this permission refers. |
| `enforceExpansiveAccess` | int64 |  | Whether the request should enforce expansive access rules. See also https://developers.google.com/workspace/drive/api... |
| `enforceExpansiveAccess_ALL` | bool |  | Same as enforceExpansiveAccess but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | Id of the file or drive |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `permissionId` | int64 |  | The ID of the permission. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

##### `gsm permissions delete recursive`

Recursively deletes a permission from a folder and all of its children.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `domain` | string |  | The domain to which this permission refers. |
| `emailAddress` | string |  | The email address of the user or group to which this permission refers. |
| `enforceExpansiveAccess` | bool |  | Whether the request should enforce expansive access rules. See also https://developers.google.com/workspace/drive/api... |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `permissionId` | string |  | The ID of the permission. |
| `useDomainAdminAccess` | bool |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |

#### `gsm permissions get`

Gets a permission by ID.

##### `gsm permissions get batch`

Batch gets permissions by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domain` | int64 |  | The domain to which this permission refers. |
| `domain_ALL` | string |  | Same as domain but value is applied to all lines in the CSV file |
| `emailAddress` | int64 |  | The email address of the user or group to which this permission refers. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | Id of the file or drive |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `permissionId` | int64 |  | The ID of the permission. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

#### `gsm permissions list`

Lists a file's or shared drive's permissions.

##### `gsm permissions list batch`

Batch lists permissions by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | Id of the file or drive |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `includePermissionsForView` | int64 |  | Specifies which additional view's permissions to include in the response. Only 'published' is supported. |
| `includePermissionsForView_ALL` | string |  | Same as includePermissionsForView but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

##### `gsm permissions list recursive`

Recursively lists permissions on a folder and all of its children.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `useDomainAdminAccess` | bool |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |

#### `gsm permissions update`

Updates a permission with patch semantics.

##### `gsm permissions update batch`

Batch updates permissions for a file or shared drive using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `allowFileDiscovery` | int64 |  | Whether the permission allows the file to be discovered through search. This is only applicable for permissions of ty... |
| `allowFileDiscovery_ALL` | bool |  | Same as allowFileDiscovery but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domain` | int64 |  | The domain to which this permission refers. |
| `domain_ALL` | string |  | Same as domain but value is applied to all lines in the CSV file |
| `emailAddress` | int64 |  | The email address of the user or group to which this permission refers. |
| `emailMessage` | int64 |  | A plain text custom message to include in the notification email |
| `emailMessage_ALL` | string |  | Same as emailMessage but value is applied to all lines in the CSV file |
| `enforceExpansiveAccess` | int64 |  | Whether the request should enforce expansive access rules. See also https://developers.google.com/workspace/drive/api... |
| `enforceExpansiveAccess_ALL` | bool |  | Same as enforceExpansiveAccess but value is applied to all lines in the CSV file |
| `expirationTime` | int64 |  | The time at which this permission will expire (RFC 3339 date-time). Expiration times have the following restrictions:... |
| `expirationTime_ALL` | string |  | Same as expirationTime but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | Id of the file or drive |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `moveToNewOwnersRoot` | int64 |  | This parameter only takes effect if the item is not in a shared drive and the request is attempting to transfer the o... |
| `moveToNewOwnersRoot_ALL` | bool |  | Same as moveToNewOwnersRoot but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `permissionId` | int64 |  | The ID of the permission. |
| `removeExpiration` | int64 |  | Whether to remove the expiration date. |
| `removeExpiration_ALL` | bool |  | Same as removeExpiration but value is applied to all lines in the CSV file |
| `role` | int64 |  | The role granted by this permission. While new values may be supported in the future, the following are currently all... |
| `role_ALL` | string |  | Same as role but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `transferOwnership` | int64 |  | Whether to transfer ownership to the specified user and downgrade the current owner to a writer. This parameter is re... |
| `transferOwnership_ALL` | bool |  | Same as transferOwnership but value is applied to all lines in the CSV file |
| `type` | int64 |  | The type of the grantee. [user\|group\|domain\|anyone]. When creating a permission, if type is user or group, you mus... |
| `type_ALL` | string |  | Same as type but value is applied to all lines in the CSV file |
| `useDomainAdminAccess` | int64 |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |
| `useDomainAdminAccess_ALL` | bool |  | Same as useDomainAdminAccess but value is applied to all lines in the CSV file |

##### `gsm permissions update recursive`

Recursively updates a permission on a folder and all of its children.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `allowFileDiscovery` | bool |  | Whether the permission allows the file to be discovered through search. This is only applicable for permissions of ty... |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `domain` | string |  | The domain to which this permission refers. |
| `emailAddress` | string |  | The email address of the user or group to which this permission refers. |
| `emailMessage` | string |  | A plain text custom message to include in the notification email |
| `enforceExpansiveAccess` | bool |  | Whether the request should enforce expansive access rules. See also https://developers.google.com/workspace/drive/api... |
| `excludeFolders` | stringSlice |  | Ids of folders to exclude. Note that due to the way permissions are automatically inherited in Drive, this may not ha... |
| `expirationTime` | string |  | The time at which this permission will expire (RFC 3339 date-time). Expiration times have the following restrictions:... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `folderId` | string |  | File id of the folder. |
| `includeRoot` | bool |  | If set to true, the root (specified parent) is included in the results |
| `moveToNewOwnersRoot` | bool |  | This parameter only takes effect if the item is not in a shared drive and the request is attempting to transfer the o... |
| `permissionId` | string |  | The ID of the permission. |
| `removeExpiration` | bool |  | Whether to remove the expiration date. |
| `role` | string |  | The role granted by this permission. While new values may be supported in the future, the following are currently all... |
| `transferOwnership` | bool |  | Whether to transfer ownership to the specified user and downgrade the current owner to a writer. This parameter is re... |
| `type` | string |  | The type of the grantee. [user\|group\|domain\|anyone]. When creating a permission, if type is user or group, you mus... |
| `useDomainAdminAccess` | bool |  | Issue the request as a domain administrator; if set to true, then the requester will be granted access if the file ID... |

### `gsm postmasterDomains`

Use Gmail Postmaster Tools to manage domain (Part of Gmail Postmaster API)

#### `gsm postmasterDomains get`

Gets a specific domain registered by the client. Returns NOT_FOUND if the domain does not exist.

##### `gsm postmasterDomains get batch`

Batch gets domains by fully qualified name using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | Fully qualified domain name. |
| `name_ALL` | string |  | Same as name but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm postmasterDomains list`

Lists the domains that have been registered by the client.
The order of domains in the response is unspecified and non-deterministic.
Newly created domains will not necessarily be added to the end of this list.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

### `gsm postmasterTrafficStats`

Use Gmail Postmaster Tools to view email traffic statistics (Part of Gmail Postmaster API)

#### `gsm postmasterTrafficStats get`

Get traffic statistics for a domain on a specific date.
Returns PERMISSION_DENIED if user does not have permission to access TrafficStats for the domain.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string | ✓ | The resource name of the traffic statistics to get. E.g., domains/mymail.mydomain.com/trafficStats/20160807. |

#### `gsm postmasterTrafficStats list`

List traffic statistics for all available days.
Returns PERMISSION_DENIED if user does not have permission to access TrafficStats for the domain.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `endDateDay` | int64 |  | The day of the most recent date of the metrics to retrieve inclusive. If you specify one date flag, you must specify ... |
| `endDateMonth` | int64 |  | The month of the most recent date of the metrics to retrieve inclusive. If you specify one date flag, you must specif... |
| `endDateYear` | int64 |  | The year of the most recent date of the metrics to retrieve inclusive. If you specify one date flag, you must specify... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string | ✓ | Fully qualified domain name. |
| `startDateDay` | int64 |  | The day of the earliest date of the metrics to retrieve inclusive. If you specify one date flag, you must specify ALL... |
| `startDateMonth` | int64 |  | The month of the earliest date of the metrics to retrieve inclusive. If you specify one date flag, you must specify A... |
| `startDateYear` | int64 |  | The year of the earliest date of the metrics to retrieve inclusive. If you specify one date flag, you must specify AL... |

### `gsm privileges`

Manage (list) Privileges (Part of Admin SDK API)

#### `gsm privileges list`

Retrieves a paginated list of all privileges for a customer.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | Immutable ID of the Workspace account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

### `gsm replies`

Manage replies to comments (Part of Drive API)

#### `gsm replies create`

Creates a new reply to a comment.

##### `gsm replies create batch`

Batch creates new replies to comments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `action` | int64 |  | The action the reply performed to the parent comment. [resolve\|reopen] |
| `action_ALL` | string |  | Same as action but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commentId` | int64 |  | The ID of the comment. |
| `commentId_ALL` | string |  | Same as commentId but value is applied to all lines in the CSV file |
| `content` | int64 |  | The plain text content of the comment. This field is used for setting the content, while htmlContent should be displa... |
| `content_ALL` | string |  | Same as content but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm replies delete`

Deletes a reply.

##### `gsm replies delete batch`

Batch deletes replies by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commentId` | int64 |  | The ID of the comment. |
| `commentId_ALL` | string |  | Same as commentId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `replyId` | int64 |  | The ID of the reply. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm replies get`

Gets a reply by ID.

##### `gsm replies get batch`

Batch gets replies by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commentId` | int64 |  | The ID of the comment. |
| `commentId_ALL` | string |  | Same as commentId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `replyId` | int64 |  | The ID of the reply. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm replies list`

Lists a comment's replies.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `commentId` | string | ✓ | The ID of the comment. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fileId` | string | ✓ | The ID of the file. |

#### `gsm replies update`

Updates a reply with patch semantics.

##### `gsm replies update batch`

Batch updates replies to comments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `action` | int64 |  | The action the reply performed to the parent comment. [resolve\|reopen] |
| `action_ALL` | string |  | Same as action but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `commentId` | int64 |  | The ID of the comment. |
| `commentId_ALL` | string |  | Same as commentId but value is applied to all lines in the CSV file |
| `content` | int64 |  | The plain text content of the comment. This field is used for setting the content, while htmlContent should be displa... |
| `content_ALL` | string |  | Same as content but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `replyId` | int64 |  | The ID of the reply. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm revisions`

Manage revisions of non-Google files (Part of Drive API)

#### `gsm revisions delete`

Permanently deletes a file version.

##### `gsm revisions delete batch`

Batch deletes revisions using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `acknowledgeAbuse` | int64 |  | Whether the user is acknowledging the risk of downloading known malware or other abusive files. This is only applicab... |
| `acknowledgeAbuse_ALL` | bool |  | Same as acknowledgeAbuse but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `revisionId` | int64 |  | The ID of the revision. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm revisions get`

Gets a revision's metadata or content by ID.

##### `gsm revisions get batch`

Batch gets revisions' metadata or content by ID using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `acknowledgeAbuse` | int64 |  | Whether the user is acknowledging the risk of downloading known malware or other abusive files. This is only applicab... |
| `acknowledgeAbuse_ALL` | bool |  | Same as acknowledgeAbuse but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `revisionId` | int64 |  | The ID of the revision. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm revisions list`

Lists a file's revisions.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fileId` | string | ✓ | The ID of the file. |

#### `gsm revisions update`

Updates a revision with patch semantics.

##### `gsm revisions update batch`

Batch updates revisions using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `acknowledgeAbuse` | int64 |  | Whether the user is acknowledging the risk of downloading known malware or other abusive files. This is only applicab... |
| `acknowledgeAbuse_ALL` | bool |  | Same as acknowledgeAbuse but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `fileId` | int64 |  | The ID of the file. |
| `fileId_ALL` | string |  | Same as fileId but value is applied to all lines in the CSV file |
| `keepForever` | int64 |  | Whether to keep this revision forever, even if it is no longer the head revision. If not set, the revision will be au... |
| `keepForever_ALL` | bool |  | Same as keepForever but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `publishAuto` | int64 |  | Whether subsequent revisions will be automatically republished. This is only applicable to Google Docs. |
| `publishAuto_ALL` | bool |  | Same as publishAuto but value is applied to all lines in the CSV file |
| `published` | int64 |  | Whether this revision is published. This is only applicable to Google Docs. |
| `publishedOutsideDomain` | int64 |  | Whether this revision is published outside the domain. This is only applicable to Google Docs. |
| `publishedOutsideDomain_ALL` | bool |  | Same as publishedOutsideDomain but value is applied to all lines in the CSV file |
| `published_ALL` | bool |  | Same as published but value is applied to all lines in the CSV file |
| `revisionId` | int64 |  | The ID of the revision. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm roleAssignments`

Manage Role Assignments (Part of Admin SDK API)

#### `gsm roleAssignments delete`

Deletes a role assignment.

##### `gsm roleAssignments delete batch`

Batch deletes role assignments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roleAssignmentId` | int64 |  | Immutable ID of the role assignment. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm roleAssignments get`

Retrieve a role assignment.

##### `gsm roleAssignments get batch`

Batch retrieve role assignments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roleAssignmentId` | int64 |  | Immutable ID of the role assignment. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm roleAssignments insert`

Creates a role assignment.

##### `gsm roleAssignments insert batch`

Batch inserts role assignments using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `assignedTo` | int64 |  | The unique ID of the user this role is assigned to. |
| `assignedTo_ALL` | string |  | Same as assignedTo but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `orgUnitId` | int64 |  | If the role is restricted to an organization unit, this contains the ID for the organization unit the exercise of thi... |
| `orgUnitId_ALL` | string |  | Same as orgUnitId but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roleId` | int64 |  | The ID of the role that is assigned. |
| `roleId_ALL` | int64 |  | Same as roleId but value is applied to all lines in the CSV file |
| `scopeType` | int64 |  | The scope in which this role is assigned. Acceptable values are: CUSTOMER ORG_UNIT |
| `scopeType_ALL` | string |  | Same as scopeType but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

##### `gsm roleAssignments insert recursive`

Creates role assignments for users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `customer` | string |  | Immutable ID of the Workspace account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `orgUnitId` | string |  | If the role is restricted to an organization unit, this contains the ID for the organization unit the exercise of thi... |
| `roleId` | int64 | ✓ | The ID of the role that is assigned. |
| `scopeType` | string |  | The scope in which this role is assigned. Acceptable values are: CUSTOMER ORG_UNIT |

#### `gsm roleAssignments list`

Retrieves a paginated list of all roleAssignments.

##### `gsm roleAssignments list recursive`

List users' role assignments by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `customer` | string |  | Immutable ID of the Workspace account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

### `gsm roles`

Manage roles (Part of Admin SDK API)

#### `gsm roles delete`

Deletes a role.

##### `gsm roles delete batch`

Batch deletes roles using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roleId` | int64 |  | Immutable ID of the role. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm roles get`

Retrieves a role.

##### `gsm roles get batch`

Batch retrieves roles using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roleId` | int64 |  | Immutable ID of the role. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm roles insert`

Creates a role.

##### `gsm roles insert batch`

Batch inserts roles using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roleName` | int64 |  | Name of the role. |
| `rolePrivileges` | int64 |  | The set of privileges that are granted to this role. |
| `rolePrivileges_ALL` | stringSlice |  | Same as rolePrivileges but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm roles list`

Retrieves a paginated list of all the roles in a domain.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | Immutable ID of the Workspace account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm roles patch`

Updates a role. This method supports patch semantics.

##### `gsm roles patch batch`

Batch patches roles using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customer` | int64 |  | Immutable ID of the Workspace account. |
| `customer_ALL` | string |  | Same as customer but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `roleId` | int64 |  | Immutable ID of the role. |
| `roleName` | int64 |  | Name of the role. |
| `rolePrivileges` | int64 |  | The set of privileges that are granted to this role. |
| `rolePrivileges_ALL` | stringSlice |  | Same as rolePrivileges but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm schemas`

Manage custom schemas for user accounts (Part of Admin SDK API)

#### `gsm schemas delete`

Delete a custom schema.

##### `gsm schemas delete batch`

Batch deletes custom schemas using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | Immutable ID of the Workspace account. |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `schemaKey` | int64 |  | Name or immutable ID of the schema. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm schemas get`

Retrieve a custom schema.

##### `gsm schemas get batch`

Batch gets custom schemas using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | Immutable ID of the Workspace account. |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `schemaKey` | int64 |  | Name or immutable ID of the schema. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm schemas insert`

Create a custom schema.

##### `gsm schemas insert batch`

Batch inserts custom schemas using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | Immutable ID of the Workspace account. |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `displayName` | int64 |  | Display name for the schema. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `schemaFields` | int64 |  | The fields that should be present in this schema. Can be used multiple times in the form of: "--schemaFields "fieldNa... |
| `schemaName` | int64 |  | The schema's name. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm schemas list`

List custom schemas

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | Immutable ID of the Workspace account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |

#### `gsm schemas patch`

Patches a custom schema

##### `gsm schemas patch batch`

Batch patches schemas using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customerId` | int64 |  | Immutable ID of the Workspace account. |
| `customerId_ALL` | string |  | Same as customerId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `displayName` | int64 |  | Display name for the schema. |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `schemaFields` | int64 |  | The fields that should be present in this schema. Can be used multiple times in the form of: "--schemaFields "fieldNa... |
| `schemaKey` | int64 |  | Name or immutable ID of the schema. |
| `schemaName` | int64 |  | The schema's name. |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm sendAs`

Manage send-as settings for users (Part of Gmail API)

#### `gsm sendAs create`

Creates a custom "from" send-as alias

##### `gsm sendAs create batch`

Batch creates custom "from" send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `displayName` | int64 |  | A name that appears in the "From:" header for mail sent using this alias. For custom "from" addresses, when this is e... |
| `displayName_ALL` | string |  | Same as displayName but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `host` | int64 |  | The hostname of the SMTP service. Required for SMTP. |
| `host_ALL` | string |  | Same as host but value is applied to all lines in the CSV file |
| `isDefault` | int64 |  | Whether this address is selected as the default "From:" address in situations such as composing a new message or send... |
| `isDefault_ALL` | bool |  | Same as isDefault but value is applied to all lines in the CSV file |
| `password` | int64 |  | The password that will be used for authentication with the SMTP service. |
| `password_ALL` | string |  | Same as password but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `port` | int64 |  | The port of the SMTP service. Required for SMTP. |
| `port_ALL` | int64 |  | Same as port but value is applied to all lines in the CSV file |
| `replyToAddress` | int64 |  | An optional email address that is included in a "Reply-To:" header for mail sent using this alias. If this is empty, ... |
| `replyToAddress_ALL` | string |  | Same as replyToAddress but value is applied to all lines in the CSV file |
| `securityMode` | int64 |  | The protocol that will be used to secure communication with the SMTP service. Required for SMTP. [NONE\|SSL\|STARTTLS... |
| `securityMode_ALL` | string |  | Same as securityMode but value is applied to all lines in the CSV file |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `signature` | int64 |  | An optional HTML signature that is included in messages composed with this alias in the Gmail web UI. |
| `signature_ALL` | string |  | Same as signature but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `treatAsAlias` | int64 |  | Whether Gmail should treat this address as an alias for the user's primary email address. This setting only applies t... |
| `treatAsAlias_ALL` | bool |  | Same as treatAsAlias but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |
| `username` | int64 |  | The username that will be used for authentication with the SMTP service. |
| `username_ALL` | string |  | Same as username but value is applied to all lines in the CSV file |

#### `gsm sendAs delete`

Deletes the specified send-as alias. Revokes any verification that may have been required for using it.

##### `gsm sendAs delete batch`

Batch deletes the specified send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm sendAs get`

Gets the specified send-as alias.

##### `gsm sendAs get batch`

Batch gets the specified send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm sendAs list`

Lists the send-as aliases for the specified account.
The result includes the primary send-as address associated with the account as well as any custom "from" aliases.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `userId` | string |  | The user's email address. The special value me can be used to indicate the authenticated user. |

#### `gsm sendAs patch`

Patch the specified send-as alias.

##### `gsm sendAs patch batch`

Batch patches custom "from" send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `displayName` | int64 |  | A name that appears in the "From:" header for mail sent using this alias. For custom "from" addresses, when this is e... |
| `displayName_ALL` | string |  | Same as displayName but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `host` | int64 |  | The hostname of the SMTP service. Required for SMTP. |
| `host_ALL` | string |  | Same as host but value is applied to all lines in the CSV file |
| `isDefault` | int64 |  | Whether this address is selected as the default "From:" address in situations such as composing a new message or send... |
| `isDefault_ALL` | bool |  | Same as isDefault but value is applied to all lines in the CSV file |
| `password` | int64 |  | The password that will be used for authentication with the SMTP service. |
| `password_ALL` | string |  | Same as password but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `port` | int64 |  | The port of the SMTP service. Required for SMTP. |
| `port_ALL` | int64 |  | Same as port but value is applied to all lines in the CSV file |
| `replyToAddress` | int64 |  | An optional email address that is included in a "Reply-To:" header for mail sent using this alias. If this is empty, ... |
| `replyToAddress_ALL` | string |  | Same as replyToAddress but value is applied to all lines in the CSV file |
| `securityMode` | int64 |  | The protocol that will be used to secure communication with the SMTP service. Required for SMTP. [NONE\|SSL\|STARTTLS... |
| `securityMode_ALL` | string |  | Same as securityMode but value is applied to all lines in the CSV file |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `signature` | int64 |  | An optional HTML signature that is included in messages composed with this alias in the Gmail web UI. |
| `signature_ALL` | string |  | Same as signature but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `treatAsAlias` | int64 |  | Whether Gmail should treat this address as an alias for the user's primary email address. This setting only applies t... |
| `treatAsAlias_ALL` | bool |  | Same as treatAsAlias but value is applied to all lines in the CSV file |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |
| `username` | int64 |  | The username that will be used for authentication with the SMTP service. |
| `username_ALL` | string |  | Same as username but value is applied to all lines in the CSV file |

#### `gsm sendAs verify`

Sends a verification email to the specified send-as alias address. The verification status must be pending.

##### `gsm sendAs verify batch`

Batch sends verification emails for send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

### `gsm sharedContacts`

Manage Domain Shared Contacts (Part of Shared Contacts API - not Admin SDK!)

#### `gsm sharedContacts create`

Create a Domain Shared Contact

##### `gsm sharedContacts create batch`

Batch create Domain Shared Contacts

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `additionalName` | int64 |  | Additional name of the person, eg. middle name. |
| `additionalName_ALL` | string |  | Same as additionalName but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domain` | int64 |  | DNS domain of the shared contact |
| `domain_ALL` | string |  | Same as domain but value is applied to all lines in the CSV file |
| `email` | int64 |  | Email. Must be in the form of "address=user@domain.com;displayName=Some Name;primary=[true\|false];label=[Work\|Home]... |
| `email_ALL` | stringSlice |  | Same as email but value is applied to all lines in the CSV file |
| `extendedProperty` | int64 |  | Extended Properties Must be in the form of "name=Some Name;Value=Some Value;Realm=Some Realm" |
| `extendedProperty_ALL` | stringSlice |  | Same as extendedProperty but value is applied to all lines in the CSV file |
| `familyName` | int64 |  | Person's family name. |
| `familyName_ALL` | string |  | Same as familyName but value is applied to all lines in the CSV file |
| `fullName` | int64 |  | Unstructured representation of the name. |
| `fullName_ALL` | string |  | Same as fullName but value is applied to all lines in the CSV file |
| `givenName` | int64 |  | Person's given name. |
| `givenName_ALL` | string |  | Same as givenName but value is applied to all lines in the CSV file |
| `im` | int64 |  | IM addresses. Must be in the form of "protocol=http://schemas.google.com/g/2005#GOOGLE_TALK;address=some@address.com;... |
| `im_ALL` | stringSlice |  | Same as im but value is applied to all lines in the CSV file |
| `json` | int64 |  | Output as JSON" |
| `json_ALL` | bool |  | Same as json but value is applied to all lines in the CSV file |
| `namePrefix` | int64 |  | Honorific prefix, eg. 'Mr' or 'Mrs'. |
| `namePrefix_ALL` | string |  | Same as namePrefix but value is applied to all lines in the CSV file |
| `nameSuffix` | int64 |  | Honorific suffix, eg. 'san' or 'III'. |
| `nameSuffix_ALL` | string |  | Same as nameSuffix but value is applied to all lines in the CSV file |
| `organization` | int64 |  | Organization of the contact. Must be in the form of "orgName=Some Company;orgDepartment=Some Department;orgTitle=Some... |
| `organization_ALL` | stringSlice |  | Same as organization but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `phoneNumber` | int64 |  | Phone number. Must be in the form of "phoneNumber=+1 212 213181;primary=[true\|false]label=[Work\|Home\|Mobile]". Can... |
| `phoneNumber_ALL` | stringSlice |  | Same as phoneNumber but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `structuredPostalAddress` | int64 |  | Structured Postal Address Must be in the form of "mailClass=...;label=...;usage=...;primary=[true\|false];agent=...;h... |
| `structuredPostalAddress_ALL` | stringSlice |  | Same as structuredPostalAddress but value is applied to all lines in the CSV file |

#### `gsm sharedContacts delete`

Delete a shared contact by referencing its id url (must begin with https://)

##### `gsm sharedContacts delete batch`

Batch deletes Domain Shared Contacts via URL / ID using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domain` | int64 |  | DNS domain of the shared contact |
| `domain_ALL` | string |  | Same as domain but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `url` | int64 |  | URL of the Shared Contact (Retrieve with "list" and look for "id"). MUST BE https://! |

#### `gsm sharedContacts get`

Gets a Domain Shared Contact via its URL / ID

##### `gsm sharedContacts get batch`

Batch gets Domain Shared Contacts via URL / ID using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `domain` | int64 |  | DNS domain of the shared contact |
| `domain_ALL` | string |  | Same as domain but value is applied to all lines in the CSV file |
| `json` | int64 |  | Output as JSON" |
| `json_ALL` | bool |  | Same as json but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `url` | int64 |  | URL of the Shared Contact (Retrieve with "list" and look for "id"). MUST BE https://! |

#### `gsm sharedContacts list`

List all shared contacts in your domain

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `domain` | string | ✓ | DNS domain of the shared contact |
| `json` | bool |  | Output as JSON" |

#### `gsm sharedContacts update`

Update a shared contact

##### `gsm sharedContacts update batch`

Batch updates Domain Shared Contacts using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `additionalName` | int64 |  | Additional name of the person, eg. middle name. |
| `additionalName_ALL` | string |  | Same as additionalName but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `email` | int64 |  | Email. Must be in the form of "address=user@domain.com;displayName=Some Name;primary=[true\|false];label=[Work\|Home]... |
| `email_ALL` | stringSlice |  | Same as email but value is applied to all lines in the CSV file |
| `extendedProperty` | int64 |  | Extended Properties Must be in the form of "name=Some Name;Value=Some Value;Realm=Some Realm" |
| `extendedProperty_ALL` | stringSlice |  | Same as extendedProperty but value is applied to all lines in the CSV file |
| `familyName` | int64 |  | Person's family name. |
| `familyName_ALL` | string |  | Same as familyName but value is applied to all lines in the CSV file |
| `fullName` | int64 |  | Unstructured representation of the name. |
| `fullName_ALL` | string |  | Same as fullName but value is applied to all lines in the CSV file |
| `givenName` | int64 |  | Person's given name. |
| `givenName_ALL` | string |  | Same as givenName but value is applied to all lines in the CSV file |
| `im` | int64 |  | IM addresses. Must be in the form of "protocol=http://schemas.google.com/g/2005#GOOGLE_TALK;address=some@address.com;... |
| `im_ALL` | stringSlice |  | Same as im but value is applied to all lines in the CSV file |
| `json` | int64 |  | Output as JSON" |
| `json_ALL` | bool |  | Same as json but value is applied to all lines in the CSV file |
| `namePrefix` | int64 |  | Honorific prefix, eg. 'Mr' or 'Mrs'. |
| `namePrefix_ALL` | string |  | Same as namePrefix but value is applied to all lines in the CSV file |
| `nameSuffix` | int64 |  | Honorific suffix, eg. 'san' or 'III'. |
| `nameSuffix_ALL` | string |  | Same as nameSuffix but value is applied to all lines in the CSV file |
| `organization` | int64 |  | Organization of the contact. Must be in the form of "orgName=Some Company;orgDepartment=Some Department;orgTitle=Some... |
| `organization_ALL` | stringSlice |  | Same as organization but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `phoneNumber` | int64 |  | Phone number. Must be in the form of "phoneNumber=+1 212 213181;primary=[true\|false]label=[Work\|Home\|Mobile]". Can... |
| `phoneNumber_ALL` | stringSlice |  | Same as phoneNumber but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `structuredPostalAddress` | int64 |  | Structured Postal Address Must be in the form of "mailClass=...;label=...;usage=...;primary=[true\|false];agent=...;h... |
| `structuredPostalAddress_ALL` | stringSlice |  | Same as structuredPostalAddress but value is applied to all lines in the CSV file |
| `url` | int64 |  | URL of the Shared Contact (Retrieve with "list" and look for "id"). MUST BE https://! |

### `gsm smimeInfo`

Manage users' S/MIME configs for send-as aliases

#### `gsm smimeInfo delete`

Deletes the specified S/MIME config for the specified send-as alias.

##### `gsm smimeInfo delete batch`

Batch deletes the specified S/MIME configs for the specified send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `id` | int64 |  | The immutable ID for the SmimeInfo. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm smimeInfo get`

Gets the specified S/MIME config for the specified send-as alias.

##### `gsm smimeInfo get batch`

Batch gets the specified S/MIME config for the specified send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | The immutable ID for the SmimeInfo. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm smimeInfo insert`

Insert (upload) the given S/MIME config for the specified send-as alias.
Note that pkcs12 format is required for the key.

##### `gsm smimeInfo insert batch`

Batch inserts S/MIME configs for the specified send-as aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `encryptedKeyPassword` | int64 |  | Encrypted key password, when key is encrypted. |
| `encryptedKeyPassword_ALL` | string |  | Same as encryptedKeyPassword but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `pkcs12` | int64 |  | Path to a PKCS#12 format file containing a single private/public key pair and certificate chain. This format is only ... |
| `pkcs12_ALL` | string |  | Same as pkcs12 but value is applied to all lines in the CSV file |
| `sendAsEmail` | int64 |  | The email address that appears in the "From:" header for mail sent using this alias. |
| `sendAsEmail_ALL` | string |  | Same as sendAsEmail but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm smimeInfo list`

Lists S/MIME configs for the specified send-as alias.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `sendAsEmail` | string | ✓ | The email address that appears in the "From:" header for mail sent using this alias. |
| `userId` | string |  | The user's email address. The special value me can be used to indicate the authenticated user. |

#### `gsm smimeInfo setDefault`

Sets the default S/MIME config for the specified send-as alias.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `id` | string | ✓ | The immutable ID for the SmimeInfo. |
| `sendAsEmail` | string | ✓ | The email address that appears in the "From:" header for mail sent using this alias. |
| `userId` | string |  | The user's email address. The special value me can be used to indicate the authenticated user. |

### `gsm spreadsheets`

Manage Google Sheets spreadsheets (Part of Sheets API)

#### `gsm spreadsheets batchUpdate`

Applies one or more updates to the spreadsheet.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `csvFileToUpload` | stringSlice |  | A list of CSV files that should be added to the spreadsheet as new sheets. Can be used multiple times in the form of ... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `spreadsheetId` | string | ✓ | The ID of the spreadsheet |

#### `gsm spreadsheets create`

Creates a spreadsheet, returning the newly created spreadsheet.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `csvFileToUpload` | stringSlice |  | A list of CSV files that should be added to the spreadsheet as new sheets. Can be used multiple times in the form of ... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `title` | string | ✓ | The ID of the spreadsheet |

#### `gsm spreadsheets get`

Returns the spreadsheet at the given ID.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `includeGridData` | bool |  | True if grid data should be returned. This parameter is ignored if a field mask was set in the request. |
| `ranges` | stringSlice |  | The ranges to retrieve from the spreadsheet. |
| `spreadsheetId` | string | ✓ | The ID of the spreadsheet |

### `gsm ssoAssignments`

Manage inbound SAML SSO assignments (Part of Cloud Identity API)

#### `gsm ssoAssignments create`

Creates an InboundSsoAssignment for users and devices in a Customer under a given Group or OrgUnit.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | The customer. For example: customers/C0123abc. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `inboundSamlSsoProfile` | string | ✓ | Name of the InboundSamlSsoProfile to use. Must be of the form inboundSamlSsoProfiles/{inboundSamlSsoProfile}. |
| `rank` | int64 |  | Must be zero (which is the default value so it can be omitted) for assignments with targetOrgUnit set and must be gre... |
| `ssoMode` | string |  | Inbound SSO behaviors. May be one of the following: - SSO_OFF                      - Disable SSO for the targeted use... |
| `targetGroup` | string |  | Must be of the form groups/{group} Only ONE of --targetGroup and --targetOrgUnit may be specified. |
| `targetOrgUnit` | string |  | Must be of the form orgUnits/{orgUnit}. Only ONE of --targetGroup and --targetOrgUnit may be specified. |

#### `gsm ssoAssignments delete`

Deletes an InboundSsoAssignment.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | The resource name of the InboundSsoAssignment. Format: inboundSsoAssignments/{assignment} |

#### `gsm ssoAssignments get`

Gets an InboundSsoAssignment.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string | ✓ | The resource name of the InboundSsoAssignment. Format: inboundSsoAssignments/{assignment} |

#### `gsm ssoAssignments list`

Lists an InboundSsoAssignment.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | A CEL expression to filter the results. The only supported filter is filtering by customer. For example: customer==cu... |

#### `gsm ssoAssignments patch`

Updates an InboundSsoAssignment.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customer` | string |  | The customer. For example: customers/C0123abc. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `inboundSamlSsoProfile` | string |  | Name of the InboundSamlSsoProfile to use. Must be of the form inboundSamlSsoProfiles/{inboundSamlSsoProfile}. |
| `name` | string | ✓ | The resource name of the InboundSsoAssignment. Format: inboundSsoAssignments/{assignment} |
| `rank` | int64 |  | Must be zero (which is the default value so it can be omitted) for assignments with targetOrgUnit set and must be gre... |
| `ssoMode` | string |  | Inbound SSO behaviors. May be one of the following: - SSO_OFF                      - Disable SSO for the targeted use... |
| `targetGroup` | string |  | Must be of the form groups/{group} Only ONE of --targetGroup and --targetOrgUnit may be specified. |
| `targetOrgUnit` | string |  | Must be of the form orgUnits/{orgUnit}. Only ONE of --targetGroup and --targetOrgUnit may be specified. |

### `gsm ssoProfileCredentials`

Manage inbound SAML SSO profile IdP Credentials (Part of Cloud Identity API)

#### `gsm ssoProfileCredentials add`

Adds an IdpCredential. Up to 2 credentials are allowed.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string | ✓ | The InboundSamlSsoProfile that owns the IdpCredential(s). Format: inboundSamlSsoProfiles/{sso_profile_id} If you don'... |
| `pemFile` | string | ✓ | The file path to a PEM encoded x509 certificate containing the public key for verifying IdP signatures. |

#### `gsm ssoProfileCredentials delete`

Deletes an IdpCredential.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | The resource name of the IdpCredential. Format: inboundSamlSsoProfiles/{sso_profile_id}/idpCredentials/{idp_credentia... |

#### `gsm ssoProfileCredentials get`

Gets an IdpCredential.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string | ✓ | The resource name of the IdpCredential. Format: inboundSamlSsoProfiles/{sso_profile_id}/idpCredentials/{idp_credentia... |

#### `gsm ssoProfileCredentials list`

Returns a list of IdpCredentials in an InboundSamlSsoProfile.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `parent` | string | ✓ | The InboundSamlSsoProfile that owns the IdpCredential(s). Format: inboundSamlSsoProfiles/{sso_profile_id} If you don'... |

### `gsm ssoProfiles`

Manage inbound SAML SSO profiles (Part of Cloud Identity API)

#### `gsm ssoProfiles create`

Creates an InboundSamlSsoProfile for a customer.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `changePasswordUri` | string |  | The Change Password URL of the identity provider. Users will be sent to this URL when changing their passwords at mya... |
| `customer` | string |  | The customer. For example: customers/C0123abc. |
| `displayName` | string |  | Human-readable name of the SAML SSO profile. |
| `entityId` | string | ✓ | The SAML Entity ID of the identity provider. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `logoutRedirectUri` | string |  | The Logout Redirect URL (sign-out page URL) of the identity provider. When a user clicks the sign-out link on a Googl... |
| `singleSignOnServiceUri` | string | ✓ | The SingleSignOnService endpoint location (sign-in page URL) of the identity provider. This is the URL where the Auth... |

#### `gsm ssoProfiles delete`

Deletes an InboundSamlSsoProfile

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `name` | string | ✓ | The resource name of the InboundSamlSsoProfile to delete. Format: inboundSamlSsoProfiles/{sso_profile_id}. If you don... |

#### `gsm ssoProfiles get`

Gets an InboundSamlSsoProfile.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `name` | string | ✓ | The resource name of the InboundSamlSsoProfile to delete. Format: inboundSamlSsoProfiles/{sso_profile_id}. If you don... |

#### `gsm ssoProfiles list`

Lists InboundSamlSsoProfiles for a customer.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | A Common Expression Language expression to filter the results. The only supported filter is filtering by customer. Fo... |

#### `gsm ssoProfiles patch`

Updates an InboundSamlSsoProfile.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `changePasswordUri` | string |  | The Change Password URL of the identity provider. Users will be sent to this URL when changing their passwords at mya... |
| `customer` | string |  | The customer. For example: customers/C0123abc. |
| `displayName` | string |  | Human-readable name of the SAML SSO profile. |
| `entityId` | string |  | The SAML Entity ID of the identity provider. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `logoutRedirectUri` | string |  | The Logout Redirect URL (sign-out page URL) of the identity provider. When a user clicks the sign-out link on a Googl... |
| `name` | string | ✓ | The resource name of the InboundSamlSsoProfile to delete. Format: inboundSamlSsoProfiles/{sso_profile_id}. If you don... |
| `singleSignOnServiceUri` | string |  | The SingleSignOnService endpoint location (sign-in page URL) of the identity provider. This is the URL where the Auth... |

### `gsm threads`

Manage threads in users' mailboxes (Part of Gmail API)

#### `gsm threads delete`

Immediately and permanently deletes the specified thread.
This operation cannot be undone. Prefer threads trash instead.

##### `gsm threads delete batch`

Batch deletes the specified threads using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `id` | int64 |  | ID of the Thread. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm threads get`

Gets the specified thread.

##### `gsm threads get batch`

Batch gets the specified threads using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `format` | int64 |  | The format to return the message in. [MINIMAL\|FULL\|RAW\|METADATA] MINIMAL   - Returns only email message ID and lab... |
| `format_ALL` | string |  | Same as format but value is applied to all lines in the CSV file |
| `id` | int64 |  | ID of the Thread. |
| `metadataHeaders` | int64 |  | When given and format is METADATA, only include headers specified. |
| `metadataHeaders_ALL` | string |  | Same as metadataHeaders but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm threads list`

Lists the threads in the user's mailbox.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `includeSpamTrash` | bool |  | Include threads from SPAM and TRASH in the results. |
| `labelIds` | stringSlice |  | Only return threads with labels that match all of the specified label IDs. |
| `q` | string |  | Only return threads matching the specified query. Supports the same query format as the Gmail search box. For example... |
| `userId` | string |  | The user's email address. The special value me can be used to indicate the authenticated user. |

#### `gsm threads modify`

Modifies the labels applied to the thread. This applies to all messages in the thread.

##### `gsm threads modify batch`

Batch modifies threads using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addLabelIds` | int64 |  | A list of label IDs to add to threads. |
| `addLabelIds_ALL` | stringSlice |  | Same as addLabelIds but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | ID of the Thread. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `removeLabelIds` | int64 |  | A list of label IDs to remove from threads. |
| `removeLabelIds_ALL` | stringSlice |  | Same as removeLabelIds but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm threads trash`

Moves the specified thread to the trash.

##### `gsm threads trash batch`

Batch trashes threads using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | ID of the Thread. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

#### `gsm threads untrash`

Removes the specified thread from the trash.

##### `gsm threads untrash batch`

Batch untrashes threads using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `id` | int64 |  | ID of the Thread. |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userId` | int64 |  | The user's email address. The special value me can be used to indicate the authenticated user. |
| `userId_ALL` | string |  | Same as userId but value is applied to all lines in the CSV file |

### `gsm tokens`

Managed OAuth access tokens for users (Part of Admin SDK API)

#### `gsm tokens delete`

Delete all access tokens issued by a user for an application.

##### `gsm tokens delete batch`

Batch delete access tokens issued by a user using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `clientId` | int64 |  | The Client ID of the application the token is issued to. |
| `clientId_ALL` | string |  | Same as clientId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

##### `gsm tokens delete recursive`

Deletes a token issued to a 3rd party application from users' tokens by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `clientId` | string | ✓ | The Client ID of the application the token is issued to. |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm tokens get`

Get information about an access token issued by a user.

##### `gsm tokens get batch`

Batch get information about access tokens issued by a user using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `clientId` | int64 |  | The Client ID of the application the token is issued to. |
| `clientId_ALL` | string |  | Same as clientId but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

#### `gsm tokens list`

Returns the set of tokens specified user has issued to 3rd party applications.

##### `gsm tokens list batch`

Batch list access tokens issued by a user using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

##### `gsm tokens list recursive`

Returns a list of tokens issued to 3rd party applications by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

### `gsm twoStepVerification`

Manage Two Step Verification for users (Park of Admin SDK)

#### `gsm twoStepVerification turnOff`

Turn off 2-Step Verification for user.

##### `gsm twoStepVerification turnOff batch`

Batch turns off two step verification for users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm twoStepVerification turnOff recursive`

Turns off two step verification for users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

### `gsm userAliases`

Manage user aliases, which are alternative email addresses (Part of Admin SDK - not Gmail API!)

#### `gsm userAliases delete`

Removes an alias.

##### `gsm userAliases delete batch`

Batch deletes user aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `alias` | int64 |  | The alias email address. |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

#### `gsm userAliases insert`

Adds an alias.

##### `gsm userAliases insert batch`

Batch insert user aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `alias` | int64 |  | The alias email address. |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

#### `gsm userAliases list`

Lists all aliases for a user.

##### `gsm userAliases list batch`

Batch lists user aliases using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `userKey_ALL` | string |  | Same as userKey but value is applied to all lines in the CSV file |

##### `gsm userAliases list recursive`

Lists user aliases by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

### `gsm userInvitations`

Manage user invitations for unmanaged accounts (Part of Cloud Identity API)

#### `gsm userInvitations cancel`

Cancels a UserInvitation that was already sent.

##### `gsm userInvitations cancel batch`

Batch cancels user invitations using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `name` | int64 |  | UserInvitation name in the format customers/{customer}/userinvitations/{user_email_address} |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm userInvitations get`

Retrieves a UserInvitation resource.

##### `gsm userInvitations get batch`

Batch gets user invitations using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | UserInvitation name in the format customers/{customer}/userinvitations/{user_email_address} |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm userInvitations isInvitableUser`

Verifies whether a user account is eligible to receive a UserInvitation (is an unmanaged account).

##### `gsm userInvitations isInvitableUser batch`

Batch checks if users are invitable using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `name` | int64 |  | UserInvitation name in the format customers/{customer}/userinvitations/{user_email_address} |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

#### `gsm userInvitations list`

Retrieves a list of UserInvitation resources.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filter` | string |  | A query string for filtering UserInvitation results by their current state, in the format: "state=='invited'". |
| `orderBy` | string |  | The sort order of the list results.  You can sort the results in descending order based on either email or last updat... |
| `parent` | string |  | The customer ID of the Google Workspace or Cloud Identity account the UserInvitation resources are associated with. |

#### `gsm userInvitations send`

Sends a UserInvitation to email.

##### `gsm userInvitations send batch`

Batch sends user invitations using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `name` | int64 |  | UserInvitation name in the format customers/{customer}/userinvitations/{user_email_address} |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |

### `gsm userPhotos`

Manage user photos (Part of Admin SDK API)

#### `gsm userPhotos delete`

Removes the user's photo.

##### `gsm userPhotos delete batch`

Batch remove user photos using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm userPhotos delete recursive`

Removes user photos by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm userPhotos get`

Retrieves the user's photo.

##### `gsm userPhotos get batch`

Batch get user photos using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm userPhotos get recursive`

Gets user photos by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm userPhotos update`

Adds a photo for the user.

##### `gsm userPhotos update batch`

Batch remove user photos using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `photo` | int64 |  | Path to the photo file. Allowed formats are: jpeg, png, gif, bmp and tiff. |
| `photo_ALL` | string |  | Same as photo but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm userPhotos update recursive`

Updates user photos by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `photo` | string | ✓ | Path to the photo file. Allowed formats are: jpeg, png, gif, bmp and tiff. |

### `gsm userUsageReports`

Manage (get) User Usage Reports (Part of Admin SDK API)

#### `gsm userUsageReports get`

Retrieves a report which is a collection of properties and statistics for a set of users with the account.
For more information, see the User Usage Report guide.
For more information about the user report's parameters, see the Users Usage parameters reference guides.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customerId` | string |  | The unique ID of the customer to retrieve data for. |
| `date` | string |  | Represents the date the usage occurred. The timestamp is in the ISO 8601 format, yyyy-mm-dd. We recommend you use you... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `filters` | string |  | The filters query string is a comma-separated list of an application's event parameters where the parameter's value i... |
| `groupIdFilter` | string |  | Comma separated group ids (obfuscated) on which user activities are filtered, i.e, the response will contain activiti... |
| `orgUnitId` | string |  | ID of the organizational unit to report on. Activity records will be shown only for users who belong to the specified... |
| `parameters` | string |  | The parameters query string is a comma-separated list of event parameters that refine a report's results. The paramet... |
| `userKey` | string |  | Represents the profile ID or the user email for which the data should be filtered. Can be "all" for all information, ... |

### `gsm users`

Manage Users (Park of Admin SDK)

#### `gsm users delete`

Deletes a user.

##### `gsm users delete batch`

Batch deletes users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm users delete recursive`

Deletes users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm users get`

Retrieves a user.

##### `gsm users get batch`

Batch retrieves users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `customFieldMask` | int64 |  | A comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when proje... |
| `customFieldMask_ALL` | string |  | Same as customFieldMask but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `projection` | int64 |  | What subset of fields to fetch for this user.  Acceptable values are: basic   - Do not include any custom fields for ... |
| `projection_ALL` | string |  | Same as projection but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `viewType` | int64 |  | Whether to fetch the administrator-only or domain-wide public view of the user. For more information, see https://dev... |
| `viewType_ALL` | string |  | Same as viewType but value is applied to all lines in the CSV file |

##### `gsm users get recursive`

Gets users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `customFieldMask` | string |  | A comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when proje... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `projection` | string |  | What subset of fields to fetch for this user.  Acceptable values are: basic   - Do not include any custom fields for ... |
| `viewType` | string |  | Whether to fetch the administrator-only or domain-wide public view of the user. For more information, see https://dev... |

#### `gsm users insert`

Creates a user.

##### `gsm users insert batch`

Batch creates users using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addresses` | int64 |  | Specifies addresses for the user. May be used multiple times in the form of: '--addresses "country=...;countryCode=..... |
| `addresses_ALL` | stringSlice |  | Same as addresses but value is applied to all lines in the CSV file |
| `archived` | int64 |  | Indicates if user is archived. |
| `archived_ALL` | bool |  | Same as archived but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `changePasswordAtNextLogin` | int64 |  | Indicates if the user is forced to change their password at next login. This setting doesn't apply when the user sign... |
| `changePasswordAtNextLogin_ALL` | bool |  | Same as changePasswordAtNextLogin but value is applied to all lines in the CSV file |
| `customGender` | int64 |  | Custom gender. |
| `customGender_ALL` | string |  | Same as customGender but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `emails` | int64 |  | Specifies email addresses for the user. May be used multiple times in the form of: '--emails "address=...;customType=... |
| `emails_ALL` | stringSlice |  | Same as emails but value is applied to all lines in the CSV file |
| `externalIds` | int64 |  | Specifies externalIds for the user. May be used multiple times in the form of: '--externalIds "customType=...;type=..... |
| `externalIds_ALL` | stringSlice |  | Same as externalIds but value is applied to all lines in the CSV file |
| `familyName` | int64 |  | The user's last name. Required when creating a user account. |
| `familyName_ALL` | string |  | Same as familyName but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `genderType` | int64 |  | Gender. Acceptable values are: female male other unknown |
| `genderType_ALL` | string |  | Same as genderType but value is applied to all lines in the CSV file |
| `givenName` | int64 |  | The user's first name. Required when creating a user account. |
| `givenName_ALL` | string |  | Same as givenName but value is applied to all lines in the CSV file |
| `hashFunction` | int64 |  | Stores the hash format of the password property. We recommend sending the password property value as a base 16 bit he... |
| `hashFunction_ALL` | string |  | Same as hashFunction but value is applied to all lines in the CSV file |
| `ims` | int64 |  | The user's Instant Messenger (IM) accounts. A user account can have multiple ims properties. But, only one of these i... |
| `ims_ALL` | stringSlice |  | Same as ims but value is applied to all lines in the CSV file |
| `includeInGlobalAddressList` | int64 |  | Indicates if the user's profile is visible in the Workspace global address list when the contact sharing feature is e... |
| `includeInGlobalAddressList_ALL` | bool |  | Same as includeInGlobalAddressList but value is applied to all lines in the CSV file |
| `ipWhitelisted` | int64 |  | If true, the user's IP address is white listed. |
| `ipWhitelisted_ALL` | bool |  | Same as ipWhitelisted but value is applied to all lines in the CSV file |
| `keywords` | int64 |  | The user's keywords. The maximum allowed data size for this field is 1Kb. May be used multiple times in the form of: ... |
| `keywords_ALL` | stringSlice |  | Same as keywords but value is applied to all lines in the CSV file |
| `languages` | int64 |  | The user's languages. The maximum allowed data size for this field is 1Kb. May be used multiple times in the form of:... |
| `languages_ALL` | stringSlice |  | Same as languages but value is applied to all lines in the CSV file |
| `locations` | int64 |  | The user's locations. The maximum allowed data size for this field is 10Kb. May be used multiple times in the form of... |
| `locations_ALL` | stringSlice |  | Same as locations but value is applied to all lines in the CSV file |
| `notesContentType` | int64 |  | Content type of note, either plain text or HTML. Default is plain text. Possible values are: text_plain text_html |
| `notesContentType_ALL` | string |  | Same as notesContentType but value is applied to all lines in the CSV file |
| `notesValue` | int64 |  | Contents of notes. |
| `notesValue_ALL` | string |  | Same as notesValue but value is applied to all lines in the CSV file |
| `orgUnitPath` | int64 |  | The full path of the parent organization associated with the user. If the parent organization is the top-level, it is... |
| `orgUnitPath_ALL` | string |  | Same as orgUnitPath but value is applied to all lines in the CSV file |
| `organizations` | int64 |  | A list of organizations the user belongs to. The maximum allowed data size for this field is 10Kb. May be used multip... |
| `organizations_ALL` | stringSlice |  | Same as organizations but value is applied to all lines in the CSV file |
| `password` | int64 |  | Stores the password for the user account. The user's password value is required when creating a user account. It is o... |
| `password_ALL` | string |  | Same as password but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `phones` | int64 |  | A list of the user's phone numbers. The maximum allowed data size for this field is 1Kb. May be used multiple times i... |
| `phones_ALL` | stringSlice |  | Same as phones but value is applied to all lines in the CSV file |
| `posixAccounts` | int64 |  | A list of POSIX account information for the user. May be used multiple times in the form of: '--posixAccounts "accoun... |
| `posixAccounts_ALL` | stringSlice |  | Same as posixAccounts but value is applied to all lines in the CSV file |
| `primaryEmail` | int64 |  | The user's primary email address. This property is required in a request to create a user account. The primaryEmail m... |
| `recoveryEmail` | int64 |  | Recovery email of the user. |
| `recoveryEmail_ALL` | string |  | Same as recoveryEmail but value is applied to all lines in the CSV file |
| `recoveryPhone` | int64 |  | Recovery phone of the user. The phone number must be in the E.164 format, starting with the plus sign (+). Example: +... |
| `recoveryPhone_ALL` | string |  | Same as recoveryPhone but value is applied to all lines in the CSV file |
| `relations` | int64 |  | A list of the user's relationships to other users. The maximum allowed data size for this field is 2Kb. May be used m... |
| `relations_ALL` | stringSlice |  | Same as relations but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sshPublicKeys` | int64 |  | A list of SSH public keys. May be used multiple times in the form of: '--sshPublicKeys "expirationTimeUsec=...;key=..... |
| `sshPublicKeys_ALL` | stringSlice |  | Same as sshPublicKeys but value is applied to all lines in the CSV file |
| `suspended` | int64 |  | Indicates if user is suspended. |
| `suspended_ALL` | bool |  | Same as suspended but value is applied to all lines in the CSV file |
| `websites` | int64 |  | The user's websites. The maximum allowed data size for this field is 2Kb. May be used multiple times in the form of: ... |
| `websites_ALL` | stringSlice |  | Same as websites but value is applied to all lines in the CSV file |

#### `gsm users list`

Retrieves a paginated list of either deleted users or all users in a domain.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `customFieldMask` | string |  | A comma-separated list of schema names. All fields from these schemas are fetched. This should only be set when proje... |
| `customer` | string |  | The unique ID for the customer's Workspace account. In case of a multi-domain account, to fetch all groups for a cust... |
| `domain` | string |  | The domain name. Use this field to get fields from only one domain. To return all domains for a customer account, use... |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `orderBy` | string |  | Property to use for sorting results. Acceptable values are: email       - Primary email of the user. familyName  - Us... |
| `projection` | string |  | What subset of fields to fetch for this user.  Acceptable values are: basic   - Do not include any custom fields for ... |
| `query` | string |  | Query string for searching user fields. For more information on constructing user queries, see https://developers.goo... |
| `showDeleted` | bool |  | If set to true, retrieves the list of deleted users. |
| `sortOrder` | string |  | Whether to return results in ascending or descending order. Acceptable values are: ASCENDING   - Ascending order. DES... |
| `viewType` | string |  | Whether to fetch the administrator-only or domain-wide public view of the user. For more information, see https://dev... |

#### `gsm users makeAdmin`

(un)Makes a user a super administrator.

##### `gsm users makeAdmin batch`

Batch makes users admins using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `unmake` | int64 |  | Use to remove admin access. |
| `unmake_ALL` | bool |  | Same as unmake but value is applied to all lines in the CSV file |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm users makeAdmin recursive`

Grants or removes the Super Admin role to/from users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `unmake` | bool |  | Use to remove admin access. |

#### `gsm users signOut`

Sign a user out of all web and device sessions and reset their sign-in cookies.
User will have to sign in by authenticating again.

##### `gsm users signOut batch`

Batch signs out users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm users signOut recursive`

Signs out users from all devices by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm users undelete`

Undeletes a deleted user.

##### `gsm users undelete batch`

Batch undeletes users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `orgUnitPath` | int64 |  | The full path of the parent organization associated with the user. If the parent organization is the top-level, it is... |
| `orgUnitPath_ALL` | string |  | Same as orgUnitPath but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

#### `gsm users update`

Updates a user using patch semantics.

##### `gsm users update batch`

Batch updates users using a CSV file as input

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addresses` | int64 |  | Specifies addresses for the user. May be used multiple times in the form of: '--addresses "country=...;countryCode=..... |
| `addresses_ALL` | stringSlice |  | Same as addresses but value is applied to all lines in the CSV file |
| `archived` | int64 |  | Indicates if user is archived. |
| `archived_ALL` | bool |  | Same as archived but value is applied to all lines in the CSV file |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `changePasswordAtNextLogin` | int64 |  | Indicates if the user is forced to change their password at next login. This setting doesn't apply when the user sign... |
| `changePasswordAtNextLogin_ALL` | bool |  | Same as changePasswordAtNextLogin but value is applied to all lines in the CSV file |
| `customGender` | int64 |  | Custom gender. |
| `customGender_ALL` | string |  | Same as customGender but value is applied to all lines in the CSV file |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `emails` | int64 |  | Specifies email addresses for the user. May be used multiple times in the form of: '--emails "address=...;customType=... |
| `emails_ALL` | stringSlice |  | Same as emails but value is applied to all lines in the CSV file |
| `externalIds` | int64 |  | Specifies externalIds for the user. May be used multiple times in the form of: '--externalIds "customType=...;type=..... |
| `externalIds_ALL` | stringSlice |  | Same as externalIds but value is applied to all lines in the CSV file |
| `familyName` | int64 |  | The user's last name. Required when creating a user account. |
| `familyName_ALL` | string |  | Same as familyName but value is applied to all lines in the CSV file |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `genderType` | int64 |  | Gender. Acceptable values are: female male other unknown |
| `genderType_ALL` | string |  | Same as genderType but value is applied to all lines in the CSV file |
| `givenName` | int64 |  | The user's first name. Required when creating a user account. |
| `givenName_ALL` | string |  | Same as givenName but value is applied to all lines in the CSV file |
| `hashFunction` | int64 |  | Stores the hash format of the password property. We recommend sending the password property value as a base 16 bit he... |
| `hashFunction_ALL` | string |  | Same as hashFunction but value is applied to all lines in the CSV file |
| `ims` | int64 |  | The user's Instant Messenger (IM) accounts. A user account can have multiple ims properties. But, only one of these i... |
| `ims_ALL` | stringSlice |  | Same as ims but value is applied to all lines in the CSV file |
| `includeInGlobalAddressList` | int64 |  | Indicates if the user's profile is visible in the Workspace global address list when the contact sharing feature is e... |
| `includeInGlobalAddressList_ALL` | bool |  | Same as includeInGlobalAddressList but value is applied to all lines in the CSV file |
| `ipWhitelisted` | int64 |  | If true, the user's IP address is white listed. |
| `ipWhitelisted_ALL` | bool |  | Same as ipWhitelisted but value is applied to all lines in the CSV file |
| `keywords` | int64 |  | The user's keywords. The maximum allowed data size for this field is 1Kb. May be used multiple times in the form of: ... |
| `keywords_ALL` | stringSlice |  | Same as keywords but value is applied to all lines in the CSV file |
| `languages` | int64 |  | The user's languages. The maximum allowed data size for this field is 1Kb. May be used multiple times in the form of:... |
| `languages_ALL` | stringSlice |  | Same as languages but value is applied to all lines in the CSV file |
| `locations` | int64 |  | The user's locations. The maximum allowed data size for this field is 10Kb. May be used multiple times in the form of... |
| `locations_ALL` | stringSlice |  | Same as locations but value is applied to all lines in the CSV file |
| `notesContentType` | int64 |  | Content type of note, either plain text or HTML. Default is plain text. Possible values are: text_plain text_html |
| `notesContentType_ALL` | string |  | Same as notesContentType but value is applied to all lines in the CSV file |
| `notesValue` | int64 |  | Contents of notes. |
| `notesValue_ALL` | string |  | Same as notesValue but value is applied to all lines in the CSV file |
| `orgUnitPath` | int64 |  | The full path of the parent organization associated with the user. If the parent organization is the top-level, it is... |
| `orgUnitPath_ALL` | string |  | Same as orgUnitPath but value is applied to all lines in the CSV file |
| `organizations` | int64 |  | A list of organizations the user belongs to. The maximum allowed data size for this field is 10Kb. May be used multip... |
| `organizations_ALL` | stringSlice |  | Same as organizations but value is applied to all lines in the CSV file |
| `password` | int64 |  | Stores the password for the user account. The user's password value is required when creating a user account. It is o... |
| `password_ALL` | string |  | Same as password but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `phones` | int64 |  | A list of the user's phone numbers. The maximum allowed data size for this field is 1Kb. May be used multiple times i... |
| `phones_ALL` | stringSlice |  | Same as phones but value is applied to all lines in the CSV file |
| `posixAccounts` | int64 |  | A list of POSIX account information for the user. May be used multiple times in the form of: '--posixAccounts "accoun... |
| `posixAccounts_ALL` | stringSlice |  | Same as posixAccounts but value is applied to all lines in the CSV file |
| `primaryEmail` | int64 |  | The user's primary email address. This property is required in a request to create a user account. The primaryEmail m... |
| `recoveryEmail` | int64 |  | Recovery email of the user. |
| `recoveryEmail_ALL` | string |  | Same as recoveryEmail but value is applied to all lines in the CSV file |
| `recoveryPhone` | int64 |  | Recovery phone of the user. The phone number must be in the E.164 format, starting with the plus sign (+). Example: +... |
| `recoveryPhone_ALL` | string |  | Same as recoveryPhone but value is applied to all lines in the CSV file |
| `relations` | int64 |  | A list of the user's relationships to other users. The maximum allowed data size for this field is 2Kb. May be used m... |
| `relations_ALL` | stringSlice |  | Same as relations but value is applied to all lines in the CSV file |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `sshPublicKeys` | int64 |  | A list of SSH public keys. May be used multiple times in the form of: '--sshPublicKeys "expirationTimeUsec=...;key=..... |
| `sshPublicKeys_ALL` | stringSlice |  | Same as sshPublicKeys but value is applied to all lines in the CSV file |
| `suspended` | int64 |  | Indicates if user is suspended. |
| `suspended_ALL` | bool |  | Same as suspended but value is applied to all lines in the CSV file |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |
| `websites` | int64 |  | The user's websites. The maximum allowed data size for this field is 2Kb. May be used multiple times in the form of: ... |
| `websites_ALL` | stringSlice |  | Same as websites but value is applied to all lines in the CSV file |

##### `gsm users update recursive`

Updates users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `addresses` | stringSlice |  | Specifies addresses for the user. May be used multiple times in the form of: '--addresses "country=...;countryCode=..... |
| `archived` | bool |  | Indicates if user is archived. |
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `changePasswordAtNextLogin` | bool |  | Indicates if the user is forced to change their password at next login. This setting doesn't apply when the user sign... |
| `customGender` | string |  | Custom gender. |
| `familyName` | string |  | The user's last name. Required when creating a user account. |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `genderType` | string |  | Gender. Acceptable values are: female male other unknown |
| `givenName` | string |  | The user's first name. Required when creating a user account. |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `includeInGlobalAddressList` | bool |  | Indicates if the user's profile is visible in the Workspace global address list when the contact sharing feature is e... |
| `ipWhitelisted` | bool |  | If true, the user's IP address is white listed. |
| `keywords` | stringSlice |  | The user's keywords. The maximum allowed data size for this field is 1Kb. May be used multiple times in the form of: ... |
| `languages` | stringSlice |  | The user's languages. The maximum allowed data size for this field is 1Kb. May be used multiple times in the form of:... |
| `locations` | stringSlice |  | The user's locations. The maximum allowed data size for this field is 10Kb. May be used multiple times in the form of... |
| `notesContentType` | string |  | Content type of note, either plain text or HTML. Default is plain text. Possible values are: text_plain text_html |
| `notesValue` | string |  | Contents of notes. |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |
| `orgUnitPath` | string |  | The full path of the parent organization associated with the user. If the parent organization is the top-level, it is... |
| `organizations` | stringSlice |  | A list of organizations the user belongs to. The maximum allowed data size for this field is 10Kb. May be used multip... |
| `password` | string |  | Stores the password for the user account. The user's password value is required when creating a user account. It is o... |
| `phones` | stringSlice |  | A list of the user's phone numbers. The maximum allowed data size for this field is 1Kb. May be used multiple times i... |
| `recoveryEmail` | string |  | Recovery email of the user. |
| `recoveryPhone` | string |  | Recovery phone of the user. The phone number must be in the E.164 format, starting with the plus sign (+). Example: +... |
| `relations` | stringSlice |  | A list of the user's relationships to other users. The maximum allowed data size for this field is 2Kb. May be used m... |
| `suspended` | bool |  | Indicates if user is suspended. |
| `websites` | stringSlice |  | The user's websites. The maximum allowed data size for this field is 2Kb. May be used multiple times in the form of: ... |

### `gsm verificationCodes`

Manage backup Verification Codes for Users (Part of Admin SDK API)

#### `gsm verificationCodes generate`

Generate new backup verification codes for the user.

##### `gsm verificationCodes generate batch`

Batch generates backup verification codes for users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm verificationCodes generate recursive`

Generate new backup verification codes for users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm verificationCodes invalidate`

Invalidate the current backup verification codes for the user.

##### `gsm verificationCodes invalidate batch`

Batch invalidates backup verification codes for users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm verificationCodes invalidate recursive`

Invalidate the current backup verification codes for users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

#### `gsm verificationCodes list`

Returns the current set of valid backup verification codes for the specified user.

##### `gsm verificationCodes list batch`

Batch lists backup verification codes for users using a CSV file as input.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for batch commands (overrides value in config file. Max 16) |
| `delimiter` | string |  | Delimiter to use for CSV columns. Must be exactly one character. Default is ';' |
| `fields` | int64 |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `fields_ALL` | string |  | Same as fields but value is applied to all lines in the CSV file |
| `path` | string | ✓ | Path of the import file (CSV) |
| `skipHeader` | bool |  | Whether to skip the first row (header) |
| `userKey` | int64 |  | Identifies the user in the API request. The value can be the user's primary email address, alias email address, or un... |

##### `gsm verificationCodes list recursive`

Returns the current set of valid backup verification codes for users by referencing one or more organizational units and/or groups.

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `batchThreads` | int64 |  | Specify the number of threads that should be used for recursive commands (overrides value in config file. Max 16) |
| `fields` | string |  | Fields allows partial responses to be retrieved. See https://developers.google.com/gdata/docs/2.0/basics#PartialRespo... |
| `groupEmail` | stringSlice |  | An email address of a group. Can be used multiple times. Note that a group will include recursive memberships! |
| `orgUnit` | stringSlice |  | Path of an orgUnit. Can be used multiple times. Note that an orgUnit always includes all of its children! |

