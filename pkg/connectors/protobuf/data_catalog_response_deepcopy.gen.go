// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data_catalog_response.proto

package connectors

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using DataComponentMetadata within kubernetes types, where deepcopy-gen is used.
func (in *DataComponentMetadata) DeepCopyInto(out *DataComponentMetadata) {
	p := proto.Clone(in).(*DataComponentMetadata)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataComponentMetadata. Required by controller-gen.
func (in *DataComponentMetadata) DeepCopy() *DataComponentMetadata {
	if in == nil {
		return nil
	}
	out := new(DataComponentMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DataComponentMetadata. Required by controller-gen.
func (in *DataComponentMetadata) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DatasetMetadata within kubernetes types, where deepcopy-gen is used.
func (in *DatasetMetadata) DeepCopyInto(out *DatasetMetadata) {
	p := proto.Clone(in).(*DatasetMetadata)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetMetadata. Required by controller-gen.
func (in *DatasetMetadata) DeepCopy() *DatasetMetadata {
	if in == nil {
		return nil
	}
	out := new(DatasetMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DatasetMetadata. Required by controller-gen.
func (in *DatasetMetadata) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DataStore within kubernetes types, where deepcopy-gen is used.
func (in *DataStore) DeepCopyInto(out *DataStore) {
	p := proto.Clone(in).(*DataStore)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStore. Required by controller-gen.
func (in *DataStore) DeepCopy() *DataStore {
	if in == nil {
		return nil
	}
	out := new(DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DataStore. Required by controller-gen.
func (in *DataStore) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using DatasetDetails within kubernetes types, where deepcopy-gen is used.
func (in *DatasetDetails) DeepCopyInto(out *DatasetDetails) {
	p := proto.Clone(in).(*DatasetDetails)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetDetails. Required by controller-gen.
func (in *DatasetDetails) DeepCopy() *DatasetDetails {
	if in == nil {
		return nil
	}
	out := new(DatasetDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new DatasetDetails. Required by controller-gen.
func (in *DatasetDetails) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using CatalogDatasetInfo within kubernetes types, where deepcopy-gen is used.
func (in *CatalogDatasetInfo) DeepCopyInto(out *CatalogDatasetInfo) {
	p := proto.Clone(in).(*CatalogDatasetInfo)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogDatasetInfo. Required by controller-gen.
func (in *CatalogDatasetInfo) DeepCopy() *CatalogDatasetInfo {
	if in == nil {
		return nil
	}
	out := new(CatalogDatasetInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CatalogDatasetInfo. Required by controller-gen.
func (in *CatalogDatasetInfo) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Db2DataStore within kubernetes types, where deepcopy-gen is used.
func (in *Db2DataStore) DeepCopyInto(out *Db2DataStore) {
	p := proto.Clone(in).(*Db2DataStore)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Db2DataStore. Required by controller-gen.
func (in *Db2DataStore) DeepCopy() *Db2DataStore {
	if in == nil {
		return nil
	}
	out := new(Db2DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Db2DataStore. Required by controller-gen.
func (in *Db2DataStore) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using S3DataStore within kubernetes types, where deepcopy-gen is used.
func (in *S3DataStore) DeepCopyInto(out *S3DataStore) {
	p := proto.Clone(in).(*S3DataStore)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3DataStore. Required by controller-gen.
func (in *S3DataStore) DeepCopy() *S3DataStore {
	if in == nil {
		return nil
	}
	out := new(S3DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new S3DataStore. Required by controller-gen.
func (in *S3DataStore) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using KafkaDataStore within kubernetes types, where deepcopy-gen is used.
func (in *KafkaDataStore) DeepCopyInto(out *KafkaDataStore) {
	p := proto.Clone(in).(*KafkaDataStore)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaDataStore. Required by controller-gen.
func (in *KafkaDataStore) DeepCopy() *KafkaDataStore {
	if in == nil {
		return nil
	}
	out := new(KafkaDataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new KafkaDataStore. Required by controller-gen.
func (in *KafkaDataStore) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}