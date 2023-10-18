// Copyright (C) 2022 Specter Ops, Inc.
//
// This file is part of AzureHound.
//
// AzureHound is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// AzureHound is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package azure

import (
	"https://github.com/AbderrahimBouhdida/AzH/enums"
)

// Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in
// user.
type AutomaticRepliesSetting struct {
	// The set of audience external to the signed-in user's organization who will receive the {@link
	// ExternalReplyMessage}.
	ExternalAudience enums.ExternalAudienceScope `json:"externalAudience"`

	// The automatic reply to send to the specified eternal audience.
	ExternalReplyMessage string `json:"externalReplyMessage"`

	// The automatic reply to send to the audience internal to the signed-in user's organization.
	InternalReplyMessage string `json:"internalReplyMessage"`

	// The date and time that automatic replies are set to end.
	ScheduledEndDateTime DateTimeTimeZone `json:"scheduledEndDateTime"`

	// The date and time that automatic replies are set to begin.
	ScheduledStartDateTime DateTimeTimeZone `json:"scheduledStartDateTime"`

	// Configuration status for automatic replies.
	Status enums.AutoReplyStatus `json:"status"`
}
