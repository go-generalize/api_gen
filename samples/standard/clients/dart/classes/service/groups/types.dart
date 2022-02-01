// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import 'common/types.dart' as external_ff6726e;

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

class CompanyConverter implements JsonConverter<Company, Map<String, dynamic>> {
  const CompanyConverter();

  @override
  Company fromJson(dynamic s) {
    return Company.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(Company s) {
    return s.toJson();
  }
}

class Company {
  external_ff6726e.Metadata metadata;

  Company({
    required this.metadata,
  });

  factory Company.fromJson(Map<String, dynamic> json) {
    return Company(
      metadata:
          const external_ff6726e.MetadataConverter().fromJson(json['Metadata']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Metadata': const external_ff6726e.MetadataConverter().toJson(metadata),
    };
  }
}

class DepartmentConverter
    implements JsonConverter<Department, Map<String, dynamic>> {
  const DepartmentConverter();

  @override
  Department fromJson(dynamic s) {
    return Department.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(Department s) {
    return s.toJson();
  }
}

class Department {
  external_ff6726e.Metadata metadata;

  Department({
    required this.metadata,
  });

  factory Department.fromJson(Map<String, dynamic> json) {
    return Department(
      metadata:
          const external_ff6726e.MetadataConverter().fromJson(json['Metadata']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Metadata': const external_ff6726e.MetadataConverter().toJson(metadata),
    };
  }
}

class GetGroupsRequestConverter
    implements JsonConverter<GetGroupsRequest, Map<String, dynamic>> {
  const GetGroupsRequestConverter();

  @override
  GetGroupsRequest fromJson(dynamic s) {
    return GetGroupsRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetGroupsRequest s) {
    return s.toJson();
  }
}

class GetGroupsRequest {
  GetGroupsRequest();

  factory GetGroupsRequest.fromJson(Map<String, dynamic> json) {
    return GetGroupsRequest();
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{};
  }
}

class GetGroupsResponseConverter
    implements JsonConverter<GetGroupsResponse, Map<String, dynamic>> {
  const GetGroupsResponseConverter();

  @override
  GetGroupsResponse fromJson(dynamic s) {
    return GetGroupsResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetGroupsResponse s) {
    return s.toJson();
  }
}

class GetGroupsResponse {
  List<Company>? companies;
  List<Department>? departments;

  GetGroupsResponse({
    this.companies,
    this.departments,
  });

  factory GetGroupsResponse.fromJson(Map<String, dynamic> json) {
    return GetGroupsResponse(
      companies:
          const NullableConverter<List<Company>, List<Map<String, dynamic>>>(
                  ListConverter<Company, Map<String, dynamic>>(
                      CompanyConverter()))
              .fromJson(json['Companies']),
      departments:
          const NullableConverter<List<Department>, List<Map<String, dynamic>>>(
                  ListConverter<Department, Map<String, dynamic>>(
                      DepartmentConverter()))
              .fromJson(json['Departments']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Companies':
          const NullableConverter<List<Company>, List<Map<String, dynamic>>>(
                  ListConverter<Company, Map<String, dynamic>>(
                      CompanyConverter()))
              .toJson(companies),
      'Departments':
          const NullableConverter<List<Department>, List<Map<String, dynamic>>>(
                  ListConverter<Department, Map<String, dynamic>>(
                      DepartmentConverter()))
              .toJson(departments),
    };
  }
}
