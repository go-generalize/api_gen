// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../common.dart' as external_1ad882c;

class PosConverter
    implements external_1ad882c.JsonConverter<Pos, Map<String, dynamic>> {
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
      x: const external_1ad882c.DoNothingConverter<int>().fromJson(json['X']),
      y: const external_1ad882c.DoNothingConverter<int>().fromJson(json['Y']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'X': const external_1ad882c.DoNothingConverter<int>().toJson(x),
      'Y': const external_1ad882c.DoNothingConverter<int>().toJson(y),
    };
  }
}

class TableConverter
    implements external_1ad882c.JsonConverter<Table, Map<String, dynamic>> {
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
      pos: const PosConverter().fromJson(json['Pos']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Pos': const PosConverter().toJson(pos),
    };
  }
}
