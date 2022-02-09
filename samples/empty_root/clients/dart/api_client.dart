// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)
import 'dart:convert';
import 'package:http/http.dart' as http;
import './classes/foo/bar/types.dart' as types__foo_bar;

class FooBarClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;

  FooBarClient(
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

  Future<types__foo_bar.PostUserResponse> postUser(
      types__foo_bar.PostUserRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/foo/bar/user';

    final resp = await client.post(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types__foo_bar.PostUserResponse();

    return res;
  }
}

class FooClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;
  late FooBarClient bar;

  FooClient(
    this.baseURL,
    this.headers,
    this.client,
  ) {
    bar = FooBarClient(
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
}

class APIClient {
  Map<String, String> headers = {};
  String baseURL;
  late http.Client client;

  late FooClient foo;

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

    foo = FooClient(
      baseURL,
      this.headers,
      this.client,
    );
  }

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};

    copied.removeWhere((key, value) => routingPath.contains(key));

    return copied;
  }
}

class ApiError extends Error {
  final http.BaseResponse response;

  ApiError(this.response);

  int get statusCode => response.statusCode;
  String? get reasonPhrase => response.reasonPhrase;
}
