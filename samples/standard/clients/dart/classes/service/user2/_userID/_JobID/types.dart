// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import '../../../../../common.dart' as external_47e933d;

class PutJobRequestConverter
    implements
        external_47e933d.JsonConverter<PutJobRequest, Map<String, dynamic>> {
  const PutJobRequestConverter();

  @override
  PutJobRequest fromJson(dynamic s) {
    return PutJobRequest.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PutJobRequest s) {
    return s.toJson();
  }
}

class PutJobRequest {
  String jobID;
  String userID;

  PutJobRequest({
    this.jobID = '',
    this.userID = '',
  });

  factory PutJobRequest.fromJson(Map<String, dynamic> json) {
    return PutJobRequest(
      jobID: const external_47e933d.DoNothingConverter<String>()
          .fromJson(json['JobID']),
      userID: const external_47e933d.DoNothingConverter<String>()
          .fromJson(json['UserID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'JobID':
          const external_47e933d.DoNothingConverter<String>().toJson(jobID),
      'UserID':
          const external_47e933d.DoNothingConverter<String>().toJson(userID),
    };
  }
}

class PutJobResponseConverter
    implements
        external_47e933d.JsonConverter<PutJobResponse, Map<String, dynamic>> {
  const PutJobResponseConverter();

  @override
  PutJobResponse fromJson(dynamic s) {
    return PutJobResponse.fromJson(Map<String, dynamic>.from(s));
  }

  @override
  Map<String, dynamic> toJson(PutJobResponse s) {
    return s.toJson();
  }
}

class PutJobResponse {
  String jobID;
  DateTime requestTime;
  String userID;

  PutJobResponse({
    this.jobID = '',
    required this.requestTime,
    this.userID = '',
  });

  factory PutJobResponse.fromJson(Map<String, dynamic> json) {
    return PutJobResponse(
      jobID: const external_47e933d.DoNothingConverter<String>()
          .fromJson(json['JobID']),
      requestTime: const external_47e933d.DateTimeConverter()
          .fromJson(json['RequestTime']),
      userID: const external_47e933d.DoNothingConverter<String>()
          .fromJson(json['UserID']),
    );
  }

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'JobID':
          const external_47e933d.DoNothingConverter<String>().toJson(jobID),
      'RequestTime':
          const external_47e933d.DateTimeConverter().toJson(requestTime),
      'UserID':
          const external_47e933d.DoNothingConverter<String>().toJson(userID),
    };
  }
}
