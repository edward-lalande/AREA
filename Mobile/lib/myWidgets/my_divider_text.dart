import 'package:flutter/material.dart';

class MyDividerText2 extends StatelessWidget {
    const MyDividerText2({
        super.key,
        required this.padding,
        required this.textBetween,
    });

    final EdgeInsets padding;
    final String textBetween;

    @override
    Widget build(BuildContext context) {

        final theme = Theme.of(context);

        return Container(
            color: theme.scaffoldBackgroundColor,
            child: Padding(
                padding: padding,
                child: Row(
                    children: [
                        Expanded(
                            child: Padding(
                                padding: const EdgeInsets.only(right: 10),
                                child: Divider(
                                    color: theme.dividerColor,
                                    thickness: 1.5,
                                ),
                            ),
                        ),
                        Text(
                            textBetween,
                            style: theme.textTheme.bodyLarge?.copyWith(
                                fontFamily: "Avenir",
                            ),
                        ),
                        Expanded(
                            child: Padding(
                                padding: const EdgeInsets.only(left: 10),
                                child: Divider(
                                color: theme.dividerColor,
                                thickness: 1.5,
                                ),
                            ),
                        ),
                    ],
                ),
            ),
        );
    }
}
