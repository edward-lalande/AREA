import 'dart:convert';

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
                    thumbColor: Theme.of(context).primaryColor,
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
                                    onTap: (context) {
                                        context.go("/password");
                                    }, firstTitle: "Forget your're", secondTitle: "Password",
                                    padding: EdgeInsets.only(left: 20)
                                ),
                                SizedBox(height: 20,),
                                MyButton2(title: "Log in",
                                     onPressed: (context) async {
                                        //if (emailController.text.isEmpty || passwordController.text.isEmpty) {
                                        //    ScaffoldMessenger.of(context).showSnackBar(
                                        //        SnackBar(
                                        //            backgroundColor: Colors.grey,
                                        //            duration: Duration(seconds: 3),
                                        //            content: Text(
                                        //                'Please enter your email and password',
                                        //                style: TextStyle(color: Colors.white, fontFamily: "avenir"),
                                        //            ),
                                        //        ),
                                        //    );
                                        //    return;
                                        //}
                                        //bool tmp = await sendSignUp(
                                        //    url: 'http://$host:8080/login',
                                        //    body: {
                                        //        "mail": emailController.text,
                                        //        "password": passwordController.text
                                        //    }
                                        //);
                                        final String actionsString = await classicGet(
                                            url: "http://10.0.2.2:8080/actions",
                                        );
                                        final String servString = await classicGet(
                                            url: "http://10.0.2.2:8080/services",
                                        );
                                        servicesMap = jsonDecode(servString);
                                        services = parseServices(actionsString);
                                        context.go("/home");
                                        /*if (tmp) {
                                            if (context.mounted) {
                                            }
                                        } else {
                                            if (context.mounted) {
                                                ScaffoldMessenger.of(context).showSnackBar(
                                                    SnackBar(
                                                        backgroundColor: Colors.grey,
                                                        duration: Duration(seconds: 3),
                                                        content: Text(
                                                            'Wrong email or password',
                                                            style: TextStyle(color: Colors.white, fontFamily: "avenir"),
                                                        ),
                                                    ),
                                                );
                                                context.go("/login");
                                            }
                                        }*/
                                    }
                                ),
                                const MyDividerText2(
                                    padding: EdgeInsets.only(top: 35, right: 35, left: 35, bottom: 35),
                                    textBetween: "Or continue with",
                                ),
                                OAuthButtonsRow(host: "10.0.2.2"),
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
