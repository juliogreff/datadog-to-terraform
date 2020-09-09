package types

import (
	"encoding/json"
	"fmt"
)

const (
	ALERT_GRAPH_WIDGET             = "alert_graph"
	ALERT_VALUE_WIDGET             = "alert_value"
	CHANGE_WIDGET                  = "change"
	CHECK_STATUS_WIDGET            = "check_status"
	DISTRIBUTION_WIDGET            = "distribution"
	EVENT_STREAM_WIDGET            = "event_stream"
	EVENT_TIMELINE_WIDGET          = "event_timeline"
	FREE_TEXT_WIDGET               = "free_text"
	GROUP_WIDGET                   = "group"
	HEATMAP_WIDGET                 = "heatmap"
	HOSTMAP_WIDGET                 = "hostmap"
	IFRAME_WIDGET                  = "iframe"
	IMAGE_WIDGET                   = "image"
	LOG_STREAM_WIDGET              = "log_stream"
	MANAGE_STATUS_WIDGET           = "manage_status"
	NOTE_WIDGET                    = "note"
	QUERY_VALUE_WIDGET             = "query_value"
	QUERY_TABLE_WIDGET             = "query_table"
	SCATTERPLOT_WIDGET             = "scatterplot"
	SERVICE_LEVEL_OBJECTIVE_WIDGET = "slo"
	TIMESERIES_WIDGET              = "timeseries"
	TOPLIST_WIDGET                 = "toplist"
	TRACE_SERVICE_WIDGET           = "trace_service"
)

// BoardWidget represents the structure of any widget. However, the widget Definition structure is
// different according to widget type.
type BoardWidget struct {
	Id     *int64        `json:"id,omitempty" hcle:"omit"`
	Layout *WidgetLayout `json:"layout,omitempty" hcl:"layout,attribute" hcle:"omitempty"`

	*AlertGraphDefinition            `hcl:"alert_graph_definition" hcle:"omitempty"`
	*AlertValueDefinition            `hcl:"alert_value_definition" hcle:"omitempty"`
	*ChangeDefinition                `hcl:"change_definition" hcle:"omitempty"`
	*CheckStatusDefinition           `hcl:"check_status_definition" hcle:"omitempty"`
	*DistributionDefinition          `hcl:"distribution_definition" hcle:"omitempty"`
	*EventStreamDefinition           `hcl:"event_stream_definition" hcle:"omitempty"`
	*EventTimelineDefinition         `hcl:"event_timeline_definition" hcle:"omitempty"`
	*FreeTextDefinition              `hcl:"free_text_definition" hcle:"omitempty"`
	*GroupDefinition                 `hcl:"group_definition" hcle:"omitempty"`
	*HeatmapDefinition               `hcl:"heatmap_definition" hcle:"omitempty"`
	*HostmapDefinition               `hcl:"hostmap_definition" hcle:"omitempty"`
	*IframeDefinition                `hcl:"iframe_definition" hcle:"omitempty"`
	*ImageDefinition                 `hcl:"image_definition" hcle:"omitempty"`
	*LogStreamDefinition             `hcl:"log_stream_definition" hcle:"omitempty"`
	*ManageStatusDefinition          `hcl:"manage_status_definition" hcle:"omitempty"`
	*NoteDefinition                  `hcl:"note_definition" hcle:"omitempty"`
	*QueryValueDefinition            `hcl:"query_value_definition" hcle:"omitempty"`
	*QueryTableDefinition            `hcl:"query_table_definition" hcle:"omitempty"`
	*ScatterplotDefinition           `hcl:"scatterplot_definition" hcle:"omitempty"`
	*ServiceLevelObjectiveDefinition `hcl:"service_level_objective_definition" hcle:"omitempty"`
	*TimeseriesDefinition            `hcl:"timeseries_definition" hcle:"omitempty"`
	*ToplistDefinition               `hcl:"toplist_definition" hcle:"omitempty"`
	*TraceServiceDefinition          `hcl:"trace_service_definition" hcle:"omitempty"`
}

// WidgetLayout represents the layout for a widget on a "free" dashboard
type WidgetLayout struct {
	X      *float64 `json:"x,omitempty" hcl:"x"`
	Y      *float64 `json:"y,omitempty" hcl:"y"`
	Height *float64 `json:"height,omitempty" hcl:"height"`
	Width  *float64 `json:"width,omitempty" hcl:"width"`
}

// AlertGraphDefinition represents the definition for an Alert Graph widget
type AlertGraphDefinition struct {
	AlertId    *string     `json:"alert_id" hcl:"alert_id"`
	VizType    *string     `json:"viz_type" hcl:"viz_type"`
	Title      *string     `json:"title,omitempty" hcl:"title"`
	TitleSize  *string     `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string     `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime `json:"time,omitempty" hcl:"time"`
}

// AlertValueDefinition represents the definition for an Alert Value widget
type AlertValueDefinition struct {
	AlertId    *string `json:"alert_id" hcl:"alert_id"`
	Precision  *int    `json:"precision,omitempty" hcl:"precision"`
	Unit       *string `json:"unit,omitempty" hcl:"unit"`
	TextAlign  *string `json:"text_align,omitempty" hcl:"text_align"`
	Title      *string `json:"title,omitempty" hcl:"title"`
	TitleSize  *string `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string `json:"title_align,omitempty" hcl:"title_align"`
}

// ChangeDefinition represents the definition for a Change widget
type ChangeDefinition struct {
	Requests   []ChangeRequest `json:"requests" hcl:"request,block"`
	Title      *string         `json:"title,omitempty" hcl:"title"`
	TitleSize  *string         `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string         `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime     `json:"time,omitempty" hcl:"time"`
}

type ChangeRequest struct {
	ChangeType   *string `json:"change_type,omitempty" hcl:"change_type"`
	CompareTo    *string `json:"compare_to,omitempty" hcl:"compare_to"`
	IncreaseGood *bool   `json:"increase_good,omitempty" hcl:"increase_good"`
	OrderBy      *string `json:"order_by,omitempty" hcl:"order_by"`
	OrderDir     *string `json:"order_dir,omitempty" hcl:"order_dir"`
	ShowPresent  *bool   `json:"show_present,omitempty" hcl:"show_present"`
	// A ChangeRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

// CheckStatusDefinition represents the definition for a Check Status widget
type CheckStatusDefinition struct {
	Check      *string     `json:"check" hcl:"check"`
	Grouping   *string     `json:"grouping" hcl:"grouping"`
	Group      *string     `json:"group,omitempty" hcl:"group"`
	GroupBy    []string    `json:"group_by,omitempty" hcl:"group_by"`
	Tags       []string    `json:"tags,omitempty" hcl:"tags"`
	Title      *string     `json:"title,omitempty" hcl:"title"`
	TitleSize  *string     `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string     `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime `json:"time,omitempty" hcl:"time"`
}

// DistributionDefinition represents the definition for a Distribution widget
type DistributionDefinition struct {
	Requests   []DistributionRequest `json:"requests" hcl:"request,block"`
	Title      *string               `json:"title,omitempty" hcl:"title"`
	TitleSize  *string               `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string               `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime           `json:"time,omitempty" hcl:"time"`
}

type DistributionRequest struct {
	Style *WidgetRequestStyle `json:"style,omitempty" hcl:"style"`
	// A DistributionRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

// EventStreamDefinition represents the definition for an Event Stream widget
type EventStreamDefinition struct {
	Query         *string     `json:"query" hcl:"query"`
	TagsExecution *string     `json:"tags_execution,omitempty" hcl:"tags_execution"`
	EventSize     *string     `json:"event_size,omitempty" hcl:"event_size"`
	Title         *string     `json:"title,omitempty" hcl:"title"`
	TitleSize     *string     `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign    *string     `json:"title_align,omitempty" hcl:"title_align"`
	Time          *WidgetTime `json:"time,omitempty" hcl:"time"`
}

// EventTimelineDefinition represents the definition for an Event Timeline widget
type EventTimelineDefinition struct {
	Query         *string     `json:"query" hcl:"query"`
	TagsExecution *string     `json:"tags_execution,omitempty" hcl:"tags_execution"`
	Title         *string     `json:"title,omitempty" hcl:"title"`
	TitleSize     *string     `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign    *string     `json:"title_align,omitempty" hcl:"title_align"`
	Time          *WidgetTime `json:"time,omitempty" hcl:"time"`
}

// FreeTextDefinition represents the definition for a Free Text widget
type FreeTextDefinition struct {
	Text      *string `json:"text" hcl:"text"`
	Color     *string `json:"color,omitempty" hcl:"color"`
	FontSize  *string `json:"font_size,omitempty" hcl:"font_size"`
	TextAlign *string `json:"text_align,omitempty" hcl:"text_align"`
}

// GroupDefinition represents the definition for an Group widget
type GroupDefinition struct {
	LayoutType *string       `json:"layout_type" hcl:"layout_type"`
	Widgets    []BoardWidget `json:"widgets" hcl:"widget,block"`
	Title      *string       `json:"title,omitempty" hcl:"title"`
}

// HeatmapDefinition represents the definition for a Heatmap widget
type HeatmapDefinition struct {
	Requests   []HeatmapRequest `json:"requests" hcl:"request,block"`
	Yaxis      *WidgetAxis      `json:"yaxis,omitempty" hcl:"yaxis"`
	Events     []WidgetEvent    `json:"events,omitempty" hcl:"event,block"`
	Title      *string          `json:"title,omitempty" hcl:"title"`
	TitleSize  *string          `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string          `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime      `json:"time,omitempty" hcl:"time"`
}

type HeatmapRequest struct {
	Style *WidgetRequestStyle `json:"style,omitempty" hcl:"style"`
	// A HeatmapRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

// HostmapDefinition represents the definition for a Hostmap widget
type HostmapDefinition struct {
	Requests      *HostmapRequests `json:"requests" hcl:"request,block"`
	NodeType      *string          `json:"node_type,omitempty" hcl:"node_type"`
	NoMetricHosts *bool            `json:"no_metric_hosts,omitempty" hcl:"no_metric_hosts"`
	NoGroupHosts  *bool            `json:"no_group_hosts,omitempty" hcl:"no_group_hosts"`
	Group         []string         `json:"group,omitempty" hcl:"group"`
	Scope         []string         `json:"scope,omitempty" hcl:"scope"`
	Style         *HostmapStyle    `json:"style,omitempty" hcl:"style"`
	Title         *string          `json:"title,omitempty" hcl:"title"`
	TitleSize     *string          `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign    *string          `json:"title_align,omitempty" hcl:"title_align"`
}

type HostmapRequests struct {
	Fill *HostmapRequest `json:"fill,omitempty" hcl:"fill"`
	Size *HostmapRequest `json:"size,omitempty" hcl:"size"`
}

type HostmapRequest struct {
	// A HostmapRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

type HostmapStyle struct {
	Palette     *string `json:"palette,omitempty" hcl:"palette"`
	PaletteFlip *bool   `json:"palette_flip,omitempty" hcl:"palette_flip"`
	FillMin     *string `json:"fill_min,omitempty" hcl:"fill_min"`
	FillMax     *string `json:"fill_max,omitempty" hcl:"fill_max"`
}

// IframeDefinition represents the definition for an Iframe widget
type IframeDefinition struct {
	Url *string `json:"url" hcl:"url"`
}

// ImageDefinition represents the definition for an Image widget
type ImageDefinition struct {
	Url    *string `json:"url" hcl:"url"`
	Sizing *string `json:"sizing,omitempty" hcl:"sizing"`
	Margin *string `json:"margin,omitempty" hcl:"margin"`
}

// LogStreamDefinition represents the definition for a Log Stream widget
type LogStreamDefinition struct {
	Logset            *string          `json:"logset" hcl:"logset"`
	Indexes           []string         `json:"indexes,omitempty" hcl:"indexes"`
	Query             *string          `json:"query,omitempty" hcl:"query"`
	Columns           []string         `json:"columns,omitempty" hcl:"columns"`
	Title             *string          `json:"title,omitempty" hcl:"title"`
	TitleSize         *string          `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign        *string          `json:"title_align,omitempty" hcl:"title_align"`
	Time              *WidgetTime      `json:"time,omitempty" hcl:"time"`
	ShowDateColumn    *bool            `json:"show_date_column,omitempty" hcl:"show_date_column"`
	ShowMessageColumn *bool            `json:"show_message_column,omitempty" hcl:"show_message_column"`
	MessageDisplay    *string          `json:"message_display,omitempty" hcl:"message_display"`
	Sort              *WidgetFieldSort `json:"sort,omitempty" hcl:"sort"`
}

type WidgetFieldSort struct {
	Column *string `json:"column" hcl:"column"`
	Order  *string `json:"order" hcl:"order"`
}

// ManageStatusDefinition represents the definition for a Manage Status widget
type ManageStatusDefinition struct {
	SummaryType       *string `json:"summary_type,omitempty" hcl:"summary_type"`
	Query             *string `json:"query" hcl:"query"`
	Sort              *string `json:"sort,omitempty" hcl:"sort"`
	Count             *int    `json:"count,omitempty" hcl:"count"`
	Start             *int    `json:"start,omitempty" hcl:"start"`
	DisplayFormat     *string `json:"display_format,omitempty" hcl:"display_format"`
	ColorPreference   *string `json:"color_preference,omitempty" hcl:"color_preference"`
	HideZeroCounts    *bool   `json:"hide_zero_counts,omitempty" hcl:"hide_zero_counts"`
	ShowLastTriggered *bool   `json:"show_last_triggered,omitempty" hcl:"show_last_triggered"`
	Title             *string `json:"title,omitempty" hcl:"title"`
	TitleSize         *string `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign        *string `json:"title_align,omitempty" hcl:"title_align"`
}

// NoteDefinition represents the definition for a Note widget
type NoteDefinition struct {
	Content         *string `json:"content" hcl:"content"`
	BackgroundColor *string `json:"background_color,omitempty" hcl:"background_color"`
	FontSize        *string `json:"font_size,omitempty" hcl:"font_size"`
	TextAlign       *string `json:"text_align,omitempty" hcl:"text_align"`
	ShowTick        *bool   `json:"show_tick,omitempty" hcl:"show_tick"`
	TickPos         *string `json:"tick_pos,omitempty" hcl:"tick_pos"`
	TickEdge        *string `json:"tick_edge,omitempty" hcl:"tick_edge"`
}

// QueryValueDefinition represents the definition for a Query Value widget
type QueryValueDefinition struct {
	Requests   []QueryValueRequest `json:"requests" hcl:"request,block"`
	Autoscale  *bool               `json:"autoscale,omitempty" hcl:"autoscale"`
	CustomUnit *string             `json:"custom_unit,omitempty" hcl:"custom_unit"`
	Precision  *int                `json:"precision,omitempty" hcl:"precision"`
	TextAlign  *string             `json:"text_align,omitempty" hcl:"text_align"`
	Title      *string             `json:"title,omitempty" hcl:"title"`
	TitleSize  *string             `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string             `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime         `json:"time,omitempty" hcl:"time"`
}

type QueryValueRequest struct {
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty" hcl:"conditional_formats,block"`
	Aggregator         *string                   `json:"aggregator,omitempty" hcl:"aggregator"`
	// A QueryValueRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

// QueryTableDefinition represents the definition for a Table widget
type QueryTableDefinition struct {
	Requests   []QueryTableRequest `json:"requests" hcl:"request,block"`
	Title      *string             `json:"title,omitempty" hcl:"title"`
	TitleSize  *string             `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string             `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime         `json:"time,omitempty" hcl:"time"`
}

type QueryTableRequest struct {
	Alias              *string                   `json:"alias,omitempty" hcl:"alias"`
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty" hcl:"conditional_formats"`
	Aggregator         *string                   `json:"aggregator,omitempty" hcl:"aggregator"`
	Limit              *int                      `json:"limit,omitempty" hcl:"limit"`
	Order              *string                   `json:"order,omitempty" hcl:"order"`
	// A QueryTableRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

// ScatterplotDefinition represents the definition for a Scatterplot widget
type ScatterplotDefinition struct {
	Requests      *ScatterplotRequests `json:"requests" hcl:"request,block"`
	Xaxis         *WidgetAxis          `json:"xaxis,omitempty" hcl:"xaxis"`
	Yaxis         *WidgetAxis          `json:"yaxis,omitempty" hcl:"yaxis"`
	ColorByGroups []string             `json:"color_by_groups,omitempty" hcl:"color_by_groups"`
	Title         *string              `json:"title,omitempty" hcl:"title"`
	TitleSize     *string              `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign    *string              `json:"title_align,omitempty" hcl:"title_align"`
	Time          *WidgetTime          `json:"time,omitempty" hcl:"time"`
}

type ScatterplotRequests struct {
	X *ScatterplotRequest `json:"x" hcl:"x"`
	Y *ScatterplotRequest `json:"y" hcl:"y"`
}

type ScatterplotRequest struct {
	Aggregator *string `json:"aggregator,omitempty" hcl:"aggregator"`
	// A ScatterplotRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

// ServiceLevelObjectiveDefinition represents the definition for a Service Level Objective widget
type ServiceLevelObjectiveDefinition struct {
	Title                   *string  `json:"title,omitempty" hcl:"title"`
	TitleSize               *string  `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign              *string  `json:"title_align,omitempty" hcl:"title_align"`
	ViewType                *string  `json:"view_type,omitempty" hcl:"view_type"`
	ServiceLevelObjectiveID *string  `json:"slo_id,omitempty" hcl:"slo_id"`
	ShowErrorBudget         *bool    `json:"show_error_budget,omitempty" hcl:"show_error_budget"`
	ViewMode                *string  `json:"view_mode,omitempty" hcl:"view_mode"`
	TimeWindows             []string `json:"time_windows,omitempty" hcl:"time_windows"`
}

// TimeseriesDefinition represents the definition for a Timeseries widget
type TimeseriesDefinition struct {
	Requests   []TimeseriesRequest `json:"requests" hcl:"request,block"`
	Yaxis      *WidgetAxis         `json:"yaxis,omitempty" hcl:"yaxis"`
	Events     []WidgetEvent       `json:"events,omitempty" hcl:"event,block"`
	Markers    []WidgetMarker      `json:"markers,omitempty" hcl:"marker,block"`
	Title      *string             `json:"title,omitempty" hcl:"title"`
	TitleSize  *string             `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string             `json:"title_align,omitempty" hcl:"title_align"`
	ShowLegend *bool               `json:"show_legend,omitempty" hcl:"show_legend"`
	LegendSize *string             `json:"legend_size,omitempty" hcl:"legend_size"`
	Time       *WidgetTime         `json:"time,omitempty" hcl:"time"`
}

type TimeseriesRequest struct {
	Style       *TimeseriesRequestStyle `json:"style,omitempty" hcl:"style"`
	Metadata    []WidgetMetadata        `json:"metadata,omitempty" hcl:"metadata,block"`
	DisplayType *string                 `json:"display_type,omitempty" hcl:"display_type"`
	// A TimeseriesRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}
type TimeseriesRequestStyle struct {
	Palette   *string `json:"palette,omitempty" hcl:"palette"`
	LineType  *string `json:"line_type,omitempty" hcl:"line_type"`
	LineWidth *string `json:"line_width,omitempty" hcl:"line_width"`
}

// ToplistDefinition represents the definition for a Top list widget
type ToplistDefinition struct {
	Requests   []ToplistRequest `json:"requests" hcl:"request,block"`
	Title      *string          `json:"title,omitempty" hcl:"title"`
	TitleSize  *string          `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign *string          `json:"title_align,omitempty" hcl:"title_align"`
	Time       *WidgetTime      `json:"time,omitempty" hcl:"time"`
}

type ToplistRequest struct {
	ConditionalFormats []WidgetConditionalFormat `json:"conditional_formats,omitempty" hcl:"conditional_formats"`
	Style              *WidgetRequestStyle       `json:"style,omitempty" hcl:"style"`
	// A ToplistRequest should implement exactly one of the following query types
	MetricQuery   *string              `json:"q,omitempty" hcl:"q"`
	ApmQuery      *WidgetApmOrLogQuery `json:"apm_query,omitempty" hcl:"apm_query" hcle:"omitempty"`
	LogQuery      *WidgetApmOrLogQuery `json:"log_query,omitempty" hcl:"log_query" hcle:"omitempty"`
	ProcessQuery  *WidgetProcessQuery  `json:"process_query,omitempty" hcl:"process_query" hcle:"omitempty"`
	RumQuery      *WidgetApmOrLogQuery `json:"rum_query,omitempty" hcl:"rum_query" hcle:"omitempty"`
	SecurityQuery *WidgetApmOrLogQuery `json:"security_query,omitempty" hcl:"security_query" hcle:"omitempty"`
}

// TraceServiceDefinition represents the definition for a Trace Service widget
type TraceServiceDefinition struct {
	Env              *string     `json:"env" hcl:"env"`
	Service          *string     `json:"service" hcl:"service"`
	SpanName         *string     `json:"span_name" hcl:"span_name"`
	ShowHits         *bool       `json:"show_hits,omitempty" hcl:"show_hits"`
	ShowErrors       *bool       `json:"show_errors,omitempty" hcl:"show_errors"`
	ShowLatency      *bool       `json:"show_latency,omitempty" hcl:"show_latency"`
	ShowBreakdown    *bool       `json:"show_breakdown,omitempty" hcl:"show_breakdown"`
	ShowDistribution *bool       `json:"show_distribution,omitempty" hcl:"show_distribution"`
	ShowResourceList *bool       `json:"show_resource_list,omitempty" hcl:"show_resource_list"`
	SizeFormat       *string     `json:"size_format,omitempty" hcl:"size_format"`
	DisplayFormat    *string     `json:"display_format,omitempty" hcl:"display_format"`
	Title            *string     `json:"title,omitempty" hcl:"title"`
	TitleSize        *string     `json:"title_size,omitempty" hcl:"title_size"`
	TitleAlign       *string     `json:"title_align,omitempty" hcl:"title_align"`
	Time             *WidgetTime `json:"time,omitempty" hcl:"time"`
}

// UnmarshalJSON is a Custom Unmarshal for BoardWidget. If first tries to unmarshal the data in a light
// struct that allows to get the widget type. Then based on the widget type, it will try to unmarshal the
// data using the corresponding widget struct.
func (widget *BoardWidget) UnmarshalJSON(data []byte) error {
	var widgetHandler struct {
		Definition *struct {
			Type *string `json:"type"`
		} `json:"definition"`
		Id     *int64        `json:"id,omitempty"`
		Layout *WidgetLayout `json:"layout,omitempty"`
	}
	if err := json.Unmarshal(data, &widgetHandler); err != nil {
		return err
	}

	// Get the widget id
	widget.Id = widgetHandler.Id

	// Get the widget layout
	widget.Layout = widgetHandler.Layout

	// Get the widget definition based on the widget type
	switch *widgetHandler.Definition.Type {
	case ALERT_GRAPH_WIDGET:
		var alertGraphWidget struct {
			Definition AlertGraphDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &alertGraphWidget); err != nil {
			return err
		}
		widget.AlertGraphDefinition = &alertGraphWidget.Definition
	case ALERT_VALUE_WIDGET:
		var alertValueWidget struct {
			Definition AlertValueDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &alertValueWidget); err != nil {
			return err
		}
		widget.AlertValueDefinition = &alertValueWidget.Definition
	case CHANGE_WIDGET:
		var changeWidget struct {
			Definition ChangeDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &changeWidget); err != nil {
			return err
		}
		widget.ChangeDefinition = &changeWidget.Definition
	case CHECK_STATUS_WIDGET:
		var checkStatusWidget struct {
			Definition CheckStatusDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &checkStatusWidget); err != nil {
			return err
		}
		widget.CheckStatusDefinition = &checkStatusWidget.Definition
	case DISTRIBUTION_WIDGET:
		var distributionWidget struct {
			Definition DistributionDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &distributionWidget); err != nil {
			return err
		}
		widget.DistributionDefinition = &distributionWidget.Definition
	case EVENT_STREAM_WIDGET:
		var eventStreamWidget struct {
			Definition EventStreamDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &eventStreamWidget); err != nil {
			return err
		}
		widget.EventStreamDefinition = &eventStreamWidget.Definition
	case EVENT_TIMELINE_WIDGET:
		var eventTimelineWidget struct {
			Definition EventTimelineDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &eventTimelineWidget); err != nil {
			return err
		}
		widget.EventTimelineDefinition = &eventTimelineWidget.Definition
	case FREE_TEXT_WIDGET:
		var freeTextWidget struct {
			Definition FreeTextDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &freeTextWidget); err != nil {
			return err
		}
		widget.FreeTextDefinition = &freeTextWidget.Definition
	case GROUP_WIDGET:
		var groupWidget struct {
			Definition GroupDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &groupWidget); err != nil {
			return err
		}
		widget.GroupDefinition = &groupWidget.Definition
	case HEATMAP_WIDGET:
		var heatmapWidget struct {
			Definition HeatmapDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &heatmapWidget); err != nil {
			return err
		}
		widget.HeatmapDefinition = &heatmapWidget.Definition
	case HOSTMAP_WIDGET:
		var hostmapWidget struct {
			Definition HostmapDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &hostmapWidget); err != nil {
			return err
		}
		widget.HostmapDefinition = &hostmapWidget.Definition
	case IFRAME_WIDGET:
		var iframeWidget struct {
			Definition IframeDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &iframeWidget); err != nil {
			return err
		}
		widget.IframeDefinition = &iframeWidget.Definition
	case IMAGE_WIDGET:
		var imageWidget struct {
			Definition ImageDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &imageWidget); err != nil {
			return err
		}
		widget.ImageDefinition = &imageWidget.Definition
	case LOG_STREAM_WIDGET:
		var logStreamWidget struct {
			Definition LogStreamDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &logStreamWidget); err != nil {
			return err
		}
		widget.LogStreamDefinition = &logStreamWidget.Definition
	case MANAGE_STATUS_WIDGET:
		var manageStatusWidget struct {
			Definition ManageStatusDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &manageStatusWidget); err != nil {
			return err
		}
		widget.ManageStatusDefinition = &manageStatusWidget.Definition
	case NOTE_WIDGET:
		var noteWidget struct {
			Definition NoteDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &noteWidget); err != nil {
			return err
		}
		widget.NoteDefinition = &noteWidget.Definition
	case QUERY_VALUE_WIDGET:
		var queryValueWidget struct {
			Definition QueryValueDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &queryValueWidget); err != nil {
			return err
		}
		widget.QueryValueDefinition = &queryValueWidget.Definition
	case QUERY_TABLE_WIDGET:
		var queryTableWidget struct {
			Definition QueryTableDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &queryTableWidget); err != nil {
			return err
		}
		widget.QueryTableDefinition = &queryTableWidget.Definition
	case SCATTERPLOT_WIDGET:
		var scatterplotWidget struct {
			Definition ScatterplotDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &scatterplotWidget); err != nil {
			return err
		}
		widget.ScatterplotDefinition = &scatterplotWidget.Definition
	case SERVICE_LEVEL_OBJECTIVE_WIDGET:
		var serviceLevelObjectiveWidget struct {
			Definition ServiceLevelObjectiveDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &serviceLevelObjectiveWidget); err != nil {
			return err
		}
		widget.ServiceLevelObjectiveDefinition = &serviceLevelObjectiveWidget.Definition
	case TIMESERIES_WIDGET:
		var timeseriesWidget struct {
			Definition TimeseriesDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &timeseriesWidget); err != nil {
			return err
		}
		widget.TimeseriesDefinition = &timeseriesWidget.Definition
	case TOPLIST_WIDGET:
		var toplistWidget struct {
			Definition ToplistDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &toplistWidget); err != nil {
			return err
		}
		widget.ToplistDefinition = &toplistWidget.Definition
	case TRACE_SERVICE_WIDGET:
		var traceServiceWidget struct {
			Definition TraceServiceDefinition `json:"definition"`
		}
		if err := json.Unmarshal(data, &traceServiceWidget); err != nil {
			return err
		}
		widget.TraceServiceDefinition = &traceServiceWidget.Definition
	default:
		return fmt.Errorf("Cannot unmarshal widget of type: %s", *widgetHandler.Definition.Type)
	}

	return nil
}

//
// List of structs common to multiple widget definitions
//

type WidgetTime struct {
	LiveSpan *string `json:"live_span,omitempty" hcl:"live_span"`
}

type WidgetAxis struct {
	Label       *string `json:"label,omitempty" hcl:"label"`
	Scale       *string `json:"scale,omitempty" hcl:"scale"`
	Min         *string `json:"min,omitempty" hcl:"min"`
	Max         *string `json:"max,omitempty" hcl:"max"`
	IncludeZero *bool   `json:"include_zero,omitempty" hcl:"include_zero"`
}

type WidgetEvent struct {
	Query *string `json:"q" hcl:"q"`
}

type WidgetMarker struct {
	Value       *string `json:"value" hcl:"value"`
	DisplayType *string `json:"display_type,omitempty" hcl:"display_type"`
	Label       *string `json:"label,omitempty" hcl:"label"`
}

type WidgetMetadata struct {
	Expression *string `json:"expression" hcl:"expression"`
	AliasName  *string `json:"alias_name,omitempty" hcl:"alias_name"`
}

type WidgetConditionalFormat struct {
	Comparator    *string  `json:"comparator" hcl:"comparator"`
	Value         *float64 `json:"value" hcl:"value"`
	Palette       *string  `json:"palette" hcl:"palette"`
	CustomBgColor *string  `json:"custom_bg_color,omitempty" hcl:"custom_bg_color"`
	CustomFgColor *string  `json:"custom_fg_color,omitempty" hcl:"custom_fg_color"`
	ImageUrl      *string  `json:"image_url,omitempty" hcl:"image_url"`
	HideValue     *bool    `json:"hide_value,omitempty" hcl:"hide_value"`
	Timeframe     *string  `json:"timeframe,omitempty" hcl:"timeframe"`
	Metric        *string  `json:"metric,omitempty" hcl:"metric"`
}

// WidgetApmOrLogQuery represents an APM or a Log query
type WidgetApmOrLogQuery struct {
	Index        *string                `json:"index" hcl:"index"`
	Compute      *ApmOrLogQueryCompute  `json:"compute,omitempty" hcl:"compute,attribute"`
	MultiCompute []ApmOrLogQueryCompute `json:"multi_compute,omitempty" hcl:"multi_compute"`
	Search       *ApmOrLogQuerySearch   `json:"search,omitempty" hcl:"search,attribute"`
	GroupBy      []ApmOrLogQueryGroupBy `json:"group_by,omitempty" hcl:"group_by"`
}
type ApmOrLogQueryCompute struct {
	Aggregation *string `json:"aggregation" hcl:"aggregation"`
	Facet       *string `json:"facet,omitempty" hcl:"facet"`
	Interval    *int    `json:"interval,omitempty" hcl:"interval"`
}
type ApmOrLogQuerySearch struct {
	Query *string `json:"query" hcl:"query"`
}
type ApmOrLogQueryGroupBy struct {
	Facet *string                   `json:"facet" hcl:"facet"`
	Limit *int                      `json:"limit,omitempty" hcl:"limit"`
	Sort  *ApmOrLogQueryGroupBySort `json:"sort,omitempty" hcl:"sort"`
}
type ApmOrLogQueryGroupBySort struct {
	Aggregation *string `json:"aggregation" hcl:"aggregation"`
	Order       *string `json:"order" hcl:"order"`
	Facet       *string `json:"facet,omitempty" hcl:"facet"`
}

// WidgetProcessQuery represents a Process query
type WidgetProcessQuery struct {
	Metric   *string  `json:"metric" hcl:"metric"`
	SearchBy *string  `json:"search_by,omitempty" hcl:"search_by"`
	FilterBy []string `json:"filter_by,omitempty" hcl:"filter_by"`
	Limit    *int     `json:"limit,omitempty" hcl:"limit"`
}

// WidgetRequestStyle represents the style that can be apply to a request
type WidgetRequestStyle struct {
	Palette *string `json:"palette,omitempty" hcl:"palette"`
}
