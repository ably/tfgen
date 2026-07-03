package ably_test

import (
	"reflect"
	"testing"

	"github.com/ably/tfgen/providers/ably"
)

func TestProviderVersion(t *testing.T) {
	t.Run("defaults to v1.0.0", func(t *testing.T) {
		req := ably.ProviderVersion()

		if req.Source != "ably/ably" {
			t.Errorf("source = %q, want ably/ably", req.Source)
		}
		if req.Version != "1.0.0" {
			t.Errorf("version = %q, want 1.0.0", req.Version)
		}
	})

	t.Run("accepts an explicit constraint", func(t *testing.T) {
		req := ably.ProviderVersion("~> 1.0")

		if req.Version != "~> 1.0" {
			t.Errorf("version = %q, want ~> 1.0", req.Version)
		}
	})
}

func TestProvider(t *testing.T) {
	p := ably.Provider()

	if p.Name != "ably" {
		t.Errorf("name = %q, want ably", p.Name)
	}
}

func TestApp(t *testing.T) {
	t.Run("label seeds the name and options are omitted", func(t *testing.T) {
		r := ably.App("main")

		if r.Type != "ably_app" {
			t.Errorf("type = %q, want ably_app", r.Type)
		}
		if r.Name != "main" {
			t.Errorf("label = %q, want main", r.Name)
		}
		if r.Config["name"] != "main" {
			t.Errorf("name = %v, want main (from label)", r.Config["name"])
		}
		if _, ok := r.Config["status"]; ok {
			t.Error("status should be omitted when unset")
		}
	})

	t.Run("options override defaults", func(t *testing.T) {
		r := ably.App("main", ably.AppOptions{
			Name:    "production",
			Status:  "enabled",
			TLSOnly: true,
		})

		if r.Config["name"] != "production" {
			t.Errorf("name = %v, want production", r.Config["name"])
		}
		if r.Config["tls_only"] != true {
			t.Errorf("tls_only = %v, want true", r.Config["tls_only"])
		}
	})
}

func TestAPIKey(t *testing.T) {
	caps := map[string][]string{"chat:*": {"publish", "subscribe"}}
	r := ably.APIKey("publisher", ably.APIKeyOptions{
		AppID:        "${ably_app.main.id}",
		Capabilities: caps,
	})

	if r.Type != "ably_api_key" {
		t.Errorf("type = %q, want ably_api_key", r.Type)
	}
	if r.Config["name"] != "publisher" {
		t.Errorf("name = %v, want publisher (from label)", r.Config["name"])
	}
	if r.Config["app_id"] != "${ably_app.main.id}" {
		t.Errorf("app_id = %v", r.Config["app_id"])
	}
	if !reflect.DeepEqual(r.Config["capabilities"], caps) {
		t.Errorf("capabilities = %v, want %v", r.Config["capabilities"], caps)
	}
}

func TestNamespace(t *testing.T) {
	r := ably.Namespace("chat", ably.NamespaceOptions{
		AppID:     "${ably_app.main.id}",
		Persisted: true,
	})

	if r.Type != "ably_namespace" {
		t.Errorf("type = %q, want ably_namespace", r.Type)
	}
	if r.Config["id"] != "chat" {
		t.Errorf("id = %v, want chat (from label)", r.Config["id"])
	}
	if r.Config["persisted"] != true {
		t.Errorf("persisted = %v, want true", r.Config["persisted"])
	}
}

func TestQueue(t *testing.T) {
	r := ably.Queue("events", ably.QueueOptions{
		AppID:     "${ably_app.main.id}",
		Region:    "us-east-1-a",
		MaxLength: 10000,
		TTL:       60,
	})

	if r.Type != "ably_queue" {
		t.Errorf("type = %q, want ably_queue", r.Type)
	}
	if r.Config["name"] != "events" {
		t.Errorf("name = %v, want events (from label)", r.Config["name"])
	}
	if r.Config["max_length"] != 10000 {
		t.Errorf("max_length = %v, want 10000", r.Config["max_length"])
	}
}

func TestResourceRef(t *testing.T) {
	r := ably.App("main")

	got := r.Ref("id")
	want := "${ably_app.main.id}"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
