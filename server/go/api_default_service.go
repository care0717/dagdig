/*
 * Workflow tool
 *
 * PFN 2019 Internship Challenge
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API. 
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	cache map[string][]Job
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService(cache map[string][]Job) DefaultApiServicer {
	return &DefaultApiService{cache}
}

// JobsGet - 
func (s *DefaultApiService) JobsGet(ctx context.Context, created string) (ImplResponse, error) {
	if created == "" {
		var allJobs []Job
		for _, jobs := range s.cache {
			allJobs = append(allJobs, jobs...)
		}
		return Response(http.StatusOK, allJobs), nil
	}
	if targetJobs, ok := s.cache[created]; ok {
		return Response(http.StatusOK, targetJobs), nil
	} else {
		return Response(http.StatusOK, []Job{}), nil
	}
}

