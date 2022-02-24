// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../common.dart' as external_9298cca;
import 'package:http/http.dart' as external_d8d61af;

enum CreatedType {
  createdTypeGuest,
  createdTypeMember,
  createdTypeOwner,
}

class CreatedTypeConverter
    implements external_9298cca.JsonConverter<CreatedType, int> {
  const CreatedTypeConverter();

  @override
  CreatedType fromJson(dynamic s) {
    return CreatedTypeExtension.fromJson(s as int);
  }

  @override
  int toJson(CreatedType s) {
    return s.toJson();
  }
}

extension CreatedTypeExtension on CreatedType {
  static final enumToValue = {
    CreatedType.createdTypeGuest: 2,
    CreatedType.createdTypeMember: 1,
    CreatedType.createdTypeOwner: 0,
  };
  static final valueToEnum = {
    2: CreatedType.createdTypeGuest,
    1: CreatedType.createdTypeMember,
    0: CreatedType.createdTypeOwner,
  };

  static CreatedType fromJson(dynamic d) {
    return CreatedTypeExtension.valueToEnum[d]!;
  }

  int toJson() {
    return CreatedTypeExtension.enumToValue[this]!;
  }
}

enum Enum {
  enumA,
  enumB,
  enumC,
}

class EnumConverter implements external_9298cca.JsonConverter<Enum, String> {
  const EnumConverter();

  @override
  Enum fromJson(dynamic s) {
    return EnumExtension.fromJson(s as String);
  }

  @override
  String toJson(Enum s) {
    return s.toJson();
  }
}

extension EnumExtension on Enum {
  static final enumToValue = {
    Enum.enumA: 'A',
    Enum.enumB: 'B',
    Enum.enumC: 'C',
  };
  static final valueToEnum = {
    'A': Enum.enumA,
    'B': Enum.enumB,
    'C': Enum.enumC,
  };

  static Enum fromJson(dynamic d) {
    return EnumExtension.valueToEnum[d]!;
  }

  String toJson() {
    return EnumExtension.enumToValue[this]!;
  }
}

class GetRequestConverter
    implements
        external_9298cca.JsonConverter<GetRequest, Map<String, dynamic>> {
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
  Enum enum_;
  String param;
  DateTime time;

  GetRequest({
    this.enum_ = Enum.enumA,
    this.param = '',
    required this.time,
  });

  factory GetRequest.fromJson(Map<String, dynamic> json) {
    return GetRequest(
      enum_: const EnumConverter().fromJson(json['Enum']),
      param: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Param']),
      time: const external_9298cca.DateTimeConverter().fromJson(json['Time']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Enum': const EnumConverter().toJson(enum_),
      'Param':
          const external_9298cca.DoNothingConverter<String>().toJson(param),
      'Time': const external_9298cca.DateTimeConverter().toJson(time),
    };
  }
}

class GetResponseConverter
    implements
        external_9298cca.JsonConverter<GetResponse, Map<String, dynamic>> {
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
  String data;

  GetResponse({
    this.data = '',
  });

  factory GetResponse.fromJson(Map<String, dynamic> json) {
    return GetResponse(
      data: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Data']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Data': const external_9298cca.DoNothingConverter<String>().toJson(data),
    };
  }
}

class PostCreateTableRequestConverter
    implements
        external_9298cca
            .JsonConverter<PostCreateTableRequest, Map<String, dynamic>> {
  const PostCreateTableRequestConverter();

  @override
  PostCreateTableRequest fromJson(dynamic s) {
    return PostCreateTableRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostCreateTableRequest s) {
    return s.toJson();
  }
}

class PostCreateTableRequest {
  external_d8d61af.MultipartFile? file;
  List<external_d8d61af.MultipartFile>? files;
  int flag;
  String id;
  Map<int, int>? map;
  String text;

  PostCreateTableRequest({
    this.file,
    this.files,
    this.flag = 0,
    this.id = '',
    this.map,
    this.text = '',
  });

  factory PostCreateTableRequest.fromJson(Map<String, dynamic> json) {
    return PostCreateTableRequest(
      flag: const external_9298cca.DoNothingConverter<int>()
          .fromJson(json['Flag']),
      id: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['ID']),
      map: const external_9298cca
                  .NullableConverter<Map<int, int>, Map<int, int>>(
              external_9298cca.MapConverter<int, int, int>(
                  external_9298cca.DoNothingConverter<int>()))
          .fromJson(json['map']),
      text: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Text']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Flag': const external_9298cca.DoNothingConverter<int>().toJson(flag),
      'ID': const external_9298cca.DoNothingConverter<String>().toJson(id),
      'map': const external_9298cca
                  .NullableConverter<Map<int, int>, Map<int, int>>(
              external_9298cca.MapConverter<int, int, int>(
                  external_9298cca.DoNothingConverter<int>()))
          .toJson(map),
      'Text': const external_9298cca.DoNothingConverter<String>().toJson(text),
    };
  }
}

class PostCreateTableResponseConverter
    implements
        external_9298cca
            .JsonConverter<PostCreateTableResponse, Map<String, dynamic>> {
  const PostCreateTableResponseConverter();

  @override
  PostCreateTableResponse fromJson(dynamic s) {
    return PostCreateTableResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostCreateTableResponse s) {
    return s.toJson();
  }
}

class PostCreateTableResponse {
  String id;
  DateTime requestTime;

  PostCreateTableResponse({
    this.id = '',
    required this.requestTime,
  });

  factory PostCreateTableResponse.fromJson(Map<String, dynamic> json) {
    return PostCreateTableResponse(
      id: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['ID']),
      requestTime: const external_9298cca.DateTimeConverter()
          .fromJson(json['RequestTime']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': const external_9298cca.DoNothingConverter<String>().toJson(id),
      'RequestTime':
          const external_9298cca.DateTimeConverter().toJson(requestTime),
    };
  }
}

class PostCreateUserRequestConverter
    implements
        external_9298cca
            .JsonConverter<PostCreateUserRequest, Map<String, dynamic>> {
  const PostCreateUserRequestConverter();

  @override
  PostCreateUserRequest fromJson(dynamic s) {
    return PostCreateUserRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostCreateUserRequest s) {
    return s.toJson();
  }
}

class PostCreateUserRequest {
  DateTime birthday;
  int gender;
  String id;
  String password;
  List<Role?>? roles;

  PostCreateUserRequest({
    required this.birthday,
    this.gender = 0,
    this.id = '',
    this.password = '',
    this.roles,
  });

  factory PostCreateUserRequest.fromJson(Map<String, dynamic> json) {
    return PostCreateUserRequest(
      birthday:
          const external_9298cca.DateTimeConverter().fromJson(json['Birthday']),
      gender: const external_9298cca.DoNothingConverter<int>()
          .fromJson(json['Gender']),
      id: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['ID']),
      password: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Password']),
      roles: const external_9298cca
                  .NullableConverter<List<Role?>, List<Map<String, dynamic>?>>(
              external_9298cca.ListConverter<Role?, Map<String, dynamic>?>(
                  external_9298cca.NullableConverter<Role,
                      Map<String, dynamic>>(RoleConverter())))
          .fromJson(json['Roles']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Birthday': const external_9298cca.DateTimeConverter().toJson(birthday),
      'Gender': const external_9298cca.DoNothingConverter<int>().toJson(gender),
      'ID': const external_9298cca.DoNothingConverter<String>().toJson(id),
      'Password':
          const external_9298cca.DoNothingConverter<String>().toJson(password),
      'Roles': const external_9298cca
                  .NullableConverter<List<Role?>, List<Map<String, dynamic>?>>(
              external_9298cca.ListConverter<Role?, Map<String, dynamic>?>(
                  external_9298cca.NullableConverter<Role,
                      Map<String, dynamic>>(RoleConverter())))
          .toJson(roles),
    };
  }
}

class PostCreateUserResponseConverter
    implements
        external_9298cca
            .JsonConverter<PostCreateUserResponse, Map<String, dynamic>> {
  const PostCreateUserResponseConverter();

  @override
  PostCreateUserResponse fromJson(dynamic s) {
    return PostCreateUserResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostCreateUserResponse s) {
    return s.toJson();
  }
}

class PostCreateUserResponse {
  CreatedType createdType;
  String message;
  DateTime requestedAt;
  bool status;

  PostCreateUserResponse({
    this.createdType = CreatedType.createdTypeGuest,
    this.message = '',
    required this.requestedAt,
    this.status = false,
  });

  factory PostCreateUserResponse.fromJson(Map<String, dynamic> json) {
    return PostCreateUserResponse(
      createdType: const CreatedTypeConverter().fromJson(json['CreatedType']),
      message: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Message']),
      requestedAt: const external_9298cca.DateTimeConverter()
          .fromJson(json['RequestedAt']),
      status: const external_9298cca.DoNothingConverter<bool>()
          .fromJson(json['Status']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'CreatedType': const CreatedTypeConverter().toJson(createdType),
      'Message':
          const external_9298cca.DoNothingConverter<String>().toJson(message),
      'RequestedAt':
          const external_9298cca.DateTimeConverter().toJson(requestedAt),
      'Status':
          const external_9298cca.DoNothingConverter<bool>().toJson(status),
    };
  }
}

class RoleConverter
    implements external_9298cca.JsonConverter<Role, Map<String, dynamic>> {
  const RoleConverter();

  @override
  Role fromJson(dynamic s) {
    return Role.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(Role s) {
    return s.toJson();
  }
}

class Role {
  int id;
  String name;
  List<Role>? recursionRoles;

  Role({
    this.id = 0,
    this.name = '',
    this.recursionRoles,
  });

  factory Role.fromJson(Map<String, dynamic> json) {
    return Role(
      id: const external_9298cca.DoNothingConverter<int>().fromJson(json['ID']),
      name: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Name']),
      recursionRoles: const external_9298cca
                  .NullableConverter<List<Role>, List<Map<String, dynamic>>>(
              external_9298cca.ListConverter<Role, Map<String, dynamic>>(
                  RoleConverter()))
          .fromJson(json['RecursionRoles']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': const external_9298cca.DoNothingConverter<int>().toJson(id),
      'Name': const external_9298cca.DoNothingConverter<String>().toJson(name),
      'RecursionRoles': const external_9298cca
                  .NullableConverter<List<Role>, List<Map<String, dynamic>>>(
              external_9298cca.ListConverter<Role, Map<String, dynamic>>(
                  RoleConverter()))
          .toJson(recursionRoles),
    };
  }
}
