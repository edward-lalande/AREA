import 'package:flutter/material.dart';
import 'dart:convert';
import '../utils/post_request.dart';
import 'package:http/http.dart' as http;

import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_title.dart';

class CreateArea extends StatefulWidget {
  const CreateArea({super.key});

  @override
  State<CreateArea> createState() => _CreateAreaState();
}

class _CreateAreaState extends State<CreateArea> {
    String action = "";
    String reaction = "";
    String page = "Create";

    int hour = 0;
    int minute = 0;

    String channel = "";
    String message = "";

    final List<String> services = ["Time"];
    final List<String> actions = ["Every day at"];
    final List<String> reactions = ["Send a message on channel"];

  void navigateToPage(String newPage) {
        setState(() {
        page = newPage;
        });
  }

  Future<void> createArea() async {
    String? accessToken = await storage.read(key: "accesToken");

    if (accessToken == null) {
        print("Access token not found.");
        return;
    }

    const String url = "http://10.0.2.2:8080/area";

    final data = [
        {
            "user_token": "AREA",
            "action": {
            "action_id": 1,
            "action_type": 0,
            "continent": "Europe",
            "city": "Paris",
            "hour": hour,
            "minute": minute
            },
            "reactions": [
            {
                "reaction_id": 2,
                "reaction_type": 0,
                "channel_id": channel,
                "message": message
            }
            ]
        }
    ];

    try {
        final response = await http.post(
            Uri.parse(url),
            headers: {
            "Authorization": "Bearer $accessToken",
            "Content-Type": "application/json",
            },
            body: jsonEncode(data),
        );

        if (response.statusCode == 200) {
            print("Area created successfully!");
        } else {
            print("Error: ${response.statusCode}, ${response.body}");
        }
    } catch (e) {
        print("Error: $e");
    }
  }

  @override
  Widget build(BuildContext context) {
        return SafeArea(
            child: Scaffold(
                backgroundColor: Colors.white,
                resizeToAvoidBottomInset: false,
                body: SingleChildScrollView(
                    physics: const AlwaysScrollableScrollPhysics(),
                    child: Column(
                        children: [
                            const MyTitle(
                                title: "AREA",
                                fontSize: 45,
                                padding: EdgeInsets.only(top: 80),
                                color: Colors.black
                            ),
                            const MyTitle(
                                title: "Create Area",
                                fontSize: 30,
                                padding: EdgeInsets.only(top: 30, bottom: 50),
                                color: Colors.black
                            ),
                            MyButton(
                                padding: const EdgeInsets.only(left: 35, right: 35, top: 60),
                                title: "If  this     (add)",
                                backgroundColor: Colors.black,
                                textColor: Colors.white,
                                fontSize: 30,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) {

                                },
                            ),
                            MyButton(
                                padding: const EdgeInsets.only(left: 35, right: 35, top: 30),
                                title: "Then that  (add)",
                                backgroundColor: Colors.grey,
                                textColor: Colors.white,
                                fontSize: 30,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) {

                                },
                            ),
                        ],
                    ),
                ),
            )
        );
  }
}
