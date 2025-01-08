import 'package:http/http.dart' as http;

import 'dart:convert';

import 'package:second_app/utils/my_secure_storage.dart';

final SecureStorageService stockData = SecureStorageService();
final host = StringBuffer("10.0.2.2");
bool isChanged = false;
Map<String, dynamic> servicesMap = {};
Map<String, dynamic> actionsMap = {};
Map<String, dynamic> reactionsMap = {};
Map<String, String> userData = {};

String parseGetToken(String body)
{
    StringBuffer result = StringBuffer();

    for (var i = 25; i <= body.length; i++) {
        if (body[i] == '"') {
            break;
        }
        result.write(body[i]);
    }
    return result.toString();
}

Future<bool> sendSignUp({Map<String, dynamic>? body, Map<String, String>? headers, required String url}) async
{

    try {
        final response = await http.post(
            Uri.parse(url),
            headers: headers,
            body: json.encode(body),
        );
        if (response.statusCode == 200) {
            print(response.body);
            stockData.write("token", parseGetToken(response.body));

            //final token = await stockData.read('token');
                //if (token != null) {
                //  print('Token: $token');
                //}

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

Future<bool> postArea({Map<String, String>? body, Map<String, String>? headers, required String url }) async
{
    try {

        final response = await http.post(
            Uri.parse(url),
            headers: headers,
            body: body,
        );

        if (response.statusCode == 200) {
            return true;
        } else {
            print('ERROR: ${response.statusCode}, ${response.body}');
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
