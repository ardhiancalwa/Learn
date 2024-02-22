import 'package:get/get.dart';

class CounterController extends GetxController {
  var counter = 0.obs;
  increament() => counter++;
  decreament() => counter--;

  var isDark = false.obs;

  void changeTheme() => isDark.value = !isDark.value;
}
