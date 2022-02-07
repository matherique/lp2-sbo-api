// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: internal/proto/book/book.proto

package book

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

type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BookList) Reset() {
	*x = BookList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_book_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookList) ProtoMessage() {}

func (x *BookList) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_book_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookList.ProtoReflect.Descriptor instead.
func (*BookList) Descriptor() ([]byte, []int) {
	return file_internal_proto_book_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookList) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_book_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_book_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_internal_proto_book_book_proto_rawDescGZIP(), []int{1}
}

type FindData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *FindData) Reset() {
	*x = FindData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_book_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindData) ProtoMessage() {}

func (x *FindData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_book_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindData.ProtoReflect.Descriptor instead.
func (*FindData) Descriptor() ([]byte, []int) {
	return file_internal_proto_book_book_proto_rawDescGZIP(), []int{2}
}

func (x *FindData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FindData) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Photo       string `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
	AuthorId    int32  `protobuf:"varint,4,opt,name=authorId,proto3" json:"authorId,omitempty"`
	ShortDesc   string `protobuf:"bytes,5,opt,name=shortDesc,proto3" json:"shortDesc,omitempty"`
	Longdesc    string `protobuf:"bytes,6,opt,name=longdesc,proto3" json:"longdesc,omitempty"`
	CategoryId  int32  `protobuf:"varint,7,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	PaperBack   int32  `protobuf:"varint,8,opt,name=paperBack,proto3" json:"paperBack,omitempty"`
	PublisherId int32  `protobuf:"varint,9,opt,name=publisherId,proto3" json:"publisherId,omitempty"`
	CreatedAt   string `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,11,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_book_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_book_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_internal_proto_book_book_proto_rawDescGZIP(), []int{3}
}

func (x *Book) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Book) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Book) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *Book) GetShortDesc() string {
	if x != nil {
		return x.ShortDesc
	}
	return ""
}

func (x *Book) GetLongdesc() string {
	if x != nil {
		return x.Longdesc
	}
	return ""
}

func (x *Book) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Book) GetPaperBack() int32 {
	if x != nil {
		return x.PaperBack
	}
	return 0
}

func (x *Book) GetPublisherId() int32 {
	if x != nil {
		return x.PublisherId
	}
	return 0
}

func (x *Book) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Book) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Photo       string `protobuf:"bytes,2,opt,name=photo,proto3" json:"photo,omitempty"`
	AuthorId    int32  `protobuf:"varint,3,opt,name=authorId,proto3" json:"authorId,omitempty"`
	ShortDesc   string `protobuf:"bytes,4,opt,name=shortDesc,proto3" json:"shortDesc,omitempty"`
	Longdesc    string `protobuf:"bytes,5,opt,name=longdesc,proto3" json:"longdesc,omitempty"`
	CategoryId  int32  `protobuf:"varint,6,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	PaperBack   int32  `protobuf:"varint,7,opt,name=paperBack,proto3" json:"paperBack,omitempty"`
	PublisherId int32  `protobuf:"varint,8,opt,name=publisherId,proto3" json:"publisherId,omitempty"`
}

func (x *CreateData) Reset() {
	*x = CreateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_book_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateData) ProtoMessage() {}

func (x *CreateData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_book_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateData.ProtoReflect.Descriptor instead.
func (*CreateData) Descriptor() ([]byte, []int) {
	return file_internal_proto_book_book_proto_rawDescGZIP(), []int{4}
}

func (x *CreateData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateData) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateData) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *CreateData) GetShortDesc() string {
	if x != nil {
		return x.ShortDesc
	}
	return ""
}

func (x *CreateData) GetLongdesc() string {
	if x != nil {
		return x.Longdesc
	}
	return ""
}

func (x *CreateData) GetCategoryId() int32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *CreateData) GetPaperBack() int32 {
	if x != nil {
		return x.PaperBack
	}
	return 0
}

func (x *CreateData) GetPublisherId() int32 {
	if x != nil {
		return x.PublisherId
	}
	return 0
}

type UpdateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *UpdateData) Reset() {
	*x = UpdateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_book_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateData) ProtoMessage() {}

func (x *UpdateData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_book_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateData.ProtoReflect.Descriptor instead.
func (*UpdateData) Descriptor() ([]byte, []int) {
	return file_internal_proto_book_book_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateData) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_internal_proto_book_book_proto protoreflect.FileDescriptor

var file_internal_proto_book_book_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x27, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x38, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0xb2, 0x02, 0x0a,
	0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x6e,
	0x67, 0x64, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x6e,
	0x67, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x70, 0x65, 0x72, 0x42, 0x61,
	0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x70, 0x65, 0x72, 0x42,
	0x61, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xec, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44,
	0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x70, 0x65, 0x72, 0x42, 0x61, 0x63, 0x6b, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x70, 0x65, 0x72, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x37, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x32, 0x89, 0x01, 0x0a, 0x0b, 0x42, 0x6f,
	0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x1e, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x1a, 0x0a, 0x04, 0x46, 0x69, 0x6e,
	0x64, 0x12, 0x09, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x05, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x65, 0x72, 0x69, 0x71, 0x75, 0x65, 0x2f, 0x6c,
	0x70, 0x32, 0x2d, 0x73, 0x62, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_book_book_proto_rawDescOnce sync.Once
	file_internal_proto_book_book_proto_rawDescData = file_internal_proto_book_book_proto_rawDesc
)

func file_internal_proto_book_book_proto_rawDescGZIP() []byte {
	file_internal_proto_book_book_proto_rawDescOnce.Do(func() {
		file_internal_proto_book_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_book_book_proto_rawDescData)
	})
	return file_internal_proto_book_book_proto_rawDescData
}

var file_internal_proto_book_book_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_proto_book_book_proto_goTypes = []interface{}{
	(*BookList)(nil),   // 0: BookList
	(*Empty)(nil),      // 1: Empty
	(*FindData)(nil),   // 2: FindData
	(*Book)(nil),       // 3: Book
	(*CreateData)(nil), // 4: CreateData
	(*UpdateData)(nil), // 5: UpdateData
}
var file_internal_proto_book_book_proto_depIdxs = []int32{
	3, // 0: BookList.books:type_name -> Book
	3, // 1: UpdateData.book:type_name -> Book
	4, // 2: BookService.Create:input_type -> CreateData
	5, // 3: BookService.Update:input_type -> UpdateData
	1, // 4: BookService.FindAll:input_type -> Empty
	2, // 5: BookService.Find:input_type -> FindData
	3, // 6: BookService.Create:output_type -> Book
	3, // 7: BookService.Update:output_type -> Book
	0, // 8: BookService.FindAll:output_type -> BookList
	3, // 9: BookService.Find:output_type -> Book
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_book_book_proto_init() }
func file_internal_proto_book_book_proto_init() {
	if File_internal_proto_book_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_book_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookList); i {
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
		file_internal_proto_book_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_internal_proto_book_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindData); i {
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
		file_internal_proto_book_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_internal_proto_book_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateData); i {
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
		file_internal_proto_book_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateData); i {
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
			RawDescriptor: file_internal_proto_book_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_book_book_proto_goTypes,
		DependencyIndexes: file_internal_proto_book_book_proto_depIdxs,
		MessageInfos:      file_internal_proto_book_book_proto_msgTypes,
	}.Build()
	File_internal_proto_book_book_proto = out.File
	file_internal_proto_book_book_proto_rawDesc = nil
	file_internal_proto_book_book_proto_goTypes = nil
	file_internal_proto_book_book_proto_depIdxs = nil
}