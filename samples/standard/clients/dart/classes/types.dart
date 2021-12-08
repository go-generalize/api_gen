// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import 'dart:convert';

import 'package:intl/intl.dart';

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
    return List<dynamic>.from(arr).map((e) => internalConverter.fromJson(e)).toList();
  }

  @override
  List<Base> toJson(List<T> arr) {
    return arr.map((e) => internalConverter.toJson(e) as Base).toList();
  }
}

class MapConverter<K, T, Base> implements JsonConverter<Map<K, T>, Map<K, Base>> {
  const MapConverter(this.internalConverter);

  final JsonConverter<T, Base> internalConverter;

  @override 
  Map<K, T> fromJson(dynamic m) {
    return Map<K, dynamic>.from(m).map((key, value) => MapEntry<K, T>(key, internalConverter.fromJson(value)));
  }

  @override
  Map<K, Base> toJson(Map<K, T> m) {
    return m.map((key, value) => MapEntry<K, Base>(key, internalConverter.toJson(value)));
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
    return (dt ?? DateTime.fromMillisecondsSinceEpoch(0, isUtc: true)).toUtc().toIso8601String();
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
    CreatedTypeGuest,
    CreatedTypeMember,
    CreatedTypeOwner,
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
    CreatedType.CreatedTypeGuest: 2,
    CreatedType.CreatedTypeMember: 1,
    CreatedType.CreatedTypeOwner: 0,
  };
  static final valueToEnum = {
    2: CreatedType.CreatedTypeGuest,
    1: CreatedType.CreatedTypeMember,
    0: CreatedType.CreatedTypeOwner,
  };

  static CreatedType fromJson(dynamic d) {
    return CreatedTypeExtension.valueToEnum[d]!;
  }

  int toJson() {
    return CreatedTypeExtension.enumToValue[this]!;
  }
}

enum Enum {
    EnumA,
    EnumB,
    EnumC,
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
    Enum.EnumA: 'A',
    Enum.EnumB: 'B',
    Enum.EnumC: 'C',
  };
  static final valueToEnum = {
    'A': Enum.EnumA,
    'B': Enum.EnumB,
    'C': Enum.EnumC,
  };

  static Enum fromJson(dynamic d) {
    return EnumExtension.valueToEnum[d]!;
  }

  String toJson() {
    return EnumExtension.enumToValue[this]!;
  }
}

class GetRequestConverter implements JsonConverter<GetRequest, Map<String, dynamic>> {
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
  Enum Enum;
  String Param;
  DateTime? Time;


  GetRequest({
    this.Enum = Enum.EnumA,
    this.Param = '',
    this.Time,
  });

  factory GetRequest.fromJson(Map<String, dynamic> json) {
    return GetRequest(
      Enum: EnumConverter().fromJson(json['Enum']),
      Param: DoNothingConverter<String>().fromJson(json['Param']),
      Time: DateTimeConverter().fromJson(json['Time']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Enum': EnumConverter().toJson(Enum),
      'Param': DoNothingConverter<String>().toJson(Param),
      'Time': DateTimeConverter().toJson(Time),
    };
  }
}

class GetResponseConverter implements JsonConverter<GetResponse, Map<String, dynamic>> {
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
  String Data;


  GetResponse({
    this.Data = '',
  });

  factory GetResponse.fromJson(Map<String, dynamic> json) {
    return GetResponse(
      Data: DoNothingConverter<String>().fromJson(json['Data']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Data': DoNothingConverter<String>().toJson(Data),
    };
  }
}

class PostCreateTableRequestConverter implements JsonConverter<PostCreateTableRequest, Map<String, dynamic>> {
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
  int Flag;
  String ID;
  Map<int, int> map;
  String Text;


  PostCreateTableRequest({
    this.Flag = 0,
    this.ID = '',
    this.map = const {},
    this.Text = '',
  });

  factory PostCreateTableRequest.fromJson(Map<String, dynamic> json) {
    return PostCreateTableRequest(
      Flag: DoNothingConverter<int>().fromJson(json['Flag']),
      ID: DoNothingConverter<String>().fromJson(json['ID']),
      map: MapConverter<int, int, int>(DoNothingConverter<int>()).fromJson(json['map']),
      Text: DoNothingConverter<String>().fromJson(json['Text']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Flag': DoNothingConverter<int>().toJson(Flag),
      'ID': DoNothingConverter<String>().toJson(ID),
      'map': MapConverter<int, int, int>(DoNothingConverter<int>()).toJson(map),
      'Text': DoNothingConverter<String>().toJson(Text),
    };
  }
}

class PostCreateTableResponseConverter implements JsonConverter<PostCreateTableResponse, Map<String, dynamic>> {
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
  String ID;
  DateTime? RequestTime;


  PostCreateTableResponse({
    this.ID = '',
    this.RequestTime,
  });

  factory PostCreateTableResponse.fromJson(Map<String, dynamic> json) {
    return PostCreateTableResponse(
      ID: DoNothingConverter<String>().fromJson(json['ID']),
      RequestTime: DateTimeConverter().fromJson(json['RequestTime']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': DoNothingConverter<String>().toJson(ID),
      'RequestTime': DateTimeConverter().toJson(RequestTime),
    };
  }
}

class PostCreateUserRequestConverter implements JsonConverter<PostCreateUserRequest, Map<String, dynamic>> {
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
  DateTime? Birthday;
  int Gender;
  String ID;
  String Password;
  List<Role?>? Roles;


  PostCreateUserRequest({
    this.Birthday,
    this.Gender = 0,
    this.ID = '',
    this.Password = '',
    this.Roles,
  });

  factory PostCreateUserRequest.fromJson(Map<String, dynamic> json) {
    return PostCreateUserRequest(
      Birthday: DateTimeConverter().fromJson(json['Birthday']),
      Gender: DoNothingConverter<int>().fromJson(json['Gender']),
      ID: DoNothingConverter<String>().fromJson(json['ID']),
      Password: DoNothingConverter<String>().fromJson(json['Password']),
      Roles: NullableConverter<List<Role?>, List<Map<String, dynamic>?>>(ListConverter<Role?, Map<String, dynamic>?>(NullableConverter<Role, Map<String, dynamic>>(RoleConverter()))).fromJson(json['Roles']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Birthday': DateTimeConverter().toJson(Birthday),
      'Gender': DoNothingConverter<int>().toJson(Gender),
      'ID': DoNothingConverter<String>().toJson(ID),
      'Password': DoNothingConverter<String>().toJson(Password),
      'Roles': NullableConverter<List<Role?>, List<Map<String, dynamic>?>>(ListConverter<Role?, Map<String, dynamic>?>(NullableConverter<Role, Map<String, dynamic>>(RoleConverter()))).toJson(Roles),
    };
  }
}

class PostCreateUserResponseConverter implements JsonConverter<PostCreateUserResponse, Map<String, dynamic>> {
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
  CreatedType CreatedType;
  String Message;
  DateTime? RequestedAt;
  bool Status;


  PostCreateUserResponse({
    this.CreatedType = CreatedType.CreatedTypeGuest,
    this.Message = '',
    this.RequestedAt,
    this.Status = false,
  });

  factory PostCreateUserResponse.fromJson(Map<String, dynamic> json) {
    return PostCreateUserResponse(
      CreatedType: CreatedTypeConverter().fromJson(json['CreatedType']),
      Message: DoNothingConverter<String>().fromJson(json['Message']),
      RequestedAt: DateTimeConverter().fromJson(json['RequestedAt']),
      Status: DoNothingConverter<bool>().fromJson(json['Status']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'CreatedType': CreatedTypeConverter().toJson(CreatedType),
      'Message': DoNothingConverter<String>().toJson(Message),
      'RequestedAt': DateTimeConverter().toJson(RequestedAt),
      'Status': DoNothingConverter<bool>().toJson(Status),
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
  int ID;
  String Name;
  List<Role>? RecursionRoles;


  Role({
    this.ID = 0,
    this.Name = '',
    this.RecursionRoles,
  });

  factory Role.fromJson(Map<String, dynamic> json) {
    return Role(
      ID: DoNothingConverter<int>().fromJson(json['ID']),
      Name: DoNothingConverter<String>().fromJson(json['Name']),
      RecursionRoles: NullableConverter<List<Role>, List<Map<String, dynamic>>>(ListConverter<Role, Map<String, dynamic>>(RoleConverter())).fromJson(json['RecursionRoles']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': DoNothingConverter<int>().toJson(ID),
      'Name': DoNothingConverter<String>().toJson(Name),
      'RecursionRoles': NullableConverter<List<Role>, List<Map<String, dynamic>>>(ListConverter<Role, Map<String, dynamic>>(RoleConverter())).toJson(RecursionRoles),
    };
  }
}

