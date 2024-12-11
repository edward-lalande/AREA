import 'package:flutter/material.dart';

import '../myWidgets/my_title.dart';

class AccountPage extends StatefulWidget {
  const AccountPage({super.key});

  @override
  State<AccountPage> createState() => _AccountPageState();
}

class _AccountPageState extends State<AccountPage> {
  @override
  Widget build(BuildContext context) {
    return SafeArea(
        child: Scaffold(
            backgroundColor: Colors.white,
        body: Column(
            children: [
                MyTitle(
                    title: "AREA",
                    fontSize: 45,
                    padding: EdgeInsets.only(top: 80),
                    color: const Color.fromARGB(255, 47, 0, 255)
                ),
            ],
        ),
      ),
    );
  }
}