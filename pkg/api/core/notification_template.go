package core

type NotificationTemplateSetting struct {
	Tag     string `json:"tag,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

type NotificationTemplateConfig struct {
	DefaultEnabledRuleLevel string `json:"default_enabled_rule_level,omitempty"`
	DefaultType             string `json:"default_type,omitempty"`
	IsDefault               bool   `json:"is_default,omitempty"`
}

type NotificationTemplate struct {
	TemplateID   int64                         `json:"template_id,omitempty"`
	Name         string                        `json:"name,omitempty"`
	TemplateName string                        `json:"template_name,omitempty"`
	IsDefault    bool                          `json:"is_default,omitempty"`
	Config       NotificationTemplateConfig    `json:"config,omitempty"`
	Settings     []NotificationTemplateSetting `json:"settings,omitempty"`
}

func (t NotificationTemplate) EffectiveName() string {
	if t.TemplateName != "" {
		return t.TemplateName
	}
	return t.Name
}

func (t NotificationTemplate) DefaultTemplate() bool {
	return t.IsDefault || t.Config.IsDefault
}

type NotificationTemplateListResponse struct {
	Templates []NotificationTemplate `json:"templates,omitempty"`
}

type NotificationTemplateGetRequest struct {
	TemplateID int64 `url:"template_id"`
}

type NotificationTemplateCreateRequest struct {
	TemplateName string                        `url:"template_name,quoted"`
	Settings     []NotificationTemplateSetting `url:"settings,json"`
}

type NotificationTemplateCreateResponse struct {
	TemplateID int64 `json:"template_id,omitempty"`
}

type NotificationTemplateSetRequest struct {
	TemplateID   int64                         `url:"template_id"`
	TemplateName string                        `url:"template_name,quoted"`
	Settings     []NotificationTemplateSetting `url:"settings,json"`
}

type NotificationTemplateDeleteRequest struct {
	TemplateID int64 `url:"template_id"`
}
