import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'package:webview_flutter/webview_flutter.dart';

class WebViewPage extends StatefulWidget {
    final String url;

    const WebViewPage({
        super.key,
        required this.url
    });

    @override
    State<WebViewPage> createState() => _WebViewPageState();
}

class _WebViewPageState extends State<WebViewPage> {
    late final WebViewController _controller;

    @override
    void initState() {

        super.initState();
        _controller = WebViewController()
        ..setJavaScriptMode(JavaScriptMode.unrestricted)
        ..loadRequest(Uri.parse(widget.url));
    }

    @override
    Widget build(BuildContext context) {
        return Scaffold(

            appBar: AppBar(
                title: const Text('AREA'),
                leading: IconButton(
                    icon: const Icon(Icons.arrow_back),
                    onPressed: () => context.pop(),
                ),
            ),

            body: WebViewWidget(
                controller: _controller
            ),

        );
    }
}
