import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/utils/post_request.dart';

import '../myWidgets/my_title.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';

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
    final scrollController = ScrollController();

    @override
    Widget build(BuildContext context) {
        return Scaffold(
          backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            appBar: AppBar(
                shadowColor: Theme.of(context).scaffoldBackgroundColor,
                foregroundColor: Theme.of(context).scaffoldBackgroundColor,
                backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                surfaceTintColor: Theme.of(context).scaffoldBackgroundColor,
                elevation: 0,
                leading: Padding(
                    padding: const EdgeInsets.only(left: 20.0),
                    child: IconButton(
                        icon: const Icon(Icons.arrow_back),
                        onPressed: () {
                            context.go("/login");
                        },
                    ),
                ),
            ),
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
                        physics: const AlwaysScrollableScrollPhysics(),
                        child: Column(
                            children: [
                                const MyTitle2(
                                    title: "AREA",
                                    fontSize: 45,
                                    padding: EdgeInsets.only(top: 30),

                                ),
                                const MyTitle2(
                                    title: "Sign in",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),

                                ),
                                MyTextField2(
                                    color: Theme.of(context).scaffoldBackgroundColor,
                                    hintText: "First name",
                                    controller: firstNameController,
                                    prefixIcon: Icon(Icons.account_circle_sharp),

                                ),
                                SizedBox(height: 20),
                                MyTextField2(
                                    color: Theme.of(context).scaffoldBackgroundColor,
                                    hintText: "Last name",
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
                                    title: "Sign in",
                                    onPressed: (context) async {
                                         if (emailController.text.isEmpty || passwordController.text.isEmpty
                                         || lastNameController.text.isEmpty || firstNameController.text.isEmpty) {
                                            showCustomSnackBar(context, "Please fill all the fields");
                                            return;
                                        }
                                        await sendSignUp(
                                            delim: 18,
                                            url: "$host/sign-up",
                                            body: {
                                                "mail": emailController.text,
                                                "password": passwordController.text,
                                                "name": firstNameController.text,
                                                "lastname": lastNameController.text,
                                            }
                                        );
                                        context.go("/login");
                                    }
                                ),
                            ],
                        )
                    )
                ),
            ),
        );
    }
}
