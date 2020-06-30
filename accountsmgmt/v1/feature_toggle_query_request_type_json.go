/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalFeatureToggleQueryRequest writes a value of the 'feature_toggle_query_request' type to the given writer.
func MarshalFeatureToggleQueryRequest(object *FeatureToggleQueryRequest, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeFeatureToggleQueryRequest(object, stream)
	stream.Flush()
	return stream.Error
}

// writeFeatureToggleQueryRequest writes a value of the 'feature_toggle_query_request' type to the given stream.
func writeFeatureToggleQueryRequest(object *FeatureToggleQueryRequest, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	if count > 0 {
		stream.WriteMore()
	}
	stream.WriteObjectField("kind")
	if object.link {
		stream.WriteString(FeatureToggleQueryRequestLinkKind)
	} else {
		stream.WriteString(FeatureToggleQueryRequestKind)
	}
	count++
	if object.id != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(*object.id)
		count++
	}
	if object.href != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(*object.href)
		count++
	}
	if object.organizationID != nil {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("organization_id")
		stream.WriteString(*object.organizationID)
		count++
	}
	stream.WriteObjectEnd()
}

// UnmarshalFeatureToggleQueryRequest reads a value of the 'feature_toggle_query_request' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalFeatureToggleQueryRequest(source interface{}) (object *FeatureToggleQueryRequest, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readFeatureToggleQueryRequest(iterator)
	err = iterator.Error
	return
}

// readFeatureToggleQueryRequest reads a value of the 'feature_toggle_query_request' type from the given iterator.
func readFeatureToggleQueryRequest(iterator *jsoniter.Iterator) *FeatureToggleQueryRequest {
	object := &FeatureToggleQueryRequest{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			object.link = value == FeatureToggleQueryRequestLinkKind
		case "id":
			value := iterator.ReadString()
			object.id = &value
		case "href":
			value := iterator.ReadString()
			object.href = &value
		case "organization_id":
			value := iterator.ReadString()
			object.organizationID = &value
		default:
			iterator.ReadAny()
		}
	}
	return object
}
