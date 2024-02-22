import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'controller/counter_controller.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  final counterController = Get.put(CounterController());

  @override
  Widget build(BuildContext context) {
    return Obx(() => MaterialApp(
          debugShowCheckedModeBanner: false,
          theme: counterController.isDark.value
              ? ThemeData.dark()
              : ThemeData.light(),
          home: HomePage(),
        ));
  }
}

class HomePage extends StatelessWidget {
  final counter_controller = Get.find<CounterController>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
      body: Center(
        child: Obx(
          () => Text(
            "Angka ${counter_controller.counter}",
            style: const TextStyle(fontSize: 35),
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => counter_controller.changeTheme(),
        child: const Icon(Icons.add),
      ),
    );
  }
}
