/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	"fmt"
	"sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
  map:
    fields:
    - name: lastTransitionTime
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: message
      type:
        scalar: string
      default: ""
    - name: observedGeneration
      type:
        scalar: numeric
    - name: reason
      type:
        scalar: string
      default: ""
    - name: status
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
          elementRelationship: atomic
    - name: matchLabels
      type:
        map:
          elementType:
            scalar: string
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: io.k8s.sigs.gateway-api.apis.v1.AllowedRoutes
  map:
    fields:
    - name: kinds
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.RouteGroupKind
          elementRelationship: atomic
    - name: namespaces
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.RouteNamespaces
- name: io.k8s.sigs.gateway-api.apis.v1.BackendObjectReference
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
- name: io.k8s.sigs.gateway-api.apis.v1.BackendRef
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
    - name: weight
      type:
        scalar: numeric
- name: io.k8s.sigs.gateway-api.apis.v1.FrontendTLSValidation
  map:
    fields:
    - name: caCertificateRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.ObjectReference
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCBackendRef
  map:
    fields:
    - name: filters
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteFilter
          elementRelationship: atomic
    - name: group
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
    - name: weight
      type:
        scalar: numeric
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCHeaderMatch
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCMethodMatch
  map:
    fields:
    - name: method
      type:
        scalar: string
    - name: service
      type:
        scalar: string
    - name: type
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCRoute
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteFilter
  map:
    fields:
    - name: extensionRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.LocalObjectReference
    - name: requestHeaderModifier
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPHeaderFilter
    - name: requestMirror
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRequestMirrorFilter
    - name: responseHeaderModifier
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPHeaderFilter
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteMatch
  map:
    fields:
    - name: headers
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCHeaderMatch
          elementRelationship: associative
          keys:
          - name
    - name: method
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCMethodMatch
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteRule
  map:
    fields:
    - name: backendRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCBackendRef
          elementRelationship: atomic
    - name: filters
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteFilter
          elementRelationship: atomic
    - name: matches
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteMatch
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteSpec
  map:
    fields:
    - name: hostnames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: parentRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.ParentReference
          elementRelationship: atomic
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteRule
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteStatus
  map:
    fields:
    - name: parents
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.RouteParentStatus
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1.Gateway
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewaySpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayAddress
  map:
    fields:
    - name: type
      type:
        scalar: string
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayClassSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayClassStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayClassSpec
  map:
    fields:
    - name: controllerName
      type:
        scalar: string
      default: ""
    - name: description
      type:
        scalar: string
    - name: parametersRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.ParametersReference
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayClassStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: supportedFeatures
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayInfrastructure
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: parametersRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.LocalParametersReference
- name: io.k8s.sigs.gateway-api.apis.v1.GatewaySpec
  map:
    fields:
    - name: addresses
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayAddress
          elementRelationship: atomic
    - name: gatewayClassName
      type:
        scalar: string
      default: ""
    - name: infrastructure
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayInfrastructure
    - name: listeners
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.Listener
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayStatus
  map:
    fields:
    - name: addresses
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayStatusAddress
          elementRelationship: atomic
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: listeners
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.ListenerStatus
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayStatusAddress
  map:
    fields:
    - name: type
      type:
        scalar: string
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.GatewayTLSConfig
  map:
    fields:
    - name: certificateRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.SecretObjectReference
          elementRelationship: atomic
    - name: frontendValidation
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.FrontendTLSValidation
    - name: mode
      type:
        scalar: string
    - name: options
      type:
        map:
          elementType:
            scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPBackendRef
  map:
    fields:
    - name: filters
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteFilter
          elementRelationship: atomic
    - name: group
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
    - name: weight
      type:
        scalar: numeric
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPHeader
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPHeaderFilter
  map:
    fields:
    - name: add
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPHeader
          elementRelationship: associative
          keys:
          - name
    - name: remove
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: set
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPHeader
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPHeaderMatch
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPPathMatch
  map:
    fields:
    - name: type
      type:
        scalar: string
    - name: value
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPPathModifier
  map:
    fields:
    - name: replaceFullPath
      type:
        scalar: string
    - name: replacePrefixMatch
      type:
        scalar: string
    - name: type
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPQueryParamMatch
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: type
      type:
        scalar: string
    - name: value
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRequestMirrorFilter
  map:
    fields:
    - name: backendRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.BackendObjectReference
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRequestRedirectFilter
  map:
    fields:
    - name: hostname
      type:
        scalar: string
    - name: path
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPPathModifier
    - name: port
      type:
        scalar: numeric
    - name: scheme
      type:
        scalar: string
    - name: statusCode
      type:
        scalar: numeric
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRoute
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteFilter
  map:
    fields:
    - name: extensionRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.LocalObjectReference
    - name: requestHeaderModifier
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPHeaderFilter
    - name: requestMirror
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRequestMirrorFilter
    - name: requestRedirect
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRequestRedirectFilter
    - name: responseHeaderModifier
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPHeaderFilter
    - name: type
      type:
        scalar: string
      default: ""
    - name: urlRewrite
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPURLRewriteFilter
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteMatch
  map:
    fields:
    - name: headers
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPHeaderMatch
          elementRelationship: associative
          keys:
          - name
    - name: method
      type:
        scalar: string
    - name: path
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPPathMatch
    - name: queryParams
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPQueryParamMatch
          elementRelationship: associative
          keys:
          - name
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteRule
  map:
    fields:
    - name: backendRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPBackendRef
          elementRelationship: atomic
    - name: filters
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteFilter
          elementRelationship: atomic
    - name: matches
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteMatch
          elementRelationship: atomic
    - name: timeouts
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteTimeouts
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteSpec
  map:
    fields:
    - name: hostnames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: parentRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.ParentReference
          elementRelationship: atomic
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteRule
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteStatus
  map:
    fields:
    - name: parents
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.RouteParentStatus
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteTimeouts
  map:
    fields:
    - name: backendRequest
      type:
        scalar: string
    - name: request
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1.HTTPURLRewriteFilter
  map:
    fields:
    - name: hostname
      type:
        scalar: string
    - name: path
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPPathModifier
- name: io.k8s.sigs.gateway-api.apis.v1.Listener
  map:
    fields:
    - name: allowedRoutes
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.AllowedRoutes
    - name: hostname
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: port
      type:
        scalar: numeric
      default: 0
    - name: protocol
      type:
        scalar: string
      default: ""
    - name: tls
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayTLSConfig
- name: io.k8s.sigs.gateway-api.apis.v1.ListenerStatus
  map:
    fields:
    - name: attachedRoutes
      type:
        scalar: numeric
      default: 0
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: name
      type:
        scalar: string
      default: ""
    - name: supportedKinds
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.RouteGroupKind
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1.LocalObjectReference
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.LocalParametersReference
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.ObjectReference
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1.ParametersReference
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1.ParentReference
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: port
      type:
        scalar: numeric
    - name: sectionName
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1.RouteGroupKind
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1.RouteNamespaces
  map:
    fields:
    - name: from
      type:
        scalar: string
    - name: selector
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
- name: io.k8s.sigs.gateway-api.apis.v1.RouteParentStatus
  map:
    fields:
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: controllerName
      type:
        scalar: string
      default: ""
    - name: parentRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.ParentReference
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1.SecretObjectReference
  map:
    fields:
    - name: group
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.BackendTLSPolicy
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.BackendTLSPolicySpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.PolicyStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.BackendTLSPolicyConfig
  map:
    fields:
    - name: caCertRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.LocalObjectReference
          elementRelationship: atomic
    - name: hostname
      type:
        scalar: string
      default: ""
    - name: wellKnownCACerts
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.BackendTLSPolicySpec
  map:
    fields:
    - name: targetRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.PolicyTargetReferenceWithSectionName
      default: {}
    - name: tls
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.BackendTLSPolicyConfig
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.GRPCRoute
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GRPCRouteStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.PolicyAncestorStatus
  map:
    fields:
    - name: ancestorRef
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.ParentReference
      default: {}
    - name: conditions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Condition
          elementRelationship: associative
          keys:
          - type
    - name: controllerName
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.PolicyStatus
  map:
    fields:
    - name: ancestors
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.PolicyAncestorStatus
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.PolicyTargetReferenceWithSectionName
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
    - name: sectionName
      type:
        scalar: string
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.ReferenceGrant
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrantSpec
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TCPRoute
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.TCPRouteSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.TCPRouteStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TCPRouteRule
  map:
    fields:
    - name: backendRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.BackendRef
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TCPRouteSpec
  map:
    fields:
    - name: parentRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.ParentReference
          elementRelationship: atomic
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.TCPRouteRule
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TCPRouteStatus
  map:
    fields:
    - name: parents
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.RouteParentStatus
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TLSRoute
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.TLSRouteSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.TLSRouteStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TLSRouteRule
  map:
    fields:
    - name: backendRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.BackendRef
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TLSRouteSpec
  map:
    fields:
    - name: hostnames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: parentRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.ParentReference
          elementRelationship: atomic
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.TLSRouteRule
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.TLSRouteStatus
  map:
    fields:
    - name: parents
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.RouteParentStatus
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.UDPRoute
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.UDPRouteSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.UDPRouteStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.UDPRouteRule
  map:
    fields:
    - name: backendRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.BackendRef
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.UDPRouteSpec
  map:
    fields:
    - name: parentRefs
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.ParentReference
          elementRelationship: atomic
    - name: rules
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1alpha2.UDPRouteRule
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1alpha2.UDPRouteStatus
  map:
    fields:
    - name: parents
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1.RouteParentStatus
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1beta1.Gateway
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewaySpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1beta1.GatewayClass
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayClassSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.GatewayClassStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1beta1.HTTPRoute
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteSpec
      default: {}
    - name: status
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1.HTTPRouteStatus
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrant
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrantSpec
      default: {}
- name: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrantFrom
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
- name: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrantSpec
  map:
    fields:
    - name: from
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrantFrom
          elementRelationship: atomic
    - name: to
      type:
        list:
          elementType:
            namedType: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrantTo
          elementRelationship: atomic
- name: io.k8s.sigs.gateway-api.apis.v1beta1.ReferenceGrantTo
  map:
    fields:
    - name: group
      type:
        scalar: string
      default: ""
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
