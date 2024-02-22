import 'package:flutter/material.dart';
import 'package:get/get.dart';
// import './models/person.dart';
// import './controllers/person_controller.dart';
import './controllers/counter_controller.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // var person = Person(name: "Calwa", age: 18).obs;
  // final personController = Get.put(PersonController());
  final countController = Get.put(CounterController());

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar: AppBar(),
        body: Center(
          // child: Obx(
          //   () => Text(
          //     "Nama saya ${personController.person.value.name}",
          //     style: const TextStyle(
          //       fontSize: 35,
          //     ),
          //   ),
          // ),
          // child: GetX<CounterController>(
          //   init: CounterController(),
          //   initState: (_) {},
          //   builder: (_) => Text(
          //     "Angka ${countController.count.value}",
          //     style: const TextStyle(
          //       fontSize: 35,
          //     ),
          //   ),
          // ),
          child: GetBuilder<CounterController>(
            // init: CounterController(),
            initState: (_) {},
            builder: (_) => Text(
              "Angka ${countController.count}",
              style: const TextStyle(
                fontSize: 35,
              ),
            ),
          ),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            // person.update((val) {
            //   person.value.name = person.value.name.toString().toUpperCase();
            // });

            // personController.changeUpperCase();
            // countController.increment();
            // Get.find<CounterController>().increment();
            countController.increment();
          },
          child: const Icon(Icons.add),
        ),
      ),
    );
  }
}
