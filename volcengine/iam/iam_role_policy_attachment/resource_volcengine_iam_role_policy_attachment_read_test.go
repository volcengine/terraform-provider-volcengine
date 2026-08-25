package iam_role_policy_attachment

import (
	"errors"
	"fmt"
	"testing"

	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

func TestRolePolicyAttachmentParentMissing(t *testing.T) {
	missing := volcengineerr.NewRequestFailure(
		volcengineerr.New("RoleNotExist", "role is absent", nil), 400, "request-id",
	)
	if !rolePolicyAttachmentParentMissing(fmt.Errorf("wrapped: %w", missing)) {
		t.Fatal("RoleNotExist was not classified as a missing attachment parent")
	}
	for _, err := range []error{
		errors.New("RoleNotExist: untyped text"),
		volcengineerr.NewRequestFailure(volcengineerr.New("AccessDenied", "denied", nil), 400, "request-id"),
	} {
		if rolePolicyAttachmentParentMissing(err) {
			t.Fatalf("unrelated error was classified as a missing attachment parent: %v", err)
		}
	}
}

func TestRolePolicyAttachmentAlreadyExists(t *testing.T) {
	conflict := volcengineerr.NewRequestFailure(
		volcengineerr.New("PolicyAttachConflict", "already attached", nil), 409, "request-id",
	)
	if !rolePolicyAttachmentAlreadyExists(fmt.Errorf("wrapped: %w", conflict)) {
		t.Fatal("PolicyAttachConflict was not classified as an existing attachment")
	}
	for _, err := range []error{
		errors.New("PolicyAttachConflict: untyped text"),
		volcengineerr.NewRequestFailure(volcengineerr.New("AccessDenied", "denied", nil), 403, "request-id"),
	} {
		if rolePolicyAttachmentAlreadyExists(err) {
			t.Fatalf("unrelated error was classified as an existing attachment: %v", err)
		}
	}
}
