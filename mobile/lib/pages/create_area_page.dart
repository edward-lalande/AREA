import 'package:flutter/material.dart';

import '../myWidgets/my_title.dart';

class CreateArea extends StatelessWidget {
  const CreateArea({super.key});

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
                    color: const Color.fromARGB(255, 0, 0, 0)
                ),
            ],
        ),
      ),
    );
  }
}