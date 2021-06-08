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

package notebooks_test

import (
	"context"

	notebooks "cloud.google.com/go/notebooks/apiv1beta1"
	"google.golang.org/api/iterator"
	notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"
)

func ExampleNewNotebookClient() {
	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleNotebookClient_ListInstances() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.ListInstancesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListInstances(ctx, req)
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

func ExampleNotebookClient_GetInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.GetInstanceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_CreateInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.CreateInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_RegisterInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.RegisterInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RegisterInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_SetInstanceAccelerator() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.SetInstanceAcceleratorRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.SetInstanceAccelerator(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_SetInstanceMachineType() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.SetInstanceMachineTypeRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.SetInstanceMachineType(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_SetInstanceLabels() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.SetInstanceLabelsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.SetInstanceLabels(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_DeleteInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.DeleteInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleNotebookClient_StartInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.StartInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.StartInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_StopInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.StopInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.StopInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_ResetInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.ResetInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ResetInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_ReportInstanceInfo() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.ReportInstanceInfoRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ReportInstanceInfo(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_IsInstanceUpgradeable() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.IsInstanceUpgradeableRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.IsInstanceUpgradeable(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_UpgradeInstance() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.UpgradeInstanceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpgradeInstance(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_UpgradeInstanceInternal() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.UpgradeInstanceInternalRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpgradeInstanceInternal(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_ListEnvironments() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.ListEnvironmentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListEnvironments(ctx, req)
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

func ExampleNotebookClient_GetEnvironment() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.GetEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_CreateEnvironment() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.CreateEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleNotebookClient_DeleteEnvironment() {
	// import notebookspb "google.golang.org/genproto/googleapis/cloud/notebooks/v1beta1"

	ctx := context.Background()
	c, err := notebooks.NewNotebookClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &notebookspb.DeleteEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
