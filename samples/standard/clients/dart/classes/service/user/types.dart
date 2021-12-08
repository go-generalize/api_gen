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


  GetRequest();

  factory GetRequest.fromJson(Map<String, dynamic> json) {
    return GetRequest(
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
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


  GetResponse();

  factory GetResponse.fromJson(Map<String, dynamic> json) {
    return GetResponse(
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
    };
  }
}

class PostUpdateUserNameRequestConverter implements JsonConverter<PostUpdateUserNameRequest, Map<String, dynamic>> {
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
  String Name;


  PostUpdateUserNameRequest({
    this.Name = '',
  });

  factory PostUpdateUserNameRequest.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserNameRequest(
      Name: DoNothingConverter<String>().fromJson(json['Name']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Name': DoNothingConverter<String>().toJson(Name),
    };
  }
}

class PostUpdateUserNameResponseConverter implements JsonConverter<PostUpdateUserNameResponse, Map<String, dynamic>> {
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
  String Message;
  DateTime? RequestTime;
  bool Status;


  PostUpdateUserNameResponse({
    this.Message = '',
    this.RequestTime,
    this.Status = false,
  });

  factory PostUpdateUserNameResponse.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserNameResponse(
      Message: DoNothingConverter<String>().fromJson(json['Message']),
      RequestTime: DateTimeConverter().fromJson(json['RequestTime']),
      Status: DoNothingConverter<bool>().fromJson(json['Status']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Message': DoNothingConverter<String>().toJson(Message),
      'RequestTime': DateTimeConverter().toJson(RequestTime),
      'Status': DoNothingConverter<bool>().toJson(Status),
    };
  }
}

class PostUpdateUserPasswordRequestConverter implements JsonConverter<PostUpdateUserPasswordRequest, Map<String, dynamic>> {
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
  String Password;
  String PasswordConfirm;


  PostUpdateUserPasswordRequest({
    this.Password = '',
    this.PasswordConfirm = '',
  });

  factory PostUpdateUserPasswordRequest.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserPasswordRequest(
      Password: DoNothingConverter<String>().fromJson(json['Password']),
      PasswordConfirm: DoNothingConverter<String>().fromJson(json['PasswordConfirm']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Password': DoNothingConverter<String>().toJson(Password),
      'PasswordConfirm': DoNothingConverter<String>().toJson(PasswordConfirm),
    };
  }
}

class PostUpdateUserPasswordResponseConverter implements JsonConverter<PostUpdateUserPasswordResponse, Map<String, dynamic>> {
  const PostUpdateUserPasswordResponseConverter();

  @override 
  PostUpdateUserPasswordResponse fromJson(dynamic s) {
    return PostUpdateUserPasswordResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostUpdateUserPasswordResponse s) {
    return s.toJson();
  }
}

class PostUpdateUserPasswordResponse {
  String Message;
  DateTime? RequestTime;
  bool Status;


  PostUpdateUserPasswordResponse({
    this.Message = '',
    this.RequestTime,
    this.Status = false,
  });

  factory PostUpdateUserPasswordResponse.fromJson(Map<String, dynamic> json) {
    return PostUpdateUserPasswordResponse(
      Message: DoNothingConverter<String>().fromJson(json['Message']),
      RequestTime: DateTimeConverter().fromJson(json['RequestTime']),
      Status: DoNothingConverter<bool>().fromJson(json['Status']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Message': DoNothingConverter<String>().toJson(Message),
      'RequestTime': DateTimeConverter().toJson(RequestTime),
      'Status': DoNothingConverter<bool>().toJson(Status),
    };
  }
}

