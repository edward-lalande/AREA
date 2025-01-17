import 'package:flutter/material.dart';

import 'package:second_app/myWidgets/my_title.dart';

class MyArea extends StatefulWidget {
    const MyArea({super.key});

    @override
    State<MyArea> createState() => _MyAreaState();
}

class _MyAreaState extends State<MyArea> {
    @override
    Widget build(BuildContext context) {

        final scrollController = ScrollController();

        return Scaffold(
            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            body: Padding(
                padding: EdgeInsets.only(left: 8, right: 14),
                child: RawScrollbar(
                    radius: Radius.circular(10),
                    thumbColor: Theme.of(context).primaryColor,
                    thickness: 5,
                    controller: scrollController,
                    thumbVisibility: true,
                    child: SingleChildScrollView(
                        controller: scrollController,
                        physics: const AlwaysScrollableScrollPhysics(),
                        child: Column(
                            children: [
                                SizedBox(height: 100),
                                const MyTitle2(
                                    title: "AREA",
                                    fontSize: 45,
                                    padding: EdgeInsets.only(top: 30),
                                ),
                                const MyTitle2(
                                    title: "My Area's",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),
                                ),
                            ]
                        )
                    )
                )
            )
        );
    }
}