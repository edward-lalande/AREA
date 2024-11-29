import 'package:flutter/material.dart';
import '../myWidgets/area_app_bar.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_divider_text.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const MyAppBarArea(
        appbartitle: Padding(
            padding: EdgeInsets.only(top: 45),
            child: Text(
              "AREA",
              style: TextStyle(
                fontFamily: "Avenir",
                fontSize: 65,
                ),
              ),
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

          const MyTextField(
            hintText: "Email",
            hintTextColor: Colors.white,
            bgColor: Colors.white,
            fieldBgColor: Colors.black,
            padding: EdgeInsets.only(
              top: 50, bottom: 0, left: 35, right: 35
            ),
            inputColor: Colors.white,
            prefixIcon: Icon(
                Icons.email,
                color: Colors.white,
              ),
          ),
          const MyTextField(
            hintText: "Password",
            hintTextColor: Colors.white,
            bgColor: Colors.white,
            fieldBgColor: Colors.black,
            padding: EdgeInsets.only(
              top: 40, bottom: 0, left: 35, right: 35
            ),
            inputColor: Colors.white,
            prefixIcon: Icon(
                Icons.lock,
                color: Colors.white,
              ),
          ),
          const MyButton(
              padding: EdgeInsets.only(
                left: 35, right: 35, top: 35
              ),
              title: "Log in",
              backgroundColor: Colors.black,
              textColor: Colors.white,
              fontSize: 20,
              spaceBetweenIconAndText: 10,
          ),
          const MyDividerText(
            bgColor: Colors.white,
            padding: EdgeInsets.only(top: 35, right: 35, left: 35),
            textBetween: "Or"
          ),
          MyButton(
              padding: const EdgeInsets.only(
                left: 35, right: 35, top: 35
              ),
              title: "Continue with Google",
              backgroundColor: Colors.black,
              textColor: Colors.white,
              fontSize: 15,
              spaceBetweenIconAndText: 10,
              prefixIcon: Container(
                        width: 35,
                        height: 35,
                        padding: const EdgeInsets.all(5),
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(15),
                          color: Colors.white38,
                        ),
                        child: Image.asset('assets/google.png'),
                      ),
          ),
        ],
      ),
    );
  }
}
