import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'package:webview_flutter/webview_flutter.dart';
import '../utils/post_request.dart';

//import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class WebViewPage extends StatefulWidget {
  final String url;

  const WebViewPage({
    super.key,
    required this.url,
  });

  @override
  State<WebViewPage> createState() => _WebViewPageState();
}

class _WebViewPageState extends State<WebViewPage> {
  late final WebViewController _controller;
  //final FlutterSecureStorage _secureStorage = const FlutterSecureStorage();
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();
    _controller = WebViewController()
      ..setJavaScriptMode(JavaScriptMode.unrestricted)
      ..setNavigationDelegate(
        NavigationDelegate(
          onPageStarted: (String url) {
            setState(() {
              _isLoading = true;
            });
          },
          onPageFinished: (String url) async {
            setState(() {
              _isLoading = false;
            });

            if (url.contains("code=")) {
              final Uri uri = Uri.parse(url);
              final String? code = uri.queryParameters["code"];

              if (code != null) {

                //await _secureStorage.write(key: "discord_code", value: code);
                bool tmp = await sendSignUp(
                    url: "http://10.0.2.2:8080/discord",
                    body: {
                        "routes": "access-token",
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

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Authentification'),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => Navigator.pop(context),
        ),
      ),
      body: Stack(
        children: [
          WebViewWidget(controller: _controller),
          if (_isLoading)
            const Center(
              child: CircularProgressIndicator(),
            ),
        ],
      ),
    );
  }
}

