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

package client

import (
	"context"
	"fmt"
	"net/url"

	"https://github.com/AbderrahimBouhdida/AzH/client/query"
	"https://github.com/AbderrahimBouhdida/AzH/client/rest"
	"https://github.com/AbderrahimBouhdida/AzH/models/azure"
)

func (s *azureClient) GetAzureKeyVault(ctx context.Context, subscriptionId, groupName, vaultName string) (*azure.KeyVault, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.KeyVault/vaults/%s", subscriptionId, groupName, vaultName)
		params   = query.Params{ApiVersion: "2019-09-01"}.AsMap()
		headers  map[string]string
		response azure.KeyVault
	)
	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return nil, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return nil, err
	} else {
		return &response, nil
	}
}

func (s *azureClient) GetAzureKeyVaults(ctx context.Context, subscriptionId string, top int32) (azure.KeyVaultList, error) {
	var (
		path     = fmt.Sprintf("/subscriptions/%s/providers/Microsoft.KeyVault/vaults", subscriptionId)
		params   = query.Params{ApiVersion: "2019-09-01", Top: top}.AsMap()
		headers  map[string]string
		response azure.KeyVaultList
	)

	if res, err := s.resourceManager.Get(ctx, path, params, headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureKeyVaults(ctx context.Context, subscriptionId string, top int32) <-chan azure.KeyVaultResult {
	out := make(chan azure.KeyVaultResult)

	go func() {
		defer close(out)

		var (
			errResult = azure.KeyVaultResult{
				SubscriptionId: subscriptionId,
			}
			nextLink string
		)

		if result, err := s.GetAzureKeyVaults(ctx, subscriptionId, top); err != nil {
			errResult.Error = err
			out <- errResult
		} else {
			for _, u := range result.Value {
				out <- azure.KeyVaultResult{
					SubscriptionId: subscriptionId,
					Ok:             u,
				}
			}

			nextLink = result.NextLink
			for nextLink != "" {
				var list azure.KeyVaultList
				if url, err := url.Parse(nextLink); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if req, err := rest.NewRequest(ctx, "GET", url, nil, nil, nil); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if res, err := s.resourceManager.Send(req); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if err := rest.Decode(res.Body, &list); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else {
					for _, u := range list.Value {
						out <- azure.KeyVaultResult{
							SubscriptionId: subscriptionId,
							Ok:             u,
						}
					}
					nextLink = list.NextLink
				}
			}
		}
	}()
	return out
}
