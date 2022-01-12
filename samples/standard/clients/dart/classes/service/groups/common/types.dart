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

class MetadataConverter
    implements JsonConverter<Metadata, Map<String, dynamic>> {
  const MetadataConverter();

  @override
  Metadata fromJson(dynamic s) {
    return Metadata.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(Metadata s) {
    return s.toJson();
  }
}

class Metadata {
  DateTime? createdAt;
  String id;
  String name;
  DateTime? updatedAt;

  Metadata({
    this.createdAt,
    this.id = '',
    this.name = '',
    this.updatedAt,
  });

  factory Metadata.fromJson(Map<String, dynamic> json) {
    return Metadata(
      createdAt: const DateTimeConverter().fromJson(json['CreatedAt']),
      id: const DoNothingConverter<String>().fromJson(json['ID']),
      name: const DoNothingConverter<String>().fromJson(json['Name']),
      updatedAt: const DateTimeConverter().fromJson(json['UpdatedAt']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'CreatedAt': const DateTimeConverter().toJson(createdAt),
      'ID': const DoNothingConverter<String>().toJson(id),
      'Name': const DoNothingConverter<String>().toJson(name),
      'UpdatedAt': const DateTimeConverter().toJson(updatedAt),
    };
  }
}
