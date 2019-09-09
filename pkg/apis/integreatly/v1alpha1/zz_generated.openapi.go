// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/integreatly/v1alpha1.BlobStorage":       schema_pkg_apis_integreatly_v1alpha1_BlobStorage(ref),
		"./pkg/apis/integreatly/v1alpha1.BlobStorageSpec":   schema_pkg_apis_integreatly_v1alpha1_BlobStorageSpec(ref),
		"./pkg/apis/integreatly/v1alpha1.BlobStorageStatus": schema_pkg_apis_integreatly_v1alpha1_BlobStorageStatus(ref),
	}
}

func schema_pkg_apis_integreatly_v1alpha1_BlobStorage(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BlobStorage is the Schema for the blobstorages API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/integreatly/v1alpha1.BlobStorageSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/integreatly/v1alpha1.BlobStorageStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/integreatly/v1alpha1.BlobStorageSpec", "./pkg/apis/integreatly/v1alpha1.BlobStorageStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_BlobStorageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BlobStorageSpec defines the desired state of BlobStorage",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tier": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/integreatly/v1alpha1.SecretRef"),
						},
					},
				},
				Required: []string{"type", "tier", "secretRef"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/integreatly/v1alpha1.SecretRef"},
	}
}

func schema_pkg_apis_integreatly_v1alpha1_BlobStorageStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BlobStorageStatus defines the observed state of BlobStorage",
				Properties: map[string]spec.Schema{
					"strategy": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"provider": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/integreatly/v1alpha1.SecretRef"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/integreatly/v1alpha1.SecretRef"},
	}
}
