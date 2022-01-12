// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

abstract class JsonConverter<T, S> {
  const JsonConverter();

  T fromJson(dynamic json);
  S toJson(T object);
}

class ListConverter<T, Base> implements JsonConverter<List<T>, List<Base>> {
  const ListConverter(this.internalConverter);

  final JsonConverter<T, Base> internalConverter;

  @override
  List<T> fromJson(dynamic arr) {
    return List<dynamic>.from(arr)
        .map((e) => internalConverter.fromJson(e))
        .toList();
  }

  @override
  List<Base> toJson(List<T> arr) {
    return arr.map((e) => internalConverter.toJson(e)).toList();
  }
}

class MapConverter<K, T, Base>
    implements JsonConverter<Map<K, T>, Map<K, Base>> {
  const MapConverter(this.internalConverter);

  final JsonConverter<T, Base> internalConverter;

  @override
  Map<K, T> fromJson(dynamic m) {
    return Map<K, dynamic>.from(m).map(
        (key, value) => MapEntry<K, T>(key, internalConverter.fromJson(value)));
  }

  @override
  Map<K, Base> toJson(Map<K, T> m) {
    return m.map((key, value) =>
        MapEntry<K, Base>(key, internalConverter.toJson(value)));
  }
}

class DateTimeConverter implements JsonConverter<DateTime, String> {
  const DateTimeConverter();

  @override
  DateTime fromJson(dynamic s) {
    return DateTime.parse(s as String);
  }

  @override
  String toJson(DateTime? dt) {
    return (dt ?? DateTime.fromMillisecondsSinceEpoch(0, isUtc: true))
        .toUtc()
        .toIso8601String();
  }
}

class NullableConverter<T, Base> implements JsonConverter<T?, Base?> {
  const NullableConverter(this.internalConverter);

  final JsonConverter<T, Base> internalConverter;

  @override
  T? fromJson(dynamic s) {
    return s == null ? null : internalConverter.fromJson(s);
  }

  @override
  Base? toJson(T? dt) {
    return dt == null ? null : internalConverter.toJson(dt);
  }
}

class DoNothingConverter<T> implements JsonConverter<T, T> {
  const DoNothingConverter();

  @override
  T fromJson(dynamic s) {
    return s as T;
  }

  @override
  T toJson(T d) {
    return d;
  }
}

enum CreatedType {
  createdTypeGuest,
  createdTypeMember,
  createdTypeOwner,
}

class CreatedTypeConverter implements JsonConverter<CreatedType, int> {
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

class EnumConverter implements JsonConverter<Enum, String> {
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
    implements JsonConverter<GetRequest, Map<String, dynamic>> {
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
  DateTime? time;

  GetRequest({
    this.enum_ = Enum.enumA,
    this.param = '',
    this.time,
  });

  factory GetRequest.fromJson(Map<String, dynamic> json) {
    return GetRequest(
      enum_: const EnumConverter().fromJson(json['Enum']),
      param: const DoNothingConverter<String>().fromJson(json['Param']),
      time: const DateTimeConverter().fromJson(json['Time']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Enum': const EnumConverter().toJson(enum_),
      'Param': const DoNothingConverter<String>().toJson(param),
      'Time': const DateTimeConverter().toJson(time),
    };
  }
}

class GetResponseConverter
    implements JsonConverter<GetResponse, Map<String, dynamic>> {
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
      data: const DoNothingConverter<String>().fromJson(json['Data']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Data': const DoNothingConverter<String>().toJson(data),
    };
  }
}

class PostCreateTableRequestConverter
    implements JsonConverter<PostCreateTableRequest, Map<String, dynamic>> {
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
  int flag;
  String id;
  Map<int, int> map;
  String text;

  PostCreateTableRequest({
    this.flag = 0,
    this.id = '',
    this.map = const {},
    this.text = '',
  });

  factory PostCreateTableRequest.fromJson(Map<String, dynamic> json) {
    return PostCreateTableRequest(
      flag: const DoNothingConverter<int>().fromJson(json['Flag']),
      id: const DoNothingConverter<String>().fromJson(json['ID']),
      map: const MapConverter<int, int, int>(DoNothingConverter<int>())
          .fromJson(json['map']),
      text: const DoNothingConverter<String>().fromJson(json['Text']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Flag': const DoNothingConverter<int>().toJson(flag),
      'ID': const DoNothingConverter<String>().toJson(id),
      'map': const MapConverter<int, int, int>(DoNothingConverter<int>())
          .toJson(map),
      'Text': const DoNothingConverter<String>().toJson(text),
    };
  }
}

class PostCreateTableResponseConverter
    implements JsonConverter<PostCreateTableResponse, Map<String, dynamic>> {
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
  DateTime? requestTime;

  PostCreateTableResponse({
    this.id = '',
    this.requestTime,
  });

  factory PostCreateTableResponse.fromJson(Map<String, dynamic> json) {
    return PostCreateTableResponse(
      id: const DoNothingConverter<String>().fromJson(json['ID']),
      requestTime: const DateTimeConverter().fromJson(json['RequestTime']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': const DoNothingConverter<String>().toJson(id),
      'RequestTime': const DateTimeConverter().toJson(requestTime),
    };
  }
}

class PostCreateUserRequestConverter
    implements JsonConverter<PostCreateUserRequest, Map<String, dynamic>> {
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
  DateTime? birthday;
  int gender;
  String id;
  String password;
  List<Role?>? roles;

  PostCreateUserRequest({
    this.birthday,
    this.gender = 0,
    this.id = '',
    this.password = '',
    this.roles,
  });

  factory PostCreateUserRequest.fromJson(Map<String, dynamic> json) {
    return PostCreateUserRequest(
      birthday: const DateTimeConverter().fromJson(json['Birthday']),
      gender: const DoNothingConverter<int>().fromJson(json['Gender']),
      id: const DoNothingConverter<String>().fromJson(json['ID']),
      password: const DoNothingConverter<String>().fromJson(json['Password']),
      roles: const NullableConverter<List<Role?>, List<Map<String, dynamic>?>>(
              ListConverter<Role?, Map<String, dynamic>?>(
                  NullableConverter<Role, Map<String, dynamic>>(
                      RoleConverter())))
          .fromJson(json['Roles']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Birthday': const DateTimeConverter().toJson(birthday),
      'Gender': const DoNothingConverter<int>().toJson(gender),
      'ID': const DoNothingConverter<String>().toJson(id),
      'Password': const DoNothingConverter<String>().toJson(password),
      'Roles':
          const NullableConverter<List<Role?>, List<Map<String, dynamic>?>>(
                  ListConverter<Role?, Map<String, dynamic>?>(
                      NullableConverter<Role, Map<String, dynamic>>(
                          RoleConverter())))
              .toJson(roles),
    };
  }
}

class PostCreateUserResponseConverter
    implements JsonConverter<PostCreateUserResponse, Map<String, dynamic>> {
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
  DateTime? requestedAt;
  bool status;

  PostCreateUserResponse({
    this.createdType = CreatedType.createdTypeGuest,
    this.message = '',
    this.requestedAt,
    this.status = false,
  });

  factory PostCreateUserResponse.fromJson(Map<String, dynamic> json) {
    return PostCreateUserResponse(
      createdType: const CreatedTypeConverter().fromJson(json['CreatedType']),
      message: const DoNothingConverter<String>().fromJson(json['Message']),
      requestedAt: const DateTimeConverter().fromJson(json['RequestedAt']),
      status: const DoNothingConverter<bool>().fromJson(json['Status']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'CreatedType': const CreatedTypeConverter().toJson(createdType),
      'Message': const DoNothingConverter<String>().toJson(message),
      'RequestedAt': const DateTimeConverter().toJson(requestedAt),
      'Status': const DoNothingConverter<bool>().toJson(status),
    };
  }
}

class RoleConverter implements JsonConverter<Role, Map<String, dynamic>> {
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
      id: const DoNothingConverter<int>().fromJson(json['ID']),
      name: const DoNothingConverter<String>().fromJson(json['Name']),
      recursionRoles:
          const NullableConverter<List<Role>, List<Map<String, dynamic>>>(
                  ListConverter<Role, Map<String, dynamic>>(RoleConverter()))
              .fromJson(json['RecursionRoles']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': const DoNothingConverter<int>().toJson(id),
      'Name': const DoNothingConverter<String>().toJson(name),
      'RecursionRoles':
          const NullableConverter<List<Role>, List<Map<String, dynamic>>>(
                  ListConverter<Role, Map<String, dynamic>>(RoleConverter()))
              .toJson(recursionRoles),
    };
  }
}
