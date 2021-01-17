// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0
package connector

import (
	"encoding/json"
	"fmt"
	"strings"
)

func emptyIfNil(val *string) string {
	if val == nil {
		return ""
	}
	return *val
}

func emptyArrayIfNil(val *[]string) []string {
	if val == nil {
		return []string{}
	}
	return *val
}

func decodeToStruct(m interface{}, s interface{}) error {
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, s)
	if err != nil {
		return err
	}
	return nil
}

func splitNamespacedName(value string) (namespace string, name string, err error) {
	identifier := strings.SplitN(value, "/", 2)
	if len(identifier) != 2 {
		err = fmt.Errorf("Expected <namespace>/<name> format but got %s", value)
		return
	}
	namespace, name = identifier[0], identifier[1]
	return
}
