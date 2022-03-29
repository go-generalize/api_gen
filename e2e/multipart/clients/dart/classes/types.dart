// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../common.dart' as external_9298cca;
import 'package:http/http.dart' as external_d8d61af;

class PostARequestConverter
    implements
        external_9298cca.JsonConverter<PostARequest, Map<String, dynamic>> {
  const PostARequestConverter();

  @override
  PostARequest fromJson(dynamic s) {
    return PostARequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostARequest s) {
    return s.toJson();
  }
}

class PostARequest {
  external_d8d61af.MultipartFile? file;
  List<external_d8d61af.MultipartFile>? files;
  String payload;

  PostARequest({
    this.file,
    this.files,
    this.payload = '',
  });

  factory PostARequest.fromJson(Map<String, dynamic> json) {
    return PostARequest(
      payload: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Payload']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Payload':
          const external_9298cca.DoNothingConverter<String>().toJson(payload),
    };
  }
}

class PostAResponseConverter
    implements
        external_9298cca.JsonConverter<PostAResponse, Map<String, dynamic>> {
  const PostAResponseConverter();

  @override
  PostAResponse fromJson(dynamic s) {
    return PostAResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PostAResponse s) {
    return s.toJson();
  }
}

class PostAResponse {
  String message;

  PostAResponse({
    this.message = '',
  });

  factory PostAResponse.fromJson(Map<String, dynamic> json) {
    return PostAResponse(
      message: const external_9298cca.DoNothingConverter<String>()
          .fromJson(json['Message']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'Message':
          const external_9298cca.DoNothingConverter<String>().toJson(message),
    };
  }
}
