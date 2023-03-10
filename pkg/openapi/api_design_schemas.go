// Copyright 2022 Cisco Systems, Inc. and its affiliates
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

/*
 * Flame REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/cisco-open/flame/pkg/openapi/constants"
)

// DesignSchemasApiController binds http requests to an api service and writes the service results to the http response
type DesignSchemasApiController struct {
	service      DesignSchemasApiServicer
	errorHandler ErrorHandler
}

// DesignSchemasApiOption for how the controller is set up.
type DesignSchemasApiOption func(*DesignSchemasApiController)

// WithDesignSchemasApiErrorHandler inject ErrorHandler into controller
func WithDesignSchemasApiErrorHandler(h ErrorHandler) DesignSchemasApiOption {
	return func(c *DesignSchemasApiController) {
		c.errorHandler = h
	}
}

// NewDesignSchemasApiController creates a default api controller
func NewDesignSchemasApiController(s DesignSchemasApiServicer, opts ...DesignSchemasApiOption) Router {
	controller := &DesignSchemasApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DesignSchemasApiController
func (c *DesignSchemasApiController) Routes() Routes {
	return Routes{
		{
			"CreateDesignSchema",
			strings.ToUpper("Post"),
			"/users/{user}/designs/{designId}/schemas",
			c.CreateDesignSchema,
		},
		{
			"DeleteDesignSchema",
			strings.ToUpper("Delete"),
			"/users/{user}/designs/{designId}/schemas/{version}",
			c.DeleteDesignSchema,
		},
		{
			"GetDesignSchema",
			strings.ToUpper("Get"),
			"/users/{user}/designs/{designId}/schemas/{version}",
			c.GetDesignSchema,
		},
		{
			"GetDesignSchemas",
			strings.ToUpper("Get"),
			"/users/{user}/designs/{designId}/schemas",
			c.GetDesignSchemas,
		},
		{
			"UpdateDesignSchema",
			strings.ToUpper("Put"),
			"/users/{user}/designs/{designId}/schemas/{version}",
			c.UpdateDesignSchema,
		},
	}
}

// CreateDesignSchema - Update a design schema
func (c *DesignSchemasApiController) CreateDesignSchema(w http.ResponseWriter, r *http.Request) {
	designSchemaParam := DesignSchema{}

	params := mux.Vars(r)
	userParam := params[constants.ParamUser]
	designIdParam := params[constants.ParamDesignID]

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	if err := d.Decode(&designSchemaParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	if err := AssertDesignSchemaRequired(designSchemaParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}

	result, err := c.service.CreateDesignSchema(r.Context(), userParam, designIdParam, designSchemaParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}

	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetDesignSchema - Get a design schema owned by user
func (c *DesignSchemasApiController) GetDesignSchema(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	designIdParam := params[constants.ParamDesignID]

	versionParam := params["version"]

	result, err := c.service.GetDesignSchema(r.Context(), userParam, designIdParam, versionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetDesignSchemas - Get all design schemas in a design
func (c *DesignSchemasApiController) GetDesignSchemas(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	designIdParam := params[constants.ParamDesignID]

	result, err := c.service.GetDesignSchemas(r.Context(), userParam, designIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateDesignSchema - Update a schema for a given design
func (c *DesignSchemasApiController) UpdateDesignSchema(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	designIdParam := params[constants.ParamDesignID]

	versionParam := params["version"]

	designSchemaParam := DesignSchema{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&designSchemaParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDesignSchemaRequired(designSchemaParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateDesignSchema(r.Context(), userParam, designIdParam, versionParam, designSchemaParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteDesignSchema - Delete a schema for a given design
func (c *DesignSchemasApiController) DeleteDesignSchema(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userParam := params[constants.ParamUser]

	designIdParam := params[constants.ParamDesignID]

	versionParam := params["version"]

	result, err := c.service.DeleteDesignSchema(r.Context(), userParam, designIdParam, versionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
