import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../myWidgets/area_app_bar.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_divider_text.dart';

class SignUpPage extends StatelessWidget {
  SignUpPage({super.key});

  final firstNameController = TextEditingController();
  final lastNameController = TextEditingController();
  final ageController = TextEditingController();
  final emailController = TextEditingController();
  final passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
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
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            children: [
              Container(
                color: Colors.white,
                height: 100,
                width: MediaQuery.sizeOf(context).width,
                child: const Text(
                  "Sign Up",
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontFamily: "Avenir",
                    fontSize: 35,
                  ),
                ),
              ),
              // Prénom
              MyTextField(
                controller: firstNameController,
                obscureText: false,
                hintText: "First Name",
                hintTextColor: Colors.black,
                bgColor: Colors.white,
                fieldBgColor: Colors.white,
                padding: const EdgeInsets.only(top: 50, bottom: 0, left: 35, right: 35),
                inputColor: Colors.black,
                prefixIcon: const Icon(
                  Icons.person,
                  color: Colors.black,
                ),
              ),
              // Nom
              MyTextField(
                controller: lastNameController,
                obscureText: false,
                hintText: "Last Name",
                hintTextColor: Colors.black,
                bgColor: Colors.white,
                fieldBgColor: Colors.white,
                padding: const EdgeInsets.only(top: 20, bottom: 0, left: 35, right: 35),
                inputColor: Colors.black,
                prefixIcon: const Icon(
                  Icons.person,
                  color: Colors.black,
                ),
              ),
              // Âge
              MyTextField(
                controller: ageController,
                obscureText: false,
                hintText: "Age",
                hintTextColor: Colors.black,
                bgColor: Colors.white,
                fieldBgColor: Colors.white,
                padding: const EdgeInsets.only(top: 20, bottom: 0, left: 35, right: 35),
                inputColor: Colors.black,
                prefixIcon: const Icon(
                  Icons.cake,
                  color: Colors.black,
                ),
              ),
              // Email
              MyTextField(
                controller: emailController,
                obscureText: false,
                hintText: "Email",
                hintTextColor: Colors.black,
                bgColor: Colors.white,
                fieldBgColor: Colors.white,
                padding: const EdgeInsets.only(top: 20, bottom: 0, left: 35, right: 35),
                inputColor: Colors.black,
                prefixIcon: const Icon(
                  Icons.email,
                  color: Colors.black,
                ),
              ),
              // Mot de passe
              MyTextField(
                controller: passwordController,
                obscureText: true,
                hintText: "Password",
                hintTextColor: Colors.black,
                bgColor: Colors.white,
                fieldBgColor: Colors.white,
                padding: const EdgeInsets.only(top: 20, bottom: 0, left: 35, right: 35),
                inputColor: Colors.black,
                prefixIcon: const Icon(
                  Icons.lock,
                  color: Colors.black,
                ),
              ),
              // Bouton Sign Up
              MyButton(
                padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
                title: "Sign Up",
                backgroundColor: Colors.black,
                textColor: Colors.white,
                fontSize: 20,
                spaceBetweenIconAndText: 10,
                onPressed: (context) {
                  if (firstNameController.text.isNotEmpty &&
                      lastNameController.text.isNotEmpty &&
                      ageController.text.isNotEmpty &&
                      emailController.text.isNotEmpty &&
                      passwordController.text.isNotEmpty) {
                   
                    context.go('/home');
                  } else {
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(content: Text('Veuillez entrer toutes les informations nécessaires')),
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
                height: 210,
                color: Colors.white,
                child: Padding(
                  padding: const EdgeInsets.only(bottom: 20),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Text(
                        "Already have an account?",
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
                          "Log in ",
                          style: TextStyle(
                            decoration: TextDecoration.underline,
                            color: Colors.blue,
                            fontFamily: "Avenir",
                            fontSize: 15,
                            decorationColor: Colors.blue,
                            decorationThickness: 2,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
