import 'package:flutter/material.dart';

class MySwitchButton extends StatefulWidget {
    const MySwitchButton({
        super.key,
        required this.padding
    });

    final EdgeInsetsGeometry padding;
    @override
    State<MySwitchButton> createState() => _MySwitchButtonState();
}

class _MySwitchButtonState extends State<MySwitchButton> {

    static bool light = true;

    @override
    Widget build(BuildContext context) {
        return  Padding(
            padding: widget.padding,
            child: Switch(
                activeColor: Colors.grey,
                thumbIcon: WidgetStatePropertyAll(
                    Icon(
                        light ? Icons.light_mode : Icons.dark_mode,
                        color: light ? Colors.yellow : Colors.white,
                    ),
                ),
                value: light,
                onChanged: (bool value) {
                    setState(() {
                        light = value;
                    });
                },
            ),
        );
    }
}