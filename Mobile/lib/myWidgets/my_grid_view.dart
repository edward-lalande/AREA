import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_card.dart';
import 'package:second_app/myWidgets/my_title.dart';
import 'package:second_app/utils/post_request.dart';

class MyGridView extends StatefulWidget {
    const MyGridView({
        super.key,
        required this.map,
        required this.typeKey,
        required this.appbarVisible,
        required this.homeAnimation,
        this.gridClick,
    });

    final Map<String, dynamic> map;
    final String typeKey;
    final bool appbarVisible;
    final bool homeAnimation;
    final Function(int idx)? gridClick;

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
        widget.gridClick!(index);

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
                        delegate: widget.homeAnimation ? SliverChildBuilderDelegate(
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
                                            color: Colors.grey,
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
                                return InkWell(
                                    onTap: () => _onCardTap(index, tmp),
                                    child: Card(
                                            elevation: 7,
                                            color: Color(0XFF5865F2),
                                            child: MyCard(
                                                title: tmp["name"],
                                                padding: const EdgeInsets.all(8),
                                    ),
                                )
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

    class MyGridViewActionsName extends StatefulWidget {

    const MyGridViewActionsName({
        super.key,
        this.gridClick,
    });

    final Function(int idx)? gridClick;

    @override
    State<MyGridViewActionsName> createState() => _MyGridViewActionsNameState();
}

class _MyGridViewActionsNameState extends State<MyGridViewActionsName> {
    int selectedIndex = -1;

    void _onCardTap(int index, dynamic service) {
        setState(() {
        selectedIndex = selectedIndex == index ? -1 : index;
        });
        widget.gridClick!(index);
    }

    @override
    Widget build(BuildContext context) {
        List<String> keysList = actionsMap.keys.toList();

        return SingleChildScrollView(
            padding: EdgeInsets.all(50),
            child: GridView.builder(
                shrinkWrap: true,
                physics: NeverScrollableScrollPhysics(),
                gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 2,
                    crossAxisSpacing: 10,
                    mainAxisSpacing: 10,
                ),
                itemCount: keysList.length,
                itemBuilder: (context, index) {
                    return InkWell(
                        onTap: () {
                            _onCardTap(index, actionsMap[keysList[index]]);
                        },
                        child: Card(
                            color: Colors.grey,
                            elevation: 7,
                            child: MyCard(
                                title: keysList[index],
                                padding: const EdgeInsets.all(8),
                            ),
                        ),
                    );
                },
            ),
        );
    }
}
