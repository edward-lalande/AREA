import 'package:flutter/material.dart';
import 'package:second_app/theme/theme_provider.dart';

final ThemeData lightTheme = ThemeData(
    brightness: Brightness.light,
    scaffoldBackgroundColor: Colors.white,
    fontFamily: "Avenir",
    cardColor: Colors.white,
    primaryColor: Colors.black,
    primaryColorLight: Colors.grey[300],
    textTheme: const TextTheme(
        bodyLarge: TextStyle(color: Colors.black),
        bodyMedium: TextStyle(color: Colors.black),
        bodySmall: TextStyle(color: Colors.black54),
    ),
    appBarTheme: const AppBarTheme(
        backgroundColor: Colors.white,
        iconTheme: IconThemeData(color: Colors.black),
    ),
);

ThemeData darkTheme(ThemeProvider themeProvider)
{
    return ThemeData(
        brightness: Brightness.dark,
        scaffoldBackgroundColor: themeProvider.customDarkBackgroundColor,
        fontFamily: "Avenir",
        cardColor: themeProvider.customDarkPrimaryColor == Colors.black ? Colors.grey[900] : themeProvider.customDarkPrimaryColor.withOpacity(0.9),
        primaryColor: themeProvider.customDarkPrimaryColor,
        primaryColorLight: themeProvider.customDarkPrimaryColor.withOpacity(0.3),
        textTheme: const TextTheme(
            bodyLarge: TextStyle(color: Colors.white),
            bodyMedium: TextStyle(color: Colors.white),
            bodySmall: TextStyle(color: Colors.white70),
        ),
        appBarTheme: AppBarTheme(
            backgroundColor: themeProvider.customDarkPrimaryColor,
            iconTheme: const IconThemeData(color: Colors.white),
        ),
    );
}
