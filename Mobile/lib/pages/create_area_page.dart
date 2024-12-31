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
    bool _selected = false;

    void _handleItemTap(int index) {
        setState(() {
        _isGridVisible = false;
        _selected = true;

        });
  }

Widget _buildActionsGrid() {

    List<Map<String, String>> actionsList = [];
    actionsMap.forEach((serviceName, serviceData) {
        for (var action in serviceData['actions']) {
            actionsList.add({'name': action['name']});
        }
    });

    return GridView.builder(
        shrinkWrap: true,
        physics: NeverScrollableScrollPhysics(),
        gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
            crossAxisCount: 3,
            crossAxisSpacing: 10,
            mainAxisSpacing: 10,
        ),
        itemCount: actionsList.length,
        itemBuilder: (context, index) {
            return Card(
                color: Colors.blueAccent,
                child: Center(
                    child: Text(
                        actionsList[index]['name']!,
                        style: TextStyle(color: Colors.white, fontSize: 18),
                    ),
                ),
            );
        },
    );
}

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
                    padding: _isGridVisible
                        ? const EdgeInsets.only(left: 35, right: 35, top: 60, bottom: 20)
                        : const EdgeInsets.only(left: 35, right: 35, top: 60),
                    title: "If this (add)",
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
                    duration: const Duration(milliseconds: 500),
                    transitionBuilder: (child, animation) {
                        return SizeTransition(
                            sizeFactor: animation,
                            axis: Axis.vertical,
                            child: child,
                        );
                    },
                    child: _isGridVisible && !_selected
                        ? SizedBox(
                            height: 400,
                            child: MyGridView(
                                homeAnimation: false,
                                appbarVisible: false,
                                map: servicesMap,
                                typeKey: "services",
                                gridClick: _handleItemTap,
                            ),
                            )
                        : const SizedBox.shrink(),
                ),
                AnimatedSwitcher(
                    duration: const Duration(milliseconds: 500),
                    transitionBuilder: (child, animation) {
                        return SizeTransition(
                            sizeFactor: animation,
                            axis: Axis.vertical,
                            child: child,
                        );
                    },
                    child: _selected
                        ? SizedBox(
                            height: 400,
                            child: _buildActionsGrid(),
                            )
                        : const SizedBox.shrink(),
                ),
                MyButton(
                    padding: const EdgeInsets.only(left: 35, right: 35, top: 30),
                    title: "Then that (add)",
                    backgroundColor: Colors.grey,
                    textColor: Colors.white,
                    fontSize: 30,
                    spaceBetweenIconAndText: 10,
                    onPressed: (context) {},
                ),
            ],
            ),
        ),
        ),
    );
    }
}
