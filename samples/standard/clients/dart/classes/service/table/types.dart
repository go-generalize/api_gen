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
    return arr.map((e) => internalConverter.toJson(e) as Base).toList();
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

class PosConverter implements JsonConverter<Pos, Map<String, dynamic>> {
  const PosConverter();

  @override
  Pos fromJson(dynamic s) {
    return Pos.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(Pos s) {
    return s.toJson();
  }
}

class Pos {
  int x;
  int y;

  Pos({
    this.x = 0,
    this.y = 0,
  });

  factory Pos.fromJson(Map<String, dynamic> json) {
    return Pos(
      x: DoNothingConverter<int>().fromJson(json['X']),
      y: DoNothingConverter<int>().fromJson(json['Y']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'X': DoNothingConverter<int>().toJson(x),
      'Y': DoNothingConverter<int>().toJson(y),
    };
  }
}

class TableConverter implements JsonConverter<Table, Map<String, dynamic>> {
  const TableConverter();

  @override
  Table fromJson(dynamic s) {
    return Table.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(Table s) {
    return s.toJson();
  }
}

class Table {
  Pos pos;

  Table({
    required this.pos,
  });

  factory Table.fromJson(Map<String, dynamic> json) {
    return Table(
      pos: PosConverter().fromJson(json['Pos']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Pos': PosConverter().toJson(pos),
    };
  }
}
