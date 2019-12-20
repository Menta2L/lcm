// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lcm.proto

package lcm

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type IssuerRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Issuer:
	//	*IssuerRequest_LetsEncryptIssuer
	//	*IssuerRequest_SelfSignedIssuer
	Issuer               isIssuerRequest_Issuer `protobuf_oneof:"issuer"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IssuerRequest) Reset()         { *m = IssuerRequest{} }
func (m *IssuerRequest) String() string { return proto.CompactTextString(m) }
func (*IssuerRequest) ProtoMessage()    {}
func (*IssuerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{0}
}

func (m *IssuerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssuerRequest.Unmarshal(m, b)
}
func (m *IssuerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssuerRequest.Marshal(b, m, deterministic)
}
func (m *IssuerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssuerRequest.Merge(m, src)
}
func (m *IssuerRequest) XXX_Size() int {
	return xxx_messageInfo_IssuerRequest.Size(m)
}
func (m *IssuerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssuerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssuerRequest proto.InternalMessageInfo

func (m *IssuerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isIssuerRequest_Issuer interface {
	isIssuerRequest_Issuer()
}

type IssuerRequest_LetsEncryptIssuer struct {
	LetsEncryptIssuer *LetsEncryptIssuerRequest `protobuf:"bytes,2,opt,name=lets_encrypt_issuer,json=letsEncryptIssuer,proto3,oneof"`
}

type IssuerRequest_SelfSignedIssuer struct {
	SelfSignedIssuer *SelfSignedIssuerRequest `protobuf:"bytes,3,opt,name=self_signed_issuer,json=selfSignedIssuer,proto3,oneof"`
}

func (*IssuerRequest_LetsEncryptIssuer) isIssuerRequest_Issuer() {}

func (*IssuerRequest_SelfSignedIssuer) isIssuerRequest_Issuer() {}

func (m *IssuerRequest) GetIssuer() isIssuerRequest_Issuer {
	if m != nil {
		return m.Issuer
	}
	return nil
}

func (m *IssuerRequest) GetLetsEncryptIssuer() *LetsEncryptIssuerRequest {
	if x, ok := m.GetIssuer().(*IssuerRequest_LetsEncryptIssuer); ok {
		return x.LetsEncryptIssuer
	}
	return nil
}

func (m *IssuerRequest) GetSelfSignedIssuer() *SelfSignedIssuerRequest {
	if x, ok := m.GetIssuer().(*IssuerRequest_SelfSignedIssuer); ok {
		return x.SelfSignedIssuer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*IssuerRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*IssuerRequest_LetsEncryptIssuer)(nil),
		(*IssuerRequest_SelfSignedIssuer)(nil),
	}
}

type LetsEncryptIssuerRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Server               string   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	SolverRef            string   `protobuf:"bytes,3,opt,name=solverRef,proto3" json:"solverRef,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LetsEncryptIssuerRequest) Reset()         { *m = LetsEncryptIssuerRequest{} }
func (m *LetsEncryptIssuerRequest) String() string { return proto.CompactTextString(m) }
func (*LetsEncryptIssuerRequest) ProtoMessage()    {}
func (*LetsEncryptIssuerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{1}
}

func (m *LetsEncryptIssuerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LetsEncryptIssuerRequest.Unmarshal(m, b)
}
func (m *LetsEncryptIssuerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LetsEncryptIssuerRequest.Marshal(b, m, deterministic)
}
func (m *LetsEncryptIssuerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LetsEncryptIssuerRequest.Merge(m, src)
}
func (m *LetsEncryptIssuerRequest) XXX_Size() int {
	return xxx_messageInfo_LetsEncryptIssuerRequest.Size(m)
}
func (m *LetsEncryptIssuerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LetsEncryptIssuerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LetsEncryptIssuerRequest proto.InternalMessageInfo

func (m *LetsEncryptIssuerRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LetsEncryptIssuerRequest) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *LetsEncryptIssuerRequest) GetSolverRef() string {
	if m != nil {
		return m.SolverRef
	}
	return ""
}

type Solver struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// Types that are valid to be assigned to DnsSolvers:
	//	*Solver_CloudFlareSolver
	//	*Solver_GoogleCloudSolver
	DnsSolvers           isSolver_DnsSolvers `protobuf_oneof:"dns_solvers"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Solver) Reset()         { *m = Solver{} }
func (m *Solver) String() string { return proto.CompactTextString(m) }
func (*Solver) ProtoMessage()    {}
func (*Solver) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{2}
}

func (m *Solver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Solver.Unmarshal(m, b)
}
func (m *Solver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Solver.Marshal(b, m, deterministic)
}
func (m *Solver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Solver.Merge(m, src)
}
func (m *Solver) XXX_Size() int {
	return xxx_messageInfo_Solver.Size(m)
}
func (m *Solver) XXX_DiscardUnknown() {
	xxx_messageInfo_Solver.DiscardUnknown(m)
}

var xxx_messageInfo_Solver proto.InternalMessageInfo

func (m *Solver) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Solver) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type isSolver_DnsSolvers interface {
	isSolver_DnsSolvers()
}

type Solver_CloudFlareSolver struct {
	CloudFlareSolver *CloudFlareSolver `protobuf:"bytes,3,opt,name=cloud_flare_solver,json=cloudFlareSolver,proto3,oneof"`
}

type Solver_GoogleCloudSolver struct {
	GoogleCloudSolver *GoogleCloudSolver `protobuf:"bytes,4,opt,name=google_cloud_solver,json=googleCloudSolver,proto3,oneof"`
}

func (*Solver_CloudFlareSolver) isSolver_DnsSolvers() {}

func (*Solver_GoogleCloudSolver) isSolver_DnsSolvers() {}

func (m *Solver) GetDnsSolvers() isSolver_DnsSolvers {
	if m != nil {
		return m.DnsSolvers
	}
	return nil
}

func (m *Solver) GetCloudFlareSolver() *CloudFlareSolver {
	if x, ok := m.GetDnsSolvers().(*Solver_CloudFlareSolver); ok {
		return x.CloudFlareSolver
	}
	return nil
}

func (m *Solver) GetGoogleCloudSolver() *GoogleCloudSolver {
	if x, ok := m.GetDnsSolvers().(*Solver_GoogleCloudSolver); ok {
		return x.GoogleCloudSolver
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Solver) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Solver_CloudFlareSolver)(nil),
		(*Solver_GoogleCloudSolver)(nil),
	}
}

type SolverShort struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SolverShort) Reset()         { *m = SolverShort{} }
func (m *SolverShort) String() string { return proto.CompactTextString(m) }
func (*SolverShort) ProtoMessage()    {}
func (*SolverShort) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{3}
}

func (m *SolverShort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SolverShort.Unmarshal(m, b)
}
func (m *SolverShort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SolverShort.Marshal(b, m, deterministic)
}
func (m *SolverShort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SolverShort.Merge(m, src)
}
func (m *SolverShort) XXX_Size() int {
	return xxx_messageInfo_SolverShort.Size(m)
}
func (m *SolverShort) XXX_DiscardUnknown() {
	xxx_messageInfo_SolverShort.DiscardUnknown(m)
}

var xxx_messageInfo_SolverShort proto.InternalMessageInfo

func (m *SolverShort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SolverShort) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type CloudFlareSolver struct {
	Apikey               string   `protobuf:"bytes,2,opt,name=apikey,proto3" json:"apikey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloudFlareSolver) Reset()         { *m = CloudFlareSolver{} }
func (m *CloudFlareSolver) String() string { return proto.CompactTextString(m) }
func (*CloudFlareSolver) ProtoMessage()    {}
func (*CloudFlareSolver) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{4}
}

func (m *CloudFlareSolver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloudFlareSolver.Unmarshal(m, b)
}
func (m *CloudFlareSolver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloudFlareSolver.Marshal(b, m, deterministic)
}
func (m *CloudFlareSolver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudFlareSolver.Merge(m, src)
}
func (m *CloudFlareSolver) XXX_Size() int {
	return xxx_messageInfo_CloudFlareSolver.Size(m)
}
func (m *CloudFlareSolver) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudFlareSolver.DiscardUnknown(m)
}

var xxx_messageInfo_CloudFlareSolver proto.InternalMessageInfo

func (m *CloudFlareSolver) GetApikey() string {
	if m != nil {
		return m.Apikey
	}
	return ""
}

type GoogleCloudSolver struct {
	Sa                   []byte   `protobuf:"bytes,2,opt,name=sa,proto3" json:"sa,omitempty"`
	Project              string   `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleCloudSolver) Reset()         { *m = GoogleCloudSolver{} }
func (m *GoogleCloudSolver) String() string { return proto.CompactTextString(m) }
func (*GoogleCloudSolver) ProtoMessage()    {}
func (*GoogleCloudSolver) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{5}
}

func (m *GoogleCloudSolver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleCloudSolver.Unmarshal(m, b)
}
func (m *GoogleCloudSolver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleCloudSolver.Marshal(b, m, deterministic)
}
func (m *GoogleCloudSolver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleCloudSolver.Merge(m, src)
}
func (m *GoogleCloudSolver) XXX_Size() int {
	return xxx_messageInfo_GoogleCloudSolver.Size(m)
}
func (m *GoogleCloudSolver) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleCloudSolver.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleCloudSolver proto.InternalMessageInfo

func (m *GoogleCloudSolver) GetSa() []byte {
	if m != nil {
		return m.Sa
	}
	return nil
}

func (m *GoogleCloudSolver) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type SelfSignedIssuerRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelfSignedIssuerRequest) Reset()         { *m = SelfSignedIssuerRequest{} }
func (m *SelfSignedIssuerRequest) String() string { return proto.CompactTextString(m) }
func (*SelfSignedIssuerRequest) ProtoMessage()    {}
func (*SelfSignedIssuerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{6}
}

func (m *SelfSignedIssuerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelfSignedIssuerRequest.Unmarshal(m, b)
}
func (m *SelfSignedIssuerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelfSignedIssuerRequest.Marshal(b, m, deterministic)
}
func (m *SelfSignedIssuerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelfSignedIssuerRequest.Merge(m, src)
}
func (m *SelfSignedIssuerRequest) XXX_Size() int {
	return xxx_messageInfo_SelfSignedIssuerRequest.Size(m)
}
func (m *SelfSignedIssuerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SelfSignedIssuerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SelfSignedIssuerRequest proto.InternalMessageInfo

type ListSolversRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSolversRequest) Reset()         { *m = ListSolversRequest{} }
func (m *ListSolversRequest) String() string { return proto.CompactTextString(m) }
func (*ListSolversRequest) ProtoMessage()    {}
func (*ListSolversRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{7}
}

func (m *ListSolversRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSolversRequest.Unmarshal(m, b)
}
func (m *ListSolversRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSolversRequest.Marshal(b, m, deterministic)
}
func (m *ListSolversRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSolversRequest.Merge(m, src)
}
func (m *ListSolversRequest) XXX_Size() int {
	return xxx_messageInfo_ListSolversRequest.Size(m)
}
func (m *ListSolversRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSolversRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSolversRequest proto.InternalMessageInfo

type ListSolversResponse struct {
	Solver               []*SolverShort `protobuf:"bytes,1,rep,name=solver,proto3" json:"solver,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListSolversResponse) Reset()         { *m = ListSolversResponse{} }
func (m *ListSolversResponse) String() string { return proto.CompactTextString(m) }
func (*ListSolversResponse) ProtoMessage()    {}
func (*ListSolversResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{8}
}

func (m *ListSolversResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSolversResponse.Unmarshal(m, b)
}
func (m *ListSolversResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSolversResponse.Marshal(b, m, deterministic)
}
func (m *ListSolversResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSolversResponse.Merge(m, src)
}
func (m *ListSolversResponse) XXX_Size() int {
	return xxx_messageInfo_ListSolversResponse.Size(m)
}
func (m *ListSolversResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSolversResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSolversResponse proto.InternalMessageInfo

func (m *ListSolversResponse) GetSolver() []*SolverShort {
	if m != nil {
		return m.Solver
	}
	return nil
}

// Request for a certificate to be issued
type CertificateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IssuerRef            string   `protobuf:"bytes,2,opt,name=issuer_ref,json=issuerRef,proto3" json:"issuer_ref,omitempty"`
	Domain               []string `protobuf:"bytes,3,rep,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{9}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CertificateRequest) GetIssuerRef() string {
	if m != nil {
		return m.IssuerRef
	}
	return ""
}

func (m *CertificateRequest) GetDomain() []string {
	if m != nil {
		return m.Domain
	}
	return nil
}

type CertificateResponse struct {
	// A blob containing the cert
	Certificate string `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// A blob containing the private key
	PrivateKey string `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	// The DNS names this certificate is valid for
	Domains              []string `protobuf:"bytes,3,rep,name=domains,proto3" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cf1c4246d481985, []int{10}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

func (m *CertificateResponse) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *CertificateResponse) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func init() {
	proto.RegisterType((*IssuerRequest)(nil), "lcm.IssuerRequest")
	proto.RegisterType((*LetsEncryptIssuerRequest)(nil), "lcm.LetsEncryptIssuerRequest")
	proto.RegisterType((*Solver)(nil), "lcm.Solver")
	proto.RegisterType((*SolverShort)(nil), "lcm.SolverShort")
	proto.RegisterType((*CloudFlareSolver)(nil), "lcm.CloudFlareSolver")
	proto.RegisterType((*GoogleCloudSolver)(nil), "lcm.GoogleCloudSolver")
	proto.RegisterType((*SelfSignedIssuerRequest)(nil), "lcm.SelfSignedIssuerRequest")
	proto.RegisterType((*ListSolversRequest)(nil), "lcm.ListSolversRequest")
	proto.RegisterType((*ListSolversResponse)(nil), "lcm.ListSolversResponse")
	proto.RegisterType((*CertificateRequest)(nil), "lcm.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "lcm.CertificateResponse")
}

func init() { proto.RegisterFile("lcm.proto", fileDescriptor_7cf1c4246d481985) }

var fileDescriptor_7cf1c4246d481985 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x4e, 0xdb, 0x4c,
	0x10, 0x26, 0x09, 0x7f, 0xfe, 0x66, 0x0c, 0x15, 0x4c, 0x68, 0x70, 0x53, 0x68, 0x23, 0x9f, 0xa2,
	0x1e, 0x1c, 0x29, 0x15, 0xb7, 0x56, 0x95, 0x82, 0x28, 0x54, 0x8d, 0x54, 0xc9, 0xb9, 0xf5, 0x62,
	0x19, 0x67, 0x1c, 0xb6, 0xdd, 0xb5, 0x8d, 0x77, 0x83, 0x94, 0x27, 0xe8, 0xcb, 0xf5, 0xd4, 0x27,
	0xaa, 0xbc, 0xbb, 0x86, 0xc4, 0x01, 0x6e, 0x9e, 0xf9, 0x66, 0xbe, 0x99, 0xf9, 0x66, 0xbc, 0xd0,
	0xe1, 0xb1, 0xf0, 0xf3, 0x22, 0x53, 0x19, 0xb6, 0x78, 0x2c, 0xfa, 0x6f, 0x16, 0x59, 0xb6, 0xe0,
	0x34, 0xd2, 0xae, 0xeb, 0x65, 0x32, 0x22, 0x91, 0xab, 0x95, 0x89, 0xe8, 0xbf, 0xab, 0x83, 0x8a,
	0x09, 0x92, 0x2a, 0x12, 0xb9, 0x09, 0xf0, 0xfe, 0x34, 0x60, 0xff, 0xab, 0x94, 0x4b, 0x2a, 0x02,
	0xba, 0x5d, 0x92, 0x54, 0x88, 0xb0, 0x9b, 0x46, 0x82, 0xdc, 0xc6, 0xa0, 0x31, 0xec, 0x04, 0xfa,
	0x1b, 0xbf, 0x43, 0x97, 0x93, 0x92, 0x21, 0xa5, 0x71, 0xb1, 0xca, 0x55, 0xc8, 0x74, 0x86, 0xdb,
	0x1c, 0x34, 0x86, 0xce, 0xf8, 0xd4, 0x2f, 0x3b, 0x9a, 0x92, 0x92, 0x17, 0x06, 0xde, 0xe0, 0xbb,
	0xda, 0x09, 0x0e, 0x79, 0x1d, 0xc3, 0x29, 0xa0, 0x24, 0x9e, 0x84, 0x92, 0x2d, 0x52, 0x9a, 0x57,
	0x7c, 0x2d, 0xcd, 0x77, 0xa2, 0xf9, 0x66, 0xc4, 0x93, 0x99, 0x46, 0xeb, 0x74, 0x07, 0xb2, 0x06,
	0x4d, 0x5e, 0x40, 0xdb, 0x30, 0x78, 0x09, 0xb8, 0x4f, 0x35, 0x82, 0x47, 0xf0, 0x1f, 0x89, 0x88,
	0x71, 0x3b, 0x99, 0x31, 0xb0, 0x07, 0x6d, 0x49, 0xc5, 0x9d, 0x9d, 0xa6, 0x13, 0x58, 0x0b, 0x4f,
	0xa0, 0x23, 0x33, 0x7e, 0x57, 0xa6, 0x27, 0xba, 0xb1, 0x4e, 0xf0, 0xe0, 0xf0, 0xfe, 0x36, 0xa0,
	0x3d, 0xd3, 0xd6, 0xa3, 0x7a, 0xdd, 0x97, 0x6a, 0xae, 0x97, 0xba, 0x00, 0x8c, 0x79, 0xb6, 0x9c,
	0x87, 0x09, 0x8f, 0x0a, 0x0a, 0x0d, 0x9b, 0x1d, 0xfa, 0x95, 0x1e, 0xfa, 0xbc, 0x84, 0xbf, 0x94,
	0xa8, 0x21, 0x2f, 0xa7, 0x8d, 0x6b, 0x3e, 0xbc, 0x82, 0xae, 0xd9, 0x6a, 0x68, 0xd8, 0x2c, 0xcf,
	0xae, 0xe6, 0xe9, 0x69, 0x9e, 0x4b, 0x8d, 0x6b, 0xb6, 0x7b, 0xa2, 0xc3, 0x45, 0xdd, 0x39, 0xd9,
	0x07, 0x67, 0x9e, 0x4a, 0x4b, 0x20, 0xbd, 0x33, 0x70, 0x0c, 0x30, 0xbb, 0xc9, 0x8a, 0xc7, 0x0f,
	0x01, 0x61, 0x57, 0xad, 0x72, 0xb2, 0x73, 0xe9, 0x6f, 0xef, 0x3d, 0x1c, 0xd4, 0xfb, 0x2e, 0x55,
	0x8d, 0x72, 0xf6, 0x8b, 0x56, 0x95, 0xaa, 0xc6, 0xf2, 0x3e, 0xc1, 0xe1, 0x56, 0x6f, 0xf8, 0x12,
	0x9a, 0x32, 0xd2, 0x81, 0x7b, 0x41, 0x53, 0x46, 0xe8, 0xc2, 0xff, 0x79, 0x91, 0xfd, 0xa4, 0x58,
	0x59, 0xe1, 0x2b, 0xd3, 0x7b, 0x0d, 0xc7, 0x4f, 0xdc, 0x85, 0x77, 0x04, 0x38, 0x65, 0x52, 0x19,
	0x4a, 0x59, 0x79, 0x3f, 0x43, 0x77, 0xc3, 0x2b, 0xf3, 0x2c, 0x95, 0x84, 0x43, 0x68, 0x5b, 0xd5,
	0x1a, 0x83, 0xd6, 0xd0, 0x19, 0x1f, 0x98, 0x93, 0x7b, 0x18, 0x3e, 0xb0, 0xb8, 0x17, 0x02, 0x9e,
	0x53, 0xa1, 0x58, 0xc2, 0xe2, 0x48, 0xd1, 0x73, 0xff, 0xc8, 0x29, 0x80, 0x39, 0xc2, 0xb0, 0xa0,
	0xc4, 0x8e, 0xdd, 0x61, 0xb6, 0xc7, 0xa4, 0x54, 0x64, 0x9e, 0x89, 0x88, 0xa5, 0x6e, 0x6b, 0xd0,
	0x2a, 0x15, 0x31, 0x96, 0x77, 0x0b, 0xdd, 0x8d, 0x02, 0xb6, 0xc3, 0x01, 0x38, 0xf1, 0x83, 0xdb,
	0x16, 0x5a, 0x77, 0xe1, 0x5b, 0x80, 0xbc, 0x60, 0x77, 0x91, 0xa2, 0x6f, 0xf7, 0x32, 0xaf, 0x79,
	0x4a, 0x15, 0x4d, 0x09, 0x69, 0x2b, 0x56, 0xe6, 0xf8, 0x77, 0x13, 0x5a, 0xd3, 0x58, 0xe0, 0x04,
	0x9c, 0x52, 0x1c, 0xbb, 0x7e, 0x3c, 0x36, 0xff, 0xf1, 0x96, 0x88, 0x7d, 0x77, 0x1b, 0x30, 0x5d,
	0x7a, 0x3b, 0x78, 0x06, 0x7b, 0xe7, 0x05, 0x45, 0xaa, 0x5a, 0xbc, 0xb3, 0xa6, 0x64, 0xbf, 0xe7,
	0x9b, 0x9b, 0xf3, 0xab, 0xe7, 0xc7, 0xbf, 0x28, 0xdf, 0x26, 0x6f, 0x07, 0x3f, 0x56, 0x69, 0xf6,
	0x3d, 0x40, 0x9d, 0xb6, 0xb1, 0xd1, 0x67, 0xb2, 0x2f, 0x01, 0x6d, 0xd0, 0x9a, 0x74, 0xb6, 0xff,
	0xed, 0x6d, 0x3d, 0x4d, 0x34, 0x71, 0x7f, 0xf4, 0x16, 0x4c, 0xdd, 0x2c, 0xaf, 0xfd, 0x38, 0x13,
	0x23, 0x41, 0xa9, 0x8a, 0xc6, 0x7c, 0xc4, 0x63, 0x71, 0xdd, 0xd6, 0xb1, 0x1f, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0x85, 0x56, 0x6d, 0x6e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LcmClient is the client API for Lcm service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LcmClient interface {
	Listsolvers(ctx context.Context, in *ListSolversRequest, opts ...grpc.CallOption) (*ListSolversResponse, error)
	CreateSolver(ctx context.Context, in *Solver, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateIssuer(ctx context.Context, in *IssuerRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RequestCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type lcmClient struct {
	cc *grpc.ClientConn
}

func NewLcmClient(cc *grpc.ClientConn) LcmClient {
	return &lcmClient{cc}
}

func (c *lcmClient) Listsolvers(ctx context.Context, in *ListSolversRequest, opts ...grpc.CallOption) (*ListSolversResponse, error) {
	out := new(ListSolversResponse)
	err := c.cc.Invoke(ctx, "/lcm.Lcm/Listsolvers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcmClient) CreateSolver(ctx context.Context, in *Solver, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/lcm.Lcm/CreateSolver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcmClient) CreateIssuer(ctx context.Context, in *IssuerRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/lcm.Lcm/CreateIssuer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lcmClient) RequestCertificate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/lcm.Lcm/RequestCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LcmServer is the server API for Lcm service.
type LcmServer interface {
	Listsolvers(context.Context, *ListSolversRequest) (*ListSolversResponse, error)
	CreateSolver(context.Context, *Solver) (*empty.Empty, error)
	CreateIssuer(context.Context, *IssuerRequest) (*empty.Empty, error)
	RequestCertificate(context.Context, *CertificateRequest) (*empty.Empty, error)
}

// UnimplementedLcmServer can be embedded to have forward compatible implementations.
type UnimplementedLcmServer struct {
}

func (*UnimplementedLcmServer) Listsolvers(ctx context.Context, req *ListSolversRequest) (*ListSolversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Listsolvers not implemented")
}
func (*UnimplementedLcmServer) CreateSolver(ctx context.Context, req *Solver) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSolver not implemented")
}
func (*UnimplementedLcmServer) CreateIssuer(ctx context.Context, req *IssuerRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssuer not implemented")
}
func (*UnimplementedLcmServer) RequestCertificate(ctx context.Context, req *CertificateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestCertificate not implemented")
}

func RegisterLcmServer(s *grpc.Server, srv LcmServer) {
	s.RegisterService(&_Lcm_serviceDesc, srv)
}

func _Lcm_Listsolvers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSolversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcmServer).Listsolvers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcm.Lcm/Listsolvers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcmServer).Listsolvers(ctx, req.(*ListSolversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lcm_CreateSolver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Solver)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcmServer).CreateSolver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcm.Lcm/CreateSolver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcmServer).CreateSolver(ctx, req.(*Solver))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lcm_CreateIssuer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssuerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcmServer).CreateIssuer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcm.Lcm/CreateIssuer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcmServer).CreateIssuer(ctx, req.(*IssuerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lcm_RequestCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LcmServer).RequestCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lcm.Lcm/RequestCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LcmServer).RequestCertificate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lcm_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lcm.Lcm",
	HandlerType: (*LcmServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Listsolvers",
			Handler:    _Lcm_Listsolvers_Handler,
		},
		{
			MethodName: "CreateSolver",
			Handler:    _Lcm_CreateSolver_Handler,
		},
		{
			MethodName: "CreateIssuer",
			Handler:    _Lcm_CreateIssuer_Handler,
		},
		{
			MethodName: "RequestCertificate",
			Handler:    _Lcm_RequestCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lcm.proto",
}
