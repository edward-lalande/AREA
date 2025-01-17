import 'package:flutter/material.dart';

class MyTextButton extends StatelessWidget {

    const MyTextButton({
        super.key,
        required this.onTap,
        required this.firstTitle,
        required this.secondTitle,
        required this.padding,
        this.mainAxisAlignment,
    });

    final void Function(BuildContext) onTap;
    final String firstTitle;
    final String secondTitle;
    final EdgeInsets padding;
    final MainAxisAlignment? mainAxisAlignment;


    @override
    Widget build(BuildContext context) {
        return Container(
            color: Theme.of(context).scaffoldBackgroundColor,
            child: Padding(
                padding: padding,
                child: Row(
                    mainAxisAlignment: mainAxisAlignment ?? MainAxisAlignment.start,
                    children: [
                        Text(
                            firstTitle,
                            style: TextStyle(fontFamily: "Avenir", fontWeight: FontWeight.w300),
                        ),
                        const SizedBox(width: 7),
                        InkWell(
                            onTap: () => onTap(context),
                            child: Text(
                                secondTitle,
                                style: TextStyle(
                                    fontSize: 16,
                                    fontFamily: 'Avenir',
                                    color: Colors.blue,
                                    decoration: TextDecoration.underline,
                                    decorationColor: Colors.blue,
                                    decorationThickness: 2,
                                ),
                            ),
                        ),
                    ],
                ),
            ),
        );
    }
}