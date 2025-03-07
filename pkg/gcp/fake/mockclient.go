/*
 * Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package mock

import (
	"context"
	"fmt"
	"net/http"

	compute "google.golang.org/api/compute/v1"
	option "google.golang.org/api/option"
	corev1 "k8s.io/api/core/v1"

	api "github.com/gardener/machine-controller-manager-provider-gcp/pkg/api/v1alpha1"
)

// PluginSPIImpl is the mock implementation of PluginSPIImpl
type PluginSPIImpl struct {
	Client *http.Client
}

// NewComputeService creates a compute service instance using the mock
func (ms *PluginSPIImpl) NewComputeService(secrets *corev1.Secret) (context.Context, *compute.Service, error) {
	ctx := context.Background()

	_, serviceAccountJSON := secrets.Data[api.GCPServiceAccountJSON]
	_, serviceAccountJSONAlternative := secrets.Data[api.GCPAlternativeServiceAccountJSON]
	if !serviceAccountJSON && !serviceAccountJSONAlternative {
		return nil, nil, fmt.Errorf("Missing secrets to connect to compute service")
	}

	// create a compute service using a mockclient work
	client := option.WithHTTPClient(ms.Client)
	endpoint := option.WithEndpoint("http://127.0.0.1:6666")

	computeService, err := compute.NewService(ctx, client, endpoint)
	if err != nil {
		return nil, nil, err
	}

	return ctx, computeService, nil
}
