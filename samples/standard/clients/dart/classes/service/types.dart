// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../common.dart' as external_254f1ba;

class GetArticleRequestConverter
    implements
        external_254f1ba
            .JsonConverter<GetArticleRequest, Map<String, dynamic>> {
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
      id: const external_254f1ba.DoNothingConverter<int>().fromJson(json['ID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'ID': const external_254f1ba.DoNothingConverter<int>().toJson(id),
    };
  }
}

class GetArticleResponseConverter
    implements
        external_254f1ba
            .JsonConverter<GetArticleResponse, Map<String, dynamic>> {
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
  DateTime requestTime;

  GetArticleResponse({
    this.body = '',
    this.group,
    this.id = 0,
    required this.requestTime,
  });

  factory GetArticleResponse.fromJson(Map<String, dynamic> json) {
    return GetArticleResponse(
      body: const external_254f1ba.DoNothingConverter<String>()
          .fromJson(json['Body']),
      group:
          const external_254f1ba.NullableConverter<List<String>, List<String>>(
                  external_254f1ba.ListConverter<String, String>(
                      external_254f1ba.DoNothingConverter<String>()))
              .fromJson(json['Group']),
      id: const external_254f1ba.DoNothingConverter<int>().fromJson(json['ID']),
      requestTime: const external_254f1ba.DateTimeConverter()
          .fromJson(json['RequestTime']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Body': const external_254f1ba.DoNothingConverter<String>().toJson(body),
      'Group':
          const external_254f1ba.NullableConverter<List<String>, List<String>>(
                  external_254f1ba.ListConverter<String, String>(
                      external_254f1ba.DoNothingConverter<String>()))
              .toJson(group),
      'ID': const external_254f1ba.DoNothingConverter<int>().toJson(id),
      'RequestTime':
          const external_254f1ba.DateTimeConverter().toJson(requestTime),
    };
  }
}
