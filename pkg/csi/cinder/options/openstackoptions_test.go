/*
Copyright 2026 The Kubernetes Authors.

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

package options

import "testing"

func TestNewOpenstackOptions(t *testing.T) {
	t.Parallel()

	opts, err := NewOpenstackOptions(map[string]string{
		"os-authURL":      "https://keystone.example.com/v3",
		"os-userName":     "admin",
		"os-password":     "secret",
		"os-projectName":  "demo",
		"os-domainName":   "Default",
		"os-endpointType": "public",
	})
	if err != nil {
		t.Fatalf("NewOpenstackOptions returned error: %v", err)
	}

	if opts.AuthURL != "https://keystone.example.com/v3" {
		t.Fatalf("expected auth URL to be populated, got %q", opts.AuthURL)
	}
	if opts.Username != "admin" {
		t.Fatalf("expected username to be populated, got %q", opts.Username)
	}
	if opts.Password != "secret" {
		t.Fatalf("expected password to be populated, got %q", opts.Password)
	}
	if opts.TenantName != "demo" {
		t.Fatalf("expected project name to be populated, got %q", opts.TenantName)
	}
	if opts.DomainName != "Default" {
		t.Fatalf("expected domain name to be populated, got %q", opts.DomainName)
	}
	if string(opts.EndpointType) != "public" {
		t.Fatalf("expected endpoint type to be populated, got %q", opts.EndpointType)
	}
}

func TestNewOpenstackOptionsMissingDependency(t *testing.T) {
	t.Parallel()

	_, err := NewOpenstackOptions(map[string]string{
		"os-authURL":  "https://keystone.example.com/v3",
		"os-userName": "admin",
	})
	if err == nil {
		t.Fatal("expected validation error for incomplete credentials")
	}
}

func TestNewOpenstackOptionsWithCloudsYAML(t *testing.T) {
	t.Parallel()

	opts, err := NewOpenstackOptions(map[string]string{
		"os-useClouds": "true",
		"os-cloud":     "demo",
	})
	if err != nil {
		t.Fatalf("NewOpenstackOptions returned error: %v", err)
	}

	if !opts.UseClouds {
		t.Fatal("expected UseClouds to be enabled")
	}
	if opts.Cloud != "demo" {
		t.Fatalf("expected cloud name to be populated, got %q", opts.Cloud)
	}
}
