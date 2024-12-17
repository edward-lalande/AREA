import 'package:flutter/material.dart';
import 'package:second_app/myWidgets/my_card.dart';

class MyGridView extends StatelessWidget {
    const MyGridView({
        super.key,
        required this.title,
        required this.icon,
        required
    });

    //assets & color back-end

    // recup json avec tout
    // passer cette string en param ici
    // parser la string
    // boucler sur les infos (nb services) => full generique
    

    //c'est temp en dessous

    final Widget icon;
    final String title;
    final int color;
    @override
    Widget build(BuildContext context) {
        return Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
                Card(
                    elevation: 7,
                    color: Color(),
                    child: MyCard(
                        title: title,
                        icon: icon,
                        padding: const EdgeInsets.only(),
                    )
                ),
                const SizedBox(width: 5,),
                Card(
                    elevation: 7,
                    child: MyCard(
                        title: "Time User",
                        icon: icon,
                        padding: const EdgeInsets.only(),
                    )
                ),
            ],
        );
    }
}