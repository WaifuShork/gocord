package internal

import (
	"fmt"
	"strconv"
)

const APIVersion = "9"

// Base API endpoints, starting at "https://discord.com/api/v9"

const (
	EndpointBase    string = "https://discord.com/"
	EndpointAPIBase string = EndpointBase + "api/v" + APIVersion + "/" // https://discord.com/api/v9/

	EndpointGuilds     string = EndpointAPIBase + "guilds/"   // https://discord.com/api/v9/guilds/
	EndpointChannels   string = EndpointAPIBase + "channels/" // https://discord.com/api/v9/channels/
	EndpointUsers      string = EndpointAPIBase + "users/"    // https://discord.com/api/v9/users/
	EndpointWebhooks   string = EndpointAPIBase + "webhooks/" // https://discord.com/api/v9/webhooks/
	EndpointGateway    string = EndpointAPIBase + "gateway"   // https://discord.com/api/v9/gateway
	EndpointGatewayBot string = EndpointGateway + "/bot"      // https://discord.com/api/v9/gateway/bot

	EndpointAuthBase       string = EndpointAPIBase + "auth/"     // https://discord.com/api/v9/auth/
	EndpointLogin          string = EndpointAuthBase + "login"    // https://discord.com/api/v9/auth/login
	EndpointLogout         string = EndpointAuthBase + "logout"   // https://discord.com/api/v9/auth/logout
	EndpointForgotPassword string = EndpointAuthBase + "forgot"   // https://discord.com/api/v9/auth/forgot
	EndpointResetPassword  string = EndpointAuthBase + "reset"    // https://discord.com/api/v9/auth/reset
	EndpointRegister       string = EndpointAuthBase + "register" // https://discord.com/api/v9/auth/register
	EndpointVerify         string = EndpointAuthBase + "verify"   // https://discord.com/api/v9/auth/verify
	EndpointVerifyResend   string = EndpointVerify + "/resend"    // https://discord.com/api/v9/auth/verify/resend

	EndpointVoiceBase    string = EndpointAPIBase + "/voice/"   // https://discord.com/api/v9/voice/
	EndpointVoiceRegions string = EndpointVoiceBase + "regions" // https://discord.com/api/v9/voice/regions
	EndpointVoiceIce     string = EndpointVoiceBase + "ice"     // https://discord.com/api/v9/voice/ice

	EndpointTutorialBase       string = EndpointAPIBase + "tutorial/"       // https://discord.com/api/v9/tutorial/
	EndpointTutorialIndicators string = EndpointTutorialBase + "indicators" // https://discord.com/api/v9/tutorial/indicators

	EndpointTrack        string = EndpointAPIBase + "track"        // https://discord.com/api/v9/track
	EndpointSingleSignOn string = EndpointAPIBase + "sso"          // https://discord.com/api/v9/sso
	EndpointReport       string = EndpointAPIBase + "report"       // https://discord.com/api/v9/report
	EndpointIntegrations string = EndpointAPIBase + "integrations" // https://discord.com/api/v9/integrations

	EndpointContentBase         string = "https://cdn.discordapp.com/"
	EndpointContentAttachments  string = EndpointContentBase + "attachments/"   // https://cdn.discordapp.com/attachments/
	EndpointContentAvatars      string = EndpointContentBase + "avatars/"       // https://cdn.discordapp.com/avatars/
	EndpointContentIcons        string = EndpointContentBase + "icons/"         // https://cdn.discordapp.com/icons/
	EndpointContentSplashes     string = EndpointContentBase + "splashes/"      // https://cdn.discordapp.com/splashes/
	EndpointContentChannelIcons string = EndpointContentBase + "channel-icons/" // https://cdn.discordapp.com/channel-icons/
	EndpointContentBanners      string = EndpointContentBase + "banners/"       // https://cdn.discordapp.com/banners/
)

var (
	// https://discord.com/api/v9/users/userId
	EndpointUser               = func(userId string) string { return EndpointUsers + userId }

	// https://cdn.discordapp.com/avatars/userId/avatarId.png
	EndpointUserAvatar         = func(userId, avatarId string) string { return EndpointContentAvatars + userId + "/" + avatarId + ".png" }

	// https://cdn.discordapp.com/avatars/userId/avatarId.gif
	EndpointUserAvatarAnimated = func(userId, avatarId string) string { return EndpointContentAvatars + userId + "/" + avatarId + ".gif" }

	// https://cdn.discordapp.com/embed/avatars/
	EndpointDefaultUserAvatar  = func(userDiscriminator string) string {
		userDiscriminatorInt, err := strconv.Atoi(userDiscriminator)
		if err != nil { 
			return fmt.Sprintf("Unable to get user avatar using %s", userDiscriminator)
		}
		return EndpointContentBase + "embed/avatars/" + strconv.Itoa(userDiscriminatorInt % 5) + "png"
	}

	// https://discord.com/api/v9/users/userId/settings
	EndpointUserSettings = func(userId string) string { return EndpointUsers + userId + "/settings" }

	// https://discord.com/api/v9/users/userId/guilds
	EndpointUserGuilds = func(userId string) string { return EndpointUsers + userId + "/guilds" }

	// https://discord.com/api/v9/users/userId/guilds/guildId
	EndpointUserGuild = func(userId, guildId string) string { return EndpointUsers + userId + "/guilds/" + guildId }

	// https://discord.com/api/v9/users/userId/guilds/guildId/settings
	EndpointUserGuildSettings = func(userId, guildId string) string { return EndpointUsers + userId + "/guilds/" + guildId + "/settings" }

	// https://discord.com/api/v9/users/userId/channels
	EndpointUserChannels = func(userId string) string { return EndpointUsers + userId + "/channels" }

	// https://discord.com/api/v9/users/userId/devices
	EndpointUserDevices = func(userId string) string { return EndpointUsers + userId + "/devices" }

	// https://discord.com/api/v9/users/userId/connections
	EndpointUserConnections = func(userId string) string { return EndpointUsers + userId + "/connections" }

	// https://discord.com/api/v9/users/@me/notes/userId
	EndpointUserNotes = func(userId string) string { return EndpointUsers + "@me/notes/" + userId }

	// https://discord.com/api/v9/guilds/guildId
	EndpointGuild = func(guildId string) string { return EndpointGuilds + guildId }

	// https://discord.com/api/v9/guilds/guildId/preview 
	EndpointGuildPreview = func(guildId string) string { return EndpointGuild(guildId) + "/preview" }

	// https://discord.com/api/v9/guilds/guildId/channels
	EndpointGuildChannels = func(guildId string) string { return EndpointGuild(guildId) + "/channels" }

	// https://discord.com/api/v9/guilds/guildId/members
	EndpointGuildMembers = func(guildId string) string { return EndpointGuild(guildId) + "/members" }

	// https://discord.com/api/v9/guilds/guildId/members/userId
	EndpointGuildMember = func(guildId, userId string) string { return EndpointGuildMembers(guildId) + "/" + userId }

	// https://discord.com/api/v9/guilds/guildId/members/userId/roles/roleId
	EndpointGuildMemberRole = func(guildId, userId, roleId string) string { return EndpointGuildMember(guildId, userId) + "/roles/" + roleId }

	// https://discord.com/api/v9/guilds/guildId/bans
	EndpointGuildBans = func(guildId string) string { return EndpointGuild(guildId) + "/bans" } 

	// https://discord.com/api/v9/guilds/guildId/bans/userId
	EndpointGuildBan = func(guildId, userId string) string { return EndpointGuildBans(guildId) + "/" + userId }

	// https://discord.com/api/v9/guilds/guildId/integrations
	EndpointGuildIntegrations = func(guildId string) string { return EndpointGuild(guildId) + "/integrations" }

	// https://discord.com/api/v9/guilds/guildId/integrations/integrationId
	EndpointGuildIntegration = func(guildId, integrationId string) string { return EndpointGuildIntegrations(guildId) + "/" + integrationId }

	// https://discord.com/api/v9/guilds/guildId/integrations/integrationId/sync
	EndpointGuildIntegrationSync = func(guildId, integrationId string) string { return EndpointGuildIntegration(guildId, integrationId) + "/sync" } 

	// https://discord.com/api/v9/guilds/guildId/roles
	EndpointGuildRoles = func(guildId string) string { return EndpointGuild(guildId) + "/roles" }

	// https://discord.com/api/v9/guilds/guildId/roles/roleId
	EndpointGuildRole = func(guildId, roleId string) string { return EndpointGuildRoles(guildId) + "/" + roleId }

	// https://discord.com/api/v9/guilds/guildId/invites
	EndpointGuildInvites = func(guildId string) string { return EndpointGuild(guildId) + "/invites" }

	// https://discord.com/api/v9/guilds/guildId/widget
	EndpointGuildWidget = func(guildId string) string { return EndpointGuild(guildId) + "/widget" }

	// https://discord.com/api/v9/guilds/guildId/widget
	EndpointGuildEmbed = EndpointGuildWidget

	// https://discord.com/api/v9/guilds/guildId/prune
	EndpointGuildPrune = func(guildId string) string { return EndpointGuild(guildId) + "/prune" }

	// https://cdn.discordapp.com/icons/guildId/hash.png
	EndpointGuildIcon = func(guildId, hash string) string { return EndpointContentIcons + guildId + "/" + hash + ".png" }

	// https://cdn.discordapp.com/icons/guildId/hash.gif
	EndpointGuildIconAnimated = func(guildId, hash string) string { return EndpointContentIcons + guildId + "/" + hash + ".gif" }

	// https://cdn.discordapp.com/splashes/guildId/hash.png
	EndpointGuildSplash = func(guildId, hash string) string { return EndpointContentSplashes + guildId + "/" + hash + ".png" }

	// https://discord.com/api/v9/guilds/guildId/webhooks
	EndpointGuildWebhooks = func(guildId string) string { return EndpointGuild(guildId) + "/webhooks" }

	// https://discord.com/api/v9/guilds/guildId/audit-logs
	EndpointGuildAuditLogs = func(guildId string) string { return EndpointGuild(guildId) + "/audit-logs" }

	// https://discord.com/api/v9/guilds/guildId/emojis
	EndpointGuildEmojis = func(guildId string) string { return EndpointGuild(guildId) + "/emojis" }

	// https://discord.com/api/v9/guilds/guildId/emojis/emojiId
	EndpointGuildEmoji = func(guildId, emojiId string) string { return EndpointGuildEmojis(guildId) + "/" + emojiId }

	// https://cdn.discordapp.com/banners/guildId/hash.png
	EndpointGuildBanner = func(guildId, hash string) string { return EndpointContentBanners + guildId + "/" + hash + ".png" }

	// https://discord.com/api/v9/channels/channelId
	EndpointChannel = func(channelId string) string { return EndpointChannels + channelId }

	// https://discord.com/api/v9/channels/channelId/permissions
	EndpointChannelPermissions = func(channelId string) string { return EndpointChannel(channelId) + "/permissions" }

	// https://discord.com/api/v9/channels/channelId/permissions/teamId
	EndpointChannelPermission = func(channelId, teamId string) string { return EndpointChannelPermissions(channelId) + "/" + teamId }

	// https://discord.com/api/v9/channels/channelId/invites
	EndpointChannelInvites = func(channelId string) string { return EndpointChannel(channelId) + "invites" }

	// https://discord.com/api/v9/channels/channelId/typing
	EndpointChannelTyping = func(channelId string) string { return EndpointChannel(channelId) + "/typing" }

	// https://discord.com/api/v9/channels/channelId/messages
	EndpointChannelMessages = func(channelId string) string { return EndpointChannel(channelId) + "/messages" }

	// https://discord.com/api/v9/channels/channelId/messages/messageId
	EndpointChannelMessage = func(channelId, messageId string) string { return EndpointChannelMessages(channelId) + "/" + messageId }

	// https://discord.com/api/v9/channels/channelId/messages/messageId/ack
	EndpointChannelMessageAck = func(channelId, messageId string) string { return EndpointChannelMessage(channelId, messageId) + "/ack" }

	// https://discord.com/api/v9/channels/channelId/messages/bulk-delete
	EndpointChannelMessagesBulkDelete = func(channelId string) string { return EndpointChannelMessages(channelId) + "/bulk-delete" }  

	// https://discord.com/api/v9/channels/channelId/pins
	EndpointChannelMessagesPins = func(channelId string) string { return EndpointChannel(channelId) + "/pins" }

	// https://discord.com/api/v9/channels/channelId/pins/messageId
	EndpointChannelMessagePin = func(channelId, messageId string) string { return EndpointChannelMessagesPins(channelId) + "/" + messageId }

	// https://discord.com/api/v9/channels/channelId/messages/messageId/crosspost
	EndpointChannelMessageCrosspost = func(channelId, messageId string) string { return EndpointChannelMessage(channelId, messageId) + "/crosspost" }

	// https://discord.com/api/v9/channels/channelId/followers
	EndpointChannelFollow = func(channelId string) string { return EndpointChannel(channelId) + "/followers" }

	// https://cdn.discordapp.com/channel-icons/channelId/hash.png
	EndpointGroupIcon = func(channelId, hash string) string { return EndpointContentChannelIcons + channelId + "/" + hash + ".png" }

	// https://discord.com/api/v9/channels/channelId/webhooks
	EndpointChannelWebhooks = func(channelId string) string { return EndpointChannel(channelId) + "/webhooks" }

	// https://discord.com/api/v9/webhooks/webhookId
	EndpointWebhook = func(webhookId string) string { return EndpointWebhooks + webhookId }

	// https://discord.com/api/v9/webhooks/webhookId/token
	EndpointWebhookToken = func(webhookId, token string) string { return EndpointWebhook(webhookId) + "/" + token }

	// https://discord.com/api/v9/webhooks/webhookId/messages/messageId
	EndpointWebhookMessage = func(webhookId, token, messageId string) string { 
		return EndpointWebhookToken(webhookId, token) + "/messages/" + messageId
	} 

	// https://discord.com/api/v9/channels/channelId/messages/messageId/reactions
	EndpointMessageReactionsAll = func(channelId, messageId string) string {
		return EndpointChannelMessage(channelId, messageId) + "/reactions"
	}

	// https://discord.com/api/v9/channels/channelId/messages/messageId/reactions/emojiId
	EndpointMessageReactions = func(channelId, messageId, emojiId string) string { 
		return EndpointMessageReactionsAll(channelId, messageId) + "/" + emojiId
	}

	// https://discord.com/api/v9/channels/channelId/messages/messageId/reactions/emojiId/userId
	EndpointMessageReaction = func(channelId, messageId, emojiId, userId string) string { 
		return EndpointMessageReactions(channelId, messageId, emojiId) + "/" + userId
	}

	// https://discord.com/api/v9/applications/applicationId/commands
	EndpointApplicationGlobalCommands = func(applicationId string) string { 
		return EndpointApplication(applicationId) + "/commands"
	}

	// https://discord.com/api/v9/applications/applicationId/commands/channelId
	EndpointApplicationGlobalCommand = func(applicationId, channelId string) string { 
		return EndpointApplicationGlobalCommands(applicationId) + "/" + channelId
	}

	// https://discord.com/api/v9/applications/applicationId/guilds/guildId/commands
	EndpointApplicationGuildCommands = func(applicationId, guildId string) string { 
		return EndpointApplication(applicationId) + "/guilds/" + guildId + "/commands"
	}

	// https://discord.com/api/v9/channels/
	EndpointApplicationGuildCommand = func(applicationId, guildId, channelId string) string { 
		return EndpointApplicationGuildCommands(applicationId, guildId) + "/" + channelId
	}

	// https://discord.com/api/v9/interactions/applicationId/interactionToken
	EndpointInteraction = func(applicationId, interactionToken string) string { 
		return EndpointAPIBase + "interactions/" + applicationId + "/" + interactionToken
	}

	// https://discord.com/api/v9/interactions/applicationId/interactionToken/callback
	EndpointInteractionResponse = func(interactionId, interactionToken string) string { 
		return EndpointInteraction(interactionId, interactionToken) + "/callback"
	}

	// https://discord.com/api/v9/webhooks/webhookId/messages/@original
	EndpointInteractionResponseActions = func(applicationId, interactionToken string) string { 
		return EndpointWebhookMessage(applicationId, interactionToken, "@original")
	}

	// https://discord.com/api/v9/webhooks/webhookId/token
	EndpointFollowupMessage = func(applicationId, interactionToken string) string { 
		return EndpointWebhookToken(applicationId, interactionToken)
	}

	// https://discord.com/api/v9/webhooks/webhookId/messages/messageId
	EndpointFollowupMessageActions = func(applicationId, interactionToken, messageId string) string { 
		return EndpointWebhookMessage(applicationId, interactionToken, messageId)
	}

	// https://discord.com/api/v9/users/@me/relationships
	EndpointRelationships = func() string { return EndpointUsers + "@me" + "/relationships" }

	// https://discord.com/api/v9/users/@me/relationships/userId
	EndpointRelationship = func(userId string) string { return  EndpointRelationships() + "/" + userId }

	// https://discord.com/api/v9/users/userId/relationships
	EndpointRelationshipMutual = func(userId string) string { return EndpointUsers + userId + "/relationships" }

	// https://discord.com/api/v9/guilds
	EndpointGuildCreate = EndpointAPIBase + "guilds"

	// https://discord.com/api/v9/invites/inviteId
	EndpointInvite = func(inviteId string) string { return EndpointAPIBase + "invites/" + inviteId }

	// https://discord.com/api/v9/integrations/integrationId/join
	EndpointIntegrationsJoin = func(integrationId string) string { return EndpointIntegrations + "/" + integrationId + "/join"}

	// https://cdn.discordapp.com/emojis/emojiId.png
	EndpointEmoji = func(emojiId string) string { return EndpointContentBase + "emojis/" + emojiId + ".png" }

	// https://cdn.discordapp.com/emojis/emojiId.gif
	EndpointEmojiAnimated = func(emojiId string) string { return EndpointContentBase + "emojis/" + emojiId + ".gif" }

	// https://discord.com/api/v9/applications
	EndpointApplications = EndpointAPIBase + "applications"

	// https://discord.com/api/v9/applications/applicationId 
	EndpointApplication = func(applicationId string) string { return EndpointApplications + "/" + applicationId }

	// https://discord.com/api/v9/oauth2/
	EndpointOAuth2 = EndpointAPIBase + "oauth2/"

	// https://discord.com/api/v9/oauth2/applications
	EndpointOAuth2Applications = EndpointOAuth2 + "applications"

	// https://discord.com/api/v9/oauth2/applications/applicationId
	EndpointOAuth2Application = func(applicationId string) string { return EndpointOAuth2Applications + "/" + applicationId } 

	// https://discord.com/api/v9/oauth2/applications/applicationId/bot
	EndpointOAuth2ApplicationsBot = func(applicationId string) string { return EndpointApplication(applicationId) + "/bot" }

	// https://discord.com/api/v9/oauth2/applications/applicationId/assets
	EndpointOAuth2ApplicationsAssets = func(applicationId string) string { return EndpointApplication(applicationId) + "/assets" }
)