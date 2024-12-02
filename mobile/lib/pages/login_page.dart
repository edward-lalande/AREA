import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/area_app_bar.dart';
import '../myWidgets/my_divider_text.dart';

class LoginPage extends StatelessWidget {
  LoginPage({super.key});

  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      appBar: MyAppBarArea(
        appbartitle: Padding(
          padding: const EdgeInsets.only(top: 45),
          child: Column(
            children: [
              MyButton(
                padding: const EdgeInsets.only(left: 10),
                title: "",
                prefixIcon: const Icon(Icons.settings, color: Colors.black,),
                backgroundColor: Colors.white,
                textColor: Colors.white,
                fontSize: 0,
                spaceBetweenIconAndText: 0,
                onPressed: (context) {
                    context.go('/home');
                },
              ),
              const Text(
                "AREA",
                style: TextStyle(
                  fontFamily: "Avenir",
                  fontSize: 65,
                ),
              ),
            ],
          ),
          ),
        ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(0),
          child: Column(
        children: [
          Container(
            color: Colors.white,
            height: 100,
            width: MediaQuery.sizeOf(context).width,
            child: const Text(
              "Log in",
              textAlign: TextAlign.center,
              style: TextStyle(
                fontFamily: "Avenir",
                fontSize: 35,
              ),
            ),
          ),
          MyTextField(
            controller: usernameController,
            obscureText: false,
            hintText: "Email",
            hintTextColor: Colors.black,
            bgColor: Colors.white,
            fieldBgColor: Colors.white,
            padding: const EdgeInsets.only(top: 50, bottom: 0, left: 35, right: 35),
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
            onPressed: (context) {
              if (usernameController.text.isNotEmpty
              && passwordController.text.isNotEmpty) {
                context.go('/home');
              } else {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Veuillez entrer vos informations de connexion')),
                );
              }
            },
          ),
          const MyDividerText(
            bgColor: Colors.white,
            padding: EdgeInsets.only(top: 35, right: 35, left: 35),
            textBetween: "Or",
          ),
          MyButton(
            padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
            title: "Continue with Google",
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
              child: Image.asset('assets/google.png'),
            ),
            onPressed: (context) {
              // OATH2
              context.go('/home');
            },
          ),
          Container(
            height: 250,
            color: Colors.white,
            child: Padding(
              padding: const EdgeInsets.only(),
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
          ),
        ],
      ),
        )
      )
    );
  }
}
