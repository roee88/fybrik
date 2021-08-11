// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

/*
 * Policy Manager Service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPoliciesDecisions - getPoliciesDecisions
func GetPoliciesDecisions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}