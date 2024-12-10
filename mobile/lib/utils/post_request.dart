import 'package:http/http.dart' as http;
import 'dart:convert';

Future<bool> sendSignUp({Map<String, dynamic>? body, Map<String, String>? headers, required String url}) async
{
  //'http://127.0.0.1:8080/user'

  try {
    final response = await http.post(
      Uri.parse(url),
      headers: headers,
      body: json.encode(body),
    );
    if (response.statusCode == 200) {
      return true;
    } else {
      print('ERRRORR : ${response.statusCode}, ${response.body}');
      return false;
    }
  } catch (e) {
    print('ERRORRRRR : $e');
    return false;
  }
}


Future<String> getOAuthUrl(String  service) async
{
    //url en dure discord => MVP dans 2 jours
    // "http://10.0.2.2:8083/oauth2"

    final apiUrl = "http://10.0.2.2:8083/oauth2";
    try {

      final response = await http.get(Uri.parse(apiUrl));

      if (response.statusCode == 200) {
        return response.body;
      }
      else {
        throw Exception('Failed to load data: ${response.statusCode}');
      }

    } catch (e) {

      throw Exception('Error fetching data: $e');

    }
}

//class exeption + secure+strorage

/*void logByOAuth(BuildContext context, String apiUrl) async
{

    try {
        String fetchedUrl = await getOAuthUrl(apiUrl);

        if (context.mounted) {
            print(fetchedUrl);
        }
    } catch (e) {
        print("ERRORRRR");
    }
}*/