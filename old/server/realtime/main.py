#!/usr/bin/env python3

import logging
from logging.handlers import RotatingFileHandler
from dotenv import load_dotenv
import os
from pathlib import Path
from common.motors.movement import Movement

class Configurator:
    def __init__(self):
        env_path = str(Path('.')) + '/.env'
        env_path_local = str(Path('.')) + '/.env.local'

        if os.path.exists(env_path_local):
            load_dotenv(dotenv_path=env_path_local)
        elif os.path.exists(env_path):
            load_dotenv(dotenv_path=env_path)


class Application:
    def __init__(self):
        Configurator()

        self.logger = self._init_logs()
        self._movement = Movement(self.logger)

    def _init_logs(self):
        logger = logging.getLogger()
        logger.setLevel(logging.INFO)
        formatter = logging.Formatter('%(asctime)s :: %(levelname)s :: %(message)s')

        file_handler = RotatingFileHandler(os.getenv('LOG_PATH'), 'a', 1000000, 1)
        file_handler.setLevel(logging.INFO)
        file_handler.setFormatter(formatter)
        logger.addHandler(file_handler)

        steam_handler = logging.StreamHandler()
        steam_handler.setLevel(logging.INFO)
        steam_handler.setFormatter(formatter)
        logger.addHandler(steam_handler)

        return logger

    def _start_movement(self):
        self._movement.start()

    def start(self):
        self.logger.info('Start realtime Roby management')
        self._start_movement()


def main():
    application = Application()
    application.start()


if __name__ == '__main__':
    main()

# c = Consumer('http://127.0.0.1:8181/hub', ['movement'], callback)
# c.start_consumption()
#
# data = json.dumps({'status': 'test'})
# msg = Message(['mytopicname'], data)
# publisher = SyncPublisher(
#      'http://127.0.0.1:8181/hub',
#      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtZXJjdXJlIjp7InB1Ymxpc2giOlsiKiJdfX0.AeCSJEGE8f_gdPLQBxgQlznmq_Mu071r_wly5gCLKug'
# )
#
# publisher.publish(msg)
