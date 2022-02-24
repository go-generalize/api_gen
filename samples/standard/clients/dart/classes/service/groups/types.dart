// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../common.dart' as external_1ad882c;
import 'common/types.dart' as external_ff6726e;

class CompanyConverter
    implements external_1ad882c.JsonConverter<Company, Map<String, dynamic>> {
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
    implements
        external_1ad882c.JsonConverter<Department, Map<String, dynamic>> {
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
    implements
        external_1ad882c.JsonConverter<GetGroupsRequest, Map<String, dynamic>> {
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
    implements
        external_1ad882c
            .JsonConverter<GetGroupsResponse, Map<String, dynamic>> {
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
      companies: const external_1ad882c
                  .NullableConverter<List<Company>, List<Map<String, dynamic>>>(
              external_1ad882c.ListConverter<Company, Map<String, dynamic>>(
                  CompanyConverter()))
          .fromJson(json['Companies']),
      departments: const external_1ad882c.NullableConverter<List<Department>,
                  List<Map<String, dynamic>>>(
              external_1ad882c.ListConverter<Department, Map<String, dynamic>>(
                  DepartmentConverter()))
          .fromJson(json['Departments']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Companies': const external_1ad882c
                  .NullableConverter<List<Company>, List<Map<String, dynamic>>>(
              external_1ad882c.ListConverter<Company, Map<String, dynamic>>(
                  CompanyConverter()))
          .toJson(companies),
      'Departments': const external_1ad882c.NullableConverter<List<Department>,
                  List<Map<String, dynamic>>>(
              external_1ad882c.ListConverter<Department, Map<String, dynamic>>(
                  DepartmentConverter()))
          .toJson(departments),
    };
  }
}
