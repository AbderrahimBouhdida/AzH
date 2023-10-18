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

type VMPublicIPConfigProperties struct {
	// Specify what happens to the public IP address when the VM is deleted.
	DeleteOption enums.VMDeleteOption `json:"deleteOption"`

	// The dns settings to be applied on the publicIP addresses.
	DNSSettings VMPublicIPDNSSettings `json:"dnsSettings"`

	// The idle timeout of the public IP address.
	IdleTimeoutInMinutes int `json:"idleTimeoutInMinutes"`

	// The list of IP tags associated with the public IP address.
	IPTags []VMIPTag `json:"ipTags"`

	// Available from Api-Version 2019-07-01 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6.
	// Default is taken as IPv4. Possible values are: 'IPv4' and 'IPv6'.
	PublicIPAddressVersion string `json:"publicIpAddressVersion"`

	// Specify the public IP allocation type.
	PublicIPAllocationMethod enums.IPAllocationMethod `json:"publicIpAllocationMethod"`
}
