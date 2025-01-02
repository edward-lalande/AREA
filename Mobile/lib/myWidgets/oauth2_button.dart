import 'package:flutter/material.dart';

class OauthButton extends StatelessWidget {
    const OauthButton({
        super.key,
        required this.iconPath,
        required this.resize,
        this.resizePadding,
    });

    final String iconPath;
    final bool resize;
    final EdgeInsetsGeometry? resizePadding;

    @override
    Widget build(BuildContext context) {
        return Container(
            padding: !resize ? EdgeInsets.only(
                top: 20,
                bottom: 20,
                left: 20,
                right:  20,

            ) : resizePadding,
           
            decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(16),
                border: Border.all(
                    width: 2,
                    color: Colors.black
                    ),
                ),
                child: Image.asset(
                    iconPath,
                    height: 40,
             ),
        );
    }
}
