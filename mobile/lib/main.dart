import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'pages/home_page.dart';
import 'pages/login_page.dart';
import 'pages/password_page.dart';
import 'pages/signup_page.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      routerConfig: appRouter(context),
    );
  }
}

GoRouter appRouter(BuildContext context) {
  return GoRouter(
    routes: [
       GoRoute(
        path: '/',
        builder: (context, state) => LoginPage(),
      ),
      GoRoute(
        path: '/login',
        builder: (context, state) => LoginPage(),
      ),
      GoRoute(
        path: '/home',
        builder: (context, state) => const HomePage(),
      ),
      GoRoute(
        path: '/password',
        builder: (context, state) => PasswordPage(),
      ),
      GoRoute(
        path: '/signup',
        builder: (context, state) => SignUpPage(),
      ),
    ],
     errorBuilder: (context, state) => const  Scaffold(
      body: Center(
        child: Text('Page introuvable'),
      ),
    ),
    );
}