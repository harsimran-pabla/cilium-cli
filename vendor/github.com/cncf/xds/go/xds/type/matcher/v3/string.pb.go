// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.1
// source: xds/type/matcher/v3/string.proto

package v3

import (
	v3 "github.com/cncf/xds/go/xds/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StringMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchPattern:
	//
	//	*StringMatcher_Exact
	//	*StringMatcher_Prefix
	//	*StringMatcher_Suffix
	//	*StringMatcher_SafeRegex
	//	*StringMatcher_Contains
	//	*StringMatcher_Custom
	MatchPattern isStringMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	IgnoreCase   bool                         `protobuf:"varint,6,opt,name=ignore_case,json=ignoreCase,proto3" json:"ignore_case,omitempty"`
}

func (x *StringMatcher) Reset() {
	*x = StringMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringMatcher) ProtoMessage() {}

func (x *StringMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringMatcher.ProtoReflect.Descriptor instead.
func (*StringMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_string_proto_rawDescGZIP(), []int{0}
}

func (m *StringMatcher) GetMatchPattern() isStringMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (x *StringMatcher) GetExact() string {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *StringMatcher) GetPrefix() string {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *StringMatcher) GetSuffix() string {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Suffix); ok {
		return x.Suffix
	}
	return ""
}

func (x *StringMatcher) GetSafeRegex() *RegexMatcher {
	if x, ok := x.GetMatchPattern().(*StringMatcher_SafeRegex); ok {
		return x.SafeRegex
	}
	return nil
}

func (x *StringMatcher) GetContains() string {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Contains); ok {
		return x.Contains
	}
	return ""
}

func (x *StringMatcher) GetCustom() *v3.TypedExtensionConfig {
	if x, ok := x.GetMatchPattern().(*StringMatcher_Custom); ok {
		return x.Custom
	}
	return nil
}

func (x *StringMatcher) GetIgnoreCase() bool {
	if x != nil {
		return x.IgnoreCase
	}
	return false
}

type isStringMatcher_MatchPattern interface {
	isStringMatcher_MatchPattern()
}

type StringMatcher_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type StringMatcher_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type StringMatcher_Suffix struct {
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3,oneof"`
}

type StringMatcher_SafeRegex struct {
	SafeRegex *RegexMatcher `protobuf:"bytes,5,opt,name=safe_regex,json=safeRegex,proto3,oneof"`
}

type StringMatcher_Contains struct {
	Contains string `protobuf:"bytes,7,opt,name=contains,proto3,oneof"`
}

type StringMatcher_Custom struct {
	Custom *v3.TypedExtensionConfig `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*StringMatcher_Exact) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Prefix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Suffix) isStringMatcher_MatchPattern() {}

func (*StringMatcher_SafeRegex) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Contains) isStringMatcher_MatchPattern() {}

func (*StringMatcher_Custom) isStringMatcher_MatchPattern() {}

type ListStringMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patterns []*StringMatcher `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
}

func (x *ListStringMatcher) Reset() {
	*x = ListStringMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_string_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStringMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStringMatcher) ProtoMessage() {}

func (x *ListStringMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_string_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStringMatcher.ProtoReflect.Descriptor instead.
func (*ListStringMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_string_proto_rawDescGZIP(), []int{1}
}

func (x *ListStringMatcher) GetPatterns() []*StringMatcher {
	if x != nil {
		return x.Patterns
	}
	return nil
}

var File_xds_type_matcher_v3_string_proto protoreflect.FileDescriptor

var file_xds_type_matcher_v3_string_proto_rawDesc = []byte{
	0x0a, 0x20, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1b, 0x78, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x02, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x21, 0x0a, 0x06, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x4c,
	0x0a, 0x0a, 0x73, 0x61, 0x66, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48,
	0x00, 0x52, 0x09, 0x73, 0x61, 0x66, 0x65, 0x52, 0x65, 0x67, 0x65, 0x78, 0x12, 0x25, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x73,
	0x65, 0x42, 0x14, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x5d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x08,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x70, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x42, 0x5b, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_type_matcher_v3_string_proto_rawDescOnce sync.Once
	file_xds_type_matcher_v3_string_proto_rawDescData = file_xds_type_matcher_v3_string_proto_rawDesc
)

func file_xds_type_matcher_v3_string_proto_rawDescGZIP() []byte {
	file_xds_type_matcher_v3_string_proto_rawDescOnce.Do(func() {
		file_xds_type_matcher_v3_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_type_matcher_v3_string_proto_rawDescData)
	})
	return file_xds_type_matcher_v3_string_proto_rawDescData
}

var file_xds_type_matcher_v3_string_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xds_type_matcher_v3_string_proto_goTypes = []interface{}{
	(*StringMatcher)(nil),           // 0: xds.type.matcher.v3.StringMatcher
	(*ListStringMatcher)(nil),       // 1: xds.type.matcher.v3.ListStringMatcher
	(*RegexMatcher)(nil),            // 2: xds.type.matcher.v3.RegexMatcher
	(*v3.TypedExtensionConfig)(nil), // 3: xds.core.v3.TypedExtensionConfig
}
var file_xds_type_matcher_v3_string_proto_depIdxs = []int32{
	2, // 0: xds.type.matcher.v3.StringMatcher.safe_regex:type_name -> xds.type.matcher.v3.RegexMatcher
	3, // 1: xds.type.matcher.v3.StringMatcher.custom:type_name -> xds.core.v3.TypedExtensionConfig
	0, // 2: xds.type.matcher.v3.ListStringMatcher.patterns:type_name -> xds.type.matcher.v3.StringMatcher
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_xds_type_matcher_v3_string_proto_init() }
func file_xds_type_matcher_v3_string_proto_init() {
	if File_xds_type_matcher_v3_string_proto != nil {
		return
	}
	file_xds_type_matcher_v3_regex_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_xds_type_matcher_v3_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringMatcher); i {
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
		file_xds_type_matcher_v3_string_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStringMatcher); i {
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
	file_xds_type_matcher_v3_string_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StringMatcher_Exact)(nil),
		(*StringMatcher_Prefix)(nil),
		(*StringMatcher_Suffix)(nil),
		(*StringMatcher_SafeRegex)(nil),
		(*StringMatcher_Contains)(nil),
		(*StringMatcher_Custom)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xds_type_matcher_v3_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_type_matcher_v3_string_proto_goTypes,
		DependencyIndexes: file_xds_type_matcher_v3_string_proto_depIdxs,
		MessageInfos:      file_xds_type_matcher_v3_string_proto_msgTypes,
	}.Build()
	File_xds_type_matcher_v3_string_proto = out.File
	file_xds_type_matcher_v3_string_proto_rawDesc = nil
	file_xds_type_matcher_v3_string_proto_goTypes = nil
	file_xds_type_matcher_v3_string_proto_depIdxs = nil
}
