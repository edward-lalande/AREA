import 'package:flutter/material.dart';

class ThemeProvider with ChangeNotifier {

    bool _isDarkMode = true;
    bool get isDarkMode => _isDarkMode;

    Color _customDarkPrimaryColor = Colors.black;
    Color get customDarkPrimaryColor => _customDarkPrimaryColor;

    Color _customDarkBackgroundColor = Colors.black;
    Color get customDarkBackgroundColor => _customDarkBackgroundColor;

    void toggleTheme()
    {
        _isDarkMode = !_isDarkMode;
        notifyListeners();
    }

    void updateCustomDarkPrimaryColor(Color color)
    {
        if (_isDarkMode) {
        _customDarkPrimaryColor = color;
        _customDarkBackgroundColor = _adjustBackgroundColor(color);
        notifyListeners();
        }
    }

  Color _adjustBackgroundColor(Color color)
  {
    return Color.alphaBlend(Colors.black.withOpacity(0.8), color);
  }

  ThemeMode get currentTheme => _isDarkMode ? ThemeMode.dark : ThemeMode.light;
}


