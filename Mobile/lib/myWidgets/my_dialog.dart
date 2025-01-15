import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_text_fields.dart';
import 'package:second_app/utils/post_request.dart';


class ActionDialog extends StatelessWidget {
    final ActionServ action;

    const ActionDialog({
        super.key,
        required this.action,
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
                    Text(action.description, style: TextStyle(fontSize: 15),),
                    SizedBox(height: 20),
                    ...List.generate(action.arguments.length, (index) {
                        final arg = action.arguments[index];
                        return Padding(
                            padding: const EdgeInsets.only(bottom: 16.0),
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
            TextButton(
                onPressed: () {
                    for (int i = 0; i < controllers.length; i++) {
                        print('${action.arguments[i].name}: ${controllers[i].text}');
                    }
                    Navigator.of(context).pop();
                },
                child: Center(child: Text('Save options', textAlign: TextAlign.center, style: TextStyle(color: Theme.of(context).textTheme.bodyLarge?.color),),),)
            ],
        );
    }
}

class ReactionDialog extends StatelessWidget {
    final Reaction reaction;

    const ReactionDialog({
        super.key,
        required this.reaction,
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
                        Text(reaction.description, style: TextStyle(fontSize: 15),),
                        SizedBox(height: 20,),
                        ...List.generate(reaction.arguments.length, (index) {
                            final arg = reaction.arguments[index];
                            return Padding(
                                padding: const EdgeInsets.only(bottom: 16.0),
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
                    TextButton(
                        onPressed: () {
                            for (int i = 0; i < controllers.length; i++) {
                                print('${reaction.arguments[i].name}: ${controllers[i].text}');
                            }
                            Navigator.of(context).pop();
                        },
                        child: Center(child: Text("Save options", textAlign: TextAlign.center, style: TextStyle(color: Theme.of(context).textTheme.bodyLarge?.color),),),
                ),
            ],
        );
    }
}