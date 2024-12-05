import 'package:flutter/material.dart';

class MyAppBarArea extends StatelessWidget implements PreferredSizeWidget {
  const MyAppBarArea({super.key, required this.appbartitle});
    final Widget appbartitle;
  @override
  Widget build(BuildContext context) {
    return AppBar(
        toolbarHeight: 200,
        centerTitle: true,
        backgroundColor: Colors.white,
        title: appbartitle,
        scrolledUnderElevation: 0.0,
        surfaceTintColor: Colors.transparent,
    );
  }
  @override
  Size get preferredSize => const Size.fromHeight(160.0);

}