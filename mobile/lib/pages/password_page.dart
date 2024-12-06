import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';
import '../myWidgets/my_divider_text.dart';

class PasswordPage extends StatelessWidget {
  PasswordPage({super.key});

  final emailController = TextEditingController();

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
          const MyTitle(
              title: "AREA",
              fontSize: 45,
              padding: EdgeInsets.only(top: 80),
              color: Colors.black
            ),
            const MyTitle(
              title: "Reset Password",
              fontSize: 30,
              padding: EdgeInsets.only(top: 30, bottom: 50),
              color: Colors.black
            ),
          MyTextField(
            controller: emailController,
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
          MyButton(
            padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
            title: "Send Reset Link",
            backgroundColor: Colors.black,
            textColor: Colors.white,
            fontSize: 17,
            spaceBetweenIconAndText: 10,
            onPressed: (context) {
              if (emailController.text.isNotEmpty) {
               // idée on lui fait la reaction
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('A reset link has been sent to your email.')),
                );
                context.go('/login');
              } else {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Please enter your email address')),
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
              context.go('/home');
            },
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
                    "Remembered your password?",
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
                      "Log in",
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
          Container(
            height: 150,
            color: Colors.white,
          )
        ],
      ),
      )
    ),
    );
  }
}
