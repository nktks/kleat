// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gitrepo_artifact_provider.proto

package proto

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

// Configuration for the Git repo artifact provider.
type GitRepoArtifactProvider struct {
	// Whether the Git repo artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Git Repo accounts.
	Accounts             []*GitRepoArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GitRepoArtifactProvider) Reset()         { *m = GitRepoArtifactProvider{} }
func (m *GitRepoArtifactProvider) String() string { return proto.CompactTextString(m) }
func (*GitRepoArtifactProvider) ProtoMessage()    {}
func (*GitRepoArtifactProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_bde915b29f1fac20, []int{0}
}

func (m *GitRepoArtifactProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitRepoArtifactProvider.Unmarshal(m, b)
}
func (m *GitRepoArtifactProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitRepoArtifactProvider.Marshal(b, m, deterministic)
}
func (m *GitRepoArtifactProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitRepoArtifactProvider.Merge(m, src)
}
func (m *GitRepoArtifactProvider) XXX_Size() int {
	return xxx_messageInfo_GitRepoArtifactProvider.Size(m)
}
func (m *GitRepoArtifactProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_GitRepoArtifactProvider.DiscardUnknown(m)
}

var xxx_messageInfo_GitRepoArtifactProvider proto.InternalMessageInfo

func (m *GitRepoArtifactProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *GitRepoArtifactProvider) GetAccounts() []*GitRepoArtifactAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// Configuration for a Git repo artifact account. An account maps to a
// credential that is able to authenticate against a Git repository hosted by a
// Git hosting service. Either `username` and `password`,
// `usernamePasswordFile`, `token`, `tokenFile`, or `sshPrivateKeyFilePath` and
// `sshPrivateKeyPassphrase` must be set.
type GitRepoArtifactAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The username of the account.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The password of the account.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// The path to a file containing the username and password of the account
	// in the format `${username}:${password}`.
	UsernamePasswordFile string `protobuf:"bytes,4,opt,name=usernamePasswordFile,proto3" json:"usernamePasswordFile,omitempty"`
	// The Git repo access token.
	Token string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	// The path to a file containing the Git repo access token.
	TokenFile string `protobuf:"bytes,6,opt,name=tokenFile,proto3" json:"tokenFile,omitempty"`
	// The path to an SSH private key to be used when connecting with the Git
	// repo over SSH.
	SshPrivateKeyFilePath string `protobuf:"bytes,7,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// The passphrase to an SSH private key to be used when connecting with
	// the Git repo over SSH.
	SshPrivateKeyPassphrase string `protobuf:"bytes,8,opt,name=sshPrivateKeyPassphrase,proto3" json:"sshPrivateKeyPassphrase,omitempty"`
	// The path to a `known_hosts` file to be used when connecting with a
	// Git repository over SSH.
	SshKnownHostsFilePath string `protobuf:"bytes,9,opt,name=sshKnownHostsFilePath,proto3" json:"sshKnownHostsFilePath,omitempty"`
	// If `true`, Spinnaker can connect with a Git repository over SSH without
	// verifying the server's IP address against a `known_hosts` file.
	SshTrustUnknownHosts bool     `protobuf:"varint,10,opt,name=sshTrustUnknownHosts,proto3" json:"sshTrustUnknownHosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitRepoArtifactAccount) Reset()         { *m = GitRepoArtifactAccount{} }
func (m *GitRepoArtifactAccount) String() string { return proto.CompactTextString(m) }
func (*GitRepoArtifactAccount) ProtoMessage()    {}
func (*GitRepoArtifactAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_bde915b29f1fac20, []int{1}
}

func (m *GitRepoArtifactAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitRepoArtifactAccount.Unmarshal(m, b)
}
func (m *GitRepoArtifactAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitRepoArtifactAccount.Marshal(b, m, deterministic)
}
func (m *GitRepoArtifactAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitRepoArtifactAccount.Merge(m, src)
}
func (m *GitRepoArtifactAccount) XXX_Size() int {
	return xxx_messageInfo_GitRepoArtifactAccount.Size(m)
}
func (m *GitRepoArtifactAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_GitRepoArtifactAccount.DiscardUnknown(m)
}

var xxx_messageInfo_GitRepoArtifactAccount proto.InternalMessageInfo

func (m *GitRepoArtifactAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetUsernamePasswordFile() string {
	if m != nil {
		return m.UsernamePasswordFile
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetTokenFile() string {
	if m != nil {
		return m.TokenFile
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetSshPrivateKeyFilePath() string {
	if m != nil {
		return m.SshPrivateKeyFilePath
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetSshPrivateKeyPassphrase() string {
	if m != nil {
		return m.SshPrivateKeyPassphrase
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetSshKnownHostsFilePath() string {
	if m != nil {
		return m.SshKnownHostsFilePath
	}
	return ""
}

func (m *GitRepoArtifactAccount) GetSshTrustUnknownHosts() bool {
	if m != nil {
		return m.SshTrustUnknownHosts
	}
	return false
}

func init() {
	proto.RegisterType((*GitRepoArtifactProvider)(nil), "proto.GitRepoArtifactProvider")
	proto.RegisterType((*GitRepoArtifactAccount)(nil), "proto.GitRepoArtifactAccount")
}

func init() { proto.RegisterFile("gitrepo_artifact_provider.proto", fileDescriptor_bde915b29f1fac20) }

var fileDescriptor_bde915b29f1fac20 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xd5, 0xff, 0xf4, 0x7c, 0x9b, 0xd5, 0x8f, 0x5a, 0x08, 0x44, 0xd5, 0xa9, 0x53, 0x87,
	0xc2, 0x00, 0x63, 0x17, 0x40, 0xea, 0x12, 0x45, 0x30, 0x57, 0x6e, 0x7b, 0x20, 0x56, 0x8b, 0x6d,
	0xf9, 0x38, 0xad, 0xb8, 0x21, 0xae, 0x13, 0xe5, 0xe4, 0x07, 0x01, 0x61, 0x8a, 0xcf, 0xfb, 0x3c,
	0xf6, 0x1b, 0xd9, 0x70, 0xf5, 0xaa, 0x83, 0x47, 0x67, 0xd7, 0xca, 0x07, 0xfd, 0xa2, 0xb6, 0x61,
	0xed, 0xbc, 0x3d, 0xea, 0x1d, 0xfa, 0xb9, 0xf3, 0x36, 0x58, 0xd1, 0xe3, 0xcf, 0xd4, 0xc0, 0xf8,
	0x41, 0x87, 0x04, 0x9d, 0x5d, 0x96, 0x62, 0x5c, 0x7a, 0x42, 0xc2, 0x00, 0x8d, 0xda, 0x1c, 0x70,
	0x27, 0x5b, 0x93, 0xd6, 0x2c, 0x4a, 0xaa, 0x51, 0xdc, 0x41, 0xa4, 0xb6, 0x5b, 0x9b, 0x99, 0x40,
	0xb2, 0x3d, 0xe9, 0xcc, 0xfe, 0x2d, 0x2e, 0x8b, 0x53, 0xe7, 0x3f, 0xce, 0x5a, 0x16, 0x56, 0x52,
	0xeb, 0xd3, 0x8f, 0x0e, 0x9c, 0x35, 0x4b, 0x42, 0x40, 0xd7, 0xa8, 0x37, 0xe4, 0xb2, 0x61, 0xc2,
	0x6b, 0x71, 0x0e, 0x51, 0x46, 0xe8, 0x39, 0x6f, 0x73, 0x5e, 0xcf, 0x39, 0x73, 0x8a, 0xe8, 0x64,
	0xfd, 0x4e, 0x76, 0x0a, 0x56, 0xcd, 0x62, 0x01, 0xa3, 0xca, 0x8b, 0xcb, 0xec, 0x5e, 0x1f, 0x50,
	0x76, 0xd9, 0x6b, 0x64, 0x62, 0x04, 0xbd, 0x60, 0xf7, 0x68, 0x64, 0x8f, 0xa5, 0x62, 0x10, 0x17,
	0x30, 0xe4, 0x05, 0x6f, 0xef, 0x33, 0xf9, 0x0a, 0xc4, 0x0d, 0xfc, 0x27, 0x4a, 0x63, 0xaf, 0x8f,
	0x2a, 0xe0, 0x0a, 0xdf, 0xf3, 0x30, 0x56, 0x21, 0x95, 0x03, 0x36, 0x9b, 0xa1, 0xb8, 0x85, 0xf1,
	0x37, 0x90, 0xff, 0x86, 0x4b, 0xbd, 0x22, 0x94, 0x11, 0xef, 0xfb, 0x0b, 0x97, 0x7d, 0x2b, 0x63,
	0x4f, 0xe6, 0xd1, 0x52, 0xa0, 0xba, 0x6f, 0x58, 0xf7, 0xfd, 0x86, 0xf9, 0x6d, 0x10, 0xa5, 0x4f,
	0x3e, 0xa3, 0xf0, 0x6c, 0xf6, 0x35, 0x97, 0xc0, 0xcf, 0xda, 0xc8, 0x36, 0x7d, 0x7e, 0xd0, 0xeb,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x80, 0xfe, 0x86, 0x49, 0x02, 0x00, 0x00,
}
