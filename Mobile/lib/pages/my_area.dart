import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:second_app/myWidgets/my_title.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:second_app/utils/post_request.dart';

class MyArea extends StatefulWidget {
    const MyArea({super.key});

    @override
    State<MyArea> createState() => _MyAreaState();
}

class _MyAreaState extends State<MyArea> {

    final _storage = FlutterSecureStorage();
    List<dynamic> areas = [];
    bool isLoading = true;

    @override
    void initState() {
        super.initState();
        _fetchAreas();
    }

    Future<void> _fetchAreas() async
    {
        String url = "http://$host:8080/areas";

        try {

            final token = await _storage.read(key: "token");
            if (token == null) {
                throw Exception("No token found in secure storage.");
            }

            final response = await http.get(
                Uri.parse(url),
                headers: {
                    "Content-Type": "application/json",
                    "token": token,
                },
            );

            if (response.statusCode == 200) {
                final Map<String, dynamic> data = json.decode(response.body);
                setState(() {
                    areas = data["areas"] ?? [];
                    isLoading = false;
                });
            } else {
                throw Exception('Error fetching areas: ${response.statusCode}');
            }
        } catch (e) {
            setState(() {
                isLoading = false;
            });
            print("Error: $e");
            ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(content: Text("Failed to fetch areas: $e")),
            );
        }
    }

    @override
    Widget build(BuildContext context) {
        final scrollController = ScrollController();

        return Scaffold(
            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            body: Padding(
                padding: const EdgeInsets.only(left: 8, right: 14),
                child: RawScrollbar(
                    radius: Radius.circular(10),
                    thumbColor: Theme.of(context).primaryColor,
                    thickness: 5,
                    controller: scrollController,
                    thumbVisibility: true,
                    child: SingleChildScrollView(
                        controller: scrollController,
                        physics: const AlwaysScrollableScrollPhysics(),
                        child: Column(
                            children: [
                                const SizedBox(height: 100),
                                const MyTitle2(
                                    title: "AREA",
                                    fontSize: 45,
                                    padding: EdgeInsets.only(top: 30),
                                ),
                                const MyTitle2(
                                    title: "My Area's",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),
                                ),
                                const SizedBox(height: 20),
                                if (isLoading)
                                    const Center(child: CircularProgressIndicator())
                                else if (areas.isEmpty)
                                const Center(
                                    child: Padding(
                                    padding: EdgeInsets.all(16.0),
                                    child: Text(
                                        "No Areas yet.",
                                        style: TextStyle(fontSize: 18),
                                    ),
                                    ),
                                )
                                else
                                ListView.builder(
                                    shrinkWrap: true,
                                    physics: const NeverScrollableScrollPhysics(),
                                    itemCount: areas.length,
                                    itemBuilder: (context, index) {
                                        final area = areas[index];
                                        return Card(
                                            elevation: 3,
                                            margin: const EdgeInsets.symmetric(vertical: 8),
                                            child: ListTile(
                                                title: Text(
                                                    "If ${area['action_name']} Then ${area['reaction_name']}",
                                                    style: const TextStyle(fontWeight: FontWeight.bold),
                                                ),
                                                subtitle: Text("Area ID: ${area['area_id']}"),
                                            ),
                                        );
                                    },
                                ),
                            ],
                        ),
                    ),
                ),
            ),
        );
    }
}