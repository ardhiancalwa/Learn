import 'package:flutter/material.dart';

void main() => runApp(const PhotoPage());

class PhotoPage extends StatelessWidget {
  static const nameRoute = '/photoPage';

  const PhotoPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Photo Page'),
      ),
      body: const Center(
        child: Text(
          'Photo Page',
          style: TextStyle(fontSize: 50),
        ),
      ),
    );
  }
}
