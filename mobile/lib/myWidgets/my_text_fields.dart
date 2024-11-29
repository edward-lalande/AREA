import 'package:flutter/material.dart';

class MyTextField extends StatelessWidget {
  const MyTextField({super.key, required this.hintText, required this.bgColor,
  required this.fieldBgColor, required this.hintTextColor, required this.inputColor,
  required this.padding, required this.prefixIcon});

  final String hintText;
  final Color hintTextColor;
  final Color inputColor;
  final Color bgColor;
  final Color fieldBgColor;
  final EdgeInsets padding;
  final Widget prefixIcon;

  @override
  Widget build(BuildContext context) {
    return Container(
      color: bgColor,

      child: Padding(
        padding: padding,
        child: Container(
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(12),
            color: fieldBgColor,
          ),
          child: TextField(
            textAlignVertical: TextAlignVertical.center,
            style: TextStyle(color: inputColor),
            decoration: InputDecoration(
              border: InputBorder.none,
              prefixIcon: prefixIcon,
              hintText: hintText,
              hintStyle: TextStyle(color: hintTextColor),
            ),
          ),
        ),
      ),
    );
  }
}
