// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../common.dart' as external_1ad882c;

class PostUserRequestConverter
    implements
        external_1ad882c.JsonConverter<PostUserRequest, Map<String, dynamic>> {
  const PostUserRequestConverter();

  @override
  PostUserRequest fromJson(dynamic s) {
    return PostUserRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostUserRequest s) {
    return s.toJson();
  }
}

class PostUserRequest {
  PostUserRequest();

  factory PostUserRequest.fromJson(Map<String, dynamic> json) {
    return PostUserRequest();
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{};
  }
}

class PostUserResponseConverter
    implements
        external_1ad882c.JsonConverter<PostUserResponse, Map<String, dynamic>> {
  const PostUserResponseConverter();

  @override
  PostUserResponse fromJson(dynamic s) {
    return PostUserResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostUserResponse s) {
    return s.toJson();
  }
}

class PostUserResponse {
  PostUserResponse();

  factory PostUserResponse.fromJson(Map<String, dynamic> json) {
    return PostUserResponse();
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{};
  }
}
