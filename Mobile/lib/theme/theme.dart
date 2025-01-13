import 'package:flutter/material.dart';

final ThemeData lightTheme = ThemeData(
  brightness: Brightness.light,
  scaffoldBackgroundColor: Colors.white,
  fontFamily: "Avenir",
  cardColor: Colors.white,
  primaryColor: Colors.black, // 👈 Boutons et éléments interactifs en mode clair
  primaryColorLight: Colors.grey[300], // 👈 Couleur utilisée pour les états survolés
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

final ThemeData darkTheme = ThemeData(
  brightness: Brightness.dark,
  scaffoldBackgroundColor: Colors.black,
  fontFamily: "Avenir",
  cardColor: Colors.grey[900],
  primaryColor: Colors.black, // 👈 Boutons et éléments interactifs en mode sombre
  primaryColorLight: Colors.grey[700], // 👈 Couleur utilisée pour les états survolés
  textTheme: const TextTheme(
    bodyLarge: TextStyle(color: Colors.white),
    bodyMedium: TextStyle(color: Colors.white),
    bodySmall: TextStyle(color: Colors.white70),
  ),
  appBarTheme: const AppBarTheme(
    backgroundColor: Colors.black,
    iconTheme: IconThemeData(color: Colors.white),
  ),
);




