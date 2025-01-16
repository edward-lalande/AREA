import 'package:flutter/material.dart';

class MyTitle2 extends StatelessWidget {
    const MyTitle2({
        super.key,
        required this.title,
        required this.fontSize,
        required this.padding,
        this.margin,
    });

    final String title;
    final double fontSize;
    final EdgeInsets padding;
    final EdgeInsetsGeometry? margin;

    @override
    Widget build(BuildContext context) {
        return Container(
            color: Theme.of(context).scaffoldBackgroundColor,
            width: MediaQuery.sizeOf(context).width,
            margin: margin,
            child: Padding(
                padding: padding,
                child: Text(
                    title,
                    textAlign: TextAlign.center,
                    style: TextStyle(
                        fontFamily: "Avenir",
                        fontSize: fontSize,
                        color: Theme.of(context).textTheme.bodyLarge?.color,
                    ),
                ),
            ),
        );
    }
}
