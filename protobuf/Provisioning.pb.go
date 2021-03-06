// Code generated by protoc-gen-go.
// source: Provisioning.proto
// DO NOT EDIT!

package signalservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProvisionEnvelope struct {
	PublicKey        []byte `protobuf:"bytes,1,opt,name=publicKey" json:"publicKey,omitempty"`
	Body             []byte `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProvisionEnvelope) Reset()                    { *m = ProvisionEnvelope{} }
func (m *ProvisionEnvelope) String() string            { return proto.CompactTextString(m) }
func (*ProvisionEnvelope) ProtoMessage()               {}
func (*ProvisionEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ProvisionEnvelope) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *ProvisionEnvelope) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type ProvisionMessage struct {
	IdentityKeyPublic  []byte  `protobuf:"bytes,1,opt,name=identityKeyPublic" json:"identityKeyPublic,omitempty"`
	IdentityKeyPrivate []byte  `protobuf:"bytes,2,opt,name=identityKeyPrivate" json:"identityKeyPrivate,omitempty"`
	Number             *string `protobuf:"bytes,3,opt,name=number" json:"number,omitempty"`
	ProvisioningCode   *string `protobuf:"bytes,4,opt,name=provisioningCode" json:"provisioningCode,omitempty"`
	UserAgent          *string `protobuf:"bytes,5,opt,name=userAgent" json:"userAgent,omitempty"`
	ProfileKey         []byte  `protobuf:"bytes,6,opt,name=profileKey" json:"profileKey,omitempty"`
	ReadReceipts       *bool   `protobuf:"varint,7,opt,name=readReceipts" json:"readReceipts,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *ProvisionMessage) Reset()                    { *m = ProvisionMessage{} }
func (m *ProvisionMessage) String() string            { return proto.CompactTextString(m) }
func (*ProvisionMessage) ProtoMessage()               {}
func (*ProvisionMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ProvisionMessage) GetIdentityKeyPublic() []byte {
	if m != nil {
		return m.IdentityKeyPublic
	}
	return nil
}

func (m *ProvisionMessage) GetIdentityKeyPrivate() []byte {
	if m != nil {
		return m.IdentityKeyPrivate
	}
	return nil
}

func (m *ProvisionMessage) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

func (m *ProvisionMessage) GetProvisioningCode() string {
	if m != nil && m.ProvisioningCode != nil {
		return *m.ProvisioningCode
	}
	return ""
}

func (m *ProvisionMessage) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

func (m *ProvisionMessage) GetProfileKey() []byte {
	if m != nil {
		return m.ProfileKey
	}
	return nil
}

func (m *ProvisionMessage) GetReadReceipts() bool {
	if m != nil && m.ReadReceipts != nil {
		return *m.ReadReceipts
	}
	return false
}

func init() {
	proto.RegisterType((*ProvisionEnvelope)(nil), "signalservice.ProvisionEnvelope")
	proto.RegisterType((*ProvisionMessage)(nil), "signalservice.ProvisionMessage")
}

func init() { proto.RegisterFile("Provisioning.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0x6a, 0xf2, 0x40,
	0x10, 0x85, 0x89, 0xbf, 0xbf, 0xad, 0x83, 0x05, 0x9d, 0x8b, 0xb2, 0x17, 0xa5, 0x88, 0x57, 0x52,
	0xca, 0xbe, 0x43, 0x6d, 0xbd, 0x92, 0x82, 0xe4, 0x0d, 0xa2, 0x99, 0xc6, 0x81, 0xb8, 0xbb, 0xec,
	0x6c, 0x52, 0xf2, 0x22, 0x7d, 0xde, 0x92, 0x25, 0x68, 0x82, 0xbd, 0xdb, 0xfd, 0xce, 0xe1, 0xec,
	0xce, 0x1c, 0xc0, 0xbd, 0xb7, 0x35, 0x0b, 0x5b, 0xc3, 0xa6, 0xd0, 0xce, 0xdb, 0x60, 0xf1, 0x41,
	0xb8, 0x30, 0x59, 0x29, 0xe4, 0x6b, 0x3e, 0xd2, 0x6a, 0x0b, 0x8b, 0x8b, 0x69, 0x6b, 0x6a, 0x2a,
	0xad, 0x23, 0x7c, 0x82, 0xa9, 0xab, 0x0e, 0x25, 0x1f, 0x77, 0xd4, 0xa8, 0x64, 0x99, 0xac, 0x67,
	0xe9, 0x15, 0x20, 0xc2, 0xf8, 0x60, 0xf3, 0x46, 0x8d, 0xa2, 0x10, 0xcf, 0xab, 0x9f, 0x11, 0xcc,
	0x2f, 0x39, 0x9f, 0x24, 0x92, 0x15, 0x84, 0xaf, 0xb0, 0xe0, 0x9c, 0x4c, 0xe0, 0xd0, 0xec, 0xa8,
	0xd9, 0xc7, 0x80, 0x2e, 0xee, 0x56, 0x40, 0x0d, 0xd8, 0x87, 0x9e, 0xeb, 0x2c, 0x50, 0xf7, 0xc8,
	0x1f, 0x0a, 0x3e, 0xc2, 0xc4, 0x54, 0xe7, 0x03, 0x79, 0xf5, 0x6f, 0x99, 0xac, 0xa7, 0x69, 0x77,
	0xc3, 0x17, 0x98, 0xbb, 0xde, 0xd8, 0xef, 0x36, 0x27, 0x35, 0x8e, 0x8e, 0x1b, 0xde, 0x0e, 0x5a,
	0x09, 0xf9, 0xb7, 0x82, 0x4c, 0x50, 0xff, 0xa3, 0xe9, 0x0a, 0xf0, 0x19, 0xc0, 0x79, 0xfb, 0xc5,
	0x25, 0xb5, 0x7b, 0x98, 0xc4, 0x9f, 0xf4, 0x08, 0xae, 0x60, 0xe6, 0x29, 0xcb, 0x53, 0x3a, 0x12,
	0xbb, 0x20, 0xea, 0x6e, 0x99, 0xac, 0xef, 0xd3, 0x01, 0xdb, 0x7c, 0x80, 0xb6, 0xbe, 0xd0, 0xdf,
	0x27, 0x16, 0x47, 0x5e, 0x1a, 0x09, 0x74, 0x16, 0x3d, 0xe8, 0x40, 0xb3, 0x09, 0xe4, 0x4d, 0x56,
	0x6a, 0x57, 0xc9, 0x69, 0x33, 0x28, 0x6d, 0xdf, 0x76, 0x26, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0x84, 0x2a, 0x4b, 0xc9, 0x01, 0x00, 0x00,
}
