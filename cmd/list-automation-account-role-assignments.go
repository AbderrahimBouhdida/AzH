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
	"os"
	"os/signal"
	"path"
	"sync"
	"time"

	"github.com/AbderrahimBouhdida/AzH/client"
	"github.com/AbderrahimBouhdida/AzH/enums"
	"github.com/AbderrahimBouhdida/AzH/models"
	"github.com/AbderrahimBouhdida/AzH/pipeline"
	"github.com/spf13/cobra"
)

func init() {
	listRootCmd.AddCommand(listAutomationAccountRoleAssignment)
}

var listAutomationAccountRoleAssignment = &cobra.Command{
	Use:          "automation-account-role-assignments",
	Long:         "Lists Azure Automation Account Role Assignments",
	Run:          listAutomationAccountRoleAssignmentImpl,
	SilenceUsage: true,
}

func listAutomationAccountRoleAssignmentImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	azClient := connectAndCreateClient()
	log.Info("collecting azure automation account role assignments...")
	start := time.Now()
	subscriptions := listSubscriptions(ctx, azClient)
	stream := listAutomationAccountRoleAssignments(ctx, azClient, listAutomationAccounts(ctx, azClient, subscriptions))
	outputStream(ctx, stream)
	duration := time.Since(start)
	log.Info("collection completed", "duration", duration.String())
}

func listAutomationAccountRoleAssignments(ctx context.Context, client client.AzureClient, automationAccounts <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), automationAccounts) {
			if automationAccount, ok := result.(AzureWrapper).Data.(models.AutomationAccount); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating automation account role assignments", "result", result)
				return
			} else {
				ids <- automationAccount.Id
			}
		}
	}()

	wg.Add(len(streams))
	for i := range streams {
		stream := streams[i]
		go func() {
			defer wg.Done()
			for id := range stream {
				var (
					automationAccountRoleAssignments = models.AzureRoleAssignments{
						ObjectId: id,
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id, "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing role assignments for this automation account", "automationAccountId", id)
					} else {
						roleDefinitionId := path.Base(item.Ok.Properties.RoleDefinitionId)

						automationAccountRoleAssignment := models.AzureRoleAssignment{
							Assignee:         item.Ok,
							ObjectId:         item.ParentId,
							RoleDefinitionId: roleDefinitionId,
						}
						log.V(2).Info("found automation account role assignment", "automationAccountRoleAssignment", automationAccountRoleAssignment)
						count++
						automationAccountRoleAssignments.RoleAssignments = append(automationAccountRoleAssignments.RoleAssignments, automationAccountRoleAssignment)
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZAutomationAccountRoleAssignment,
					Data: automationAccountRoleAssignments,
				}
				log.V(1).Info("finished listing automation account role assignments", "automationAccountId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all automation account role assignments")
	}()

	return out
}
