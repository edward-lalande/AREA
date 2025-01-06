import 'package:flutter/material.dart';

class CustomFormDialog extends StatefulWidget {
    final int numberOfFields;
    final int actionId;
    final int actionType;
    final List<String> fieldLabels;
    final Function(Map<String, String>) onSubmit;
    final bool isActionChose;

    const CustomFormDialog({
        super.key,
        required this.numberOfFields,
        required this.fieldLabels,
        required this.onSubmit,
        required this.actionId,
        required this.actionType,
        required this.isActionChose
    });

    @override
    _CustomFormDialogState createState() => _CustomFormDialogState();
}

class _CustomFormDialogState extends State<CustomFormDialog> {

    final Map<String, String> actionForm = {};
    final Map<String, String> reactionForm = {};
    final List<TextEditingController> controllers = [];

    @override
    void initState() {
        super.initState();
        controllers.addAll(List.generate(widget.numberOfFields, (_) => TextEditingController()));
    }

    @override
    void dispose() {
        for (var controller in controllers) {
            controller.dispose();
        }
        super.dispose();
    }

    @override
    Widget build(BuildContext context) {
        return Dialog(
            shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(20),
            ),
            backgroundColor: Colors.white.withOpacity(0.9),
            child: SingleChildScrollView(
                child: Padding(
                    padding: const EdgeInsets.all(40),
                    child: Column(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                            Text(
                                'Options Requirement',
                                textAlign: TextAlign.center,
                                style: const TextStyle(
                                    fontSize: 17,
                                    fontFamily: "avenir",
                                    fontWeight: FontWeight.bold,
                                    color: Colors.black,
                                ),
                            ),
                            const SizedBox(height: 20),
                            ...List.generate(
                                widget.numberOfFields,
                                (index) => Padding(
                                    padding: const EdgeInsets.only(bottom: 30),
                                    child: TextFormField(
                                        controller: controllers[index],
                                        decoration: InputDecoration(
                                            labelText: widget.fieldLabels[index],
                                            labelStyle: const TextStyle(
                                                fontFamily: "avenir",
                                                fontSize: 15,
                                                color: Colors.black,
                                            ),
                                            hintStyle: const TextStyle(
                                                fontFamily: "avenir",
                                                color: Colors.black,
                                            ),
                                            border: OutlineInputBorder(
                                                borderRadius: BorderRadius.circular(10),
                                                borderSide: const BorderSide(
                                                    color: Colors.black,
                                                    width: 2,
                                                ),
                                            ),
                                            focusedBorder: OutlineInputBorder(
                                                borderRadius: BorderRadius.circular(10),
                                                borderSide: const BorderSide(
                                                    color: Colors.black,
                                                    width: 2,
                                                ),
                                            ),
                                        ),
                                        style: const TextStyle(
                                            fontFamily: "avenir",
                                            color: Colors.black,
                                        ),
                                    ),
                                ),
                            ),
                            const SizedBox(height: 20),
                            ElevatedButton(
                                style: ElevatedButton.styleFrom(
                                    backgroundColor: Colors.black,
                                    shape: RoundedRectangleBorder(
                                        borderRadius: BorderRadius.circular(30),
                                    ),
                                    padding: const EdgeInsets.symmetric(
                                        vertical: 15,
                                        horizontal: 30,
                                    ),
                                ),
                                onPressed: () {
                                    if (widget.isActionChose) {
                                        reactionForm["reaction_id"] = widget.actionId.toString();
                                        reactionForm["reaction_type"] = widget.actionType.toString();
                                        for (int i = 0; i < widget.numberOfFields; i++) {
                                            reactionForm[widget.fieldLabels[i].toLowerCase()] = controllers[i].text;
                                        }
                                        widget.onSubmit(reactionForm);
                                    }
                                    else {
                                        actionForm["action_id"] = widget.actionId.toString();
                                        actionForm["action_type"] = widget.actionType.toString();
                                        for (int i = 0; i < widget.numberOfFields; i++) {
                                            actionForm[widget.fieldLabels[i].toLowerCase()] = controllers[i].text;
                                        }
                                        widget.onSubmit(actionForm);
                                    }
                                    Navigator.of(context).pop();
                                    ScaffoldMessenger.of(context).showSnackBar(
                                                SnackBar(
                                                    backgroundColor: Colors.lightGreen,
                                                    duration: Duration(seconds: 3),
                                                    content: Text(
                                                        'Options Saved !',
                                                        style: TextStyle(color: Colors.white, fontFamily: "avenir"),
                                                    ),
                                                ));
                                },
                                child: const Text(
                                    'Submit',
                                    style: TextStyle(
                                        fontSize: 16,
                                        fontFamily: 'avenir',
                                        color: Colors.white,
                                    ),
                                ),
                            ),
                        ],
                    ),
                ),
            ),
        );
    }
}
