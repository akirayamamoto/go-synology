package core

import (
	"strings"
	"testing"

	"github.com/synology-community/go-synology/pkg/query"
)

func TestNotificationTemplate_EffectiveName(t *testing.T) {
	tpl := NotificationTemplate{
		Name:         "fallback-name",
		TemplateName: "primary-name",
	}

	if got := tpl.EffectiveName(); got != "primary-name" {
		t.Fatalf("expected template_name to win, got %q", got)
	}

	tpl.TemplateName = ""
	if got := tpl.EffectiveName(); got != "fallback-name" {
		t.Fatalf("expected name fallback, got %q", got)
	}
}

func TestNotificationTemplate_DefaultTemplate(t *testing.T) {
	tests := []struct {
		name string
		tpl  NotificationTemplate
		want bool
	}{
		{
			name: "default from top-level flag",
			tpl: NotificationTemplate{
				IsDefault: false,
				Config: NotificationTemplateConfig{
					IsDefault: true,
				},
			},
			want: true,
		},
		{
			name: "default from config flag",
			tpl: NotificationTemplate{
				IsDefault: true,
			},
			want: true,
		},
		{
			name: "custom template",
			tpl: NotificationTemplate{
				IsDefault: false,
				Config: NotificationTemplateConfig{
					IsDefault: false,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tpl.DefaultTemplate(); got != tt.want {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestNotificationTemplateSetRequest_EncodesEnabledFalseInSettingsJSON(t *testing.T) {
	request := NotificationTemplateSetRequest{
		TemplateID:   4,
		TemplateName: "rule",
		Settings: []NotificationTemplateSetting{
			{Tag: "docker_container_unexpected_exit", Enabled: false},
		},
	}

	values, err := query.Values(request)
	if err != nil {
		t.Fatalf("failed to encode query values: %v", err)
	}

	encodedSettings := values.Get("settings")
	if encodedSettings == "" {
		t.Fatalf("expected settings query value to be present")
	}

	if !strings.Contains(encodedSettings, `"enabled":false`) {
		t.Fatalf("expected settings JSON to include enabled=false, got %s", encodedSettings)
	}
}
