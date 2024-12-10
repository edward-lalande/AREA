import 'package:flutter/material.dart';
import '../myWidgets/my_title.dart';


class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
    int index = 0;

    @override
    Widget build(BuildContext context) {
        return SafeArea(
            child: Scaffold(
            backgroundColor: Colors.white,
            bottomNavigationBar: BottomNavigationBar(
                onTap: (value) {
                    setState(() {
                      index = value;
                    });
                },
                currentIndex: index,
                backgroundColor: Colors.white,
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
                            color: Colors.black,
                            Icons.home,
                        ),
                        label: "Home",
                    ),
                    BottomNavigationBarItem(
                        activeIcon: Icon(Icons.create_outlined),
                        icon: Icon(
                            color: Colors.black,
                            Icons.create,
                        ),
                        label: "Create",
                    ),
                    BottomNavigationBarItem(
                        activeIcon: Icon(Icons.my_library_add_outlined),
                        icon: Icon(
                            color: Colors.black,
                            Icons.my_library_add,
                        ),
                        label: "MyArea's",
                    ),
                    BottomNavigationBarItem(
                        activeIcon: Icon(Icons.account_box_outlined),
                        icon: Icon(
                            color: Colors.black,
                            Icons.account_box,
                        ),
                        label: "Account",
                    ),
                ],
                    ),
            body: Column(
                children: [
                    const MyTitle(
                      title: "AREA",
                      fontSize: 45,
                      padding: EdgeInsets.only(top: 30),
                      color: Colors.black
                    ),
                ],
            ),
        )
            );
    }
}
