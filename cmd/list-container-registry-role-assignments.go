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

	"github.com/spf13/cobra"
	"https://github.com/AbderrahimBouhdida/AzH/client"
	"https://github.com/AbderrahimBouhdida/AzH/enums"
	"https://github.com/AbderrahimBouhdida/AzH/models"
	"https://github.com/AbderrahimBouhdida/AzH/pipeline"
)

func init() {
	listRootCmd.AddCommand(listContainerRegistryRoleAssignment)
}

var listContainerRegistryRoleAssignment = &cobra.Command{
	Use:          "container-registry-role-assignments",
	Long:         "Lists Azure Container Registry Role Assignments",
	Run:          listContainerRegistryRoleAssignmentImpl,
	SilenceUsage: true,
}

func listContainerRegistryRoleAssignmentImpl(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, os.Kill)
	defer gracefulShutdown(stop)

	log.V(1).Info("testing connections")
	if err := testConnections(); err != nil {
		exit(err)
	} else if azClient, err := newAzureClient(); err != nil {
		exit(err)
	} else {
		log.Info("collecting azure container registry role assignments...")
		start := time.Now()
		subscriptions := listSubscriptions(ctx, azClient)
		stream := listContainerRegistryRoleAssignments(ctx, azClient, listContainerRegistries(ctx, azClient, subscriptions))
		outputStream(ctx, stream)
		duration := time.Since(start)
		log.Info("collection completed", "duration", duration.String())
	}
}

func listContainerRegistryRoleAssignments(ctx context.Context, client client.AzureClient, containerRegistries <-chan interface{}) <-chan interface{} {
	var (
		out     = make(chan interface{})
		ids     = make(chan string)
		streams = pipeline.Demux(ctx.Done(), ids, 25)
		wg      sync.WaitGroup
	)

	go func() {
		defer close(ids)

		for result := range pipeline.OrDone(ctx.Done(), containerRegistries) {
			if containerRegistry, ok := result.(AzureWrapper).Data.(models.ContainerRegistry); !ok {
				log.Error(fmt.Errorf("failed type assertion"), "unable to continue enumerating container registry role assignments", "result", result)
				return
			} else {
				ids <- containerRegistry.Id
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
					containerRegistryRoleAssignments = models.AzureRoleAssignments{
						ObjectId: id,
					}
					count = 0
				)
				for item := range client.ListRoleAssignmentsForResource(ctx, id, "") {
					if item.Error != nil {
						log.Error(item.Error, "unable to continue processing role assignments for this container registry", "containerRegistryId", id)
					} else {
						roleDefinitionId := path.Base(item.Ok.Properties.RoleDefinitionId)

						containerRegistryRoleAssignment := models.AzureRoleAssignment{
							Assignee:         item.Ok,
							ObjectId:         item.ParentId,
							RoleDefinitionId: roleDefinitionId,
						}
						log.V(2).Info("found container registry role assignment", "containerRegistryRoleAssignment", containerRegistryRoleAssignment)
						count++
						containerRegistryRoleAssignments.RoleAssignments = append(containerRegistryRoleAssignments.RoleAssignments, containerRegistryRoleAssignment)
					}
				}
				out <- AzureWrapper{
					Kind: enums.KindAZContainerRegistryRoleAssignment,
					Data: containerRegistryRoleAssignments,
				}
				log.V(1).Info("finished listing container registry role assignments", "containerRegistryId", id, "count", count)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		log.Info("finished listing all container registry role assignments")
	}()

	return out
}
