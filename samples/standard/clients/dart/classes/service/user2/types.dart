// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../common.dart' as external_1ad882c;

class DeleteUserRequestConverter
    implements
        external_1ad882c
            .JsonConverter<DeleteUserRequest, Map<String, dynamic>> {
  const DeleteUserRequestConverter();

  @override
  DeleteUserRequest fromJson(dynamic s) {
    return DeleteUserRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(DeleteUserRequest s) {
    return s.toJson();
  }
}

class DeleteUserRequest {
  String id;

  DeleteUserRequest({
    this.id = '',
  });

  factory DeleteUserRequest.fromJson(Map<String, dynamic> json) {
    return DeleteUserRequest(
      id: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['id']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'id': const external_1ad882c.DoNothingConverter<String>().toJson(id),
    };
  }
}

class DeleteUserResponseConverter
    implements
        external_1ad882c
            .JsonConverter<DeleteUserResponse, Map<String, dynamic>> {
  const DeleteUserResponseConverter();

  @override
  DeleteUserResponse fromJson(dynamic s) {
    return DeleteUserResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(DeleteUserResponse s) {
    return s.toJson();
  }
}

class DeleteUserResponse {
  DeleteUserResponse();

  factory DeleteUserResponse.fromJson(Map<String, dynamic> json) {
    return DeleteUserResponse();
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{};
  }
}

class EmbeddingConverter
    implements external_1ad882c.JsonConverter<Embedding, Map<String, dynamic>> {
  const EmbeddingConverter();

  @override
  Embedding fromJson(dynamic s) {
    return Embedding.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(Embedding s) {
    return s.toJson();
  }
}

class Embedding {
  String page;

  Embedding({
    this.page = '',
  });

  factory Embedding.fromJson(Map<String, dynamic> json) {
    return Embedding(
      page: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['page']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'page': const external_1ad882c.DoNothingConverter<String>().toJson(page),
    };
  }
}

class GetUserRequestConverter
    implements
        external_1ad882c.JsonConverter<GetUserRequest, Map<String, dynamic>> {
  const GetUserRequestConverter();

  @override
  GetUserRequest fromJson(dynamic s) {
    return GetUserRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetUserRequest s) {
    return s.toJson();
  }
}

class GetUserRequest {
  String id;
  String page;
  String searchRequest;

  GetUserRequest({
    this.id = '',
    this.page = '',
    this.searchRequest = '',
  });

  factory GetUserRequest.fromJson(Map<String, dynamic> json) {
    return GetUserRequest(
      id: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['id']),
      page: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['page']),
      searchRequest: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['search_request']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'id': const external_1ad882c.DoNothingConverter<String>().toJson(id),
      'page': const external_1ad882c.DoNothingConverter<String>().toJson(page),
      'search_request': const external_1ad882c.DoNothingConverter<String>()
          .toJson(searchRequest),
    };
  }
}

class GetUserResponseConverter
    implements
        external_1ad882c.JsonConverter<GetUserResponse, Map<String, dynamic>> {
  const GetUserResponseConverter();

  @override
  GetUserResponse fromJson(dynamic s) {
    return GetUserResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetUserResponse s) {
    return s.toJson();
  }
}

class GetUserResponse {
  String id;
  DateTime requestTime;
  String searchRequest;

  GetUserResponse({
    this.id = '',
    required this.requestTime,
    this.searchRequest = '',
  });

  factory GetUserResponse.fromJson(Map<String, dynamic> json) {
    return GetUserResponse(
      id: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['ID']),
      requestTime: const external_1ad882c.DateTimeConverter()
          .fromJson(json['RequestTime']),
      searchRequest: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['SearchRequest']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': const external_1ad882c.DoNothingConverter<String>().toJson(id),
      'RequestTime':
          const external_1ad882c.DateTimeConverter().toJson(requestTime),
      'SearchRequest': const external_1ad882c.DoNothingConverter<String>()
          .toJson(searchRequest),
    };
  }
}

class PostUpdateUserNameRequestConverter
    implements
        external_1ad882c
            .JsonConverter<PostUpdateUserNameRequest, Map<String, dynamic>> {
  const PostUpdateUserNameRequestConverter();

  @override
  PostUpdateUserNameRequest fromJson(dynamic s) {
    return PostUpdateUserNameRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostUpdateUserNameRequest s) {
    return s.toJson();
  }
}

class PostUpdateUserNameRequest {
  String name;

  PostUpdateUserNameRequest({
    this.name = '',
  });

  factory PostUpdateUserNameRequest.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserNameRequest(
      name: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['Name']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Name': const external_1ad882c.DoNothingConverter<String>().toJson(name),
    };
  }
}

class PostUpdateUserNameResponseConverter
    implements
        external_1ad882c
            .JsonConverter<PostUpdateUserNameResponse, Map<String, dynamic>> {
  const PostUpdateUserNameResponseConverter();

  @override
  PostUpdateUserNameResponse fromJson(dynamic s) {
    return PostUpdateUserNameResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostUpdateUserNameResponse s) {
    return s.toJson();
  }
}

class PostUpdateUserNameResponse {
  String message;
  DateTime requestTime;
  bool status;

  PostUpdateUserNameResponse({
    this.message = '',
    required this.requestTime,
    this.status = false,
  });

  factory PostUpdateUserNameResponse.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserNameResponse(
      message: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['Message']),
      requestTime: const external_1ad882c.DateTimeConverter()
          .fromJson(json['RequestTime']),
      status: const external_1ad882c.DoNothingConverter<bool>()
          .fromJson(json['Status']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Message':
          const external_1ad882c.DoNothingConverter<String>().toJson(message),
      'RequestTime':
          const external_1ad882c.DateTimeConverter().toJson(requestTime),
      'Status':
          const external_1ad882c.DoNothingConverter<bool>().toJson(status),
    };
  }
}

class PostUpdateUserPasswordRequestConverter
    implements
        external_1ad882c.JsonConverter<PostUpdateUserPasswordRequest,
            Map<String, dynamic>> {
  const PostUpdateUserPasswordRequestConverter();

  @override
  PostUpdateUserPasswordRequest fromJson(dynamic s) {
    return PostUpdateUserPasswordRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostUpdateUserPasswordRequest s) {
    return s.toJson();
  }
}

class PostUpdateUserPasswordRequest {
  String password;
  String passwordConfirm;

  PostUpdateUserPasswordRequest({
    this.password = '',
    this.passwordConfirm = '',
  });

  factory PostUpdateUserPasswordRequest.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserPasswordRequest(
      password: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['Password']),
      passwordConfirm: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['PasswordConfirm']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Password':
          const external_1ad882c.DoNothingConverter<String>().toJson(password),
      'PasswordConfirm': const external_1ad882c.DoNothingConverter<String>()
          .toJson(passwordConfirm),
    };
  }
}

class PostUpdateUserPasswordResponseConverter
    implements
        external_1ad882c.JsonConverter<PostUpdateUserPasswordResponse,
            Map<String, dynamic>> {
  const PostUpdateUserPasswordResponseConverter();

  @override
  PostUpdateUserPasswordResponse fromJson(dynamic s) {
    return PostUpdateUserPasswordResponse.fromJson(
        Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostUpdateUserPasswordResponse s) {
    return s.toJson();
  }
}

class PostUpdateUserPasswordResponse {
  String message;
  DateTime requestTime;
  bool status;

  PostUpdateUserPasswordResponse({
    this.message = '',
    required this.requestTime,
    this.status = false,
  });

  factory PostUpdateUserPasswordResponse.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserPasswordResponse(
      message: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['Message']),
      requestTime: const external_1ad882c.DateTimeConverter()
          .fromJson(json['RequestTime']),
      status: const external_1ad882c.DoNothingConverter<bool>()
          .fromJson(json['Status']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Message':
          const external_1ad882c.DoNothingConverter<String>().toJson(message),
      'RequestTime':
          const external_1ad882c.DateTimeConverter().toJson(requestTime),
      'Status':
          const external_1ad882c.DoNothingConverter<bool>().toJson(status),
    };
  }
}
