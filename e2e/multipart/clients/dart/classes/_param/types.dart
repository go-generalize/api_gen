// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../common.dart' as external_254f1ba;
import 'package:http/http.dart' as external_d8d61af;

class PostBRequestConverter
    implements
        external_254f1ba.JsonConverter<PostBRequest, Map<String, dynamic>> {
  const PostBRequestConverter();

  @override
  PostBRequest fromJson(dynamic s) {
    return PostBRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostBRequest s) {
    return s.toJson();
  }
}

class PostBRequest {
  external_d8d61af.MultipartFile? file;
  List<external_d8d61af.MultipartFile>? files;
  String param;
  String payload;

  PostBRequest({
    this.file,
    this.files,
    this.param = '',
    this.payload = '',
  });

  factory PostBRequest.fromJson(Map<String, dynamic> json) {
    return PostBRequest(
      param: const external_254f1ba.DoNothingConverter<String>()
          .fromJson(json['Param']),
      payload: const external_254f1ba.DoNothingConverter<String>()
          .fromJson(json['payload']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Param':
          const external_254f1ba.DoNothingConverter<String>().toJson(param),
      'payload':
          const external_254f1ba.DoNothingConverter<String>().toJson(payload),
    };
  }
}

class PostBResponseConverter
    implements
        external_254f1ba.JsonConverter<PostBResponse, Map<String, dynamic>> {
  const PostBResponseConverter();

  @override
  PostBResponse fromJson(dynamic s) {
    return PostBResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostBResponse s) {
    return s.toJson();
  }
}

class PostBResponse {
  String message;

  PostBResponse({
    this.message = '',
  });

  factory PostBResponse.fromJson(Map<String, dynamic> json) {
    return PostBResponse(
      message: const external_254f1ba.DoNothingConverter<String>()
          .fromJson(json['message']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'message':
          const external_254f1ba.DoNothingConverter<String>().toJson(message),
    };
  }
}
