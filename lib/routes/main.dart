import 'package:flutter/material.dart';
import './pages/galleryPage.dart';
import './pages/photoPage.dart';
import './pages/homePage.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: HomePage(),
      initialRoute: HomePage.nameRoute,
      routes: {
        HomePage.nameRoute: (context) => const HomePage(),
        GalleryPage.nameRoute: (context) => const GalleryPage(),
        PhotoPage.nameRoute: (context) => const PhotoPage(),
      },
    );
  }
}
