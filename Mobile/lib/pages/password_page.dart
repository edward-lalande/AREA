import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:second_app/myWidgets/my_text_button.dart';
import 'package:second_app/utils/post_request.dart';
import '../myWidgets/my_button.dart';
import '../myWidgets/my_text_fields.dart';
import '../myWidgets/my_title.dart';

class PasswordPage extends StatefulWidget {
    const PasswordPage({super.key});

    @override
    State<PasswordPage> createState() => _PasswordPageState();
}

class _PasswordPageState extends State<PasswordPage> {

    final scrollController = ScrollController();
    final emailController = TextEditingController();

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
                                    title: "Reset password",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),

                                ),
                                MyTextField2(
                                    color: Theme.of(context).scaffoldBackgroundColor,

                                    hintText: "New password",
                                    controller: emailController,
                                    prefixIcon: Icon(Icons.lock),

                                ),
                                SizedBox(height: 7),
                                MyTextButton(
                                    onTap: (context) {
                                        context.go("/login");
                                    },
                                    firstTitle: "Remembered your're", secondTitle: "password ?",
                                    padding: EdgeInsets.only(left:20),
                                ),
                                SizedBox(
                                  height: 30,
                                ),
                                MyButton2(
                                    title: "Save password",
                                    onPressed: (context) {
                                        if (emailController.text.isNotEmpty) {
                                            showCustomSnackBar(context, "You're password has been saved.");
                                            context.go('/login');
                                        } else {
                                            showCustomSnackBar(context, "No changes.");
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
