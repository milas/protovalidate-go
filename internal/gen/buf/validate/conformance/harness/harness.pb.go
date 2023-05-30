// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: buf/validate/conformance/harness/harness.proto

package harness

import (
	validate "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestConformanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fdset *descriptorpb.FileDescriptorSet `protobuf:"bytes,2,opt,name=fdset,proto3" json:"fdset,omitempty"`
	Cases map[string]*anypb.Any           `protobuf:"bytes,3,rep,name=cases,proto3" json:"cases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestConformanceRequest) Reset() {
	*x = TestConformanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_harness_harness_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConformanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConformanceRequest) ProtoMessage() {}

func (x *TestConformanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_harness_harness_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConformanceRequest.ProtoReflect.Descriptor instead.
func (*TestConformanceRequest) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_harness_harness_proto_rawDescGZIP(), []int{0}
}

func (x *TestConformanceRequest) GetFdset() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.Fdset
	}
	return nil
}

func (x *TestConformanceRequest) GetCases() map[string]*anypb.Any {
	if x != nil {
		return x.Cases
	}
	return nil
}

type TestConformanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results map[string]*TestResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestConformanceResponse) Reset() {
	*x = TestConformanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_harness_harness_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConformanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConformanceResponse) ProtoMessage() {}

func (x *TestConformanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_harness_harness_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConformanceResponse.ProtoReflect.Descriptor instead.
func (*TestConformanceResponse) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_harness_harness_proto_rawDescGZIP(), []int{1}
}

func (x *TestConformanceResponse) GetResults() map[string]*TestResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type TestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*TestResult_Success
	//	*TestResult_ValidationError
	//	*TestResult_CompilationError
	//	*TestResult_RuntimeError
	//	*TestResult_UnexpectedError
	Result isTestResult_Result `protobuf_oneof:"result"`
}

func (x *TestResult) Reset() {
	*x = TestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_harness_harness_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResult) ProtoMessage() {}

func (x *TestResult) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_harness_harness_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResult.ProtoReflect.Descriptor instead.
func (*TestResult) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_harness_harness_proto_rawDescGZIP(), []int{2}
}

func (m *TestResult) GetResult() isTestResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *TestResult) GetSuccess() bool {
	if x, ok := x.GetResult().(*TestResult_Success); ok {
		return x.Success
	}
	return false
}

func (x *TestResult) GetValidationError() *validate.Violations {
	if x, ok := x.GetResult().(*TestResult_ValidationError); ok {
		return x.ValidationError
	}
	return nil
}

func (x *TestResult) GetCompilationError() string {
	if x, ok := x.GetResult().(*TestResult_CompilationError); ok {
		return x.CompilationError
	}
	return ""
}

func (x *TestResult) GetRuntimeError() string {
	if x, ok := x.GetResult().(*TestResult_RuntimeError); ok {
		return x.RuntimeError
	}
	return ""
}

func (x *TestResult) GetUnexpectedError() string {
	if x, ok := x.GetResult().(*TestResult_UnexpectedError); ok {
		return x.UnexpectedError
	}
	return ""
}

type isTestResult_Result interface {
	isTestResult_Result()
}

type TestResult_Success struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3,oneof"`
}

type TestResult_ValidationError struct {
	ValidationError *validate.Violations `protobuf:"bytes,2,opt,name=validation_error,json=validationError,proto3,oneof"`
}

type TestResult_CompilationError struct {
	CompilationError string `protobuf:"bytes,3,opt,name=compilation_error,json=compilationError,proto3,oneof"`
}

type TestResult_RuntimeError struct {
	RuntimeError string `protobuf:"bytes,4,opt,name=runtime_error,json=runtimeError,proto3,oneof"`
}

type TestResult_UnexpectedError struct {
	UnexpectedError string `protobuf:"bytes,5,opt,name=unexpected_error,json=unexpectedError,proto3,oneof"`
}

func (*TestResult_Success) isTestResult_Result() {}

func (*TestResult_ValidationError) isTestResult_Result() {}

func (*TestResult_CompilationError) isTestResult_Result() {}

func (*TestResult_RuntimeError) isTestResult_Result() {}

func (*TestResult_UnexpectedError) isTestResult_Result() {}

var File_buf_validate_conformance_harness_harness_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_harness_harness_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x20, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x1a, 0x1d, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd,
	0x01, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x66, 0x64, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x05, 0x66, 0x64,
	0x73, 0x65, 0x74, 0x12, 0x59, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61,
	0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x73,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x4e,
	0x0a, 0x0a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe5,
	0x01, 0x0a, 0x17, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x68, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x42,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfc, 0x01, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x45, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0d, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0c, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b,
	0x0a, 0x10, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x75, 0x6e, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0xac, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x0c,
	0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x52,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0xa2, 0x02, 0x04, 0x42, 0x56, 0x43, 0x48, 0xaa, 0x02, 0x20, 0x42, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0xca, 0x02, 0x20, 0x42,
	0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x48, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0xe2,
	0x02, 0x2c, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x48, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x23, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x3a,
	0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x48, 0x61, 0x72,
	0x6e, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_harness_harness_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_harness_harness_proto_rawDescData = file_buf_validate_conformance_harness_harness_proto_rawDesc
)

func file_buf_validate_conformance_harness_harness_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_harness_harness_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_harness_harness_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_harness_harness_proto_rawDescData)
	})
	return file_buf_validate_conformance_harness_harness_proto_rawDescData
}

var file_buf_validate_conformance_harness_harness_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_buf_validate_conformance_harness_harness_proto_goTypes = []interface{}{
	(*TestConformanceRequest)(nil),         // 0: buf.validate.conformance.harness.TestConformanceRequest
	(*TestConformanceResponse)(nil),        // 1: buf.validate.conformance.harness.TestConformanceResponse
	(*TestResult)(nil),                     // 2: buf.validate.conformance.harness.TestResult
	nil,                                    // 3: buf.validate.conformance.harness.TestConformanceRequest.CasesEntry
	nil,                                    // 4: buf.validate.conformance.harness.TestConformanceResponse.ResultsEntry
	(*descriptorpb.FileDescriptorSet)(nil), // 5: google.protobuf.FileDescriptorSet
	(*validate.Violations)(nil),            // 6: buf.validate.Violations
	(*anypb.Any)(nil),                      // 7: google.protobuf.Any
}
var file_buf_validate_conformance_harness_harness_proto_depIdxs = []int32{
	5, // 0: buf.validate.conformance.harness.TestConformanceRequest.fdset:type_name -> google.protobuf.FileDescriptorSet
	3, // 1: buf.validate.conformance.harness.TestConformanceRequest.cases:type_name -> buf.validate.conformance.harness.TestConformanceRequest.CasesEntry
	4, // 2: buf.validate.conformance.harness.TestConformanceResponse.results:type_name -> buf.validate.conformance.harness.TestConformanceResponse.ResultsEntry
	6, // 3: buf.validate.conformance.harness.TestResult.validation_error:type_name -> buf.validate.Violations
	7, // 4: buf.validate.conformance.harness.TestConformanceRequest.CasesEntry.value:type_name -> google.protobuf.Any
	2, // 5: buf.validate.conformance.harness.TestConformanceResponse.ResultsEntry.value:type_name -> buf.validate.conformance.harness.TestResult
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_harness_harness_proto_init() }
func file_buf_validate_conformance_harness_harness_proto_init() {
	if File_buf_validate_conformance_harness_harness_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_harness_harness_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConformanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_validate_conformance_harness_harness_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConformanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_validate_conformance_harness_harness_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_buf_validate_conformance_harness_harness_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TestResult_Success)(nil),
		(*TestResult_ValidationError)(nil),
		(*TestResult_CompilationError)(nil),
		(*TestResult_RuntimeError)(nil),
		(*TestResult_UnexpectedError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_harness_harness_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_harness_harness_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_harness_harness_proto_depIdxs,
		MessageInfos:      file_buf_validate_conformance_harness_harness_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_harness_harness_proto = out.File
	file_buf_validate_conformance_harness_harness_proto_rawDesc = nil
	file_buf_validate_conformance_harness_harness_proto_goTypes = nil
	file_buf_validate_conformance_harness_harness_proto_depIdxs = nil
}
