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

class GetArticleRequestConverter
    implements JsonConverter<GetArticleRequest, Map<String, dynamic>> {
  const GetArticleRequestConverter();

  @override
  GetArticleRequest fromJson(dynamic s) {
    return GetArticleRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetArticleRequest s) {
    return s.toJson();
  }
}

class GetArticleRequest {
  int id;

  GetArticleRequest({
    this.id = 0,
  });

  factory GetArticleRequest.fromJson(Map<String, dynamic> json) {
    return GetArticleRequest(
      id: const DoNothingConverter<int>().fromJson(json['ID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': const DoNothingConverter<int>().toJson(id),
    };
  }
}

class GetArticleResponseConverter
    implements JsonConverter<GetArticleResponse, Map<String, dynamic>> {
  const GetArticleResponseConverter();

  @override
  GetArticleResponse fromJson(dynamic s) {
    return GetArticleResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetArticleResponse s) {
    return s.toJson();
  }
}

class GetArticleResponse {
  String body;
  List<String>? group;
  int id;
  DateTime? requestTime;

  GetArticleResponse({
    this.body = '',
    this.group,
    this.id = 0,
    this.requestTime,
  });

  factory GetArticleResponse.fromJson(Map<String, dynamic> json) {
    return GetArticleResponse(
      body: const DoNothingConverter<String>().fromJson(json['Body']),
      group: const NullableConverter<List<String>, List<String>>(
              ListConverter<String, String>(DoNothingConverter<String>()))
          .fromJson(json['Group']),
      id: const DoNothingConverter<int>().fromJson(json['ID']),
      requestTime: const DateTimeConverter().fromJson(json['RequestTime']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Body': const DoNothingConverter<String>().toJson(body),
      'Group': const NullableConverter<List<String>, List<String>>(
              ListConverter<String, String>(DoNothingConverter<String>()))
          .toJson(group),
      'ID': const DoNothingConverter<int>().toJson(id),
      'RequestTime': const DateTimeConverter().toJson(requestTime),
    };
  }
}
