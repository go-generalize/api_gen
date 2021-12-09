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


class GetUserJobGetRequestConverter implements JsonConverter<GetUserJobGetRequest, Map<String, dynamic>> {
  const GetUserJobGetRequestConverter();

  @override 
  GetUserJobGetRequest fromJson(dynamic s) {
    return GetUserJobGetRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetUserJobGetRequest s) {
    return s.toJson();
  }
}

class GetUserJobGetRequest {
  String userID;


  GetUserJobGetRequest({
    this.userID = '',
  });

  factory GetUserJobGetRequest.fromJson(Map<String, dynamic> json) {
    return GetUserJobGetRequest(
      userID: DoNothingConverter<String>().fromJson(json['UserID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'UserID': DoNothingConverter<String>().toJson(userID),
    };
  }
}

class GetUserJobGetResponseConverter implements JsonConverter<GetUserJobGetResponse, Map<String, dynamic>> {
  const GetUserJobGetResponseConverter();

  @override 
  GetUserJobGetResponse fromJson(dynamic s) {
    return GetUserJobGetResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetUserJobGetResponse s) {
    return s.toJson();
  }
}

class GetUserJobGetResponse {
  String jobName;
  DateTime? requestTime;


  GetUserJobGetResponse({
    this.jobName = '',
    this.requestTime,
  });

  factory GetUserJobGetResponse.fromJson(Map<String, dynamic> json) {
    return GetUserJobGetResponse(
      jobName: DoNothingConverter<String>().fromJson(json['JobName']),
      requestTime: DateTimeConverter().fromJson(json['RequestTime']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'JobName': DoNothingConverter<String>().toJson(jobName),
      'RequestTime': DateTimeConverter().toJson(requestTime),
    };
  }
}

