import 'package:flutter/material.dart';

class MyTextField2 extends StatefulWidget {
    const MyTextField2({
        super.key,
        required this.hintText,
        required this.controller,
        this.prefixIcon,
        this.obscureText = false,
        this.padding = const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
    });

    final String hintText;
    final TextEditingController controller;
    final Widget? prefixIcon;
    final bool obscureText;
    final EdgeInsets padding;

    @override
    State<MyTextField2> createState() => _MyTextField2State();
}

class _MyTextField2State extends State<MyTextField2> {

    final FocusNode _focusNode = FocusNode();
    bool _isFocused = false;

    @override
    void initState() {
        super.initState();
        _focusNode.addListener(() {
            setState(() {
                _isFocused = _focusNode.hasFocus;
            });
        });
    }

    @override
    void dispose() {
        _focusNode.dispose();
        super.dispose();
    }

    @override
    Widget build(BuildContext context) {

        final theme = Theme.of(context);

        return Container(
            color: theme.scaffoldBackgroundColor,
            padding: widget.padding,
            child: AnimatedContainer(
                duration: const Duration(milliseconds: 200),
                decoration: BoxDecoration(
                    color: theme.cardColor,
                    borderRadius: BorderRadius.circular(12),
                    border: Border.all(
                        color: _isFocused ? theme.primaryColor : theme.dividerColor,
                        width: 2,
                    ),
                    boxShadow: [
                        BoxShadow(
                            color: theme.shadowColor.withOpacity(0.1),
                            spreadRadius: 1,
                            blurRadius: 5,
                            offset: const Offset(0, 3),
                        ),
                    ],
                ),
                child: TextField(
                    focusNode: _focusNode,
                    controller: widget.controller,
                    obscureText: widget.obscureText,
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
                    ),
                ),
            ),
        );
    }
}
