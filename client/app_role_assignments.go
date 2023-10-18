package client

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/AbderrahimBouhdida/AzH/client/query"
	"github.com/AbderrahimBouhdida/AzH/client/rest"
	"github.com/AbderrahimBouhdida/AzH/constants"
	"github.com/AbderrahimBouhdida/AzH/models/azure"
)

func (s *azureClient) GetAzureADAppRoleAssignments(ctx context.Context, servicePrincipalId string, filter, search, orderBy, expand string, selectCols []string, top int32, count bool) (azure.AppRoleAssignmentList, error) {
	var (
		path     = fmt.Sprintf("/%s/servicePrincipals/%s/appRoleAssignedTo", constants.GraphApiVersion, servicePrincipalId)
		params   = query.Params{Filter: filter, Search: search, OrderBy: orderBy, Select: selectCols, Top: top, Count: count, Expand: expand}
		headers  map[string]string
		response azure.AppRoleAssignmentList
	)

	count = count || search != "" || (filter != "" && orderBy != "") || strings.Contains(filter, "endsWith")
	if count {
		headers = make(map[string]string)
		headers["ConsistencyLevel"] = "eventual"
	}
	if res, err := s.msgraph.Get(ctx, path, params.AsMap(), headers); err != nil {
		return response, err
	} else if err := rest.Decode(res.Body, &response); err != nil {
		return response, err
	} else {
		return response, nil
	}
}

func (s *azureClient) ListAzureADAppRoleAssignments(ctx context.Context, servicePrincipal, filter, search, orderBy, expand string, selectCols []string) <-chan azure.AppRoleAssignmentResult {
	out := make(chan azure.AppRoleAssignmentResult)

	go func() {
		defer close(out)

		var (
			errResult = azure.AppRoleAssignmentResult{}
			nextLink  string
		)

		if list, err := s.GetAzureADAppRoleAssignments(ctx, servicePrincipal, filter, search, orderBy, expand, selectCols, 999, false); err != nil {
			errResult.Error = err
			out <- errResult
		} else {
			for _, u := range list.Value {
				out <- azure.AppRoleAssignmentResult{Ok: u}
			}

			nextLink = list.NextLink
			for nextLink != "" {
				var list azure.AppRoleAssignmentList
				if url, err := url.Parse(nextLink); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if req, err := rest.NewRequest(ctx, "GET", url, nil, nil, nil); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if res, err := s.msgraph.Send(req); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else if err := rest.Decode(res.Body, &list); err != nil {
					errResult.Error = err
					out <- errResult
					nextLink = ""
				} else {
					for _, u := range list.Value {
						out <- azure.AppRoleAssignmentResult{Ok: u}
					}
					nextLink = list.NextLink
				}
			}
		}
	}()
	return out
}
