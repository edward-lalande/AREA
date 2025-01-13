import 'package:flutter/material.dart';

class ThemeProvider with ChangeNotifier
{

    bool _isDarkMode = true;
    bool get isDarkMode => _isDarkMode;

    void toggleTheme() {
        _isDarkMode = !_isDarkMode;
        notifyListeners();
    }

    ThemeMode get currentTheme => _isDarkMode ? ThemeMode.dark : ThemeMode.light;
}
