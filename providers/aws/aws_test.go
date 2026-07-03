package aws_test

import (
	"reflect"
	"testing"

	"github.com/ably/tfgen"
	"github.com/ably/tfgen/providers/aws"
)

func TestProviderVersion(t *testing.T) {
	req := aws.ProviderVersion("~> 5.0")

	if req.Source != "hashicorp/aws" {
		t.Errorf("source = %q, want hashicorp/aws", req.Source)
	}
	if req.Version != "~> 5.0" {
		t.Errorf("version = %q, want ~> 5.0", req.Version)
	}
}

func TestProvider(t *testing.T) {
	tags := tfgen.Config{"env": "prod"}
	p := aws.Provider("primary", "eu-west-1", tags)

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

func TestProviderWithAssumeRole(t *testing.T) {
	tags := tfgen.Config{"env": "prod"}
	roleARN := "arn:aws:iam::123456789012:role/Example"
	p := aws.ProviderWithAssumeRole("management", "eu-west-1", roleARN, tags)

	assumeRole, ok := p.Config["assume_role"].([]tfgen.Config)
	if !ok {
		t.Fatal("assume_role should be []tfgen.Config")
	}

	if assumeRole[0]["role_arn"] != roleARN {
		t.Errorf("role_arn = %v, want %s", assumeRole[0]["role_arn"], roleARN)
	}
}

func TestAssumeRole(t *testing.T) {
	roleARN := "arn:aws:iam::123456789012:role/Example"
	got := aws.AssumeRole(roleARN)
	want := []tfgen.Config{
		{"role_arn": roleARN},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
