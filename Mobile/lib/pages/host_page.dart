import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_button.dart';
import 'package:second_app/myWidgets/my_text_fields.dart';
import 'package:second_app/myWidgets/my_title.dart';
import 'package:second_app/utils/post_request.dart';

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
                    padding: const EdgeInsets.only(left: 20.0),
                    child: IconButton(
                        icon: const Icon(Icons.arrow_back),
                        onPressed: () {
                            showCustomSnackBar(context, "No changes.");
                            context.go("/login");
                        },
                    ),
                ),
            ),
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
                                    color: Theme.of(context).scaffoldBackgroundColor,

                                    hintText: "10.0.2.2 (default)",
                                    controller: emailController,
                                    prefixIcon: Icon(Icons.network_wifi),

                                ),
                                SizedBox(height: 7),
                                SizedBox(
                                  height: 30,
                                ),
                                MyButton2(
                                    title: "Save address",
                                    onPressed: (context) {
                                        if (emailController.text.isNotEmpty) {
                                            host = emailController.text;
                                            showCustomSnackBar(context, "Network location has been changed to $host");
                                            context.go('/login');
                                        } else {
                                            showCustomSnackBar(context, "Default network location $host");
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