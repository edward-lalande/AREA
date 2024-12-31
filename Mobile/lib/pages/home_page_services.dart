import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_grid_view.dart';
import 'package:second_app/utils/post_request.dart';

class HomePageServices extends StatefulWidget {
    const HomePageServices({super.key});

    @override
    State<HomePageServices> createState() => _HomePageServicesState();
}

class _HomePageServicesState extends State<HomePageServices> {
    @override
    Widget build(BuildContext context) {
        return SafeArea(
            child: Scaffold(
                backgroundColor: Colors.white,
                body: MyGridView(
                    homeAnimation: true,
                    appbarVisible: true,
                    map: servicesMap,
                    typeKey: "services",
                ),
            ),
        );
    }
}
