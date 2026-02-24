package core

import "testing"

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
