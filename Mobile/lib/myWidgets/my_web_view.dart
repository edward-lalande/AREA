import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'package:webview_flutter/webview_flutter.dart';
import 'package:second_app/utils/post_request.dart';

class WebViewPage extends StatefulWidget {
    final String url;
    final String serv;

    const WebViewPage({
        super.key,
        required this.url,
        required this.serv,
    });
    @override
    State<WebViewPage> createState() => _WebViewPageState();
}

class _WebViewPageState extends State<WebViewPage> {

    late final WebViewController _controller;

    bool _isLoading = true;

    @override
    void initState() {
        super.initState();
        String servRoot = widget.serv;
        _controller = WebViewController()
        ..setJavaScriptMode(JavaScriptMode.unrestricted)
        ..setUserAgent("random")
        ..setNavigationDelegate(
            NavigationDelegate(
                onPageStarted: (String url) {
                    setState(() {
                    _isLoading = true;
                    });
                },
                onPageFinished: (String url) async {

                    await getDatas();

                    setState(() {
                        _isLoading = false;
                    });

                    if (url.contains("code=")) {

                        final Uri uri = Uri.parse(url);
                        final String? code = uri.queryParameters["code"];
                        if (code != null) {
                            bool tmp = await sendSignUp(
                                delim: 9,
                                url: "$host/$servRoot/access-token",
                                body: {
                                    "code": code,
                                }
                            );
                            if (tmp) {
                                if (context.mounted) {
                                    context.go("/home");
                                }
                            } else {
                                if (context.mounted) {
                                    context.go("/login");
                                }
                            }
                        }
                    }
                },
            ),
        )
        ..loadRequest(Uri.parse(widget.url));
    }

    String getService(String url)
    {
        StringBuffer service = StringBuffer();

        for (var i = url.length - 7; i >= 0; i--) {
            if (url[i] == '/') {
                break;
            }
            service.write(url[i]);
        }
        return service.toString().split('').reversed.join();
    }
    @override
    Widget build(BuildContext context) {
        return Scaffold(
            appBar: AppBar(
                backgroundColor: Theme.of(context).scaffoldBackgroundColor,
                    title: const Text('Authentification'),
                    leading: IconButton(
                        icon: const Icon(Icons.arrow_back),
                        onPressed: () => Navigator.pop(context),
                    ),
            ),
            body: Stack(
                children: [
                    WebViewWidget(
                        controller: _controller
                    ),
                    if (_isLoading)
                        Center(
                            child: CircularProgressIndicator(color: Theme.of(context).scaffoldBackgroundColor,),
                        ),
                ],
            ),
        );
    }
}

