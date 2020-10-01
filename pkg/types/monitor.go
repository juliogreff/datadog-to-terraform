package types

import (
	"encoding/json"
	"strconv"
)

type ThresholdCount struct {
	Ok               *json.Number `json:"ok,omitempty" hcl:"ok"`
	Critical         *json.Number `json:"critical,omitempty" hcl:"critical"`
	Warning          *json.Number `json:"warning,omitempty" hcl:"warning"`
	Unknown          *json.Number `json:"unknown,omitempty" hcl:"unknown"`
	CriticalRecovery *json.Number `json:"critical_recovery,omitempty" hcl:"critical_recovery"`
	WarningRecovery  *json.Number `json:"warning_recovery,omitempty" hcl:"warning_recovery"`
}

type ThresholdWindows struct {
	RecoveryWindow *string `json:"recovery_window,omitempty" hcl:"recovery_window"`
	TriggerWindow  *string `json:"trigger_window,omitempty" hcl:"trigger_window"`
}

type NoDataTimeframe int

func (tf *NoDataTimeframe) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "false" || s == "null" {
		*tf = 0
	} else {
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}
		*tf = NoDataTimeframe(i)
	}
	return nil
}

type Options struct {
	NoDataTimeframe   NoDataTimeframe   `json:"no_data_timeframe,omitempty" hcl:"no_data_timeframe"`
	NotifyAudit       *bool             `json:"notify_audit,omitempty" hcl:"notify_audit"`
	NotifyNoData      *bool             `json:"notify_no_data,omitempty" hcl:"notify_no_data"`
	RenotifyInterval  *int              `json:"renotify_interval,omitempty" hcl:"renotify_interval"`
	NewHostDelay      *int              `json:"new_host_delay,omitempty" hcl:"new_host_delay"`
	EvaluationDelay   *int              `json:"evaluation_delay,omitempty" hcl:"evaluation_delay"`
	Silenced          map[string]int    `json:"silenced,omitempty" hcl:"silenced"`
	TimeoutH          *int              `json:"timeout_h,omitempty" hcl:"timeout_h"`
	EscalationMessage *string           `json:"escalation_message,omitempty" hcl:"escalation_message"`
	Thresholds        *ThresholdCount   `json:"thresholds,omitempty" hcl:"thresholds,attribute"`
	ThresholdWindows  *ThresholdWindows `json:"threshold_windows,omitempty" hcl:"threshold_windows"`
	IncludeTags       *bool             `json:"include_tags,omitempty" hcl:"include_tags"`
	RequireFullWindow *bool             `json:"require_full_window,omitempty" hcl:"require_full_window"`
	Locked            *bool             `json:"locked,omitempty" hcl:"locked"`
	EnableLogsSample  *bool             `json:"enable_logs_sample,omitempty" hcl:"enable_logs_sample"`
}

// Monitor allows watching a metric or check that you care about,
// notifying your team when some defined threshold is exceeded
type Monitor struct {
	Type     *string  `json:"type,omitempty" hcl:"type"`
	Query    *string  `json:"query,omitempty" hcl:"query"`
	Name     *string  `json:"name,omitempty" hcl:"name"`
	Message  *string  `json:"message,omitempty" hcl:"message"`
	Tags     []string `json:"tags" hcl:"tags"`
	*Options `json:"options,omitempty" hcl:",squash"`
}
