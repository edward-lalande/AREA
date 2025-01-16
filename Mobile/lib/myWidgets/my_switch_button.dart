import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:second_app/theme/theme_provider.dart';

class MySwitchButton extends StatelessWidget {
    const MySwitchButton({
        super.key,
        required this.padding,
    });

    final EdgeInsetsGeometry padding;

    @override
    Widget build(BuildContext context) {

        final themeProvider = Provider.of<ThemeProvider>(context);

        return Padding(
            padding: padding,
            child: Switch(
                activeColor: Colors.grey,
                thumbIcon: WidgetStatePropertyAll(
                    Icon(
                        themeProvider.isDarkMode ? Icons.dark_mode : Icons.light_mode,
                        color: themeProvider.isDarkMode ? Colors.white : Colors.yellow,
                    ),
                ),
                value: themeProvider.isDarkMode,
                onChanged: (value) {
                    themeProvider.toggleTheme();
                },
            ),
        );
    }
}
