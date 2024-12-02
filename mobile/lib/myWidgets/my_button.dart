import 'package:flutter/material.dart';

class MyButton extends StatelessWidget {
  const MyButton({
    super.key,
    required this.title,
    required this.backgroundColor,
    required this.textColor,
    required this.padding,
    required this.fontSize,
    required this.spaceBetweenIconAndText,
    required this.onPressed,
    this.prefixIcon,
  });

  final String title;
  final Color backgroundColor;
  final Color textColor;
  final EdgeInsets padding;
  final Widget? prefixIcon;
  final double fontSize;
  final double spaceBetweenIconAndText;
  final void Function(BuildContext)? onPressed;

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      width: double.infinity,
      child: Padding(
        padding: padding,
        child: TextButton(
          onPressed: () {
            if (onPressed != null) {
              onPressed!(context);
            }
          },
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
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              if (prefixIcon != null) ...[
                prefixIcon!,
              ],
              if (prefixIcon == null) ...[
                 SizedBox(
                  width: spaceBetweenIconAndText,
                  ),
              ],
              Center(child: Text(
                title,
                style: TextStyle(
                  fontFamily: "Avenir",
                  fontSize: fontSize,
                ),
              ),
          ),
                SizedBox(
                  width: spaceBetweenIconAndText
                )
            ],
          ),
        ),
      ),
    );
  }
}