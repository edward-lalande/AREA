import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_text_button.dart';
import 'package:second_app/utils/post_request.dart';

import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';

class AccountPage extends StatefulWidget {
    @override
    _AccountPageState createState() => _AccountPageState();
}

class _AccountPageState extends State<AccountPage> {

    late TextEditingController emailController;
    late TextEditingController passwordController;
    final scrollController = ScrollController();

    @override
    void initState() {
        super.initState();
        emailController = TextEditingController(text: "");
        passwordController = TextEditingController(text: "");
        fetchAndSetUserData();
    }

    Future<void> fetchAndSetUserData() async {

        const String url = "http://10.0.2.2:8085/user";
        //String? token = await stockData.read("token");

        try {
            final userData = await fetchUserData(url);


            setState(() {
                emailController = TextEditingController(text: userData['mail']);
                passwordController = TextEditingController(text: userData['password'] ?? '');
            });
        } catch (e) {
            print("Failed to fetch user data: $e");
            showCustomSnackBar(context, "Failed to load user data.");
        }
    }

    @override
    Widget build(BuildContext context) {
        return Scaffold(
            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            body: Padding(
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
                                    hintText: "Email",
                                    controller: emailController,
                                    prefixIcon: Icon(Icons.email),
                                ),
                                SizedBox(height: 20),
                                MyTextField2(
                                    color: Theme.of(context).scaffoldBackgroundColor,
                                    hintText: "Password",
                                    controller: passwordController,
                                    obscureText: true,
                                    prefixIcon: Icon(Icons.lock),
                                ),
                                SizedBox(height: 30),
                                MyButton2(
                                    title: "Save",
                                    onPressed: (context) {
                                        if (emailController.text.isEmpty || passwordController.text.isEmpty) {
                                            showCustomSnackBar(context, "Please fill the fields.");
                                            return;
                                        }
                                        showCustomSnackBar(context, "Informations have been saved.");
                                    },
                                ),
                                SizedBox(height: 30),
                                MyTextButton(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    onTap: (context) {
                                        context.go("/login");
                                    },
                                    firstTitle: "See you soon ?", secondTitle: "Log out",
                                    padding: EdgeInsets.only(left: 15),
                                ),
                            ],
                        ),
                    ),
                ),
            ),
        );
    }
}
