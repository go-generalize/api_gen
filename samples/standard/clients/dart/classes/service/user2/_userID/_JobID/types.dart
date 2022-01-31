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
  String toJson(DateTime dt) {
    return dt.toUtc().toIso8601String();
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

class PutJobRequestConverter
    implements JsonConverter<PutJobRequest, Map<String, dynamic>> {
  const PutJobRequestConverter();

  @override
  PutJobRequest fromJson(dynamic s) {
    return PutJobRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PutJobRequest s) {
    return s.toJson();
  }
}

class PutJobRequest {
  String jobID;
  String userID;

  PutJobRequest({
    this.jobID = '',
    this.userID = '',
  });

  factory PutJobRequest.fromJson(Map<String, dynamic> json) {
    return PutJobRequest(
      jobID: const DoNothingConverter<String>().fromJson(json['JobID']),
      userID: const DoNothingConverter<String>().fromJson(json['UserID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'JobID': const DoNothingConverter<String>().toJson(jobID),
      'UserID': const DoNothingConverter<String>().toJson(userID),
    };
  }
}

class PutJobResponseConverter
    implements JsonConverter<PutJobResponse, Map<String, dynamic>> {
  const PutJobResponseConverter();

  @override
  PutJobResponse fromJson(dynamic s) {
    return PutJobResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PutJobResponse s) {
    return s.toJson();
  }
}

class PutJobResponse {
  String jobID;
  DateTime requestTime;
  String userID;

  PutJobResponse({
    this.jobID = '',
    required this.requestTime,
    this.userID = '',
  });

  factory PutJobResponse.fromJson(Map<String, dynamic> json) {
    return PutJobResponse(
      jobID: const DoNothingConverter<String>().fromJson(json['JobID']),
      requestTime: const DateTimeConverter().fromJson(json['RequestTime']),
      userID: const DoNothingConverter<String>().fromJson(json['UserID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'JobID': const DoNothingConverter<String>().toJson(jobID),
      'RequestTime': const DateTimeConverter().toJson(requestTime),
      'UserID': const DoNothingConverter<String>().toJson(userID),
    };
  }
}
