// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: piratas.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ResultadoCompra_Estado int32

const (
	ResultadoCompra_EXITO              ResultadoCompra_Estado = 0
	ResultadoCompra_FRAUDE             ResultadoCompra_Estado = 1
	ResultadoCompra_ATAQUE_MERCENARIOS ResultadoCompra_Estado = 2
	ResultadoCompra_CONFISCADO         ResultadoCompra_Estado = 3
)

// Enum value maps for ResultadoCompra_Estado.
var (
	ResultadoCompra_Estado_name = map[int32]string{
		0: "EXITO",
		1: "FRAUDE",
		2: "ATAQUE_MERCENARIOS",
		3: "CONFISCADO",
	}
	ResultadoCompra_Estado_value = map[string]int32{
		"EXITO":              0,
		"FRAUDE":             1,
		"ATAQUE_MERCENARIOS": 2,
		"CONFISCADO":         3,
	}
)

func (x ResultadoCompra_Estado) Enum() *ResultadoCompra_Estado {
	p := new(ResultadoCompra_Estado)
	*p = x
	return p
}

func (x ResultadoCompra_Estado) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultadoCompra_Estado) Descriptor() protoreflect.EnumDescriptor {
	return file_piratas_proto_enumTypes[0].Descriptor()
}

func (ResultadoCompra_Estado) Type() protoreflect.EnumType {
	return &file_piratas_proto_enumTypes[0]
}

func (x ResultadoCompra_Estado) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultadoCompra_Estado.Descriptor instead.
func (ResultadoCompra_Estado) EnumDescriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{1, 0}
}

// Información del pirata ofrecido por el cazarrecompensas
type PirataCapturado struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre             string                 `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Recompensa         int32                  `protobuf:"varint,3,opt,name=recompensa,proto3" json:"recompensa,omitempty"`
	CazarrecompensasId string                 `protobuf:"bytes,4,opt,name=cazarrecompensas_id,json=cazarrecompensasId,proto3" json:"cazarrecompensas_id,omitempty"`
	Peligrosidad       string                 `protobuf:"bytes,5,opt,name=peligrosidad,proto3" json:"peligrosidad,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *PirataCapturado) Reset() {
	*x = PirataCapturado{}
	mi := &file_piratas_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PirataCapturado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PirataCapturado) ProtoMessage() {}

func (x *PirataCapturado) ProtoReflect() protoreflect.Message {
	mi := &file_piratas_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PirataCapturado.ProtoReflect.Descriptor instead.
func (*PirataCapturado) Descriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{0}
}

func (x *PirataCapturado) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PirataCapturado) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *PirataCapturado) GetRecompensa() int32 {
	if x != nil {
		return x.Recompensa
	}
	return 0
}

func (x *PirataCapturado) GetCazarrecompensasId() string {
	if x != nil {
		return x.CazarrecompensasId
	}
	return ""
}

func (x *PirataCapturado) GetPeligrosidad() string {
	if x != nil {
		return x.Peligrosidad
	}
	return ""
}

// Resultado de la compra
type ResultadoCompra struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Estado        ResultadoCompra_Estado `protobuf:"varint,1,opt,name=estado,proto3,enum=onepiece.ResultadoCompra_Estado" json:"estado,omitempty"`
	MontoPagado   int32                  `protobuf:"varint,2,opt,name=monto_pagado,json=montoPagado,proto3" json:"monto_pagado,omitempty"` // Solo relevante si estado = EXITO
	Mensaje       string                 `protobuf:"bytes,3,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResultadoCompra) Reset() {
	*x = ResultadoCompra{}
	mi := &file_piratas_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResultadoCompra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultadoCompra) ProtoMessage() {}

func (x *ResultadoCompra) ProtoReflect() protoreflect.Message {
	mi := &file_piratas_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultadoCompra.ProtoReflect.Descriptor instead.
func (*ResultadoCompra) Descriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{1}
}

func (x *ResultadoCompra) GetEstado() ResultadoCompra_Estado {
	if x != nil {
		return x.Estado
	}
	return ResultadoCompra_EXITO
}

func (x *ResultadoCompra) GetMontoPagado() int32 {
	if x != nil {
		return x.MontoPagado
	}
	return 0
}

func (x *ResultadoCompra) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

// Mensajes para el Gobierno
type SolicitudLista struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CazarrecompensaId string                 `protobuf:"bytes,1,opt,name=cazarrecompensa_id,json=cazarrecompensaId,proto3" json:"cazarrecompensa_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *SolicitudLista) Reset() {
	*x = SolicitudLista{}
	mi := &file_piratas_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolicitudLista) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudLista) ProtoMessage() {}

func (x *SolicitudLista) ProtoReflect() protoreflect.Message {
	mi := &file_piratas_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudLista.ProtoReflect.Descriptor instead.
func (*SolicitudLista) Descriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{2}
}

func (x *SolicitudLista) GetCazarrecompensaId() string {
	if x != nil {
		return x.CazarrecompensaId
	}
	return ""
}

type PirataInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre        string                 `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Recompensa    uint64                 `protobuf:"varint,3,opt,name=recompensa,proto3" json:"recompensa,omitempty"`
	Peligrosidad  string                 `protobuf:"bytes,4,opt,name=peligrosidad,proto3" json:"peligrosidad,omitempty"`
	Estado        string                 `protobuf:"bytes,5,opt,name=estado,proto3" json:"estado,omitempty"` // "Buscado" o "Capturado"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PirataInfo) Reset() {
	*x = PirataInfo{}
	mi := &file_piratas_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PirataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PirataInfo) ProtoMessage() {}

func (x *PirataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_piratas_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PirataInfo.ProtoReflect.Descriptor instead.
func (*PirataInfo) Descriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{3}
}

func (x *PirataInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PirataInfo) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *PirataInfo) GetRecompensa() uint64 {
	if x != nil {
		return x.Recompensa
	}
	return 0
}

func (x *PirataInfo) GetPeligrosidad() string {
	if x != nil {
		return x.Peligrosidad
	}
	return ""
}

func (x *PirataInfo) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type ListaPiratas struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Piratas       []*PirataInfo          `protobuf:"bytes,1,rep,name=piratas,proto3" json:"piratas,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListaPiratas) Reset() {
	*x = ListaPiratas{}
	mi := &file_piratas_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListaPiratas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListaPiratas) ProtoMessage() {}

func (x *ListaPiratas) ProtoReflect() protoreflect.Message {
	mi := &file_piratas_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListaPiratas.ProtoReflect.Descriptor instead.
func (*ListaPiratas) Descriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{4}
}

func (x *ListaPiratas) GetPiratas() []*PirataInfo {
	if x != nil {
		return x.Piratas
	}
	return nil
}

type ReporteCaptura struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CazarrecompensaId string                 `protobuf:"bytes,1,opt,name=cazarrecompensa_id,json=cazarrecompensaId,proto3" json:"cazarrecompensa_id,omitempty"`
	PirataId          string                 `protobuf:"bytes,2,opt,name=pirata_id,json=pirataId,proto3" json:"pirata_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ReporteCaptura) Reset() {
	*x = ReporteCaptura{}
	mi := &file_piratas_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReporteCaptura) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReporteCaptura) ProtoMessage() {}

func (x *ReporteCaptura) ProtoReflect() protoreflect.Message {
	mi := &file_piratas_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReporteCaptura.ProtoReflect.Descriptor instead.
func (*ReporteCaptura) Descriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{5}
}

func (x *ReporteCaptura) GetCazarrecompensaId() string {
	if x != nil {
		return x.CazarrecompensaId
	}
	return ""
}

func (x *ReporteCaptura) GetPirataId() string {
	if x != nil {
		return x.PirataId
	}
	return ""
}

type ConfirmacionCaptura struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Exito         bool                   `protobuf:"varint,1,opt,name=exito,proto3" json:"exito,omitempty"`
	Mensaje       string                 `protobuf:"bytes,2,opt,name=mensaje,proto3" json:"mensaje,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfirmacionCaptura) Reset() {
	*x = ConfirmacionCaptura{}
	mi := &file_piratas_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmacionCaptura) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmacionCaptura) ProtoMessage() {}

func (x *ConfirmacionCaptura) ProtoReflect() protoreflect.Message {
	mi := &file_piratas_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmacionCaptura.ProtoReflect.Descriptor instead.
func (*ConfirmacionCaptura) Descriptor() ([]byte, []int) {
	return file_piratas_proto_rawDescGZIP(), []int{6}
}

func (x *ConfirmacionCaptura) GetExito() bool {
	if x != nil {
		return x.Exito
	}
	return false
}

func (x *ConfirmacionCaptura) GetMensaje() string {
	if x != nil {
		return x.Mensaje
	}
	return ""
}

var File_piratas_proto protoreflect.FileDescriptor

const file_piratas_proto_rawDesc = "" +
	"\n" +
	"\rpiratas.proto\x12\bonepiece\"\xae\x01\n" +
	"\x0fPirataCapturado\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06nombre\x18\x02 \x01(\tR\x06nombre\x12\x1e\n" +
	"\n" +
	"recompensa\x18\x03 \x01(\x05R\n" +
	"recompensa\x12/\n" +
	"\x13cazarrecompensas_id\x18\x04 \x01(\tR\x12cazarrecompensasId\x12\"\n" +
	"\fpeligrosidad\x18\x05 \x01(\tR\fpeligrosidad\"\xd1\x01\n" +
	"\x0fResultadoCompra\x128\n" +
	"\x06estado\x18\x01 \x01(\x0e2 .onepiece.ResultadoCompra.EstadoR\x06estado\x12!\n" +
	"\fmonto_pagado\x18\x02 \x01(\x05R\vmontoPagado\x12\x18\n" +
	"\amensaje\x18\x03 \x01(\tR\amensaje\"G\n" +
	"\x06Estado\x12\t\n" +
	"\x05EXITO\x10\x00\x12\n" +
	"\n" +
	"\x06FRAUDE\x10\x01\x12\x16\n" +
	"\x12ATAQUE_MERCENARIOS\x10\x02\x12\x0e\n" +
	"\n" +
	"CONFISCADO\x10\x03\"?\n" +
	"\x0eSolicitudLista\x12-\n" +
	"\x12cazarrecompensa_id\x18\x01 \x01(\tR\x11cazarrecompensaId\"\x90\x01\n" +
	"\n" +
	"PirataInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06nombre\x18\x02 \x01(\tR\x06nombre\x12\x1e\n" +
	"\n" +
	"recompensa\x18\x03 \x01(\x04R\n" +
	"recompensa\x12\"\n" +
	"\fpeligrosidad\x18\x04 \x01(\tR\fpeligrosidad\x12\x16\n" +
	"\x06estado\x18\x05 \x01(\tR\x06estado\">\n" +
	"\fListaPiratas\x12.\n" +
	"\apiratas\x18\x01 \x03(\v2\x14.onepiece.PirataInfoR\apiratas\"\\\n" +
	"\x0eReporteCaptura\x12-\n" +
	"\x12cazarrecompensa_id\x18\x01 \x01(\tR\x11cazarrecompensaId\x12\x1b\n" +
	"\tpirata_id\x18\x02 \x01(\tR\bpirataId\"E\n" +
	"\x13ConfirmacionCaptura\x12\x14\n" +
	"\x05exito\x18\x01 \x01(\bR\x05exito\x12\x18\n" +
	"\amensaje\x18\x02 \x01(\tR\amensaje2X\n" +
	"\x0fSubmundoService\x12E\n" +
	"\rComprarPirata\x12\x19.onepiece.PirataCapturado\x1a\x19.onepiece.ResultadoCompra2\xa6\x01\n" +
	"\x0fGobiernoService\x12G\n" +
	"\x13ObtenerListaPiratas\x12\x18.onepiece.SolicitudLista\x1a\x16.onepiece.ListaPiratas\x12J\n" +
	"\x0fReportarCaptura\x12\x18.onepiece.ReporteCaptura\x1a\x1d.onepiece.ConfirmacionCapturaB\x13Z\x11grpc-server/protob\x06proto3"

var (
	file_piratas_proto_rawDescOnce sync.Once
	file_piratas_proto_rawDescData []byte
)

func file_piratas_proto_rawDescGZIP() []byte {
	file_piratas_proto_rawDescOnce.Do(func() {
		file_piratas_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_piratas_proto_rawDesc), len(file_piratas_proto_rawDesc)))
	})
	return file_piratas_proto_rawDescData
}

var file_piratas_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_piratas_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_piratas_proto_goTypes = []any{
	(ResultadoCompra_Estado)(0), // 0: onepiece.ResultadoCompra.Estado
	(*PirataCapturado)(nil),     // 1: onepiece.PirataCapturado
	(*ResultadoCompra)(nil),     // 2: onepiece.ResultadoCompra
	(*SolicitudLista)(nil),      // 3: onepiece.SolicitudLista
	(*PirataInfo)(nil),          // 4: onepiece.PirataInfo
	(*ListaPiratas)(nil),        // 5: onepiece.ListaPiratas
	(*ReporteCaptura)(nil),      // 6: onepiece.ReporteCaptura
	(*ConfirmacionCaptura)(nil), // 7: onepiece.ConfirmacionCaptura
}
var file_piratas_proto_depIdxs = []int32{
	0, // 0: onepiece.ResultadoCompra.estado:type_name -> onepiece.ResultadoCompra.Estado
	4, // 1: onepiece.ListaPiratas.piratas:type_name -> onepiece.PirataInfo
	1, // 2: onepiece.SubmundoService.ComprarPirata:input_type -> onepiece.PirataCapturado
	3, // 3: onepiece.GobiernoService.ObtenerListaPiratas:input_type -> onepiece.SolicitudLista
	6, // 4: onepiece.GobiernoService.ReportarCaptura:input_type -> onepiece.ReporteCaptura
	2, // 5: onepiece.SubmundoService.ComprarPirata:output_type -> onepiece.ResultadoCompra
	5, // 6: onepiece.GobiernoService.ObtenerListaPiratas:output_type -> onepiece.ListaPiratas
	7, // 7: onepiece.GobiernoService.ReportarCaptura:output_type -> onepiece.ConfirmacionCaptura
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_piratas_proto_init() }
func file_piratas_proto_init() {
	if File_piratas_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_piratas_proto_rawDesc), len(file_piratas_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_piratas_proto_goTypes,
		DependencyIndexes: file_piratas_proto_depIdxs,
		EnumInfos:         file_piratas_proto_enumTypes,
		MessageInfos:      file_piratas_proto_msgTypes,
	}.Build()
	File_piratas_proto = out.File
	file_piratas_proto_goTypes = nil
	file_piratas_proto_depIdxs = nil
}
