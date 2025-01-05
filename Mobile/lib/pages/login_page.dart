import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_web_view.dart';
import 'package:second_app/myWidgets/oauth2_button.dart';

import '../utils/post_request.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';
import '../myWidgets/my_divider_text.dart';

class LoginPage extends StatefulWidget {

    const LoginPage({super.key});

    @override
    State<LoginPage> createState() => _LoginPageState();

}

class _LoginPageState extends State<LoginPage> {

    final usernameController = TextEditingController();
    final passwordController = TextEditingController();
    String tmp = '';

    @override
    Widget build(BuildContext context) {
        return SafeArea(
            child: Scaffold(
                backgroundColor: Colors.white,
                resizeToAvoidBottomInset: false,
                body: SingleChildScrollView(
                    physics: const AlwaysScrollableScrollPhysics(),
                    child: Column(
                        children: [
                            MyButton(title: "",
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
                            })
                            ,
                            const MyTitle(
                                title: "AREA",
                                fontSize: 45,
                                padding: EdgeInsets.only(top: 30),
                                color: Colors.black
                            ),
                            const MyTitle(
                                title: "Log in",
                                fontSize: 30,
                                padding: EdgeInsets.only(top: 30, bottom: 50),
                                color: Colors.black
                            ),
                            MyTextField(
                                controller: usernameController,
                                obscureText: false,
                                hintText: "Email",
                                hintTextColor: Colors.black,
                                bgColor: Colors.white,
                                fieldBgColor: Colors.white,
                                padding: const EdgeInsets.only(top: 0, bottom: 0, left: 35, right: 35),
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
                                padding: const EdgeInsets.only(top: 50, bottom: 0, left: 35, right: 35),
                                inputColor: Colors.black,
                                prefixIcon: const Icon(
                                    Icons.lock,
                                    color: Colors.black,
                                ),
                            ),
                            Container(
                                color: Colors.white,
                                child: Padding(
                                    padding: const EdgeInsets.only(left: 36, top: 6),
                                    child: Row(
                                    children: [
                                        const Text(
                                        "Forget your",
                                        style: TextStyle(fontFamily: "Avenir", fontWeight: FontWeight.w300),
                                        ),
                                        const SizedBox(width: 7),
                                        GestureDetector(
                                        onTap: () {
                                            context.go('/password');
                                        },
                                        child: const Text(
                                            "Password",
                                            style: TextStyle(
                                            fontFamily: 'Avenir',
                                            color: Colors.blue,
                                            decoration: TextDecoration.underline,
                                            decorationColor: Colors.blue,
                                            decorationThickness: 2,
                                            ),
                                        ),
                                        ),
                                        const Text(
                                            "  ?",
                                            style: TextStyle(fontFamily: "Avenir", fontWeight: FontWeight.w900),
                                        ),
                                    ],
                                    ),
                                ),
                            ),
                            MyButton(
                                padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
                                title: "Log in",
                                backgroundColor: Colors.black,
                                textColor: Colors.white,
                                fontSize: 20,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) async {
                                    if (usernameController.text.isEmpty ||
                                        passwordController.text.isEmpty) {
                                        ScaffoldMessenger.of(context).showSnackBar(
                                            SnackBar(
                                                backgroundColor: Colors.grey,
                                                duration: Duration(seconds: 3),
                                                content: Text(
                                                    'Please enter your email and password',
                                                    style: TextStyle(color: Colors.white, fontFamily: "avenir"),
                                                ),
                                            ),
                                        );
                                        return;
                                    }
                                    bool tmp = await sendSignUp(
                                        url: 'http://10.0.2.2:8080/login',
                                        body: {
                                            "mail": usernameController.text,
                                            "password": passwordController.text
                                        }
                                    );
                                    final String servString = await classicGet(
                                        url: "http://10.0.2.2:8080/services",
                                    );
                                    final String actionsString = await classicGet(
                                        url: "http://10.0.2.2:8080/actions",
                                    );
                                    List<dynamic> data = jsonDecode(actionsString);
                                    actionsMap = {
                                        for (var service in data.where((element) => element != null))
                                            service['name']: {
                                                'actions': service['actions'].map((action) {
                                                    return {
                                                        'name': action['name'],
                                                        'arguments': action['arguments'],
                                                    };
                                                }).toList(),
                                            }
                                    };
                                    servicesMap = jsonDecode(servString);
                                    if (tmp) {
                                        if (context.mounted) {
                                            context.go("/home");
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
                                    }
                                }
                            ),
                            const MyDividerText(
                                bgColor: Colors.white,
                                padding: EdgeInsets.only(top: 35, right: 35, left: 35, bottom: 35),
                                textBetween: "Or continue with",
                            ),

                            Row(
                                mainAxisAlignment: MainAxisAlignment.center,
                                children: [
                                    OauthButton(iconPath: "assets/google.png", resize: false,
                                        onPressed: (context) async{
                                            String url = await classicGet(url: "http://10.0.2.2:8080/google/oauth");
                                            if (context.mounted) {
                                                Navigator.push(
                                                    context,
                                                    MaterialPageRoute(
                                                        builder: (context) => WebViewPage(url: url, serv: "google"),
                                                    ),
                                                );
                                            }
                                        },
                                    ),
                                    SizedBox(width: 20,),
                                    OauthButton(iconPath: "assets/discord.png", resize: false,
                                       onPressed: (context) async {
                                           
                                            String url = await classicGet(url: "http://10.0.2.2:8080/discord/oauth");
                                            if (context.mounted) {
                                                Navigator.push(
                                                    context,
                                                    MaterialPageRoute(
                                                        builder: (context) => WebViewPage(url: url, serv: "discord",),
                                                    ),
                                                );
                                            }
                                        },
                                    ),
                                    SizedBox(width: 20),
                                    OauthButton(iconPath: "assets/spotify.png", resize: false,
                                        onPressed: (context) async {
                                            String url = await classicGet(url: "http://10.0.2.2:8080/spotify/oauth");
                                            if (context.mounted) {
                                                Navigator.push(
                                                    context,
                                                    MaterialPageRoute(
                                                        builder: (context) => WebViewPage(url: url, serv: "spotify",),
                                                    ),
                                                );
                                            }
                                        },
                                    ),
                                ],
                            ),
                            SizedBox(height: 20,),
                            Row(
                                mainAxisAlignment: MainAxisAlignment.center,
                                children: [
                                    OauthButton(iconPath: "assets/github.png", resize: true,
                                        resizePadding: EdgeInsets.only(top: 20, bottom: 20, left: 5, right: 5),
                                        onPressed: (context) async {

                                           String url = await classicGet(url: "http://10.0.2.2:8080/github/oauth");
                                            if (context.mounted) {
                                                Navigator.push(
                                                    context,
                                                    MaterialPageRoute(
                                                        builder: (context) => WebViewPage(url: url, serv: "github",),
                                                    ),
                                                );
                                            }
                                        },
                                    ),
                                    SizedBox(width: 20,),
                                    OauthButton(iconPath: "assets/gitlab.png", resize: false,
                                        onPressed: (context) async {
                                            String url = await classicGet(url: "http://10.0.2.2:8080/gitlab/oauth");
                                            if (context.mounted) {
                                                Navigator.push(
                                                    context,
                                                    MaterialPageRoute(
                                                        builder: (context) => WebViewPage(url: url, serv: "gitlab",),
                                                    ),
                                                );
                                            }
                                            },
                                    ),
                                    SizedBox(width: 20),
                                    OauthButton(iconPath: "assets/dropbox.png", resize: false,
                                        onPressed: (context) async {
                                               String url = await classicGet(url: "http://10.0.2.2:8080/dropbox/oauth");
                                            if (context.mounted) {
                                                Navigator.push(
                                                    context,
                                                    MaterialPageRoute(
                                                        builder: (context) => WebViewPage(url: url, serv: "dropbox",),
                                                    ),
                                                );
                                            }
                                        },
                                    ),
                                ],
                            ),
                            Container(
                                height: 130,
                                color: Colors.white,
                                child: Row(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                        const Text(
                                            "No account ?",
                                            style: TextStyle(
                                                fontFamily: "Avenir",
                                                fontSize: 16,
                                            ),
                                        ),
                                        const SizedBox(width: 5),
                                        GestureDetector(
                                            onTap: () {
                                                context.go('/signup');
                                            },
                                            child: const Text(
                                                "Sign-up ",
                                                style: TextStyle(
                                                    decoration: TextDecoration.underline,
                                                    color: Colors.blue,
                                                    fontFamily: "Avenir",
                                                    fontSize: 15,
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
            )
        );
    }
}
