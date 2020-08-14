package types

// Template variable represents a template variable that might exist on a dashboard
type TemplateVariable struct {
	Name    *string `json:"name,omitempty" hcl:"name"`
	Prefix  *string `json:"prefix,omitempty" hcl:"prefix"`
	Default *string `json:"default,omitempty" hcl:"default"`
}

// Template variable preset represents a set of template variable values on a dashboard
// Not available to timeboards and screenboards
type TemplateVariablePreset struct {
	Name              *string                       `json:"name,omitempty" hcl:"name"`
	TemplateVariables []TemplateVariablePresetValue `json:"template_variables" hcl:"template_variable,block"`
}

// Template variable preset value represents the value for "name" template variable to assume
type TemplateVariablePresetValue struct {
	Name  *string `json:"name,omitempty" hcl:"name"`
	Value *string `json:"value,omitempty" hcl:"value"`
}

// Board represents a user created dashboard. This is the full dashboard
// struct when we load a dashboard in detail.
type Board struct {
	Title                   *string                  `json:"title" hcl:"title"`
	Widgets                 []BoardWidget            `json:"widgets" hcl:"widget,block"`
	LayoutType              *string                  `json:"layout_type" hcl:"layout_type"`
	Id                      *string                  `json:"id,omitempty" hcle:"omit"`
	Description             *string                  `json:"description,omitempty" hcl:"description"`
	TemplateVariables       []TemplateVariable       `json:"template_variables,omitempty" hcl:"template_variable,block"`
	TemplateVariablePresets []TemplateVariablePreset `json:"template_variable_presets,omitempty" hcl:"template_variable_preset,block"`
	IsReadOnly              *bool                    `json:"is_read_only,omitempty" hcl:"is_read_only"`
	NotifyList              []string                 `json:"notify_list,omitempty" hcl:"notify_list" hcle:"omitempty"`
	AuthorHandle            *string                  `json:"author_handle,omitempty" hcle:"omit"`
	Url                     *string                  `json:"url,omitempty" hcle:"omit"`
	CreatedAt               *string                  `json:"created_at,omitempty" hcle:"omit"`
	ModifiedAt              *string                  `json:"modified_at,omitempty" hcle:"omit"`
}
