// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../../common.dart' as external_2c4f542;

class GetUserJobGetRequestConverter
    implements
        external_2c4f542
            .JsonConverter<GetUserJobGetRequest, Map<String, dynamic>> {
  const GetUserJobGetRequestConverter();

  @override
  GetUserJobGetRequest fromJson(dynamic s) {
    return GetUserJobGetRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetUserJobGetRequest s) {
    return s.toJson();
  }
}

class GetUserJobGetRequest {
  String userID;

  GetUserJobGetRequest({
    this.userID = '',
  });

  factory GetUserJobGetRequest.fromJson(Map<String, dynamic> json) {
    return GetUserJobGetRequest(
      userID: const external_2c4f542.DoNothingConverter<String>()
          .fromJson(json['UserID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'UserID':
          const external_2c4f542.DoNothingConverter<String>().toJson(userID),
    };
  }
}

class GetUserJobGetResponseConverter
    implements
        external_2c4f542
            .JsonConverter<GetUserJobGetResponse, Map<String, dynamic>> {
  const GetUserJobGetResponseConverter();

  @override
  GetUserJobGetResponse fromJson(dynamic s) {
    return GetUserJobGetResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(GetUserJobGetResponse s) {
    return s.toJson();
  }
}

class GetUserJobGetResponse {
  String jobName;
  DateTime requestTime;

  GetUserJobGetResponse({
    this.jobName = '',
    required this.requestTime,
  });

  factory GetUserJobGetResponse.fromJson(Map<String, dynamic> json) {
    return GetUserJobGetResponse(
      jobName: const external_2c4f542.DoNothingConverter<String>()
          .fromJson(json['JobName']),
      requestTime: const external_2c4f542.DateTimeConverter()
          .fromJson(json['RequestTime']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'JobName':
          const external_2c4f542.DoNothingConverter<String>().toJson(jobName),
      'RequestTime':
          const external_2c4f542.DateTimeConverter().toJson(requestTime),
    };
  }
}
