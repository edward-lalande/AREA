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

        emailController = TextEditingController(text: userData['mail']);
        passwordController = TextEditingController(text: userData['password']);

    }

    void updateUserInfo() {
        setState(() {

            userData['mail'] = emailController.text;
            userData['password'] = passwordController.text;

        });
  }

    @override
    Widget build(BuildContext context) {
        return Scaffold(
          backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            body: Padding(
                padding: EdgeInsets.only(left: 8, right: 14),
                child: RawScrollbar(
                        radius: Radius.circular(10),
                        thumbColor: Colors.black,
                        thickness: 5,
                        controller: scrollController,
                        thumbVisibility: true,
                        child: SingleChildScrollView(
                            controller: scrollController,
                            child: Column(
                                children: [
                                    SizedBox(height: 100,),
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
                                            ScaffoldMessenger.of(context).showSnackBar(
                                                SnackBar(
                                                    backgroundColor: Colors.grey,
                                                    duration: Duration(seconds: 3),
                                                    content: Text(
                                                        'Informations saved.',
                                                        style: TextStyle(color: Colors.white, fontFamily: "avenir"),
                                                    ),
                                                ),
                                            );
                                        }
                                    ),
                                    SizedBox(height: 30),
                                    MyTextButton(
                                        mainAxisAlignment: MainAxisAlignment.center,
                                        onTap: (context) {
                                            context.go("/login");

                                        },
                                        firstTitle: "See you soon ?", secondTitle: "Log out",
                                        padding: EdgeInsets.only(left: 15),
                                    )
                                ],
                        ),
                    ),
                ),
            )
        );
    }
}
