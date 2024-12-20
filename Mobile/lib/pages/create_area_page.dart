import 'package:flutter/material.dart';

import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_grid_view.dart';
import 'package:second_app/myWidgets/my_title.dart';
import 'package:second_app/utils/post_request.dart';

class CreateArea extends StatefulWidget {
  const CreateArea({super.key});

  @override
  State<CreateArea> createState() => _CreateAreaState();
}

class _CreateAreaState extends State<CreateArea> {
    bool _isGridVisible = false;

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
                                color: Colors.black,
                            ),
                            const MyTitle(
                                title: "Create Area",
                                fontSize: 30,
                                padding: EdgeInsets.only(top: 30, bottom: 50),
                                color: Colors.black,
                            ),
                            MyButton(
                                padding: _isGridVisible ? const EdgeInsets.only(
                                    left: 35,
                                    right: 35,
                                    top: 60,
                                    bottom: 20
                                )
                                : const EdgeInsets.only(
                                    left: 35,
                                    right: 35,
                                    top: 60
                                ),
                                title: "If  this     (add)",
                                backgroundColor: Colors.black,
                                textColor: Colors.white,
                                fontSize: 30,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) {
                                setState(() {
                                    _isGridVisible = !_isGridVisible;
                                });
                                },
                            ),
                            AnimatedSwitcher(
                                duration: const Duration(milliseconds: 1000),
                                transitionBuilder: (child, animation) {
                                    return SizeTransition(
                                        sizeFactor: animation,
                                        axis: Axis.vertical,
                                        child: child,
                                    );
                                },
                                child: _isGridVisible
                                    ? SizedBox(
                                        height: 400,
                                        child: MyGridView(
                                            needAnimation: false,
                                            appbarVisible: false,
                                            map: servicesMap,
                                            typeKey: "services",
                                        ),
                                    )
                                    : const SizedBox.shrink(),
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
            ),
        );
    }
}