import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
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
        final passwordController = TextEditingController();

    @override
    Widget build(BuildContext context) {
      final theme = Theme.of(context);
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
                                    title: "Reset password",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),

                                ),
                                MyTextField2(
                                    hintText: "New password",
                                    controller: emailController,
                                    prefixIcon: Icon(Icons.email),

                                ),
                                SizedBox(height: 7),
                                Container(
                                    color: Theme.of(context).scaffoldBackgroundColor,
                                    child: Padding(
                                        padding: const EdgeInsets.only(left: 20, top: 6, bottom: 20),
                                        child: Row(
                                            children: [
                                                const Text(
                                                    "Remember you're",
                                                    style: TextStyle(fontFamily: "Avenir", fontWeight: FontWeight.w300),
                                                ),
                                                const SizedBox(width: 7),
                                                InkWell(
                                                    onTap: () {
                                                        context.go('/login');
                                                    },
                                                    child: const Text(
                                                        "Password",
                                                        style: TextStyle(
                                                            fontFamily: 'Avenir',
                                                            color: Colors.blue,
                                                            decoration: TextDecoration.underline,
                                                            decorationColor: Colors.blue,
                                                            decorationThickness: 2,
                                                        ),
                                                    ),
                                                ),
                                                const Text(
                                                    "  ?",
                                                    style: TextStyle(fontFamily: "Avenir", fontWeight: FontWeight.w900),
                                                ),
                                            ],
                                        ),
                                    ),
                                ),
                                SizedBox(
                                  height: 30,
                                ),
                                MyButton2(
                                    title: "Save password",
                                    onPressed: (context) {
                                        if (emailController.text.isNotEmpty) {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                                const SnackBar(content: Text("You're password has been saved.", style: TextStyle(fontFamily: "avenir"))),
                                            );
                                            context.go('/login');
                                        } else {
                                            ScaffoldMessenger.of(context).showSnackBar(
                                                const SnackBar(content: Text('Please enter a new password.', style: TextStyle(fontFamily: "avenir"))),
                                            );
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

/*class PasswordPage extends StatefulWidget {
  const PasswordPage({super.key});

  @override
  State<PasswordPage> createState() => _PasswordPageState();
}

class _PasswordPageState extends State<PasswordPage> {
    final emailController = TextEditingController();

    @override
    Widget build(BuildContext context) {
      final theme = Theme.of(context);
            return Scaffold(
                backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                resizeToAvoidBottomInset: false,
                appBar: AppBar(
                backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                elevation: 0,
                leading: Padding(
                    padding: const EdgeInsets.only(left: 35.0),
                    child: IconButton(
                        icon: const Icon(Icons.arrow_back),
                        onPressed: () {
                            context.go("/login");
                        },
                    ),
                ),
                
            ),
                body: SingleChildScrollView(
                    physics: const AlwaysScrollableScrollPhysics(),
                    child: Column(
                        children: [
                           
                             const MyTitle2(
                                    title: "AREA",
                                    fontSize: 45,
                                    padding: EdgeInsets.only(top: 30),

                                ),
                                const MyTitle2(
                                    title: "Peset Password",
                                    fontSize: 30,
                                    padding: EdgeInsets.only(top: 30, bottom: 50),

                                ),
                            MyTextField(
                                controller: emailController,
                                obscureText: false,
                                hintText: "Email",
                                hintTextColor: Colors.black,
                                bgColor: Colors.white,
                                fieldBgColor: Colors.white,
                                padding: const EdgeInsets.only(top: 50, bottom: 0, left: 35, right: 35),
                                inputColor: Colors.black,
                                prefixIcon: const Icon(
                                    Icons.email,
                                    color: Colors.black,
                                ),
                            ),
                            MyButton(
                                padding: const EdgeInsets.only(left: 35, right: 35, top: 35),
                                title: "Send Reset Link",
                                backgroundColor: Colors.black,
                                textColor: Colors.white,
                                fontSize: 17,
                                spaceBetweenIconAndText: 10,
                                onPressed: (context) {
                                    if (emailController.text.isNotEmpty) {
                                        ScaffoldMessenger.of(context).showSnackBar(
                                        const SnackBar(content: Text('A reset link has been sent to your email.', style: TextStyle(fontFamily: "avenir"))),
                                        );
                                        context.go('/login');
                                    } else {
                                        ScaffoldMessenger.of(context).showSnackBar(
                                        const SnackBar(content: Text('Please enter your email address', style: TextStyle(fontFamily: "avenir"))),
                                        );
                                    }
                                },
                            ),
                            const MyDividerText(
                                bgColor: Colors.white,
                                padding: EdgeInsets.only(top: 35, right: 35, left: 35),
                                textBetween: "Or",
                            ),
                            Container(
                                height: 130,
                                color: Colors.white,
                                child: Padding(
                                padding: const EdgeInsets.only(),
                                child: Row(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                    const Text(
                                        "Remembered your password?",
                                        style: TextStyle(
                                        fontFamily: "Avenir",
                                        fontSize: 16,
                                        ),
                                    ),
                                    const SizedBox(width: 5),
                                    GestureDetector(
                                        onTap: () {
                                        context.go('/login');
                                        },
                                        child: const Text(
                                        "Log in",
                                        style: TextStyle(
                                            decoration: TextDecoration.underline,
                                            color: Colors.blue,
                                            fontFamily: "Avenir",
                                            fontSize: 16,
                                            decorationColor: Colors.blue,
                                            decorationThickness: 2,
                                        ),
                                        ),
                                    ),
                                    ],
                                ),
                                ),
                            ),
                            Container(
                                height: 150,
                                color: Colors.white,
                            )
                        ],
                    ),
                )
            );
    }
}*/
