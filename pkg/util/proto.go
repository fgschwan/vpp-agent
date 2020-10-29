//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package util

import (
	protoV2 "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"

	"github.com/golang/protobuf/proto"
)

func ExtractProtos(from ...interface{}) (protos []proto.Message) {
	for _, v := range from {
		if reflect.ValueOf(v).IsNil() {
			continue
		}
		val := reflect.ValueOf(v).Elem()
		typ := val.Type()
		if typ.Kind() != reflect.Struct {
			return
		}
		for i := 0; i < typ.NumField(); i++ {
			field := val.Field(i)
			if field.Kind() == reflect.Slice {
				for idx := 0; idx < field.Len(); idx++ {
					elem := field.Index(idx)
					if msg, ok := elem.Interface().(proto.Message); ok {
						protos = append(protos, msg)
					}
				}
			} else if field.Kind() == reflect.Ptr && !field.IsNil() {
				if msg, ok := field.Interface().(proto.Message); ok && !field.IsNil() {
					protos = append(protos, msg)
				}
			}
		}
	}
	return
}

func PlaceProtos(protos map[string]proto.Message, dsts ...interface{}) {
	for _, prot := range protos {
		protTyp := reflect.TypeOf(prot)
		for _, dst := range dsts {
			dstVal := reflect.ValueOf(dst).Elem()
			dstTyp := dstVal.Type()
			if dstTyp.Kind() != reflect.Struct {
				return
			}
			for i := 0; i < dstTyp.NumField(); i++ {
				field := dstVal.Field(i)
				if field.Kind() == reflect.Slice {
					if protTyp.AssignableTo(field.Type().Elem()) {
						field.Set(reflect.Append(field, reflect.ValueOf(prot)))
					}
				} else {
					if field.Type() == protTyp {
						field.Set(reflect.ValueOf(prot))
					}
				}
			}
		}
	}
	return
}

// PlaceProtosIntoProtos fills dsts proto messages (direct or transitive) fields with protos values.
// The matching is done by message descriptor's full name.
func PlaceProtosIntoProtos(protos []protoV2.Message, dsts ...protoV2.Message) {
	protosMap := make(map[string][]protoV2.Message)
	for _, protoMsg := range protos {
		protoName := string(protoMsg.ProtoReflect().Descriptor().FullName())
		protosMap[protoName] = append(protosMap[protoName], protoMsg)
	}
	for _, dst := range dsts {
		placeProtosInProto(dst, protosMap)
	}
}

// placeProtosInProto fills dst proto message (direct or transitive) fields with protos values from protoMap
// (convenient map[proto descriptor full name]= proto value). The matching is done by message descriptor's
// full name. The function is recursive and one run is handling only one level of proto message structure tree
// (only handling Message references and ignoring scalar/enum/... values)
// Currently unsupported are maps as fields.
func placeProtosInProto(dst protoV2.Message, protosMap map[string][]protoV2.Message) {
	fields := dst.ProtoReflect().Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		fieldMessageDesc := field.Message()
		if fieldMessageDesc != nil { // only interested in MessageKind or GroupKind fields
			if protoMsgsForField, typeMatch := protosMap[string(fieldMessageDesc.FullName())]; typeMatch {
				// fill value(s)
				if field.IsList() {
					list := dst.ProtoReflect().Mutable(field).List()
					for _, protoMsg := range protoMsgsForField {
						list.Append(protoreflect.ValueOf(protoMsg))
					}
				} else if field.IsMap() { // unsupported
				} else {
					dst.ProtoReflect().Set(field, protoreflect.ValueOf(protoMsgsForField[0]))
				}
			} else {
				// no type match -> check deeper structure layers
				if field.IsList() {
					list := dst.ProtoReflect().Mutable(field).List()
					for j:=0; j < list.Len(); j++ {
						placeProtosInProto(list.Get(j).Message().Interface(), protosMap)
					}
				} else if field.IsMap() { // unsupported
				} else {
					placeProtosInProto(dst.ProtoReflect().Mutable(field).Message().Interface(), protosMap)
				}
			}
		}
	}
}