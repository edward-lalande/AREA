import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';

class AccountPage extends StatefulWidget {
    const AccountPage({super.key});

    @override
    State<AccountPage> createState() => _AccountPageState();
}

class _AccountPageState extends State<AccountPage> {
    final Map<String, String> userInfo = {
        'email': 'user@example.com',
        'name': 'John Doe',
        'phone': '+1234567890',
    };

    final TextEditingController emailController = TextEditingController();
    final TextEditingController nameController = TextEditingController();
    final TextEditingController phoneController = TextEditingController();

    @override
    void initState() {

        super.initState();

        emailController.text = userInfo['email']!;
        nameController.text = userInfo['name']!;
        phoneController.text = userInfo['phone']!;
    }

    void updateUserInfo() {

        setState(() {
        userInfo['email'] = emailController.text;
        userInfo['name'] = nameController.text;
        userInfo['phone'] = phoneController.text;
        });

        print('User info updated: $userInfo');
    }

    @override
    Widget build(BuildContext context) {
        return SafeArea(
        child: Scaffold(
            backgroundColor: Colors.white,
            body: SingleChildScrollView(
            physics: const AlwaysScrollableScrollPhysics(),
            child: Column(
                children: [
                MyButton(
                    title: "",
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
                    },
                ),
                const MyTitle(
                    title: "Account",
                    fontSize: 30,
                    padding: EdgeInsets.only(top: 30, bottom: 50),
                    color: Colors.black,
                ),
                MyTextField(
                    controller: emailController,
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
                    controller: nameController,
                    obscureText: false,
                    hintText: "Full Name",
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
                MyTextField(
                    controller: phoneController,
                    obscureText: false,
                    hintText: "Phone Number",
                    hintTextColor: Colors.black,
                    bgColor: Colors.white,
                    fieldBgColor: Colors.white,
                    padding: const EdgeInsets.only(top: 50, bottom: 0, left: 35, right: 35),
                    inputColor: Colors.black,
                    prefixIcon: const Icon(
                        Icons.phone,
                        color: Colors.black,
                    ),
                ),
                MyButton(
                    padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
                    title: "Save Changes",
                    backgroundColor: Colors.black,
                    textColor: Colors.white,
                    fontSize: 20,
                    spaceBetweenIconAndText: 10,
                    onPressed: (context) {
                        updateUserInfo();
                    },
                ),
                ],
            ),
            ),
        ),
        );
    }
    }