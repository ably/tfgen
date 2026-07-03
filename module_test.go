package tfgen

import (
	"testing"
)

func TestModuleSourceRegistrySource(t *testing.T) {
	ms := ModuleSource{Name: "vpc", Provider: "aws"}

	got := ms.RegistrySource("app.terraform.io/example-org")
	want := "app.terraform.io/example-org/vpc/aws"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestModuleSourceLocalSource(t *testing.T) {
	ms := ModuleSource{Name: "eks-cluster", Provider: "aws"}

	got := ms.LocalSource("/home/user/terraform/modules")
	want := "/home/user/terraform/modules/aws/terraform-aws-eks-cluster"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestModuleSourceLocalSourceWithSubDir(t *testing.T) {
	ms := ModuleSource{Name: "vpc", Provider: "aws", SubDir: "networking"}

	got := ms.LocalSource("/home/user/terraform/modules")
	want := "/home/user/terraform/modules/aws/networking/terraform-aws-vpc"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestModuleOutputRef(t *testing.T) {
	m := &Module{Name: "my_cluster"}

	got := m.OutputRef("vpc_id")
	want := "${module.my_cluster.vpc_id}"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
