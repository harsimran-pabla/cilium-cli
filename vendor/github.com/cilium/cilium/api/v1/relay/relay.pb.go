// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: relay/relay.proto

package relay

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeState int32

const (
	// UNKNOWN_NODE_STATE indicates that the state of this node is unknown.
	NodeState_UNKNOWN_NODE_STATE NodeState = 0
	// NODE_CONNECTED indicates that we have established a connection
	// to this node. The client can expect to observe flows from this node.
	NodeState_NODE_CONNECTED NodeState = 1
	// NODE_UNAVAILABLE indicates that the connection to this
	// node is currently unavailable. The client can expect to not see any
	// flows from this node until either the connection is re-established or
	// the node is gone.
	NodeState_NODE_UNAVAILABLE NodeState = 2
	// NODE_GONE indicates that a node has been removed from the
	// cluster. No reconnection attempts will be made.
	NodeState_NODE_GONE NodeState = 3
	// NODE_ERROR indicates that a node has reported an error while processing
	// the request. No reconnection attempts will be made.
	NodeState_NODE_ERROR NodeState = 4
)

// Enum value maps for NodeState.
var (
	NodeState_name = map[int32]string{
		0: "UNKNOWN_NODE_STATE",
		1: "NODE_CONNECTED",
		2: "NODE_UNAVAILABLE",
		3: "NODE_GONE",
		4: "NODE_ERROR",
	}
	NodeState_value = map[string]int32{
		"UNKNOWN_NODE_STATE": 0,
		"NODE_CONNECTED":     1,
		"NODE_UNAVAILABLE":   2,
		"NODE_GONE":          3,
		"NODE_ERROR":         4,
	}
)

func (x NodeState) Enum() *NodeState {
	p := new(NodeState)
	*p = x
	return p
}

func (x NodeState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeState) Descriptor() protoreflect.EnumDescriptor {
	return file_relay_relay_proto_enumTypes[0].Descriptor()
}

func (NodeState) Type() protoreflect.EnumType {
	return &file_relay_relay_proto_enumTypes[0]
}

func (x NodeState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeState.Descriptor instead.
func (NodeState) EnumDescriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{0}
}

// NodeStatusEvent is a message sent by hubble-relay to inform clients about
// the state of a particular node.
type NodeStatusEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// state_change contains the new node state
	StateChange NodeState `protobuf:"varint,1,opt,name=state_change,json=stateChange,proto3,enum=relay.NodeState" json:"state_change,omitempty"`
	// node_names is the list of nodes for which the above state changes applies
	NodeNames []string `protobuf:"bytes,2,rep,name=node_names,json=nodeNames,proto3" json:"node_names,omitempty"`
	// message is an optional message attached to the state change (e.g. an
	// error message). The message applies to all nodes in node_names.
	Message       string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeStatusEvent) Reset() {
	*x = NodeStatusEvent{}
	mi := &file_relay_relay_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeStatusEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatusEvent) ProtoMessage() {}

func (x *NodeStatusEvent) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatusEvent.ProtoReflect.Descriptor instead.
func (*NodeStatusEvent) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStatusEvent) GetStateChange() NodeState {
	if x != nil {
		return x.StateChange
	}
	return NodeState_UNKNOWN_NODE_STATE
}

func (x *NodeStatusEvent) GetNodeNames() []string {
	if x != nil {
		return x.NodeNames
	}
	return nil
}

func (x *NodeStatusEvent) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_relay_relay_proto protoreflect.FileDescriptor

var file_relay_relay_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x7f, 0x0a, 0x0f, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x6c, 0x0a, 0x09, 0x4e,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f,
	0x44, 0x45, 0x5f, 0x47, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x44,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x63,
	0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_relay_relay_proto_rawDescOnce sync.Once
	file_relay_relay_proto_rawDescData []byte
)

func file_relay_relay_proto_rawDescGZIP() []byte {
	file_relay_relay_proto_rawDescOnce.Do(func() {
		file_relay_relay_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_relay_relay_proto_rawDesc), len(file_relay_relay_proto_rawDesc)))
	})
	return file_relay_relay_proto_rawDescData
}

var file_relay_relay_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relay_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_relay_relay_proto_goTypes = []any{
	(NodeState)(0),          // 0: relay.NodeState
	(*NodeStatusEvent)(nil), // 1: relay.NodeStatusEvent
}
var file_relay_relay_proto_depIdxs = []int32{
	0, // 0: relay.NodeStatusEvent.state_change:type_name -> relay.NodeState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relay_relay_proto_init() }
func file_relay_relay_proto_init() {
	if File_relay_relay_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_relay_relay_proto_rawDesc), len(file_relay_relay_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relay_relay_proto_goTypes,
		DependencyIndexes: file_relay_relay_proto_depIdxs,
		EnumInfos:         file_relay_relay_proto_enumTypes,
		MessageInfos:      file_relay_relay_proto_msgTypes,
	}.Build()
	File_relay_relay_proto = out.File
	file_relay_relay_proto_goTypes = nil
	file_relay_relay_proto_depIdxs = nil
}
