# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/auth/v1/auth.proto](#proto_auth_v1_auth-proto)
    - [UserRequest](#auth-v1-UserRequest)
    - [UserResponse](#auth-v1-UserResponse)
  
    - [AuthService](#auth-v1-AuthService)
  
- [proto/learning/v1/learning.proto](#proto_learning_v1_learning-proto)
    - [AnswerRequest](#learning-v1-AnswerRequest)
    - [AnswerResponse](#learning-v1-AnswerResponse)
  
    - [LearningService](#learning-v1-LearningService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_auth_v1_auth-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/auth/v1/auth.proto



<a name="auth-v1-UserRequest"></a>

### UserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_name | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="auth-v1-UserResponse"></a>

### UserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_name | [string](#string) |  |  |
| token | [string](#string) |  |  |





 

 

 


<a name="auth-v1-AuthService"></a>

### AuthService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUser | [UserRequest](#auth-v1-UserRequest) | [UserResponse](#auth-v1-UserResponse) |  |

 



<a name="proto_learning_v1_learning-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/learning/v1/learning.proto



<a name="learning-v1-AnswerRequest"></a>

### AnswerRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sentence | [string](#string) |  |  |
| user_id | [string](#string) |  |  |






<a name="learning-v1-AnswerResponse"></a>

### AnswerResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sentence | [string](#string) |  |  |





 

 

 


<a name="learning-v1-LearningService"></a>

### LearningService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Answer | [AnswerRequest](#learning-v1-AnswerRequest) | [AnswerResponse](#learning-v1-AnswerResponse) |  |
| CreateAnswer | [AnswerRequest](#learning-v1-AnswerRequest) | [AnswerResponse](#learning-v1-AnswerResponse) |  |

 



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

