import 'package:flutter/material.dart';

import '../pages/home_page_services.dart';
import '../pages/create_area_page.dart';
import '../pages/my_area.dart';
import '../pages/account_page.dart';

class HomePage extends StatefulWidget {
    const HomePage({super.key});

    @override
    State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {

    int index = 0;
    final _pageOptions = [
        HomePageServices(),
        CreateArea(),
        MyArea(),
        AccountPage(),
    ];

    @override
    Widget build(BuildContext context) {
        return Scaffold(
            backgroundColor: Theme.of(context).scaffoldBackgroundColor,
            bottomNavigationBar: BottomNavigationBar(
                elevation: 40,
                onTap: (value) {
                    setState(() {
                        index = value;
                    });
                },
                currentIndex: index,
                backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                type: BottomNavigationBarType.fixed,
                iconSize: 25,
                selectedLabelStyle: TextStyle(
                    fontFamily: "Avenir",
                ),
                unselectedLabelStyle: TextStyle(
                    fontFamily: "Avenir",
                ),
                items: [
                    BottomNavigationBarItem(
                        activeIcon: Icon(Icons.home_outlined),
                        icon: Icon(
                            Icons.home,
                        ),
                        label: "Home",
                    ),
                    BottomNavigationBarItem(
                        activeIcon: Icon(Icons.create_outlined),
                        icon: Icon(
                            Icons.create,
                        ),
                        label: "Create",
                    ),
                    BottomNavigationBarItem(
                        activeIcon: Icon(Icons.my_library_add_outlined),
                        icon: Icon(
                            Icons.my_library_books,
                        ),
                        label: "MyArea's",
                    ),
                    BottomNavigationBarItem(
                        activeIcon: Icon(Icons.account_box_outlined),
                        icon: Icon(
                            Icons.account_box,
                        ),
                        label: "Account",
                    ),
                ],
            ),
            body: _pageOptions[index],
        );
    }
}
