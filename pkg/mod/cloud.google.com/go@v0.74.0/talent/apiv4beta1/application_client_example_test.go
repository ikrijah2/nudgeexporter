// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package talent_test

import (
	"context"

	talent "cloud.google.com/go/talent/apiv4beta1"
	"google.golang.org/api/iterator"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
)

func ExampleNewApplicationClient() {
	ctx := context.Background()
	c, err := talent.NewApplicationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleApplicationClient_CreateApplication() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

	ctx := context.Background()
	c, err := talent.NewApplicationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.CreateApplicationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateApplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleApplicationClient_GetApplication() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

	ctx := context.Background()
	c, err := talent.NewApplicationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.GetApplicationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetApplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleApplicationClient_UpdateApplication() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

	ctx := context.Background()
	c, err := talent.NewApplicationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.UpdateApplicationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateApplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleApplicationClient_DeleteApplication() {
	ctx := context.Background()
	c, err := talent.NewApplicationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.DeleteApplicationRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteApplication(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleApplicationClient_ListApplications() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := talent.NewApplicationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.ListApplicationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListApplications(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
