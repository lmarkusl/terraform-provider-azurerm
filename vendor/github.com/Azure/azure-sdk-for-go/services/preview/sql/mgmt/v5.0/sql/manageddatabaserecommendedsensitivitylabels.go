package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ManagedDatabaseRecommendedSensitivityLabelsClient is the the Azure SQL Database management API provides a RESTful
// set of web services that interact with Azure SQL Database services to manage your databases. The API enables you to
// create, retrieve, update, and delete databases.
type ManagedDatabaseRecommendedSensitivityLabelsClient struct {
	BaseClient
}

// NewManagedDatabaseRecommendedSensitivityLabelsClient creates an instance of the
// ManagedDatabaseRecommendedSensitivityLabelsClient client.
func NewManagedDatabaseRecommendedSensitivityLabelsClient(subscriptionID string) ManagedDatabaseRecommendedSensitivityLabelsClient {
	return NewManagedDatabaseRecommendedSensitivityLabelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedDatabaseRecommendedSensitivityLabelsClientWithBaseURI creates an instance of the
// ManagedDatabaseRecommendedSensitivityLabelsClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewManagedDatabaseRecommendedSensitivityLabelsClientWithBaseURI(baseURI string, subscriptionID string) ManagedDatabaseRecommendedSensitivityLabelsClient {
	return ManagedDatabaseRecommendedSensitivityLabelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Update update recommended sensitivity labels states of a given database using an operations batch.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// managedInstanceName - the name of the managed instance.
// databaseName - the name of the database.
func (client ManagedDatabaseRecommendedSensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedDatabaseRecommendedSensitivityLabelsClient.Update")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, managedInstanceName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseRecommendedSensitivityLabelsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseRecommendedSensitivityLabelsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ManagedDatabaseRecommendedSensitivityLabelsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ManagedDatabaseRecommendedSensitivityLabelsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, parameters RecommendedSensitivityLabelUpdateList) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":        autorest.Encode("path", databaseName),
		"managedInstanceName": autorest.Encode("path", managedInstanceName),
		"resourceGroupName":   autorest.Encode("path", resourceGroupName),
		"subscriptionId":      autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/recommendedSensitivityLabels", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedDatabaseRecommendedSensitivityLabelsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ManagedDatabaseRecommendedSensitivityLabelsClient) UpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}