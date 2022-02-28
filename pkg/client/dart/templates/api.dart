// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: {{ .AppVersion }}
import 'dart:convert';
import 'package:http/http.dart' as http;
{{- if .ImportHTTPParser }}
import 'package:http_parser/http_parser.dart' as http_parser;
{{- end }}
{{- range $index, $elem := .Imports }}
import '{{ $elem.Path }}.dart' as {{ $elem.Alias }};
{{- end }}

{{- range $index, $elem := .ChildrenClients }}

class {{ $elem.Name }} {
	Map<String, String> headers = {};
  String baseURL;
  http.Client client;

{{- 	range $index, $elem := .Children }}
	late {{ $elem.ClassName }} {{ $elem.Name }};
{{- 	end }}

  {{ $elem.Name }}(
    this.baseURL,
    this.headers,
    this.client,
  ){{ if .Children }} {
{{- 	range $index, $elem := .Children }}
    {{ $elem.Name }} = {{ $elem.ClassName }}(
			baseURL,
			headers,
			client,
		);
{{- 	end }}
  }{{ else }};{{ end }}

  Map<String, dynamic> getRequestObject(Map<String, dynamic> obj, List<String> routingPath, bool isGET) {
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

{{- range $index, $method := .Methods }}

	Future<{{ $method.ResponseType }}> {{ $method.Name }}(
		{{ $method.RequestType }} param,
    {
      Map<String, String>? headers,
      http.Client? client,
      Object? mockOption
    }
  ) async {
    client = client ?? this.client;

    final excludeParams = <String>[{{ range $param := $method.URLParams }}'{{ $param }}', {{ end }}];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
		}

{{- 	if eq $method.Method "GET" }}
    headers.remove('Content-Type');
		final url = baseURL + '{{ $method.Endpoint }}' + Uri(queryParameters: getRequestObject(param.toJson(), excludeParams, true)).toString();
{{- 	else }}
		final url = baseURL + '{{ $method.Endpoint }}';
{{      end }}

{{- 	if or (eq $method.Method "GET") (eq $method.Method "HEAD") }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
    );
{{- else if $method.Multipart }}
    final request = http.MultipartRequest('{{ $method.Method }}', Uri.parse(url))
      ..headers = headers
      ..files.add(http.MultipartFile.fromString(
        'x-multipart-json-binder-request-json', jsonEncode(getRequestObject(param.toJson(), excludeParams, false)),
        filename: 'x-multipart-json-binder-request-json', contentType: http_parser.MediaType.parse('application/json')
      ));
{{- range $index, $field := $method.FileFields }}
{{ if $field.IsArray }}
      param.{{ $field.StructField }}.forEach((http.MultipartFile file) {
        request.files.add(http.MultipartFile(
          '{{ $field.MultipartField }}', file.finalize(), file.length,
          filename: file.filename ?? 'untitled', contentType: file.contentType
        ));
      });
{{ else }}
    if (param.{{ $field.StructField }} != null) {
      final file = param.{{ $field.StructField }}!;
      request.files.add(http.MultipartFile(
        '{{ $field.MultipartField }}', file.finalize(), file.length,
        filename: file.filename ?? 'untitled', contentType: file.contentType
      ));
    }
{{- end }}
{{- end }}
    final resp = await client.send(request);
{{- else }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams, false)),
    );
{{- end }}

		if (resp.statusCode ~/ 100 != 2) {
			throw ApiError(resp);
		}

{{- if eq .HasFields true }}
		final res = {{ $method.ResponseType }}.fromJson(jsonDecode({{if .Multipart}}await resp.stream.bytesToString(){{else}}resp.body{{end}}));
{{- else }}
		final res = {{ $method.ResponseType }}();
{{- end }}

		return res;
	}
{{- end }}

}
{{- end }}

class APIClient {
	Map<String, String> headers = {};
	String baseURL;
  late http.Client client;

{{- 	range $index, $elem := .Children }}

	late {{ $elem.ClassName }} {{ $elem.Name }};
{{- 	end }}

	APIClient(
    this.baseURL,
    {
      String? token,
      Map<String, String>? headers,
      http.Client? client,
    }
	) {
    if (headers != null) {
      this.headers = {...headers};
    }

    if (token != null) {
      this.headers['Authorization'] = 'Bearer ' + token;
    }

    this.client = client ?? http.Client();

    this.headers['Content-Type'] = 'application/json';
{{- 	range $index, $elem := .Children }}

		{{ $elem.Name }} = {{ $elem.ClassName }}(
			baseURL,
			this.headers,
			this.client,
		);
{{- 	end }}
	}

	Map<String, dynamic> getRequestObject(Map<String, dynamic> obj, List<String> routingPath, bool isGET) {
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

{{- range $index, $method := .Methods }}

	Future<{{ $method.ResponseType }}> {{ $method.Name }}(
		{{ $method.RequestType }} param,
    {
      Map<String, String>? headers,
      http.Client? client,
      Object? mockOption
    }
  ) async {
    client = client ?? this.client;

    final excludeParams = <String>[{{ range $param := $method.URLParams }}'{{ $param }}', {{ end }}];

    headers = {...this.headers, ...headers ?? {}};

    if (mockOption != null) {
      headers['Api-Gen-Option'] = jsonEncode(mockOption);
		}

{{- 	if eq $method.Method "GET" }}
    headers.remove('Content-Type');
		final url = baseURL + '{{ $method.Endpoint }}' + Uri(queryParameters: getRequestObject(param.toJson(), excludeParams, true)).toString();
{{- 	else }}
		final url = baseURL + '{{ $method.Endpoint }}';
{{      end }}

{{- 	if or (eq $method.Method "GET") (eq $method.Method "HEAD") }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
    );
{{- else if $method.Multipart }}
    final request = http.MultipartRequest('{{ $method.Method }}', Uri.parse(url))
      ..headers.addAll(headers)
      ..files.add(http.MultipartFile.fromString(
        'x-multipart-json-binder-request-json', jsonEncode(getRequestObject(param.toJson(), excludeParams, false)),
        filename: 'x-multipart-json-binder-request-json', contentType: http_parser.MediaType.parse('application/json')
      ));
{{- range $index, $field := $method.FileFields }}
{{ if $field.IsArray }}
    if (param.{{ $field.StructField }} != null) {
      param.{{ $field.StructField }}!.forEach((http.MultipartFile file) {
        request.files.add(http.MultipartFile(
          '{{ $field.MultipartField }}', file.finalize(), file.length,
          filename: file.filename ?? 'untitled', contentType: file.contentType
        ));
      });
    }
{{ else }}
    if (param.{{ $field.StructField }} != null) {
      final file = param.{{ $field.StructField }}!;
      request.files.add(http.MultipartFile(
        '{{ $field.MultipartField }}', file.finalize(), file.length,
        filename: file.filename ?? 'untitled', contentType: file.contentType
      ));
    }
{{- end }}
{{- end }}
    final resp = await client.send(request);
{{- else }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams, false)),
    );
{{- end }}

		if (resp.statusCode ~/ 100 != 2) {
			throw ApiError(resp);
		}

{{- if eq .HasFields true }}
		final res = {{ $method.ResponseType }}.fromJson(jsonDecode({{if .Multipart}}await resp.stream.bytesToString(){{else}}resp.body{{end}}));
{{- else }}
		final res = {{ $method.ResponseType }}();
{{- end }}

		return res;
	}
{{- end }}
}

class ApiError extends Error {
	final http.BaseResponse response;

	ApiError(this.response);

	int get statusCode => response.statusCode;
  String? get reasonPhrase => response.reasonPhrase;
}
