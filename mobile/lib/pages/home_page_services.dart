import 'package:flutter/material.dart';

import '../myWidgets/my_title.dart';

class HomePageServices extends StatefulWidget {
  const HomePageServices({super.key});

  @override
  State<HomePageServices> createState() => _HomePageServicesState();
}

class _HomePageServicesState extends State<HomePageServices> {
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
                    color: Colors.purple
                ),
            ],
        ),
      ),
    );
  }
}