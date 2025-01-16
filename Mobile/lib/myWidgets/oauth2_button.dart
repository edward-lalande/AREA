import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_web_view.dart';
import 'package:second_app/utils/post_request.dart';

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

enum OauthService {
    google,
    discord,
    spotify,
    github,
    gitlab,
    dropbox
}

final oauthButtonsData = [
    {
        'service': OauthService.google,
        'iconPath': 'assets/google.png',
        'url': 'google/oauth',
        'resize': false,
    },
    {
        'service': OauthService.discord,
        'iconPath': 'assets/discord.png',
        'url': 'discord/oauth',
        'resize': false,
    },
    {
        'service': OauthService.spotify,
        'iconPath': 'assets/spotify.png',
        'url': 'spotify/oauth',
        'resize': false,
    },
    {
        'service': OauthService.github,
        'iconPath': 'assets/github.png',
        'url': 'github/oauth',
        'resize': true,
        'resizePadding': const EdgeInsets.only(top: 20, bottom: 20, left: 5, right: 5),
    },
    {
        'service': OauthService.gitlab,
        'iconPath': 'assets/gitlab.png',
        'url': 'gitlab/oauth',
        'resize': false,
    },
    {
        'service': OauthService.dropbox,
        'iconPath': 'assets/dropbox.png',
        'url': 'dropbox/oauth',
        'resize': false,
    },
];

class OAuthButtonsRow extends StatelessWidget {
    final String host;

    const OAuthButtonsRow({super.key, required this.host});

    @override
    Widget build(BuildContext context) {
        return Wrap(
            spacing: 20,
            runSpacing: 20,
            alignment: WrapAlignment.center,
            children: oauthButtonsData.map((buttonData) {
                final service = buttonData['service'] as OauthService;
                final iconPath = buttonData['iconPath'] as String;
                final url = buttonData['url'] as String;
                final resize = buttonData['resize'] as bool;
                final resizePadding = buttonData['resizePadding'] as EdgeInsets?;

                return OauthButton(
                    iconPath: iconPath,
                    resize: resize,
                    resizePadding: resizePadding,
                    onPressed: (context) async {
                        String oauthUrl = await classicGet(url: 'http://$host:8080/$url');
                        if (context.mounted) {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => WebViewPage(url: oauthUrl, serv: service.name),
                            ),
                          );
                        }
                    },
                );
            }).toList(),
        );
    }
}
