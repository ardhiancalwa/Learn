import 'package:get/get.dart';
import '../models/person.dart';

class PersonController extends GetxController {
  var person = Person(name: "Calwa", age: 18).obs;

  void changeUpperCase() {
    person.update((val) {
      person.value.name = person.value.name.toString().toUpperCase();
    });
  }
}
