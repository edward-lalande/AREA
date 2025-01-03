import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_text_fields.dart';
import 'package:second_app/myWidgets/my_title.dart';

class HostPage extends StatelessWidget {
  HostPage({super.key});

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
              title: "Change network location",
              fontSize: 26,
              padding: EdgeInsets.only(top: 30, bottom: 50),
              color: Colors.black
            ),
          MyTextField(
            controller: emailController,
            obscureText: false,
            hintText: "10.0.2.2  (default)",
          
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
            title: "Submit",
            backgroundColor: Colors.black,
            textColor: Colors.white,
            fontSize: 18,
            spaceBetweenIconAndText: 10,
            onPressed: (context) {
              if (emailController.text.isNotEmpty) {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Network location has been changed !', style: TextStyle(fontFamily: "avenir"))),
                );
                context.go('/login');
              } else {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('No changes', style: TextStyle(fontFamily: "avenir"),)),
                );
                context.go('/login');
              }
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
                    "Don't want to change ?",
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
                      "Log page",
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