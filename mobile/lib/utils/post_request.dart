import 'package:http/http.dart' as http;
import 'dart:convert';

Future<bool> sendSignUp(String name, String lastname,
  String email, String password) async {

    final response = await http.post(
        Uri.parse('http://127.0.0.1:8080/user'),

        body: {
          "routes": "sign-up",
          "mail": email,
          "password": password,
          "name": name,
          "lastname": lastname,
        },
    );

    if (response.statusCode == 200) {
        return true;
    } else {
        return false;
    }
}