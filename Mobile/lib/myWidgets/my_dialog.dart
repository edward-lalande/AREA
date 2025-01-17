import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_text_fields.dart';
import 'package:second_app/utils/post_request.dart';


class ActionDialog extends StatelessWidget {

    final ActionServ action;
    final VoidCallback done;

    const ActionDialog({
        super.key,
        required this.action,
        required this.done,
    });

    @override
    Widget build(BuildContext context) {

    final controllers = action.arguments.map((arg) => TextEditingController()).toList();

    return AlertDialog(

        backgroundColor: Theme.of(context).primaryColorLight,

        title: Text(action.name),
        content: SingleChildScrollView(
            child: Column(
                children: [
                    Text(action.description, style: TextStyle(fontSize: 16),),
                    SizedBox(height: 25),
                    ...List.generate(action.arguments.length, (index) {
                        final arg = action.arguments[index];
                        return Padding(
                            padding: const EdgeInsets.only(bottom: 20.0),
                            child: MyTextField2(
                                padding: EdgeInsets.only(left: 0, right: 0, top: 8, bottom: 8),
                                color: Theme.of(context).primaryColorLight,
                                hintText: arg.display,
                                controller: controllers[index],
                            ),
                        );
                    }
                ),
                ]
            ),
        ),
        actions: [
            MyButton2(
                padding: EdgeInsets.only(left: 16, right: 16, bottom: 20),
                title: "Save options",
                onPressed: (_) {
                    actionDone = true;
                    actionData = {
                        "action_name": action.name,
                        "action_id": action.actionId,
                        "action_type": action.actionType,
                        ...getFormsData(controllers, action.arguments),
                    };
                    Navigator.of(context).pop();
                    Navigator.of(context).pop();
                    done();
                },
            ),
        ],
        );
    }
}

class ReactionDialog extends StatelessWidget {

    final Reaction reaction;
    final VoidCallback done;

    const ReactionDialog({
        super.key,
        required this.reaction,
        required this.done,
    });

    @override
    Widget build(BuildContext context) {

        final controllers = reaction.arguments.map((arg) => TextEditingController()).toList();

        return AlertDialog(
            backgroundColor: Theme.of(context).primaryColorLight,
            title: Text(reaction.name),
            content: SingleChildScrollView(
                child: Column(
                    children: [
                        Text(reaction.description, style: TextStyle(fontSize: 16),),
                        SizedBox(height: 25,),
                        ...List.generate(reaction.arguments.length, (index) {
                            final arg = reaction.arguments[index];
                            return Padding(
                                padding: const EdgeInsets.only(bottom: 20.0),
                                child: MyTextField2(
                                    padding: EdgeInsets.only(left: 0, right: 0, top: 8, bottom: 8),
                                    color: Theme.of(context).primaryColorLight,
                                    hintText: arg.display,
                                    controller: controllers[index],
                                ),
                            );
                        }),
                    ]
                    ),
                ),
                actions: [
                    MyButton2(
                        padding: EdgeInsets.only(left: 16, right: 16, bottom: 20),
                        title: "Save options",
                        onPressed: (_) {
                            reactionDone = true;
                            reactionData = {
                                "reaction_name": reaction.name,
                                "reaction_id": reaction.reactionId,
                                "reaction_type": reaction.reactionType,
                                ...getFormsData(controllers, reaction.arguments),
                            };
                            Navigator.of(context).pop();
                            Navigator.of(context).pop();
                            done();
                        }
                    ),
            ],
        );
    }
}

void myOauthDialog(BuildContext context, String serviceName)
{
    showDialog(
        context: context,
        builder: (context) {
            return AlertDialog(
                title: Text("Authentification required"),
                content: Text("You're not logged in with $serviceName. You have to.",
                    style: TextStyle(fontSize: 16),),
                actions: [
                    MyButton2(
                        title: "Log in",
                        onPressed: (context) {
                            context.go("/login");
                        },
                    ),
                ],
            );
        },
    );
}

Map<String, dynamic> getFormsData(List<TextEditingController> controllers, List<Argument> arguments)
{
    final Map<String, dynamic> result = {};

    for (int i = 0; i < arguments.length; i++) {
        final arg = arguments[i];
        final value = controllers[i].text;

        if (arg.type == "number") {
            result[arg.name] = int.tryParse(value) ?? 0;
        } else if (arg.type == "string") {
            result[arg.name] = value;
        } else {
            result[arg.name] = value;
        }
    }
    return result;
}
