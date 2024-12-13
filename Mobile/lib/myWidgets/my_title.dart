import 'package:flutter/material.dart';

class MyTitle extends StatelessWidget {

  const MyTitle({
    super.key,
    required this.title,
    required this.fontSize,
    required this.padding,
    required this.color,
  });

  final String title;
  final double fontSize;
  final EdgeInsets padding;
  final Color color;

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      width: MediaQuery.sizeOf(context).width,
      child: Padding(
        padding: padding,
        child: Text(
          title,
          textAlign: TextAlign.center,
          style: TextStyle(
            fontFamily: "Avenir",
            fontSize: fontSize,
            color: color,
          ),
        ),
      ),
    );
  }
}