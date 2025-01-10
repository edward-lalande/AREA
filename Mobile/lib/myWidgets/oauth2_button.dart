import 'package:flutter/material.dart';

class OauthButton extends StatefulWidget {
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
    final void Function(BuildContext)? onPressed;

    @override
    _OauthButtonState createState() => _OauthButtonState();
}

class _OauthButtonState extends State<OauthButton> {
    late FocusNode _focusNode;

    @override
    void initState() {
        super.initState();
        _focusNode = FocusNode();
    }

    @override
    void dispose() {
        _focusNode.dispose();
        super.dispose();
    }

    @override
    Widget build(BuildContext context) {
        return Card(
            elevation: 5,
            shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(16),
            ),
            child: Focus(
                focusNode: _focusNode,
                child: InkWell(
                    borderRadius: BorderRadius.circular(16),
                    onTap: () {
                        if (widget.onPressed != null) {
                        widget.onPressed!(context);
                        }
                    },
                    onFocusChange: (hasFocus) {
                        setState(() {});
                    },
                    child: Padding(
                        padding: !widget.resize
                            ? const EdgeInsets.all(20)
                            : widget.resizePadding ?? const EdgeInsets.all(20),
                        child: Image.asset(
                            widget.iconPath,
                            height: 40,
                            ),
                    ),
                ),
            ),
        );
    }
}
