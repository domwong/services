// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/posts/proto/posts/post.proto

package go_micro_service_posts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug                 string   `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TagNames             []string `protobuf:"bytes,6,rep,name=tagNames,proto3" json:"tagNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d32cca1c2f74735, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Post) GetTagNames() []string {
	if m != nil {
		return m.TagNames
	}
	return nil
}

type QueryRequest struct {
	Slug                 string   `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d32cca1c2f74735, []int{1}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *QueryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type QueryResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d32cca1c2f74735, []int{2}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type PostRequest struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostRequest) Reset()         { *m = PostRequest{} }
func (m *PostRequest) String() string { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()    {}
func (*PostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d32cca1c2f74735, []int{3}
}

func (m *PostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostRequest.Unmarshal(m, b)
}
func (m *PostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostRequest.Marshal(b, m, deterministic)
}
func (m *PostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostRequest.Merge(m, src)
}
func (m *PostRequest) XXX_Size() int {
	return xxx_messageInfo_PostRequest.Size(m)
}
func (m *PostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostRequest proto.InternalMessageInfo

func (m *PostRequest) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type PostResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d32cca1c2f74735, []int{4}
}

func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (m *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(m, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d32cca1c2f74735, []int{5}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d32cca1c2f74735, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Post)(nil), "go.micro.service.posts.Post")
	proto.RegisterType((*QueryRequest)(nil), "go.micro.service.posts.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "go.micro.service.posts.QueryResponse")
	proto.RegisterType((*PostRequest)(nil), "go.micro.service.posts.PostRequest")
	proto.RegisterType((*PostResponse)(nil), "go.micro.service.posts.PostResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.service.posts.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.service.posts.DeleteResponse")
}

func init() { proto.RegisterFile("blog/posts/proto/posts/post.proto", fileDescriptor_2d32cca1c2f74735) }

var fileDescriptor_2d32cca1c2f74735 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdb, 0x4a, 0xc3, 0x40,
	0x14, 0x34, 0x57, 0xed, 0xe9, 0x05, 0x59, 0xa4, 0x2c, 0xa5, 0x60, 0x8c, 0x55, 0xf2, 0x94, 0x4a,
	0xfd, 0x00, 0x05, 0x7d, 0x96, 0x1a, 0x41, 0xf0, 0xb1, 0x97, 0x6d, 0x58, 0x48, 0xba, 0x31, 0x7b,
	0x2a, 0xf8, 0x15, 0xbe, 0xf8, 0xc1, 0x92, 0x93, 0xa4, 0x56, 0xb1, 0x2d, 0xbe, 0xed, 0x0c, 0xc3,
	0xcc, 0xec, 0x24, 0x0b, 0x67, 0xd3, 0x44, 0xc5, 0xc3, 0x4c, 0x69, 0xd4, 0xc3, 0x2c, 0x57, 0xa8,
	0xea, 0xb3, 0xd2, 0x18, 0x12, 0xc1, 0xba, 0xb1, 0x0a, 0x53, 0x39, 0xcb, 0x55, 0xa8, 0x45, 0xfe,
	0x26, 0x67, 0x22, 0x24, 0x89, 0xff, 0x69, 0x80, 0x3d, 0x56, 0x1a, 0x59, 0x07, 0x4c, 0x39, 0xe7,
	0x86, 0x67, 0x04, 0x8d, 0xc8, 0x94, 0x73, 0x76, 0x02, 0x0e, 0x4a, 0x4c, 0x04, 0x37, 0x89, 0x2a,
	0x01, 0x63, 0x60, 0xeb, 0x64, 0x15, 0x73, 0x8b, 0x48, 0x3a, 0x33, 0x0e, 0x87, 0x33, 0xb5, 0x44,
	0xb1, 0x44, 0x6e, 0x13, 0x5d, 0x43, 0xd6, 0x87, 0x06, 0xca, 0x54, 0x68, 0x9c, 0xa4, 0x19, 0x77,
	0x3c, 0x23, 0xb0, 0xa2, 0x6f, 0x82, 0xf5, 0xe0, 0x08, 0x27, 0xf1, 0xc3, 0x24, 0x15, 0x9a, 0xbb,
	0x9e, 0x15, 0x34, 0xa2, 0x35, 0xf6, 0xc7, 0xd0, 0x7a, 0x5c, 0x89, 0xfc, 0x3d, 0x12, 0xaf, 0x2b,
	0xa1, 0x71, 0x9d, 0x6b, 0x6c, 0xe4, 0x76, 0xc1, 0x55, 0x8b, 0x85, 0x16, 0x48, 0x15, 0xad, 0xa8,
	0x42, 0x45, 0xf3, 0x44, 0xa6, 0x12, 0xa9, 0xa4, 0x15, 0x95, 0xc0, 0xbf, 0x83, 0x76, 0xe5, 0xa8,
	0x33, 0xb5, 0xd4, 0x82, 0x8d, 0xc0, 0xa1, 0x09, 0xb8, 0xe1, 0x59, 0x41, 0x73, 0xd4, 0x0f, 0xff,
	0x5e, 0x28, 0x2c, 0xd6, 0x89, 0x4a, 0xa9, 0x7f, 0x03, 0x4d, 0x82, 0x55, 0xab, 0x2b, 0xb0, 0x0b,
	0x9e, 0x5a, 0xed, 0x73, 0x20, 0xa5, 0x7f, 0x0b, 0xad, 0xd2, 0xa0, 0x2a, 0xf1, 0x7f, 0x87, 0x53,
	0x68, 0xdf, 0x8b, 0x44, 0xa0, 0xa8, 0x4b, 0xfc, 0xfa, 0x70, 0xfe, 0x31, 0x74, 0x6a, 0x41, 0x19,
	0x32, 0xfa, 0x30, 0xc1, 0x29, 0x1c, 0x34, 0x7b, 0x06, 0x87, 0x46, 0x60, 0x83, 0x6d, 0x49, 0x9b,
	0xab, 0xf7, 0x2e, 0xf6, 0xa8, 0x4a, 0x7f, 0xff, 0x80, 0x3d, 0x55, 0x3f, 0xd1, 0xf9, 0xce, 0x0b,
	0x54, 0xae, 0x83, 0xdd, 0xa2, 0xb5, 0xe9, 0x0b, 0xb8, 0xe5, 0x45, 0xd8, 0xd6, 0x1e, 0x3f, 0x96,
	0xe8, 0x5d, 0xee, 0x93, 0xd5, 0xd6, 0x53, 0x97, 0x1e, 0xc5, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x85, 0x6c, 0xbc, 0xeb, 0x39, 0x03, 0x00, 0x00,
}