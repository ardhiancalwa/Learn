import 'package:flutter/material.dart';
import 'package:get/get.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      debugShowCheckedModeBanner: false,
      home: HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  // final countController = Get.put(CounterController());

  // @override
  // void dispose() {
  //   countController.dispose(); // Dispose the controller
  //   super.dispose();
  // }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Home Page"),
        actions: [
          IconButton(
            onPressed: () => Get.to(() => TextPage()),
            icon: Icon(Icons.refresh),
          ),
        ],
      ),
      body: Center(
        child: Text("Home Page"),
      ),
      // floatingActionButton: FloatingActionButton(
      //   onPressed: () {
      //     setState(() {
      //       countController.add();
      //     });
      //   },
      //   child: Icon(Icons.add),
      // ),
    );
  }
}

// class CountWidget extends StatelessWidget {
//   final countController = Get.put(CounterController());

//   @override
//   Widget build(BuildContext context) {
//     return GetBuilder<CounterController>(
//       initState: (_) => print("initState"),
//       didChangeDependencies: (state) => print("didChangeDependencies"),
//       didUpdateWidget: (oldWidget, state) => print("didUpdate"),
//       dispose: (state) => print('dispose'),
//       builder: (c) => Text('Angka ${c.count}'),
//     );
//   }
// }

// class OtherPage extends StatelessWidget {
//   const OtherPage({super.key});

//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       appBar: AppBar(),
//     );
//   }
// }

class TextPage extends StatelessWidget {
  final textC = Get.put(TextController()); // Move this line into TextPageState

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Text Page"),
      ),
      body: TextField(
        controller: textC.textController,
      ),
    );
  }
}

class TextController extends GetxController {
  final textController = TextEditingController();
}

// class CounterController extends GetxController {
//   var count = 0;

//   void add() {
//     count++;
//     update();
//   }
// }
