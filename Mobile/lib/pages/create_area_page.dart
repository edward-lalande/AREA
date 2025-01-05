import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_card.dart';
import 'package:second_app/myWidgets/my_dialog.dart';
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
    bool _actionChosen = false;
    int which = 0;

    void handleServiceChose(int index) {
        setState(() {
            _isGridVisible = false;
            _selected = true;
            which = index;
        });
    }
    void handleActionChose(Map<String, String> formData) {
        setState(() {
            _selected = false;
            _isGridVisible = false;
            _actionChosen = true;
        });
    }

    Widget showServicesActionsGrid(int serviceIndex) {
        List<Map<String, String>> actionsList = [];

        String selectedServiceName = actionsMap.keys.elementAt(serviceIndex);

        actionsMap.forEach((serviceName, serviceData) {
            if (serviceName == selectedServiceName) {
                for (var action in serviceData['actions']) {
                    actionsList.add({'name': action['name']});
                }
            }
        });

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
                itemCount: actionsList.length,
                itemBuilder: (context, index) {
                    return GestureDetector(
                        onTap: () => handleCardTap(context, index, serviceIndex),
                        child: Card(
                            color: Colors.grey,
                            elevation: 7,
                            child: MyCard(
                                title: actionsList[index]['name']!,
                                padding: const EdgeInsets.all(8),
                            ),
                        ),
                    );
                },
            ),
        );
    }

    void handleCardTap(BuildContext context, int index, int serviceIndex) {

        if (serviceIndex == 0) {
            switch (index) {
                case 0:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 4,
                                fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                    break;
                default:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 3,
                                fieldLabels: ['City', 'Country', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                break;
            }
        }
        if (serviceIndex == 1) {
            switch (index) {
                case 0:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 2,
                                fieldLabels: ['Chanel ID', 'Message ID'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                    break;
                default:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 2,
                                fieldLabels: ['Chanel ID', 'Message ID'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                break;
            }
        }
        if (serviceIndex == 2) {
            switch (index) {
                case 0:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 4,
                                fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                    break;
                default:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 3,
                                fieldLabels: ['City', 'Country', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                break;
            }
        }
        if (serviceIndex == 3) {
            switch (index) {
                case 0:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 4,
                                fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                    break;
                default:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 3,
                                fieldLabels: ['City', 'Country', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                break;
            }
        }
        if (serviceIndex == 4) {
            switch (index) {
                case 0:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 4,
                                fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                    break;
                default:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 3,
                                fieldLabels: ['City', 'Country', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                break;
            }
        }
        if (serviceIndex == 5) {
            switch (index) {
                case 0:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 4,
                                fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                    break;
                default:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 3,
                                fieldLabels: ['City', 'Country', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                break;
            }
        }
        if (serviceIndex == 6) {
            switch (index) {
                case 0:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 4,
                                fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                    break;
                default:
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                            return CustomFormDialog(
                                numberOfFields: 3,
                                fieldLabels: ['City', 'Country', 'Continent'],
                                onSubmit: handleActionChose
                            );
                        },
                    );
                break;
            }
        }
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
                                title: _actionChosen ? "if this Discord services": "If this (add)",
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
                                        child: MyGridViewActionsName(gridClick: handleServiceChose),
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
                                        height: 600,
                                        child: showServicesActionsGrid(which),
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
