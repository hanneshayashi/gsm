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
	"google.golang.org/api/groupssettings/v1"
)

// groupSettingsCmd represents the groupSettings command
var groupSettingsCmd = &cobra.Command{
	Use:   "groupSettings",
	Short: "Manage Group Settings (Part of Admin SDK)",
	Long:  "https://developers.google.com/admin-sdk/groups-settings/v1/reference/groups",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func ignoreDeprecatedGroupSettings(groupSettings *groupssettings.Groups) *groupssettings.Groups {
	groupSettings.WhoCanInvite = ""
	groupSettings.WhoCanAdd = ""
	groupSettings.MaxMessageBytes = 0
	groupSettings.ShowInGroupDirectory = ""
	groupSettings.AllowGoogleCommunication = ""
	groupSettings.MessageDisplayFont = ""
	groupSettings.WhoCanAddReferences = ""
	groupSettings.WhoCanAssignTopics = ""
	groupSettings.WhoCanUnassignTopic = ""
	groupSettings.WhoCanTakeTopics = ""
	groupSettings.WhoCanMarkDuplicate = ""
	groupSettings.WhoCanMarkNoResponseNeeded = ""
	groupSettings.WhoCanMarkFavoriteReplyOnAnyTopic = ""
	groupSettings.WhoCanMarkFavoriteReplyOnOwnTopic = ""
	groupSettings.WhoCanUnmarkFavoriteReplyOnAnyTopic = ""
	groupSettings.WhoCanEnterFreeFormTags = ""
	groupSettings.WhoCanModifyTagsAndCategories = ""
	groupSettings.WhoCanModifyMembers = ""
	groupSettings.WhoCanApproveMessages = ""
	groupSettings.WhoCanDeleteAnyPost = ""
	groupSettings.WhoCanDeleteTopics = ""
	groupSettings.WhoCanLockTopics = ""
	groupSettings.WhoCanMoveTopicsIn = ""
	groupSettings.WhoCanMoveTopicsOut = ""
	groupSettings.WhoCanPostAnnouncements = ""
	groupSettings.WhoCanHideAbuse = ""
	groupSettings.WhoCanMakeTopicsSticky = ""
	return groupSettings
}

var groupSettingFlags map[string]*gsmhelpers.Flag = map[string]*gsmhelpers.Flag{
	"groupUniqueId": {
		AvailableFor:   []string{"get", "patch"},
		Type:           "string",
		Description:    `The group's email address.`,
		Required:       []string{"get", "patch"},
		ExcludeFromAll: true,
	},
	"whoCanJoin": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Permission to join group.
[ANYONE_CAN_JOIN|ALL_IN_DOMAIN_CAN_JOIN|INVITED_CAN_JOIN|CAN_REQUEST_TO_JOIN]
ANYONE_CAN_JOIN         - Any Internet user, both inside and outside your domain, can join the group.
ALL_IN_DOMAIN_CAN_JOIN  - Anyone in the account domain can join. This includes accounts with multiple domains.
INVITED_CAN_JOIN        - Candidates for membership can be invited to join.
CAN_REQUEST_TO_JOIN     - Non members can request an invitation to join.`,
	},
	"whoCanViewMembership": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Permissions to view membership.
[ALL_IN_DOMAIN_CAN_VIEW|ALL_MEMBERS_CAN_VIEW|ALL_MANAGERS_CAN_VIEW]
ALL_IN_DOMAIN_CAN_VIEW  - Anyone in the account can view the group members list.
                          If a group already has external members, those members can still send email to this group.
ALL_MEMBERS_CAN_VIEW    - The group members can view the group members list.
ALL_MANAGERS_CAN_VIEW   - The group managers can view group members list.`,
	},
	"whoCanViewGroup": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Permissions to view group messages.
[ANYONE_CAN_VIEW|ALL_IN_DOMAIN_CAN_VIEW|ALL_MEMBERS_CAN_VIEW|ALL_OWNERS_CAN_VIEW]
ANYONE_CAN_VIEW         - Any Internet user can view the group's messages.
ALL_IN_DOMAIN_CAN_VIEW  - Anyone in your account can view this group's messages.
ALL_MEMBERS_CAN_VIEW    - All group members can view the group's messages.
ALL_MANAGERS_CAN_VIEW   - Any group manager can view this group's messages.
ALL_OWNERS_CAN_VIEW     - Any group owner can view this group's messages.`,
	},
	"allowExternalMembers": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Identifies whether members external to your organization can join the group.
true   - Workspace users external to your organization can become members of this group.
false  - Users not belonging to the organization are not allowed to become members of this group.`,
	},
	"whoCanPostMessage": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Permissions to post messages.
[NONE_CAN_POST|ALL_MANAGERS_CAN_POST|ALL_MEMBERS_CAN_POST|ALL_OWNERS_CAN_POST|ALL_IN_DOMAIN_CAN_POST|ANYONE_CAN_POST]
NONE_CAN_POST           - The group is disabled and archived. No one can post a message to this group.
                            - When archiveOnly is false, updating whoCanPostMessage to NONE_CAN_POST, results in an error.
                            - If archiveOnly is reverted from true to false, whoCanPostMessages is set to ALL_MANAGERS_CAN_POST.
ALL_MANAGERS_CAN_POST   - Managers, including group owners, can post messages.
ALL_MEMBERS_CAN_POST    - Any group member can post a message.
ALL_OWNERS_CAN_POST     - Only group owners can post a message.
ALL_IN_DOMAIN_CAN_POST  - Anyone in the account can post a message.
ANYONE_CAN_POST         - Any Internet user who outside your account can access your Google Groups service and post a message.`,
	},
	"allowWebPosting": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Allows posting from web.
true   - Allows any member to post to the group forum.
false  - Members only use Gmail to communicate with the group.`,
	},
	"primaryLanguage": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `The primary language for group.
For a group's primary language use the language tags from the Workspace languages found at Workspace Email Settings API Email Language Tags.`,
	},
	"isArchived": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Allows the Group contents to be archived.
true   - Archive messages sent to the group.
false  - Do not keep an archive of messages sent to this group.
         If false, previously archived messages remain in the archive.`,
	},
	"archiveOnly": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Allows the group to be archived only.
true   - Group is archived and the group is inactive. New messages to this group are rejected. The older archived messages are browseable and searchable.
           - If true, the whoCanPostMessage property is set to NONE_CAN_POST.
           - If reverted from true to false, whoCanPostMessages is set to ALL_MANAGERS_CAN_POST.
false  - The group is active and can receive messages.
		   - When false, updating whoCanPostMessage to NONE_CAN_POST, results in an error.`,
	},
	"messageModerationLevel": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Moderation level of incoming messages.
[MODERATE_ALL_MESSAGES|MODERATE_NON_MEMBERS|MODERATE_NEW_MEMBERS|MODERATE_NONE]
MODERATE_ALL_MESSAGES  - All messages are sent to the group owner's email address for approval. If approved, the message is sent to the group.
MODERATE_NON_MEMBERS   - All messages from non group members are sent to the group owner's email address for approval. If approved, the message is sent to the group.
MODERATE_NEW_MEMBERS   - All messages from new members are sent to the group owner's email address for approval. If approved, the message is sent to the group.
MODERATE_NONE          - No moderator approval is required. Messages are delivered directly to the group.`,
	},
	"spamModerationLevel": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies moderation levels for messages detected as spam.
[ALLOW|MODERATE|SILENTLY_MODERATE|REJECT]
ALLOW              - Post the message to the group.
MODERATE           - Send the message to the moderation queue. This is the default.
SILENTLY_MODERATE  - Send the message to the moderation queue, but do not send notification to moderators.
REJECT             - Immediately reject the message.`,
	},
	"replyTo": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who the default reply should go to.
[REPLY_TO_CUSTOM|REPLY_TO_SENDER|REPLY_TO_LIST|REPLY_TO_OWNER|REPLY_TO_IGNORE|REPLY_TO_MANAGERS]
REPLY_TO_CUSTOM    - For replies to messages, use the group's custom email address.
					   - When the group's ReplyTo property is set to REPLY_TO_CUSTOM, the customReplyTo property holds the custom email address used when replying to a message.
						 If the group's ReplyTo property is set to REPLY_TO_CUSTOM, the customReplyTo property must have a value.
						 Otherwise an error is returned.

REPLY_TO_SENDER    - The reply sent to author of message.
REPLY_TO_LIST      - This reply message is sent to the group.
REPLY_TO_OWNER     - The reply is sent to the owners of the group. This doesn't include the group's managers.
REPLY_TO_IGNORE    - Group users individually decide where the message reply is sent.
REPLY_TO_MANAGERS  - This reply message is sent to the group's managers, which includes all managers and the group owner.`,
	},
	"customReplyTo": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `An email address used when replying to a message if the replyTo property is set to REPLY_TO_CUSTOM.
This address is defined by an account administrator.
When the group's ReplyTo property is set to REPLY_TO_CUSTOM, the customReplyTo property holds a custom email address used when replying to a message.
If the group's ReplyTo property is set to REPLY_TO_CUSTOM, the customReplyTo property must have a text value or an error is returned.`,
	},
	"includeCustomFooter": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `Whether to include custom footer.`,
	},
	"customFooterText": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Set the content of custom footer text.
The maximum number of characters is 1000.`,
	},
	"sendMessageDenyNotification": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Allows a member to be notified if the member's message to the group is denied by the group owner.
true   - When a message is rejected, send the deny message notification to the message author.
         The defaultMessageDenyNotificationText property is dependent on the sendMessageDenyNotification property being true.

false  - When a message is rejected, no notification is sent.`,
	},
	"defaultMessageDenyNotificationText": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `When a message is rejected, this is text for the rejection notification sent to the message's author.
By default, this property is empty and has no value in the API's response body.
The maximum notification text size is 10,000 characters.`,
	},
	"membersCanPostAsTheGroup": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Enables members to post messages as the group.
true   - Group member can post messages using the group's email address instead of their own email address.
         Message appear to originate from the group itself.
         Note: When true, any message moderation settings on individual users or new members do not apply to posts made on behalf of the group.
false  - Members can not post in behalf of the group's email address.`,
	},
	"includeInGlobalAddressList": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Enables the group to be included in the Global Address List. For more information, see the help center.
true   - Group is included in the Global Address List.
false  - Group is not included in the Global Address List.`,
	},
	"whoCanLeaveGroup": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who can leave the group.
[ALL_MANAGERS_CAN_LEAVE|ALL_MEMBERS_CAN_LEAVE|NONE_CAN_LEAVE]`,
	},
	"whoCanContactOwner": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who can contact the group owner.
[ALL_IN_DOMAIN_CAN_CONTACT|ALL_MANAGERS_CAN_CONTACT|ALL_MEMBERS_CAN_CONTACT|ANYONE_CAN_CONTACT]`,
	},
	"favoriteRepliesOnTop": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Indicates if favorite replies should be displayed above other replies.
true   - Favorite replies will be displayed above other replies.
false  - Favorite replies will not be displayed above other replies.`,
	},
	"whoCanModerateMembers": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who can manage members.
[ALL_MEMBERS|OWNERS_AND_MANAGERS|OWNERS_ONLY|NONE]`,
	},
	"whoCanModerateContent": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who can moderate content.
[ALL_MEMBERS|OWNERS_AND_MANAGERS|OWNERS_ONLY|NONE]`,
	},
	"whoCanAssistContent": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who can moderate metadata.
[ALL_MEMBERS|OWNERS_AND_MANAGERS|OWNERS_ONLY|NONE]`,
	},
	"enableCollaborativeInbox": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description:  `Specifies whether a collaborative inbox will remain turned on for the group.`,
	},
	"whoCanDiscoverGroup": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies the set of users for whom this group is discoverable.
[ANYONE_CAN_DISCOVER|ALL_IN_DOMAIN_CAN_DISCOVER|ALL_MEMBERS_CAN_DISCOVER]`,
	},
	"whoCanApproveMembers": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who can approve members who ask to join groups.
This permission will be deprecated once it is merged into the new whoCanModerateMembers setting.
[ALL_MEMBERS_CAN_APPROVE]ALL_MANAGERS_CAN_APPROVE]ALL_OWNERS_CAN_APPROVE|NONE_CAN_APPROVE]`,
	},
	"whoCanBanUsers": {
		AvailableFor: []string{"patch"},
		Type:         "string",
		Description: `Specifies who can deny membership to users.
This permission will be deprecated once it is merged into the new whoCanModerateMembers setting.
[ALL_MEMBERS|OWNERS_AND_MANAGERS|OWNERS_ONLY|NONE]`,
	},
	"fields": {
		AvailableFor: []string{"get", "patch"},
		Type:         "string",
		Description: `Fields allows partial responses to be retrieved.
See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse for more information.`,
	},
	"ignoreDeprecated": {
		AvailableFor: []string{"get", "patch"},
		Type:         "bool",
		Description:  `Ignore deprecated fields.`,
		Defaults:     map[string]interface{}{"get": true, "patch": true},
	},
}
var groupSettingFlagsALL = gsmhelpers.GetAllFlags(groupSettingFlags)

func init() {
	rootCmd.AddCommand(groupSettingsCmd)
}

func mapToGroupSettings(flags map[string]*gsmhelpers.Value) (*groupssettings.Groups, error) {
	groups := &groupssettings.Groups{}
	if flags["whoCanJoin"].IsSet() {
		groups.WhoCanJoin = flags["whoCanJoin"].GetString()
		if groups.WhoCanJoin == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanJoin")
		}
	}
	if flags["whoCanViewMembership"].IsSet() {
		groups.WhoCanViewMembership = flags["whoCanViewMembership"].GetString()
		if groups.WhoCanViewMembership == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanViewMembership")
		}
	}
	if flags["whoCanViewGroup"].IsSet() {
		groups.WhoCanViewGroup = flags["whoCanViewGroup"].GetString()
		if groups.WhoCanViewGroup == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanViewGroup")
		}
	}
	if flags["allowExternalMembers"].IsSet() {
		groups.AllowExternalMembers = flags["allowExternalMembers"].GetString()
		if groups.AllowExternalMembers == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "AllowExternalMembers")
		}
	}
	if flags["whoCanPostMessage"].IsSet() {
		groups.WhoCanPostMessage = flags["whoCanPostMessage"].GetString()
		if groups.WhoCanPostMessage == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanPostMessage")
		}
	}
	if flags["allowWebPosting"].IsSet() {
		groups.AllowWebPosting = flags["allowWebPosting"].GetString()
		if groups.AllowWebPosting == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "AllowWebPosting")
		}
	}
	if flags["primaryLanguage"].IsSet() {
		groups.PrimaryLanguage = flags["primaryLanguage"].GetString()
		if groups.PrimaryLanguage == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "PrimaryLanguage")
		}
	}
	if flags["isArchived"].IsSet() {
		groups.IsArchived = flags["isArchived"].GetString()
		if groups.IsArchived == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "IsArchived")
		}
	}
	if flags["archiveOnly"].IsSet() {
		groups.ArchiveOnly = flags["archiveOnly"].GetString()
		if groups.ArchiveOnly == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "ArchiveOnly")
		}
	}
	if flags["messageModerationLevel"].IsSet() {
		groups.MessageModerationLevel = flags["messageModerationLevel"].GetString()
		if groups.MessageModerationLevel == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "MessageModerationLevel")
		}
	}
	if flags["spamModerationLevel"].IsSet() {
		groups.SpamModerationLevel = flags["spamModerationLevel"].GetString()
		if groups.SpamModerationLevel == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "SpamModerationLevel")
		}
	}
	if flags["replyTo"].IsSet() {
		groups.ReplyTo = flags["replyTo"].GetString()
		if groups.ReplyTo == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "ReplyTo")
		}
	}
	if flags["customReplyTo"].IsSet() {
		groups.CustomReplyTo = flags["customReplyTo"].GetString()
		if groups.CustomReplyTo == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "CustomReplyTo")
		}
	}
	if flags["includeCustomFooter"].IsSet() {
		groups.IncludeCustomFooter = flags["includeCustomFooter"].GetString()
		if groups.IncludeCustomFooter == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "IncludeCustomFooter")
		}
	}
	if flags["customFooterText"].IsSet() {
		groups.CustomFooterText = flags["customFooterText"].GetString()
		if groups.CustomFooterText == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "CustomFooterText")
		}
	}
	if flags["sendMessageDenyNotification"].IsSet() {
		groups.SendMessageDenyNotification = flags["sendMessageDenyNotification"].GetString()
		if groups.SendMessageDenyNotification == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "SendMessageDenyNotification")
		}
	}
	if flags["defaultMessageDenyNotificationText"].IsSet() {
		groups.DefaultMessageDenyNotificationText = flags["defaultMessageDenyNotificationText"].GetString()
		if groups.DefaultMessageDenyNotificationText == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "DefaultMessageDenyNotificationText")
		}
	}
	if flags["membersCanPostAsTheGroup"].IsSet() {
		groups.MembersCanPostAsTheGroup = flags["membersCanPostAsTheGroup"].GetString()
		if groups.MembersCanPostAsTheGroup == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "MembersCanPostAsTheGroup")
		}
	}
	if flags["includeInGlobalAddressList"].IsSet() {
		groups.IncludeInGlobalAddressList = flags["includeInGlobalAddressList"].GetString()
		if groups.IncludeInGlobalAddressList == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "IncludeInGlobalAddressList")
		}
	}
	if flags["whoCanLeaveGroup"].IsSet() {
		groups.WhoCanLeaveGroup = flags["whoCanLeaveGroup"].GetString()
		if groups.WhoCanLeaveGroup == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanLeaveGroup")
		}
	}
	if flags["whoCanContactOwner"].IsSet() {
		groups.WhoCanContactOwner = flags["whoCanContactOwner"].GetString()
		if groups.WhoCanContactOwner == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanContactOwner")
		}
	}
	if flags["favoriteRepliesOnTop"].IsSet() {
		groups.FavoriteRepliesOnTop = flags["favoriteRepliesOnTop"].GetString()
		if groups.FavoriteRepliesOnTop == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "FavoriteRepliesOnTop")
		}
	}
	if flags["whoCanModerateMembers"].IsSet() {
		groups.WhoCanModerateMembers = flags["whoCanModerateMembers"].GetString()
		if groups.WhoCanModerateMembers == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanModerateMembers")
		}
	}
	if flags["whoCanModerateContent"].IsSet() {
		groups.WhoCanModerateContent = flags["whoCanModerateContent"].GetString()
		if groups.WhoCanModerateContent == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanModerateContent")
		}
	}
	if flags["whoCanAssistContent"].IsSet() {
		groups.WhoCanAssistContent = flags["whoCanAssistContent"].GetString()
		if groups.WhoCanAssistContent == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanAssistContent")
		}
	}
	if flags["enableCollaborativeInbox"].IsSet() {
		groups.EnableCollaborativeInbox = flags["enableCollaborativeInbox"].GetString()
		if groups.EnableCollaborativeInbox == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "EnableCollaborativeInbox")
		}
	}
	if flags["whoCanDiscoverGroup"].IsSet() {
		groups.WhoCanDiscoverGroup = flags["whoCanDiscoverGroup"].GetString()
		if groups.WhoCanDiscoverGroup == "" {
			groups.ForceSendFields = append(groups.ForceSendFields, "WhoCanDiscoverGroup")
		}
	}
	return groups, nil
}
