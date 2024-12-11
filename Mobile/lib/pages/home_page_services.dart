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
                    color: Colors.black
                ),
                MyTitle(
                    title: "Services available",
                    fontSize: 35,
                    padding: EdgeInsets.only(top: 30, bottom: 50),
                    color: Colors.black
                ),
                Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                       Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                            Card(
                                elevation: 7,
                                color: Color(0XFF5865F2),
                                child: SampleCard(
                                    title: "Discord",
                                    icon: Icon(
                                        color: Colors.white,
                                        Icons.discord,
                                        size: 50,
                                    ),
                                    padding: const EdgeInsets.only(),
                                )
                            ),
                            const SizedBox(width: 5,),
                            Card(
                                elevation: 7,
                                child: SampleCard(
                                    title: "Time User",
                                    icon: Icon(
                                        Icons.av_timer,
                                        size: 50,
                                    ),
                                    padding: const EdgeInsets.only(),
                                )
                            ),
                        ],
                      ),
                    ],
                )
            ],
        ),
      ),
    );
  }
}

class SampleCard extends StatelessWidget {
    const SampleCard({
        super.key,
        required this.icon,
        required this.title,
        this.padding = const EdgeInsets.all(0),
    });

    final Widget icon;
    final String title;
    final EdgeInsetsGeometry padding;

    @override
    Widget build(BuildContext context) {
        return Padding(
            padding: padding,
            child: SizedBox(
                width: 150,
                height: 100,
                child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                        icon,
                        const SizedBox(
                            height: 8
                        ),
                        Text(
                            title,
                            textAlign: TextAlign.center,
                            style: TextStyle(
                                fontSize: 14,
                                color: Colors.black,
                                fontFamily: "Avenir"
                            ),
                        ),
                    ],
                ),
            ),
        );
    }
}
