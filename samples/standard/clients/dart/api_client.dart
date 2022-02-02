// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)
import 'dart:convert';
import 'package:http/http.dart' as http;
import './classes/service/groups/types.dart' as types__service_groups;
import './classes/service/static_page/types.dart' as types__service_static_page;
import './classes/service/types.dart' as types__service;
import './classes/service/user/types.dart' as types__service_user;
import './classes/service/user2/_userID/_JobID/types.dart'
    as types__service_user_2__user_id__job_id;
import './classes/service/user2/_userID/types.dart'
    as types__service_user_2__user_id;
import './classes/service/user2/types.dart' as types__service_user_2;
import './classes/types.dart' as types_;

class ServiceClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;
  late ServiceGroupsClient groups;
  late ServiceStaticPageClient static_page;
  late ServiceTableClient table;
  late ServiceUserClient user;
  late ServiceUser2Client user2;

  ServiceClient(
    this.baseURL,
    this.headers,
    this.client,
  ) {
    groups = ServiceGroupsClient(
      baseURL,
      headers,
      client,
    );
    static_page = ServiceStaticPageClient(
      baseURL,
      headers,
      client,
    );
    table = ServiceTableClient(
      baseURL,
      headers,
      client,
    );
    user = ServiceUserClient(
      baseURL,
      headers,
      client,
    );
    user2 = ServiceUser2Client(
      baseURL,
      headers,
      client,
    );
  }

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }

  Future<types__service.GetArticleResponse> getArticle(
      types__service.GetArticleRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    headers.remove('Content-Type');
    final url = baseURL +
        '/service/article' +
        Uri(queryParameters: getRequestObject(param.toJson(), excludeParams))
            .toString();
    final resp = await client.get(
      Uri.parse(url),
      headers: headers,
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res =
        types__service.GetArticleResponse.fromJson(jsonDecode(resp.body));

    return res;
  }
}

class ServiceGroupsClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;
  late ServiceGroupsCommonClient common;

  ServiceGroupsClient(
    this.baseURL,
    this.headers,
    this.client,
  ) {
    common = ServiceGroupsCommonClient(
      baseURL,
      headers,
      client,
    );
  }

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }

  Future<types__service_groups.GetGroupsResponse> getGroups(
      types__service_groups.GetGroupsRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    headers.remove('Content-Type');
    final url = baseURL +
        '/service/groups/groups' +
        Uri(queryParameters: getRequestObject(param.toJson(), excludeParams))
            .toString();
    final resp = await client.get(
      Uri.parse(url),
      headers: headers,
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res =
        types__service_groups.GetGroupsResponse.fromJson(jsonDecode(resp.body));

    return res;
  }
}

class ServiceGroupsCommonClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;

  ServiceGroupsCommonClient(
    this.baseURL,
    this.headers,
    this.client,
  );

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }
}

class ServiceStaticPageClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;

  ServiceStaticPageClient(
    this.baseURL,
    this.headers,
    this.client,
  );

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }

  Future<types__service_static_page.GetStaticPageResponse> getStaticPage(
      types__service_static_page.GetStaticPageRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    headers.remove('Content-Type');
    final url = baseURL +
        '/service/static_page/static_page' +
        Uri(queryParameters: getRequestObject(param.toJson(), excludeParams))
            .toString();
    final resp = await client.get(
      Uri.parse(url),
      headers: headers,
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_static_page.GetStaticPageResponse.fromJson(
        jsonDecode(resp.body));

    return res;
  }
}

class ServiceTableClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;

  ServiceTableClient(
    this.baseURL,
    this.headers,
    this.client,
  );

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }
}

class ServiceUser2Client {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;
  late ServiceUser2UserIDClient userID;

  ServiceUser2Client(
    this.baseURL,
    this.headers,
    this.client,
  ) {
    userID = ServiceUser2UserIDClient(
      baseURL,
      headers,
      client,
    );
  }

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }

  Future<types__service_user_2.DeleteUserResponse> deleteUser(
      types__service_user_2.DeleteUserRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[
      'id',
    ];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url =
        baseURL + '/service/user2/${Uri.encodeComponent(param.id.toString())}';

    final resp = await client.delete(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user_2.DeleteUserResponse();

    return res;
  }

  Future<types__service_user_2.GetUserResponse> getUser(
      types__service_user_2.GetUserRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[
      'id',
    ];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    headers.remove('Content-Type');
    final url = baseURL +
        '/service/user2/${Uri.encodeComponent(param.id.toString())}' +
        Uri(queryParameters: getRequestObject(param.toJson(), excludeParams))
            .toString();
    final resp = await client.get(
      Uri.parse(url),
      headers: headers,
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res =
        types__service_user_2.GetUserResponse.fromJson(jsonDecode(resp.body));

    return res;
  }

  Future<types__service_user_2.PostUpdateUserNameResponse> postUpdateUserName(
      types__service_user_2.PostUpdateUserNameRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/service/user2/update_user_name';

    final resp = await client.post(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user_2.PostUpdateUserNameResponse.fromJson(
        jsonDecode(resp.body));

    return res;
  }

  Future<types__service_user_2.PostUpdateUserPasswordResponse>
      postUpdateUserPassword(
          types__service_user_2.PostUpdateUserPasswordRequest param,
          {Map<String, String>? headers,
          http.Client? client,
          Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/service/user2/update_user_password';

    final resp = await client.post(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user_2.PostUpdateUserPasswordResponse.fromJson(
        jsonDecode(resp.body));

    return res;
  }
}

class ServiceUser2UserIDClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;
  late ServiceUser2UserIDJobIDClient JobID;

  ServiceUser2UserIDClient(
    this.baseURL,
    this.headers,
    this.client,
  ) {
    JobID = ServiceUser2UserIDJobIDClient(
      baseURL,
      headers,
      client,
    );
  }

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }

  Future<types__service_user_2__user_id.GetUserJobGetResponse> getUserJobGet(
      types__service_user_2__user_id.GetUserJobGetRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[
      'UserID',
    ];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    headers.remove('Content-Type');
    final url = baseURL +
        '/service/user2/${Uri.encodeComponent(param.userID.toString())}/user_job_get' +
        Uri(queryParameters: getRequestObject(param.toJson(), excludeParams))
            .toString();
    final resp = await client.get(
      Uri.parse(url),
      headers: headers,
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user_2__user_id.GetUserJobGetResponse.fromJson(
        jsonDecode(resp.body));

    return res;
  }
}

class ServiceUser2UserIDJobIDClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;

  ServiceUser2UserIDJobIDClient(
    this.baseURL,
    this.headers,
    this.client,
  );

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }

  Future<types__service_user_2__user_id__job_id.PutJobResponse> putJob(
      types__service_user_2__user_id__job_id.PutJobRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[
      'JobID',
      'UserID',
    ];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL +
        '/service/user2/${Uri.encodeComponent(param.userID.toString())}/${Uri.encodeComponent(param.jobID.toString())}/job';

    final resp = await client.put(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user_2__user_id__job_id.PutJobResponse.fromJson(
        jsonDecode(resp.body));

    return res;
  }
}

class ServiceUserClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;

  ServiceUserClient(
    this.baseURL,
    this.headers,
    this.client,
  );

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }

  Future<types__service_user.GetResponse> get(
      types__service_user.GetRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    headers.remove('Content-Type');
    final url = baseURL +
        '/service/user/' +
        Uri(queryParameters: getRequestObject(param.toJson(), excludeParams))
            .toString();
    final resp = await client.get(
      Uri.parse(url),
      headers: headers,
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user.GetResponse();

    return res;
  }

  Future<types__service_user.PostUpdateUserNameResponse> postUpdateUserName(
      types__service_user.PostUpdateUserNameRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/service/user/update_user_name';

    final resp = await client.post(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user.PostUpdateUserNameResponse.fromJson(
        jsonDecode(resp.body));

    return res;
  }

  Future<types__service_user.PostUpdateUserPasswordResponse>
      postUpdateUserPassword(
          types__service_user.PostUpdateUserPasswordRequest param,
          {Map<String, String>? headers,
          http.Client? client,
          Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/service/user/update_user_password';

    final resp = await client.post(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__service_user.PostUpdateUserPasswordResponse.fromJson(
        jsonDecode(resp.body));

    return res;
  }
}

class APIClient {
  Map<String, String> headers = {};
  String baseURL;
  late http.Client client;

  late ServiceClient service;

  APIClient(
    this.baseURL, {
    String? token,
    Map<String, String>? headers,
    http.Client? client,
  }) {
    if (headers != null) {
      this.headers = {...headers};
    }

    if (token != null) {
      this.headers['Authorization'] = 'Bearer ' + token;
    }

    this.client = client ?? http.Client();

    this.headers['Content-Type'] = 'application/json';

    service = ServiceClient(
      baseURL,
      this.headers,
      this.client,
    );
  }

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.forEach((key, value) {
      if (routingPath.contains(key))
        copied.remove(key);
      else
        copied[key] = value.toString();
    });

    return copied;
  }

  Future<types_.GetResponse> get(types_.GetRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    headers.remove('Content-Type');
    final url = baseURL +
        '/' +
        Uri(queryParameters: getRequestObject(param.toJson(), excludeParams))
            .toString();
    final resp = await client.get(
      Uri.parse(url),
      headers: headers,
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types_.GetResponse.fromJson(jsonDecode(resp.body));

    return res;
  }

  Future<types_.PostCreateTableResponse> postCreateTable(
      types_.PostCreateTableRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/create_table';

    final resp = await client.post(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types_.PostCreateTableResponse.fromJson(jsonDecode(resp.body));

    return res;
  }

  Future<types_.PostCreateUserResponse> postCreateUser(
      types_.PostCreateUserRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/create_user';

    final resp = await client.post(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types_.PostCreateUserResponse.fromJson(jsonDecode(resp.body));

    return res;
  }
}

class ApiError extends Error {
  final http.Response response;

  ApiError(this.response);

  int get statusCode => response.statusCode;
  String? get reasonPhrase => response.reasonPhrase;
}
