from pymercure.consumer import Consumer
from gpiozero import Robot
import json
import time

class Movement:
    def __init__(self, logger):
        self._logger = logger
        self._robot = Robot(left=(24, 23, 21), right=(20, 16, 18))
        self._consumer = Consumer('http://192.168.1.183:8181/hub', ['movement'], self._move)

    def _up(self):
        self._robot.forward()
        time.sleep(2)
        self._robot.stop()
        print('up')

    def _down(self):
        self._robot.backward()
        time.sleep(2)
        self._robot.stop()
        print('down')

    def _left(self):
        self._robot.left()
        time.sleep(2)
        self._robot.stop()
        print('left')

    def _right(self):
        self._robot.right()
        time.sleep(2)
        self._robot.stop()
        print('right')

    def _move(self, message):
        switcher = {
            'up': self._up,
            'down': self._down,
            'left': self._left,
            'right': self._right,
        }

        try:
            movement_data = json.loads(message.data)

            # Get the function from switcher dictionary
            movement = switcher.get(movement_data['direction'], lambda: "Invalid month")
            # Execute the function
            movement()
        except Exception as e:
            self._logger.error(str(e))

    def start(self):
        self._consumer.start_consumption()
        self._logger.info("Movemement connected")
