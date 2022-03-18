import 'package:http/http.dart';
import '../../clients/dart/api_client.dart' as api_client;
import '../../clients/dart/classes/types.dart' as types;
import '../../clients/dart/classes/_param/types.dart' as ptypes;

void main(List<String> args) async {
  final client = api_client.APIClient(args[0]);

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
