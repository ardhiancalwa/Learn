import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  List<Tab> myTab = [
    Tab(
      text: "Home",
      icon: Icon(Icons.home),
    ),
    Tab(
      text: "Shop",
      icon: Icon(Icons.shopping_bag_rounded),
    ),
    Tab(
      text: "GOO",
      icon: Icon(Icons.shopping_basket_rounded),
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: DefaultTabController(
        initialIndex: 1,
        length: myTab.length,
        child: Scaffold(
          appBar: AppBar(
            title: const Text('Tab Bar'),
            centerTitle: true,
            bottom: TabBar(
              labelColor: Colors.blue[700],
              labelStyle: TextStyle(fontWeight: FontWeight.bold),
              unselectedLabelColor: Colors.blue[200],
              unselectedLabelStyle: TextStyle(fontWeight: FontWeight.normal),
              indicator: BoxDecoration(
                // color: Colors.purple[100],
                border: Border(
                  bottom: BorderSide(
                    color: Colors.yellow,
                    width: 3
                  )
                )
              ),
              tabs: myTab,
            ),
          ),
          body: TabBarView(children: [
            Center(
              child: Text("Home", style: TextStyle(fontSize: 50),),
            ),
            Center(
              child: Text("Shop", style: TextStyle(fontSize: 50),),
            ),
            Center(
              child: Text("GOO", style: TextStyle(fontSize: 50),),
            )
          ]),
        ),
      ),
    );
  }
}
