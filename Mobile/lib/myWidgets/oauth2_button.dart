import 'package:flutter/material.dart';

class OauthButton extends StatelessWidget {
    const OauthButton({
        super.key,
        required this.iconPath,
        required this.resize,
        this.resizePadding,
        this.onPressed,
    });

    final String iconPath;
    final bool resize;
    final EdgeInsetsGeometry? resizePadding;
    final VoidCallback? onPressed;

    @override
    Widget build(BuildContext context) {
        return GestureDetector(
            onTap: onPressed,
            child: Card(
                elevation:5, // Définit l'élévation
                shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(16),
                ),
                child: Padding(
                    padding: !resize
                        ? const EdgeInsets.all(20)
                        : resizePadding ?? const EdgeInsets.all(20),
                    child: Image.asset(
                        iconPath,
                        height: 40,
                    ),
                ),
            ),
        );
    }
}

