import 'dart:math';

import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  final List<Container> myList = List.generate(
    90,
    (index) {
      return Container(
        color: Color.fromARGB(
          255,
          Random().nextInt(256),
          Random().nextInt(256),
          Random().nextInt(256),
        ),
      );
    },
  );
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Grid View'),
        ),
        body: GridView.count(
          crossAxisCount: 3,
          childAspectRatio: 3 / 4,
          padding: EdgeInsets.all(15),
          crossAxisSpacing: 10,
          mainAxisSpacing: 10,
          children: myList,
        ),
      ),
    );
  }
}
