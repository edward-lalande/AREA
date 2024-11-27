import 'package:flutter/material.dart';
import 'package:flutter_application_1/myWidgets/my_button.dart';
import '../myWidgets/my_app_bar_area.dart';
import '../myWidgets/my_email_text_field.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const MyAppBarArea(
        appbartitle: Padding(
          padding: EdgeInsets.only(top: 45),
          child: Text(
                  "Area",
                  style: TextStyle(
                    fontFamily: "Avenir",
                    fontSize: 65,

                  ),),
        ),
      ),
      body: Column(
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
          Container(
            color: Colors.white,
            child: 
            Container(
              color: Colors.white,
              child: const MyTextField(hintText: "Email",)
            ),
          
          ),
          Container(
            color: Colors.white,
            child: const MyTextField(hintText: "Password",)
          ),
          const Padding(padding: EdgeInsets.only(top: 30),
          child: SizedBox(height: 30,)),
          const SizedBox(
            height: 50,
            width: 375,
            child: MyButton(title: "Log in",)
          ),
        ],
      ),
    );
  } 
}
