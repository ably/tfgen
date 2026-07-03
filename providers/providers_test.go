package providers_test

import (
	"reflect"
	"testing"

	"github.com/ably/tfgen/providers"
)

func TestAWS(t *testing.T) {
	tags := map[string]any{"env": "prod"}
	p := providers.AWS("primary", "eu-west-1", tags)

	if p.Name != "aws" {
		t.Errorf("name = %q, want aws", p.Name)
	}

	if p.Config["alias"] != "primary" {
		t.Errorf("alias = %v, want primary", p.Config["alias"])
	}

	if p.Config["region"] != "eu-west-1" {
		t.Errorf("region = %v, want eu-west-1", p.Config["region"])
	}

	if _, ok := p.Config["assume_role"]; ok {
		t.Error("should not have assume_role")
	}
}

func TestAWSWithAssumeRole(t *testing.T) {
	tags := map[string]any{"env": "prod"}
	roleARN := "arn:aws:iam::123456789012:role/Example"
	p := providers.AWSWithAssumeRole("management", "eu-west-1", roleARN, tags)

	assumeRole, ok := p.Config["assume_role"].([]map[string]any)
	if !ok {
		t.Fatal("assume_role should be []map[string]any")
	}

	if assumeRole[0]["role_arn"] != roleARN {
		t.Errorf("role_arn = %v, want %s", assumeRole[0]["role_arn"], roleARN)
	}
}

func TestCloudflare(t *testing.T) {
	p := providers.Cloudflare()

	if p.Name != "cloudflare" {
		t.Errorf("name = %q, want cloudflare", p.Name)
	}
}

func TestAssumeRole(t *testing.T) {
	roleARN := "arn:aws:iam::123456789012:role/Example"
	got := providers.AssumeRole(roleARN)
	want := []map[string]any{
		{"role_arn": roleARN},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
