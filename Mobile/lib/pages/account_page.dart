import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/utils/post_request.dart';

import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';

class AccountPage extends StatefulWidget {
    @override
    _AccountPageState createState() => _AccountPageState();
}

class _AccountPageState extends State<AccountPage> {
    late TextEditingController firstNameController;
    late TextEditingController lastnameController;
    late TextEditingController emailController;
    late TextEditingController passwordController;

    @override
    void initState() {

        super.initState();
        firstNameController = TextEditingController(text: userData['name']);
        lastnameController = TextEditingController(text: userData['lastname']);
        emailController = TextEditingController(text: userData['mail']);
        passwordController = TextEditingController(text: userData['password']);

    }

    void updateUserInfo() {
        setState(() {

            userData['name'] = firstNameController.text;
            userData['lastname'] = lastnameController.text;
            userData['mail'] = emailController.text;
            userData['password'] = passwordController.text;

        });
  }

    @override
    Widget build(BuildContext context) {
        return SafeArea(
            child: Scaffold(
                backgroundColor: Colors.white,
                body: SingleChildScrollView(
                    physics: const AlwaysScrollableScrollPhysics(),
                    child: Column(
                        children: [
                            MyButton(
                                title: "",
                                backgroundColor: Colors.white,
                                textColor: Colors.black,
                                padding: const EdgeInsets.only(top: 30, left: 25),
                                fontSize: 0,
                                spaceBetweenIconAndText: 0,
                                prefixIcon: const Icon(
                                size: 30,
                                Icons.settings,
                                ),
                                onPressed: (context) {
                                context.go("/host");
                                },
                            ),
                            const MyTitle(
                                title: "Account",
                                fontSize: 30,
                                padding: EdgeInsets.only(top: 30, bottom: 50),
                                color: Colors.black,
                            ),
                            MyTextField(
                                controller: firstNameController,
                                obscureText: false,
                                hintText: "First Name",
                                hintTextColor: Colors.black,
                                bgColor: Colors.white,
                                fieldBgColor: Colors.white,
                                padding: const EdgeInsets.only(top: 50, bottom: 0, left: 35, right: 35),
                                inputColor: Colors.black,
                                prefixIcon: const Icon(
                                Icons.person,
                                color: Colors.black,
                                ),
                            ),
                            MyTextField(
                                controller: lastnameController,
                                obscureText: false,
                                hintText: "Last Name",
                                hintTextColor: Colors.black,
                                bgColor: Colors.white,
                                fieldBgColor: Colors.white,
                                padding: const EdgeInsets.only(top: 35, bottom: 0, left: 35, right: 35),
                                inputColor: Colors.black,
                                prefixIcon: const Icon(
                                Icons.person,
                                color: Colors.black,
                                ),
                            ),
                            MyTextField(
                                controller: emailController,
                                obscureText: false,
                                hintText: "Email",
                                hintTextColor: Colors.black,
                                bgColor: Colors.white,
                                fieldBgColor: Colors.white,
                                padding: const EdgeInsets.only(top: 35, bottom: 0, left: 35, right: 35),
                                inputColor: Colors.black,
                                prefixIcon: const Icon(
                                Icons.email,
                                color: Colors.black,
                                ),
                            ),
                            MyTextField(
                                controller: passwordController,
                                obscureText: true,
                                hintText: "Password",
                                hintTextColor: Colors.black,
                                bgColor: Colors.white,
                                fieldBgColor: Colors.white,
                                padding: const EdgeInsets.only(top: 35, bottom: 0, left: 35, right: 35),
                                inputColor: Colors.black,
                                prefixIcon: const Icon(
                                Icons.lock,
                                color: Colors.black,
                                ),
                            ),
                            MyButton(
                                padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
                                title: "Save Changes",
                                backgroundColor: Colors.black,
                                textColor: Colors.white,
                                fontSize: 20,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) {
                                    updateUserInfo();
                                },
                            ),
                            Container(
                                height: 100,
                                color: Colors.white,
                                child: Row(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                        const Text(
                                            "Log out ?",
                                            style: TextStyle(
                                                fontFamily: "Avenir",
                                                fontSize: 16,
                                            ),
                                        ),
                                        const SizedBox(width: 5),
                                        GestureDetector(
                                            onTap: () {
                                                userData.clear();
                                                context.go('/login');
                                            },
                                            child: const Text(
                                                "See you soon",
                                                style: TextStyle(
                                                    decoration: TextDecoration.underline,
                                                    color: Colors.blue,
                                                    fontFamily: "Avenir",
                                                    fontSize: 16,
                                                    decorationColor: Colors.blue,
                                                    decorationThickness: 2,
                                                ),
                                            ),
                                        )
                                    ],
                                ),
                            ),
                        ],
                    ),
                ),
            ),
        );
    }
}
