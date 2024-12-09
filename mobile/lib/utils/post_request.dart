import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

Future<bool> sendSignUp({Map<String, dynamic>? body}) async
{

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


Future<String> getOAuthUrl(String apiUrl) async
{
    try {

      final response = await http.get(Uri.parse(apiUrl));

      if (response.statusCode == 200) {
        return response.body;
      }
      else {
        throw Exception('Failed to load data: ${response.statusCode}');
      }

    } catch (e) {

      throw Exception('Error fetching data: $e');

    }
}

//class exeption + secure+strorage

void logByOAuth(BuildContext context, String apiUrl) async
{
    try {
        String fetchedUrl = await getOAuthUrl(apiUrl);

        if (context.mounted) {
            context.go(fetchedUrl);
        }
    } catch (e) {
        print("ERRORRRR");
    }
}