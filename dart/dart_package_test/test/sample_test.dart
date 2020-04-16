import 'package:test/test.dart';

import 'package:dart_package_test/sample.dart';

void main() {
  test('String.split() splits the string on the delimiter', () {
    var string = 'foo,bar,baz';
    expect(string.split(','), equals(['foo', 'bar', 'baz']));
  });

  test('sample_lib.helloWorld() print name', () {
    var name = 'kekekekeke';
    helloWorld(name);
  });
}