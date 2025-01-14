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
List<Service> services = [];

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
            stockData.write("token", parseGetToken(response.body));
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

Future<bool> classicPost({List<Map<String, dynamic>>? body, Map<String, String>? headers, required String url}) async
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

class Service {

    final String name;
    final List<Action> actions;

    Service({required this.name, required this.actions});


    factory Service.fromJson(Map<String, dynamic> json) {
        return Service(
            name: json['name'],
            actions: (json['actions'] as List)
                .map((action) => Action.fromJson(action))
                .toList(),
        );
    }
}

class Action {

    final int actionId;
    final int actionType;
    final String name;
    final String description;
    final List<Argument> arguments;

    Action({
        required this.actionId,
        required this.actionType,
        required this.name,
        required this.description,
        required this.arguments,
    });

    factory Action.fromJson(Map<String, dynamic> json)
    {
        return Action(
            actionId: json['action_id'],
            actionType: json['action_type'],
            name: json['name'],
            description: json['description'],
            arguments: (json['arguments'] as List)
                .map((arg) => Argument.fromJson(arg))
                .toList(),
        );
    }
}

class Argument {

    final String display;
    final String name;
    final String type;

    Argument({
        required this.display,
        required this.name,
        required this.type
    });

    factory Argument.fromJson(Map<String, dynamic> json)
    {
        return Argument(
            display: json['display'],
            name: json['name'],
            type: json['type'],
        );
    }
}


List<Service> parseServices(String jsonString)
{
    final List<dynamic> jsonData = jsonDecode(jsonString);
    return jsonData.map((service) => Service.fromJson(service)).toList();
}
