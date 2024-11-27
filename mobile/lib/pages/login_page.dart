import 'package:flutter/material.dart';
import '../myWidgets/area_app_bar.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(

        appBar: const MyAppBarArea(
                        appbartitle:

                            Text(
                              "AREA",
                              style: TextStyle(
                                fontWeight: FontWeight.bold,
                                fontSize: 55.0,
                                fontFamily: 'Avenir'

                              ),
                            ),

                      ),

        body: SingleChildScrollView(

          child: Column(
            children: [

              Container(
                color: Colors.white,
                height: 40.0,
                width: MediaQuery.sizeOf(context).width,

                child: const Text(
                  "Log in",
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    fontSize: 30.0,
                    fontFamily: 'Avenir'

                  ),),
              ),
              Padding(
                  padding: const EdgeInsets.only(top: 100, left: 35, right: 35),
                  child: Container(
                    decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(12),
                        color: Colors.black,
                    ),
                    child: const TextField(
                        style: TextStyle(color: Colors.white),
                        decoration: InputDecoration(
                            border: InputBorder.none,
                                    prefixIcon: Icon(
                                        Icons.email,
                                        color: Colors.white,
                                    ),
                            hintText: 'Email',
                            hintStyle: TextStyle(color: Colors.white),
                        ),
                    ),
                  ),
              ),
            ],
          ),
        ),
      );
  }
}
