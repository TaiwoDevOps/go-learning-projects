main() {
  final startTimer = Stopwatch()..start();
  print('The start timer is ${startTimer.elapsedMilliseconds} milliseconds');
  for (var i = 0; i < 1000000; i++) {
    print("The value is $i");
  }

  final stopTimer = Stopwatch()..stop();
  print('The stop timer is ${stopTimer.elapsedMilliseconds} milliseconds');

  print('The total time is ${startTimer.elapsedMilliseconds} milliseconds');
}

/***!SECTION
The value is 999999
The stop timer is 0 milliseconds
The total time is 2702 milliseconds*/