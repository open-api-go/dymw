// Code generated by protoc-gen-go. DO NOT EDIT.
// source: douyin-webapp/account.proto

package douyin_webapp // import "github.com/dev-openapi/douyin-webapp"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetUserInfoReq struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	OpenId               string   `protobuf:"bytes,2,opt,name=open_id,json=openId" json:"open_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoReq) Reset()         { *m = GetUserInfoReq{} }
func (m *GetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoReq) ProtoMessage()    {}
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_84294f619ecb4a71, []int{0}
}
func (m *GetUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoReq.Unmarshal(m, b)
}
func (m *GetUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoReq.Marshal(b, m, deterministic)
}
func (dst *GetUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoReq.Merge(dst, src)
}
func (m *GetUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoReq.Size(m)
}
func (m *GetUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoReq proto.InternalMessageInfo

func (m *GetUserInfoReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *GetUserInfoReq) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

type UserInfo struct {
	Avatar               string   `protobuf:"bytes,1,opt,name=avatar" json:"avatar,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	OpenId               string   `protobuf:"bytes,5,opt,name=open_id,json=openId" json:"open_id,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	EAccountRole         string   `protobuf:"bytes,7,opt,name=e_account_role,json=eAccountRole" json:"e_account_role,omitempty"`
	ErrorCode            int32    `protobuf:"varint,8,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	UnionId              string   `protobuf:"bytes,9,opt,name=union_id,json=unionId" json:"union_id,omitempty"`
	EncryptMobile        string   `protobuf:"bytes,100,opt,name=encrypt_mobile,json=encryptMobile" json:"encrypt_mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_84294f619ecb4a71, []int{1}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfo) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *UserInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UserInfo) GetEAccountRole() string {
	if m != nil {
		return m.EAccountRole
	}
	return ""
}

func (m *UserInfo) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *UserInfo) GetUnionId() string {
	if m != nil {
		return m.UnionId
	}
	return ""
}

func (m *UserInfo) GetEncryptMobile() string {
	if m != nil {
		return m.EncryptMobile
	}
	return ""
}

type GetUserInfoRes struct {
	Message              string    `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Data                 *UserInfo `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetUserInfoRes) Reset()         { *m = GetUserInfoRes{} }
func (m *GetUserInfoRes) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoRes) ProtoMessage()    {}
func (*GetUserInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_84294f619ecb4a71, []int{2}
}
func (m *GetUserInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoRes.Unmarshal(m, b)
}
func (m *GetUserInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoRes.Marshal(b, m, deterministic)
}
func (dst *GetUserInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoRes.Merge(dst, src)
}
func (m *GetUserInfoRes) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoRes.Size(m)
}
func (m *GetUserInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoRes proto.InternalMessageInfo

func (m *GetUserInfoRes) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetUserInfoRes) GetData() *UserInfo {
	if m != nil {
		return m.Data
	}
	return nil
}

type FansCheckReq struct {
	FollowerOpenId       string   `protobuf:"bytes,1,opt,name=follower_open_id,json=followerOpenId" json:"follower_open_id,omitempty"`
	OpenId               string   `protobuf:"bytes,2,opt,name=open_id,json=openId" json:"open_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FansCheckReq) Reset()         { *m = FansCheckReq{} }
func (m *FansCheckReq) String() string { return proto.CompactTextString(m) }
func (*FansCheckReq) ProtoMessage()    {}
func (*FansCheckReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_84294f619ecb4a71, []int{3}
}
func (m *FansCheckReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FansCheckReq.Unmarshal(m, b)
}
func (m *FansCheckReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FansCheckReq.Marshal(b, m, deterministic)
}
func (dst *FansCheckReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FansCheckReq.Merge(dst, src)
}
func (m *FansCheckReq) XXX_Size() int {
	return xxx_messageInfo_FansCheckReq.Size(m)
}
func (m *FansCheckReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FansCheckReq.DiscardUnknown(m)
}

var xxx_messageInfo_FansCheckReq proto.InternalMessageInfo

func (m *FansCheckReq) GetFollowerOpenId() string {
	if m != nil {
		return m.FollowerOpenId
	}
	return ""
}

func (m *FansCheckReq) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

type FansCheckRes struct {
	Extra                *Extra             `protobuf:"bytes,1,opt,name=extra" json:"extra,omitempty"`
	Data                 *FansCheckRes_Data `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FansCheckRes) Reset()         { *m = FansCheckRes{} }
func (m *FansCheckRes) String() string { return proto.CompactTextString(m) }
func (*FansCheckRes) ProtoMessage()    {}
func (*FansCheckRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_84294f619ecb4a71, []int{4}
}
func (m *FansCheckRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FansCheckRes.Unmarshal(m, b)
}
func (m *FansCheckRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FansCheckRes.Marshal(b, m, deterministic)
}
func (dst *FansCheckRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FansCheckRes.Merge(dst, src)
}
func (m *FansCheckRes) XXX_Size() int {
	return xxx_messageInfo_FansCheckRes.Size(m)
}
func (m *FansCheckRes) XXX_DiscardUnknown() {
	xxx_messageInfo_FansCheckRes.DiscardUnknown(m)
}

var xxx_messageInfo_FansCheckRes proto.InternalMessageInfo

func (m *FansCheckRes) GetExtra() *Extra {
	if m != nil {
		return m.Extra
	}
	return nil
}

func (m *FansCheckRes) GetData() *FansCheckRes_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type FansCheckRes_Data struct {
	IsFollower           bool     `protobuf:"varint,1,opt,name=is_follower,json=isFollower" json:"is_follower,omitempty"`
	FollowTime           int64    `protobuf:"varint,2,opt,name=follow_time,json=followTime" json:"follow_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FansCheckRes_Data) Reset()         { *m = FansCheckRes_Data{} }
func (m *FansCheckRes_Data) String() string { return proto.CompactTextString(m) }
func (*FansCheckRes_Data) ProtoMessage()    {}
func (*FansCheckRes_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_84294f619ecb4a71, []int{4, 0}
}
func (m *FansCheckRes_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FansCheckRes_Data.Unmarshal(m, b)
}
func (m *FansCheckRes_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FansCheckRes_Data.Marshal(b, m, deterministic)
}
func (dst *FansCheckRes_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FansCheckRes_Data.Merge(dst, src)
}
func (m *FansCheckRes_Data) XXX_Size() int {
	return xxx_messageInfo_FansCheckRes_Data.Size(m)
}
func (m *FansCheckRes_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_FansCheckRes_Data.DiscardUnknown(m)
}

var xxx_messageInfo_FansCheckRes_Data proto.InternalMessageInfo

func (m *FansCheckRes_Data) GetIsFollower() bool {
	if m != nil {
		return m.IsFollower
	}
	return false
}

func (m *FansCheckRes_Data) GetFollowTime() int64 {
	if m != nil {
		return m.FollowTime
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUserInfoReq)(nil), "open.douyin.com.GetUserInfoReq")
	proto.RegisterType((*UserInfo)(nil), "open.douyin.com.UserInfo")
	proto.RegisterType((*GetUserInfoRes)(nil), "open.douyin.com.GetUserInfoRes")
	proto.RegisterType((*FansCheckReq)(nil), "open.douyin.com.FansCheckReq")
	proto.RegisterType((*FansCheckRes)(nil), "open.douyin.com.FansCheckRes")
	proto.RegisterType((*FansCheckRes_Data)(nil), "open.douyin.com.FansCheckRes.Data")
}

func init() {
	proto.RegisterFile("douyin-webapp/account.proto", fileDescriptor_account_84294f619ecb4a71)
}

var fileDescriptor_account_84294f619ecb4a71 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdf, 0x4e, 0x13, 0x4f,
	0x14, 0x4e, 0xf9, 0x41, 0xff, 0x9c, 0xe5, 0x57, 0xc8, 0x5c, 0xe0, 0xb6, 0x4a, 0xc0, 0x0d, 0x1a,
	0x42, 0xa0, 0x9b, 0x60, 0xe2, 0x05, 0x77, 0x8a, 0xa2, 0x24, 0x1a, 0xe3, 0x8a, 0x17, 0x7a, 0xb3,
	0x99, 0xce, 0x9e, 0x96, 0x09, 0xed, 0x9c, 0x65, 0x66, 0x0a, 0x72, 0x67, 0x7c, 0x05, 0xdf, 0xc7,
	0x97, 0xf0, 0x15, 0xf4, 0x3d, 0xcc, 0xcc, 0x6c, 0x49, 0x0b, 0x11, 0x2f, 0xcf, 0x77, 0xbe, 0xf9,
	0xe6, 0x3b, 0xdf, 0x99, 0x81, 0xfb, 0x05, 0x4d, 0xae, 0xa4, 0xda, 0xbb, 0xc4, 0x3e, 0x2f, 0xcb,
	0x94, 0x0b, 0x41, 0x13, 0x65, 0x7b, 0xa5, 0x26, 0x4b, 0x6c, 0x85, 0x4a, 0x54, 0xbd, 0xc0, 0xe8,
	0x09, 0x1a, 0x77, 0x1f, 0x0c, 0x89, 0x86, 0x23, 0x4c, 0x79, 0x29, 0x53, 0xae, 0x14, 0x59, 0x6e,
	0x25, 0x29, 0x13, 0xe8, 0xdd, 0xce, 0xbc, 0x16, 0x7e, 0xb1, 0x9a, 0x87, 0x56, 0xf2, 0x06, 0xda,
	0xaf, 0xd0, 0x7e, 0x34, 0xa8, 0x8f, 0xd5, 0x80, 0x32, 0x3c, 0x67, 0x0f, 0x61, 0x99, 0x0b, 0x81,
	0xc6, 0xe4, 0x96, 0xce, 0x50, 0xc5, 0xb5, 0xcd, 0xda, 0x76, 0x2b, 0x8b, 0x02, 0x76, 0xe2, 0x20,
	0x76, 0x0f, 0x1a, 0xce, 0x40, 0x2e, 0x8b, 0x78, 0xc1, 0x77, 0xeb, 0xae, 0x3c, 0x2e, 0x92, 0xaf,
	0x0b, 0xd0, 0x9c, 0x6a, 0xb1, 0x35, 0xa8, 0xf3, 0x0b, 0x6e, 0xb9, 0xae, 0x24, 0xaa, 0x8a, 0x75,
	0xa1, 0xa9, 0xa4, 0x38, 0x53, 0x7c, 0x8c, 0xd5, 0xf1, 0xeb, 0x7a, 0x56, 0x79, 0x69, 0x56, 0x99,
	0x6d, 0x42, 0x54, 0xa0, 0x11, 0x5a, 0x96, 0x6e, 0xb0, 0xb8, 0x1e, 0x4c, 0xcd, 0x40, 0x6c, 0x0b,
	0xda, 0x98, 0x57, 0x31, 0xe5, 0x9a, 0x46, 0x18, 0x37, 0x3c, 0x69, 0x19, 0x9f, 0x05, 0x30, 0xa3,
	0x11, 0xb2, 0x75, 0x00, 0xd4, 0x9a, 0x74, 0x2e, 0xa8, 0xc0, 0xb8, 0xb9, 0x59, 0xdb, 0x5e, 0xca,
	0x5a, 0x1e, 0x39, 0xa4, 0x02, 0x59, 0x07, 0x9a, 0x13, 0x25, 0xc9, 0x1b, 0x68, 0xf9, 0xe3, 0x0d,
	0x5f, 0x1f, 0x17, 0xec, 0x11, 0xb4, 0x51, 0x09, 0x7d, 0x55, 0xda, 0x7c, 0x4c, 0x7d, 0x39, 0xc2,
	0xb8, 0xf0, 0x84, 0xff, 0x2b, 0xf4, 0xad, 0x07, 0x93, 0x4f, 0x37, 0x02, 0x35, 0x2c, 0x86, 0xc6,
	0x18, 0x8d, 0xe1, 0x43, 0xac, 0x82, 0x98, 0x96, 0x6c, 0x0f, 0x16, 0x0b, 0x6e, 0xb9, 0x4f, 0x21,
	0xda, 0xef, 0xf4, 0x6e, 0x6c, 0xb5, 0x77, 0xad, 0xe2, 0x69, 0xc9, 0x7b, 0x58, 0x3e, 0xe2, 0xca,
	0x1c, 0x9e, 0xa2, 0x38, 0x73, 0x9b, 0xda, 0x86, 0xd5, 0x01, 0x8d, 0x46, 0x74, 0x89, 0x3a, 0x9f,
	0xa6, 0x16, 0x6e, 0x68, 0x4f, 0xf1, 0x77, 0x21, 0xbd, 0xbf, 0x2e, 0xec, 0x47, 0x6d, 0x4e, 0xd3,
	0xb0, 0x5d, 0x58, 0xf2, 0xcf, 0xc3, 0x0b, 0x45, 0xfb, 0x6b, 0xb7, 0x3c, 0xbd, 0x74, 0xdd, 0x2c,
	0x90, 0xd8, 0xd3, 0xb9, 0x01, 0x92, 0x5b, 0xe4, 0x59, 0xe9, 0xde, 0x0b, 0x6e, 0x79, 0x98, 0xa4,
	0xfb, 0x1a, 0x16, 0x5d, 0xc5, 0x36, 0x20, 0x92, 0x26, 0x9f, 0x9a, 0xf5, 0x77, 0x36, 0x33, 0x90,
	0xe6, 0xa8, 0x42, 0x1c, 0x21, 0x74, 0x73, 0x2b, 0xab, 0xe7, 0xf2, 0x5f, 0x06, 0x01, 0x3a, 0x91,
	0x63, 0xdc, 0xff, 0x5d, 0x83, 0x76, 0xb5, 0xdf, 0x0f, 0xa8, 0x2f, 0xa4, 0x40, 0x56, 0x42, 0x34,
	0xb3, 0x01, 0xb6, 0x71, 0xcb, 0xd5, 0xfc, 0x83, 0xef, 0xfe, 0x83, 0x60, 0x92, 0x8d, 0x6f, 0x3f,
	0x7f, 0x7d, 0x5f, 0xe8, 0x24, 0x2b, 0x29, 0xf1, 0x89, 0x3d, 0x4d, 0x27, 0x06, 0xb5, 0x54, 0x03,
	0x3a, 0xa8, 0xef, 0xec, 0x0e, 0x48, 0x8f, 0x19, 0x87, 0xd6, 0xf5, 0xa4, 0x6c, 0xfd, 0xae, 0x14,
	0xce, 0xbb, 0x77, 0xb6, 0x4d, 0xb2, 0xe6, 0xef, 0x5a, 0x4d, 0xa2, 0x74, 0xc0, 0x95, 0x49, 0x85,
	0xc3, 0x0f, 0x6a, 0x3b, 0xcf, 0x1f, 0x7f, 0xde, 0x1a, 0x4a, 0x7b, 0x3a, 0xe9, 0xbb, 0x23, 0x69,
	0x81, 0x17, 0x7b, 0x4e, 0xc6, 0x7d, 0xf7, 0xb9, 0xbf, 0xdd, 0xaf, 0xfb, 0x6f, 0xfd, 0xe4, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0xef, 0xac, 0xed, 0x3f, 0x04, 0x00, 0x00,
}
