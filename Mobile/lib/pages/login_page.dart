import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_switch_button.dart';
import 'package:second_app/myWidgets/my_text_button.dart';
import 'package:second_app/myWidgets/oauth2_button.dart';
import 'package:second_app/utils/post_request.dart';

import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';
import '../myWidgets/my_divider_text.dart';

class LoginPage extends StatelessWidget {
    @override
    Widget build(BuildContext context) {

        final scrollController = ScrollController();
        final emailController = TextEditingController();
        final passwordController = TextEditingController();

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
                        icon: const Icon(Icons.settings),
                        onPressed: () {
                            context.go("/host");
                        },
                    ),
                ),
                actions: [
                    Padding(
                        padding: const EdgeInsets.only(right: 35.0),
                        child: const MySwitchButton(padding: EdgeInsets.all(8.0)),
                    ),
                ],
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
                                    title: "Log in",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),

                                ),
                                MyTextField2(
                                    color: Theme.of(context).scaffoldBackgroundColor,

                                    hintText: "Email",
                                    controller: emailController,
                                    prefixIcon: Icon(Icons.email),

                                ),
                                SizedBox(height: 30),
                                MyTextField2(
                                    color: Theme.of(context).scaffoldBackgroundColor,

                                    hintText: "Password",
                                    controller: passwordController,
                                    obscureText: true,
                                    prefixIcon: Icon(Icons.lock),

                                ),
                                SizedBox(height: 7,),
                                MyTextButton(
                                    firstTitle: "Forget your're",
                                    secondTitle: "Password",
                                    onTap: (context) {
                                        context.go("/password");
                                    },
                                    padding: EdgeInsets.only(left: 20)
                                ),
                                SizedBox(height: 20,),
                                MyButton2(
                                    title: "Log in",
                                    onPressed: (context) async {
                                        if (emailController.text.isEmpty ||
                                        passwordController.text.isEmpty) {
                                            showCustomSnackBar(
                                                context,
                                                "Please enter your email and password."
                                            );
                                            return;
                                        }
                                        bool res = await sendSignUp(
                                            delim: 18,
                                            url: '$host/login',
                                            body: {
                                                "mail": emailController.text,
                                                "password": passwordController.text
                                            }
                                        );
                                        await getDatas();
                                        if (res) {
                                            if (context.mounted) {
                                                context.go("/home");
                                            }
                                        } else {
                                            showCustomSnackBar(context, "Wrong email or password.");
                                            return;
                                        }
                                    }
                                ),
                                const MyDividerText2(
                                    padding: EdgeInsets.only(top: 35, right: 35, left: 35, bottom: 35),
                                    textBetween: "Or continue with",
                                ),
                                OAuthButtonsRow(host: host.toString()),
                                MyTextButton(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    onTap: (context) {
                                        context.go("/signup");
                                    }, firstTitle: "No account ?", secondTitle: "Create one",
                                    padding: EdgeInsets.only(top: 35, left: 20, bottom: 35)
                                )
                            ],
                        )
                    )
                ),
            ),
        );
    }
}
