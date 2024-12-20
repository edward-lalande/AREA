import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_card.dart';
import 'package:second_app/myWidgets/my_title.dart';

class MyGridView extends StatefulWidget {
    const MyGridView({
        super.key,
        required this.map,
        required this.typeKey,
        required this.appbarVisible,
        required this.needAnimation,
    });

    final Map<String, dynamic> map;
    final String typeKey;
    final bool appbarVisible;
    final bool needAnimation;

    @override
    State<MyGridView> createState() => _MyGridViewState();
}

class _MyGridViewState extends State<MyGridView> {
    int selectedIndex = -1;
    bool  isBig = false;

    void _onCardTap(int index, dynamic service) {
        setState(() {
            selectedIndex = selectedIndex == index ? -1 : index;
            isBig = !isBig;
        });

        Future.delayed(Duration(milliseconds: 200), () {
            setState(() {
                isBig = false;
            });
        });
    }

    @override
    Widget build(BuildContext context) {
        final List<dynamic> services = widget.map[widget.typeKey];

        return CustomScrollView(
            slivers: [
                if (widget.appbarVisible)
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
                if (widget.appbarVisible)
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
                        delegate: widget.needAnimation ? SliverChildBuilderDelegate(
                            childCount: services.length,
                            (context, index) {
                                final tmp = services[index];
                                return InkWell(
                                    onTap: () => _onCardTap(index, tmp),
                                    child: AnimatedScale(
                                        scale: selectedIndex == index ? 1.1 : 1.0,
                                        duration: Duration(milliseconds: 200),
                                        child: Card(
                                            elevation: 7,
                                            color: Color(0XFF5865F2),
                                            child: MyCard(
                                                title: tmp["name"],
                                                padding: const EdgeInsets.all(8),
                                                icon: Icon(
                                                    color: Colors.white,
                                                    Icons.discord,
                                                    size: 60,
                                                ),
                                            ),
                                        ),
                                    ),
                                );
                            },
                        ) : SliverChildBuilderDelegate(
                            childCount: services.length,
                            (context, index) {
                                final tmp = services[index];
                                return Card(
                                            elevation: 7,
                                            color: Color(0XFF5865F2),
                                            child: MyCard(
                                                title: tmp["name"],
                                                padding: const EdgeInsets.all(8),
                                                icon: Icon(
                                                    color: Colors.white,
                                                    Icons.discord,
                                                    size: 60,
                                                ),
                                    ),
                                );
                            },
                        ),
                    ),
                ),
                SliverPadding(
                    padding: const EdgeInsets.only(bottom: 50),
                    sliver: SliverToBoxAdapter(
                        child: SizedBox(
                            height: 20,
                        ),
                    ),
                ),
            ],
        );
    }
}

