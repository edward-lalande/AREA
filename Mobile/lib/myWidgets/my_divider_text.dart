import 'package:flutter/material.dart';

class MyDividerText extends StatelessWidget {

  const MyDividerText({
    super.key,
    required this.bgColor,
    required this.padding,
    required this.textBetween
  });

  final Color bgColor;
  final EdgeInsets padding;
  final String textBetween;


  @override
  Widget build(BuildContext context) {
    return Container(
      color: bgColor,
      child: Padding(
        padding: padding,
        child: Row(
        children: [
          const Expanded(
            child: Padding(
              padding: EdgeInsets.only(right: 10),
              child: Divider(),
              )
          ),
          Text(textBetween,
          style: const TextStyle(fontFamily: "Avenir"),),
          const Expanded(
            child: Padding(
              padding: EdgeInsets.only(left: 10),
              child: Divider(),
              )
          ),
        ],
      )
      ),
    );
  }
}