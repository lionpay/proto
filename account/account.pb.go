// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.2
// source: account/account.proto

package account

import (
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

// 账户请求定义
type FundPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderNo    string `protobuf:"bytes,1,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`          //订单号
	MerchantId string `protobuf:"bytes,2,opt,name=merchant_id,json=merchantId,proto3" json:"merchant_id,omitempty"` //商户id
	Currency   string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`                       //币种
}

func (x *FundPaymentRequest) Reset() {
	*x = FundPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundPaymentRequest) ProtoMessage() {}

func (x *FundPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundPaymentRequest.ProtoReflect.Descriptor instead.
func (*FundPaymentRequest) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{0}
}

func (x *FundPaymentRequest) GetOrderNo() string {
	if x != nil {
		return x.OrderNo
	}
	return ""
}

func (x *FundPaymentRequest) GetMerchantId() string {
	if x != nil {
		return x.MerchantId
	}
	return ""
}

func (x *FundPaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

// 返回结构体
type JsonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`    //返回消息
	Code int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"` //业务码
}

func (x *JsonResponse) Reset() {
	*x = JsonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonResponse) ProtoMessage() {}

func (x *JsonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonResponse.ProtoReflect.Descriptor instead.
func (*JsonResponse) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{1}
}

func (x *JsonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *JsonResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_account_account_proto protoreflect.FileDescriptor

var file_account_account_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x6c, 0x0a, 0x12, 0x46, 0x75, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x34,
	0x0a, 0x0c, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x32, 0x4c, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x41, 0x0a, 0x0b, 0x46, 0x75, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_account_proto_rawDescOnce sync.Once
	file_account_account_proto_rawDescData = file_account_account_proto_rawDesc
)

func file_account_account_proto_rawDescGZIP() []byte {
	file_account_account_proto_rawDescOnce.Do(func() {
		file_account_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_account_proto_rawDescData)
	})
	return file_account_account_proto_rawDescData
}

var file_account_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_account_account_proto_goTypes = []interface{}{
	(*FundPaymentRequest)(nil), // 0: account.FundPaymentRequest
	(*JsonResponse)(nil),       // 1: account.JsonResponse
}
var file_account_account_proto_depIdxs = []int32{
	0, // 0: account.Account.FundPayment:input_type -> account.FundPaymentRequest
	1, // 1: account.Account.FundPayment:output_type -> account.JsonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_account_proto_init() }
func file_account_account_proto_init() {
	if File_account_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundPaymentRequest); i {
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
		file_account_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonResponse); i {
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
			RawDescriptor: file_account_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_account_proto_goTypes,
		DependencyIndexes: file_account_account_proto_depIdxs,
		MessageInfos:      file_account_account_proto_msgTypes,
	}.Build()
	File_account_account_proto = out.File
	file_account_account_proto_rawDesc = nil
	file_account_account_proto_goTypes = nil
	file_account_account_proto_depIdxs = nil
}
