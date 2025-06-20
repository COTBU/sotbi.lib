// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: counterparty/counterparty.proto

package counterparty

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CounterpartiesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Counterparty        `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CounterpartiesResponse) Reset() {
	*x = CounterpartiesResponse{}
	mi := &file_counterparty_counterparty_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CounterpartiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterpartiesResponse) ProtoMessage() {}

func (x *CounterpartiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_counterparty_counterparty_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterpartiesResponse.ProtoReflect.Descriptor instead.
func (*CounterpartiesResponse) Descriptor() ([]byte, []int) {
	return file_counterparty_counterparty_proto_rawDescGZIP(), []int{0}
}

func (x *CounterpartiesResponse) GetItems() []*Counterparty {
	if x != nil {
		return x.Items
	}
	return nil
}

type Link struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Link) Reset() {
	*x = Link{}
	mi := &file_counterparty_counterparty_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_counterparty_counterparty_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_counterparty_counterparty_proto_rawDescGZIP(), []int{1}
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Контрагент
type Counterparty struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ID              int64                  `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Kind            bool                   `protobuf:"varint,3,opt,name=Kind,json=kind,proto3" json:"Kind,omitempty"`                   // ЮЛ/ФЛ
	Name            string                 `protobuf:"bytes,6,opt,name=Name,json=name,proto3" json:"Name,omitempty"`                    // Краткое наименование
	FullName        *string                `protobuf:"bytes,9,opt,name=FullName,json=full_name,proto3,oneof" json:"FullName,omitempty"` // Полное наименование
	Alias           *string                `protobuf:"bytes,12,opt,name=Alias,json=alias,proto3,oneof" json:"Alias,omitempty"`          // Погоняло
	INN             string                 `protobuf:"bytes,15,opt,name=INN,json=inn,proto3" json:"INN,omitempty"`
	KPP             *string                `protobuf:"bytes,18,opt,name=KPP,json=kpp,proto3,oneof" json:"KPP,omitempty"`
	OGRN            *string                `protobuf:"bytes,21,opt,name=OGRN,json=ogrn,proto3,oneof" json:"OGRN,omitempty"`
	Address         *string                `protobuf:"bytes,24,opt,name=Address,json=address,proto3,oneof" json:"Address,omitempty"`                          // Юридический адрес
	PostAddress     *string                `protobuf:"bytes,27,opt,name=PostAddress,json=post_address,proto3,oneof" json:"PostAddress,omitempty"`             // Почтовый адрес
	Contacts        *string                `protobuf:"bytes,30,opt,name=Contacts,json=contacts,proto3,oneof" json:"Contacts,omitempty"`                       // Контакты
	Email           *string                `protobuf:"bytes,33,opt,name=Email,json=email,proto3,oneof" json:"Email,omitempty"`                                // Электронная почта
	CEO             *string                `protobuf:"bytes,36,opt,name=CEO,json=ceo,proto3,oneof" json:"CEO,omitempty"`                                      // Руководитель
	ChiefAccountant *string                `protobuf:"bytes,39,opt,name=ChiefAccountant,json=chief_accountant,proto3,oneof" json:"ChiefAccountant,omitempty"` // Главный бухгалтер
	Links           []*Link                `protobuf:"bytes,42,rep,name=links,proto3" json:"links,omitempty"`                                                 // Ссылки
	Access          []string               `protobuf:"bytes,45,rep,name=access,proto3" json:"access,omitempty"`                                               // Доступы
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,48,opt,name=CreatedAt,json=created_at,proto3,oneof" json:"CreatedAt,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=UpdatedAt,json=updated_at,proto3,oneof" json:"UpdatedAt,omitempty"`
	DeletedAt       *timestamppb.Timestamp `protobuf:"bytes,54,opt,name=DeletedAt,json=deleted_at,proto3,oneof" json:"DeletedAt,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Counterparty) Reset() {
	*x = Counterparty{}
	mi := &file_counterparty_counterparty_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Counterparty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counterparty) ProtoMessage() {}

func (x *Counterparty) ProtoReflect() protoreflect.Message {
	mi := &file_counterparty_counterparty_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counterparty.ProtoReflect.Descriptor instead.
func (*Counterparty) Descriptor() ([]byte, []int) {
	return file_counterparty_counterparty_proto_rawDescGZIP(), []int{2}
}

func (x *Counterparty) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Counterparty) GetKind() bool {
	if x != nil {
		return x.Kind
	}
	return false
}

func (x *Counterparty) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Counterparty) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *Counterparty) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

func (x *Counterparty) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

func (x *Counterparty) GetKPP() string {
	if x != nil && x.KPP != nil {
		return *x.KPP
	}
	return ""
}

func (x *Counterparty) GetOGRN() string {
	if x != nil && x.OGRN != nil {
		return *x.OGRN
	}
	return ""
}

func (x *Counterparty) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *Counterparty) GetPostAddress() string {
	if x != nil && x.PostAddress != nil {
		return *x.PostAddress
	}
	return ""
}

func (x *Counterparty) GetContacts() string {
	if x != nil && x.Contacts != nil {
		return *x.Contacts
	}
	return ""
}

func (x *Counterparty) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *Counterparty) GetCEO() string {
	if x != nil && x.CEO != nil {
		return *x.CEO
	}
	return ""
}

func (x *Counterparty) GetChiefAccountant() string {
	if x != nil && x.ChiefAccountant != nil {
		return *x.ChiefAccountant
	}
	return ""
}

func (x *Counterparty) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Counterparty) GetAccess() []string {
	if x != nil {
		return x.Access
	}
	return nil
}

func (x *Counterparty) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Counterparty) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Counterparty) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

// Добавить контрагента
type AddCounterpartyRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Kind            bool                   `protobuf:"varint,3,opt,name=Kind,json=kind,proto3" json:"Kind,omitempty"`                   // ЮЛ/ФЛ
	Name            string                 `protobuf:"bytes,6,opt,name=Name,json=name,proto3" json:"Name,omitempty"`                    // Краткое наименование
	FullName        *string                `protobuf:"bytes,9,opt,name=FullName,json=full_name,proto3,oneof" json:"FullName,omitempty"` // Полное наименование
	Alias           *string                `protobuf:"bytes,12,opt,name=Alias,json=alias,proto3,oneof" json:"Alias,omitempty"`          // Погоняло
	INN             string                 `protobuf:"bytes,15,opt,name=INN,json=inn,proto3" json:"INN,omitempty"`
	KPP             *string                `protobuf:"bytes,18,opt,name=KPP,json=kpp,proto3,oneof" json:"KPP,omitempty"`
	OGRN            *string                `protobuf:"bytes,21,opt,name=OGRN,json=ogrn,proto3,oneof" json:"OGRN,omitempty"`
	Address         *string                `protobuf:"bytes,24,opt,name=Address,json=address,proto3,oneof" json:"Address,omitempty"`                          // Юридический адрес
	PostAddress     *string                `protobuf:"bytes,27,opt,name=PostAddress,json=post_address,proto3,oneof" json:"PostAddress,omitempty"`             // Почтовый адрес
	Contacts        *string                `protobuf:"bytes,30,opt,name=Contacts,json=contacts,proto3,oneof" json:"Contacts,omitempty"`                       // Контакты
	Email           *string                `protobuf:"bytes,33,opt,name=Email,json=email,proto3,oneof" json:"Email,omitempty"`                                // Электронная почта
	CEO             *string                `protobuf:"bytes,36,opt,name=CEO,json=ceo,proto3,oneof" json:"CEO,omitempty"`                                      // Руководитель
	ChiefAccountant *string                `protobuf:"bytes,39,opt,name=ChiefAccountant,json=chief_accountant,proto3,oneof" json:"ChiefAccountant,omitempty"` // Главный бухгалтер
	Links           []*Link                `protobuf:"bytes,42,rep,name=links,proto3" json:"links,omitempty"`                                                 // Ссылки
	Access          []string               `protobuf:"bytes,45,rep,name=access,proto3" json:"access,omitempty"`                                               // Доступы
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AddCounterpartyRequest) Reset() {
	*x = AddCounterpartyRequest{}
	mi := &file_counterparty_counterparty_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddCounterpartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCounterpartyRequest) ProtoMessage() {}

func (x *AddCounterpartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counterparty_counterparty_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCounterpartyRequest.ProtoReflect.Descriptor instead.
func (*AddCounterpartyRequest) Descriptor() ([]byte, []int) {
	return file_counterparty_counterparty_proto_rawDescGZIP(), []int{3}
}

func (x *AddCounterpartyRequest) GetKind() bool {
	if x != nil {
		return x.Kind
	}
	return false
}

func (x *AddCounterpartyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddCounterpartyRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *AddCounterpartyRequest) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

func (x *AddCounterpartyRequest) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

func (x *AddCounterpartyRequest) GetKPP() string {
	if x != nil && x.KPP != nil {
		return *x.KPP
	}
	return ""
}

func (x *AddCounterpartyRequest) GetOGRN() string {
	if x != nil && x.OGRN != nil {
		return *x.OGRN
	}
	return ""
}

func (x *AddCounterpartyRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *AddCounterpartyRequest) GetPostAddress() string {
	if x != nil && x.PostAddress != nil {
		return *x.PostAddress
	}
	return ""
}

func (x *AddCounterpartyRequest) GetContacts() string {
	if x != nil && x.Contacts != nil {
		return *x.Contacts
	}
	return ""
}

func (x *AddCounterpartyRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *AddCounterpartyRequest) GetCEO() string {
	if x != nil && x.CEO != nil {
		return *x.CEO
	}
	return ""
}

func (x *AddCounterpartyRequest) GetChiefAccountant() string {
	if x != nil && x.ChiefAccountant != nil {
		return *x.ChiefAccountant
	}
	return ""
}

func (x *AddCounterpartyRequest) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *AddCounterpartyRequest) GetAccess() []string {
	if x != nil {
		return x.Access
	}
	return nil
}

// Отредактировать контрагента
type UpdateCounterpartyRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind            bool                   `protobuf:"varint,3,opt,name=Kind,json=kind,proto3" json:"Kind,omitempty"`                   // ЮЛ/ФЛ
	Name            string                 `protobuf:"bytes,6,opt,name=Name,json=name,proto3" json:"Name,omitempty"`                    // Краткое наименование
	FullName        *string                `protobuf:"bytes,9,opt,name=FullName,json=full_name,proto3,oneof" json:"FullName,omitempty"` // Полное наименование
	Alias           *string                `protobuf:"bytes,12,opt,name=Alias,json=alias,proto3,oneof" json:"Alias,omitempty"`          // Погоняло
	INN             string                 `protobuf:"bytes,15,opt,name=INN,json=inn,proto3" json:"INN,omitempty"`
	KPP             *string                `protobuf:"bytes,18,opt,name=KPP,json=kpp,proto3,oneof" json:"KPP,omitempty"`
	OGRN            *string                `protobuf:"bytes,21,opt,name=OGRN,json=ogrn,proto3,oneof" json:"OGRN,omitempty"`
	Address         *string                `protobuf:"bytes,24,opt,name=Address,json=address,proto3,oneof" json:"Address,omitempty"`                          // Юридический адрес
	PostAddress     *string                `protobuf:"bytes,27,opt,name=PostAddress,json=post_address,proto3,oneof" json:"PostAddress,omitempty"`             // Почтовый адрес
	Contacts        *string                `protobuf:"bytes,30,opt,name=Contacts,json=contacts,proto3,oneof" json:"Contacts,omitempty"`                       // Контакты
	Email           *string                `protobuf:"bytes,33,opt,name=Email,json=email,proto3,oneof" json:"Email,omitempty"`                                // Электронная почта
	CEO             *string                `protobuf:"bytes,36,opt,name=CEO,json=ceo,proto3,oneof" json:"CEO,omitempty"`                                      // Руководитель
	ChiefAccountant *string                `protobuf:"bytes,39,opt,name=ChiefAccountant,json=chief_accountant,proto3,oneof" json:"ChiefAccountant,omitempty"` // Главный бухгалтер
	Links           []*Link                `protobuf:"bytes,42,rep,name=links,proto3" json:"links,omitempty"`                                                 // Ссылки
	Access          []string               `protobuf:"bytes,45,rep,name=access,proto3" json:"access,omitempty"`                                               // Доступы
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdateCounterpartyRequest) Reset() {
	*x = UpdateCounterpartyRequest{}
	mi := &file_counterparty_counterparty_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCounterpartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCounterpartyRequest) ProtoMessage() {}

func (x *UpdateCounterpartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_counterparty_counterparty_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCounterpartyRequest.ProtoReflect.Descriptor instead.
func (*UpdateCounterpartyRequest) Descriptor() ([]byte, []int) {
	return file_counterparty_counterparty_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCounterpartyRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCounterpartyRequest) GetKind() bool {
	if x != nil {
		return x.Kind
	}
	return false
}

func (x *UpdateCounterpartyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetKPP() string {
	if x != nil && x.KPP != nil {
		return *x.KPP
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetOGRN() string {
	if x != nil && x.OGRN != nil {
		return *x.OGRN
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetPostAddress() string {
	if x != nil && x.PostAddress != nil {
		return *x.PostAddress
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetContacts() string {
	if x != nil && x.Contacts != nil {
		return *x.Contacts
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetCEO() string {
	if x != nil && x.CEO != nil {
		return *x.CEO
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetChiefAccountant() string {
	if x != nil && x.ChiefAccountant != nil {
		return *x.ChiefAccountant
	}
	return ""
}

func (x *UpdateCounterpartyRequest) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *UpdateCounterpartyRequest) GetAccess() []string {
	if x != nil {
		return x.Access
	}
	return nil
}

var File_counterparty_counterparty_proto protoreflect.FileDescriptor

const file_counterparty_counterparty_proto_rawDesc = "" +
	"\n" +
	"\x1fcounterparty/counterparty.proto\x12\fcounterparty\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"J\n" +
	"\x16CounterpartiesResponse\x120\n" +
	"\x05items\x18\x01 \x03(\v2\x1a.counterparty.CounterpartyR\x05items\",\n" +
	"\x04Link\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x10\n" +
	"\x03url\x18\x02 \x01(\tR\x03url\"\xc3\x06\n" +
	"\fCounterparty\x12\x1f\n" +
	"\x02ID\x18\x01 \x01(\x03B\x0f\x92A\f\x9a\x02\x01\x03\xa2\x02\x05int64R\x02id\x12\x12\n" +
	"\x04Kind\x18\x03 \x01(\bR\x04kind\x12\x12\n" +
	"\x04Name\x18\x06 \x01(\tR\x04name\x12 \n" +
	"\bFullName\x18\t \x01(\tH\x00R\tfull_name\x88\x01\x01\x12\x19\n" +
	"\x05Alias\x18\f \x01(\tH\x01R\x05alias\x88\x01\x01\x12\x10\n" +
	"\x03INN\x18\x0f \x01(\tR\x03inn\x12\x15\n" +
	"\x03KPP\x18\x12 \x01(\tH\x02R\x03kpp\x88\x01\x01\x12\x17\n" +
	"\x04OGRN\x18\x15 \x01(\tH\x03R\x04ogrn\x88\x01\x01\x12\x1d\n" +
	"\aAddress\x18\x18 \x01(\tH\x04R\aaddress\x88\x01\x01\x12&\n" +
	"\vPostAddress\x18\x1b \x01(\tH\x05R\fpost_address\x88\x01\x01\x12\x1f\n" +
	"\bContacts\x18\x1e \x01(\tH\x06R\bcontacts\x88\x01\x01\x12\x19\n" +
	"\x05Email\x18! \x01(\tH\aR\x05email\x88\x01\x01\x12\x15\n" +
	"\x03CEO\x18$ \x01(\tH\bR\x03ceo\x88\x01\x01\x12.\n" +
	"\x0fChiefAccountant\x18' \x01(\tH\tR\x10chief_accountant\x88\x01\x01\x12(\n" +
	"\x05links\x18* \x03(\v2\x12.counterparty.LinkR\x05links\x12\x16\n" +
	"\x06access\x18- \x03(\tR\x06access\x12>\n" +
	"\tCreatedAt\x180 \x01(\v2\x1a.google.protobuf.TimestampH\n" +
	"R\n" +
	"created_at\x88\x01\x01\x12>\n" +
	"\tUpdatedAt\x183 \x01(\v2\x1a.google.protobuf.TimestampH\vR\n" +
	"updated_at\x88\x01\x01\x12>\n" +
	"\tDeletedAt\x186 \x01(\v2\x1a.google.protobuf.TimestampH\fR\n" +
	"deleted_at\x88\x01\x01B\v\n" +
	"\t_FullNameB\b\n" +
	"\x06_AliasB\x06\n" +
	"\x04_KPPB\a\n" +
	"\x05_OGRNB\n" +
	"\n" +
	"\b_AddressB\x0e\n" +
	"\f_PostAddressB\v\n" +
	"\t_ContactsB\b\n" +
	"\x06_EmailB\x06\n" +
	"\x04_CEOB\x12\n" +
	"\x10_ChiefAccountantB\f\n" +
	"\n" +
	"_CreatedAtB\f\n" +
	"\n" +
	"_UpdatedAtB\f\n" +
	"\n" +
	"_DeletedAt\"\xc2\x04\n" +
	"\x16AddCounterpartyRequest\x12\x12\n" +
	"\x04Kind\x18\x03 \x01(\bR\x04kind\x12\x12\n" +
	"\x04Name\x18\x06 \x01(\tR\x04name\x12 \n" +
	"\bFullName\x18\t \x01(\tH\x00R\tfull_name\x88\x01\x01\x12\x19\n" +
	"\x05Alias\x18\f \x01(\tH\x01R\x05alias\x88\x01\x01\x12\x10\n" +
	"\x03INN\x18\x0f \x01(\tR\x03inn\x12\x15\n" +
	"\x03KPP\x18\x12 \x01(\tH\x02R\x03kpp\x88\x01\x01\x12\x17\n" +
	"\x04OGRN\x18\x15 \x01(\tH\x03R\x04ogrn\x88\x01\x01\x12\x1d\n" +
	"\aAddress\x18\x18 \x01(\tH\x04R\aaddress\x88\x01\x01\x12&\n" +
	"\vPostAddress\x18\x1b \x01(\tH\x05R\fpost_address\x88\x01\x01\x12\x1f\n" +
	"\bContacts\x18\x1e \x01(\tH\x06R\bcontacts\x88\x01\x01\x12\x19\n" +
	"\x05Email\x18! \x01(\tH\aR\x05email\x88\x01\x01\x12\x15\n" +
	"\x03CEO\x18$ \x01(\tH\bR\x03ceo\x88\x01\x01\x12.\n" +
	"\x0fChiefAccountant\x18' \x01(\tH\tR\x10chief_accountant\x88\x01\x01\x12(\n" +
	"\x05links\x18* \x03(\v2\x12.counterparty.LinkR\x05links\x12\x16\n" +
	"\x06access\x18- \x03(\tR\x06accessB\v\n" +
	"\t_FullNameB\b\n" +
	"\x06_AliasB\x06\n" +
	"\x04_KPPB\a\n" +
	"\x05_OGRNB\n" +
	"\n" +
	"\b_AddressB\x0e\n" +
	"\f_PostAddressB\v\n" +
	"\t_ContactsB\b\n" +
	"\x06_EmailB\x06\n" +
	"\x04_CEOB\x12\n" +
	"\x10_ChiefAccountant\"\xfe\x04\n" +
	"\x19UpdateCounterpartyRequest\x12\"\n" +
	"\x02id\x18\x01 \x01(\x03B\x12\x92A\f\x9a\x02\x01\x03\xa2\x02\x05int64\xe0A\x02R\x02id\x12\x12\n" +
	"\x04Kind\x18\x03 \x01(\bR\x04kind\x12\x12\n" +
	"\x04Name\x18\x06 \x01(\tR\x04name\x12 \n" +
	"\bFullName\x18\t \x01(\tH\x00R\tfull_name\x88\x01\x01\x12\x19\n" +
	"\x05Alias\x18\f \x01(\tH\x01R\x05alias\x88\x01\x01\x12%\n" +
	"\x03INN\x18\x0f \x01(\tB\x13\x92A\x10x\f\x80\x01\n" +
	"\x8a\x01\b^[0-9]+$R\x03inn\x12\x15\n" +
	"\x03KPP\x18\x12 \x01(\tH\x02R\x03kpp\x88\x01\x01\x12\x17\n" +
	"\x04OGRN\x18\x15 \x01(\tH\x03R\x04ogrn\x88\x01\x01\x12\x1d\n" +
	"\aAddress\x18\x18 \x01(\tH\x04R\aaddress\x88\x01\x01\x12&\n" +
	"\vPostAddress\x18\x1b \x01(\tH\x05R\fpost_address\x88\x01\x01\x12\x1f\n" +
	"\bContacts\x18\x1e \x01(\tH\x06R\bcontacts\x88\x01\x01\x12\x19\n" +
	"\x05Email\x18! \x01(\tH\aR\x05email\x88\x01\x01\x12\x15\n" +
	"\x03CEO\x18$ \x01(\tH\bR\x03ceo\x88\x01\x01\x12.\n" +
	"\x0fChiefAccountant\x18' \x01(\tH\tR\x10chief_accountant\x88\x01\x01\x12(\n" +
	"\x05links\x18* \x03(\v2\x12.counterparty.LinkR\x05links\x12\x16\n" +
	"\x06access\x18- \x03(\tR\x06accessB\v\n" +
	"\t_FullNameB\b\n" +
	"\x06_AliasB\x06\n" +
	"\x04_KPPB\a\n" +
	"\x05_OGRNB\n" +
	"\n" +
	"\b_AddressB\x0e\n" +
	"\f_PostAddressB\v\n" +
	"\t_ContactsB\b\n" +
	"\x06_EmailB\x06\n" +
	"\x04_CEOB\x12\n" +
	"\x10_ChiefAccountantB1Z/github.com/COTBU/sotbi.lib/pkg/api/counterpartyb\x06proto3"

var (
	file_counterparty_counterparty_proto_rawDescOnce sync.Once
	file_counterparty_counterparty_proto_rawDescData []byte
)

func file_counterparty_counterparty_proto_rawDescGZIP() []byte {
	file_counterparty_counterparty_proto_rawDescOnce.Do(func() {
		file_counterparty_counterparty_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_counterparty_counterparty_proto_rawDesc), len(file_counterparty_counterparty_proto_rawDesc)))
	})
	return file_counterparty_counterparty_proto_rawDescData
}

var file_counterparty_counterparty_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_counterparty_counterparty_proto_goTypes = []any{
	(*CounterpartiesResponse)(nil),    // 0: counterparty.CounterpartiesResponse
	(*Link)(nil),                      // 1: counterparty.Link
	(*Counterparty)(nil),              // 2: counterparty.Counterparty
	(*AddCounterpartyRequest)(nil),    // 3: counterparty.AddCounterpartyRequest
	(*UpdateCounterpartyRequest)(nil), // 4: counterparty.UpdateCounterpartyRequest
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
}
var file_counterparty_counterparty_proto_depIdxs = []int32{
	2, // 0: counterparty.CounterpartiesResponse.items:type_name -> counterparty.Counterparty
	1, // 1: counterparty.Counterparty.links:type_name -> counterparty.Link
	5, // 2: counterparty.Counterparty.CreatedAt:type_name -> google.protobuf.Timestamp
	5, // 3: counterparty.Counterparty.UpdatedAt:type_name -> google.protobuf.Timestamp
	5, // 4: counterparty.Counterparty.DeletedAt:type_name -> google.protobuf.Timestamp
	1, // 5: counterparty.AddCounterpartyRequest.links:type_name -> counterparty.Link
	1, // 6: counterparty.UpdateCounterpartyRequest.links:type_name -> counterparty.Link
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_counterparty_counterparty_proto_init() }
func file_counterparty_counterparty_proto_init() {
	if File_counterparty_counterparty_proto != nil {
		return
	}
	file_counterparty_counterparty_proto_msgTypes[2].OneofWrappers = []any{}
	file_counterparty_counterparty_proto_msgTypes[3].OneofWrappers = []any{}
	file_counterparty_counterparty_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_counterparty_counterparty_proto_rawDesc), len(file_counterparty_counterparty_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_counterparty_counterparty_proto_goTypes,
		DependencyIndexes: file_counterparty_counterparty_proto_depIdxs,
		MessageInfos:      file_counterparty_counterparty_proto_msgTypes,
	}.Build()
	File_counterparty_counterparty_proto = out.File
	file_counterparty_counterparty_proto_goTypes = nil
	file_counterparty_counterparty_proto_depIdxs = nil
}
