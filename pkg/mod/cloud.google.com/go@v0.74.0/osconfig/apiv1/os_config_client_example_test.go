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

package osconfig_test

import (
	"context"

	osconfig "cloud.google.com/go/osconfig/apiv1"
	"google.golang.org/api/iterator"
	osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_ExecutePatchJob() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.ExecutePatchJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ExecutePatchJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetPatchJob() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.GetPatchJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetPatchJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CancelPatchJob() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.CancelPatchJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CancelPatchJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListPatchJobs() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.ListPatchJobsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListPatchJobs(ctx, req)
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

func ExampleClient_ListPatchJobInstanceDetails() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.ListPatchJobInstanceDetailsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListPatchJobInstanceDetails(ctx, req)
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

func ExampleClient_CreatePatchDeployment() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.CreatePatchDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreatePatchDeployment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetPatchDeployment() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.GetPatchDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetPatchDeployment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListPatchDeployments() {
	// import osconfigpb "google.golang.org/genproto/googleapis/cloud/osconfig/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.ListPatchDeploymentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListPatchDeployments(ctx, req)
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

func ExampleClient_DeletePatchDeployment() {
	ctx := context.Background()
	c, err := osconfig.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &osconfigpb.DeletePatchDeploymentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeletePatchDeployment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
