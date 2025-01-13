import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_text_fields.dart';
import 'package:second_app/myWidgets/my_title.dart';

class HostPage extends StatefulWidget {
    const HostPage({super.key});

    @override
    State<HostPage> createState() => _HostPageState();
}

class _HostPageState extends State<HostPage> {

    final emailController = TextEditingController();
    final scrollController = ScrollController();

    @override
    Widget build(BuildContext context) {
        return Scaffold(
            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            appBar: AppBar(
                shadowColor: Theme.of(context).scaffoldBackgroundColor,
                foregroundColor: Theme.of(context).scaffoldBackgroundColor,
                backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                surfaceTintColor: Theme.of(context).scaffoldBackgroundColor,
                elevation: 0,
                leading: Padding(
                    padding: const EdgeInsets.only(left: 35.0),
                    child: IconButton(
                        icon: const Icon(Icons.arrow_back),
                        onPressed: () {
                            ScaffoldMessenger.of(context).showSnackBar(
                                const SnackBar(content: Text('No changes.', style: TextStyle(fontFamily: "avenir"))),
                            );
                            context.go("/login");
                        },
                    ),
                ),
            ),
            body: Padding(
                padding: EdgeInsets.only(left: 8, right: 14),
                child: RawScrollbar(
                    radius: Radius.circular(10),
                    thumbColor: Colors.black,
                    thickness: 5,
                    controller: scrollController,
                    thumbVisibility: true,
                    child: SingleChildScrollView(
                        controller: scrollController,
                        physics: const AlwaysScrollableScrollPhysics(),
                        child: Column(
                            children: [
                                const MyTitle2(
                                    title: "AREA",
                                    fontSize: 45,
                                    padding: EdgeInsets.only(top: 30),

                                ),
                                const MyTitle2(
                                    title: "Network location",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),

                                ),
                                MyTextField2(
                                    hintText: "10.0.2.2 (default)",
                                    controller: emailController,
                                    prefixIcon: Icon(Icons.email),

                                ),
                                SizedBox(height: 7),
                                SizedBox(
                                  height: 30,
                                ),
                                MyButton2(
                                    title: "Save address",
                                    onPressed: (context) {
                                        if (emailController.text.isNotEmpty) {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                                const SnackBar(content: Text("Netword location has been saved.", style: TextStyle(fontFamily: "avenir"))),
                                            );
                                            context.go('/login');
                                        } else {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                                const SnackBar(content: Text('No changes.', style: TextStyle(fontFamily: "avenir"))),
                                            );
                                            context.go('/login');
                                        }
                                    },
                                ),
                            ],
                        )
                    )
                ),
            ),
        );
    }
}