// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mimapi/mimapi.proto

/*
Package mimapi is a generated protocol buffer package.

It is generated from these files:
	mimapi/mimapi.proto

It has these top-level messages:
	Provable
	Updateable
	BoardOwner
	Board
	Thread
	Post
	Vote
	Key
	Truststate
	Address
	Subprotocol
	Protocol
	Client
*/
package mimapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Provable struct {
	Fingerprint string `protobuf:"bytes,1,opt,name=Fingerprint" json:"Fingerprint,omitempty"`
	Creation    int64  `protobuf:"varint,2,opt,name=Creation" json:"Creation,omitempty"`
	ProofOfWork string `protobuf:"bytes,3,opt,name=ProofOfWork" json:"ProofOfWork,omitempty"`
	Signature   string `protobuf:"bytes,4,opt,name=Signature" json:"Signature,omitempty"`
}

func (m *Provable) Reset()                    { *m = Provable{} }
func (m *Provable) String() string            { return proto.CompactTextString(m) }
func (*Provable) ProtoMessage()               {}
func (*Provable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Provable) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *Provable) GetCreation() int64 {
	if m != nil {
		return m.Creation
	}
	return 0
}

func (m *Provable) GetProofOfWork() string {
	if m != nil {
		return m.ProofOfWork
	}
	return ""
}

func (m *Provable) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type Updateable struct {
	LastUpdate        int64  `protobuf:"varint,1,opt,name=LastUpdate" json:"LastUpdate,omitempty"`
	UpdateProofOfWork string `protobuf:"bytes,2,opt,name=UpdateProofOfWork" json:"UpdateProofOfWork,omitempty"`
	UpdateSignature   string `protobuf:"bytes,3,opt,name=UpdateSignature" json:"UpdateSignature,omitempty"`
}

func (m *Updateable) Reset()                    { *m = Updateable{} }
func (m *Updateable) String() string            { return proto.CompactTextString(m) }
func (*Updateable) ProtoMessage()               {}
func (*Updateable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Updateable) GetLastUpdate() int64 {
	if m != nil {
		return m.LastUpdate
	}
	return 0
}

func (m *Updateable) GetUpdateProofOfWork() string {
	if m != nil {
		return m.UpdateProofOfWork
	}
	return ""
}

func (m *Updateable) GetUpdateSignature() string {
	if m != nil {
		return m.UpdateSignature
	}
	return ""
}

type BoardOwner struct {
	KeyFingerprint string `protobuf:"bytes,1,opt,name=KeyFingerprint" json:"KeyFingerprint,omitempty"`
	Expiry         int64  `protobuf:"varint,2,opt,name=Expiry" json:"Expiry,omitempty"`
	Level          int32  `protobuf:"varint,3,opt,name=Level" json:"Level,omitempty"`
}

func (m *BoardOwner) Reset()                    { *m = BoardOwner{} }
func (m *BoardOwner) String() string            { return proto.CompactTextString(m) }
func (*BoardOwner) ProtoMessage()               {}
func (*BoardOwner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BoardOwner) GetKeyFingerprint() string {
	if m != nil {
		return m.KeyFingerprint
	}
	return ""
}

func (m *BoardOwner) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *BoardOwner) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type Board struct {
	Provable       *Provable     `protobuf:"bytes,1,opt,name=Provable" json:"Provable,omitempty"`
	Name           string        `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	BoardOwners    []*BoardOwner `protobuf:"bytes,3,rep,name=BoardOwners" json:"BoardOwners,omitempty"`
	Description    string        `protobuf:"bytes,4,opt,name=Description" json:"Description,omitempty"`
	Owner          string        `protobuf:"bytes,5,opt,name=Owner" json:"Owner,omitempty"`
	OwnerPublicKey string        `protobuf:"bytes,6,opt,name=OwnerPublicKey" json:"OwnerPublicKey,omitempty"`
	EntityVersion  int32         `protobuf:"varint,7,opt,name=EntityVersion" json:"EntityVersion,omitempty"`
	Language       string        `protobuf:"bytes,8,opt,name=Language" json:"Language,omitempty"`
	Meta           string        `protobuf:"bytes,9,opt,name=Meta" json:"Meta,omitempty"`
	RealmId        string        `protobuf:"bytes,10,opt,name=RealmId" json:"RealmId,omitempty"`
	EncrContent    string        `protobuf:"bytes,11,opt,name=EncrContent" json:"EncrContent,omitempty"`
	Updateable     *Updateable   `protobuf:"bytes,12,opt,name=Updateable" json:"Updateable,omitempty"`
}

func (m *Board) Reset()                    { *m = Board{} }
func (m *Board) String() string            { return proto.CompactTextString(m) }
func (*Board) ProtoMessage()               {}
func (*Board) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Board) GetProvable() *Provable {
	if m != nil {
		return m.Provable
	}
	return nil
}

func (m *Board) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Board) GetBoardOwners() []*BoardOwner {
	if m != nil {
		return m.BoardOwners
	}
	return nil
}

func (m *Board) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Board) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Board) GetOwnerPublicKey() string {
	if m != nil {
		return m.OwnerPublicKey
	}
	return ""
}

func (m *Board) GetEntityVersion() int32 {
	if m != nil {
		return m.EntityVersion
	}
	return 0
}

func (m *Board) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Board) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Board) GetRealmId() string {
	if m != nil {
		return m.RealmId
	}
	return ""
}

func (m *Board) GetEncrContent() string {
	if m != nil {
		return m.EncrContent
	}
	return ""
}

func (m *Board) GetUpdateable() *Updateable {
	if m != nil {
		return m.Updateable
	}
	return nil
}

type Thread struct {
	Provable       *Provable   `protobuf:"bytes,1,opt,name=Provable" json:"Provable,omitempty"`
	Board          string      `protobuf:"bytes,2,opt,name=Board" json:"Board,omitempty"`
	Name           string      `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Body           string      `protobuf:"bytes,4,opt,name=Body" json:"Body,omitempty"`
	Link           string      `protobuf:"bytes,5,opt,name=Link" json:"Link,omitempty"`
	Owner          string      `protobuf:"bytes,6,opt,name=Owner" json:"Owner,omitempty"`
	OwnerPublicKey string      `protobuf:"bytes,7,opt,name=OwnerPublicKey" json:"OwnerPublicKey,omitempty"`
	EntityVersion  int32       `protobuf:"varint,8,opt,name=EntityVersion" json:"EntityVersion,omitempty"`
	Meta           string      `protobuf:"bytes,9,opt,name=Meta" json:"Meta,omitempty"`
	RealmId        string      `protobuf:"bytes,10,opt,name=RealmId" json:"RealmId,omitempty"`
	EncrContent    string      `protobuf:"bytes,11,opt,name=EncrContent" json:"EncrContent,omitempty"`
	Updateable     *Updateable `protobuf:"bytes,12,opt,name=Updateable" json:"Updateable,omitempty"`
}

func (m *Thread) Reset()                    { *m = Thread{} }
func (m *Thread) String() string            { return proto.CompactTextString(m) }
func (*Thread) ProtoMessage()               {}
func (*Thread) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Thread) GetProvable() *Provable {
	if m != nil {
		return m.Provable
	}
	return nil
}

func (m *Thread) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *Thread) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Thread) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Thread) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Thread) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Thread) GetOwnerPublicKey() string {
	if m != nil {
		return m.OwnerPublicKey
	}
	return ""
}

func (m *Thread) GetEntityVersion() int32 {
	if m != nil {
		return m.EntityVersion
	}
	return 0
}

func (m *Thread) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Thread) GetRealmId() string {
	if m != nil {
		return m.RealmId
	}
	return ""
}

func (m *Thread) GetEncrContent() string {
	if m != nil {
		return m.EncrContent
	}
	return ""
}

func (m *Thread) GetUpdateable() *Updateable {
	if m != nil {
		return m.Updateable
	}
	return nil
}

type Post struct {
	Provable       *Provable   `protobuf:"bytes,1,opt,name=Provable" json:"Provable,omitempty"`
	Board          string      `protobuf:"bytes,2,opt,name=Board" json:"Board,omitempty"`
	Thread         string      `protobuf:"bytes,3,opt,name=Thread" json:"Thread,omitempty"`
	Parent         string      `protobuf:"bytes,4,opt,name=Parent" json:"Parent,omitempty"`
	Body           string      `protobuf:"bytes,5,opt,name=Body" json:"Body,omitempty"`
	Owner          string      `protobuf:"bytes,6,opt,name=Owner" json:"Owner,omitempty"`
	OwnerPublicKey string      `protobuf:"bytes,7,opt,name=OwnerPublicKey" json:"OwnerPublicKey,omitempty"`
	EntityVersion  int32       `protobuf:"varint,8,opt,name=EntityVersion" json:"EntityVersion,omitempty"`
	Meta           string      `protobuf:"bytes,9,opt,name=Meta" json:"Meta,omitempty"`
	RealmId        string      `protobuf:"bytes,10,opt,name=RealmId" json:"RealmId,omitempty"`
	EncrContent    string      `protobuf:"bytes,11,opt,name=EncrContent" json:"EncrContent,omitempty"`
	Updateable     *Updateable `protobuf:"bytes,12,opt,name=Updateable" json:"Updateable,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Post) GetProvable() *Provable {
	if m != nil {
		return m.Provable
	}
	return nil
}

func (m *Post) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *Post) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *Post) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Post) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Post) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Post) GetOwnerPublicKey() string {
	if m != nil {
		return m.OwnerPublicKey
	}
	return ""
}

func (m *Post) GetEntityVersion() int32 {
	if m != nil {
		return m.EntityVersion
	}
	return 0
}

func (m *Post) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Post) GetRealmId() string {
	if m != nil {
		return m.RealmId
	}
	return ""
}

func (m *Post) GetEncrContent() string {
	if m != nil {
		return m.EncrContent
	}
	return ""
}

func (m *Post) GetUpdateable() *Updateable {
	if m != nil {
		return m.Updateable
	}
	return nil
}

type Vote struct {
	Provable       *Provable   `protobuf:"bytes,1,opt,name=Provable" json:"Provable,omitempty"`
	Board          string      `protobuf:"bytes,2,opt,name=Board" json:"Board,omitempty"`
	Thread         string      `protobuf:"bytes,3,opt,name=Thread" json:"Thread,omitempty"`
	Target         string      `protobuf:"bytes,4,opt,name=Target" json:"Target,omitempty"`
	Owner          string      `protobuf:"bytes,5,opt,name=Owner" json:"Owner,omitempty"`
	OwnerPublicKey string      `protobuf:"bytes,6,opt,name=OwnerPublicKey" json:"OwnerPublicKey,omitempty"`
	TypeClass      int32       `protobuf:"varint,7,opt,name=TypeClass" json:"TypeClass,omitempty"`
	Type           int32       `protobuf:"varint,8,opt,name=Type" json:"Type,omitempty"`
	EntityVersion  int32       `protobuf:"varint,9,opt,name=EntityVersion" json:"EntityVersion,omitempty"`
	Meta           string      `protobuf:"bytes,10,opt,name=Meta" json:"Meta,omitempty"`
	RealmId        string      `protobuf:"bytes,11,opt,name=RealmId" json:"RealmId,omitempty"`
	EncrContent    string      `protobuf:"bytes,12,opt,name=EncrContent" json:"EncrContent,omitempty"`
	Updateable     *Updateable `protobuf:"bytes,13,opt,name=Updateable" json:"Updateable,omitempty"`
}

func (m *Vote) Reset()                    { *m = Vote{} }
func (m *Vote) String() string            { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()               {}
func (*Vote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Vote) GetProvable() *Provable {
	if m != nil {
		return m.Provable
	}
	return nil
}

func (m *Vote) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *Vote) GetThread() string {
	if m != nil {
		return m.Thread
	}
	return ""
}

func (m *Vote) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Vote) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Vote) GetOwnerPublicKey() string {
	if m != nil {
		return m.OwnerPublicKey
	}
	return ""
}

func (m *Vote) GetTypeClass() int32 {
	if m != nil {
		return m.TypeClass
	}
	return 0
}

func (m *Vote) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Vote) GetEntityVersion() int32 {
	if m != nil {
		return m.EntityVersion
	}
	return 0
}

func (m *Vote) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Vote) GetRealmId() string {
	if m != nil {
		return m.RealmId
	}
	return ""
}

func (m *Vote) GetEncrContent() string {
	if m != nil {
		return m.EncrContent
	}
	return ""
}

func (m *Vote) GetUpdateable() *Updateable {
	if m != nil {
		return m.Updateable
	}
	return nil
}

type Key struct {
	Provable      *Provable   `protobuf:"bytes,1,opt,name=Provable" json:"Provable,omitempty"`
	Type          string      `protobuf:"bytes,2,opt,name=Type" json:"Type,omitempty"`
	Key           string      `protobuf:"bytes,3,opt,name=Key" json:"Key,omitempty"`
	Expiry        int64       `protobuf:"varint,4,opt,name=Expiry" json:"Expiry,omitempty"`
	Name          string      `protobuf:"bytes,5,opt,name=Name" json:"Name,omitempty"`
	Info          string      `protobuf:"bytes,6,opt,name=Info" json:"Info,omitempty"`
	EntityVersion int32       `protobuf:"varint,7,opt,name=EntityVersion" json:"EntityVersion,omitempty"`
	Meta          string      `protobuf:"bytes,8,opt,name=Meta" json:"Meta,omitempty"`
	RealmId       string      `protobuf:"bytes,9,opt,name=RealmId" json:"RealmId,omitempty"`
	EncrContent   string      `protobuf:"bytes,10,opt,name=EncrContent" json:"EncrContent,omitempty"`
	Updateable    *Updateable `protobuf:"bytes,11,opt,name=Updateable" json:"Updateable,omitempty"`
}

func (m *Key) Reset()                    { *m = Key{} }
func (m *Key) String() string            { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()               {}
func (*Key) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Key) GetProvable() *Provable {
	if m != nil {
		return m.Provable
	}
	return nil
}

func (m *Key) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Key) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *Key) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Key) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Key) GetEntityVersion() int32 {
	if m != nil {
		return m.EntityVersion
	}
	return 0
}

func (m *Key) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Key) GetRealmId() string {
	if m != nil {
		return m.RealmId
	}
	return ""
}

func (m *Key) GetEncrContent() string {
	if m != nil {
		return m.EncrContent
	}
	return ""
}

func (m *Key) GetUpdateable() *Updateable {
	if m != nil {
		return m.Updateable
	}
	return nil
}

type Truststate struct {
	Provable       *Provable   `protobuf:"bytes,1,opt,name=Provable" json:"Provable,omitempty"`
	Target         string      `protobuf:"bytes,2,opt,name=Target" json:"Target,omitempty"`
	Owner          string      `protobuf:"bytes,3,opt,name=Owner" json:"Owner,omitempty"`
	OwnerPublicKey string      `protobuf:"bytes,4,opt,name=OwnerPublicKey" json:"OwnerPublicKey,omitempty"`
	TypeClass      int32       `protobuf:"varint,5,opt,name=TypeClass" json:"TypeClass,omitempty"`
	Type           int32       `protobuf:"varint,6,opt,name=Type" json:"Type,omitempty"`
	Domain         string      `protobuf:"bytes,7,opt,name=Domain" json:"Domain,omitempty"`
	Expiry         int64       `protobuf:"varint,8,opt,name=Expiry" json:"Expiry,omitempty"`
	EntityVersion  int32       `protobuf:"varint,9,opt,name=EntityVersion" json:"EntityVersion,omitempty"`
	Meta           string      `protobuf:"bytes,10,opt,name=Meta" json:"Meta,omitempty"`
	RealmId        string      `protobuf:"bytes,11,opt,name=RealmId" json:"RealmId,omitempty"`
	EncrContent    string      `protobuf:"bytes,12,opt,name=EncrContent" json:"EncrContent,omitempty"`
	Updateable     *Updateable `protobuf:"bytes,13,opt,name=Updateable" json:"Updateable,omitempty"`
}

func (m *Truststate) Reset()                    { *m = Truststate{} }
func (m *Truststate) String() string            { return proto.CompactTextString(m) }
func (*Truststate) ProtoMessage()               {}
func (*Truststate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Truststate) GetProvable() *Provable {
	if m != nil {
		return m.Provable
	}
	return nil
}

func (m *Truststate) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Truststate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Truststate) GetOwnerPublicKey() string {
	if m != nil {
		return m.OwnerPublicKey
	}
	return ""
}

func (m *Truststate) GetTypeClass() int32 {
	if m != nil {
		return m.TypeClass
	}
	return 0
}

func (m *Truststate) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Truststate) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Truststate) GetExpiry() int64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *Truststate) GetEntityVersion() int32 {
	if m != nil {
		return m.EntityVersion
	}
	return 0
}

func (m *Truststate) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Truststate) GetRealmId() string {
	if m != nil {
		return m.RealmId
	}
	return ""
}

func (m *Truststate) GetEncrContent() string {
	if m != nil {
		return m.EncrContent
	}
	return ""
}

func (m *Truststate) GetUpdateable() *Updateable {
	if m != nil {
		return m.Updateable
	}
	return nil
}

//
// Likely, when we are communicating the address, most of these fields would be empty. Nevertheless, it's good to have it here, and in the future, for some reason we can't predict right now, we might actually need it.
type Address struct {
	Location           string    `protobuf:"bytes,1,opt,name=Location" json:"Location,omitempty"`
	Sublocation        string    `protobuf:"bytes,2,opt,name=Sublocation" json:"Sublocation,omitempty"`
	LocationType       int32     `protobuf:"varint,3,opt,name=LocationType" json:"LocationType,omitempty"`
	Port               int32     `protobuf:"varint,4,opt,name=Port" json:"Port,omitempty"`
	LastSuccessfulPing int64     `protobuf:"varint,5,opt,name=LastSuccessfulPing" json:"LastSuccessfulPing,omitempty"`
	LastSuccessfulSync int64     `protobuf:"varint,6,opt,name=LastSuccessfulSync" json:"LastSuccessfulSync,omitempty"`
	Protocol           *Protocol `protobuf:"bytes,7,opt,name=Protocol" json:"Protocol,omitempty"`
	Client             *Client   `protobuf:"bytes,8,opt,name=Client" json:"Client,omitempty"`
	EntityVersion      int32     `protobuf:"varint,9,opt,name=EntityVersion" json:"EntityVersion,omitempty"`
	RealmId            string    `protobuf:"bytes,10,opt,name=RealmId" json:"RealmId,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Address) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Address) GetSublocation() string {
	if m != nil {
		return m.Sublocation
	}
	return ""
}

func (m *Address) GetLocationType() int32 {
	if m != nil {
		return m.LocationType
	}
	return 0
}

func (m *Address) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Address) GetLastSuccessfulPing() int64 {
	if m != nil {
		return m.LastSuccessfulPing
	}
	return 0
}

func (m *Address) GetLastSuccessfulSync() int64 {
	if m != nil {
		return m.LastSuccessfulSync
	}
	return 0
}

func (m *Address) GetProtocol() *Protocol {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func (m *Address) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Address) GetEntityVersion() int32 {
	if m != nil {
		return m.EntityVersion
	}
	return 0
}

func (m *Address) GetRealmId() string {
	if m != nil {
		return m.RealmId
	}
	return ""
}

type Subprotocol struct {
	Name              string   `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	VersionMajor      int32    `protobuf:"varint,2,opt,name=VersionMajor" json:"VersionMajor,omitempty"`
	VersionMinor      int32    `protobuf:"varint,3,opt,name=VersionMinor" json:"VersionMinor,omitempty"`
	SupportedEntities []string `protobuf:"bytes,4,rep,name=SupportedEntities" json:"SupportedEntities,omitempty"`
}

func (m *Subprotocol) Reset()                    { *m = Subprotocol{} }
func (m *Subprotocol) String() string            { return proto.CompactTextString(m) }
func (*Subprotocol) ProtoMessage()               {}
func (*Subprotocol) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Subprotocol) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subprotocol) GetVersionMajor() int32 {
	if m != nil {
		return m.VersionMajor
	}
	return 0
}

func (m *Subprotocol) GetVersionMinor() int32 {
	if m != nil {
		return m.VersionMinor
	}
	return 0
}

func (m *Subprotocol) GetSupportedEntities() []string {
	if m != nil {
		return m.SupportedEntities
	}
	return nil
}

type Protocol struct {
	VersionMajor int32          `protobuf:"varint,1,opt,name=VersionMajor" json:"VersionMajor,omitempty"`
	VersionMinor int32          `protobuf:"varint,2,opt,name=VersionMinor" json:"VersionMinor,omitempty"`
	Subprotocols []*Subprotocol `protobuf:"bytes,3,rep,name=Subprotocols" json:"Subprotocols,omitempty"`
}

func (m *Protocol) Reset()                    { *m = Protocol{} }
func (m *Protocol) String() string            { return proto.CompactTextString(m) }
func (*Protocol) ProtoMessage()               {}
func (*Protocol) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Protocol) GetVersionMajor() int32 {
	if m != nil {
		return m.VersionMajor
	}
	return 0
}

func (m *Protocol) GetVersionMinor() int32 {
	if m != nil {
		return m.VersionMinor
	}
	return 0
}

func (m *Protocol) GetSubprotocols() []*Subprotocol {
	if m != nil {
		return m.Subprotocols
	}
	return nil
}

type Client struct {
	VersionMajor int32  `protobuf:"varint,1,opt,name=VersionMajor" json:"VersionMajor,omitempty"`
	VersionMinor int32  `protobuf:"varint,2,opt,name=VersionMinor" json:"VersionMinor,omitempty"`
	VersionPatch int32  `protobuf:"varint,3,opt,name=VersionPatch" json:"VersionPatch,omitempty"`
	ClientName   string `protobuf:"bytes,4,opt,name=ClientName" json:"ClientName,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Client) GetVersionMajor() int32 {
	if m != nil {
		return m.VersionMajor
	}
	return 0
}

func (m *Client) GetVersionMinor() int32 {
	if m != nil {
		return m.VersionMinor
	}
	return 0
}

func (m *Client) GetVersionPatch() int32 {
	if m != nil {
		return m.VersionPatch
	}
	return 0
}

func (m *Client) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func init() {
	proto.RegisterType((*Provable)(nil), "mimapi.Provable")
	proto.RegisterType((*Updateable)(nil), "mimapi.Updateable")
	proto.RegisterType((*BoardOwner)(nil), "mimapi.BoardOwner")
	proto.RegisterType((*Board)(nil), "mimapi.Board")
	proto.RegisterType((*Thread)(nil), "mimapi.Thread")
	proto.RegisterType((*Post)(nil), "mimapi.Post")
	proto.RegisterType((*Vote)(nil), "mimapi.Vote")
	proto.RegisterType((*Key)(nil), "mimapi.Key")
	proto.RegisterType((*Truststate)(nil), "mimapi.Truststate")
	proto.RegisterType((*Address)(nil), "mimapi.Address")
	proto.RegisterType((*Subprotocol)(nil), "mimapi.Subprotocol")
	proto.RegisterType((*Protocol)(nil), "mimapi.Protocol")
	proto.RegisterType((*Client)(nil), "mimapi.Client")
}

func init() { proto.RegisterFile("mimapi/mimapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 958 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x57, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0x7a, 0xbd, 0x1b, 0xfb, 0xd9, 0x2d, 0x65, 0x5a, 0x45, 0x2b, 0x54, 0x21, 0x6b, 0x85,
	0xaa, 0x1c, 0x4a, 0x22, 0x05, 0x24, 0xce, 0x24, 0x0d, 0x52, 0xd5, 0x94, 0x5a, 0xeb, 0x50, 0x24,
	0x6e, 0xe3, 0xf5, 0xc4, 0x59, 0xba, 0xde, 0x59, 0xcd, 0xce, 0x16, 0x7c, 0xe7, 0x80, 0xc4, 0xb5,
	0x12, 0x17, 0xae, 0xfc, 0x12, 0xfe, 0x10, 0x17, 0x38, 0xa3, 0x79, 0x33, 0xe3, 0x9d, 0x8d, 0x5d,
	0xc7, 0x11, 0x45, 0x02, 0x71, 0xf2, 0x7b, 0xdf, 0x7b, 0x9e, 0x79, 0xef, 0xfb, 0x9e, 0x67, 0xc6,
	0x70, 0x7f, 0x91, 0x2d, 0x68, 0x99, 0x1d, 0xe9, 0x8f, 0xc3, 0x52, 0x70, 0xc9, 0x49, 0xa8, 0xbd,
	0xf8, 0x47, 0x0f, 0x7a, 0x63, 0xc1, 0x5f, 0xd3, 0x69, 0xce, 0xc8, 0x08, 0x06, 0x5f, 0x64, 0xc5,
	0x9c, 0x89, 0x52, 0x64, 0x85, 0x8c, 0xbc, 0x91, 0x77, 0xd0, 0x4f, 0x5c, 0x88, 0x7c, 0x00, 0xbd,
	0x53, 0xc1, 0xa8, 0xcc, 0x78, 0x11, 0x75, 0x46, 0xde, 0x81, 0x9f, 0xac, 0x7c, 0xf5, 0xed, 0xb1,
	0xe0, 0xfc, 0xf2, 0xc5, 0xe5, 0xd7, 0x5c, 0xbc, 0x8a, 0x7c, 0xfd, 0x6d, 0x07, 0x22, 0x0f, 0xa1,
	0x3f, 0xc9, 0xe6, 0x05, 0x95, 0xb5, 0x60, 0x51, 0x17, 0xe3, 0x0d, 0x10, 0xff, 0xe0, 0x01, 0x7c,
	0x55, 0xce, 0xa8, 0x64, 0x58, 0xcc, 0x87, 0x00, 0xe7, 0xb4, 0x92, 0x1a, 0xc1, 0x5a, 0xfc, 0xc4,
	0x41, 0xc8, 0x63, 0x78, 0x5f, 0x5b, 0xee, 0xa6, 0x1d, 0x5c, 0x74, 0x3d, 0x40, 0x0e, 0xe0, 0x3d,
	0x0d, 0x36, 0x05, 0xe8, 0x02, 0xaf, 0xc3, 0xf1, 0x14, 0xe0, 0x84, 0x53, 0x31, 0x7b, 0xf1, 0x5d,
	0xc1, 0x04, 0x79, 0x04, 0x77, 0x9f, 0xb1, 0xe5, 0x3a, 0x2b, 0xd7, 0x50, 0xb2, 0x0f, 0xe1, 0xd9,
	0xf7, 0x65, 0x26, 0x96, 0x86, 0x16, 0xe3, 0x91, 0x07, 0x10, 0x9c, 0xb3, 0xd7, 0x2c, 0xc7, 0xdd,
	0x82, 0x44, 0x3b, 0xf1, 0xaf, 0x3e, 0x04, 0xb8, 0x09, 0x79, 0xdc, 0xd0, 0x8f, 0x2b, 0x0f, 0x8e,
	0xef, 0x1d, 0x1a, 0xa1, 0x2c, 0x9e, 0x34, 0x02, 0x11, 0xe8, 0x7e, 0x49, 0x17, 0xcc, 0xb4, 0x89,
	0x36, 0xf9, 0x14, 0x06, 0x4d, 0xbd, 0x55, 0xe4, 0x8f, 0xfc, 0x83, 0xc1, 0x31, 0xb1, 0x8b, 0x34,
	0xa1, 0xc4, 0x4d, 0x53, 0x62, 0x3d, 0x61, 0x55, 0x2a, 0xb2, 0x12, 0xb5, 0xd4, 0x62, 0xb8, 0x90,
	0xaa, 0x1c, 0x73, 0xa3, 0x00, 0x63, 0xc1, 0x8a, 0x0f, 0x34, 0xc6, 0xf5, 0x34, 0xcf, 0xd2, 0x67,
	0x6c, 0x19, 0x85, 0x9a, 0x8f, 0x36, 0x4a, 0x3e, 0x82, 0x3b, 0x67, 0x85, 0xcc, 0xe4, 0xf2, 0x25,
	0x13, 0x95, 0xda, 0x61, 0x0f, 0xfb, 0x6f, 0x83, 0x6a, 0x9c, 0xce, 0x69, 0x31, 0xaf, 0xe9, 0x9c,
	0x45, 0x3d, 0x5c, 0x67, 0xe5, 0xab, 0x5e, 0x9f, 0x33, 0x49, 0xa3, 0xbe, 0xee, 0x55, 0xd9, 0x24,
	0x82, 0xbd, 0x84, 0xd1, 0x7c, 0xf1, 0x74, 0x16, 0x01, 0xc2, 0xd6, 0x55, 0xfd, 0x9c, 0x15, 0xa9,
	0x38, 0xe5, 0x85, 0x64, 0x85, 0x8c, 0x06, 0xba, 0x1f, 0x07, 0x22, 0xc7, 0xee, 0x74, 0x45, 0x43,
	0xe4, 0x7a, 0x45, 0x53, 0x13, 0x49, 0x9c, 0xac, 0xf8, 0xf7, 0x0e, 0x84, 0x17, 0x57, 0x82, 0xd1,
	0xdb, 0x0a, 0xf5, 0xc0, 0xe8, 0x6b, 0x94, 0x32, 0x62, 0x5b, 0xf9, 0x7c, 0x47, 0x3e, 0x02, 0xdd,
	0x13, 0x3e, 0x5b, 0x1a, 0x05, 0xd0, 0x56, 0xd8, 0x79, 0x56, 0xbc, 0x32, 0xcc, 0xa3, 0xdd, 0xc8,
	0x11, 0x6e, 0x97, 0x63, 0x6f, 0x37, 0x39, 0x7a, 0x9b, 0xe4, 0xf8, 0x37, 0x50, 0xfe, 0x67, 0x07,
	0xba, 0x63, 0x5e, 0xc9, 0x77, 0x42, 0xf8, 0xbe, 0x95, 0xcf, 0x50, 0x6e, 0xc5, 0xdc, 0x87, 0x70,
	0x4c, 0x85, 0xaa, 0x5a, 0xd3, 0x6e, 0xbc, 0x95, 0x18, 0x81, 0x23, 0xc6, 0xff, 0x85, 0xf8, 0x37,
	0x3e, 0x74, 0x5f, 0x72, 0x3c, 0x58, 0xff, 0x51, 0xe2, 0x2f, 0xa8, 0x98, 0xb3, 0x15, 0xf1, 0xda,
	0xfb, 0x9b, 0x87, 0xcd, 0x43, 0xe8, 0x5f, 0x2c, 0x4b, 0x76, 0x9a, 0xd3, 0xaa, 0x32, 0x07, 0x4d,
	0x03, 0x28, 0x72, 0x95, 0x63, 0x98, 0x47, 0x7b, 0x5d, 0x96, 0xfe, 0x36, 0x59, 0x60, 0xb3, 0x2c,
	0x83, 0xad, 0xb2, 0x0c, 0x6f, 0x92, 0xe5, 0xce, 0x4e, 0xb2, 0xfc, 0xd6, 0x01, 0x5f, 0xf5, 0x78,
	0xeb, 0x8b, 0x02, 0x7b, 0x36, 0x17, 0x05, 0xf6, 0x7c, 0x0f, 0x17, 0x32, 0x82, 0xe0, 0x9a, 0xcd,
	0xa5, 0xd5, 0x6d, 0x5d, 0x5a, 0xf6, 0x9c, 0x0a, 0xda, 0xe7, 0xd4, 0xd3, 0xe2, 0x92, 0x1b, 0x05,
	0xd0, 0xde, 0xf1, 0x90, 0xb7, 0x2c, 0xf6, 0x36, 0xb3, 0xd8, 0xdf, 0xca, 0x22, 0xdc, 0xc4, 0xe2,
	0x60, 0x27, 0x16, 0x7f, 0xf1, 0x01, 0x2e, 0x44, 0x5d, 0xc9, 0x4a, 0xd2, 0x5b, 0x8f, 0x78, 0x33,
	0xb4, 0x9d, 0xcd, 0x43, 0xeb, 0x6f, 0x1f, 0xda, 0xee, 0xcd, 0x43, 0x1b, 0xbc, 0x6d, 0x68, 0x43,
	0x67, 0x68, 0xf7, 0x21, 0x7c, 0xc2, 0x17, 0x34, 0x2b, 0xcc, 0x59, 0x63, 0x3c, 0x47, 0xc6, 0x5e,
	0x4b, 0xc6, 0xff, 0xce, 0x90, 0xff, 0xd1, 0x81, 0xbd, 0xcf, 0x67, 0x33, 0xc1, 0xaa, 0x0a, 0xdf,
	0x04, 0x3c, 0xd5, 0x4f, 0x4c, 0xcf, 0xbc, 0x09, 0x8c, 0xaf, 0x76, 0x9f, 0xd4, 0xd3, 0xdc, 0x86,
	0xb5, 0x1c, 0x2e, 0x44, 0x62, 0x18, 0xda, 0x6c, 0xe4, 0x4f, 0x3f, 0xbb, 0x5a, 0x98, 0xea, 0x78,
	0xcc, 0x85, 0x3e, 0x82, 0x82, 0x04, 0x6d, 0x72, 0x08, 0x44, 0xbd, 0x2d, 0x27, 0x75, 0x9a, 0xb2,
	0xaa, 0xba, 0xac, 0xf3, 0x71, 0x56, 0xcc, 0x51, 0x16, 0x3f, 0xd9, 0x10, 0x59, 0xcf, 0x9f, 0x2c,
	0x8b, 0x14, 0xd5, 0x5a, 0xcb, 0x57, 0x11, 0x33, 0x71, 0x92, 0xa7, 0x3c, 0x47, 0xf5, 0xda, 0x13,
	0x87, 0x78, 0xb2, 0xca, 0x20, 0x8f, 0x20, 0x3c, 0xcd, 0x33, 0x45, 0x70, 0x0f, 0x73, 0xef, 0xda,
	0x5c, 0x8d, 0x26, 0x26, 0xba, 0xa3, 0xc2, 0x6f, 0xbd, 0x49, 0xe2, 0x9f, 0x3d, 0x24, 0xb4, 0xb4,
	0xfb, 0xda, 0x1f, 0xbe, 0xe7, 0xfc, 0xf0, 0x63, 0x18, 0x9a, 0x85, 0x9e, 0xd3, 0x6f, 0xb9, 0x40,
	0xd2, 0x83, 0xa4, 0x85, 0xb9, 0x39, 0x59, 0xc1, 0x85, 0x65, 0xdd, 0xc5, 0xd4, 0x7b, 0x7d, 0x52,
	0x97, 0x25, 0x17, 0x92, 0xcd, 0xb0, 0xbe, 0x8c, 0x55, 0x51, 0x77, 0xe4, 0xab, 0xf7, 0xfa, 0x5a,
	0x20, 0xfe, 0xc9, 0x6b, 0x08, 0x5b, 0x2b, 0xc1, 0xdb, 0xa1, 0x84, 0xce, 0x86, 0x12, 0x3e, 0x83,
	0xa1, 0xd3, 0xad, 0x7d, 0x2b, 0xdf, 0xb7, 0xe4, 0x3a, 0xb1, 0xa4, 0x95, 0x18, 0xbf, 0xf1, 0xac,
	0x20, 0xef, 0xac, 0x96, 0x26, 0x67, 0x4c, 0x65, 0x7a, 0x75, 0x8d, 0x32, 0xc4, 0xd4, 0x5f, 0x20,
	0xbd, 0x2b, 0x8a, 0xa2, 0x8f, 0x11, 0x07, 0x39, 0x89, 0xbf, 0x19, 0x51, 0x26, 0xaf, 0x98, 0xf8,
	0x38, 0xe5, 0x82, 0x1d, 0x69, 0xfb, 0x08, 0xcb, 0xae, 0xcc, 0xdf, 0xbd, 0x69, 0x88, 0xee, 0x27,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x85, 0x9b, 0x0b, 0x06, 0x0e, 0x00, 0x00,
}
