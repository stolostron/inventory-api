// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: kessel/inventory/v1beta1/k8s_cluster.proto

package v1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type K8SCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata about this resource
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Write only reporter specific data
	ReporterData *ReporterData `protobuf:"bytes,245278792,opt,name=reporter_data,json=reporterData,proto3" json:"reporter_data,omitempty"`
	// The entities that registered this item in the Kessel Asset Inventory. The
	// same resource may be registered by multiple reporters.
	Reporters    []*ReporterData   `protobuf:"bytes,353323086,rep,name=reporters,proto3" json:"reporters,omitempty"`
	ResourceData *K8SClusterDetail `protobuf:"bytes,2122698,opt,name=resource_data,json=resourceData,proto3" json:"resource_data,omitempty"`
}

func (x *K8SCluster) Reset() {
	*x = K8SCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kessel_inventory_v1beta1_k8s_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SCluster) ProtoMessage() {}

func (x *K8SCluster) ProtoReflect() protoreflect.Message {
	mi := &file_kessel_inventory_v1beta1_k8s_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SCluster.ProtoReflect.Descriptor instead.
func (*K8SCluster) Descriptor() ([]byte, []int) {
	return file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *K8SCluster) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *K8SCluster) GetReporterData() *ReporterData {
	if x != nil {
		return x.ReporterData
	}
	return nil
}

func (x *K8SCluster) GetReporters() []*ReporterData {
	if x != nil {
		return x.Reporters
	}
	return nil
}

func (x *K8SCluster) GetResourceData() *K8SClusterDetail {
	if x != nil {
		return x.ResourceData
	}
	return nil
}

var File_kessel_inventory_v1beta1_k8s_cluster_proto protoreflect.FileDescriptor

var file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6b, 0x38, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6b, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc9, 0x02, 0x0a, 0x0a, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x53, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0xc8, 0xd0, 0xfa, 0x74, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x04, 0x52, 0x0c, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x09, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x18, 0xce, 0x90, 0xbd, 0xa8, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x73, 0x12, 0x52, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0xca, 0xc7, 0x81, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4b,
	0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x83, 0x01,
	0x0a, 0x28, 0x6f, 0x72, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6b, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0f, 0x4b, 0x38, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2d, 0x6b, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x73, 0x73, 0x65,
	0x6c, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescOnce sync.Once
	file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescData = file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDesc
)

func file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescGZIP() []byte {
	file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescOnce.Do(func() {
		file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescData)
	})
	return file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDescData
}

var file_kessel_inventory_v1beta1_k8s_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_kessel_inventory_v1beta1_k8s_cluster_proto_goTypes = []any{
	(*K8SCluster)(nil),       // 0: kessel.inventory.v1beta1.K8sCluster
	(*Metadata)(nil),         // 1: kessel.inventory.v1beta1.Metadata
	(*ReporterData)(nil),     // 2: kessel.inventory.v1beta1.ReporterData
	(*K8SClusterDetail)(nil), // 3: kessel.inventory.v1beta1.K8sClusterDetail
}
var file_kessel_inventory_v1beta1_k8s_cluster_proto_depIdxs = []int32{
	1, // 0: kessel.inventory.v1beta1.K8sCluster.metadata:type_name -> kessel.inventory.v1beta1.Metadata
	2, // 1: kessel.inventory.v1beta1.K8sCluster.reporter_data:type_name -> kessel.inventory.v1beta1.ReporterData
	2, // 2: kessel.inventory.v1beta1.K8sCluster.reporters:type_name -> kessel.inventory.v1beta1.ReporterData
	3, // 3: kessel.inventory.v1beta1.K8sCluster.resource_data:type_name -> kessel.inventory.v1beta1.K8sClusterDetail
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_kessel_inventory_v1beta1_k8s_cluster_proto_init() }
func file_kessel_inventory_v1beta1_k8s_cluster_proto_init() {
	if File_kessel_inventory_v1beta1_k8s_cluster_proto != nil {
		return
	}
	file_kessel_inventory_v1beta1_k8s_cluster_detail_proto_init()
	file_kessel_inventory_v1beta1_metadata_proto_init()
	file_kessel_inventory_v1beta1_reporter_data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_kessel_inventory_v1beta1_k8s_cluster_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*K8SCluster); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kessel_inventory_v1beta1_k8s_cluster_proto_goTypes,
		DependencyIndexes: file_kessel_inventory_v1beta1_k8s_cluster_proto_depIdxs,
		MessageInfos:      file_kessel_inventory_v1beta1_k8s_cluster_proto_msgTypes,
	}.Build()
	File_kessel_inventory_v1beta1_k8s_cluster_proto = out.File
	file_kessel_inventory_v1beta1_k8s_cluster_proto_rawDesc = nil
	file_kessel_inventory_v1beta1_k8s_cluster_proto_goTypes = nil
	file_kessel_inventory_v1beta1_k8s_cluster_proto_depIdxs = nil
}
