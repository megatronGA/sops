package azurestackapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/azurestack/mgmt/2017-06-01/azurestack"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result azurestack.OperationListPage, err error)
}

var _ OperationsClientAPI = (*azurestack.OperationsClient)(nil)

// ProductsClientAPI contains the set of methods on the ProductsClient type.
type ProductsClientAPI interface {
	Get(ctx context.Context, resourceGroup string, registrationName string, productName string) (result azurestack.Product, err error)
	List(ctx context.Context, resourceGroup string, registrationName string) (result azurestack.ProductListPage, err error)
	ListDetails(ctx context.Context, resourceGroup string, registrationName string, productName string) (result azurestack.ExtendedProduct, err error)
}

var _ ProductsClientAPI = (*azurestack.ProductsClient)(nil)

// RegistrationsClientAPI contains the set of methods on the RegistrationsClient type.
type RegistrationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroup string, registrationName string, tokenParameter azurestack.RegistrationParameter) (result azurestack.Registration, err error)
	Delete(ctx context.Context, resourceGroup string, registrationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroup string, registrationName string) (result azurestack.Registration, err error)
	GetActivationKey(ctx context.Context, resourceGroup string, registrationName string) (result azurestack.ActivationKeyResult, err error)
	List(ctx context.Context, resourceGroup string) (result azurestack.RegistrationListPage, err error)
	Update(ctx context.Context, resourceGroup string, registrationName string, tokenParameter azurestack.RegistrationParameter) (result azurestack.Registration, err error)
}

var _ RegistrationsClientAPI = (*azurestack.RegistrationsClient)(nil)

// CustomerSubscriptionsClientAPI contains the set of methods on the CustomerSubscriptionsClient type.
type CustomerSubscriptionsClientAPI interface {
	Create(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string, customerCreationParameters azurestack.CustomerSubscription) (result azurestack.CustomerSubscription, err error)
	Delete(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroup string, registrationName string, customerSubscriptionName string) (result azurestack.CustomerSubscription, err error)
	List(ctx context.Context, resourceGroup string, registrationName string) (result azurestack.CustomerSubscriptionListPage, err error)
}

var _ CustomerSubscriptionsClientAPI = (*azurestack.CustomerSubscriptionsClient)(nil)
