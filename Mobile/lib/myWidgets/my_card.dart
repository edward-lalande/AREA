import 'package:flutter/material.dart';

class MyCard extends StatelessWidget {
    const MyCard({
        super.key,
        this.icon,
        required this.title,
        this.padding = const EdgeInsets.all(0),
    });

    final Widget? icon;
    final String title;
    final EdgeInsetsGeometry padding;

    @override
    Widget build(BuildContext context) {
        return Padding(
            padding: padding,
            child: SizedBox(
                width: 10,
                height: 10,
                child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                        icon ?? SizedBox(),
                        const SizedBox(
                            height: 8
                        ),
                        Text(
                            title,
                            textAlign: TextAlign.center,
                            style: TextStyle(
                                fontSize: 14,
                                color: Colors.black,
                                fontFamily: "Avenir"
                            ),
                        ),
                    ],
                ),
            ),
        );
    }
}