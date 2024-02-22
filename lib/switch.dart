import 'package:flutter/material.dart';

void main() => runApp(const MyApp());

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  bool isSwitch = false;
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Swicth'),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Switch(
                activeColor: Colors.pink,
                inactiveTrackColor: Colors.black,
                activeThumbImage: AssetImage(
                  'assets/images/switch/cool.png',
                ),
                inactiveThumbImage: AssetImage(
                  'assets/images/switch/smile.png',
                ),
                value: isSwitch,
                onChanged: (value) {
                  setState(() {
                    isSwitch = !isSwitch;
                  });
                  print(isSwitch);
                },
              ),
              Text(
                (isSwitch == true) ? "Swicth on" : "Swicth off",
                style: TextStyle(
                  fontSize: 35,
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
