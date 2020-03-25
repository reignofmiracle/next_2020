import 'dart:math';
import 'package:flutter/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.indigo,
      ),
      home: MyHomePage(title: 'Flutter Demo Home Page?'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;
  bool _reversed = false;
  List<UniqueKey> _buttonKeys = [UniqueKey(), UniqueKey()];

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  void _decrementCounter() {
    setState(() {
      _counter--;
    });
  }

  void _resetCounter(){
    setState(() {
      _counter = 0;
    });
    _swap();
  }

  void _swap() {
    setState(() {
      _reversed = !_reversed;
    });
  }

  @override
  Widget build(BuildContext context) {
    final incrementButton = FancyButton(
      key: _buttonKeys.first,
      child: Text(
        "Increment",
        style: TextStyle(color: Colors.white),
      ),
      onPressed: _incrementCounter,
    );

    final decrementButton = FancyButton(
      key: _buttonKeys.last,
      child: Text(
        "Decrement",
        style: TextStyle(color: Colors.white),
      ),
      onPressed: _decrementCounter,
    );

    List<Widget> _buttons = <Widget>[incrementButton, decrementButton];

    if (_reversed) {
      _buttons = _buttons.reversed.toList();
    }

    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Container(
              margin: EdgeInsets.only(bottom: 100.0),
              padding: EdgeInsets.all(8.0),
              decoration: BoxDecoration(
                color: Colors.blue.withOpacity(0.25),
                borderRadius: BorderRadius.circular(4.0),
              ),
              child: Image.asset(
                'images/flutter_logo.png',
                width: 100.0,
              ),
            ),
            Text(
              'You have pushed the button this many times:',
            ),
            Text(
              '$_counter',
              style: Theme.of(context).textTheme.display1,
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: <Widget>[
                PanicButton(
                  display: Text("Decrement Counter"),
                  onPressed: _decrementCounter,
                ),
                RaisedButton(
                  color:Colors.green,
                  child:Text(
                    "Increment",
                    style: TextStyle(color: Colors.white),
                  ),
                  onPressed: _incrementCounter,
                ),
              ],
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: _buttons,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _resetCounter,
        tooltip: 'Reset Counter',
        child: Icon(Icons.refresh),
      ),
    );
  }
}

class PanicButton extends StatelessWidget {
  final Widget display;
  final VoidCallback onPressed;

  PanicButton({this.display, this.onPressed});

  Widget build(BuildContext context) {
    return RaisedButton(
      color: Colors.red,
      child: display,
      onPressed: onPressed,
    );
  }
}

class FancyButton extends StatefulWidget {
  final VoidCallback onPressed;
  final Widget child;

  const FancyButton({Key key, this.onPressed, this.child}) : super(key: key);

  @override
  _FancyButtonState createState() => _FancyButtonState();
}

class _FancyButtonState extends State<FancyButton> {
  @override
  Widget build(BuildContext context) {
    return Container(
      child: RaisedButton(
        color: _getColors(),
        child: widget.child,
        onPressed: widget.onPressed,
      ),
    );
  }

  Color _getColors() {
    return _buttonColors.putIfAbsent(this, () => colors[next(0, 5)]);
  }

  Map<_FancyButtonState, Color> _buttonColors = {};
  final _random = Random();
  int next(int min, int max) => min + _random.nextInt(max - min);
  List<Color> colors = [
    Colors.blue,
    Colors.green,
    Colors.orange,
    Colors.purple,
    Colors.amber,
    Colors.lightBlue,
  ];
}