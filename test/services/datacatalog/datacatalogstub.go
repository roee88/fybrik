// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"fybrik.io/fybrik/manager/controllers/mockup"
	"fybrik.io/fybrik/pkg/model/catalog"
	"github.com/gin-gonic/gin"
)

const (
	PORT = 8080
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.POST("/getAssetInfo", func(c *gin.Context) {
		creds := ""
		if values := c.Request.Header["X-Request-Datacatalog-Cred"]; len(values) > 0 {
			creds = values[0]
		}
		log.Println("creds extracted from POST request in mockup data catalog:", creds)

		input, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		log.Println("input extracted from POST request body in mockup data catalog:", string(input))
		var dataCatalogReq catalog.GetAssetRequest

		if err := json.Unmarshal(input, &dataCatalogReq); err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		dataCatalog := mockup.NewTestCatalog()
		dataCatalogResp, err := dataCatalog.GetAssetInfo(&dataCatalogReq, creds)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, dataCatalogResp)
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Serving REST APIs as part of data catalog stub")
	})

	log.Fatal(router.Run(":" + strconv.Itoa(PORT)))
}
