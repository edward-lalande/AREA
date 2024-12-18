import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_card.dart';
import 'package:second_app/myWidgets/my_title.dart';

class MyGridView extends StatelessWidget {
    const MyGridView({
        super.key,
    });
    @override
    Widget build(BuildContext context) {
        return CustomScrollView(
                    slivers: [
                        SliverPadding(
                            padding: const EdgeInsets.only(
                                top: 80,
                                left: 20,
                                right: 20
                            ),
                            sliver: SliverToBoxAdapter(
                                child: MyTitle(
                                margin: EdgeInsets.only(bottom: 20),
                                title: "AREA",
                                fontSize: 45,
                                padding: EdgeInsets.zero,
                                color: Colors.black,
                                ),
                            ),
                        ),
                        SliverPadding(
                            padding: const EdgeInsets.only(
                                top: 20,
                                left: 20,
                                right: 20
                            ),
                            sliver: SliverToBoxAdapter(
                                child: MyTitle(
                                margin: EdgeInsets.only(bottom: 20),
                                title: "Available services",
                                fontSize: 30,
                                padding: EdgeInsets.zero,
                                color: Colors.black,
                                ),
                            ),
                        ),
                        SliverPadding(
                            padding: const EdgeInsets.symmetric(horizontal: 50),
                            sliver: SliverGrid(
                                gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                                    crossAxisCount: 2,
                                    crossAxisSpacing: 20,
                                    mainAxisSpacing: 20,
                                ),
                                delegate: SliverChildBuilderDelegate(
                                    childCount: 13,
                                    (context, index) {
                                        return Card(
                                            elevation: 7,
                                            color: Color(0XFF5865F2),
                                            child: MyCard(
                                                title: "Discord",
                                                padding: const EdgeInsets.all(8),
                                                icon: Icon(
                                                    color: Colors.white,
                                                    Icons.discord,
                                                    size: 50,
                                                ),
                                            ),
                                        );
                                    },
                                )
                            ),
                        ),
                        SliverPadding(
                            padding: const EdgeInsets.only(bottom: 50),
                            sliver: SliverToBoxAdapter(
                                child: SizedBox(
                                    height: 20
                                ),
                            ),
                        ),
                    ],
                );
    }
}