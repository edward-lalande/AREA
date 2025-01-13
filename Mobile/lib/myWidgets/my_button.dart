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

class MyButton2 extends StatelessWidget {
    const MyButton2({
        super.key,
        required this.title,
        required this.onPressed,
        this.prefixIcon,
        this.padding = const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
        this.fontSize = 25,
        this.spaceBetweenIconAndText = 8,
    });

    final String title;
    final void Function(BuildContext) onPressed;
    final Widget? prefixIcon;
    final EdgeInsets padding;
    final double fontSize;
    final double spaceBetweenIconAndText;

    @override
    Widget build(BuildContext context) {
        final theme = Theme.of(context);

        return Container(
            padding: padding,
            width: double.infinity,
            child: TextButton(
                onPressed: () => onPressed(context),
                style: ButtonStyle(
                    backgroundColor: WidgetStatePropertyAll(theme.primaryColor),
                    foregroundColor: WidgetStatePropertyAll(Colors.white),
                    overlayColor: WidgetStateProperty.resolveWith(
                        (states) {
                            if (states.contains(WidgetState.pressed) || states.contains(WidgetState.hovered)) {
                                return theme.primaryColorLight.withOpacity(0.2);
                            }
                            return null;
                        },
                    ),
                    shape: WidgetStatePropertyAll(
                        RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(25),
                            side: BorderSide(
                            color: theme.dividerColor,
                            width: 2,
                            ),
                        ),
                    ),
                ),
                child: Row(
                    mainAxisSize: MainAxisSize.min,
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                        if (prefixIcon != null) ...[
                            prefixIcon!,
                            SizedBox(width: spaceBetweenIconAndText),
                        ],
                        Text(
                            title,
                            style: TextStyle(
                                fontFamily: "Avenir",
                                fontSize: fontSize,
                                fontWeight: FontWeight.bold,
                            ),
                        ),
                    ],
                ),
            ),
        );
    }
}

