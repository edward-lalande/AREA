import 'package:flutter/material.dart';
import 'dart:convert';
import '../utils/post_request.dart';
import 'package:http/http.dart' as http;

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

  // Fonction pour créer une "area"
  Future<void> createArea() async {
    // Utilise l'instance globale pour récupérer le token d'accès
    String? accessToken = await storage.read(key: "accesToken");

    if (accessToken == null) {
      // Si le token est absent, on ne peut pas faire la requête
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
          //"Authorization": "Bearer $accessToken",
          //"Content-Type": "application/json",
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
    return Scaffold(
      appBar: AppBar(title: Text('Create Area')),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          children: [
            if (page == "Create") ...[
              Text(
                "Create an area",
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              ),
              SizedBox(height: 20),
              GestureDetector(
                onTap: () => navigateToPage("Select service"),
                child: Container(
                  padding: EdgeInsets.all(16),
                  decoration: BoxDecoration(
                    color: Colors.black,
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Text(
                    action.isNotEmpty
                        ? "If: $action at $hour:$minute"
                        : "If this",
                    style: TextStyle(color: Colors.white, fontSize: 16),
                  ),
                ),
              ),
              SizedBox(height: 10),
              GestureDetector(
                onTap: action.isNotEmpty
                    ? () => navigateToPage("Select reaction")
                    : null,
                child: Container(
                  padding: EdgeInsets.all(16),
                  decoration: BoxDecoration(
                    color: action.isNotEmpty ? Colors.grey : Colors.grey[400],
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Text(
                    reaction.isNotEmpty
                        ? "Then: $reaction"
                        : "Then that",
                    style: TextStyle(color: Colors.white, fontSize: 16),
                  ),
                ),
              ),
              SizedBox(height: 20),
              if (action.isNotEmpty && reaction.isNotEmpty)
                ElevatedButton(
                  onPressed: createArea,
                  child: Text("Create Area"),
                ),
            ] else if (page == "Select service") ...[
              Text(
                "Select a service",
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              ),
              SizedBox(height: 20),
              for (var service in services)
                ListTile(
                  title: Text(service),
                  onTap: () => navigateToPage("Select action"),
                ),
            ] else if (page == "Select action") ...[
              Text(
                "Select an action",
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              ),
              SizedBox(height: 20),
              for (var act in actions)
                ListTile(
                  title: Text(act),
                  onTap: () => setState(() {
                    action = act;
                    navigateToPage("Select hour");
                  }),
                ),
            ] else if (page == "Select hour") ...[
              Text(
                "Every day at",
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              ),
              SizedBox(height: 20),
              TextField(
                decoration: InputDecoration(labelText: "Hour"),
                keyboardType: TextInputType.number,
                onChanged: (value) => hour = int.tryParse(value) ?? 0,
              ),
              TextField(
                decoration: InputDecoration(labelText: "Minute"),
                keyboardType: TextInputType.number,
                onChanged: (value) => minute = int.tryParse(value) ?? 0,
              ),
              SizedBox(height: 20),
              ElevatedButton(
                onPressed: () => navigateToPage("Create"),
                child: Text("Confirm"),
              ),
            ] else if (page == "Select reaction") ...[
              Text(
                "Select a reaction",
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              ),
              SizedBox(height: 20),
              for (var react in reactions)
                ListTile(
                  title: Text(react),
                  onTap: () => navigateToPage("Select channel"),
                ),
            ] else if (page == "Select channel") ...[
              Text(
                "Enter channel and message",
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
              ),
              SizedBox(height: 20),
              TextField(
                decoration: InputDecoration(labelText: "Channel (id)"),
                onChanged: (value) => channel = value,
              ),
              TextField(
                decoration: InputDecoration(labelText: "Message"),
                onChanged: (value) => message = value,
              ),
              SizedBox(height: 20),
              ElevatedButton(
                onPressed: () {
                  reaction = "Send a message on channel";
                  navigateToPage("Create");
                },
                child: Text("Confirm"),
              ),
            ]
          ],
        ),
      ),
    );
  }
}
