import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:http/http.dart' as http;

import 'package:second_app/myWidgets/my_text_button.dart';
import 'package:second_app/theme/theme_provider.dart';
import 'package:second_app/utils/post_request.dart';
import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_text_fields.dart';
import 'package:second_app/myWidgets/my_title.dart';

class AccountPage extends StatefulWidget {
    @override
    _AccountPageState createState() => _AccountPageState();
}

class _AccountPageState extends State<AccountPage> {

    late TextEditingController emailController;
    late TextEditingController nameController;
    late TextEditingController lastNameController;
    final _storage = FlutterSecureStorage();
    final scrollController = ScrollController();
    bool isLoading = true;


    final Map<Color, String> themeColors = {

        Colors.red: "Red",
        Colors.blue: "Blue",
        Colors.green: "Green",
        Colors.orange: "Orange",
        Colors.purple: "Purple",
        Colors.cyan: "Cyan",
        Colors.teal: "Teal",
        Colors.pink: "Pink",
        Colors.amber: "Amber",
        Colors.brown: "Brown",
        Colors.indigo: "Indigo",
        Colors.lime: "Lime",
        Colors.black: "Black",

    };


    @override
    void initState() {
        super.initState();
        emailController = TextEditingController(text: "");
        nameController = TextEditingController(text: "");
        lastNameController = TextEditingController(text: "");
        fetchAndSetUserData();
    }

    Future<void> updateUser({required String email, required String name, required String lastName,}) async
    {
        try {
            final token = await _storage.read(key: "token");
            final emailResponse = await http.post(
                Uri.parse("http://10.0.2.2:8085/update-email"),
                headers: {
                    'Content-Type': 'application/json',
                    "token": token??"",
                },
                body: json.encode({"mail": email})
            );

            if (emailResponse.statusCode != 200) {
                throw Exception("Failed to update email");
            }

            final nameResponse = await http.post(
                Uri.parse('http://10.0.2.2:8085/update-name'),
                headers: {
                    'Content-Type': 'application/json',
                    "token": token??"",
                },
                body: json.encode({"name": name})
            );
            if (nameResponse.statusCode != 200) {
                throw Exception("Failed to update name");
            }

            final lastNameResponse = await http.post(
                Uri.parse('http://10.0.2.2:8085/update-lastname'),
                headers: {
                    'Content-Type': 'application/json',
                    "token": token??"",
                },
                body: json.encode({"lastname": lastName})
            );
            if (lastNameResponse.statusCode != 200) {
                throw Exception("Failed to update lastname");
            }


        } catch (e) {

            throw Exception("Error updating user: $e");
        }
    }


    Future<void> fetchAndSetUserData() async {

        String url = "http://10.0.2.2:8085/user";

        try {
            final userData = await fetchUserData(url);

            setState(() {
                emailController = TextEditingController(text: userData['mail']);
                nameController = TextEditingController(text: userData['name'] ?? '');
                lastNameController = TextEditingController(text: userData['lastname'] ?? '');
                isLoading = false;
            });
        } catch (e) {
            showCustomSnackBar(context, "Failed to load user data.");

        }
    }

    @override
    Widget build(BuildContext context) {

    final themeProvider = Provider.of<ThemeProvider>(context);

    return Scaffold(
        backgroundColor: Theme.of(context).scaffoldBackgroundColor,
        body: isLoading
            ? Center(child: CircularProgressIndicator(color: Theme.of(context).textTheme.bodyLarge?.color))
            : Padding(
                padding: EdgeInsets.only(left: 8, right: 14),
                child: RawScrollbar(
                radius: Radius.circular(10),
                thumbColor: Theme.of(context).textTheme.bodyLarge?.color,
                thickness: 5,
                controller: scrollController,
                thumbVisibility: true,
                child: SingleChildScrollView(
                    controller: scrollController,
                    child: Column(
                        children: [
                            SizedBox(height: 100),
                            const MyTitle2(
                                title: "AREA",
                                fontSize: 45,
                                padding: EdgeInsets.only(top: 30),
                            ),
                            const MyTitle2(
                                title: "Account settings",
                                fontSize: 30,
                                padding: EdgeInsets.only(top: 30, bottom: 50),
                            ),
                            MyTextField2(
                                color: Theme.of(context).scaffoldBackgroundColor,
                                hintText: "Firstname",
                                controller: nameController,
                                prefixIcon: Icon(Icons.account_circle_sharp),
                            ),
                            SizedBox(height: 20),
                            MyTextField2(
                                color: Theme.of(context).scaffoldBackgroundColor,
                                hintText: "Lastname",
                                controller: lastNameController,
                                prefixIcon: Icon(Icons.account_circle_sharp),
                            ),
                            SizedBox(height: 20),
                            MyTextField2(
                                color: Theme.of(context).scaffoldBackgroundColor,
                                hintText: "Email",
                                controller: emailController,
                                prefixIcon: Icon(Icons.email),
                            ),
                            SizedBox(height: 30),
                            if (themeProvider.isDarkMode)
                                Row(
                                    children: [
                                        SizedBox(width: 20,),
                                        Text(
                                            "Choose a theme color:",
                                            style: TextStyle(
                                            color: Theme.of(context).textTheme.bodyLarge?.color,
                                            fontSize: 16,
                                            fontWeight: FontWeight.bold,
                                            ),
                                        ),

                                        DropdownButton<Color>(
                                            value: themeProvider.customDarkPrimaryColor,
                                            onChanged: (selectedColor) {
                                                if (selectedColor != null) {
                                                    themeProvider.updateCustomDarkPrimaryColor(selectedColor);
                                                }
                                            },
                                            items: themeColors.entries.map((entry) {
                                                final color = entry.key;
                                                final colorName = entry.value;
                                                return DropdownMenuItem<Color>(
                                                value: color,
                                                child: Row(
                                                    children: [
                                                        Container(
                                                            width: 40,
                                                            height: 20,
                                                            decoration: BoxDecoration(
                                                            color: color,
                                                            shape: BoxShape.circle,
                                                            ),
                                                        ),
                                                        SizedBox(width: 10),
                                                        Text(
                                                            colorName,
                                                            style: TextStyle(
                                                            color: Theme.of(context).textTheme.bodyLarge?.color,
                                                            ),
                                                        ),
                                                    ],
                                                ),
                                                );
                                            }).toList(),
                                        ),
                                    ],
                                ),
                            SizedBox(height: 30),
                            MyButton2(
                                title: "Save edit",
                                onPressed: (context) async {
                                    if (emailController.text.isEmpty || nameController.text.isEmpty || lastNameController.text.isEmpty) {
                                        showCustomSnackBar(context, "Please fill the fields.");
                                        return;
                                    }
                                    try {
                                        await updateUser(
                                            email: emailController.text,
                                            name: nameController.text,
                                            lastName: lastNameController.text,
                                        );
                                        showCustomSnackBar(context, "Informations have been saved.");
                                        if(context.mounted) context.go("/login");
                                    } catch (e) {
                                        showCustomSnackBar(context, "Failed to save user information.");
                                    }
                                },
                            ),
                            SizedBox(height: 30),
                            MyTextButton(
                                mainAxisAlignment: MainAxisAlignment.center,
                                onTap: (context) {
                                    context.go("/login");
                                },
                                firstTitle: "See you soon ?",
                                secondTitle: "Log out",
                                padding: EdgeInsets.only(left: 15),
                            ),
                            SizedBox(height: 30),
                        ],
                        ),
                    ),
                ),
            ),
        );
    }
}
