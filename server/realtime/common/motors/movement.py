from pymercure.consumer import Consumer
import json

class Movement:
    def __init__(self, logger):
        self._logger = logger
        self._consumer = Consumer('http://127.0.0.1:8181/hub', ['movement'], self._move)

    def _up(self):
        print('up')

    def _down(self):
        print('down')

    def _left(self):
        print('left')

    def _right(self):
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
