// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../common.dart' as external_1ad882c;

class GetStaticPageRequestConverter
    implements
        external_1ad882c
            .JsonConverter<GetStaticPageRequest, Map<String, dynamic>> {
  const GetStaticPageRequestConverter();

  @override
  GetStaticPageRequest fromJson(dynamic s) {
    return GetStaticPageRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetStaticPageRequest s) {
    return s.toJson();
  }
}

class GetStaticPageRequest {
  GetStaticPageRequest();

  factory GetStaticPageRequest.fromJson(Map<String, dynamic> json) {
    return GetStaticPageRequest();
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{};
  }
}

class GetStaticPageResponseConverter
    implements
        external_1ad882c
            .JsonConverter<GetStaticPageResponse, Map<String, dynamic>> {
  const GetStaticPageResponseConverter();

  @override
  GetStaticPageResponse fromJson(dynamic s) {
    return GetStaticPageResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetStaticPageResponse s) {
    return s.toJson();
  }
}

class GetStaticPageResponse {
  String body;
  String title;

  GetStaticPageResponse({
    this.body = '',
    this.title = '',
  });

  factory GetStaticPageResponse.fromJson(Map<String, dynamic> json) {
    return GetStaticPageResponse(
      body: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['body']),
      title: const external_1ad882c.DoNothingConverter<String>()
          .fromJson(json['title']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'body': const external_1ad882c.DoNothingConverter<String>().toJson(body),
      'title':
          const external_1ad882c.DoNothingConverter<String>().toJson(title),
    };
  }
}
