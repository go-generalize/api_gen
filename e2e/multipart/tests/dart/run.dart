import 'package:http/http.dart';
import '../../clients/dart/api_client.dart' as api_client;
import '../../clients/dart/classes/types.dart' as types;
import '../../clients/dart/classes/_param/types.dart' as ptypes;

Future<void> postA(String baseURL) async {
  final client = api_client.APIClient(baseURL);

  final req = types.PostARequest();

  req.file = MultipartFile.fromString('', '1');
  req.files = [
    MultipartFile.fromString('', '0', filename: '0.txt'),
    MultipartFile.fromString('', '1', filename: '1.txt'),
  ];
  req.payload = 'payload';

  final result = await client.postA(req);

  if (result.message != "OK") {
    throw "Unexpected result: " + result.message;
  }
}

Future<void> postB(String baseURL) async {
  final client = api_client.APIClient(baseURL);

  final req = ptypes.PostBRequest();

  req.file = MultipartFile.fromString('', '1');
  req.files = [
    MultipartFile.fromString('', '0', filename: '0.txt'),
    MultipartFile.fromString('', '1', filename: '1.txt'),
  ];
  req.payload = 'payload';
  req.param = 'param';

  final result = await client.param.postB(req);

  if (result.message != "OK") {
    throw "Unexpected result: " + result.message;
  }
}

void main(List<String> args) async {
  await postA(args[0]);
  await postB(args[0]);
}
