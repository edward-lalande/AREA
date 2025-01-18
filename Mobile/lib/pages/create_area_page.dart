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

    bool showActionsGrid = false;
    bool showReactionsGrid = false;

    bool startSelect = false;

    @override
    Widget build(BuildContext context) {

        final scrollController = ScrollController();

        return Scaffold(
            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            body: Padding(
                padding: EdgeInsets.only(left: 8, right: 14),
                child: RawScrollbar(
                    radius: Radius.circular(10),
                    thumbColor: Theme.of(context).textTheme.bodyLarge?.color,
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
                                    title: "Create Area",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),
                                ),
                                SizedBox(height: 50),
                                Padding(
                                    padding: EdgeInsets.symmetric(horizontal: 4),
                                    child: Column(
                                        mainAxisAlignment: MainAxisAlignment.center,
                                        children: [
                                            MyButton2(
                                                title: "If this (actions)",
                                                onPressed: (context) {
                                                    setState(() {
                                                        startSelect = true;
                                                        showActionsGrid = !showActionsGrid;
                                                        showReactionsGrid = false;
                                                    });
                                                },
                                            ),
                                        ],
                                    ),
                                ),
                                SizedBox(height: 20),
                                AnimatedSwitcher(
                                    duration: const Duration(milliseconds: 300),
                                    child: startSelect && showActionsGrid && !actionDone
                                        ? ActionsGrid(
                                            onActionSelected: (isActionDone) {
                                                setState(() {
                                                  actionDone = true;
                                                });
                                            },
                                            services: services,
                                            key: ValueKey("servicesGrid")
                                        )
                                        : SizedBox.shrink(),
                                ),
                                 Padding(
                                    padding: EdgeInsets.symmetric(horizontal: 4),
                                    child: Column(
                                        mainAxisAlignment: MainAxisAlignment.center,
                                        children: [
                                            MyButton2(
                                                title: "Then that (reactions)",
                                                onPressed: (context) {
                                                    setState(() {
                                                        showActionsGrid = false;
                                                        showReactionsGrid = !showReactionsGrid;
                                                    });
                                                },
                                            ),
                                        ],
                                    ),
                                ),
                                AnimatedSwitcher(
                                    duration: const Duration(milliseconds: 300),
                                    child: showReactionsGrid && !reactionDone
                                        ? ReactionsGrid(
                                            onReactionSelected: (isReactionDone) {
                                                setState(() {
                                                    reactionDone = true;
                                                });
                                            },
                                            reactionServices: reactions,
                                            key: ValueKey("reactionsGrid")
                                        )
                                        : SizedBox.shrink(),
                                ),
                                SizedBox(height: 20,),
                                reactionDone && actionDone ? Padding(
                                    padding: EdgeInsets.symmetric(horizontal: 20),
                                    child: Column(
                                        mainAxisAlignment: MainAxisAlignment.center,
                                        children: [
                                            MyButton2(
                                                title: "Create Area",
                                                onPressed: (context) async {
                                                    bool res = await setupAreaArgs(
                                                        actionData,
                                                        [reactionData]
                                                    );
                                                    reactionDone = false;
                                                    actionDone = false;
                                                    if (res) {
                                                        showCustomSnackBar(context, "AREA Created !");
                                                    }
                                                    else {
                                                      showCustomSnackBar(context, "Creation of Area failed !");
                                                    }
                                                },
                                            ),
                                        ],
                                    ),
                                ) : SizedBox.shrink(),
                            ],
                        ),
                    ),
                ),
            ),
        );
    }
}