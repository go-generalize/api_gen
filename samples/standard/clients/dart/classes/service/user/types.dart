// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../common.dart' as external_1ad882c;

class GetRequestConverter
    implements
        external_1ad882c.JsonConverter<GetRequest, Map<String, dynamic>> {
  const GetRequestConverter();

  @override
  GetRequest fromJson(dynamic s) {
    return GetRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetRequest s) {
    return s.toJson();
  }
}

class GetRequest {
  GetRequest();

  factory GetRequest.fromJson(Map<String, dynamic> json) {
    return GetRequest();
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{};
  }
}

class GetResponseConverter
    implements
        external_1ad882c.JsonConverter<GetResponse, Map<String, dynamic>> {
  const GetResponseConverter();

  @override
  GetResponse fromJson(dynamic s) {
    return GetResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetResponse s) {
    return s.toJson();
  }
}

class GetResponse {
  GetResponse();

  factory GetResponse.fromJson(Map<String, dynamic> json) {
    return GetResponse();
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{};
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
