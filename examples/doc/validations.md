# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [my_proto.proto](#my_proto-proto)
    - [AdvancedExample](#demo-AdvancedExample)
    - [AdvancedExample.Sub](#demo-AdvancedExample-Sub)
    - [AdvancedExample.Sub.MetadataEntry](#demo-AdvancedExample-Sub-MetadataEntry)
    - [ChatMessage](#demo-ChatMessage)
    - [ColoredPoint](#demo-ColoredPoint)
    - [Point](#demo-Point)
    - [User](#demo-User)
  
    - [Color](#demo-Color)
  
- [validate.proto](#validate-proto)
    - [AnyRules](#nanopb_validate-AnyRules)
    - [AtLeastRule](#nanopb_validate-AtLeastRule)
    - [BoolRules](#nanopb_validate-BoolRules)
    - [BytesRules](#nanopb_validate-BytesRules)
    - [DoubleRules](#nanopb_validate-DoubleRules)
    - [EnumRules](#nanopb_validate-EnumRules)
    - [FieldGroup](#nanopb_validate-FieldGroup)
    - [FieldRules](#nanopb_validate-FieldRules)
    - [Fixed32Rules](#nanopb_validate-Fixed32Rules)
    - [Fixed64Rules](#nanopb_validate-Fixed64Rules)
    - [FloatRules](#nanopb_validate-FloatRules)
    - [Int32Rules](#nanopb_validate-Int32Rules)
    - [Int64Rules](#nanopb_validate-Int64Rules)
    - [MapRules](#nanopb_validate-MapRules)
    - [MessageRules](#nanopb_validate-MessageRules)
    - [RepeatedRules](#nanopb_validate-RepeatedRules)
    - [SFixed32Rules](#nanopb_validate-SFixed32Rules)
    - [SFixed64Rules](#nanopb_validate-SFixed64Rules)
    - [SInt32Rules](#nanopb_validate-SInt32Rules)
    - [SInt64Rules](#nanopb_validate-SInt64Rules)
    - [StringRules](#nanopb_validate-StringRules)
    - [UInt32Rules](#nanopb_validate-UInt32Rules)
    - [UInt64Rules](#nanopb_validate-UInt64Rules)
  
    - [File-level Extensions](#validate-proto-extensions)
    - [File-level Extensions](#validate-proto-extensions)
    - [File-level Extensions](#validate-proto-extensions)
  
- [Scalar Value Types](#scalar-value-types)



<a name="my_proto-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## my_proto.proto



<a name="demo-AdvancedExample"></a>

### AdvancedExample



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) | oneof target |  |
| age | [int32](#int32) | oneof target |  |
| raw | [bytes](#bytes) | oneof target |  |
| details | [AdvancedExample.Sub](#demo-AdvancedExample-Sub) |  | validated via nested fields |
| favorite | [Color](#demo-Color) |  |  |






#### Validated Fields

| Field | Validations |
| ----- | ----------- |
| email | string.max_len: 255<br>string.email: true<br> |
| age | int32.lte: 150<br>int32.gte: 0<br> |
| raw | bytes.min_len: 4<br>bytes.max_len: 1024<br> |
| favorite | enum.defined_only: true<br> |







<a name="demo-AdvancedExample-Sub"></a>

### AdvancedExample.Sub
Nested submessage with its own constraints


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| tags | [string](#string) | repeated |  |
| metadata | [AdvancedExample.Sub.MetadataEntry](#demo-AdvancedExample-Sub-MetadataEntry) | repeated |  |
| score | [float](#float) |  |  |






#### Validated Fields

| Field | Validations |
| ----- | ----------- |
| id | string.min_len: 3<br>string.max_len: 64<br> |
| tags | repeated.min_items: 1<br>repeated.max_items: 10<br>repeated.unique: true<br> |
| metadata | map.min_pairs: 1<br>map.max_pairs: 20<br> |
| score | float.lte: 100<br>float.gte: 0<br> |







<a name="demo-AdvancedExample-Sub-MetadataEntry"></a>

### AdvancedExample.Sub.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |







<a name="demo-ChatMessage"></a>

### ChatMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| from | [User](#demo-User) |  |  |
| to | [User](#demo-User) |  |  |
| text | [string](#string) |  |  |
| location | [ColoredPoint](#demo-ColoredPoint) |  | where it was sent from |
| theme | [Color](#demo-Color) |  |  |






#### Validated Fields

| Field | Validations |
| ----- | ----------- |
| text | string.min_len: 1<br>string.max_len: 256<br> |
| theme | enum.defined_only: true<br> |







<a name="demo-ColoredPoint"></a>

### ColoredPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| p | [Point](#demo-Point) |  |  |
| color | [Color](#demo-Color) |  |  |






#### Validated Fields

| Field | Validations |
| ----- | ----------- |
| color | enum.defined_only: true<br> |







<a name="demo-Point"></a>

### Point



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| x | [float](#float) |  |  |
| y | [float](#float) |  |  |






#### Validated Fields

| Field | Validations |
| ----- | ----------- |
| x | float.lte: 1000<br>float.gte: -1000<br> |
| y | float.lte: 1000<br>float.gte: -1000<br> |







<a name="demo-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |






#### Validated Fields

| Field | Validations |
| ----- | ----------- |
| id | string.min_len: 1<br>string.max_len: 36<br> |
| name | string.min_len: 1<br>string.max_len: 50<br> |






 <!-- end messages -->


<a name="demo-Color"></a>

### Color
Simple enum used by other messages

| Name | Number | Description |
| ---- | ------ | ----------- |
| COLOR_UNSPECIFIED | 0 |  |
| RED | 1 |  |
| GREEN | 2 |  |
| BLUE | 3 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="validate-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## validate.proto



<a name="nanopb_validate-AnyRules"></a>

### AnyRules
Rules for google.protobuf.Any fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| in | [string](#string) | repeated | type_url must be in this list |
| not_in | [string](#string) | repeated | type_url must not be in this list |







<a name="nanopb_validate-AtLeastRule"></a>

### AtLeastRule
AtLeastRule requires at least N fields from the group to be set


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| n | [uint32](#uint32) | optional | Minimum number of fields that must be set |
| fields | [string](#string) | repeated | Field names in the group |







<a name="nanopb_validate-BoolRules"></a>

### BoolRules
Rules for bool fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [bool](#bool) | optional | Value must equal this |







<a name="nanopb_validate-BytesRules"></a>

### BytesRules
Rules for bytes fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [bytes](#bytes) | optional | Value must equal this |
| min_len | [uint32](#uint32) | optional | Minimum length in bytes |
| max_len | [uint32](#uint32) | optional | Maximum length in bytes |
| pattern | [string](#string) | optional | Regex pattern (not implemented in C runtime) |
| prefix | [bytes](#bytes) | optional | Value must start with these bytes |
| suffix | [bytes](#bytes) | optional | Value must end with these bytes |
| contains | [bytes](#bytes) | optional | Value must contain these bytes |
| in | [bytes](#bytes) | repeated | Value must be in this list |
| not_in | [bytes](#bytes) | repeated | Value must not be in this list |







<a name="nanopb_validate-DoubleRules"></a>

### DoubleRules
Numeric rules for double fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [double](#double) | optional |  |
| lt | [double](#double) | optional |  |
| lte | [double](#double) | optional |  |
| gt | [double](#double) | optional |  |
| gte | [double](#double) | optional |  |
| in | [double](#double) | repeated |  |
| not_in | [double](#double) | repeated |  |







<a name="nanopb_validate-EnumRules"></a>

### EnumRules
Rules for enum fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [int32](#int32) | optional | Value must equal this |
| defined_only | [bool](#bool) | optional | Value must be a defined enum value |
| in | [int32](#int32) | repeated | Value must be in this list |
| not_in | [int32](#int32) | repeated | Value must not be in this list |







<a name="nanopb_validate-FieldGroup"></a>

### FieldGroup
FieldGroup represents a group of field names


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [string](#string) | repeated |  |







<a name="nanopb_validate-FieldRules"></a>

### FieldRules
FieldRules encapsulates all validation rules for a single field


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| float | [FloatRules](#nanopb_validate-FloatRules) | optional | Type-specific rules (only one should be set per field) |
| double | [DoubleRules](#nanopb_validate-DoubleRules) | optional |  |
| int32 | [Int32Rules](#nanopb_validate-Int32Rules) | optional |  |
| int64 | [Int64Rules](#nanopb_validate-Int64Rules) | optional |  |
| uint32 | [UInt32Rules](#nanopb_validate-UInt32Rules) | optional |  |
| uint64 | [UInt64Rules](#nanopb_validate-UInt64Rules) | optional |  |
| sint32 | [SInt32Rules](#nanopb_validate-SInt32Rules) | optional |  |
| sint64 | [SInt64Rules](#nanopb_validate-SInt64Rules) | optional |  |
| fixed32 | [Fixed32Rules](#nanopb_validate-Fixed32Rules) | optional |  |
| fixed64 | [Fixed64Rules](#nanopb_validate-Fixed64Rules) | optional |  |
| sfixed32 | [SFixed32Rules](#nanopb_validate-SFixed32Rules) | optional |  |
| sfixed64 | [SFixed64Rules](#nanopb_validate-SFixed64Rules) | optional |  |
| bool | [BoolRules](#nanopb_validate-BoolRules) | optional |  |
| string | [StringRules](#nanopb_validate-StringRules) | optional |  |
| bytes | [BytesRules](#nanopb_validate-BytesRules) | optional |  |
| enum | [EnumRules](#nanopb_validate-EnumRules) | optional |  |
| repeated | [RepeatedRules](#nanopb_validate-RepeatedRules) | optional |  |
| map | [MapRules](#nanopb_validate-MapRules) | optional |  |
| any | [AnyRules](#nanopb_validate-AnyRules) | optional | Rules for google.protobuf.Any fields |







<a name="nanopb_validate-Fixed32Rules"></a>

### Fixed32Rules
Numeric rules for fixed32 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [fixed32](#fixed32) | optional |  |
| lt | [fixed32](#fixed32) | optional |  |
| lte | [fixed32](#fixed32) | optional |  |
| gt | [fixed32](#fixed32) | optional |  |
| gte | [fixed32](#fixed32) | optional |  |
| in | [fixed32](#fixed32) | repeated |  |
| not_in | [fixed32](#fixed32) | repeated |  |







<a name="nanopb_validate-Fixed64Rules"></a>

### Fixed64Rules
Numeric rules for fixed64 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [fixed64](#fixed64) | optional |  |
| lt | [fixed64](#fixed64) | optional |  |
| lte | [fixed64](#fixed64) | optional |  |
| gt | [fixed64](#fixed64) | optional |  |
| gte | [fixed64](#fixed64) | optional |  |
| in | [fixed64](#fixed64) | repeated |  |
| not_in | [fixed64](#fixed64) | repeated |  |







<a name="nanopb_validate-FloatRules"></a>

### FloatRules
Numeric rules for float fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [float](#float) | optional | Value must equal this |
| lt | [float](#float) | optional | Value must be less than this |
| lte | [float](#float) | optional | Value must be less than or equal to this |
| gt | [float](#float) | optional | Value must be greater than this |
| gte | [float](#float) | optional | Value must be greater than or equal to this |
| in | [float](#float) | repeated | Value must be in this list |
| not_in | [float](#float) | repeated | Value must not be in this list |







<a name="nanopb_validate-Int32Rules"></a>

### Int32Rules
Numeric rules for int32 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [int32](#int32) | optional |  |
| lt | [int32](#int32) | optional |  |
| lte | [int32](#int32) | optional |  |
| gt | [int32](#int32) | optional |  |
| gte | [int32](#int32) | optional |  |
| in | [int32](#int32) | repeated |  |
| not_in | [int32](#int32) | repeated |  |







<a name="nanopb_validate-Int64Rules"></a>

### Int64Rules
Numeric rules for int64 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [int64](#int64) | optional |  |
| lt | [int64](#int64) | optional |  |
| lte | [int64](#int64) | optional |  |
| gt | [int64](#int64) | optional |  |
| gte | [int64](#int64) | optional |  |
| in | [int64](#int64) | repeated |  |
| not_in | [int64](#int64) | repeated |  |







<a name="nanopb_validate-MapRules"></a>

### MapRules
Rules for map fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_pairs | [uint32](#uint32) | optional | Minimum number of key-value pairs |
| max_pairs | [uint32](#uint32) | optional | Maximum number of key-value pairs |
| no_sparse | [bool](#bool) | optional | Disallow keys with empty values |







<a name="nanopb_validate-MessageRules"></a>

### MessageRules
MessageRules defines validation rules that apply at the message level


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| requires | [string](#string) | repeated | Cross-field constraints

If this message is set, these fields must also be set |
| mutex | [FieldGroup](#nanopb_validate-FieldGroup) | repeated | Mutually exclusive field groups |
| at_least | [AtLeastRule](#nanopb_validate-AtLeastRule) | repeated | At least N of the specified fields must be set |







<a name="nanopb_validate-RepeatedRules"></a>

### RepeatedRules
Rules for repeated fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_items | [uint32](#uint32) | optional | Minimum number of items |
| max_items | [uint32](#uint32) | optional | Maximum number of items |
| unique | [bool](#bool) | optional | All items must be unique |







<a name="nanopb_validate-SFixed32Rules"></a>

### SFixed32Rules
Numeric rules for sfixed32 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sfixed32](#sfixed32) | optional |  |
| lt | [sfixed32](#sfixed32) | optional |  |
| lte | [sfixed32](#sfixed32) | optional |  |
| gt | [sfixed32](#sfixed32) | optional |  |
| gte | [sfixed32](#sfixed32) | optional |  |
| in | [sfixed32](#sfixed32) | repeated |  |
| not_in | [sfixed32](#sfixed32) | repeated |  |







<a name="nanopb_validate-SFixed64Rules"></a>

### SFixed64Rules
Numeric rules for sfixed64 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sfixed64](#sfixed64) | optional |  |
| lt | [sfixed64](#sfixed64) | optional |  |
| lte | [sfixed64](#sfixed64) | optional |  |
| gt | [sfixed64](#sfixed64) | optional |  |
| gte | [sfixed64](#sfixed64) | optional |  |
| in | [sfixed64](#sfixed64) | repeated |  |
| not_in | [sfixed64](#sfixed64) | repeated |  |







<a name="nanopb_validate-SInt32Rules"></a>

### SInt32Rules
Numeric rules for sint32 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sint32](#sint32) | optional |  |
| lt | [sint32](#sint32) | optional |  |
| lte | [sint32](#sint32) | optional |  |
| gt | [sint32](#sint32) | optional |  |
| gte | [sint32](#sint32) | optional |  |
| in | [sint32](#sint32) | repeated |  |
| not_in | [sint32](#sint32) | repeated |  |







<a name="nanopb_validate-SInt64Rules"></a>

### SInt64Rules
Numeric rules for sint64 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [sint64](#sint64) | optional |  |
| lt | [sint64](#sint64) | optional |  |
| lte | [sint64](#sint64) | optional |  |
| gt | [sint64](#sint64) | optional |  |
| gte | [sint64](#sint64) | optional |  |
| in | [sint64](#sint64) | repeated |  |
| not_in | [sint64](#sint64) | repeated |  |







<a name="nanopb_validate-StringRules"></a>

### StringRules
Rules for string fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [string](#string) | optional | Value must equal this |
| min_len | [uint32](#uint32) | optional | Minimum length in characters |
| max_len | [uint32](#uint32) | optional | Maximum length in characters |
| min_bytes | [uint32](#uint32) | optional | Minimum length in bytes |
| max_bytes | [uint32](#uint32) | optional | Maximum length in bytes |
| pattern | [string](#string) | optional | Regex pattern (not implemented in C runtime) |
| prefix | [string](#string) | optional | Value must start with this |
| suffix | [string](#string) | optional | Value must end with this |
| contains | [string](#string) | optional | Value must contain this substring |
| in | [string](#string) | repeated | Value must be in this list |
| not_in | [string](#string) | repeated | Value must not be in this list |
| ascii | [bool](#bool) | optional | Value must be ASCII only |
| email | [bool](#bool) | optional | Value must be a valid email |
| hostname | [bool](#bool) | optional | Value must be a valid hostname |
| ip | [bool](#bool) | optional | Value must be a valid IP (v4 or v6) |
| ipv4 | [bool](#bool) | optional | Value must be a valid IPv4 |
| ipv6 | [bool](#bool) | optional | Value must be a valid IPv6 |







<a name="nanopb_validate-UInt32Rules"></a>

### UInt32Rules
Numeric rules for uint32 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [uint32](#uint32) | optional |  |
| lt | [uint32](#uint32) | optional |  |
| lte | [uint32](#uint32) | optional |  |
| gt | [uint32](#uint32) | optional |  |
| gte | [uint32](#uint32) | optional |  |
| in | [uint32](#uint32) | repeated |  |
| not_in | [uint32](#uint32) | repeated |  |







<a name="nanopb_validate-UInt64Rules"></a>

### UInt64Rules
Numeric rules for uint64 fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| const | [uint64](#uint64) | optional |  |
| lt | [uint64](#uint64) | optional |  |
| lte | [uint64](#uint64) | optional |  |
| gt | [uint64](#uint64) | optional |  |
| gte | [uint64](#uint64) | optional |  |
| in | [uint64](#uint64) | repeated |  |
| not_in | [uint64](#uint64) | repeated |  |






 <!-- end messages -->

 <!-- end enums -->


<a name="validate-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| rules | FieldRules | .google.protobuf.FieldOptions | 1011 | (nanopb_validate.rules) |
| validate | bool | .google.protobuf.FileOptions | 1011 | Enable validation for entire file |
| message | MessageRules | .google.protobuf.MessageOptions | 1011 |  |

 <!-- end HasExtensions -->

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

