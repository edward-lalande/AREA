import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../myWidgets/my_title.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_divider_text.dart';
import '../utils/post_request.dart';

class SignUpPage extends StatefulWidget {
    const SignUpPage({super.key});

    @override
    State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {

    final firstNameController = TextEditingController();
    final lastNameController = TextEditingController();
    final emailController = TextEditingController();
    final passwordController = TextEditingController();
    final _scrollController = ScrollController();

    @override
    Widget build(BuildContext context) {
        return SafeArea(
            child: Scaffold(
            resizeToAvoidBottomInset: false,
            backgroundColor: Colors.white,
            body: Padding(
                    padding: EdgeInsets.only(left: 8, right: 14),
                    child: RawScrollbar(
                        radius: Radius.circular(10),
                        thumbColor: Colors.black,
                        thickness: 5,
                        controller: _scrollController,
                        thumbVisibility: true,
                        child: SingleChildScrollView(
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
                            onPressed: (context) {
                                context.go("/login");
                            },
                            prefixIcon: const Icon(
                                size: 30,
                                Icons.arrow_back,
                            ),
                        ),
                        const MyTitle(
                            title: "AREA",
                            fontSize: 45,
                            padding: EdgeInsets.only(top: 80),
                            color: Colors.black
                        ),
                        const MyTitle(
                            title: "Sign Up",
                            fontSize: 30,
                            padding: EdgeInsets.only(top: 30, bottom: 50),
                            color: Colors.black
                        ),
                        MyTextField(
                            controller: firstNameController,
                            obscureText: false,
                            hintText: "First Name",
                            hintTextColor: Colors.black,
                            bgColor: Colors.white,
                            fieldBgColor: Colors.white,
                            padding: const EdgeInsets.only(
                                top: 0,
                                bottom: 0,
                                left: 35,
                                right: 35
                            ),
                            inputColor: Colors.black,
                            prefixIcon: const Icon(
                                Icons.person,
                                color: Colors.black,
                            ),
                        ),
                        MyTextField(
                            controller: lastNameController,
                            obscureText: false,
                            hintText: "Last Name",
                            hintTextColor: Colors.black,
                            bgColor: Colors.white,
                            fieldBgColor: Colors.white,
                            padding: const EdgeInsets.only(
                                top: 20,
                                bottom: 0,
                                left: 35,
                                right: 35
                            ),
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
                            padding: const EdgeInsets.only(top: 20, bottom: 0, left: 35, right: 35),
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
                            padding: const EdgeInsets.only(
                                top: 20,
                                bottom: 0,
                                left: 35,
                                right: 35
                            ),
                            inputColor: Colors.black,
                            prefixIcon: const Icon(
                                Icons.lock,
                                color: Colors.black,
                            ),
                        ),
                        // Bouton Sign Up
                        MyButton(
                            padding: const EdgeInsets.only(
                                left: 35,
                                right: 35,
                                top: 35
                            ),
                            title: "Sign Up",
                            backgroundColor: Colors.black,
                            textColor: Colors.white,
                            fontSize: 20,
                            spaceBetweenIconAndText: 10,

                            onPressed: (context) async {
                                if (firstNameController.text.isEmpty || lastNameController.text.isEmpty ||
                                emailController.text.isEmpty || passwordController.text.isEmpty) {
                                        ScaffoldMessenger.of(context).showSnackBar(
                                            SnackBar(
                                                backgroundColor: Colors.grey,
                                                duration: Duration(seconds: 3),
                                                content: Text(
                                                    'Please enter your Firstname, Lastname, email and your password',
                                                    style: TextStyle(color: Colors.white, fontFamily: "avenir"),
                                                ),
                                            ),
                                        );
                                        return;
                                }
                                userData['name'] = firstNameController.text;
                                userData['lastname'] = lastNameController.text;
                                userData['mail'] = emailController.text;
                                userData['password'] = passwordController.text;
                                bool tmp = await sendSignUp(
                                    url: "http://10.0.2.2:8080/sign-up",
                                    body: {
                                      "mail": emailController.text,
                                      "password": passwordController.text,
                                      "name": firstNameController.text,
                                      "lastname": lastNameController.text
                                    }
                                );
                                if (tmp) {
                                    if (context.mounted) {
                                    context.go("/login");
                                    }
                                } else {
                                    if (context.mounted) {
                                        context.go("/signup");
                                    }
                                }
                            }
                        ),
                        const MyDividerText(
                            bgColor: Colors.white,
                            textBetween: "Or",
                            padding: EdgeInsets.only(
                                top: 35,
                                right: 35,
                                left: 35
                            ),
                        ),
                        Container(
                            height: 130,
                            color: Colors.white,
                            child: Padding(
                                padding: const EdgeInsets.only(),
                                    child: Row(
                                        mainAxisAlignment: MainAxisAlignment.center,
                                        children: [
                                            const Text(
                                                "Already have an account?",
                                                style: TextStyle(
                                                fontFamily: "Avenir",
                                                fontSize: 16,
                                                ),
                                            ),
                                            const SizedBox(width: 5),
                                            GestureDetector(
                                                onTap: () {
                                                    context.go('/login');
                                                },
                                                child: const Text(
                                                    "Log in ",
                                                    style: TextStyle(
                                                        decoration: TextDecoration.underline,
                                                        color: Colors.blue,
                                                        fontFamily: "Avenir",
                                                        fontSize: 16,
                                                        decorationColor: Colors.blue,
                                                        decorationThickness: 2,
                                                    ),
                                                ),
                                            ),
                                        ],
                                    ),
                                ),
                            ),
                        ],
                    ),
                ),
            ),
            )
            )
        );
    }
}
