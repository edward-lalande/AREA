import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import 'package:second_app/utils/my_secure_storage.dart';

final SecureStorageService stockToken = SecureStorageService();

Map<String, dynamic> servicesMap = {};
Map<String, dynamic> actionsMap = {};
Map<String, dynamic> reactionsMap = {};

Map<String, String> userData = {};
Map<String, dynamic> actionData = {};
Map<String, dynamic> reactionData = {};

List<Service> services = [];
List<ReactionService> reactions = [];

bool actionDone = false;
bool reactionDone = false;
bool isOAuthStarted = false;

String host = "http://10.0.2.2:8080";

void showCustomSnackBar(BuildContext context, String message, {Color backgroundColor = Colors.grey})
{
    ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(
            backgroundColor: backgroundColor,
            duration: const Duration(seconds: 3),
            content: Text(
                message,
            ),
        ),
    );
}

Future<bool> sendSignUp({Map<String, dynamic>? body,Map<String, String>? headers,required String url, required int delim,}) async
{
    try {
        final response = await http.post(
            Uri.parse(url),
            headers: headers,
            body: json.encode(body),
        );

        if (response.statusCode == 200) {

            final tmp = json.decode(response.body);
            final token = tmp["body"] is String ? tmp["body"] : tmp["body"]?["token"];

            if (token != null) {
                await stockToken.write("token", token);
                return true;
            } else {
                print("No token found in response.");
                return false;
            }
        } else {
            print('ERRRORR : ${response.statusCode}, ${response.body}');
            return false;
        }
    } catch (e) {
        print('ERRORRRRR : $e');
        return false;
    }
}

Future<bool> setupAreaArgs(Map<String, dynamic> actionData, List<Map<String, dynamic>> reactionsData) async
{
    String? token = await stockToken.read("token");

    final body = [{
        "user_token" : token,
        "action": actionData,
        "reactions": reactionsData,
    }];

    final success = await classicPost(
        url: "$host/areas",
        body: body,
    );

    if (success) {
        return true;
    } else {
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

Future<Map<String, dynamic>> fetchUserData(String url) async
{
    String? token = await stockToken.read("token");

    try {
        final response = await http.get(
            Uri.parse(url),
            headers: {
                "Content-Type": "application/json",
                "token": token ?? "",
            },
        );

        if (response.statusCode == 200) {
            final Map<String, dynamic> responseData = json.decode(response.body);
            return responseData['user'];
        } else {
            throw Exception('Error: ${response.statusCode}, ${response.body}');
        }
    } catch (e) {
        throw Exception('Failed to fetch user data: $e');
    }
}

Future<void> getDatas() async
{
    try {
        final String actionsString = await classicGet(
            url: "$host/actions",
        );
        final String reactionsString = await classicGet(
            url: "$host/reactions",
        );
        final String servString = await classicGet(
            url: "$host/services",
        );
        servicesMap = jsonDecode(servString);
        services = parseServices(actionsString);
        reactions = parseReactionServices(reactionsString);
    }
    catch (e) {
        throw Exception("ERORR: $e");
    }
}

class Service {

    final String name;
    final List<ActionServ> actions;

    Service({required this.name, required this.actions});


    factory Service.fromJson(Map<String, dynamic> json) {
        return Service(
            name: json['name'],
            actions: (json['actions'] as List)
                .map((action) => ActionServ.fromJson(action))
                .toList(),
        );
    }
}

class ActionServ {

    final int actionId;
    final int actionType;
    final String name;
    final String description;
    final List<Argument> arguments;

    ActionServ({
        required this.actionId,
        required this.actionType,
        required this.name,
        required this.description,
        required this.arguments,
    });

    factory ActionServ.fromJson(Map<String, dynamic> json)
    {
        return ActionServ(
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

class ReactionService {

    final String name;
    final List<Reaction> reactions;

    ReactionService({required this.name, required this.reactions});

    factory ReactionService.fromJson(Map<String, dynamic> json) {
        return ReactionService(
            name: json['name'],
            reactions: (json['reactions'] as List)
                .map((reaction) => Reaction.fromJson(reaction))
                .toList(),
        );
    }
}

class Reaction {

    final String name;
    final String description;
    final int reactionId;
    final int reactionType;
    final List<Argument> arguments;

    Reaction({
        required this.name,
        required this.description,
        required this.reactionId,
        required this.reactionType,
        required this.arguments,
    });

    factory Reaction.fromJson(Map<String, dynamic> json) {
        return Reaction(
            name: json['name'] ?? 'Unknown',
            description: json['description'] ?? '',
            reactionId: json['reaction_id'] ?? 0,
            reactionType: json['reaction_type'] ?? 0,
            arguments: (json['arguments'] as List?)?.map((arg) => Argument.fromJson(arg)).toList() ?? [],
        );
    }
}


List<ReactionService> parseReactionServices(String jsonString)
{
    final List<dynamic> jsonData = jsonDecode(jsonString);

    return jsonData
      .where((item) => item != null)
      .map((item) => ReactionService.fromJson(item))
      .toList();
}
