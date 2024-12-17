import 'package:flutter/material.dart';

import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_title.dart';
import 'package:second_app/utils/post_request.dart';

class CreateArea extends StatefulWidget {
  const CreateArea({super.key});

  @override
  State<CreateArea> createState() => _CreateAreaState();
}

class _CreateAreaState extends State<CreateArea> {

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
                                title: "Create Area",
                                fontSize: 30,
                                padding: EdgeInsets.only(top: 30, bottom: 50),
                                color: Colors.black
                            ),
                            MyButton(
                                padding: const EdgeInsets.only(left: 35, right: 35, top: 60),
                                title: "If  this     (add)",
                                backgroundColor: Colors.black,
                                textColor: Colors.white,
                                fontSize: 30,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) async {
                                    final String tmp = await classicGet(
                                        url: "http://10.0.2.2:8080/services",
                                    );
                                },
                            ),
                            MyButton(
                                padding: const EdgeInsets.only(left: 35, right: 35, top: 30),
                                title: "Then that  (add)",
                                backgroundColor: Colors.grey,
                                textColor: Colors.white,
                                fontSize: 30,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) {

                                },
                            ),
                        ],
                    ),
                ),
            )
        );
  }
}
