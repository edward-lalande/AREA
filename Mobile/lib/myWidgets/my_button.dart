import 'package:flutter/material.dart';

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
                                return theme.primaryColorLight.withOpacity(0.4);
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
