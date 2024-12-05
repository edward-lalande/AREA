import 'package:http/http.dart' as http;
import 'dart:convert';

Future<bool> sendSignUp({Map<String, dynamic>? body}) async {

    final response = await http.post(
        Uri.parse('http://127.0.0.1:8080/user'),
        body: json.encode(body),
    );
    if (response.statusCode == 200) {
        return true;
    } else {
        return false;
    }
}