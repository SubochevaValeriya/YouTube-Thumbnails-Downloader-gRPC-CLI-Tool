// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/youtube_thumbnails.proto

package grpcYoutubeThumbnails

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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VideoId       string `protobuf:"bytes,3,opt,name=videoId,proto3" json:"videoId,omitempty"`
	ThumbnailLink string `protobuf:"bytes,4,opt,name=thumbnailLink,proto3" json:"thumbnailLink,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_youtube_thumbnails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_proto_youtube_thumbnails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_proto_youtube_thumbnails_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Video) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Video) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *Video) GetThumbnailLink() string {
	if x != nil {
		return x.ThumbnailLink
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_youtube_thumbnails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_youtube_thumbnails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_youtube_thumbnails_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Response) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type DownloadThumbnailLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *DownloadThumbnailLinkRequest) Reset() {
	*x = DownloadThumbnailLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_youtube_thumbnails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadThumbnailLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadThumbnailLinkRequest) ProtoMessage() {}

func (x *DownloadThumbnailLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_youtube_thumbnails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadThumbnailLinkRequest.ProtoReflect.Descriptor instead.
func (*DownloadThumbnailLinkRequest) Descriptor() ([]byte, []int) {
	return file_proto_youtube_thumbnails_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadThumbnailLinkRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type DownloadThumbnailLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DownloadThumbnailLinkResponse) Reset() {
	*x = DownloadThumbnailLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_youtube_thumbnails_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadThumbnailLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadThumbnailLinkResponse) ProtoMessage() {}

func (x *DownloadThumbnailLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_youtube_thumbnails_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadThumbnailLinkResponse.ProtoReflect.Descriptor instead.
func (*DownloadThumbnailLinkResponse) Descriptor() ([]byte, []int) {
	return file_proto_youtube_thumbnails_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadThumbnailLinkResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_proto_youtube_thumbnails_proto protoreflect.FileDescriptor

var file_proto_youtube_thumbnails_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x5f,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x1c, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22, 0x4c, 0x0a, 0x1d,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x7a, 0x0a, 0x18, 0x59, 0x6f,
	0x75, 0x74, 0x75, 0x62, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x60, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x75, 0x62, 0x6f, 0x63, 0x68, 0x65, 0x76, 0x61, 0x56, 0x61,
	0x6c, 0x65, 0x72, 0x69, 0x79, 0x61, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x2d,
	0x79, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x2d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x73, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x59, 0x6f, 0x75, 0x74, 0x75, 0x62, 0x65, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_youtube_thumbnails_proto_rawDescOnce sync.Once
	file_proto_youtube_thumbnails_proto_rawDescData = file_proto_youtube_thumbnails_proto_rawDesc
)

func file_proto_youtube_thumbnails_proto_rawDescGZIP() []byte {
	file_proto_youtube_thumbnails_proto_rawDescOnce.Do(func() {
		file_proto_youtube_thumbnails_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_youtube_thumbnails_proto_rawDescData)
	})
	return file_proto_youtube_thumbnails_proto_rawDescData
}

var file_proto_youtube_thumbnails_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_youtube_thumbnails_proto_goTypes = []interface{}{
	(*Video)(nil),                         // 0: proto.Video
	(*Response)(nil),                      // 1: proto.Response
	(*DownloadThumbnailLinkRequest)(nil),  // 2: proto.DownloadThumbnailLinkRequest
	(*DownloadThumbnailLinkResponse)(nil), // 3: proto.DownloadThumbnailLinkResponse
}
var file_proto_youtube_thumbnails_proto_depIdxs = []int32{
	1, // 0: proto.DownloadThumbnailLinkResponse.response:type_name -> proto.Response
	2, // 1: proto.YoutubeThumbnailsService.DownloadThumbnail:input_type -> proto.DownloadThumbnailLinkRequest
	3, // 2: proto.YoutubeThumbnailsService.DownloadThumbnail:output_type -> proto.DownloadThumbnailLinkResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_youtube_thumbnails_proto_init() }
func file_proto_youtube_thumbnails_proto_init() {
	if File_proto_youtube_thumbnails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_youtube_thumbnails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_proto_youtube_thumbnails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_youtube_thumbnails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadThumbnailLinkRequest); i {
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
		file_proto_youtube_thumbnails_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadThumbnailLinkResponse); i {
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
			RawDescriptor: file_proto_youtube_thumbnails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_youtube_thumbnails_proto_goTypes,
		DependencyIndexes: file_proto_youtube_thumbnails_proto_depIdxs,
		MessageInfos:      file_proto_youtube_thumbnails_proto_msgTypes,
	}.Build()
	File_proto_youtube_thumbnails_proto = out.File
	file_proto_youtube_thumbnails_proto_rawDesc = nil
	file_proto_youtube_thumbnails_proto_goTypes = nil
	file_proto_youtube_thumbnails_proto_depIdxs = nil
}