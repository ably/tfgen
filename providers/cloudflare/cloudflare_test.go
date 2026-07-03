package cloudflare_test

import (
	"testing"

	"github.com/ably/tfgen/providers/cloudflare"
)

func TestProviderVersion(t *testing.T) {
	req := cloudflare.ProviderVersion("~> 4.0")

	if req.Source != "cloudflare/cloudflare" {
		t.Errorf("source = %q, want cloudflare/cloudflare", req.Source)
	}
	if req.Version != "~> 4.0" {
		t.Errorf("version = %q, want ~> 4.0", req.Version)
	}
}

func TestProvider(t *testing.T) {
	p := cloudflare.Provider()

	if p.Name != "cloudflare" {
		t.Errorf("name = %q, want cloudflare", p.Name)
	}
}
