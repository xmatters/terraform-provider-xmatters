package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/xmatters/terraform-provider-xmatters/internal/resources/serviceDependency"
	"github.com/xmatters/terraform-provider-xmatters/internal/utils"
	"github.com/xmatters/xmatters-go"
)

func TestResourceServiceDependencyToModel(t *testing.T) {
	testServiceDependency := xmatters.ServiceDependency{
		ID: utils.RandUUIDPointer(),
		Service: &xmatters.ServiceReference{
			ID: utils.RandUUIDPointer(),
		},
		DependentService: &xmatters.ServiceReference{
			ID: utils.RandUUIDPointer(),
		},
	}
	tests := []struct {
		name       string
		serviceDep xmatters.ServiceDependency
		expected   serviceDependency.ServiceDependencyModel
	}{
		{
			name:       "empty service dependency",
			serviceDep: xmatters.ServiceDependency{},
			expected:   serviceDependency.ServiceDependencyModel{},
		},
		{
			name:       "valid service",
			serviceDep: testServiceDependency,
			expected: serviceDependency.ServiceDependencyModel{
				ID:               types.StringPointerValue(testServiceDependency.ID),
				Service:          types.StringPointerValue(testServiceDependency.Service.ID),
				DependentService: types.StringPointerValue(testServiceDependency.DependentService.ID),
			},
		},
	}
	for _, thisTest := range tests {
		actual := serviceDependency.ServiceDependencyToModel(thisTest.serviceDep)
		assert.Equal(t, thisTest.expected, actual)
	}
}

func TestServiceDependencyParams(t *testing.T) {
	testParams := xmatters.PushServiceDependencyParams{
		ID:                 utils.RandUUID(),
		ServiceID:          utils.RandUUID(),
		DependentServiceID: utils.RandUUID(),
	}
	tests := []struct {
		name     string
		model    serviceDependency.ServiceDependencyModel
		expected *xmatters.PushServiceDependencyParams
	}{
		{
			name: "create with service only",
			model: serviceDependency.ServiceDependencyModel{
				Service: types.StringValue(testParams.ServiceID),
			},
			expected: &xmatters.PushServiceDependencyParams{
				ServiceID: testParams.ServiceID,
			},
		},
		{
			name: "get dependent service only",
			model: serviceDependency.ServiceDependencyModel{
				DependentService: types.StringValue(testParams.DependentServiceID),
			},
			expected: &xmatters.PushServiceDependencyParams{
				DependentServiceID: testParams.DependentServiceID,
			},
		},
		{
			name: "get service and dependant service",
			model: serviceDependency.ServiceDependencyModel{
				Service:          types.StringValue(testParams.ServiceID),
				DependentService: types.StringValue(testParams.DependentServiceID),
			},
			expected: &xmatters.PushServiceDependencyParams{
				ServiceID:          testParams.ServiceID,
				DependentServiceID: testParams.DependentServiceID,
			},
		},
		{
			name: "modify service dependency",
			model: serviceDependency.ServiceDependencyModel{
				ID:               types.StringValue(testParams.ID),
				Service:          types.StringValue(testParams.ServiceID),
				DependentService: types.StringValue(testParams.DependentServiceID),
			},
			expected: &testParams,
		},
	}
	for _, thisTest := range tests {
		actual := thisTest.model.ServiceDependencyParams()
		assert.Equal(t, thisTest.expected, actual)
	}
}
