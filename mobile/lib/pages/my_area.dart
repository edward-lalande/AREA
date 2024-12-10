import 'package:flutter/material.dart';

import '../myWidgets/my_title.dart';

class MyArea extends StatefulWidget {
  const MyArea({super.key});

  @override
  State<MyArea> createState() => _MyAreaState();
}

class _MyAreaState extends State<MyArea> {
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
                    color: const Color.fromARGB(255, 255, 0, 0)
                ),
            ],
        ),
      ),
    );
  }
}