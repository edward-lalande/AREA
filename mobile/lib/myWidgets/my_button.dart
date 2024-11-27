import 'package:flutter/material.dart';

class MyButton extends StatelessWidget {
  const MyButton({super.key, required this.title});
  final String title;
  @override
  Widget build(BuildContext context) {
    return TextButton(
          style: ButtonStyle(
            backgroundColor: WidgetStateProperty.all<Color>(Colors.black),
            foregroundColor: WidgetStateProperty.all<Color>(const Color.fromARGB(255, 255, 255, 255)),
            overlayColor: WidgetStateProperty.resolveWith<Color?>(
              (Set<WidgetState> states) {
                if (states.contains(WidgetState.hovered)) {
                  return const Color.fromARGB(255, 112, 112, 112).withOpacity(0.12);
                }
                if (states.contains(WidgetState.focused) ||
                    states.contains(WidgetState.pressed)) {
                  return const Color.fromARGB(255, 112, 112, 112).withOpacity(0.12);
                }
                return null; // Defer to the widget's default.
              },
            ),
          ),
          onPressed: () { },
          child: Text(
            title,
            style: const TextStyle(
              fontFamily: "Avenir",
              fontSize: 20),),
      );
  }
}
