import 'package:flutter/material.dart';

class MyTextField2 extends StatefulWidget {
    const MyTextField2({
        super.key,
        required this.hintText,
        required this.controller,
        required this.color,
        this.prefixIcon,
        this.suffixIcon,
        this.onSuffixIconPressed,
        this.obscureText = false,
        this.padding = const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
    });

    final String hintText;
    final TextEditingController controller;
    final Widget? prefixIcon;
    final Widget? suffixIcon;
    final VoidCallback? onSuffixIconPressed;
    final bool obscureText;
    final EdgeInsets padding;
    final Color color;

    @override
    State<MyTextField2> createState() => _MyTextField2State();
}

class _MyTextField2State extends State<MyTextField2> {

    late bool _isObscure;

    @override
    void initState() {
        super.initState();
        _isObscure = widget.obscureText;
    }

    @override
    Widget build(BuildContext context) {
        final theme = Theme.of(context);

        return Container(
            color: widget.color,
            padding: widget.padding,
            child: AnimatedContainer(
                duration: const Duration(milliseconds: 200),
                decoration: BoxDecoration(
                    color: theme.cardColor,
                    borderRadius: BorderRadius.circular(12),
                    border: Border.all(
                        color: theme.dividerColor,
                        width: 2,
                    ),
                    boxShadow: [
                        BoxShadow(
                        color: theme.shadowColor.withOpacity(0.5),
                        spreadRadius: 1,
                        blurRadius: 5,
                        offset: const Offset(0, 3),
                        ),
                    ],
                    ),
                    child: TextField(
                      controller: widget.controller,
                      obscureText: _isObscure,
                      textAlignVertical: TextAlignVertical.center,
                      style: TextStyle(
                          color: theme.textTheme.bodyLarge?.color,
                          fontFamily: "Avenir",
                      ),
                      decoration: InputDecoration(
                          border: InputBorder.none,
                          prefixIcon: widget.prefixIcon,
                          hintText: widget.hintText,
                          hintStyle: TextStyle(color: theme.hintColor),
                          contentPadding: const EdgeInsets.symmetric(vertical: 16, horizontal: 12),
                          suffixIcon: widget.obscureText
                              ? GestureDetector(
                                  onTap: widget.onSuffixIconPressed ??
                                      () {
                                      setState(() {
                                          _isObscure = !_isObscure;
                                      });
                                      },
                                  child: Icon(
                                  _isObscure ? Icons.visibility_off : Icons.visibility,
                                  color: theme.iconTheme.color,
                                  ),
                              )
                              : null,
                    ),
                ),
            ),
        );
    }
}
