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

package cmd

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"https://github.com/AbderrahimBouhdida/AzH/client/mocks"
	"https://github.com/AbderrahimBouhdida/AzH/models/azure"
)

func init() {
	setupLogger()
}

func TestListSubscriptions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockClient := mocks.NewMockAzureClient(ctrl)
	mockChannel := make(chan azure.SubscriptionResult)
	mockTenant := azure.Tenant{}
	mockError := fmt.Errorf("I'm an error")
	mockClient.EXPECT().TenantInfo().Return(mockTenant).AnyTimes()
	mockClient.EXPECT().ListAzureSubscriptions(gomock.Any()).Return(mockChannel)

	go func() {
		defer close(mockChannel)
		mockChannel <- azure.SubscriptionResult{
			Ok: azure.Subscription{},
		}
		mockChannel <- azure.SubscriptionResult{
			Error: mockError,
		}
		mockChannel <- azure.SubscriptionResult{
			Ok: azure.Subscription{},
		}
	}()

	channel := listSubscriptions(ctx, mockClient)
	result := <-channel
	if _, ok := result.(AzureWrapper); !ok {
		t.Errorf("failed type assertion: got %T, want %T", result, AzureWrapper{})
	}

	if _, ok := <-channel; ok {
		t.Error("expected channel to close from an error result but it did not")
	}
}
