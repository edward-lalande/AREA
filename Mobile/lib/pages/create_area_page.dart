import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_card.dart';
import 'package:second_app/myWidgets/my_dialog.dart';
import 'package:second_app/myWidgets/my_grid_view.dart';
import 'package:second_app/myWidgets/my_title.dart';
import 'package:second_app/utils/post_request.dart';
import 'package:http/http.dart' as http;


    /*[{
        "user_token":"dsfsf",
        "action":{
            "action_id": 0,
            "action_type": "0",
            "hour":"17",
            "minute":"45",
            "city":"Paris",
            "continent":"Europe"},
        "reactions":[{
            "reaction_id":"0",
            "reaction_type":"0",
            "channel_id":"452345432",
            "message":"Hello"
        }]
    }]
 */


class CreateArea extends StatefulWidget {
    const CreateArea({super.key});

    @override
    State<CreateArea> createState() => _CreateAreaState();
}

class _CreateAreaState extends State<CreateArea> {
    bool _isServicesVisible = false;
    bool _selectedService = false;
    bool _actionChosen = false;
    bool _reactionChosen = false;

    bool _isReactionsVisible = false;
    bool _selectedReactions = false;
    Map<String, dynamic> actionData = {};
    Map<String, dynamic> reactionData = {};

    int which = 0;
    final _scrollController = ScrollController();

    void handleServiceChose(int index) {
        setState(() {
            _isServicesVisible = false;
            _selectedService = true;
            which = index;
        });
    }
    void handleActionChose(Map<String, dynamic> formData) {
        setState(() {
            _selectedService = false;
            _isServicesVisible = false;
            _actionChosen = true;
            actionData = formData;
        });
        //print(formData);
    }
    void handleReactionChose(Map<String, dynamic> formData) {
        setState(() {
            _isReactionsVisible = false;
            _selectedReactions = false;
            _reactionChosen = true;
            reactionData = formData;
        });
        //print(formData);
    }
    void handleServiceReactionsChose(int index) {
        setState(() {
            _isReactionsVisible = false;
            _selectedReactions = true;
            which = index;
        });
    }

    Widget showServicesActionsGrid(int serviceIndex, Map<String, dynamic> map, int type) {
        List<Map<String, String>> actionsList = [];

        String selectedServiceName = map.keys.elementAt(serviceIndex);

        if (type == 1) {

            map.forEach((serviceName, serviceData) {
                if (serviceName == selectedServiceName) {
                    for (var action in serviceData['actions']) {
                        actionsList.add({'name': action['name']});
                    }
                }
            });
        }
        else {
            map.forEach((serviceName, serviceData) {
                if (serviceName == selectedServiceName) {
                    for (var action in serviceData['reactions']) {
                        actionsList.add({'name': action['name']});
                    }
                }
            });
        }
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

        if (!_actionChosen) {
            if (serviceIndex == 0) {
                switch (index) {
                    case 0:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 3,
                                    fieldLabels: ['Minute', 'City', 'Continent'],
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
                                    isActionChose: false,
                                    actionId: serviceIndex,
                                    actionType: index,
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
        else {
            if (serviceIndex == 0) {
                switch (index) {
                    case 1 && 2:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 2,
                                    fieldLabels: ["Guild_ID", "Channel_Name"],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    case 3:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 1,
                                    fieldLabels: ["Channel_ID"],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    case 4 && 5:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 2,
                                    fieldLabels: ["Channel_ID", "Message_ID"],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    case 6 && 7:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 2,
                                    fieldLabels: ["Guild_ID", "Role_ID"],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    default:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 2,
                                    fieldLabels: ["Channel_ID", "Message"],
                                    onSubmit: handleReactionChose
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
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 2,
                                    fieldLabels: ['Chanel ID', 'Message ID'],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    default:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 2,
                                    fieldLabels: ['Chanel ID', 'Message ID'],
                                    onSubmit: handleReactionChose
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
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 4,
                                    fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    default:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 3,
                                    fieldLabels: ['City', 'Country', 'Continent'],
                                    onSubmit: handleReactionChose
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
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 4,
                                    fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    default:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 3,
                                    fieldLabels: ['City', 'Country', 'Continent'],
                                    onSubmit: handleReactionChose
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
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 4,
                                    fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    default:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 3,
                                    fieldLabels: ['City', 'Country', 'Continent'],
                                    onSubmit: handleReactionChose
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
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 4,
                                    fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    default:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 3,
                                    fieldLabels: ['City', 'Country', 'Continent'],
                                    onSubmit: handleReactionChose
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
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 4,
                                    fieldLabels: ['Hour', 'Minute', 'City', 'Continent'],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                        break;
                    default:
                        showDialog(
                            context: context,
                            builder: (BuildContext context) {
                                return CustomFormDialog(
                                    isActionChose: true,
                                    actionId: serviceIndex,
                                    actionType: index,
                                    numberOfFields: 3,
                                    fieldLabels: ['City', 'Country', 'Continent'],
                                    onSubmit: handleReactionChose
                                );
                            },
                        );
                    break;
                }
            }
        }
    }
    @override
    Widget build(BuildContext context) {
        return SafeArea(
            child: Scaffold(
                backgroundColor: Colors.white,
                resizeToAvoidBottomInset: false,
                body: Padding(
                    padding: EdgeInsets.only(left: 8, right: 14),
                    child: RawScrollbar(
                        radius: Radius.circular(10),
                        thumbColor: Colors.black,
                        thickness: 5,
                        controller: _scrollController,
                        thumbVisibility: true,
                        child: SingleChildScrollView(
                            controller: _scrollController,
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
                                        padding: _isServicesVisible
                                            ? const EdgeInsets.only(left: 35, right: 35, top: 60, bottom: 20)
                                            : const EdgeInsets.only(left: 35, right: 35, top: 60),
                                        title: "If this (add)",
                                        backgroundColor: Colors.black,
                                        textColor: Colors.white,
                                        fontSize: 30,
                                        spaceBetweenIconAndText: 10,
                                        onPressed: (context) {
                                            setState(() {
                                                _isServicesVisible = !_isServicesVisible;
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
                                        child: _isServicesVisible && !_selectedService
                                            ? SizedBox(
                                                height: 400,
                                                child: MyGridViewActionsName(gridClick: handleServiceChose, dataMap: actionsMap,),
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
                                        child: _selectedService
                                            ? SizedBox(
                                                height: 600,
                                                child: showServicesActionsGrid(which, actionsMap, 1),
                                            )
                                            : const SizedBox.shrink(),
                                    ),

                                    //reaction

                                    MyButton(
                                        padding: _isReactionsVisible
                                            ? const EdgeInsets.only(left: 35, right: 35, top: 60, bottom: 20)
                                            : const EdgeInsets.only(left: 35, right: 35, top: 60),
                                        title: "Then that (add)",
                                        backgroundColor: Colors.black,
                                        textColor: Colors.white,
                                        fontSize: 30,
                                        spaceBetweenIconAndText: 10,
                                        onPressed: (context) {
                                            setState(() {
                                                _isReactionsVisible = !_isReactionsVisible;
                                            });
                                        },
                                    ),
                                    AnimatedSwitcher(
                                        duration: const Duration(milliseconds: 100),
                                        transitionBuilder: (child, animation) {
                                            return SizeTransition(
                                                sizeFactor: animation,
                                                axis: Axis.vertical,
                                                child: child,
                                            );
                                        },
                                        child: _isReactionsVisible && !_selectedReactions
                                            ? SizedBox(
                                                height: 400,
                                                child: MyGridViewActionsName(gridClick: handleServiceReactionsChose, dataMap: reactionsMap,),
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
                                        child: _selectedReactions
                                            ? SizedBox(
                                                height: 400,
                                                child: showServicesActionsGrid(which, reactionsMap, 2),
                                            )
                                            : const SizedBox.shrink(),
                                    ),
                                    _reactionChosen && _actionChosen ? MyButton(
                                        title: "Create Area !",
                                        backgroundColor: Colors.black,
                                        textColor: Colors.white,
                                        padding: EdgeInsets.only(left: 35, right: 35, top: 60),
                                        fontSize: 30,
                                        spaceBetweenIconAndText: 0,
                                        onPressed: (context) async {
                                            final token = await stockData.read('token');
                                            //final tmp = await postArea(
                                            //    url: ""
                                            //);
                                            ScaffoldMessenger.of(context).showSnackBar(
                                                SnackBar(
                                                    backgroundColor: Colors.lightGreen,
                                                    duration: Duration(seconds: 3),
                                                    content: Text(
                                                        'Area Created !',
                                                        style: TextStyle(color: Colors.white, fontFamily: "avenir"),
                                                    ),
                                                ));
                                        }
                                    ): const SizedBox.shrink(),
                                ],
                            ),
                        ),
                    ),
                )
            ),
        );
    }
}
