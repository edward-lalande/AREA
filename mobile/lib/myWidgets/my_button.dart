import 'package:flutter/material.dart';

class MyButton extends StatelessWidget {
  const MyButton({super.key, required this.title, required this.backgroundColor,
  required this.textColor, required this.padding});
  final String title;

  final Color backgroundColor;
  final Color textColor;
  final EdgeInsets padding;
  @override
  Widget build(BuildContext context) {
    return Container(

      color: Colors.white,
      width: double.infinity,

      child: Padding(

        padding: padding,
        child: TextButton(
            style: ButtonStyle(
              backgroundColor: WidgetStateProperty.all<Color>(backgroundColor),
              foregroundColor: WidgetStateProperty.all<Color>(textColor),
              overlayColor: WidgetStateProperty.resolveWith<Color?>(
                (Set<WidgetState> states) {
                  if (states.contains(WidgetState.hovered)) {
                    return const Color.fromARGB(255, 255, 255, 255).withOpacity(0.24);
                  }
                  if (states.contains(WidgetState.focused) ||
                      states.contains(WidgetState.pressed)) {
                    return const Color.fromARGB(255, 255, 255, 255).withOpacity(0.24);
                  }
                  return null;
                },
              ),
            ),
            onPressed: () { },
            child: Text(
              title,
              style: const TextStyle(
                fontFamily: "Avenir",
                fontSize: 20),),
        ),
      ),
    );
  }
}
