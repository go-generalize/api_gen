// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../../common.dart' as external_2c4f542;

class MetadataConverter
    implements external_2c4f542.JsonConverter<Metadata, Map<String, dynamic>> {
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
  DateTime createdAt;
  DateTime? deletedAt;
  String id;
  String name;
  DateTime updatedAt;

  Metadata({
    required this.createdAt,
    this.deletedAt,
    this.id = '',
    this.name = '',
    required this.updatedAt,
  });

  factory Metadata.fromJson(Map<String, dynamic> json) {
    return Metadata(
      createdAt: const external_2c4f542.DateTimeConverter()
          .fromJson(json['CreatedAt']),
      deletedAt: const external_2c4f542.NullableConverter<DateTime, String>(
              external_2c4f542.DateTimeConverter())
          .fromJson(json['DeletedAt']),
      id: const external_2c4f542.DoNothingConverter<String>()
          .fromJson(json['ID']),
      name: const external_2c4f542.DoNothingConverter<String>()
          .fromJson(json['Name']),
      updatedAt: const external_2c4f542.DateTimeConverter()
          .fromJson(json['UpdatedAt']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'CreatedAt': const external_2c4f542.DateTimeConverter().toJson(createdAt),
      'DeletedAt': const external_2c4f542.NullableConverter<DateTime, String>(
              external_2c4f542.DateTimeConverter())
          .toJson(deletedAt),
      'ID': const external_2c4f542.DoNothingConverter<String>().toJson(id),
      'Name': const external_2c4f542.DoNothingConverter<String>().toJson(name),
      'UpdatedAt': const external_2c4f542.DateTimeConverter().toJson(updatedAt),
    };
  }
}
