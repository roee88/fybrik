// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"fybrik.io/crdoc/pkg/builder"
	"fybrik.io/fybrik/pkg/slices"
	taxonomyio "fybrik.io/fybrik/pkg/taxonomy/io"
	"fybrik.io/fybrik/pkg/taxonomy/model"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
)

var (
	matcher = regexp.MustCompile(`(taxonomy.json#/definitions/[a-zA-Z0-9]+)`)
)

func GenerateValidationObjectFromCRDs(inputDirOrFile, outputDir string) error {
	crds, err := builder.LoadCRDs(inputDirOrFile)
	if err != nil {
		return err
	}

	// create output directory if needed
	err = os.MkdirAll(filepath.Clean(outputDir), os.ModePerm)
	if err != nil {
		return err
	}

	for _, crd := range crds {
		outputFilepath := filepath.Join(outputDir, fmt.Sprintf("%s_%s.json", crd.Spec.Group, crd.Spec.Names.Plural))
		err = generateFile(crd, outputFilepath)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateFile(crd *apiextensions.CustomResourceDefinition, outputFilepath string) error {
	document, err := processCRD(crd)
	if err != nil {
		return err
	}

	if document != nil {
		err = taxonomyio.WriteDocumentToFile(document, outputFilepath)
		if err != nil {
			return err
		}
	}

	return nil
}

func getOpenAPIV3Schema(crd *apiextensions.CustomResourceDefinition) (*apiextensions.JSONSchemaProps, error) {
	for _, version := range crd.Spec.Versions {
		if !version.Storage {
			continue
		}

		// Find schema
		validation := version.Schema
		if validation == nil {
			// Fallback to resource level schema
			validation = crd.Spec.Validation
		}

		if validation == nil {
			return nil, errors.New("missing validation field in input CRD")
		}
		schema := validation.OpenAPIV3Schema
		return schema, nil
	}
	return nil, errors.New("missing storage version in CRD")
}

func processCRD(crd *apiextensions.CustomResourceDefinition) (*model.Document, error) {
	schema, err := getOpenAPIV3Schema(crd)
	if err != nil {
		return nil, err
	}

	jsonSchema := buildJSONSchema(schema)
	if jsonSchema == nil {
		return nil, nil
	}

	document := &model.Document{
		SchemaVersion: "http://json-schema.org/draft-04/schema#",
		Schema:        jsonSchema.Schema,
	}
	return document, nil
}

func buildJSONSchema(props *apiextensions.JSONSchemaProps) *model.SchemaRef {
	if props == nil {
		return nil
	}

	if groups := matcher.FindStringSubmatch(props.Description); len(groups) > 1 {
		return &model.SchemaRef{
			Ref: groups[1],
		}
	}

	out := &model.SchemaRef{
		Schema: model.Schema{
			Type:     props.Type,
			Required: []string{},
		},
	}

	if props.Properties != nil {
		for k, v := range props.Properties {
			schema := buildJSONSchema(&v)
			if schema != nil {
				if out.Properties == nil {
					out.Properties = model.Schemas{}
				}
				out.Properties[k] = schema
				if slices.ContainsString(k, props.Required) {
					out.Required = append(out.Required, k)
				}
			}
		}
	}

	if props.Items != nil {
		schema := buildJSONSchema(props.Items.Schema)
		if schema != nil {
			out.Items = schema
		}
	}

	if props.AdditionalProperties != nil {
		schema := buildJSONSchema(props.AdditionalProperties.Schema)
		if schema != nil {
			out.AdditionalProperties = &model.AdditionalPropertiesType{
				Schema: schema,
			}
		}
	}

	if out.Properties != nil && out.AdditionalProperties == nil {
		trueValue := true
		out.AdditionalProperties = &model.AdditionalPropertiesType{
			Allowed: &trueValue,
		}
	}

	if out.Properties != nil || out.AdditionalProperties != nil || out.Items != nil {
		return out
	}
	return nil
}
