// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: ride_hailing.proto

package ridehailing

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

// User ride request
type RideRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PickupLocation  string                 `protobuf:"bytes,2,opt,name=pickup_location,json=pickupLocation,proto3" json:"pickup_location,omitempty"`
	DropoffLocation string                 `protobuf:"bytes,3,opt,name=dropoff_location,json=dropoffLocation,proto3" json:"dropoff_location,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RideRequest) Reset() {
	*x = RideRequest{}
	mi := &file_ride_hailing_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideRequest) ProtoMessage() {}

func (x *RideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideRequest.ProtoReflect.Descriptor instead.
func (*RideRequest) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{0}
}

func (x *RideRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RideRequest) GetPickupLocation() string {
	if x != nil {
		return x.PickupLocation
	}
	return ""
}

func (x *RideRequest) GetDropoffLocation() string {
	if x != nil {
		return x.DropoffLocation
	}
	return ""
}

// Response when a ride is assigned
type RideResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RideId        string                 `protobuf:"bytes,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	DriverId      string                 `protobuf:"bytes,2,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	DriverName    string                 `protobuf:"bytes,3,opt,name=driver_name,json=driverName,proto3" json:"driver_name,omitempty"`
	CarModel      string                 `protobuf:"bytes,4,opt,name=car_model,json=carModel,proto3" json:"car_model,omitempty"`
	CarPlate      string                 `protobuf:"bytes,5,opt,name=car_plate,json=carPlate,proto3" json:"car_plate,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"` // "assigned", "pending"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RideResponse) Reset() {
	*x = RideResponse{}
	mi := &file_ride_hailing_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideResponse) ProtoMessage() {}

func (x *RideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideResponse.ProtoReflect.Descriptor instead.
func (*RideResponse) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{1}
}

func (x *RideResponse) GetRideId() string {
	if x != nil {
		return x.RideId
	}
	return ""
}

func (x *RideResponse) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *RideResponse) GetDriverName() string {
	if x != nil {
		return x.DriverName
	}
	return ""
}

func (x *RideResponse) GetCarModel() string {
	if x != nil {
		return x.CarModel
	}
	return ""
}

func (x *RideResponse) GetCarPlate() string {
	if x != nil {
		return x.CarPlate
	}
	return ""
}

func (x *RideResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Request for driver location streaming
type DriverLocationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RideId        string                 `protobuf:"bytes,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DriverLocationRequest) Reset() {
	*x = DriverLocationRequest{}
	mi := &file_ride_hailing_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DriverLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocationRequest) ProtoMessage() {}

func (x *DriverLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverLocationRequest.ProtoReflect.Descriptor instead.
func (*DriverLocationRequest) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{2}
}

func (x *DriverLocationRequest) GetRideId() string {
	if x != nil {
		return x.RideId
	}
	return ""
}

// Streaming driver's live location to the user
type DriverLocation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DriverId      string                 `protobuf:"bytes,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Latitude      float64                `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float64                `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Timestamp     string                 `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DriverLocation) Reset() {
	*x = DriverLocation{}
	mi := &file_ride_hailing_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DriverLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverLocation) ProtoMessage() {}

func (x *DriverLocation) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverLocation.ProtoReflect.Descriptor instead.
func (*DriverLocation) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{3}
}

func (x *DriverLocation) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *DriverLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *DriverLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *DriverLocation) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

// Driver sends location updates
type LocationUpdateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // "success" or "error"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LocationUpdateResponse) Reset() {
	*x = LocationUpdateResponse{}
	mi := &file_ride_hailing_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocationUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationUpdateResponse) ProtoMessage() {}

func (x *LocationUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationUpdateResponse.ProtoReflect.Descriptor instead.
func (*LocationUpdateResponse) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{4}
}

func (x *LocationUpdateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Completing the ride
type RideCompletionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RideId        string                 `protobuf:"bytes,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	FareAmount    float64                `protobuf:"fixed64,2,opt,name=fare_amount,json=fareAmount,proto3" json:"fare_amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RideCompletionRequest) Reset() {
	*x = RideCompletionRequest{}
	mi := &file_ride_hailing_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideCompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideCompletionRequest) ProtoMessage() {}

func (x *RideCompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideCompletionRequest.ProtoReflect.Descriptor instead.
func (*RideCompletionRequest) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{5}
}

func (x *RideCompletionRequest) GetRideId() string {
	if x != nil {
		return x.RideId
	}
	return ""
}

func (x *RideCompletionRequest) GetFareAmount() float64 {
	if x != nil {
		return x.FareAmount
	}
	return 0
}

type RideCompletionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                                    // "completed" or "failed"
	PaymentStatus string                 `protobuf:"bytes,2,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"` // "paid", "pending"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RideCompletionResponse) Reset() {
	*x = RideCompletionResponse{}
	mi := &file_ride_hailing_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideCompletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideCompletionResponse) ProtoMessage() {}

func (x *RideCompletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideCompletionResponse.ProtoReflect.Descriptor instead.
func (*RideCompletionResponse) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{6}
}

func (x *RideCompletionResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RideCompletionResponse) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_ride_hailing_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[7]
	if x != nil {
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
	return file_ride_hailing_proto_rawDescGZIP(), []int{7}
}

type RequestAppIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         string                 `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestAppIdResponse) Reset() {
	*x = RequestAppIdResponse{}
	mi := &file_ride_hailing_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestAppIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAppIdResponse) ProtoMessage() {}

func (x *RequestAppIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ride_hailing_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAppIdResponse.ProtoReflect.Descriptor instead.
func (*RequestAppIdResponse) Descriptor() ([]byte, []int) {
	return file_ride_hailing_proto_rawDescGZIP(), []int{8}
}

func (x *RequestAppIdResponse) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

var File_ride_hailing_proto protoreflect.FileDescriptor

var file_ride_hailing_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x22, 0x7a, 0x0a, 0x0b, 0x52, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x69, 0x63,
	0x6b, 0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x6f, 0x66, 0x66, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x72,
	0x6f, 0x70, 0x6f, 0x66, 0x66, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb7, 0x01,
	0x0a, 0x0c, 0x52, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x69, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x30, 0x0a, 0x15, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x69, 0x64, 0x65, 0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x30, 0x0a, 0x16, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x51, 0x0a, 0x15, 0x52, 0x69, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x69, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x72, 0x65, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x66, 0x61, 0x72, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x16, 0x52, 0x69, 0x64, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x32, 0xa8, 0x03, 0x0a, 0x0b, 0x52, 0x69, 0x64, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x69, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x72, 0x69,
	0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x21, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x70, 0x70, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x72, 0x69, 0x64,
	0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x5a, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x23, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x57, 0x0a, 0x0c, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x12, 0x22, 0x2e, 0x72, 0x69, 0x64, 0x65,
	0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x69, 0x64, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x72, 0x69, 0x64, 0x65, 0x68, 0x61, 0x69, 0x6c, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_ride_hailing_proto_rawDescOnce sync.Once
	file_ride_hailing_proto_rawDescData []byte
)

func file_ride_hailing_proto_rawDescGZIP() []byte {
	file_ride_hailing_proto_rawDescOnce.Do(func() {
		file_ride_hailing_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ride_hailing_proto_rawDesc), len(file_ride_hailing_proto_rawDesc)))
	})
	return file_ride_hailing_proto_rawDescData
}

var file_ride_hailing_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ride_hailing_proto_goTypes = []any{
	(*RideRequest)(nil),            // 0: ridehailing.RideRequest
	(*RideResponse)(nil),           // 1: ridehailing.RideResponse
	(*DriverLocationRequest)(nil),  // 2: ridehailing.DriverLocationRequest
	(*DriverLocation)(nil),         // 3: ridehailing.DriverLocation
	(*LocationUpdateResponse)(nil), // 4: ridehailing.LocationUpdateResponse
	(*RideCompletionRequest)(nil),  // 5: ridehailing.RideCompletionRequest
	(*RideCompletionResponse)(nil), // 6: ridehailing.RideCompletionResponse
	(*Empty)(nil),                  // 7: ridehailing.Empty
	(*RequestAppIdResponse)(nil),   // 8: ridehailing.RequestAppIdResponse
}
var file_ride_hailing_proto_depIdxs = []int32{
	0, // 0: ridehailing.RideService.RequestRide:input_type -> ridehailing.RideRequest
	7, // 1: ridehailing.RideService.RequestAppId:input_type -> ridehailing.Empty
	2, // 2: ridehailing.RideService.StreamDriverLocation:input_type -> ridehailing.DriverLocationRequest
	3, // 3: ridehailing.RideService.UpdateDriverLocation:input_type -> ridehailing.DriverLocation
	5, // 4: ridehailing.RideService.CompleteRide:input_type -> ridehailing.RideCompletionRequest
	1, // 5: ridehailing.RideService.RequestRide:output_type -> ridehailing.RideResponse
	8, // 6: ridehailing.RideService.RequestAppId:output_type -> ridehailing.RequestAppIdResponse
	3, // 7: ridehailing.RideService.StreamDriverLocation:output_type -> ridehailing.DriverLocation
	4, // 8: ridehailing.RideService.UpdateDriverLocation:output_type -> ridehailing.LocationUpdateResponse
	6, // 9: ridehailing.RideService.CompleteRide:output_type -> ridehailing.RideCompletionResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ride_hailing_proto_init() }
func file_ride_hailing_proto_init() {
	if File_ride_hailing_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ride_hailing_proto_rawDesc), len(file_ride_hailing_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ride_hailing_proto_goTypes,
		DependencyIndexes: file_ride_hailing_proto_depIdxs,
		MessageInfos:      file_ride_hailing_proto_msgTypes,
	}.Build()
	File_ride_hailing_proto = out.File
	file_ride_hailing_proto_goTypes = nil
	file_ride_hailing_proto_depIdxs = nil
}
