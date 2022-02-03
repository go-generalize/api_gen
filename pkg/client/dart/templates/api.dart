// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: {{ .AppVersion }}
import 'dart:convert';
import 'package:http/http.dart' as http;
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

  Map<String, dynamic> getRequestObject(Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};
    
    copied.removeWhere((key, value) => routingPath.contains(key));

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
		final url = baseURL + '{{ $method.Endpoint }}' + Uri(queryParameters: getRequestObject(param.toJson(), excludeParams)).toString();
{{- 	else }}
		final url = baseURL + '{{ $method.Endpoint }}';
{{      end }}

{{- 	if or (eq $method.Method "GET") (eq $method.Method "HEAD") }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
    );
{{- else if $method.Multipart }}
    final request = http.MultipartRequest('$method.Method', Uri.parse(url))
      ..headers = headers
      ..files.add(http.MultipartFile.fromString(
        'x-multipart-json-binder-request-json', jsonEncode(getRequestObject(param.toJson(), excludeParams)),
        filename: 'x-multipart-json-binder-request-json', contentType: contentType
      )
{{- range $index, $field := $method.FileFieldNames }}
      ..files.add(param.$field)
{{- end }};
    var response = await client.send(request);

    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );
{{- else }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );
{{- end }}

		if (resp.statusCode ~/ 100 != 2) {
			throw ApiError(resp);
		}

{{- if eq .HasFields true }}
		final res = {{ $method.ResponseType }}.fromJson(jsonDecode({{if .Multipart}}resp.stream.byteToString(){{else}}resp.body{{end}}));
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

	Map<String, dynamic> getRequestObject(Map<String, dynamic> obj, List<String> routingPath) {
    final copied = {...obj};
    
    copied.removeWhere((key, value) => routingPath.contains(key));

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
		final url = baseURL + '{{ $method.Endpoint }}' + Uri(queryParameters: getRequestObject(param.toJson(), excludeParams)).toString();
{{- 	else }}
		final url = baseURL + '{{ $method.Endpoint }}';
{{      end }}

{{- 	if or (eq $method.Method "GET") (eq $method.Method "HEAD") }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
    );
{{- else if $method.Multipart }}
    final request = http.MultipartRequest('$method.Method', Uri.parse(url))
      ..headers = headers
      ..files.add(http.MultipartFile.fromString(
        'x-multipart-json-binder-request-json', jsonEncode(getRequestObject(param.toJson(), excludeParams)),
        filename: 'x-multipart-json-binder-request-json', contentType: contentType
      ))
{{- range $index, $field := $method.FileFieldNames }}
      ..files.add(param.{{ $field }})
{{- end }};
    final resp = await client.send(request);
{{- else }}
    final resp = await client.{{ toLower $method.Method }}(
      Uri.parse(url),
      headers: headers,
      body: jsonEncode(getRequestObject(param.toJson(), excludeParams)),
    );
{{- end }}

		if (resp.statusCode ~/ 100 != 2) {
			throw ApiError(resp);
		}

{{- if eq .HasFields true }}
		final res = {{ $method.ResponseType }}.fromJson(jsonDecode({{if .Multipart}}resp.stream.byteToString(){{else}}resp.body{{end}}));
{{- else }}
		final res = {{ $method.ResponseType }}();
{{- end }}

		return res;
	}
{{- end }}
}

class ApiError extends Error {
	final http.Response response;

	ApiError(this.response);

	int get statusCode => response.statusCode;
  String? get reasonPhrase => response.reasonPhrase;
}
