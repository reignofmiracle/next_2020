class Animal {
  String name;
  int legCount;
}

class Cat extends Animal {
  String makeNoise() {
    print('sjflajklfjalsjf');
  }
}

void main() {
  Cat cat = Cat();
  cat.name = 'test';
  cat.legCount = 4;
  cat.makeNoise();
}

// class Energy {
//   int joules;

//   Energy(this.joules);

//   Energy.fromWind(int windBlows) {
//     final joules = _counvertWindToEnergy(windBlows);
//     return Energy(joules);
//   }

//   factory Energy.fromSolar(int sunbeams) {
//     if (appState.solarEnergy != null) return appState.solarEnergy;
//     final joules = _convertSunbeamsToJoules(sunbeams);
//     return appState.solarEnergy = Energy(joules)
//   }
// }