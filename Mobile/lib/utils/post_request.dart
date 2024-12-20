import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;

import 'dart:convert';

final storage = const FlutterSecureStorage();
Map<String, dynamic> servicesMap = {};

Future<bool> sendSignUp({Map<String, dynamic>? body, Map<String, String>? headers, required String url}) async
{
    try {
        final response = await http.post(
            Uri.parse(url),
            headers: headers,
            body: json.encode(body),
        );
        if (response.statusCode == 200) {

            storage.write(key: "accesToken", value: response.body);
            return true;

        } else {
            print('ERRRORR : ${response.statusCode}, ${response.body}');
            return false;
        }
    } catch (e) {
        print('ERRORRRRR : $e');
        return false;
    }
}

Future<bool> classicPost({Map<String, dynamic>? body, Map<String, String>? headers, required String url}) async
{
    try {
        final response = await http.post(
            Uri.parse(url),
            headers: headers,
            body: json.encode(body),
        );
        if (response.statusCode == 200) {
            return true;

        } else {
            print('ERRRORR : ${response.statusCode}, ${response.body}');
            return false;
        }
    } catch (e) {
        print('ERRORRRRR : $e');
        return false;
    }
}

Future<String> classicGet({required String url}) async
{

    final apiUrl = url;

    try {

      final response = await http.get(Uri.parse(apiUrl));

      if (response.statusCode == 200) {
        return response.body;
      }
      else {
        throw Exception('ERRORRR: ${response.statusCode}');
      }

    } catch (e) {

      throw Exception('ERRORRR: $e');

    }
}
