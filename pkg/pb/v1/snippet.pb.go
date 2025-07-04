// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/snippet.proto

package snippetpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type SnippetOwner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SnippetOwner) Reset() {
	*x = SnippetOwner{}
	mi := &file_v1_snippet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnippetOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnippetOwner) ProtoMessage() {}

func (x *SnippetOwner) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnippetOwner.ProtoReflect.Descriptor instead.
func (*SnippetOwner) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{0}
}

func (x *SnippetOwner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SnippetOwner) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Snippet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner         *SnippetOwner          `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	LanguageHint  string                 `protobuf:"bytes,5,opt,name=language_hint,json=languageHint,proto3" json:"language_hint,omitempty"`
	IsPublic      bool                   `protobuf:"varint,6,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	Tags          []string               `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Snippet) Reset() {
	*x = Snippet{}
	mi := &file_v1_snippet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Snippet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snippet) ProtoMessage() {}

func (x *Snippet) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snippet.ProtoReflect.Descriptor instead.
func (*Snippet) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{1}
}

func (x *Snippet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Snippet) GetOwner() *SnippetOwner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Snippet) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Snippet) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Snippet) GetLanguageHint() string {
	if x != nil {
		return x.LanguageHint
	}
	return ""
}

func (x *Snippet) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *Snippet) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Snippet) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Snippet) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type PaginationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Size          int32                  `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	TotalItems    int32                  `protobuf:"varint,2,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
	CurrentPage   int32                  `protobuf:"varint,3,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	TotalPages    int32                  `protobuf:"varint,4,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_v1_snippet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationResponse) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PaginationResponse) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

func (x *PaginationResponse) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *PaginationResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

type CreateSnippetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	LanguageHint  *string                `protobuf:"bytes,3,opt,name=language_hint,json=languageHint,proto3,oneof" json:"language_hint,omitempty"`
	IsPublic      *bool                  `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3,oneof" json:"is_public,omitempty"`
	Tags          []string               `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSnippetRequest) Reset() {
	*x = CreateSnippetRequest{}
	mi := &file_v1_snippet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnippetRequest) ProtoMessage() {}

func (x *CreateSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnippetRequest.ProtoReflect.Descriptor instead.
func (*CreateSnippetRequest) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSnippetRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateSnippetRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateSnippetRequest) GetLanguageHint() string {
	if x != nil && x.LanguageHint != nil {
		return *x.LanguageHint
	}
	return ""
}

func (x *CreateSnippetRequest) GetIsPublic() bool {
	if x != nil && x.IsPublic != nil {
		return *x.IsPublic
	}
	return false
}

func (x *CreateSnippetRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreateSnippetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Snippet       *Snippet               `protobuf:"bytes,1,opt,name=snippet,proto3" json:"snippet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSnippetResponse) Reset() {
	*x = CreateSnippetResponse{}
	mi := &file_v1_snippet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSnippetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSnippetResponse) ProtoMessage() {}

func (x *CreateSnippetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSnippetResponse.ProtoReflect.Descriptor instead.
func (*CreateSnippetResponse) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{4}
}

func (x *CreateSnippetResponse) GetSnippet() *Snippet {
	if x != nil {
		return x.Snippet
	}
	return nil
}

type GetSnippetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSnippetRequest) Reset() {
	*x = GetSnippetRequest{}
	mi := &file_v1_snippet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnippetRequest) ProtoMessage() {}

func (x *GetSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnippetRequest.ProtoReflect.Descriptor instead.
func (*GetSnippetRequest) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{5}
}

func (x *GetSnippetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSnippetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Snippet       *Snippet               `protobuf:"bytes,1,opt,name=snippet,proto3" json:"snippet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSnippetResponse) Reset() {
	*x = GetSnippetResponse{}
	mi := &file_v1_snippet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSnippetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSnippetResponse) ProtoMessage() {}

func (x *GetSnippetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSnippetResponse.ProtoReflect.Descriptor instead.
func (*GetSnippetResponse) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{6}
}

func (x *GetSnippetResponse) GetSnippet() *Snippet {
	if x != nil {
		return x.Snippet
	}
	return nil
}

type ListMySnippetsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          *int32                 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Size          *int32                 `protobuf:"varint,2,opt,name=size,proto3,oneof" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMySnippetsRequest) Reset() {
	*x = ListMySnippetsRequest{}
	mi := &file_v1_snippet_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMySnippetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMySnippetsRequest) ProtoMessage() {}

func (x *ListMySnippetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMySnippetsRequest.ProtoReflect.Descriptor instead.
func (*ListMySnippetsRequest) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{7}
}

func (x *ListMySnippetsRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListMySnippetsRequest) GetSize() int32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

type ListMySnippetsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Snippets      []*Snippet             `protobuf:"bytes,1,rep,name=snippets,proto3" json:"snippets,omitempty"`
	Pagination    *PaginationResponse    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMySnippetsResponse) Reset() {
	*x = ListMySnippetsResponse{}
	mi := &file_v1_snippet_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMySnippetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMySnippetsResponse) ProtoMessage() {}

func (x *ListMySnippetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMySnippetsResponse.ProtoReflect.Descriptor instead.
func (*ListMySnippetsResponse) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{8}
}

func (x *ListMySnippetsResponse) GetSnippets() []*Snippet {
	if x != nil {
		return x.Snippets
	}
	return nil
}

func (x *ListMySnippetsResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListPublicSnippetsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          *int32                 `protobuf:"varint,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Size          *int32                 `protobuf:"varint,2,opt,name=size,proto3,oneof" json:"size,omitempty"`
	OwnerId       *string                `protobuf:"bytes,3,opt,name=owner_id,json=ownerId,proto3,oneof" json:"owner_id,omitempty"`
	ExcludeMine   *bool                  `protobuf:"varint,4,opt,name=exclude_mine,json=excludeMine,proto3,oneof" json:"exclude_mine,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPublicSnippetsRequest) Reset() {
	*x = ListPublicSnippetsRequest{}
	mi := &file_v1_snippet_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPublicSnippetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicSnippetsRequest) ProtoMessage() {}

func (x *ListPublicSnippetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicSnippetsRequest.ProtoReflect.Descriptor instead.
func (*ListPublicSnippetsRequest) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{9}
}

func (x *ListPublicSnippetsRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *ListPublicSnippetsRequest) GetSize() int32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *ListPublicSnippetsRequest) GetOwnerId() string {
	if x != nil && x.OwnerId != nil {
		return *x.OwnerId
	}
	return ""
}

func (x *ListPublicSnippetsRequest) GetExcludeMine() bool {
	if x != nil && x.ExcludeMine != nil {
		return *x.ExcludeMine
	}
	return false
}

type ListPublicSnippetsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Snippets      []*Snippet             `protobuf:"bytes,1,rep,name=snippets,proto3" json:"snippets,omitempty"`
	Pagination    *PaginationResponse    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPublicSnippetsResponse) Reset() {
	*x = ListPublicSnippetsResponse{}
	mi := &file_v1_snippet_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPublicSnippetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPublicSnippetsResponse) ProtoMessage() {}

func (x *ListPublicSnippetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPublicSnippetsResponse.ProtoReflect.Descriptor instead.
func (*ListPublicSnippetsResponse) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{10}
}

func (x *ListPublicSnippetsResponse) GetSnippets() []*Snippet {
	if x != nil {
		return x.Snippets
	}
	return nil
}

func (x *ListPublicSnippetsResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type UpdateSnippetTags struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       []string               `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSnippetTags) Reset() {
	*x = UpdateSnippetTags{}
	mi := &file_v1_snippet_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSnippetTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnippetTags) ProtoMessage() {}

func (x *UpdateSnippetTags) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnippetTags.ProtoReflect.Descriptor instead.
func (*UpdateSnippetTags) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{11}
}

func (x *UpdateSnippetTags) GetContent() []string {
	if x != nil {
		return x.Content
	}
	return nil
}

type UpdateSnippetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         *string                `protobuf:"bytes,2,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Content       *string                `protobuf:"bytes,3,opt,name=content,proto3,oneof" json:"content,omitempty"`
	LanguageHint  *string                `protobuf:"bytes,4,opt,name=language_hint,json=languageHint,proto3,oneof" json:"language_hint,omitempty"`
	IsPublic      *bool                  `protobuf:"varint,5,opt,name=is_public,json=isPublic,proto3,oneof" json:"is_public,omitempty"`
	Tags          *UpdateSnippetTags     `protobuf:"bytes,6,opt,name=tags,proto3,oneof" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSnippetRequest) Reset() {
	*x = UpdateSnippetRequest{}
	mi := &file_v1_snippet_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnippetRequest) ProtoMessage() {}

func (x *UpdateSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnippetRequest.ProtoReflect.Descriptor instead.
func (*UpdateSnippetRequest) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateSnippetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSnippetRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdateSnippetRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *UpdateSnippetRequest) GetLanguageHint() string {
	if x != nil && x.LanguageHint != nil {
		return *x.LanguageHint
	}
	return ""
}

func (x *UpdateSnippetRequest) GetIsPublic() bool {
	if x != nil && x.IsPublic != nil {
		return *x.IsPublic
	}
	return false
}

func (x *UpdateSnippetRequest) GetTags() *UpdateSnippetTags {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateSnippetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Snippet       *Snippet               `protobuf:"bytes,1,opt,name=snippet,proto3" json:"snippet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSnippetResponse) Reset() {
	*x = UpdateSnippetResponse{}
	mi := &file_v1_snippet_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSnippetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSnippetResponse) ProtoMessage() {}

func (x *UpdateSnippetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSnippetResponse.ProtoReflect.Descriptor instead.
func (*UpdateSnippetResponse) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateSnippetResponse) GetSnippet() *Snippet {
	if x != nil {
		return x.Snippet
	}
	return nil
}

type DeleteSnippetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSnippetRequest) Reset() {
	*x = DeleteSnippetRequest{}
	mi := &file_v1_snippet_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSnippetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnippetRequest) ProtoMessage() {}

func (x *DeleteSnippetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_snippet_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnippetRequest.ProtoReflect.Descriptor instead.
func (*DeleteSnippetRequest) Descriptor() ([]byte, []int) {
	return file_v1_snippet_proto_rawDescGZIP(), []int{14}
}

func (x *DeleteSnippetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_snippet_proto protoreflect.FileDescriptor

const file_v1_snippet_proto_rawDesc = "" +
	"\n" +
	"\x10v1/snippet.proto\x12\n" +
	"snippet.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/api/annotations.proto\":\n" +
	"\fSnippetOwner\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\"\xc5\x02\n" +
	"\aSnippet\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12.\n" +
	"\x05owner\x18\x02 \x01(\v2\x18.snippet.v1.SnippetOwnerR\x05owner\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\x12#\n" +
	"\rlanguage_hint\x18\x05 \x01(\tR\flanguageHint\x12\x1b\n" +
	"\tis_public\x18\x06 \x01(\bR\bisPublic\x12\x12\n" +
	"\x04tags\x18\a \x03(\tR\x04tags\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"\x8d\x01\n" +
	"\x12PaginationResponse\x12\x12\n" +
	"\x04size\x18\x01 \x01(\x05R\x04size\x12\x1f\n" +
	"\vtotal_items\x18\x02 \x01(\x05R\n" +
	"totalItems\x12!\n" +
	"\fcurrent_page\x18\x03 \x01(\x05R\vcurrentPage\x12\x1f\n" +
	"\vtotal_pages\x18\x04 \x01(\x05R\n" +
	"totalPages\"\xc6\x01\n" +
	"\x14CreateSnippetRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12(\n" +
	"\rlanguage_hint\x18\x03 \x01(\tH\x00R\flanguageHint\x88\x01\x01\x12 \n" +
	"\tis_public\x18\x04 \x01(\bH\x01R\bisPublic\x88\x01\x01\x12\x12\n" +
	"\x04tags\x18\x05 \x03(\tR\x04tagsB\x10\n" +
	"\x0e_language_hintB\f\n" +
	"\n" +
	"_is_public\"F\n" +
	"\x15CreateSnippetResponse\x12-\n" +
	"\asnippet\x18\x01 \x01(\v2\x13.snippet.v1.SnippetR\asnippet\"#\n" +
	"\x11GetSnippetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"C\n" +
	"\x12GetSnippetResponse\x12-\n" +
	"\asnippet\x18\x01 \x01(\v2\x13.snippet.v1.SnippetR\asnippet\"[\n" +
	"\x15ListMySnippetsRequest\x12\x17\n" +
	"\x04page\x18\x01 \x01(\x05H\x00R\x04page\x88\x01\x01\x12\x17\n" +
	"\x04size\x18\x02 \x01(\x05H\x01R\x04size\x88\x01\x01B\a\n" +
	"\x05_pageB\a\n" +
	"\x05_size\"\x89\x01\n" +
	"\x16ListMySnippetsResponse\x12/\n" +
	"\bsnippets\x18\x01 \x03(\v2\x13.snippet.v1.SnippetR\bsnippets\x12>\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x1e.snippet.v1.PaginationResponseR\n" +
	"pagination\"\xc5\x01\n" +
	"\x19ListPublicSnippetsRequest\x12\x17\n" +
	"\x04page\x18\x01 \x01(\x05H\x00R\x04page\x88\x01\x01\x12\x17\n" +
	"\x04size\x18\x02 \x01(\x05H\x01R\x04size\x88\x01\x01\x12\x1e\n" +
	"\bowner_id\x18\x03 \x01(\tH\x02R\aownerId\x88\x01\x01\x12&\n" +
	"\fexclude_mine\x18\x04 \x01(\bH\x03R\vexcludeMine\x88\x01\x01B\a\n" +
	"\x05_pageB\a\n" +
	"\x05_sizeB\v\n" +
	"\t_owner_idB\x0f\n" +
	"\r_exclude_mine\"\x8d\x01\n" +
	"\x1aListPublicSnippetsResponse\x12/\n" +
	"\bsnippets\x18\x01 \x03(\v2\x13.snippet.v1.SnippetR\bsnippets\x12>\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x1e.snippet.v1.PaginationResponseR\n" +
	"pagination\"-\n" +
	"\x11UpdateSnippetTags\x12\x18\n" +
	"\acontent\x18\x01 \x03(\tR\acontent\"\xa3\x02\n" +
	"\x14UpdateSnippetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x19\n" +
	"\x05title\x18\x02 \x01(\tH\x00R\x05title\x88\x01\x01\x12\x1d\n" +
	"\acontent\x18\x03 \x01(\tH\x01R\acontent\x88\x01\x01\x12(\n" +
	"\rlanguage_hint\x18\x04 \x01(\tH\x02R\flanguageHint\x88\x01\x01\x12 \n" +
	"\tis_public\x18\x05 \x01(\bH\x03R\bisPublic\x88\x01\x01\x126\n" +
	"\x04tags\x18\x06 \x01(\v2\x1d.snippet.v1.UpdateSnippetTagsH\x04R\x04tags\x88\x01\x01B\b\n" +
	"\x06_titleB\n" +
	"\n" +
	"\b_contentB\x10\n" +
	"\x0e_language_hintB\f\n" +
	"\n" +
	"_is_publicB\a\n" +
	"\x05_tags\"F\n" +
	"\x15UpdateSnippetResponse\x12-\n" +
	"\asnippet\x18\x01 \x01(\v2\x13.snippet.v1.SnippetR\asnippet\"&\n" +
	"\x14DeleteSnippetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xc6\x05\n" +
	"\x0eSnippetService\x12q\n" +
	"\rCreateSnippet\x12 .snippet.v1.CreateSnippetRequest\x1a!.snippet.v1.CreateSnippetResponse\"\x1b\x82\xd3\xe4\x93\x02\x15:\x01*\"\x10/api/v1/snippets\x12j\n" +
	"\n" +
	"GetSnippet\x12\x1d.snippet.v1.GetSnippetRequest\x1a\x1e.snippet.v1.GetSnippetResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/api/v1/snippets/{id}\x12t\n" +
	"\x0eListMySnippets\x12!.snippet.v1.ListMySnippetsRequest\x1a\".snippet.v1.ListMySnippetsResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/api/v1/me/snippets\x12}\n" +
	"\x12ListPublicSnippets\x12%.snippet.v1.ListPublicSnippetsRequest\x1a&.snippet.v1.ListPublicSnippetsResponse\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/api/v1/snippets\x12v\n" +
	"\rUpdateSnippet\x12 .snippet.v1.UpdateSnippetRequest\x1a!.snippet.v1.UpdateSnippetResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*2\x15/api/v1/snippets/{id}\x12h\n" +
	"\rDeleteSnippet\x12 .snippet.v1.DeleteSnippetRequest\x1a\x16.google.protobuf.Empty\"\x1d\x82\xd3\xe4\x93\x02\x17*\x15/api/v1/snippets/{id}B3Z1github.com/kakitomeru/snippet/pkg/pb/v1;snippetpbb\x06proto3"

var (
	file_v1_snippet_proto_rawDescOnce sync.Once
	file_v1_snippet_proto_rawDescData []byte
)

func file_v1_snippet_proto_rawDescGZIP() []byte {
	file_v1_snippet_proto_rawDescOnce.Do(func() {
		file_v1_snippet_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_snippet_proto_rawDesc), len(file_v1_snippet_proto_rawDesc)))
	})
	return file_v1_snippet_proto_rawDescData
}

var file_v1_snippet_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_v1_snippet_proto_goTypes = []any{
	(*SnippetOwner)(nil),               // 0: snippet.v1.SnippetOwner
	(*Snippet)(nil),                    // 1: snippet.v1.Snippet
	(*PaginationResponse)(nil),         // 2: snippet.v1.PaginationResponse
	(*CreateSnippetRequest)(nil),       // 3: snippet.v1.CreateSnippetRequest
	(*CreateSnippetResponse)(nil),      // 4: snippet.v1.CreateSnippetResponse
	(*GetSnippetRequest)(nil),          // 5: snippet.v1.GetSnippetRequest
	(*GetSnippetResponse)(nil),         // 6: snippet.v1.GetSnippetResponse
	(*ListMySnippetsRequest)(nil),      // 7: snippet.v1.ListMySnippetsRequest
	(*ListMySnippetsResponse)(nil),     // 8: snippet.v1.ListMySnippetsResponse
	(*ListPublicSnippetsRequest)(nil),  // 9: snippet.v1.ListPublicSnippetsRequest
	(*ListPublicSnippetsResponse)(nil), // 10: snippet.v1.ListPublicSnippetsResponse
	(*UpdateSnippetTags)(nil),          // 11: snippet.v1.UpdateSnippetTags
	(*UpdateSnippetRequest)(nil),       // 12: snippet.v1.UpdateSnippetRequest
	(*UpdateSnippetResponse)(nil),      // 13: snippet.v1.UpdateSnippetResponse
	(*DeleteSnippetRequest)(nil),       // 14: snippet.v1.DeleteSnippetRequest
	(*timestamppb.Timestamp)(nil),      // 15: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),              // 16: google.protobuf.Empty
}
var file_v1_snippet_proto_depIdxs = []int32{
	0,  // 0: snippet.v1.Snippet.owner:type_name -> snippet.v1.SnippetOwner
	15, // 1: snippet.v1.Snippet.created_at:type_name -> google.protobuf.Timestamp
	15, // 2: snippet.v1.Snippet.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 3: snippet.v1.CreateSnippetResponse.snippet:type_name -> snippet.v1.Snippet
	1,  // 4: snippet.v1.GetSnippetResponse.snippet:type_name -> snippet.v1.Snippet
	1,  // 5: snippet.v1.ListMySnippetsResponse.snippets:type_name -> snippet.v1.Snippet
	2,  // 6: snippet.v1.ListMySnippetsResponse.pagination:type_name -> snippet.v1.PaginationResponse
	1,  // 7: snippet.v1.ListPublicSnippetsResponse.snippets:type_name -> snippet.v1.Snippet
	2,  // 8: snippet.v1.ListPublicSnippetsResponse.pagination:type_name -> snippet.v1.PaginationResponse
	11, // 9: snippet.v1.UpdateSnippetRequest.tags:type_name -> snippet.v1.UpdateSnippetTags
	1,  // 10: snippet.v1.UpdateSnippetResponse.snippet:type_name -> snippet.v1.Snippet
	3,  // 11: snippet.v1.SnippetService.CreateSnippet:input_type -> snippet.v1.CreateSnippetRequest
	5,  // 12: snippet.v1.SnippetService.GetSnippet:input_type -> snippet.v1.GetSnippetRequest
	7,  // 13: snippet.v1.SnippetService.ListMySnippets:input_type -> snippet.v1.ListMySnippetsRequest
	9,  // 14: snippet.v1.SnippetService.ListPublicSnippets:input_type -> snippet.v1.ListPublicSnippetsRequest
	12, // 15: snippet.v1.SnippetService.UpdateSnippet:input_type -> snippet.v1.UpdateSnippetRequest
	14, // 16: snippet.v1.SnippetService.DeleteSnippet:input_type -> snippet.v1.DeleteSnippetRequest
	4,  // 17: snippet.v1.SnippetService.CreateSnippet:output_type -> snippet.v1.CreateSnippetResponse
	6,  // 18: snippet.v1.SnippetService.GetSnippet:output_type -> snippet.v1.GetSnippetResponse
	8,  // 19: snippet.v1.SnippetService.ListMySnippets:output_type -> snippet.v1.ListMySnippetsResponse
	10, // 20: snippet.v1.SnippetService.ListPublicSnippets:output_type -> snippet.v1.ListPublicSnippetsResponse
	13, // 21: snippet.v1.SnippetService.UpdateSnippet:output_type -> snippet.v1.UpdateSnippetResponse
	16, // 22: snippet.v1.SnippetService.DeleteSnippet:output_type -> google.protobuf.Empty
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_v1_snippet_proto_init() }
func file_v1_snippet_proto_init() {
	if File_v1_snippet_proto != nil {
		return
	}
	file_v1_snippet_proto_msgTypes[3].OneofWrappers = []any{}
	file_v1_snippet_proto_msgTypes[7].OneofWrappers = []any{}
	file_v1_snippet_proto_msgTypes[9].OneofWrappers = []any{}
	file_v1_snippet_proto_msgTypes[12].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_snippet_proto_rawDesc), len(file_v1_snippet_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_snippet_proto_goTypes,
		DependencyIndexes: file_v1_snippet_proto_depIdxs,
		MessageInfos:      file_v1_snippet_proto_msgTypes,
	}.Build()
	File_v1_snippet_proto = out.File
	file_v1_snippet_proto_goTypes = nil
	file_v1_snippet_proto_depIdxs = nil
}
