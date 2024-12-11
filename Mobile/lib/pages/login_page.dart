import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../utils/post_request.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';
import '../myWidgets/my_divider_text.dart';
import '../myWidgets/my_web_view.dart';

class LoginPage extends StatefulWidget {

    const LoginPage({super.key});

    @override
    State<LoginPage> createState() => _LoginPageState();

}

class _LoginPageState extends State<LoginPage> {

    final usernameController = TextEditingController();
    final passwordController = TextEditingController();

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
                        bool tmp = await sendSignUp(
                        url: 'http://10.0.2.2:8080/user',
                        body: {
                            "routes": "login",
                            "mail": usernameController.text,
                            "password": passwordController.text
                            }
                        );
                        if (tmp) {
                        if (context.mounted) {
                            context.go("/home");
                        }

                        } else {
                        if (context.mounted) {
                            context.go("/login");
                        }
                        }
                    }
                    ),
                    const MyDividerText(
                    bgColor: Colors.white,
                    padding: EdgeInsets.only(top: 35, right: 35, left: 35),
                    textBetween: "Or",
                    ),
                    MyButton(
                    padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
                    title: "Continue with Discord",
                    backgroundColor: Colors.black,
                    textColor: Colors.white,
                    fontSize: 17,
                    spaceBetweenIconAndText: 10,
                    prefixIcon: Container(
                        width: 30,
                        height: 30,
                        padding: const EdgeInsets.all(5),
                        decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(15),
                        color: Colors.white38,
                        ),
                        child: Image.asset('assets/discord-logo.png'),
                    ),
                    onPressed: (context) async {
                        String url = await getOAuthUrl(url: "http://10.0.2.2:8083/oauth2");
                        if (context.mounted) {
                            Navigator.push(
                                context,
                                MaterialPageRoute(
                                    builder: (context) => WebViewPage(url: url),
                            ),
                            );
                        }
                    },
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
                    const SizedBox(
                    height: 100,
                    )
                ],
                ),
            ),
        )
        );
    }
}
