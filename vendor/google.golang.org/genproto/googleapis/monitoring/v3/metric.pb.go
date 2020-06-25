// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/monitoring/v3/metric.proto

package monitoring

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/distribution"
	label "google.golang.org/genproto/googleapis/api/label"
	metric "google.golang.org/genproto/googleapis/api/metric"
	monitoredres "google.golang.org/genproto/googleapis/api/monitoredres"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A single data point in a time series.
type Point struct {
	// The time interval to which the data point applies.  For `GAUGE` metrics,
	// the start time is optional, but if it is supplied, it must equal the
	// end time.  For `DELTA` metrics, the start
	// and end time should specify a non-zero interval, with subsequent points
	// specifying contiguous and non-overlapping intervals.  For `CUMULATIVE`
	// metrics, the start and end time should specify a non-zero interval, with
	// subsequent points specifying the same start time and increasing end times,
	// until an event resets the cumulative value to zero and sets a new start
	// time for the following points.
	Interval *TimeInterval `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	// The value of the data point.
	Value                *TypedValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{0}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetInterval() *TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Point) GetValue() *TypedValue {
	if m != nil {
		return m.Value
	}
	return nil
}

// A collection of data points that describes the time-varying values
// of a metric. A time series is identified by a combination of a
// fully-specified monitored resource and a fully-specified metric.
// This type is used for both listing and creating time series.
type TimeSeries struct {
	// The associated metric. A fully-specified metric used to identify the time
	// series.
	Metric *metric.Metric `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// The associated monitored resource.  Custom metrics can use only certain
	// monitored resource types in their time series data.
	Resource *monitoredres.MonitoredResource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// Output only. The associated monitored resource metadata. When reading a
	// a timeseries, this field will include metadata labels that are explicitly
	// named in the reduction. When creating a timeseries, this field is ignored.
	Metadata *monitoredres.MonitoredResourceMetadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The metric kind of the time series. When listing time series, this metric
	// kind might be different from the metric kind of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the metric kind of the associated metric. If the associated
	// metric's descriptor must be auto-created, then this field specifies the
	// metric kind of the new descriptor and must be either `GAUGE` (the default)
	// or `CUMULATIVE`.
	MetricKind metric.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	// The value type of the time series. When listing time series, this value
	// type might be different from the value type of the associated metric if
	// this time series is an alignment or reduction of other time series.
	//
	// When creating a time series, this field is optional. If present, it must be
	// the same as the type of the data in the `points` field.
	ValueType metric.MetricDescriptor_ValueType `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The data points of this time series. When listing time series, points are
	// returned in reverse time order.
	//
	// When creating a time series, this field must contain exactly one point and
	// the point's type must be the same as the value type of the associated
	// metric. If the associated metric's descriptor must be auto-created, then
	// the value type of the descriptor is determined by the point's type, which
	// must be `BOOL`, `INT64`, `DOUBLE`, or `DISTRIBUTION`.
	Points               []*Point `protobuf:"bytes,5,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{1}
}

func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeries.Unmarshal(m, b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
}
func (m *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(m, src)
}
func (m *TimeSeries) XXX_Size() int {
	return xxx_messageInfo_TimeSeries.Size(m)
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetMetric() *metric.Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (m *TimeSeries) GetResource() *monitoredres.MonitoredResource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TimeSeries) GetMetadata() *monitoredres.MonitoredResourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TimeSeries) GetMetricKind() metric.MetricDescriptor_MetricKind {
	if m != nil {
		return m.MetricKind
	}
	return metric.MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

func (m *TimeSeries) GetValueType() metric.MetricDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return metric.MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

// A descriptor for the labels and points in a timeseries.
type TimeSeriesDescriptor struct {
	// Descriptors for the labels.
	LabelDescriptors []*label.LabelDescriptor `protobuf:"bytes,1,rep,name=label_descriptors,json=labelDescriptors,proto3" json:"label_descriptors,omitempty"`
	// Descriptors for the point data value columns.
	PointDescriptors     []*TimeSeriesDescriptor_ValueDescriptor `protobuf:"bytes,5,rep,name=point_descriptors,json=pointDescriptors,proto3" json:"point_descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *TimeSeriesDescriptor) Reset()         { *m = TimeSeriesDescriptor{} }
func (m *TimeSeriesDescriptor) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesDescriptor) ProtoMessage()    {}
func (*TimeSeriesDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{2}
}

func (m *TimeSeriesDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesDescriptor.Unmarshal(m, b)
}
func (m *TimeSeriesDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesDescriptor.Marshal(b, m, deterministic)
}
func (m *TimeSeriesDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesDescriptor.Merge(m, src)
}
func (m *TimeSeriesDescriptor) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesDescriptor.Size(m)
}
func (m *TimeSeriesDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesDescriptor proto.InternalMessageInfo

func (m *TimeSeriesDescriptor) GetLabelDescriptors() []*label.LabelDescriptor {
	if m != nil {
		return m.LabelDescriptors
	}
	return nil
}

func (m *TimeSeriesDescriptor) GetPointDescriptors() []*TimeSeriesDescriptor_ValueDescriptor {
	if m != nil {
		return m.PointDescriptors
	}
	return nil
}

// A descriptor for the value columns in a data point.
type TimeSeriesDescriptor_ValueDescriptor struct {
	// The value key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value type.
	ValueType metric.MetricDescriptor_ValueType `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=google.api.MetricDescriptor_ValueType" json:"value_type,omitempty"`
	// The value stream kind.
	MetricKind           metric.MetricDescriptor_MetricKind `protobuf:"varint,3,opt,name=metric_kind,json=metricKind,proto3,enum=google.api.MetricDescriptor_MetricKind" json:"metric_kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *TimeSeriesDescriptor_ValueDescriptor) Reset()         { *m = TimeSeriesDescriptor_ValueDescriptor{} }
func (m *TimeSeriesDescriptor_ValueDescriptor) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesDescriptor_ValueDescriptor) ProtoMessage()    {}
func (*TimeSeriesDescriptor_ValueDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{2, 0}
}

func (m *TimeSeriesDescriptor_ValueDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesDescriptor_ValueDescriptor.Unmarshal(m, b)
}
func (m *TimeSeriesDescriptor_ValueDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesDescriptor_ValueDescriptor.Marshal(b, m, deterministic)
}
func (m *TimeSeriesDescriptor_ValueDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesDescriptor_ValueDescriptor.Merge(m, src)
}
func (m *TimeSeriesDescriptor_ValueDescriptor) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesDescriptor_ValueDescriptor.Size(m)
}
func (m *TimeSeriesDescriptor_ValueDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesDescriptor_ValueDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesDescriptor_ValueDescriptor proto.InternalMessageInfo

func (m *TimeSeriesDescriptor_ValueDescriptor) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TimeSeriesDescriptor_ValueDescriptor) GetValueType() metric.MetricDescriptor_ValueType {
	if m != nil {
		return m.ValueType
	}
	return metric.MetricDescriptor_VALUE_TYPE_UNSPECIFIED
}

func (m *TimeSeriesDescriptor_ValueDescriptor) GetMetricKind() metric.MetricDescriptor_MetricKind {
	if m != nil {
		return m.MetricKind
	}
	return metric.MetricDescriptor_METRIC_KIND_UNSPECIFIED
}

// Represents the values of a time series associated with a
// TimeSeriesDescriptor.
type TimeSeriesData struct {
	// The values of the labels in the time series identifier, given in the same
	// order as the `label_descriptors` field of the TimeSeriesDescriptor
	// associated with this object. Each value must have a value of the type
	// given in the corresponding entry of `label_descriptors`.
	LabelValues []*LabelValue `protobuf:"bytes,1,rep,name=label_values,json=labelValues,proto3" json:"label_values,omitempty"`
	// The points in the time series.
	PointData            []*TimeSeriesData_PointData `protobuf:"bytes,2,rep,name=point_data,json=pointData,proto3" json:"point_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TimeSeriesData) Reset()         { *m = TimeSeriesData{} }
func (m *TimeSeriesData) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesData) ProtoMessage()    {}
func (*TimeSeriesData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{3}
}

func (m *TimeSeriesData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesData.Unmarshal(m, b)
}
func (m *TimeSeriesData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesData.Marshal(b, m, deterministic)
}
func (m *TimeSeriesData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesData.Merge(m, src)
}
func (m *TimeSeriesData) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesData.Size(m)
}
func (m *TimeSeriesData) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesData.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesData proto.InternalMessageInfo

func (m *TimeSeriesData) GetLabelValues() []*LabelValue {
	if m != nil {
		return m.LabelValues
	}
	return nil
}

func (m *TimeSeriesData) GetPointData() []*TimeSeriesData_PointData {
	if m != nil {
		return m.PointData
	}
	return nil
}

// A point's value columns and time interval. Each point has one or more
// point values corresponding to the entries in `point_descriptors` field in
// the TimeSeriesDescriptor associated with this object.
type TimeSeriesData_PointData struct {
	// The values that make up the point.
	Values []*TypedValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// The time interval associated with the point.
	TimeInterval         *TimeInterval `protobuf:"bytes,2,opt,name=time_interval,json=timeInterval,proto3" json:"time_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TimeSeriesData_PointData) Reset()         { *m = TimeSeriesData_PointData{} }
func (m *TimeSeriesData_PointData) String() string { return proto.CompactTextString(m) }
func (*TimeSeriesData_PointData) ProtoMessage()    {}
func (*TimeSeriesData_PointData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{3, 0}
}

func (m *TimeSeriesData_PointData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeriesData_PointData.Unmarshal(m, b)
}
func (m *TimeSeriesData_PointData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeriesData_PointData.Marshal(b, m, deterministic)
}
func (m *TimeSeriesData_PointData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeriesData_PointData.Merge(m, src)
}
func (m *TimeSeriesData_PointData) XXX_Size() int {
	return xxx_messageInfo_TimeSeriesData_PointData.Size(m)
}
func (m *TimeSeriesData_PointData) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeriesData_PointData.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeriesData_PointData proto.InternalMessageInfo

func (m *TimeSeriesData_PointData) GetValues() []*TypedValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *TimeSeriesData_PointData) GetTimeInterval() *TimeInterval {
	if m != nil {
		return m.TimeInterval
	}
	return nil
}

// A label value.
type LabelValue struct {
	// The label value can be a bool, int64, or string.
	//
	// Types that are valid to be assigned to Value:
	//	*LabelValue_BoolValue
	//	*LabelValue_Int64Value
	//	*LabelValue_StringValue
	Value                isLabelValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LabelValue) Reset()         { *m = LabelValue{} }
func (m *LabelValue) String() string { return proto.CompactTextString(m) }
func (*LabelValue) ProtoMessage()    {}
func (*LabelValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{4}
}

func (m *LabelValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelValue.Unmarshal(m, b)
}
func (m *LabelValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelValue.Marshal(b, m, deterministic)
}
func (m *LabelValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelValue.Merge(m, src)
}
func (m *LabelValue) XXX_Size() int {
	return xxx_messageInfo_LabelValue.Size(m)
}
func (m *LabelValue) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelValue.DiscardUnknown(m)
}

var xxx_messageInfo_LabelValue proto.InternalMessageInfo

type isLabelValue_Value interface {
	isLabelValue_Value()
}

type LabelValue_BoolValue struct {
	BoolValue bool `protobuf:"varint,1,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type LabelValue_Int64Value struct {
	Int64Value int64 `protobuf:"varint,2,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type LabelValue_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*LabelValue_BoolValue) isLabelValue_Value() {}

func (*LabelValue_Int64Value) isLabelValue_Value() {}

func (*LabelValue_StringValue) isLabelValue_Value() {}

func (m *LabelValue) GetValue() isLabelValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LabelValue) GetBoolValue() bool {
	if x, ok := m.GetValue().(*LabelValue_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *LabelValue) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*LabelValue_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *LabelValue) GetStringValue() string {
	if x, ok := m.GetValue().(*LabelValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LabelValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LabelValue_BoolValue)(nil),
		(*LabelValue_Int64Value)(nil),
		(*LabelValue_StringValue)(nil),
	}
}

// An error associated with a query in the time series query language format.
type QueryError struct {
	// The location of the time series query language text that this error applies
	// to.
	Locator *TextLocator `protobuf:"bytes,1,opt,name=locator,proto3" json:"locator,omitempty"`
	// The error message.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryError) Reset()         { *m = QueryError{} }
func (m *QueryError) String() string { return proto.CompactTextString(m) }
func (*QueryError) ProtoMessage()    {}
func (*QueryError) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{5}
}

func (m *QueryError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryError.Unmarshal(m, b)
}
func (m *QueryError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryError.Marshal(b, m, deterministic)
}
func (m *QueryError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryError.Merge(m, src)
}
func (m *QueryError) XXX_Size() int {
	return xxx_messageInfo_QueryError.Size(m)
}
func (m *QueryError) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryError.DiscardUnknown(m)
}

var xxx_messageInfo_QueryError proto.InternalMessageInfo

func (m *QueryError) GetLocator() *TextLocator {
	if m != nil {
		return m.Locator
	}
	return nil
}

func (m *QueryError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A locator for text. Indicates a particular part of the text of a request or
// of an object referenced in the request.
//
// For example, suppose the request field `text` contains:
//
//   text: "The quick brown fox jumps over the lazy dog."
//
// Then the locator:
//
//   source: "text"
//   start_position {
//     line: 1
//     column: 17
//   }
//   end_position {
//     line: 1
//     column: 19
//   }
//
// refers to the part of the text: "fox".
type TextLocator struct {
	// The source of the text. The source may be a field in the request, in which
	// case its format is the format of the
	// google.rpc.BadRequest.FieldViolation.field field in
	// https://cloud.google.com/apis/design/errors#error_details. It may also be
	// be a source other than the request field (e.g. a macro definition
	// referenced in the text of the query), in which case this is the name of
	// the source (e.g. the macro name).
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The position of the first byte within the text.
	StartPosition *TextLocator_Position `protobuf:"bytes,2,opt,name=start_position,json=startPosition,proto3" json:"start_position,omitempty"`
	// The position of the last byte within the text.
	EndPosition *TextLocator_Position `protobuf:"bytes,3,opt,name=end_position,json=endPosition,proto3" json:"end_position,omitempty"`
	// If `source`, `start_position`, and `end_position` describe a call on
	// some object (e.g. a macro in the time series query language text) and a
	// location is to be designated in that object's text, `nested_locator`
	// identifies the location within that object.
	NestedLocator *TextLocator `protobuf:"bytes,4,opt,name=nested_locator,json=nestedLocator,proto3" json:"nested_locator,omitempty"`
	// When `nested_locator` is set, this field gives the reason for the nesting.
	// Usually, the reason is a macro invocation. In that case, the macro name
	// (including the leading '@') signals the location of the macro call
	// in the text and a macro argument name (including the leading '$') signals
	// the location of the macro argument inside the macro body that got
	// substituted away.
	NestingReason        string   `protobuf:"bytes,5,opt,name=nesting_reason,json=nestingReason,proto3" json:"nesting_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextLocator) Reset()         { *m = TextLocator{} }
func (m *TextLocator) String() string { return proto.CompactTextString(m) }
func (*TextLocator) ProtoMessage()    {}
func (*TextLocator) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{6}
}

func (m *TextLocator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextLocator.Unmarshal(m, b)
}
func (m *TextLocator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextLocator.Marshal(b, m, deterministic)
}
func (m *TextLocator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextLocator.Merge(m, src)
}
func (m *TextLocator) XXX_Size() int {
	return xxx_messageInfo_TextLocator.Size(m)
}
func (m *TextLocator) XXX_DiscardUnknown() {
	xxx_messageInfo_TextLocator.DiscardUnknown(m)
}

var xxx_messageInfo_TextLocator proto.InternalMessageInfo

func (m *TextLocator) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *TextLocator) GetStartPosition() *TextLocator_Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *TextLocator) GetEndPosition() *TextLocator_Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *TextLocator) GetNestedLocator() *TextLocator {
	if m != nil {
		return m.NestedLocator
	}
	return nil
}

func (m *TextLocator) GetNestingReason() string {
	if m != nil {
		return m.NestingReason
	}
	return ""
}

// The position of a byte within the text.
type TextLocator_Position struct {
	// The line, starting with 1, where the byte is positioned.
	Line int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	// The column within the line, starting with 1, where the byte is
	// positioned. This is a byte index even though the text is UTF-8.
	Column               int32    `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextLocator_Position) Reset()         { *m = TextLocator_Position{} }
func (m *TextLocator_Position) String() string { return proto.CompactTextString(m) }
func (*TextLocator_Position) ProtoMessage()    {}
func (*TextLocator_Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_c76199a3d2c4c21e, []int{6, 0}
}

func (m *TextLocator_Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextLocator_Position.Unmarshal(m, b)
}
func (m *TextLocator_Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextLocator_Position.Marshal(b, m, deterministic)
}
func (m *TextLocator_Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextLocator_Position.Merge(m, src)
}
func (m *TextLocator_Position) XXX_Size() int {
	return xxx_messageInfo_TextLocator_Position.Size(m)
}
func (m *TextLocator_Position) XXX_DiscardUnknown() {
	xxx_messageInfo_TextLocator_Position.DiscardUnknown(m)
}

var xxx_messageInfo_TextLocator_Position proto.InternalMessageInfo

func (m *TextLocator_Position) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *TextLocator_Position) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "google.monitoring.v3.Point")
	proto.RegisterType((*TimeSeries)(nil), "google.monitoring.v3.TimeSeries")
	proto.RegisterType((*TimeSeriesDescriptor)(nil), "google.monitoring.v3.TimeSeriesDescriptor")
	proto.RegisterType((*TimeSeriesDescriptor_ValueDescriptor)(nil), "google.monitoring.v3.TimeSeriesDescriptor.ValueDescriptor")
	proto.RegisterType((*TimeSeriesData)(nil), "google.monitoring.v3.TimeSeriesData")
	proto.RegisterType((*TimeSeriesData_PointData)(nil), "google.monitoring.v3.TimeSeriesData.PointData")
	proto.RegisterType((*LabelValue)(nil), "google.monitoring.v3.LabelValue")
	proto.RegisterType((*QueryError)(nil), "google.monitoring.v3.QueryError")
	proto.RegisterType((*TextLocator)(nil), "google.monitoring.v3.TextLocator")
	proto.RegisterType((*TextLocator_Position)(nil), "google.monitoring.v3.TextLocator.Position")
}

func init() {
	proto.RegisterFile("google/monitoring/v3/metric.proto", fileDescriptor_c76199a3d2c4c21e)
}

var fileDescriptor_c76199a3d2c4c21e = []byte{
	// 862 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0x1c, 0x35,
	0x14, 0xce, 0xee, 0x66, 0x37, 0xbb, 0x67, 0x92, 0xd0, 0x5a, 0x55, 0x19, 0x2d, 0x2a, 0x24, 0x53,
	0x15, 0xa2, 0x5e, 0xcc, 0x4a, 0x59, 0x14, 0x01, 0x95, 0x2a, 0x91, 0x52, 0x75, 0x11, 0x8d, 0x94,
	0x1a, 0x94, 0x0b, 0x14, 0x69, 0xe4, 0x9d, 0x31, 0x23, 0xab, 0x33, 0xf6, 0xc8, 0xe3, 0x59, 0x91,
	0x1b, 0xe0, 0x09, 0x78, 0x09, 0xee, 0xb8, 0x86, 0x97, 0xe0, 0x69, 0x78, 0x00, 0x2e, 0x90, 0x7f,
	0xe6, 0x67, 0x43, 0xb6, 0x0d, 0x88, 0x3b, 0xfb, 0x9c, 0xef, 0x7c, 0xe7, 0xf8, 0x3b, 0xf6, 0x31,
	0x1c, 0xa6, 0x42, 0xa4, 0x19, 0x9d, 0xe5, 0x82, 0x33, 0x25, 0x24, 0xe3, 0xe9, 0x6c, 0x35, 0x9f,
	0xe5, 0x54, 0x49, 0x16, 0x87, 0x85, 0x14, 0x4a, 0xa0, 0x7b, 0x16, 0x12, 0xb6, 0x90, 0x70, 0x35,
	0x9f, 0x3e, 0x70, 0x81, 0xa4, 0x60, 0xb3, 0x84, 0x95, 0x4a, 0xb2, 0x65, 0xa5, 0x98, 0xe0, 0x36,
	0x68, 0x7a, 0xbf, 0xe3, 0xce, 0xc8, 0x92, 0x66, 0xce, 0xfe, 0x6e, 0xc7, 0xde, 0xcd, 0x32, 0x7d,
	0xd8, 0x75, 0xd8, 0x4c, 0x34, 0x89, 0x24, 0x2d, 0x45, 0x25, 0x63, 0xea, 0x40, 0x37, 0x57, 0x1b,
	0x8b, 0x3c, 0x6f, 0x12, 0xbf, 0xef, 0x20, 0x66, 0xb7, 0xac, 0xbe, 0x9b, 0x25, 0x95, 0x24, 0x6d,
	0x61, 0xc1, 0x8f, 0x30, 0x3c, 0x17, 0x8c, 0x2b, 0xf4, 0x14, 0xc6, 0x8c, 0x2b, 0x2a, 0x57, 0x24,
	0xf3, 0x7b, 0x07, 0xbd, 0x23, 0xef, 0x38, 0x08, 0x6f, 0x3a, 0x69, 0xf8, 0x0d, 0xcb, 0xe9, 0x97,
	0x0e, 0x89, 0x9b, 0x18, 0x74, 0x02, 0xc3, 0x15, 0xc9, 0x2a, 0xea, 0xf7, 0x4d, 0xf0, 0xc1, 0x86,
	0xe0, 0xab, 0x82, 0x26, 0x17, 0x1a, 0x87, 0x2d, 0x3c, 0xf8, 0xab, 0x0f, 0xa0, 0x29, 0xbf, 0xa6,
	0x92, 0xd1, 0x12, 0x3d, 0x86, 0x91, 0xd5, 0xc1, 0x15, 0x81, 0x6a, 0x1e, 0x52, 0xb0, 0xf0, 0xcc,
	0x78, 0xb0, 0x43, 0xa0, 0x4f, 0x61, 0x5c, 0x0b, 0xe2, 0xb2, 0x3e, 0x58, 0x43, 0xd7, 0xb2, 0x61,
	0x07, 0xc2, 0x0d, 0x1c, 0x7d, 0x0e, 0xe3, 0x9c, 0x2a, 0x92, 0x10, 0x45, 0xfc, 0x1d, 0x13, 0xfa,
	0xe8, 0x8d, 0xa1, 0x67, 0x0e, 0x8c, 0x9b, 0x30, 0xb4, 0x00, 0xcf, 0xd6, 0x11, 0xbd, 0x66, 0x3c,
	0xf1, 0x07, 0x07, 0xbd, 0xa3, 0xfd, 0xe3, 0x8f, 0xfe, 0x59, 0xee, 0x17, 0xb4, 0x8c, 0x25, 0x2b,
	0x94, 0x90, 0xce, 0xf0, 0x15, 0xe3, 0x09, 0x86, 0xbc, 0x59, 0xa3, 0xe7, 0x00, 0x46, 0x8b, 0x48,
	0x5d, 0x15, 0xd4, 0xdf, 0x36, 0x44, 0x1f, 0xbe, 0x91, 0xc8, 0x28, 0xa8, 0xb5, 0xc4, 0x93, 0x55,
	0xbd, 0x44, 0x73, 0x18, 0x15, 0xba, 0x95, 0xa5, 0x3f, 0x3c, 0x18, 0x1c, 0x79, 0xc7, 0xef, 0xdd,
	0xdc, 0x02, 0xd3, 0x6e, 0xec, 0xa0, 0xc1, 0x4f, 0x03, 0xb8, 0xd7, 0xca, 0xdf, 0xa6, 0x40, 0x0b,
	0xb8, 0x6b, 0x2e, 0x6a, 0x94, 0x34, 0xb6, 0xd2, 0xef, 0xad, 0x13, 0xeb, 0xda, 0x5e, 0x6a, 0x50,
	0x1b, 0x87, 0xef, 0x64, 0xeb, 0x86, 0x12, 0xa5, 0x70, 0xd7, 0x24, 0x5b, 0x63, 0xb2, 0x25, 0x7e,
	0xb6, 0xf9, 0x8a, 0x5d, 0x2f, 0xc8, 0x9e, 0xb9, 0x9b, 0xc8, 0x90, 0x76, 0x12, 0x4d, 0x7f, 0xeb,
	0xc1, 0x3b, 0xd7, 0x50, 0xe8, 0x0e, 0x0c, 0x5e, 0xd3, 0x2b, 0x73, 0x99, 0x26, 0x58, 0x2f, 0xaf,
	0xa9, 0xdd, 0xff, 0xaf, 0x6a, 0xff, 0x6f, 0xed, 0x0f, 0x7e, 0xef, 0xc3, 0x7e, 0xe7, 0xc4, 0xfa,
	0x6e, 0x3d, 0x83, 0x5d, 0x2b, 0xbe, 0xc9, 0x57, 0xeb, 0xbe, 0xe1, 0x4d, 0x99, 0x0e, 0xd8, 0x37,
	0xe5, 0x65, 0xcd, 0xba, 0x44, 0x67, 0x00, 0x4e, 0x77, 0x7d, 0xcb, 0xfb, 0x86, 0x22, 0x7c, 0xab,
	0xe0, 0x44, 0x11, 0x7b, 0x45, 0xf4, 0x0a, 0x4f, 0x8a, 0x7a, 0x39, 0xfd, 0xb9, 0x07, 0x93, 0xc6,
	0x81, 0x3e, 0x81, 0xd1, 0x6d, 0x6a, 0xeb, 0xbc, 0x77, 0x87, 0x47, 0x2f, 0x60, 0x4f, 0xb1, 0x9c,
	0x46, 0xcd, 0xb4, 0xe9, 0xdf, 0x7a, 0xda, 0xec, 0xaa, 0xce, 0x2e, 0xf8, 0x01, 0xa0, 0x3d, 0x3a,
	0xfa, 0x00, 0x60, 0x29, 0x84, 0x53, 0xcc, 0xf4, 0x7b, 0xbc, 0xd8, 0xc2, 0x13, 0x6d, 0xb3, 0x80,
	0x43, 0xf0, 0x18, 0x57, 0x27, 0x1f, 0x47, 0xed, 0x98, 0x1a, 0x2c, 0xb6, 0x30, 0x18, 0xa3, 0x85,
	0x3c, 0x84, 0x5d, 0x3d, 0xb9, 0x79, 0xea, 0x30, 0xba, 0xa9, 0x93, 0xc5, 0x16, 0xf6, 0xac, 0xd5,
	0x80, 0x4e, 0x77, 0xdc, 0xa0, 0x0b, 0x62, 0x80, 0x57, 0x15, 0x95, 0x57, 0xcf, 0xa5, 0x14, 0x12,
	0x3d, 0x81, 0x9d, 0x4c, 0xc4, 0x44, 0x09, 0xe9, 0x26, 0xd7, 0xe1, 0x86, 0x03, 0xd1, 0xef, 0xd5,
	0x4b, 0x0b, 0xc4, 0x75, 0x04, 0xf2, 0x61, 0x27, 0xa7, 0x65, 0x49, 0x52, 0x5b, 0xd7, 0x04, 0xd7,
	0xdb, 0xe0, 0xcf, 0x3e, 0x78, 0x9d, 0x10, 0x74, 0x1f, 0x46, 0x6e, 0xe2, 0xd9, 0x2b, 0xed, 0x76,
	0xe8, 0x15, 0xec, 0x97, 0x8a, 0x48, 0x15, 0x15, 0xa2, 0x64, 0x7a, 0xbe, 0x3b, 0x59, 0x1f, 0xbf,
	0xb5, 0x8a, 0xf0, 0xdc, 0x45, 0xe0, 0x3d, 0xc3, 0x50, 0x6f, 0xd1, 0x19, 0xec, 0x52, 0x9e, 0xb4,
	0x84, 0x83, 0x7f, 0x4d, 0xe8, 0x51, 0x9e, 0x34, 0x74, 0x0b, 0xd8, 0xe7, 0xb4, 0x54, 0x34, 0x89,
	0x6a, 0x9d, 0xb6, 0x6f, 0xab, 0xd3, 0x9e, 0x0d, 0xac, 0x35, 0x78, 0x64, 0x99, 0x74, 0x9f, 0x24,
	0x25, 0xa5, 0xe0, 0xfe, 0xd0, 0x68, 0xb1, 0xe7, 0xac, 0xd8, 0x18, 0xa7, 0x27, 0x30, 0x6e, 0x92,
	0x23, 0xd8, 0xce, 0x18, 0xb7, 0xa2, 0x0d, 0xb1, 0x59, 0x6b, 0x29, 0x63, 0x91, 0x55, 0xb9, 0x95,
	0x6a, 0x88, 0xdd, 0xee, 0xf4, 0x97, 0x1e, 0xf8, 0xb1, 0xc8, 0x6f, 0x2c, 0xeb, 0xd4, 0xb3, 0x8f,
	0xf8, 0x5c, 0x7f, 0x9e, 0xe7, 0xbd, 0x6f, 0x9f, 0x3a, 0x50, 0x2a, 0x32, 0xc2, 0xd3, 0x50, 0xc8,
	0x74, 0x96, 0x52, 0x6e, 0xbe, 0xd6, 0x99, 0x75, 0x91, 0x82, 0x95, 0xeb, 0x1f, 0xf4, 0x93, 0x76,
	0xf7, 0x6b, 0x7f, 0xfa, 0xc2, 0x12, 0x3c, 0xcb, 0x44, 0x95, 0xd4, 0xff, 0x8e, 0xce, 0x75, 0x31,
	0xff, 0xa3, 0x76, 0x5e, 0x1a, 0xe7, 0x65, 0xeb, 0xbc, 0xbc, 0x98, 0x2f, 0x47, 0x26, 0xc9, 0xfc,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0xd6, 0x77, 0xe8, 0xb2, 0x08, 0x00, 0x00,
}
