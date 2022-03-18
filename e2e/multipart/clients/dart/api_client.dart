// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:http_parser/http_parser.dart' as http_parser;
import './classes/_param/types.dart' as types___param;
import './classes/types.dart' as types_;

class ParamClient {
  Map<String, String> headers = {};
  String baseURL;
  http.Client client;

  ParamClient(
    this.baseURL,
    this.headers,
    this.client,
  );

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath, bool isGET) {
    final Map<String, dynamic> copied = {};

    obj.forEach((key, value) {
      if (routingPath.contains(key))
        return;
      else if (isGET) {
        if (value == null) return;
        copied[key] = value.toString();
        return;
      }
      copied[key] = value;
    });

    return copied;
  }

  Future<types___param.PostBResponse> postB(types___param.PostBRequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[
      'Param',
    ];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/${Uri.encodeComponent(param.param.toString())}/b';

    final request = http.MultipartRequest('POST', Uri.parse(url))
      ..headers.addAll(headers)
      ..files.add(http.MultipartFile.fromString(
          'x-multipart-json-binder-request-json',
          jsonEncode(getRequestObject(param.toJson(), excludeParams, false)),
          filename: 'x-multipart-json-binder-request-json',
          contentType: http_parser.MediaType.parse('application/json')));

    if (param.file != null) {
      final file = param.file!;
      request.files.add(http.MultipartFile('file', file.finalize(), file.length,
          filename: file.filename ?? 'untitled',
          contentType: file.contentType));
    }

    if (param.files != null) {
      param.files!.forEach((http.MultipartFile file) {
        request.files.add(http.MultipartFile(
            'files', file.finalize(), file.length,
            filename: file.filename ?? 'untitled',
            contentType: file.contentType));
      });
    }

    final resp = await client.send(request);

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types___param.PostBResponse.fromJson(
        jsonDecode(await resp.stream.bytesToString()));

    return res;
  }
}

class APIClient {
  Map<String, String> headers = {};
  String baseURL;
  late http.Client client;

  late ParamClient param;

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

    param = ParamClient(
      baseURL,
      this.headers,
      this.client,
    );
  }

  Map<String, dynamic> getRequestObject(
      Map<String, dynamic> obj, List<String> routingPath, bool isGET) {
    final Map<String, dynamic> copied = {};

    obj.forEach((key, value) {
      if (routingPath.contains(key))
        return;
      else if (isGET) {
        if (value == null) return;
        copied[key] = value.toString();
        return;
      }
      copied[key] = value;
    });

    return copied;
  }

  Future<types_.PostAResponse> postA(types_.PostARequest param,
      {Map<String, String>? headers,
      http.Client? client,
      Object? mockOption}) async {
    client = client ?? this.client;

    final excludeParams = <String>[];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
    }
    final url = baseURL + '/a';

    final request = http.MultipartRequest('POST', Uri.parse(url))
      ..headers.addAll(headers)
      ..files.add(http.MultipartFile.fromString(
          'x-multipart-json-binder-request-json',
          jsonEncode(getRequestObject(param.toJson(), excludeParams, false)),
          filename: 'x-multipart-json-binder-request-json',
          contentType: http_parser.MediaType.parse('application/json')));

    if (param.file != null) {
      final file = param.file!;
      request.files.add(http.MultipartFile('file', file.finalize(), file.length,
          filename: file.filename ?? 'untitled',
          contentType: file.contentType));
    }

    if (param.files != null) {
      param.files!.forEach((http.MultipartFile file) {
        request.files.add(http.MultipartFile(
            'files', file.finalize(), file.length,
            filename: file.filename ?? 'untitled',
            contentType: file.contentType));
      });
    }

    final resp = await client.send(request);

    if (resp.statusCode ~/ 100 != 2) {
      throw ApiError(resp);
    }
    final res = types_.PostAResponse.fromJson(
        jsonDecode(await resp.stream.bytesToString()));

    return res;
  }
}

class ApiError extends Error {
  final http.BaseResponse response;

  ApiError(this.response);

  int get statusCode => response.statusCode;
  String? get reasonPhrase => response.reasonPhrase;
}
