import 'package:flutter/material.dart';

import '../myWidgets/area_app_bar.dart';


class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {

  int _selectedIndex = 0;

  final List<Widget> _pages = [
    const Center(
      child: Text(
        'Créer',
        style: TextStyle(
          fontFamily: "Avenir",
          fontSize: 30,
          fontWeight: FontWeight.w900,
        ),
      ),
    ),

    const Center(
      child: Text(
        'My AREA',
        style: TextStyle(
          fontFamily: "Avenir",
          fontSize: 30,
          fontWeight: FontWeight.w900,
        ),
      ),
    ),

    const Center(
      child: Text(
        'Account',
        style: TextStyle(
          fontFamily: "Avenir",
          fontSize: 30,
          fontWeight: FontWeight.w900,
        ),
      ),
    ),
  ];
  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const MyAppBarArea(
        appbartitle: Padding(
          padding: EdgeInsets.only(top: 45),
          child: Text(
            "AREA",
            style: TextStyle(
              fontFamily: "Avenir",
              fontSize: 65,
            ),
          ),
        ),
      ),
      body: _pages[_selectedIndex],
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _selectedIndex,
        onTap: _onItemTapped,
        selectedItemColor: Colors.black, 
        unselectedItemColor: Colors.grey,
        showUnselectedLabels: true,
        backgroundColor: Colors.white,
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.create),
            label: 'Créer',
            backgroundColor: Colors.white,
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'My AREA',
            backgroundColor: Colors.white,
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.account_circle),
            label: 'Account',
            backgroundColor: Colors.white,
          ),
        ],
      ),
    );
  }
}
