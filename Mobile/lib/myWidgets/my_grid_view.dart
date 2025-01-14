import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_dialog.dart';
import 'package:second_app/myWidgets/my_title.dart';
import 'package:second_app/utils/post_request.dart';

extension HexColor on Color {

    static Color fromHex(String hexString) {
        final buffer = StringBuffer();
        if (hexString.length == 6 || hexString.length == 7) buffer.write('ff');
            buffer.write(hexString.replaceFirst('#', ''));
        return Color(int.parse(buffer.toString(), radix: 16));
    }

}

class ServiceCard extends StatelessWidget {
    const ServiceCard({
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
                width: 80,
                height: 100,
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

    const MyGridViewHome({
        super.key,
        required this.servicesMap
    });

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

                return ServiceCard(
                    title: name,
                    iconPath: iconPath,
                );
            },
        );
    }
}

class ActionsPage extends StatelessWidget {
    final Service service;

    const ActionsPage({
        super.key,
        required this.service,
    });

    @override
    Widget build(BuildContext context) {
        return Scaffold(
            appBar: AppBar(
                title: Text("Choice of actions"),
            ),
            body: Column(
                children: [
                    SizedBox(height: 80,),
                    MyTitle2(title: service.name,
                        fontSize: 40, padding: EdgeInsets.only(bottom: 50)
                    ),
                    ListView.builder(
                        padding: EdgeInsets.only(left: 25, right: 25),
                        shrinkWrap: true,
                        itemCount: service.actions.length,
                        itemBuilder: (context, index) {
                            final action = service.actions[index];
                            return InkWell(
                                onTap: () {
                                    showDialog(
                                        context: context,
                                        builder: (context) => ActionDialog(action: action),
                                    );
                                },
                                child: Card(
                                    child: Padding(
                                        padding: const EdgeInsets.all(25.0),
                                        child: Text(
                                            action.name,
                                        style: TextStyle(fontSize: 18),
                                        ),
                                    ),
                                ),
                            );
                        },
                    ),
                ],
            )
        );
    }
}

class ReactionsPage extends StatelessWidget {

    final ReactionService service;


    const ReactionsPage({
        super.key,
        required this.service,
    });

    @override
    Widget build(BuildContext context) {
    final scrollController = ScrollController();
        return Scaffold(
            appBar: AppBar(
                title: Text("Choice of reactions"),
            ),
            body: Padding(
                padding: const EdgeInsets.only(left: 5, right: 14),
                child: RawScrollbar(
                radius: Radius.circular(10),
                thumbColor: Theme.of(context).primaryColor,
                    thickness: 5,
                    controller: scrollController,
                    thumbVisibility: true,
                child: SingleChildScrollView(
                    controller: scrollController,
                    child: Column(
                        children: [
                            SizedBox(height: 80),
                            MyTitle2(
                                title: service.name,
                                fontSize: 40,
                                padding: EdgeInsets.only(bottom: 50),
                            ),
                            ListView.builder(
                                shrinkWrap: true,
                                padding: EdgeInsets.symmetric(horizontal: 25),
                                itemCount: service.reactions.length,
                                itemBuilder: (context, index) {
                                    final reaction = service.reactions[index];
                                    return InkWell(
                                        onTap: () {
                                            showDialog(
                                                context: context,
                                                builder: (context) => ReactionDialog(reaction: reaction),
                                            );
                                        },
                                        child: Card(
                                            child: Padding(
                                                padding: const EdgeInsets.all(25),
                                                child: Text(
                                                    reaction.name,
                                                    style: TextStyle(fontSize: 18),
                                                ),
                                            ),
                                        ),
                                    );
                                },
                            ),
                        ],
                    ),
                )
            ),
        )
        );
    }
}

class ReactionsGrid extends StatelessWidget {

    final List<ReactionService> reactionServices;

    const ReactionsGrid({
        super.key,
        required this.reactionServices,
    });

    @override
    Widget build(BuildContext context) {
        return GridView.builder(
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            padding: const EdgeInsets.all(20),
            gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2,
                crossAxisSpacing: 10,
                mainAxisSpacing: 10,
                childAspectRatio: 1,
            ),
            itemCount: reactionServices.length,
            itemBuilder: (context, index) {
                final reactionService = reactionServices[index];
                return InkWell(
                    onTap: () {
                        Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) => ReactionsPage(service: reactionService),
                            ),
                        );
                    },
                    child: ServiceCard(
                        title: reactionService.name,
                        iconPath: "assets/${reactionService.name.toLowerCase().replaceAll(' ', '_')}.png",
                    ),
                );
            },
        );
    }
}


class ServicesGrid extends StatelessWidget {
    final List<Service> services;

    const ServicesGrid({
        super.key,
        required this.services
    });

    @override
    Widget build(BuildContext context) {

        return GridView.builder(
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            padding: const EdgeInsets.all(20),
            gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 2,
                crossAxisSpacing: 10,
                mainAxisSpacing: 10,
                childAspectRatio: 1,
            ),
            itemCount: services.length,
            itemBuilder: (context, index) {
                final service = services[index];
                return InkWell(
                    onTap: () {
                        Navigator.push(
                            context,
                            MaterialPageRoute(
                                builder: (context) => ActionsPage(service: service),
                            ),
                        );
                    },
                    child: ServiceCard(
                        title: service.name, iconPath: "assets/${service.name.toLowerCase().replaceAll(' ', '_')}.png")
                );
            },
        );
    }
}
