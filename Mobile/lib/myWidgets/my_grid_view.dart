import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_card.dart';

extension HexColor on Color {

    static Color fromHex(String hexString) {
        final buffer = StringBuffer();
        if (hexString.length == 6 || hexString.length == 7) buffer.write('ff');
            buffer.write(hexString.replaceFirst('#', ''));
        return Color(int.parse(buffer.toString(), radix: 16));
    }

}

class OauthButton extends StatelessWidget {
    const OauthButton({
      super.key,
      required this.title,
      required this.iconPath,
    });

    final String title;
    final String iconPath;

    @override
    Widget build(BuildContext context) {
        return Card(
            elevation: 5,
            shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(16),
            ),
            child: Container(
                width: 100,
                height: 120,
                padding: const EdgeInsets.all(12),
                child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                        Expanded(
                            child: Image.asset(
                                iconPath,
                                fit: BoxFit.contain,
                                errorBuilder: (context, error, stackTrace) {
                                return const Icon(
                                    Icons.broken_image,
                                    size: 50,
                                    color: Colors.grey,
                                );
                                },
                            ),
                        ),
                        const SizedBox(height: 8),
                        Text(
                            title,
                            textAlign: TextAlign.center,
                            style: const TextStyle(
                                fontFamily: "Avenir",
                                fontSize: 14,
                                fontWeight: FontWeight.bold,
                            ),
                        ),
                    ],
                ),
            ),
        );
    }
}

class MyGridViewHome extends StatelessWidget {
    final Map<String, dynamic> servicesMap;

    const MyGridViewHome({super.key, required this.servicesMap});

    @override
    Widget build(BuildContext context) {

        final services = servicesMap['services'] as List<dynamic>;

        return GridView.builder(
            padding: const EdgeInsets.all(20),
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2,
                crossAxisSpacing: 10,
                mainAxisSpacing: 20,
                childAspectRatio: 1,
            ),
            itemCount: services.length,
            itemBuilder: (context, index) {
                final service = services[index];
                final name = service['name'] as String;
                final iconPath = 'assets/${name.toLowerCase().replaceAll(' ', '_')}.png';

                return OauthButton(
                    title: name,
                    iconPath: iconPath,
                );
            },
        );
    }
}

class MyGridViewActionsName extends StatefulWidget {

    const MyGridViewActionsName({
        super.key,
        this.gridClick,
        required this.dataMap,
    });

    final Function(int idx)? gridClick;
    final Map<String, dynamic> dataMap;

    @override
    State<MyGridViewActionsName> createState() => _MyGridViewActionsNameState();
}

class _MyGridViewActionsNameState extends State<MyGridViewActionsName> {
    int selectedIndex = -1;

    void _onCardTap(int index, dynamic service) {
        setState(() {
            selectedIndex = selectedIndex == index ? -1 : index;
        });
        widget.gridClick!(index);
    }

    @override
    Widget build(BuildContext context) {
        List<String> keysList = widget.dataMap.keys.toList();

        return SingleChildScrollView(
            padding: EdgeInsets.only(bottom: 50, left: 50, right: 50),
            child: GridView.builder(
                shrinkWrap: true,
                physics: NeverScrollableScrollPhysics(),
                gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                    crossAxisCount: 2,
                    crossAxisSpacing: 10,
                    mainAxisSpacing: 10,
                ),
                itemCount: keysList.length,
                itemBuilder: (context, index) {
                    return InkWell(
                        onTap: () {
                            _onCardTap(index, widget.dataMap[keysList[index]]);
                        },
                        child: Card(
                            color: Colors.grey,
                            elevation: 7,
                            child: MyCard(
                                title: keysList[index],
                                padding: const EdgeInsets.all(8),
                            ),
                        ),
                    );
                },
            ),
        );
    }
}
