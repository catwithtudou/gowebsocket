// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datalabeling/v1beta1/evaluation_job.proto

package datalabeling

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// State of the job.
type EvaluationJob_State int32

const (
	EvaluationJob_STATE_UNSPECIFIED EvaluationJob_State = 0
	EvaluationJob_SCHEDULED         EvaluationJob_State = 1
	EvaluationJob_RUNNING           EvaluationJob_State = 2
	EvaluationJob_PAUSED            EvaluationJob_State = 3
	EvaluationJob_STOPPED           EvaluationJob_State = 4
)

var EvaluationJob_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "SCHEDULED",
	2: "RUNNING",
	3: "PAUSED",
	4: "STOPPED",
}

var EvaluationJob_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"SCHEDULED":         1,
	"RUNNING":           2,
	"PAUSED":            3,
	"STOPPED":           4,
}

func (x EvaluationJob_State) String() string {
	return proto.EnumName(EvaluationJob_State_name, int32(x))
}

func (EvaluationJob_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_052df044fb4ef883, []int{0, 0}
}

// Defines an evaluation job that is triggered periodically to generate
// evaluations.
type EvaluationJob struct {
	// Format: 'projects/{project_id}/evaluationJobs/{evaluation_job_id}'
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the job. The description can be up to
	// 25000 characters long.
	Description string              `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	State       EvaluationJob_State `protobuf:"varint,3,opt,name=state,proto3,enum=google.cloud.datalabeling.v1beta1.EvaluationJob_State" json:"state,omitempty"`
	// Describes the schedule on which the job will be executed. Minimum schedule
	// unit is 1 day.
	//
	// The schedule can be either of the following types:
	// * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview)
	// * English-like
	//
	// [schedule](https:
	// //cloud.google.com/scheduler/docs/configuring/cron-job-schedules)
	Schedule string `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	// The versioned model that is being evaluated here.
	// Only one job is allowed for each model name.
	// Format: 'projects/*/models/*/versions/*'
	ModelVersion string `protobuf:"bytes,5,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	// Detailed config for running this eval job.
	EvaluationJobConfig *EvaluationJobConfig `protobuf:"bytes,6,opt,name=evaluation_job_config,json=evaluationJobConfig,proto3" json:"evaluation_job_config,omitempty"`
	// Name of the AnnotationSpecSet.
	AnnotationSpecSet string `protobuf:"bytes,7,opt,name=annotation_spec_set,json=annotationSpecSet,proto3" json:"annotation_spec_set,omitempty"`
	// If a human annotation should be requested when some data don't have ground
	// truth.
	LabelMissingGroundTruth bool `protobuf:"varint,8,opt,name=label_missing_ground_truth,json=labelMissingGroundTruth,proto3" json:"label_missing_ground_truth,omitempty"`
	// Output only. Any attempts with errors happening in evaluation job runs each
	// time will be recorded here incrementally.
	Attempts []*Attempt `protobuf:"bytes,9,rep,name=attempts,proto3" json:"attempts,omitempty"`
	// Timestamp when this evaluation job was created.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EvaluationJob) Reset()         { *m = EvaluationJob{} }
func (m *EvaluationJob) String() string { return proto.CompactTextString(m) }
func (*EvaluationJob) ProtoMessage()    {}
func (*EvaluationJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_052df044fb4ef883, []int{0}
}

func (m *EvaluationJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluationJob.Unmarshal(m, b)
}
func (m *EvaluationJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluationJob.Marshal(b, m, deterministic)
}
func (m *EvaluationJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluationJob.Merge(m, src)
}
func (m *EvaluationJob) XXX_Size() int {
	return xxx_messageInfo_EvaluationJob.Size(m)
}
func (m *EvaluationJob) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluationJob.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluationJob proto.InternalMessageInfo

func (m *EvaluationJob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EvaluationJob) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EvaluationJob) GetState() EvaluationJob_State {
	if m != nil {
		return m.State
	}
	return EvaluationJob_STATE_UNSPECIFIED
}

func (m *EvaluationJob) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *EvaluationJob) GetModelVersion() string {
	if m != nil {
		return m.ModelVersion
	}
	return ""
}

func (m *EvaluationJob) GetEvaluationJobConfig() *EvaluationJobConfig {
	if m != nil {
		return m.EvaluationJobConfig
	}
	return nil
}

func (m *EvaluationJob) GetAnnotationSpecSet() string {
	if m != nil {
		return m.AnnotationSpecSet
	}
	return ""
}

func (m *EvaluationJob) GetLabelMissingGroundTruth() bool {
	if m != nil {
		return m.LabelMissingGroundTruth
	}
	return false
}

func (m *EvaluationJob) GetAttempts() []*Attempt {
	if m != nil {
		return m.Attempts
	}
	return nil
}

func (m *EvaluationJob) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type EvaluationJobConfig struct {
	// config specific to different supported human annotation use cases.
	//
	// Types that are valid to be assigned to HumanAnnotationRequestConfig:
	//	*EvaluationJobConfig_ImageClassificationConfig
	//	*EvaluationJobConfig_BoundingPolyConfig
	//	*EvaluationJobConfig_VideoClassificationConfig
	//	*EvaluationJobConfig_ObjectDetectionConfig
	//	*EvaluationJobConfig_TextClassificationConfig
	//	*EvaluationJobConfig_ObjectTrackingConfig
	HumanAnnotationRequestConfig isEvaluationJobConfig_HumanAnnotationRequestConfig `protobuf_oneof:"human_annotation_request_config"`
	// Input config for data, gcs_source in the config will be the root path for
	// data. Data should be organzied chronically under that path.
	InputConfig *InputConfig `protobuf:"bytes,1,opt,name=input_config,json=inputConfig,proto3" json:"input_config,omitempty"`
	// Config used to create evaluation.
	EvaluationConfig      *EvaluationConfig      `protobuf:"bytes,2,opt,name=evaluation_config,json=evaluationConfig,proto3" json:"evaluation_config,omitempty"`
	HumanAnnotationConfig *HumanAnnotationConfig `protobuf:"bytes,3,opt,name=human_annotation_config,json=humanAnnotationConfig,proto3" json:"human_annotation_config,omitempty"`
	// Mappings between reserved keys for bigquery import and customized tensor
	// names. Key is the reserved key, value is tensor name in the bigquery table.
	// Different annotation type has different required key mapping. See user
	// manual for more details:
	//
	// https:
	// //docs.google.com/document/d/1bg1meMIBGY
	// // 9I5QEoFoHSX6u9LsZQYBSmPt6E9SxqHZc/edit#heading=h.tfyjhxhvsqem
	BigqueryImportKeys map[string]string `protobuf:"bytes,9,rep,name=bigquery_import_keys,json=bigqueryImportKeys,proto3" json:"bigquery_import_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Max number of examples to collect in each period.
	ExampleCount int32 `protobuf:"varint,10,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	// Percentage of examples to collect in each period. 0.1 means 10% of total
	// examples will be collected, and 0.0 means no collection.
	ExampleSamplePercentage float64 `protobuf:"fixed64,11,opt,name=example_sample_percentage,json=exampleSamplePercentage,proto3" json:"example_sample_percentage,omitempty"`
	// Alert config for the evaluation job. The alert will be triggered when its
	// criteria is met.
	EvaluationJobAlertConfig *EvaluationJobAlertConfig `protobuf:"bytes,13,opt,name=evaluation_job_alert_config,json=evaluationJobAlertConfig,proto3" json:"evaluation_job_alert_config,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *EvaluationJobConfig) Reset()         { *m = EvaluationJobConfig{} }
func (m *EvaluationJobConfig) String() string { return proto.CompactTextString(m) }
func (*EvaluationJobConfig) ProtoMessage()    {}
func (*EvaluationJobConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_052df044fb4ef883, []int{1}
}

func (m *EvaluationJobConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluationJobConfig.Unmarshal(m, b)
}
func (m *EvaluationJobConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluationJobConfig.Marshal(b, m, deterministic)
}
func (m *EvaluationJobConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluationJobConfig.Merge(m, src)
}
func (m *EvaluationJobConfig) XXX_Size() int {
	return xxx_messageInfo_EvaluationJobConfig.Size(m)
}
func (m *EvaluationJobConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluationJobConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluationJobConfig proto.InternalMessageInfo

type isEvaluationJobConfig_HumanAnnotationRequestConfig interface {
	isEvaluationJobConfig_HumanAnnotationRequestConfig()
}

type EvaluationJobConfig_ImageClassificationConfig struct {
	ImageClassificationConfig *ImageClassificationConfig `protobuf:"bytes,4,opt,name=image_classification_config,json=imageClassificationConfig,proto3,oneof"`
}

type EvaluationJobConfig_BoundingPolyConfig struct {
	BoundingPolyConfig *BoundingPolyConfig `protobuf:"bytes,5,opt,name=bounding_poly_config,json=boundingPolyConfig,proto3,oneof"`
}

type EvaluationJobConfig_VideoClassificationConfig struct {
	VideoClassificationConfig *VideoClassificationConfig `protobuf:"bytes,6,opt,name=video_classification_config,json=videoClassificationConfig,proto3,oneof"`
}

type EvaluationJobConfig_ObjectDetectionConfig struct {
	ObjectDetectionConfig *ObjectDetectionConfig `protobuf:"bytes,7,opt,name=object_detection_config,json=objectDetectionConfig,proto3,oneof"`
}

type EvaluationJobConfig_TextClassificationConfig struct {
	TextClassificationConfig *TextClassificationConfig `protobuf:"bytes,8,opt,name=text_classification_config,json=textClassificationConfig,proto3,oneof"`
}

type EvaluationJobConfig_ObjectTrackingConfig struct {
	ObjectTrackingConfig *ObjectTrackingConfig `protobuf:"bytes,12,opt,name=object_tracking_config,json=objectTrackingConfig,proto3,oneof"`
}

func (*EvaluationJobConfig_ImageClassificationConfig) isEvaluationJobConfig_HumanAnnotationRequestConfig() {
}

func (*EvaluationJobConfig_BoundingPolyConfig) isEvaluationJobConfig_HumanAnnotationRequestConfig() {}

func (*EvaluationJobConfig_VideoClassificationConfig) isEvaluationJobConfig_HumanAnnotationRequestConfig() {
}

func (*EvaluationJobConfig_ObjectDetectionConfig) isEvaluationJobConfig_HumanAnnotationRequestConfig() {
}

func (*EvaluationJobConfig_TextClassificationConfig) isEvaluationJobConfig_HumanAnnotationRequestConfig() {
}

func (*EvaluationJobConfig_ObjectTrackingConfig) isEvaluationJobConfig_HumanAnnotationRequestConfig() {
}

func (m *EvaluationJobConfig) GetHumanAnnotationRequestConfig() isEvaluationJobConfig_HumanAnnotationRequestConfig {
	if m != nil {
		return m.HumanAnnotationRequestConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetImageClassificationConfig() *ImageClassificationConfig {
	if x, ok := m.GetHumanAnnotationRequestConfig().(*EvaluationJobConfig_ImageClassificationConfig); ok {
		return x.ImageClassificationConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetBoundingPolyConfig() *BoundingPolyConfig {
	if x, ok := m.GetHumanAnnotationRequestConfig().(*EvaluationJobConfig_BoundingPolyConfig); ok {
		return x.BoundingPolyConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetVideoClassificationConfig() *VideoClassificationConfig {
	if x, ok := m.GetHumanAnnotationRequestConfig().(*EvaluationJobConfig_VideoClassificationConfig); ok {
		return x.VideoClassificationConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetObjectDetectionConfig() *ObjectDetectionConfig {
	if x, ok := m.GetHumanAnnotationRequestConfig().(*EvaluationJobConfig_ObjectDetectionConfig); ok {
		return x.ObjectDetectionConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetTextClassificationConfig() *TextClassificationConfig {
	if x, ok := m.GetHumanAnnotationRequestConfig().(*EvaluationJobConfig_TextClassificationConfig); ok {
		return x.TextClassificationConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetObjectTrackingConfig() *ObjectTrackingConfig {
	if x, ok := m.GetHumanAnnotationRequestConfig().(*EvaluationJobConfig_ObjectTrackingConfig); ok {
		return x.ObjectTrackingConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetInputConfig() *InputConfig {
	if m != nil {
		return m.InputConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetEvaluationConfig() *EvaluationConfig {
	if m != nil {
		return m.EvaluationConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetHumanAnnotationConfig() *HumanAnnotationConfig {
	if m != nil {
		return m.HumanAnnotationConfig
	}
	return nil
}

func (m *EvaluationJobConfig) GetBigqueryImportKeys() map[string]string {
	if m != nil {
		return m.BigqueryImportKeys
	}
	return nil
}

func (m *EvaluationJobConfig) GetExampleCount() int32 {
	if m != nil {
		return m.ExampleCount
	}
	return 0
}

func (m *EvaluationJobConfig) GetExampleSamplePercentage() float64 {
	if m != nil {
		return m.ExampleSamplePercentage
	}
	return 0
}

func (m *EvaluationJobConfig) GetEvaluationJobAlertConfig() *EvaluationJobAlertConfig {
	if m != nil {
		return m.EvaluationJobAlertConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EvaluationJobConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EvaluationJobConfig_ImageClassificationConfig)(nil),
		(*EvaluationJobConfig_BoundingPolyConfig)(nil),
		(*EvaluationJobConfig_VideoClassificationConfig)(nil),
		(*EvaluationJobConfig_ObjectDetectionConfig)(nil),
		(*EvaluationJobConfig_TextClassificationConfig)(nil),
		(*EvaluationJobConfig_ObjectTrackingConfig)(nil),
	}
}

type EvaluationJobAlertConfig struct {
	// Required. Email of the user who will be receiving the alert.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// If a single evaluation run's aggregate mean average precision is
	// lower than this threshold, the alert will be triggered.
	MinAcceptableMeanAveragePrecision float64  `protobuf:"fixed64,2,opt,name=min_acceptable_mean_average_precision,json=minAcceptableMeanAveragePrecision,proto3" json:"min_acceptable_mean_average_precision,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *EvaluationJobAlertConfig) Reset()         { *m = EvaluationJobAlertConfig{} }
func (m *EvaluationJobAlertConfig) String() string { return proto.CompactTextString(m) }
func (*EvaluationJobAlertConfig) ProtoMessage()    {}
func (*EvaluationJobAlertConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_052df044fb4ef883, []int{2}
}

func (m *EvaluationJobAlertConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluationJobAlertConfig.Unmarshal(m, b)
}
func (m *EvaluationJobAlertConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluationJobAlertConfig.Marshal(b, m, deterministic)
}
func (m *EvaluationJobAlertConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluationJobAlertConfig.Merge(m, src)
}
func (m *EvaluationJobAlertConfig) XXX_Size() int {
	return xxx_messageInfo_EvaluationJobAlertConfig.Size(m)
}
func (m *EvaluationJobAlertConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluationJobAlertConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluationJobAlertConfig proto.InternalMessageInfo

func (m *EvaluationJobAlertConfig) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EvaluationJobAlertConfig) GetMinAcceptableMeanAveragePrecision() float64 {
	if m != nil {
		return m.MinAcceptableMeanAveragePrecision
	}
	return 0
}

// Records a failed attempt.
type Attempt struct {
	AttemptTime          *timestamp.Timestamp `protobuf:"bytes,1,opt,name=attempt_time,json=attemptTime,proto3" json:"attempt_time,omitempty"`
	PartialFailures      []*status.Status     `protobuf:"bytes,2,rep,name=partial_failures,json=partialFailures,proto3" json:"partial_failures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Attempt) Reset()         { *m = Attempt{} }
func (m *Attempt) String() string { return proto.CompactTextString(m) }
func (*Attempt) ProtoMessage()    {}
func (*Attempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_052df044fb4ef883, []int{3}
}

func (m *Attempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attempt.Unmarshal(m, b)
}
func (m *Attempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attempt.Marshal(b, m, deterministic)
}
func (m *Attempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attempt.Merge(m, src)
}
func (m *Attempt) XXX_Size() int {
	return xxx_messageInfo_Attempt.Size(m)
}
func (m *Attempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Attempt.DiscardUnknown(m)
}

var xxx_messageInfo_Attempt proto.InternalMessageInfo

func (m *Attempt) GetAttemptTime() *timestamp.Timestamp {
	if m != nil {
		return m.AttemptTime
	}
	return nil
}

func (m *Attempt) GetPartialFailures() []*status.Status {
	if m != nil {
		return m.PartialFailures
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.datalabeling.v1beta1.EvaluationJob_State", EvaluationJob_State_name, EvaluationJob_State_value)
	proto.RegisterType((*EvaluationJob)(nil), "google.cloud.datalabeling.v1beta1.EvaluationJob")
	proto.RegisterType((*EvaluationJobConfig)(nil), "google.cloud.datalabeling.v1beta1.EvaluationJobConfig")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.datalabeling.v1beta1.EvaluationJobConfig.BigqueryImportKeysEntry")
	proto.RegisterType((*EvaluationJobAlertConfig)(nil), "google.cloud.datalabeling.v1beta1.EvaluationJobAlertConfig")
	proto.RegisterType((*Attempt)(nil), "google.cloud.datalabeling.v1beta1.Attempt")
}

func init() {
	proto.RegisterFile("google/cloud/datalabeling/v1beta1/evaluation_job.proto", fileDescriptor_052df044fb4ef883)
}

var fileDescriptor_052df044fb4ef883 = []byte{
	// 1073 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdf, 0x53, 0xdb, 0x46,
	0x10, 0x8e, 0x00, 0xf3, 0xe3, 0x0c, 0xad, 0x39, 0xa0, 0x56, 0x9c, 0xce, 0xc4, 0xd0, 0xc9, 0x8c,
	0xa7, 0x0f, 0xf2, 0x84, 0x4c, 0xd3, 0x4c, 0x68, 0xa6, 0x63, 0xb0, 0x09, 0xb4, 0x81, 0xb8, 0xb2,
	0xc9, 0x43, 0x5f, 0xd4, 0x93, 0xbc, 0x88, 0x03, 0x49, 0xa7, 0xe8, 0x4e, 0x1e, 0xdc, 0xce, 0x74,
	0x3a, 0x7d, 0xc9, 0x1f, 0xd8, 0x7f, 0xa8, 0x73, 0xa7, 0x93, 0x31, 0x01, 0xd5, 0xa6, 0x2f, 0xa0,
	0xdb, 0xfd, 0xf6, 0xfb, 0x6e, 0xf7, 0xd6, 0xb7, 0x87, 0x5e, 0xfa, 0x8c, 0xf9, 0x01, 0x34, 0xbd,
	0x80, 0xa5, 0x83, 0xe6, 0x80, 0x08, 0x12, 0x10, 0x17, 0x02, 0x1a, 0xf9, 0xcd, 0xe1, 0x73, 0x17,
	0x04, 0x79, 0xde, 0x84, 0x21, 0x09, 0x52, 0x22, 0x28, 0x8b, 0x9c, 0x4b, 0xe6, 0x5a, 0x71, 0xc2,
	0x04, 0xc3, 0xdb, 0x59, 0x9c, 0xa5, 0xe2, 0xac, 0xc9, 0x38, 0x4b, 0xc7, 0xd5, 0xbe, 0xd6, 0xd4,
	0x24, 0xa6, 0x4d, 0x12, 0x45, 0x4c, 0x28, 0x0e, 0x9e, 0x11, 0xd4, 0x9a, 0xd3, 0x85, 0xa5, 0x91,
	0x83, 0xd0, 0x01, 0xbb, 0x0f, 0xd9, 0xa9, 0x8e, 0xf9, 0x71, 0x7a, 0xcc, 0x45, 0x1a, 0x92, 0xc8,
	0xb9, 0xd9, 0x9f, 0xe3, 0xb1, 0xe8, 0x9c, 0xfa, 0x9a, 0xe0, 0xa9, 0x26, 0x50, 0x2b, 0x37, 0x3d,
	0x6f, 0x0a, 0x1a, 0x02, 0x17, 0x24, 0x8c, 0x35, 0xa0, 0xaa, 0x01, 0x49, 0xec, 0x35, 0xb9, 0x20,
	0x22, 0xd5, 0xf9, 0xed, 0x7c, 0x2a, 0xa1, 0xb5, 0xce, 0x78, 0x3f, 0x3f, 0x31, 0x17, 0x63, 0xb4,
	0x10, 0x91, 0x10, 0x4c, 0xa3, 0x6e, 0x34, 0x56, 0x6c, 0xf5, 0x8d, 0xeb, 0xa8, 0x3c, 0x00, 0xee,
	0x25, 0x34, 0x96, 0x28, 0x73, 0x4e, 0xb9, 0x26, 0x4d, 0xf8, 0x1d, 0x2a, 0x49, 0x5e, 0x30, 0xe7,
	0xeb, 0x46, 0xe3, 0x8b, 0xdd, 0x97, 0xd6, 0xd4, 0xc2, 0x5b, 0xb7, 0x64, 0xad, 0x9e, 0x8c, 0xb6,
	0x33, 0x12, 0x5c, 0x43, 0xcb, 0xdc, 0xbb, 0x80, 0x41, 0x1a, 0x80, 0xb9, 0xa0, 0xc4, 0xc6, 0x6b,
	0xfc, 0x0d, 0x5a, 0x0b, 0xd9, 0x00, 0x02, 0x67, 0x08, 0x09, 0x97, 0xbb, 0x29, 0x29, 0xc0, 0xaa,
	0x32, 0x7e, 0xc8, 0x6c, 0xf8, 0x12, 0x6d, 0xdd, 0xee, 0x07, 0x5d, 0x2f, 0x73, 0xb1, 0x6e, 0x34,
	0xca, 0x0f, 0xdf, 0xde, 0x81, 0x8a, 0xb6, 0x37, 0xe0, 0xae, 0x11, 0x5b, 0x68, 0x63, 0xe2, 0x5c,
	0x78, 0x0c, 0x9e, 0xc3, 0x41, 0x98, 0x4b, 0x6a, 0x5b, 0xeb, 0x37, 0xae, 0x5e, 0x0c, 0x5e, 0x0f,
	0x04, 0xde, 0x43, 0x35, 0x25, 0xe6, 0x84, 0x94, 0x73, 0x1a, 0xf9, 0x8e, 0x9f, 0xb0, 0x34, 0x1a,
	0x38, 0x22, 0x49, 0xc5, 0x85, 0xb9, 0x5c, 0x37, 0x1a, 0xcb, 0x76, 0x55, 0x21, 0x4e, 0x32, 0xc0,
	0x5b, 0xe5, 0xef, 0x4b, 0x37, 0x3e, 0x44, 0xcb, 0x44, 0x08, 0x08, 0x63, 0xc1, 0xcd, 0x95, 0xfa,
	0x7c, 0xa3, 0xbc, 0xfb, 0xed, 0x0c, 0xb9, 0xb4, 0xb2, 0x10, 0x7b, 0x1c, 0x8b, 0xf7, 0x50, 0xd9,
	0x4b, 0x80, 0x08, 0x70, 0x64, 0xab, 0x98, 0x48, 0x95, 0xa5, 0x96, 0x53, 0xe5, 0x7d, 0x64, 0xf5,
	0xf3, 0x3e, 0xb2, 0x51, 0x06, 0x97, 0x86, 0x9d, 0x1e, 0x2a, 0xa9, 0xe3, 0xc2, 0x5b, 0x68, 0xbd,
	0xd7, 0x6f, 0xf5, 0x3b, 0xce, 0xd9, 0x69, 0xaf, 0xdb, 0x39, 0x38, 0x3e, 0x3c, 0xee, 0xb4, 0x2b,
	0x8f, 0xf0, 0x1a, 0x5a, 0xe9, 0x1d, 0x1c, 0x75, 0xda, 0x67, 0xef, 0x3a, 0xed, 0x8a, 0x81, 0xcb,
	0x68, 0xc9, 0x3e, 0x3b, 0x3d, 0x3d, 0x3e, 0x7d, 0x5b, 0x99, 0xc3, 0x08, 0x2d, 0x76, 0x5b, 0x67,
	0xbd, 0x4e, 0xbb, 0x32, 0x2f, 0x1d, 0xbd, 0xfe, 0xfb, 0x6e, 0xb7, 0xd3, 0xae, 0x2c, 0xec, 0xfc,
	0x53, 0x46, 0x1b, 0xf7, 0xd4, 0x1c, 0xff, 0x89, 0x9e, 0xd0, 0x90, 0xf8, 0xe0, 0x78, 0x01, 0xe1,
	0x9c, 0x9e, 0x53, 0x6f, 0xf2, 0x07, 0xa0, 0xda, 0xa3, 0xbc, 0xfb, 0xc3, 0x0c, 0x45, 0x38, 0x96,
	0x2c, 0x07, 0xb7, 0x48, 0x32, 0x89, 0xa3, 0x47, 0xf6, 0x63, 0x5a, 0xe4, 0xc4, 0x14, 0x6d, 0xba,
	0xb2, 0xfe, 0xf2, 0xa4, 0x62, 0x16, 0x8c, 0x72, 0xe1, 0x92, 0x12, 0xfe, 0x6e, 0x06, 0xe1, 0x7d,
	0x1d, 0xde, 0x65, 0xc1, 0x68, 0xac, 0x88, 0xdd, 0x3b, 0x56, 0x99, 0xea, 0x90, 0x0e, 0x80, 0x15,
	0xa4, 0xba, 0x38, 0x73, 0xaa, 0x1f, 0x24, 0x4b, 0x51, 0xaa, 0xc3, 0x22, 0x27, 0x4e, 0x50, 0x95,
	0xb9, 0x97, 0xe0, 0x09, 0x67, 0x00, 0x02, 0xbc, 0x49, 0xed, 0x25, 0xa5, 0xfd, 0x6a, 0x06, 0xed,
	0xf7, 0x8a, 0xa1, 0x9d, 0x13, 0x8c, 0x75, 0xb7, 0xd8, 0x7d, 0x0e, 0xfc, 0x07, 0xaa, 0x09, 0xb8,
	0x16, 0x05, 0x29, 0x2f, 0x2b, 0xd9, 0xbd, 0x19, 0x64, 0xfb, 0x70, 0x2d, 0x0a, 0x32, 0x36, 0x45,
	0x81, 0x0f, 0x33, 0xf4, 0x95, 0x4e, 0x58, 0x24, 0xc4, 0xbb, 0x92, 0x47, 0xac, 0x85, 0x57, 0x95,
	0xf0, 0xf7, 0x33, 0xe7, 0xdb, 0xd7, 0xf1, 0x63, 0xd1, 0x4d, 0x76, 0x8f, 0x1d, 0xff, 0x82, 0x56,
	0x69, 0x14, 0xa7, 0x22, 0x97, 0x31, 0x94, 0x8c, 0x35, 0x4b, 0xf7, 0xca, 0x30, 0x7d, 0x0d, 0x95,
	0xe9, 0xcd, 0x02, 0xff, 0x86, 0xd6, 0x27, 0xae, 0x3a, 0xcd, 0x3b, 0xa7, 0x78, 0x5f, 0x3c, 0xe8,
	0x9a, 0xd3, 0xe4, 0x15, 0xf8, 0xcc, 0x82, 0x63, 0x54, 0x2d, 0x18, 0x3f, 0xea, 0xb6, 0x9f, 0xad,
	0x2d, 0x8e, 0x24, 0x43, 0x6b, 0x4c, 0xa0, 0xc5, 0xb6, 0x2e, 0xee, 0x33, 0xe3, 0xbf, 0x0c, 0xb4,
	0xe9, 0x52, 0xff, 0x63, 0x0a, 0xc9, 0xc8, 0xa1, 0x61, 0xcc, 0x12, 0xe1, 0x5c, 0xc1, 0x28, 0xbf,
	0xf2, 0x4e, 0xff, 0xdf, 0xf5, 0x6d, 0xed, 0x6b, 0xca, 0x63, 0xc5, 0xf8, 0x33, 0x8c, 0x78, 0x27,
	0x12, 0xc9, 0xc8, 0xc6, 0xee, 0x1d, 0x87, 0x1c, 0x33, 0x70, 0x4d, 0xc2, 0x38, 0x00, 0xc7, 0x63,
	0x69, 0x24, 0xd4, 0x15, 0x59, 0xb2, 0x57, 0xb5, 0xf1, 0x40, 0xda, 0xf0, 0x6b, 0xf4, 0x38, 0x07,
	0xf1, 0xec, 0x5f, 0x0c, 0x89, 0x07, 0x91, 0x20, 0x3e, 0x98, 0xe5, 0xba, 0xd1, 0x30, 0xec, 0xaa,
	0x06, 0xf4, 0xd4, 0xdf, 0xee, 0xd8, 0x8d, 0x7f, 0x47, 0x4f, 0x3e, 0x1b, 0x51, 0x24, 0x80, 0x64,
	0xdc, 0x19, 0x6b, 0x33, 0x77, 0xfe, 0xad, 0x4c, 0x5b, 0x92, 0x43, 0x17, 0xd7, 0x84, 0x02, 0x4f,
	0xad, 0x83, 0xaa, 0x05, 0xb5, 0xc0, 0x15, 0x34, 0x7f, 0x05, 0x23, 0x3d, 0xfd, 0xe5, 0x27, 0xde,
	0x44, 0x25, 0xc9, 0x03, 0x7a, 0xec, 0x67, 0x8b, 0xd7, 0x73, 0xaf, 0x8c, 0xfd, 0x6d, 0xf4, 0xf4,
	0x4e, 0x63, 0x24, 0xf0, 0x31, 0x05, 0x9e, 0xa7, 0xb1, 0xf3, 0xb7, 0x81, 0xcc, 0xa2, 0x0d, 0x4a,
	0x66, 0x08, 0x09, 0x0d, 0xb4, 0x5a, 0xb6, 0xc0, 0x5d, 0xf4, 0x2c, 0xa4, 0x91, 0x43, 0x3c, 0x0f,
	0x62, 0x41, 0xdc, 0x00, 0x9c, 0x10, 0xa4, 0xc6, 0x10, 0x12, 0x39, 0x06, 0xe2, 0x04, 0x3c, 0xca,
	0xf3, 0x67, 0x88, 0x61, 0x6f, 0x87, 0x34, 0x6a, 0x8d, 0xb1, 0x27, 0x40, 0xa2, 0x56, 0x86, 0xec,
	0xe6, 0xc0, 0x9d, 0x4f, 0x06, 0x5a, 0xd2, 0x23, 0x10, 0xbf, 0x41, 0xab, 0x7a, 0x08, 0x66, 0x93,
	0xcf, 0x98, 0x3a, 0xf9, 0xca, 0x1a, 0x2f, 0x2d, 0xf8, 0x0d, 0xaa, 0xc4, 0x24, 0x11, 0x94, 0x04,
	0xce, 0x39, 0xa1, 0x41, 0x9a, 0x00, 0x37, 0xe7, 0x54, 0x53, 0xe2, 0x9c, 0x22, 0x89, 0x3d, 0xf5,
	0x9a, 0x49, 0xb9, 0xfd, 0xa5, 0xc6, 0x1e, 0x6a, 0xe8, 0xfe, 0x35, 0x7a, 0xe6, 0xb1, 0x70, 0xfa,
	0xa1, 0x76, 0x8d, 0x5f, 0x4f, 0x34, 0xc8, 0x67, 0x01, 0x89, 0x7c, 0x8b, 0x25, 0x7e, 0xd3, 0x87,
	0x48, 0xed, 0x4f, 0x3f, 0x4a, 0x49, 0x4c, 0xf9, 0x7f, 0xbc, 0x19, 0xf7, 0x26, 0x8d, 0xee, 0xa2,
	0x8a, 0x7c, 0xf1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xe1, 0xc5, 0x2b, 0x4a, 0x0b, 0x00,
	0x00,
}
